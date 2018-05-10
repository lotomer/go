package privilege

import (
	"database/sql"
	"log"
	"strconv"

	"../datastore"
)

// Use 使用数据库模式初始化
func Use(db *sql.DB) {
	// 初始化用户数据
	initUsers(db)

	// 初始化角色列表
	initRoles(db)

	// 初始化用户角色
	initUserRoles(db)

	// 初始化用户url权限
	initUserURIPrivileges(db)

	// 初始化角色url权限
	initRoleURIPrivileges(db)
}

// 初始化角色url权限
func initRoleURIPrivileges(db *sql.DB) {
	datas, err := datastore.GetAPIDatas("getValidRoleApiUris", nil)
	if err != nil {
		panic(err)
	}
	//log.Print(datas)
	roleUrls := make(map[string]map[string]uint8)
	for _, row := range datas {
		if _, ok := roleUrls[row["ROLE_ID"].(string)]; !ok {
			roleUrls[row["ROLE_ID"].(string)] = make(map[string]uint8)
		}
		if urls, ok := roleUrls[row["ROLE_ID"].(string)]; ok {
			urls[row["URL"].(string)] = 1
		}

	}

	// 替换原缓存
	CachedRoleUrls = roleUrls
}

// 初始化用户url权限
func initUserURIPrivileges(db *sql.DB) {
	datas, err := datastore.GetAPIDatas("getValidUserApiUris", nil)
	if err != nil {
		panic(err)
	}
	//log.Print(datas)
	userUrls := make(map[int]map[string]uint8)
	for _, row := range datas {
		uid, err := strconv.Atoi(row["USER_ID"].(string))
		if err != nil {
			log.Printf("Convert userId(%s) to int failed: %s", row["USER_ID"].(string), err)
			continue
		}
		if _, ok := userUrls[uid]; !ok {
			userUrls[uid] = make(map[string]uint8)
		}
		if urls, ok := userUrls[uid]; ok {
			urls[row["URL"].(string)] = 1
		}
	}

	// 替换原缓存
	CachedUserUrls = userUrls
}

// 初始化用户角色
func initUserRoles(db *sql.DB) {
	datas, err := datastore.GetAPIDatas("getValidUserRelRoles", nil)
	if err != nil {
		panic(err)
	}
	//log.Print(datas)
	for _, row := range datas {
		if user, ok := CachedUsers[row["KEY"].(string)]; ok {
			if role, ok := CachedRoles[row["ROLE_ID"].(string)]; ok {
				user.Roles = append(user.Roles, role)
			}
		}
	}
}

// 初始化角色列表
func initRoles(db *sql.DB) {
	datas, err := datastore.GetAPIDatas("getValidRoles", nil)
	if err != nil {
		panic(err)
	}
	//log.Print(datas)
	roles := make(map[string]*Role)
	for _, row := range datas {
		roles[row["ROLE_ID"].(string)] = &Role{
			ID:        row["ROLE_ID"].(string),
			Name:      row["ROLE_NAME"].(string),
			AliveTime: row["ALIVE_TIME"].(int),
		}
	}

	// 替换原缓存
	CachedRoles = roles
}

// 初始化用户列表
func initUsers(db *sql.DB) {
	datas, err := datastore.GetAPIDatas("getValidUsers", nil)
	if err != nil {
		panic(err)
	}
	//log.Print(datas)
	users := make(map[string]*User)
	for _, row := range datas {
		users[row["KEY"].(string)] = &User{
			ID:                 row["USER_ID"].(int),
			Name:               row["USER_NAME"].(string),
			NickName:           row["NICK_NAME"].(string),
			AliveTime:          row["ALIVE_TIME"].(int),
			ClientIP:           row["CLIENT_IP"].(string),
			NeedChangePassword: row["NEED_CHANGE_PASSWORD"].(string) == "1",
			Roles:              []*Role{},
			//Password:           row["USER_PASSWD"].(string),
		}
	}
	//log.Print(users)
	// 替换原缓存
	CachedUsers = users
}
