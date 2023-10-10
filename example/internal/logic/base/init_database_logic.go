package base

import (
	"context"
	"entgo.io/ent/dialect/sql/schema"
	"fmt"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/msg/logmsg"
	"github.com/suyuan32/simple-admin-example-api/ent"
	"github.com/suyuan32/simple-admin-example-api/ent/course"
	"github.com/suyuan32/simple-admin-example-api/ent/exam"
	"github.com/suyuan32/simple-admin-example-api/ent/student"
	"github.com/suyuan32/simple-admin-example-api/internal/svc"
	"github.com/suyuan32/simple-admin-example-api/internal/types"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
)

type InitDatabaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InitDatabaseLogic) InitDatabase() (resp *types.BaseMsgResp, err error) {
	if err := l.svcCtx.DB.Schema.Create(l.ctx, schema.WithForeignKeys(true), schema.WithDropColumn(true),
		schema.WithDropIndex(true)); err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		return nil, errorx.NewCodeInternalError(err.Error())
	}

	//// 获取所有学生信息 | Get All Student Data
	studentData, _ := l.svcCtx.DB.Student.Query().All(l.ctx)
	fmt.Println("All Students' data: ", studentData)

	// 获取所有学生名字和年龄信息 | Get Student's name and age  data
	studentPartialData, _ := l.svcCtx.DB.Student.Query().Select(student.FieldName, student.FieldAge).All(l.ctx)
	fmt.Println("All Students' name and age data: ", studentPartialData)

	// 给 Jack 添加一门数学课
	courseMath, err := l.svcCtx.DB.Course.Query().Where(course.NameEQ("Math")).First(l.ctx)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.DB.Student.Update().Where(student.NameEQ("Jack")).AddCourses(courseMath).Exec(l.ctx)
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.DB.Student.Update().Where(student.NameEQ("Mike")).AddCourseIDs(2).Exec(l.ctx)
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.DB.Exam.Create().SetCoursesID(2).SetStudentsID(2).SetScore(65).Exec(l.ctx)
	if err != nil {
		return nil, err
	}

	// 获取Jack的信息 | Get Jack Information
	jackData, _ := l.svcCtx.DB.Student.Query().Where(student.NameEQ("Jack")).First(l.ctx)
	fmt.Println("Jack's Data: ", jackData)

	// 获取 Jack 上的课 | Get Jack's courses
	jackCourseData, _ := l.svcCtx.DB.Student.Query().Where(student.NameEQ("Jack")).WithCourses().First(l.ctx)
	fmt.Println("Jack's Course Data: ", jackCourseData.Edges.Courses)

	// 获取选艺术课的学生 | Get students who choose Art course
	artCourse, _ := l.svcCtx.DB.Course.Query().Where(course.NameEQ("Art")).WithStudents().First(l.ctx)
	fmt.Println("Students Choosing Art course: ", artCourse.Edges.Students)

	// 获取选艺术课的学生 | Get students who choose Art course
	artStudent, _ := l.svcCtx.DB.Student.Query().Where(student.HasCoursesWith(course.NameEQ("Art"))).All(l.ctx)
	fmt.Println("Students Choosing Art course: ", artStudent)

	// 艺术课的总分 | Art course score sum
	sumArtCourse, _ := l.svcCtx.DB.Exam.Query().Where(
		exam.HasCoursesWith(course.NameEQ("Art"))).Aggregate(ent.Sum(exam.FieldScore)).Int(l.ctx)
	fmt.Println("Art course score sum: ", sumArtCourse)

	// 聚合函数的例子 | Aggregation Example
	var examAggreData []struct {
		Count     int `json:"count"`
		Sum       int `json:"sum"`
		StudentID int `json:"student_exams"`
	}
	_ = l.svcCtx.DB.Exam.Query().GroupBy(exam.StudentsColumn).Aggregate(ent.Count(),
		ent.Sum(exam.FieldScore)).Scan(l.ctx, &examAggreData)
	fmt.Printf("Exam aggregation: %+v", examAggreData)

	// 纯 sql 例子 | Raw sql example
	rows, _ := l.svcCtx.DB.QueryContext(l.ctx, "select name from students;")
	defer rows.Close()
	names := make([]string, 0)

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Fatal(err)
		}
		names = append(names, name)
	}
	fmt.Println("Raw data: ", names)

	page, err := l.svcCtx.DB.Student.Query().Page(l.ctx, 1, 10, func(pager *ent.StudentPager) {
		pager.Order = ent.Desc(student.FieldID)
	})
	if err != nil {
		return nil, err
	}

	fmt.Println(page.List)

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.Success)}, nil
}
