package mutex

type Mutex struct {
	Count int
	done  chan struct{}
}

func (m *Mutex) Unlock() {
	m.done <- struct{}{}
}

func (m *Mutex) Wait() {
	for i := 0; i < m.Count; i++ {
		<-m.done
	}
}
