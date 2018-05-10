// Package privilege URI权限
package privilege

import (
	"fmt"
)

// CheckURIPrivilege 校验用户是否有访问该uri的权限
func CheckURIPrivilege(user *User, uri string) error {
	// 1、校验用户访问该URI的权限
	// 2、校验用户对应的角色访问该URI的权限
	return fmt.Errorf("User(%s) access uri(%s) forbidden", user.Name, uri)
}

// CheckAPIPrivilege 校验用户是否有访问该API的权限
func CheckAPIPrivilege(user *User, api string) error {
	// 1、校验用户访问该api的权限
	// 2、校验用户对应的角色访问该api的权限
	return fmt.Errorf("User(%s) access api(%s) forbidden", user.Name, api)
}
