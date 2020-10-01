package listcoursework

import (
	"fmt"
	"log"
	"net/http"

	"google.golang.org/api/classroom/v1"
)

// ListCourseWork returns a list of all the coursework in a course with a specific Id
func ListCourseWork(client *http.Client, id string) {
	srv, err := classroom.New(client)
	if err != nil {
		log.Fatalf("Unable to create classroom Client %v", err)
	}
	// My attempt to get courseWork using courseId 126909787383
	coursework, err := srv.Courses.CourseWork.List(id).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve coursework. %v", err)
	}

	if len(coursework.CourseWork) > 0 {
		fmt.Print("\nCoursework:\n")
		for _, cw := range coursework.CourseWork {
			fmt.Printf("%s\n", cw.Id)
		}
	} else {
		fmt.Print("No coursework found.")
	}
}

/*JSON Response Look Like This:

{
"courseWork": [
{
	"courseId": "126909787383",
	"id": "178503269396",
	"title": "7.4.4 - Discuss",
	"description": "Please take part in a discussion addressing the questions posed in 7.4.4 on APEX, ...etc",
	"state": "PUBLISHED",
	"alternateLink": "https://classroom.google.com/c/MTI2OTA5Nzg3Mzgz/sa/MTc4NTAzMjY5Mzk2/details",
	"creationTime": "2020-09-29T14:05:12.405Z",
	"updateTime": "2020-09-29T14:06:59.230Z",
	"dueDate": {
	"year": 2020,
	"month": 10,
	"day": 16
	},
	"dueTime": {
	"hours": 3,
	"minutes": 59
	},
	"maxPoints": 15,
	"workType": "SHORT_ANSWER_QUESTION",
	"submissionModificationMode": "MODIFIABLE",
	"assigneeMode": "ALL_STUDENTS",
	"creatorUserId": "118026622292638242171",
	"topicId": "178468686835"
},

*/
