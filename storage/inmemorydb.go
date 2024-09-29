package storage

import "sync"

type InMemoryDB struct {
	data map[string][]byte
	lck sync.RWMutex
}

func NewInMemoryDB() DB {
	return &InMemoryDB{ 
		data: make(map[string][]byte)
	}
}

func (d *InMemoryDB) Get(key string) ([]byte, error) {
	d.lck.RLock()
	defer d.lck.RUnlock()
	val, ok := d.data[key]
	if !ok {
		return nil, ErrNotFound
	}
	return val, nil
}

func (d *InMemoryDB) PutKey(key string, val []byte) error {
	d.lck.WLock()
	defer d.lck.WUnlock()
	d.data[key] = val
	return nil
}



