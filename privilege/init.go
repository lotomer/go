package privilege

import (
	"database/sql"

	"../common"
)

// Use 使用数据库模式初始化
func Use(db *sql.DB) {
	// 初始化用户数据
	_, rows, err := common.LoadDatasFromDB(db, "SELECT k.`KEY`,u.USER_ID,u.USER_NAME,u.NICK_NAME,u.USER_PASSWD FROM T_OF_U_KEY k JOIN T_OF_SYS_USERS u ON (k.USER_ID=u.USER_ID) WHERE k.INVALID_TIME > NOW() and IS_VALID='1'")
	if err != nil {
		panic(err)
	}
	users := make(map[string]*User)
	for _, row := range rows {
		users[row[0].(string)] = &User{ID: row[1].(int), Name: row[2].(string), NickName: row[3].(string), Password: row[4].(string)}
	}
	CachedUsers = &users
}
