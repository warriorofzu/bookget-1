package curl

import (
	"context"
	"net/http"
	"sync/atomic"
	"time"
)

// Download holds downloadable file config and infos.
type Download struct {
	Client                    *http.Client
	URL, Dir, Dest            string
	Interval                  uint64
	Cookie                    []http.Cookie
	StopProgress              bool
	ctx                       context.Context
	totalSize, size, lastSize uint64
	startedAt                 time.Time
}

// TotalSize returns file total size (0 if unknown).
func (d *Download) TotalSize() uint64 {
	return d.totalSize
}

// Size returns downloaded size.
func (d *Download) Size() uint64 {
	return atomic.LoadUint64(&d.size)
}

// Speed returns download speed.
func (d *Download) Speed() uint64 {
	return (atomic.LoadUint64(&d.size) - atomic.LoadUint64(&d.lastSize)) / d.Interval * 1000
}

// AvgSpeed returns average download speed.
func (d *Download) AvgSpeed() uint64 {

	if totalMills := d.TotalCost().Milliseconds(); totalMills > 0 {
		return uint64(atomic.LoadUint64(&d.size) / uint64(totalMills) * 1000)
	}

	return 0
}

// TotalCost returns download duration.
func (d *Download) TotalCost() time.Duration {
	return time.Now().Sub(d.startedAt)
}

// Write updates progress size.
func (d *Download) Write(b []byte) (int, error) {
	n := len(b)
	atomic.AddUint64(&d.size, uint64(n))
	return n, nil
}
