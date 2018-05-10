// Package privilege 用户权限
package privilege

import (
	"fmt"
)

// CachedUsers 已缓存用户信息
var CachedUsers *map[string]*User

// User 用户信息
type User struct {
	ID       int    //编号
	Name     string //登陆名
	NickName string // 昵称
	Password string // 密码
}

// GetUserByKey 根据密钥key获取用户信息
func GetUserByKey(key string) (*User, error) {
	if user, ok := (*CachedUsers)[key]; ok {
		return user, nil
	}

	return nil, fmt.Errorf("The key donnot match a user: %s", key)
}
