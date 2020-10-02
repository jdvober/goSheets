package main

import (
	"fmt"

	spew "github.com/davecgh/go-spew/spew"
	auth "github.com/jdvober/goGoogleAuth"

	subs "github.com/jdvober/goClassroomTools/studentSubmissions"
	/*
		co "github.com/jdvober/goClassroomTools/courses"
		cw "github.com/jdvober/goClassroomTools/courseWork"
		st "github.com/jdvober/goClassroomTools/students"
	*/)

func main() {
	client := auth.GetClient()

	submissions := subs.List(client, "126909787383", "-")

	// courses := co.List(client)
	// courseWork := cw.List(client, "126909787383")
	// students := st.List(client, "126909787383")

	// fmt.Println("\nCourses (Course ID): >>")
	// for cl := range courseList {
	// 	spew.Dump(courses[cl])
	// }

	// fmt.Println("\nCoursework: >>")
	// for c := range courseWork {
	// 	spew.Dump(courseWork[c])
	// }

	// fmt.Println("\nStudents: >>")
	// for s := range students {
	// 	spew.Dump(students[s])
	// }

	fmt.Println("\nSubmissions: >>")
	for s := range submissions {
		// fmt.Printf("\nSubmission %d\n%+v\n", s, submissions[s])
		spew.Dump(submissions[s])
	}

	// fmt.Println("\n<<Press any key to exit...>>")
	// fmt.Scanln()

}
