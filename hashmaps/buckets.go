package hashmaps

type Bucket struct {
	bucket []Pair
}

type Pair struct {
	Key   int
	Value int
}

// Create an empty bucket
func NewBucket() *Bucket {
	return &Bucket{
		bucket: make([]Pair, 0),
	}
}

// Get key value from bucket
func (b *Bucket) Get(key int) int {
	for _, pair := range b.bucket {
		if pair.Key == key {
			return pair.Value
		}
	}
	return -1
}

// Update key in bucket
func (b *Bucket) Update(key, value int) {
	found := false
	for i, pair := range b.bucket {
		if pair.Key == key {
			b.bucket[i].Value = value
			found = true
			break
		}
	}

	if !found {
		b.bucket = append(b.bucket, Pair{key, value})
	}
}

// Remove key in bucket
func (b *Bucket) Remove(key int) {
	for i, pair := range b.bucket {
		if pair.Key == key {
			b.bucket = append(b.bucket[:i], b.bucket[i+1:]...)
			break
		}
	}
}
