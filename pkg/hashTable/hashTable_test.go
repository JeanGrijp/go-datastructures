package hashtable

import (
	"strconv"
	"testing"
)

// TestNew tests the creation of a new hash table
func TestNew(t *testing.T) {
	// Test with positive capacity
	ht := New(10)
	if ht.capacity != 10 {
		t.Errorf("Expected capacity 10, got %d", ht.capacity)
	}
	if ht.size != 0 {
		t.Errorf("Expected size 0, got %d", ht.size)
	}
	if len(ht.buckets) != 10 {
		t.Errorf("Expected 10 buckets, got %d", len(ht.buckets))
	}

	// Test with zero capacity (should default to 16)
	ht = New(0)
	if ht.capacity != 16 {
		t.Errorf("Expected default capacity 16, got %d", ht.capacity)
	}

	// Test with negative capacity (should default to 16)
	ht = New(-5)
	if ht.capacity != 16 {
		t.Errorf("Expected default capacity 16, got %d", ht.capacity)
	}
}

// TestPutAndGet tests inserting and retrieving values
func TestPutAndGet(t *testing.T) {
	ht := New(5)

	// Test putting new key
	isNew := ht.Put("key1", "value1")
	if !isNew {
		t.Error("Expected Put to return true for new key")
	}
	if ht.Size() != 1 {
		t.Errorf("Expected size 1, got %d", ht.Size())
	}

	// Test getting the value
	value, found := ht.Get("key1")
	if !found {
		t.Error("Expected to find key1")
	}
	if value != "value1" {
		t.Errorf("Expected value1, got %v", value)
	}

	// Test updating existing key
	isNew = ht.Put("key1", "newvalue1")
	if isNew {
		t.Error("Expected Put to return false for existing key")
	}
	if ht.Size() != 1 {
		t.Errorf("Expected size to remain 1, got %d", ht.Size())
	}

	// Test getting updated value
	value, found = ht.Get("key1")
	if !found {
		t.Error("Expected to find key1 after update")
	}
	if value != "newvalue1" {
		t.Errorf("Expected newvalue1, got %v", value)
	}

	// Test getting non-existent key
	_, found = ht.Get("nonexistent")
	if found {
		t.Error("Expected not to find nonexistent key")
	}
}

// TestDelete tests deleting key-value pairs
func TestDelete(t *testing.T) {
	ht := New(5)

	// Add some data
	ht.Put("key1", "value1")
	ht.Put("key2", "value2")
	ht.Put("key3", "value3")

	// Test deleting existing key
	deleted := ht.Delete("key2")
	if !deleted {
		t.Error("Expected Delete to return true for existing key")
	}
	if ht.Size() != 2 {
		t.Errorf("Expected size 2 after deletion, got %d", ht.Size())
	}

	// Test that deleted key is no longer found
	_, found := ht.Get("key2")
	if found {
		t.Error("Expected not to find deleted key")
	}

	// Test deleting non-existent key
	deleted = ht.Delete("nonexistent")
	if deleted {
		t.Error("Expected Delete to return false for non-existent key")
	}
	if ht.Size() != 2 {
		t.Errorf("Expected size to remain 2, got %d", ht.Size())
	}
}

// TestContains tests the Contains method
func TestContains(t *testing.T) {
	ht := New(5)

	// Test with empty hash table
	if ht.Contains("key1") {
		t.Error("Expected not to contain key1 in empty hash table")
	}

	// Add some data
	ht.Put("key1", "value1")
	ht.Put("key2", "value2")

	// Test existing keys
	if !ht.Contains("key1") {
		t.Error("Expected to contain key1")
	}
	if !ht.Contains("key2") {
		t.Error("Expected to contain key2")
	}

	// Test non-existent key
	if ht.Contains("key3") {
		t.Error("Expected not to contain key3")
	}
}

// TestSizeAndIsEmpty tests size-related methods
func TestSizeAndIsEmpty(t *testing.T) {
	ht := New(5)

	// Test empty hash table
	if !ht.IsEmpty() {
		t.Error("Expected hash table to be empty")
	}
	if ht.Size() != 0 {
		t.Errorf("Expected size 0, got %d", ht.Size())
	}

	// Add some data
	ht.Put("key1", "value1")
	if ht.IsEmpty() {
		t.Error("Expected hash table not to be empty")
	}
	if ht.Size() != 1 {
		t.Errorf("Expected size 1, got %d", ht.Size())
	}

	ht.Put("key2", "value2")
	if ht.Size() != 2 {
		t.Errorf("Expected size 2, got %d", ht.Size())
	}

	// Delete all data
	ht.Delete("key1")
	ht.Delete("key2")
	if !ht.IsEmpty() {
		t.Error("Expected hash table to be empty after deleting all")
	}
	if ht.Size() != 0 {
		t.Errorf("Expected size 0 after deleting all, got %d", ht.Size())
	}
}

