package data

import (
	"project/e-commerce/features/event"
	"time"

	"gorm.io/gorm"
)

type Activities struct {
	gorm.Model
	ActivityType string
	Institution  string
	When         time.Time
	Objective    string
	Remarks      string
}

func (act *Activities) toCore() event.Core {
	return event.Core{
		ID:           act.ID,
		ActivityType: act.ActivityType,
		Institution:  act.Institution,
		When:         act.When,
		Objective:    act.Objective,
		Remarks:      act.Remarks,
	}
}

func toCoreList(data []Activities) []event.Core {
	var dataCore []event.Core
	for key := range data {
		dataCore = append(dataCore, data[key].toCore())
	}
	return dataCore
}

func fromCore(dataCore event.Core) Activities {
	activitiesModel := Activities{
		ActivityType: dataCore.ActivityType,
		Institution:  dataCore.Institution,
		When:         dataCore.When,
		Objective:    dataCore.Objective,
		Remarks:      dataCore.Remarks,
	}
	return activitiesModel
}
