package event

type Core struct {
	ID           uint
	ActivityType string
	Institution  string
	When         string
	Objective    string
	Remarks      string
}

type DataInterface interface {
	SelectAll() ([]Core, error)
	SelectById(param int) (Core, error)
	UpdateData(Core) int
	CreateData(Core) int
}

type UsecaseInterface interface {
	GetAll() ([]Core, error)
	GetById(param int) (Core, error)
	PutData(Core) int
	PostData(Core) int
}
