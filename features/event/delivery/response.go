package delivery

import "project/e-commerce/features/event"

type Respon struct {
	Activities []event.Core
}

func toRespon(data []event.Core) Respon {
	var res Respon

	res.Activities = data

	return res
}
