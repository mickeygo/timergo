package repos

import (
	"sync"
)

var (
	mux                 sync.Mutex
	DefaultCacheSeconds int64 = 3600
)
