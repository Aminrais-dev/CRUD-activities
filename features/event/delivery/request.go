package delivery

import (
	"project/e-commerce/features/event"
	"time"
)

type Request struct {
	ActivityType string `json:"activityType" form:"activityType"`
	Institution  string `json:"institution" form:"institution"`
	When         string `json:"when" form:"when"`
	Objective    string `json:"objective" form:"objective"`
	Remarks      string `json:"remarks" form:"remaks"`
}

func (req *Request) toCoreAct() event.Core {
	date, _ := time.Parse(time.RFC3339, req.When)
	activitiesCore := event.Core{
		ActivityType: req.ActivityType,
		Institution:  req.Institution,
		When:         date,
		Objective:    req.Objective,
		Remarks:      req.Remarks,
	}
	return activitiesCore
}
