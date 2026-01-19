package cache

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
)

// Cache provides disk-based caching for summaries
type Cache struct {
	dir string
}

// New creates a new cache with the given directory
func New(dir string) *Cache {
	return &Cache{dir: dir}
}

// Get retrieves data from the cache by key
func (c *Cache) Get(key string) ([]byte, bool) {
	path := c.path(key)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, false
	}
	return data, true
}

// Set stores data in the cache
func (c *Cache) Set(key string, data []byte) error {
	if err := os.MkdirAll(c.dir, 0755); err != nil {
		return err
	}

	path := c.path(key)
	return os.WriteFile(path, data, 0644)
}

// Delete removes an entry from the cache
func (c *Cache) Delete(key string) error {
	path := c.path(key)
	err := os.Remove(path)
	if os.IsNotExist(err) {
		return nil
	}
	return err
}

// path returns the file path for a cache key
func (c *Cache) path(key string) string {
	// Hash the key to avoid filesystem issues with special characters
	h := sha256.Sum256([]byte(key))
	name := hex.EncodeToString(h[:])
	return filepath.Join(c.dir, name)
}

// ContentHash computes a hash of the given content
func ContentHash(content ...string) string {
	h := sha256.New()
	for _, c := range content {
		h.Write([]byte(c))
	}
	return hex.EncodeToString(h.Sum(nil))
}
