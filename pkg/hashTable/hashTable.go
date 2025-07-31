// Package hashtable implements a simple hash table with basic operations.
package hashtable

type HashTable struct {
	buckets [][]KeyValuePair
	size    int
}

type KeyValuePair struct {
	key   string
	value any
}
