package main

import (
	"fmt"

	/* co "github.com/jdvober/goClassroomTools/courses" */
	co "github.com/jdvober/goClassroomTools/courses"
	"github.com/jdvober/goClassroomTools/students"
	stu "github.com/jdvober/goClassroomTools/students"
	auth "github.com/jdvober/goGoogleAuth"
	sh "github.com/jdvober/goSheets/values"
)

type Bio struct {
	First    string
	Last     string
	Email    string
	GoogleID string
	Class    string
}

func main() {
	client := auth.Authorize()

	spreadsheetId := "1HRfK4yZERLWd-OcDZ8pJRirdzdkHln3SUtIfyGZEjNk"
	rangeData := "Master!A2"
	/* values := [][]interface{}{{"sample_A1", "sample_B1"}, {"sample_A2", "sample_B2"}, {"sample_A3", "sample_A3"}} */

	courses := co.List(client)
	var studentProfiles []students.Profile

	for _, course := range courses {
		studentList := stu.List(client, course.Id) // CourseId Email Id First Last
		for _, student := range studentList {

			studentProfiles = append(studentProfiles, student)
		}
	}

	// Get all courses
	/* courses := co.List(client) */
	values := make([][]interface{}, len(studentProfiles))
	counter := 0
	for _, course := range courses {
		students := stu.List(client, course.Id)
		for _, s := range students {
			fmt.Println(s.First, s.Last, s.CourseId)
			gradeLevel := switchGradeLevel(course.Name)
			values[counter] = []interface{}{s.First, s.Last, s.Email, course.Name, gradeLevel, s.Id, course.Id}

			counter++
		}
	}
	// Must convert students (type []students.Profile) to type []interface {}

	/* for i, v := range studentProfiles {
	 *     values[i] = []interface{}{v.First, v.Last, v.Email}
	 * } */

	sh.BatchUpdate(client, spreadsheetId, rangeData, values)
	fmt.Println("Finished main()")
}

func switchGradeLevel(name string) string {

	switch name {
	case "AP Physics", "Physics", "APEX Physics":
		return "12"
	case "APEX Honors Chemistry", "APEX Chemistry":
		return "11"
	case "APEX Physical Science":
		return "9"
	default:
		return ""
	}

}
