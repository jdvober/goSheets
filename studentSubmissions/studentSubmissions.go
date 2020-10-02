package studentsubmissions

import (
	"log"
	"net/http"

	"google.golang.org/api/classroom/v1"
)

// List returns a list of all the students in a course with a specific Id
func List(client *http.Client, id string, cwid string) []*classroom.StudentSubmission {
	srv, err := classroom.New(client)
	if err != nil {
		log.Fatalf("Unable to create classroom Client %v", err)
	}

	courseID := id
	courseWorkID := cwid
	// pToken := ""

	submissions := []*classroom.StudentSubmission{}

	// My attempt to get studentSubmissions using courseId 126909787383 and courseWorkId "-" which returns all studentSubmissions

	// https://godoc.org/google.golang.org/api/classroom/v1#CoursesCourseWorkStudentSubmissionsService
	res, err := srv.Courses.CourseWork.StudentSubmissions.List(courseID, courseWorkID).PageSize(2).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve submission. %v", err)
	}
	for _, s := range res.StudentSubmissions {
		submissions = append(submissions, s)
	}

	return submissions
}

// Submission is generated with https://mholt.github.io/json-to-go/

// StudentSubmission is used to convert from JSON resonse of call to struct
// type StudentSubmission struct {
// 	CourseID             string    `json:"courseId"`
// 	CourseWorkID         string    `json:"courseWorkId"`
// 	ID                   string    `json:"id"`
// 	UserID               string    `json:"userId"`
// 	CreationTime         time.Time `json:"creationTime"`
// 	UpdateTime           time.Time `json:"updateTime"`
// 	State                string    `json:"state"`
// 	DraftGrade           int       `json:"draftGrade"`
// 	AssignedGrade        int       `json:"assignedGrade"`
// 	AlternateLink        string    `json:"alternateLink"`
// 	CourseWorkType       string    `json:"courseWorkType"`
// 	AssignmentSubmission struct {
// 		Attachments []struct {
// 			DriveFile struct {
// 				ID            string `json:"id"`
// 				Title         string `json:"title"`
// 				AlternateLink string `json:"alternateLink"`
// 				ThumbnailURL  string `json:"thumbnailUrl"`
// 			} `json:"driveFile"`
// 		} `json:"attachments"`
// 	} `json:"assignmentSubmission,omitempty"`
// 	SubmissionHistory []struct {
// 		StateHistory struct {
// 			State          string    `json:"state"`
// 			StateTimestamp time.Time `json:"stateTimestamp"`
// 			ActorUserID    string    `json:"actorUserId"`
// 		} `json:"stateHistory,omitempty"`
// 		GradeHistory struct {
// 			PointsEarned    int       `json:"pointsEarned"`
// 			MaxPoints       int       `json:"maxPoints"`
// 			GradeTimestamp  time.Time `json:"gradeTimestamp"`
// 			ActorUserID     string    `json:"actorUserId"`
// 			GradeChangeType string    `json:"gradeChangeType"`
// 		} `json:"gradeHistory,omitempty"`
// 	} `json:"submissionHistory"`
// 	ShortAnswerSubmission struct {
// 		Answer string `json:"answer"`
// 	} `json:"shortAnswerSubmission,omitempty"`
// }

// Response is a response from Google Classroom
// type Response struct {
// 	StudentSubmissions []struct {
// 		CourseID             string    `json:"courseId"`
// 		CourseWorkID         string    `json:"courseWorkId"`
// 		ID                   string    `json:"id"`
// 		UserID               string    `json:"userId"`
// 		CreationTime         time.Time `json:"creationTime"`
// 		UpdateTime           time.Time `json:"updateTime"`
// 		State                string    `json:"state"`
// 		DraftGrade           int       `json:"draftGrade"`
// 		AssignedGrade        int       `json:"assignedGrade"`
// 		AlternateLink        string    `json:"alternateLink"`
// 		CourseWorkType       string    `json:"courseWorkType"`
// 		AssignmentSubmission struct {
// 			Attachments []struct {
// 				DriveFile struct {
// 					ID            string `json:"id"`
// 					Title         string `json:"title"`
// 					AlternateLink string `json:"alternateLink"`
// 					ThumbnailURL  string `json:"thumbnailUrl"`
// 				} `json:"driveFile"`
// 			} `json:"attachments"`
// 		} `json:"assignmentSubmission,omitempty"`
// 		SubmissionHistory []struct {
// 			StateHistory struct {
// 				State          string    `json:"state"`
// 				StateTimestamp time.Time `json:"stateTimestamp"`
// 				ActorUserID    string    `json:"actorUserId"`
// 			} `json:"stateHistory,omitempty"`
// 			GradeHistory struct {
// 				PointsEarned    int       `json:"pointsEarned"`
// 				MaxPoints       int       `json:"maxPoints"`
// 				GradeTimestamp  time.Time `json:"gradeTimestamp"`
// 				ActorUserID     string    `json:"actorUserId"`
// 				GradeChangeType string    `json:"gradeChangeType"`
// 			} `json:"gradeHistory,omitempty"`
// 		} `json:"submissionHistory"`
// 		ShortAnswerSubmission struct {
// 			Answer string `json:"answer"`
// 		} `json:"shortAnswerSubmission,omitempty"`
// 	} `json:"studentSubmissions"`
// 	NextPageToken string `json:"nextPageToken"`
// }
