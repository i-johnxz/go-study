package scheduler

import "../model"

type ReadyNotifier interface {
	WorkerReady(chan model.Request)
}
