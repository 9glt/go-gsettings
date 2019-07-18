package gsettings

import (
	"errors"
	"sync"
)

var (
	mu = &sync.RWMutex{}
	s  = make(map[string]string)

	ErrNotFound = errors.New("Key not found")
)

func Reset() {
	s = make(map[string]string)
}

func Set(key, value string) {
	mu.Lock()
	s[key] = value
	mu.Unlock()
}

func Get(key string) (string, error) {
	mu.RLock()
	value, exists := s[key]
	mu.RUnlock()
	if !exists {
		return "", ErrNotFound
	}
	return value, nil
}
