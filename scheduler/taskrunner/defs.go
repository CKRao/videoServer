package taskrunner

const (
	READY_TO_DISPATCH = "d"
	READY_TO_EXECUTE  = "e"
	CLOSE             = "c"
	VIDEO_PATH        = "C:/Users/clarkrao/go/src/videos/"
)

type controllerChan chan string

type dataChan chan interface{}

type fn func(dc dataChan) error
