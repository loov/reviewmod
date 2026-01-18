// cache/cache_test.go
package cache

import (
	"testing"
)

func TestCache_GetSet(t *testing.T) {
	dir := t.TempDir()
	c := New(dir)

	key := "test-key-123"
	data := []byte(`{"purpose": "test function"}`)

	// Should miss initially
	if _, ok := c.Get(key); ok {
		t.Error("expected cache miss")
	}

	// Set and get
	if err := c.Set(key, data); err != nil {
		t.Fatalf("Set: %v", err)
	}

	got, ok := c.Get(key)
	if !ok {
		t.Fatal("expected cache hit")
	}

	if string(got) != string(data) {
		t.Errorf("got %s, want %s", got, data)
	}
}

func TestCache_Persistence(t *testing.T) {
	dir := t.TempDir()

	key := "persist-key"
	data := []byte("persistent data")

	// Write with first cache instance
	c1 := New(dir)
	if err := c1.Set(key, data); err != nil {
		t.Fatalf("Set: %v", err)
	}

	// Read with second cache instance
	c2 := New(dir)
	got, ok := c2.Get(key)
	if !ok {
		t.Fatal("expected cache hit after reload")
	}

	if string(got) != string(data) {
		t.Errorf("got %s, want %s", got, data)
	}
}

func TestCache_Delete(t *testing.T) {
	dir := t.TempDir()
	c := New(dir)

	key := "delete-key"
	data := []byte("to be deleted")

	// Set then delete
	if err := c.Set(key, data); err != nil {
		t.Fatalf("Set: %v", err)
	}

	if err := c.Delete(key); err != nil {
		t.Fatalf("Delete: %v", err)
	}

	// Should miss after delete
	if _, ok := c.Get(key); ok {
		t.Error("expected cache miss after delete")
	}

	// Delete non-existent key should not error
	if err := c.Delete("nonexistent"); err != nil {
		t.Errorf("Delete nonexistent: %v", err)
	}
}
