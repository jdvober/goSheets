package main

import (
	auth "github.com/jdvober/goGoogleAuth"
	//courselist "github.com/jdvober/goClassroomTools"
	courselist "github.com/jdvober/goClassroomTools/listCourses"
	//courseworklist "github.com/jdvober/goClassroomTools/listCourseWork"
	courseworklist "github.com/jdvober/goClassroomTools/listCourseWork"
)

func main() {
	client := auth.GetClient()

	courselist.ListCourses(client)
	courseworklist.ListCourseWork(client, "126909787383")
	// fmt.Println("\n<<Press any key to exit...>>")
	// fmt.Scanln()

}
