package mutex

type Mutex struct {
	Count int
	done  chan struct{}
}

func New(count int) *Mutex {
	return &Mutex{
		Count: count,
		done:  make(chan struct{}, count),
	}
}

func (m *Mutex) Unlock() {
	m.done <- struct{}{}
}

func (m *Mutex) Wait() {
	for i := 0; i < m.Count; i++ {
		<-m.done
	}
}
