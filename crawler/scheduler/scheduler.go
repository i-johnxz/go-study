package scheduler

import "../model"

type Scheduler interface {
	Submit(request model.Request)
	ConfigureMasterWorkerChan(chan model.Request)
}
