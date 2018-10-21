package common

import (
	"log"
)

// Cache 缓存管理器
var Cache Cachable

// 包初始化
func init() {
	Cache = CacheInMemory{StringValues: make(map[string]string), Int64Values: make(map[string]int64)}
}

// Cachable 缓存接口
type Cachable interface {
	SetString(k, v string) error
	GetString(k string) string
	SetInt64(k string, v int64) error
	GetInt64(k string) int64
}

// CacheInMemory 本地内存缓存
type CacheInMemory struct {
	StringValues map[string]string
	Int64Values  map[string]int64
}

// SetString 设置缓存
func (c CacheInMemory) SetString(k, v string) error {
	log.Printf("Set %s = %s", k, v)
	c.StringValues[k] = v
	return nil
}

// GetString 获取缓存
func (c CacheInMemory) GetString(k string) string {
	if v, b := c.StringValues[k]; b {
		log.Printf("Get %s = %s", k, v)
		return v
	}
	log.Printf("Get failed: %s", k)
	return ""
}

// SetInt64 设置缓存
func (c CacheInMemory) SetInt64(k string, v int64) error {
	log.Printf("Set %s = %d", k, v)
	c.Int64Values[k] = v
	return nil
}

// GetInt64 获取缓存
func (c CacheInMemory) GetInt64(k string) int64 {
	if v, b := c.Int64Values[k]; b {
		log.Printf("Get %s = %d", k, v)
		return v
	}
	log.Printf("Get failed: %s", k)
	return 0
}
