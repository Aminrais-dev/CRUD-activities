package delivery

import "project/e-commerce/features/event"

type Request struct {
	ActivityType string `json:"activity type" form:"activity type"`
	Institution  string `json:"institution" form:"institution"`
	When         string `json:"when" form:"when"`
	Objective    string `json:"objective" form:"objective"`
	Remarks      string `json:"remaks" form:"remaks"`
}

func (req *Request) toCoreAct() event.Core {
	activitiesCore := event.Core{
		ActivityType: req.ActivityType,
		Institution:  req.Institution,
		When:         req.When,
		Objective:    req.Objective,
		Remarks:      req.Remarks,
	}
	return activitiesCore
}
