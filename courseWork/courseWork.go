package coursework

import (
	"fmt"
	"log"
	"net/http"

	"google.golang.org/api/classroom/v1"
)

type date struct {
	year  int64
	month int64
	day   int64
}
type timeOfDay struct {
	hours   int64
	minutes int64
}

type coursework struct {
	courseID      string
	id            string
	title         string
	description   string
	state         string
	creationTime  string
	updateTime    string
	dueDate       date
	dueTime       timeOfDay
	scheduledTime string
	maxPoints     float64
	topicID       string
}

// List returns a list of all the coursework in a course with a specific Id
func List(client *http.Client, id string) []coursework {
	srv, err := classroom.New(client)
	if err != nil {
		log.Fatalf("Unable to create classroom Client %v", err)
	}
	// My attempt to get courseWork using courseId 126909787383
	r, err := srv.Courses.CourseWork.List(id).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve coursework. %v", err)
	}
	courseworkList := []coursework{}
	if len(r.CourseWork) > 0 {
		//fmt.Print("\nCoursework:\n")
		for _, cw := range r.CourseWork {

			cwork := coursework{

				courseID:      cw.CourseId,
				id:            cw.Id,
				title:         cw.Title,
				description:   cw.Description,
				state:         cw.State,
				creationTime:  cw.CreationTime,
				updateTime:    cw.UpdateTime,
				dueDate:       date{year: cw.DueDate.Year, month: cw.DueDate.Month, day: cw.DueDate.Day},
				dueTime:       timeOfDay{hours: cw.DueTime.Hours, minutes: cw.DueTime.Minutes},
				scheduledTime: cw.ScheduledTime,
				maxPoints:     cw.MaxPoints,
				topicID:       cw.TopicId,
			}
			courseworkList = append(courseworkList, cwork)
		}

	} else {
		fmt.Print("No coursework found.")
	}
	return courseworkList
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
