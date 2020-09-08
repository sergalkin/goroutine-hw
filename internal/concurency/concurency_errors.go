package concurency

import "sync"

type errorsCnt struct {
	sync.RWMutex
	errorCounter   int
	errorThreshold int
}

func (ec *errorsCnt) Inc() {
	ec.Lock()
	ec.errorCounter++
	ec.Unlock()
}

func (ec *errorsCnt) Get() int {
	ec.RLock()
	defer ec.RUnlock()
	return ec.errorCounter
}

func (ec *errorsCnt) isThresholdReached() bool {
	ec.Lock()
	defer ec.Unlock()
	return ec.errorCounter == ec.errorThreshold
}