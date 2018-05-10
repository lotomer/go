// Package privilege 用户权限
package privilege

import (
	"fmt"
)

// CachedUsers 已缓存用户信息
var CachedUsers map[string]*User

// CachedRoles 已缓存角色信息
var CachedRoles map[string]*Role

// User 用户信息
type User struct {
	ID                 int     //用户编号
	Name               string  //登陆名
	NickName           string  // 昵称
	Password           string  // 密码
	AliveTime          int     // 会话存活时间
	ClientIP           string  // 有效客户的IP地址
	NeedChangePassword bool    // 是否需要修改密码
	Roles              []*Role // 归属角色
}

// Role 角色信息
type Role struct {
	ID        string // 角色编号
	Name      string // 角色名称
	AliveTime int    // 会话存活时间
}

// GetUserByKey 根据密钥key获取用户信息
func GetUserByKey(key string) (*User, error) {
	if user, ok := CachedUsers[key]; ok {
		return user, nil
	}

	return nil, fmt.Errorf("The key donnot match a user: %s", key)
}
