package listcourses

import (
	"fmt"
	"log"
	"net/http"

	"google.golang.org/api/classroom/v1"
)

// ListCourses returns a list of Google Classroom Courses with information such as id and name
func ListCourses(client *http.Client) {
	// https://godoc.org/google.golang.org/api/classroom/v1#New
	srv, err := classroom.New(client)
	if err != nil {
		log.Fatalf("Unable to create classroom Client %v", err)
	}

	/* I know it is Courses and not courses because srv made a new
			   type Service struct {
			client    *http.Client
			BasePath  string // API endpoint base URL
			UserAgent string // optional additional User-Agent fragment

			--> Courses *CoursesService

			Invitations *InvitationsService

			Registrations *RegistrationsService

			UserProfiles *UserProfilesService
		}

		List() then requires a variable of type *CoursesService

		func (r *CoursesService) List() *CoursesListCall {
		c := &CoursesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
		return c
	}
	*/

	//Notice that the Query Paramaters are added on as methods??
	courses, err := srv.Courses.List().PageSize(10).TeacherId("me").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve courses. %v", err)
	}

	if len(courses.Courses) > 0 {
		fmt.Print("Courses:\n")
		for _, c := range courses.Courses {
			fmt.Printf("%s (%s)\n", c.Name, c.Id)
		}
	} else {
		fmt.Print("No courses found.")
	}
}
