package trylock

type Locker struct {
	c chan struct{}
}

func New() *Locker {
	return &Locker{c: make(chan struct{})}
}
