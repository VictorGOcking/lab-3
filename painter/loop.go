package painter

import (
	"image"
	"sync"

	"golang.org/x/exp/shiny/screen"
)

// Receiver отримує текстуру, яка була підготовлена в результаті виконання команд у циклі подій.
type Receiver interface {
	Update(t screen.Texture)
}

// Loop реалізує цикл подій для формування текстури отриманої через виконання операцій отриманих з внутрішньої черги.
type Loop struct {
	Receiver Receiver

	next screen.Texture // Текстура, яка зараз формується
	prev screen.Texture // Текстура, яка була відправлення останнього разу у Receiver

	mq messageQueue

	stop    chan struct{}
	stopReq bool
}

var size = image.Pt(800, 800)

// Start запускає цикл подій. Цей метод потрібно запустити до того, як викликати на ньому будь-які інші методи.
func (l *Loop) Start(s screen.Screen) {
	l.next, _ = s.NewTexture(size)
	l.prev, _ = s.NewTexture(size)

	l.stop = make(chan struct{})

	go func() {
		for !(l.stopReq && l.mq.empty()) {
			op := l.mq.pull()

			// Значення true повертається лише тоді, коли
			// Операція хоче перемалювати вікно після свого виконання
			update := op.Do(l.next)
			if update {
				l.Receiver.Update(l.next)
				l.next, l.prev = l.prev, l.next
			}
		}
		close(l.stop)
	}()
}

// Post додає нову операцію у внутрішню чергу.
func (l *Loop) Post(op Operation) {
	l.mq.push(op)
}

// StopAndWait сигналізує про необхідність завершити цикл та блокується до моменту його повної зупинки.
func (l *Loop) StopAndWait() {
	l.Post(OperationFunc(func(texture screen.Texture) {
		l.stopReq = true
	}))
	// Виконання заблокується поки хтось не напише щось в канал
	// Або він не закриється
	<-l.stop
}

// Черга подій.
type messageQueue struct {
	messages []Operation
	mu       sync.Mutex

	signal chan struct{}
}

func (mq *messageQueue) push(op Operation) {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	mq.messages = append(mq.messages, op)

	if mq.signal != nil {
		close(mq.signal)
		mq.signal = nil
	}
}

func (mq *messageQueue) pull() Operation {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	for len(mq.messages) == 0 {
		mq.signal = make(chan struct{})
		mq.mu.Unlock()
		<-mq.signal
		//mq.signal = nil
		mq.mu.Lock()
	}

	res := mq.messages[0]
	mq.messages[0] = nil
	mq.messages = mq.messages[1:]

	return res
}

func (mq *messageQueue) empty() bool {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	return len(mq.messages) == 0
}