// TestClear tests the Clear method
func TestClear(t *testing.T) {
	ht := New(5)

	// Add some data
	ht.Put("key1", "value1")
	ht.Put("key2", "value2")
	ht.Put("key3", "value3")

	// Clear the hash table
	ht.Clear()

	// Test that it's empty
	if !ht.IsEmpty() {
		t.Error("Expected hash table to be empty after Clear")
	}
	if ht.Size() != 0 {
		t.Errorf("Expected size 0 after Clear, got %d", ht.Size())
	}

	// Test that keys are no longer found
	if ht.Contains("key1") {
		t.Error("Expected not to contain key1 after Clear")
	}
	if ht.Contains("key2") {
		t.Error("Expected not to contain key2 after Clear")
	}
	if ht.Contains("key3") {
		t.Error("Expected not to contain key3 after Clear")
	}
}

// TestKeys tests the Keys method
func TestKeys(t *testing.T) {
	ht := New(5)

	// Test empty hash table
	keys := ht.Keys()
	if len(keys) != 0 {
		t.Errorf("Expected 0 keys, got %d", len(keys))
	}

	// Add some data
	ht.Put("key1", "value1")
	ht.Put("key2", "value2")
	ht.Put("key3", "value3")

	keys = ht.Keys()
	if len(keys) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(keys))
	}

	// Check that all keys are present (order may vary)
	keyMap := make(map[string]bool)
	for _, key := range keys {
		keyMap[key] = true
	}

	if !keyMap["key1"] || !keyMap["key2"] || !keyMap["key3"] {
		t.Error("Not all expected keys found in Keys() result")
	}
}

// TestValues tests the Values method
func TestValues(t *testing.T) {
	ht := New(5)

	// Test empty hash table
	values := ht.Values()
	if len(values) != 0 {
		t.Errorf("Expected 0 values, got %d", len(values))
	}

	// Add some data
	ht.Put("key1", "value1")
	ht.Put("key2", "value2")
	ht.Put("key3", "value3")

	values = ht.Values()
	if len(values) != 3 {
		t.Errorf("Expected 3 values, got %d", len(values))
	}

	// Check that all values are present (order may vary)
	valueMap := make(map[string]bool)
	for _, value := range values {
		if str, ok := value.(string); ok {
			valueMap[str] = true
		}
	}

	if !valueMap["value1"] || !valueMap["value2"] || !valueMap["value3"] {
		t.Error("Not all expected values found in Values() result")
	}
}

// TestGetPairs tests the GetPairs method
func TestGetPairs(t *testing.T) {
	ht := New(5)

	// Test empty hash table
	pairs := ht.GetPairs()
	if len(pairs) != 0 {
		t.Errorf("Expected 0 pairs, got %d", len(pairs))
	}

	// Add some data
	ht.Put("key1", "value1")
	ht.Put("key2", "value2")

	pairs = ht.GetPairs()
	if len(pairs) != 2 {
		t.Errorf("Expected 2 pairs, got %d", len(pairs))
	}

	// Check that pairs are correct
	pairMap := make(map[string]string)
	for _, pair := range pairs {
		pairMap[pair.Key()] = pair.Value().(string)
	}

	if pairMap["key1"] != "value1" || pairMap["key2"] != "value2" {
		t.Error("Pairs do not match expected key-value mappings")
	}
}

// TestLoadFactor tests the LoadFactor method
func TestLoadFactor(t *testing.T) {
	ht := New(10)

	// Test empty hash table
	if ht.LoadFactor() != 0.0 {
		t.Errorf("Expected load factor 0.0, got %f", ht.LoadFactor())
	}

	// Add 5 elements (load factor should be 0.5)
	for i := 0; i < 5; i++ {
		ht.Put("key"+strconv.Itoa(i), "value"+strconv.Itoa(i))
	}

	expected := 0.5
	if ht.LoadFactor() != expected {
		t.Errorf("Expected load factor %f, got %f", expected, ht.LoadFactor())
	}

	// Add 5 more elements (load factor should be 1.0)
	for i := 5; i < 10; i++ {
		ht.Put("key"+strconv.Itoa(i), "value"+strconv.Itoa(i))
	}

	expected = 1.0
	if ht.LoadFactor() != expected {
		t.Errorf("Expected load factor %f, got %f", expected, ht.LoadFactor())
	}
}

