package protocol

import (
	"i2pbote/log"
	"net"
	"sync"
	"time"
)

// a measurement of tx/rx for limiting
type LimitStat struct {
	Tx uint64
	Rx uint64
}

// message limiter
type Limiter struct {
	cmtx    sync.RWMutex
	convos  map[net.Addr]*LimitStat
	RxLimit uint64
	TxLimit uint64
	ticker  *time.Ticker
}

func (l *Limiter) Start() {
	if l.ticker != nil {
		log.Warn("Limiter ticker already started")
		return
	}
	// periodically reset tx/rx stats
	l.ticker = time.NewTicker(time.Second)
	go func() {
		for {
			_, ok := <-l.ticker.C
			if !ok {
				return
			}
			l.cmtx.Lock()
			for _, st := range l.convos {
				st.Rx = 0
				st.Tx = 0
			}
			l.cmtx.Unlock()
		}
	}()
}

func (l *Limiter) Stop() {
	if l.ticker != nil {
		l.ticker.Stop()
	}
}

// safely visit stat
func (l *Limiter) VisitStat(a net.Addr, v func(*LimitStat)) {
	l.cmtx.Lock()
	defer l.cmtx.Unlock()
	v(l.convos[a])
}

// return true if we should drop because we are sending to them too fast
func (l *Limiter) CheckTx(tx int, to net.Addr) bool {
	drop := false
	l.VisitStat(to, func(st *LimitStat) {
		if tx > 0 {
			st.Tx += uint64(tx)
			if l.TxLimit > 0 && st.Tx >= l.TxLimit {
				drop = true
			}
		}
	})
	return drop
}

// return true if we should drop because they are sending too fast to us
func (l *Limiter) CheckRX(rx int, from net.Addr) bool {
	drop := false
	l.VisitStat(from, func(st *LimitStat) {
		if rx > 0 {
			st.Rx += uint64(rx)
			if l.RxLimit > 0 && st.Rx >= l.RxLimit {
				drop = true
			}
		}
	})
	return drop
}

func NewLimiter() *Limiter {
	return &Limiter{
		convos: make(map[net.Addr]*LimitStat),
	}
}
