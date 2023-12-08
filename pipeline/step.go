package pipeline

import (
	"sync"
	"sync/atomic"

	"github.com/jeessy2/s3sync/storage"
)

// StepFn implement the type of pipeline Step function.
type StepFn func(group *Group, stepNum int, input <-chan *storage.Object, output chan<- *storage.Object, errChan chan<- error)

// Step contain configuration of pipeline step and it's internal structure.
// Be careful with Config interface! Check of its type should implemented in StepFn.
// If typing fails, you get a StepConfigurationError in runtime.
type Step struct {
	Name       string
	Fn         StepFn
	AddWorkers uint
	Config     interface{}
	ChanSize   uint
	outChan    chan *storage.Object
	intOutChan chan *storage.Object
	intInChan  chan *storage.Object
	errChan    chan error
	workerWg   *sync.WaitGroup
	stats      StepStats
}

// StepStats to keep basic step statistics.
type StepStats struct {
	Input  atomic.Uint64
	Output atomic.Uint64
	Error  atomic.Uint64
}

// StepInfo is used to represent step information, statistic and the step configuration interface.
type StepInfo struct {
	Stats  *StepStats
	Name   string
	Num    int
	Config interface{}
}
