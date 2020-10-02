package students

import (
	"fmt"
	"log"
	"net/http"

	"google.golang.org/api/classroom/v1"
)

type profile struct {
	first    string
	last     string
	id       string
	email    string
	courseId string
}

// List returns a list of all the students in a course with a specific Id
func List(client *http.Client, id string) []profile {
	srv, err := classroom.New(client)
	if err != nil {
		log.Fatalf("Unable to create classroom Client %v", err)
	}
	// My attempt to get courseWork using courseId 126909787383
	students, err := srv.Courses.Students.List(id).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve students. %v", err)
	}
	roster := []profile{}
	if len(students.Students) > 0 {
		//fmt.Print("\nStudents:\n")

		for _, s := range students.Students {
			studentProfile := profile{
				first:    s.Profile.Name.GivenName,
				last:     s.Profile.Name.FamilyName,
				id:       s.Profile.Id,
				email:    s.Profile.EmailAddress,
				courseId: s.CourseId,
			}
			roster = append(roster, studentProfile)
		}
	} else {
		fmt.Print("No students found.")
	}
	return roster
}
