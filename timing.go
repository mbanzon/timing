package timing

import (
	"net/http"
	"sync/atomic"
	"time"
)

type Timing struct {
	threshold time.Duration
	good      uint64
	bad       uint64
}

func New(threshold time.Duration) *Timing {
	return &Timing{
		threshold: threshold,
	}
}

func (t *Timing) Stats() (good, bad uint64) {
	return atomic.LoadUint64(&t.good), atomic.LoadUint64(&t.bad)
}

func (t *Timing) Reset() {
	atomic.StoreUint64(&t.good, 0)
	atomic.StoreUint64(&t.bad, 0)
}

func (t *Timing) Wrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		elapsed := time.Since(start)
		if elapsed > t.threshold {
			atomic.AddUint64(&t.bad, 1)
		} else {
			atomic.AddUint64(&t.good, 1)
		}
	})
}
