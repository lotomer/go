package common

import (
	"log"
)

// Cache 缓存管理器
var Cache Cachable

// 包初始化
func init() {
	Cache = CacheInMemory{Values: make(map[string]string)}
}

// Cachable 缓存接口
type Cachable interface {
	SetString(k, v string) error
	GetString(k string) string
}

// CacheInMemory 本地内存缓存
type CacheInMemory struct {
	Values map[string]string
}

// SetString 设置缓存
func (c CacheInMemory) SetString(k, v string) error {
	log.Printf("Set %s = %s", k, v)
	c.Values[k] = v
	return nil
}

// GetString 获取缓存
func (c CacheInMemory) GetString(k string) string {
	if v, b := c.Values[k]; b {
		log.Printf("Get %s = %s", k, v)
		return v
	}
	log.Printf("Get failed: %s", k)
	return ""
}
