package event

import "time"

type Core struct {
	ID           uint      `json:"id"`
	ActivityType string    `json:"activityType"`
	Institution  string    `json:"institution"`
	When         time.Time `json:"when"`
	Objective    string    `json:"objective"`
	Remarks      string    `json:"remarks"`
}

type DataInterface interface {
	SelectAll() ([]Core, error)
	SelectById(param int) (Core, error)
	UpdateData(Core) (Core, error)
	CreateData(Core) (Core, error)
}

type UsecaseInterface interface {
	GetAll() ([]Core, error)
	GetById(param int) (Core, error)
	PutData(Core) (Core, error)
	PostData(Core) (Core, error)
}
