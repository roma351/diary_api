package resources

import (
	"github.com/roma351/diary_api/objects"
)

type RatingClass struct {
	Items  []objects.RatingCard `json:"items"`
	Class  objects.Class        `json:"class"`
	School objects.School       `json:"school"`
	*ExtraResponse
}

type RatingClasses struct {
	Items  []objects.RatingCard `json:"items"`
	School objects.School       `json:"school"`
	*ExtraResponse
}

type RatingLesson struct {
	Items   []objects.RatingCard `json:"items"`
	Subject objects.Subject      `json:"subject"`
	School  objects.School       `json:"school"`
	Day     objects.DiaryDayInfo `json:"day"`
	*ExtraResponse
}

type RatingSubjects struct {
	Subjects []objects.RatingSubject `json:"subjects"`
	Class    objects.Class           `json:"class"`
	School   objects.School          `json:"school"`
	Period   objects.Period          `json:"period"`
	*ExtraResponse
}
