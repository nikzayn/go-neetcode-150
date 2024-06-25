package hashmaps

import (
	"testing"
)

func TestHashMap(t *testing.T) {
	var m DesignHashMap
	m = Constructor()
	m.Put(1, 100)
	m.Put(2, 200)
	m.Put(2, 400)
	m.Remove(2)
	m.Get(4)
	m.Get(1)
}
