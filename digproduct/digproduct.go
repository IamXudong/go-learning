package digproduct

import (
	"io"
	"time"
)

type Artifact interface {
	Title() string
	Creators() []string
	Created() time.Time
}

type Text interface {
	Pages() int
	Words() int
	Pagesize() int
}

type Streamer interface {
	Stream() (io.ReaderCloser, error)
	RuningTime() time.Duration
	Format() string
}

type Audio interface {
	Streamer
}

type Video interface {
	Streamer
	Resolution() (x, y int)
}
