package hashmaps

type DesignHashMap struct {
	keySpace int
	buckets  []*Bucket
}

// Create Empty HashMap
func Constructor() DesignHashMap {
	hashmap := DesignHashMap{}
	hashmap.keySpace = 2069
	hashmap.buckets = make([]*Bucket, hashmap.keySpace)
	for i := 0; i < hashmap.keySpace; i++ {
		hashmap.buckets[i] = NewBucket()
	}
	return hashmap
}

// Add key value to hashmap
func (m *DesignHashMap) Put(key, value int) {
	hashKey := key % m.keySpace
	m.buckets[hashKey].Update(key, value)
}

// Add key value to hashmap
func (m *DesignHashMap) Get(key int) int {
	hashKey := key % m.keySpace
	return m.buckets[hashKey].Get(key)
}

// Add key value to hashmap
func (m *DesignHashMap) Remove(key int) {
	hashKey := key % m.keySpace
	m.buckets[hashKey].Remove(key)
}
