// Package privilege URI权限
package privilege

import (
	"fmt"
)

// CachedUserUrls 缓存的用户url
var CachedUserUrls map[int]map[string]uint8

// CachedRoleUrls 缓存的角色url
var CachedRoleUrls map[string]map[string]uint8

// CheckURIPrivilege 校验用户是否有访问该uri的权限
func CheckURIPrivilege(user *User, uri string) error {
	// 1、校验用户访问该URI的权限
	//log.Print(CachedUserUrls)
	if urls, ok := CachedUserUrls[user.ID]; ok {
		if _, ok = urls[uri]; ok {
			// 匹配上
			return nil
		}
	}
	// 2、校验用户对应的角色访问该URI的权限
	for _, role := range user.Roles {
		if urls, ok := CachedRoleUrls[role.ID]; ok {
			if _, ok = urls[uri]; ok {
				// 匹配上
				return nil
			}
		}
	}

	return fmt.Errorf("User(%s) access uri(%s) forbidden", user.Name, uri)
}
