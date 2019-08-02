package scheduler

import "../model"

type Scheduler interface {
	ReadyNotifier
	Submit(request model.Request)
	WorkerChann() chan model.Request
	Run()
}
