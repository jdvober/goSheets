module goClassroomTools

go 1.15

//Comment this line out when using for production
replace github.com/jdvober/goGoogleAuth => ../goGoogleAuth

replace github.com/jdvober/goClassroomTools/listCourses => ../goClassroomTools/listCourses

replace github.com/jdvober/goClassroomTools/listCourseWork => ../goClassroomTools/listCourseWork

require (
	github.com/jdvober/goClassroomTools/listCourseWork v0.0.0-00010101000000-000000000000
	github.com/jdvober/goClassroomTools/listCourses v0.0.0-00010101000000-000000000000
	github.com/jdvober/goGoogleAuth v0.0.0-00010101000000-000000000000
	google.golang.org/api v0.32.0
)