// TestBucketDistribution tests the BucketDistribution method
func TestBucketDistribution(t *testing.T) {
	ht := New(3)

	// Test empty hash table
	distribution := ht.BucketDistribution()
	if distribution[0] != 3 {
		t.Errorf("Expected 3 empty buckets, got %d", distribution[0])
	}

	// Add some data
	ht.Put("key1", "value1")
	ht.Put("key2", "value2")

	distribution = ht.BucketDistribution()

	// Should have some buckets with 0 items and some with 1 or more
	totalBuckets := 0
	for _, count := range distribution {
		totalBuckets += count
	}

	if totalBuckets != 3 {
		t.Errorf("Expected total of 3 buckets, got %d", totalBuckets)
	}
}

// TestKeyValuePairMethods tests the Key() and Value() methods
func TestKeyValuePairMethods(t *testing.T) {
	kvp := KeyValuePair{key: "testkey", value: "testvalue"}

	if kvp.Key() != "testkey" {
		t.Errorf("Expected key 'testkey', got '%s'", kvp.Key())
	}

	if kvp.Value() != "testvalue" {
		t.Errorf("Expected value 'testvalue', got '%v'", kvp.Value())
	}
}

// TestDifferentDataTypes tests storing different data types
func TestDifferentDataTypes(t *testing.T) {
	ht := New(10)

	// Test different data types
	ht.Put("string", "hello")
	ht.Put("int", 42)
	ht.Put("float", 3.14)
	ht.Put("bool", true)
	ht.Put("slice", []int{1, 2, 3})

	// Retrieve and check types
	if value, found := ht.Get("string"); !found || value.(string) != "hello" {
		t.Error("String value not stored/retrieved correctly")
	}

	if value, found := ht.Get("int"); !found || value.(int) != 42 {
		t.Error("Int value not stored/retrieved correctly")
	}

	if value, found := ht.Get("float"); !found || value.(float64) != 3.14 {
		t.Error("Float value not stored/retrieved correctly")
	}

	if value, found := ht.Get("bool"); !found || value.(bool) != true {
		t.Error("Bool value not stored/retrieved correctly")
	}

	if value, found := ht.Get("slice"); !found {
		t.Error("Slice value not found")
	} else if slice, ok := value.([]int); !ok || len(slice) != 3 {
		t.Error("Slice value not stored/retrieved correctly")
	}
}

// TestCollisionHandling tests that collisions are handled correctly
func TestCollisionHandling(t *testing.T) {
	// Use a very small hash table to force collisions
	ht := New(2)

	// Add many keys to force collisions
	for i := 0; i < 10; i++ {
		key := "key" + strconv.Itoa(i)
		value := "value" + strconv.Itoa(i)
		ht.Put(key, value)
	}

	// All keys should still be retrievable
	for i := 0; i < 10; i++ {
		key := "key" + strconv.Itoa(i)
		expectedValue := "value" + strconv.Itoa(i)

		if value, found := ht.Get(key); !found || value != expectedValue {
			t.Errorf("Key %s not found or has wrong value after collision", key)
		}
	}

	// Size should be correct
	if ht.Size() != 10 {
		t.Errorf("Expected size 10, got %d", ht.Size())
	}
}

// TestString tests the String method
func TestString(t *testing.T) {
	ht := New(3)
	ht.Put("key1", "value1")
	ht.Put("key2", "value2")

	str := ht.String()
	if str == "" {
		t.Error("String representation should not be empty")
	}

	// Should contain basic information
	if !contains(str, "HashTable") {
		t.Error("String should contain 'HashTable'")
	}
	if !contains(str, "size: 2") {
		t.Error("String should contain correct size")
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && (s[:len(substr)] == substr || s[len(s)-len(substr):] == substr || containsAtIndex(s, substr)))
}

func containsAtIndex(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// Benchmark tests
func BenchmarkPut(b *testing.B) {
	ht := New(1000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := "key" + strconv.Itoa(i)
		ht.Put(key, i)
	}
}

func BenchmarkGet(b *testing.B) {
	ht := New(1000)

	// Pre-populate with data
	for i := 0; i < 1000; i++ {
		key := "key" + strconv.Itoa(i)
		ht.Put(key, i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := "key" + strconv.Itoa(i%1000)
		ht.Get(key)
	}
}

func BenchmarkDelete(b *testing.B) {
	// Pre-populate with data
	ht := New(1000)
	for i := 0; i < b.N; i++ {
		key := "key" + strconv.Itoa(i)
		ht.Put(key, i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := "key" + strconv.Itoa(i)
		ht.Delete(key)
	}
}

func BenchmarkContains(b *testing.B) {
	ht := New(1000)

	// Pre-populate with data
	for i := 0; i < 1000; i++ {
		key := "key" + strconv.Itoa(i)
		ht.Put(key, i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := "key" + strconv.Itoa(i%1000)
		ht.Contains(key)
	}
}
