package instrumentation

type Segment interface {
	End() error
}
