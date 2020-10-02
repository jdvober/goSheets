package main

import (
	"fmt"

	auth "github.com/jdvober/goGoogleAuth"
	//courselist "github.com/jdvober/goClassroomTools"
	course "github.com/jdvober/goClassroomTools/courses"
	//courseworklist "github.com/jdvober/goClassroomTools/listCourseWork"
	coursework "github.com/jdvober/goClassroomTools/courseWork"
	//studentlist "github.com/jdvober/goClassroomTools/listStudents"
	students "github.com/jdvober/goClassroomTools/students"
)

func main() {
	client := auth.GetClient()

	courseList := course.List(client)
	courseWork := coursework.List(client, "126909787383")
	studentList := students.List(client, "126909787383")

	fmt.Println("\nCourses (Course ID): >>")
	for cl := range courseList {
		fmt.Println(courseList[cl])
	}

	fmt.Println("\nCoursework: >>")
	for c := range courseWork {
		fmt.Println(courseWork[c])
	}

	fmt.Println("\nStudents: >>")
	for s := range studentList {
		fmt.Println(studentList[s])
	}
	// fmt.Println("\n<<Press any key to exit...>>")
	// fmt.Scanln()

}
