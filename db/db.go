package db

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"strings"
)

// dataSourcePool 数据源连接池
var dataSourcePool map[int]*sql.DB

// ThisDataSourceID 默认自身数据源编号
const ThisDataSourceID = -999

// 文件初始化
func init() {
	dataSourcePool = make(map[int]*sql.DB)
}

// Use 保存当前数据库
func Use(db *sql.DB) {
	AddDB(ThisDataSourceID, db)
}

// AddDB 保存当前数据库
func AddDB(dsID int, db *sql.DB) {
	dataSourcePool[dsID] = db
}

// GetDB 获取默认数据库对象
func GetDB(dsID int) *sql.DB {
	if dsID < 0 {
		return dataSourcePool[ThisDataSourceID]
	} else {
		return dataSourcePool[dsID]
	}
}

// dbInfo 数据库配置信息
type dbInfo struct {
	Port        uint16 `json:"port"`
	Host        string `json:"host"`
	DBname      string `json:"dbname"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Type        string `json:"type"`
	Maxpoolsize uint16 `json:"maxPoolSize" `
	Maxidlesize uint16 `json:"maxIdleSize" `
	Urltemplate string `json:"urlTemplate" `
}

func generatURL(dbinfo *dbInfo) string {
	url := dbinfo.Urltemplate
	url = strings.Replace(url, "${port}", strconv.Itoa(int(dbinfo.Port)), -1)
	url = strings.Replace(url, "${host}", dbinfo.Host, -1)
	url = strings.Replace(url, "${dbname}", dbinfo.DBname, -1)
	url = strings.Replace(url, "${username}", dbinfo.Username, -1)
	url = strings.Replace(url, "${password}", dbinfo.Password, -1)
	return url
}

// GenerateDB 根据数据库配置信息获取数据库操作指针
func GenerateDB(dbInfo *dbInfo) (*sql.DB, error) {
	url := generatURL(dbInfo)
	log.Printf("db url: %s", url)
	return sql.Open(dbInfo.Type, url)
}

// GenerateDBWithJSONStr 根据数据库配置信息(json字符串)获取数据库操作指针
func GenerateDBWithJSONStr(dbinfoStr string) (*sql.DB, error) {
	dbInfo := &dbInfo{}
	err := json.Unmarshal([]byte(dbinfoStr), dbInfo)
	if err != nil {
		return nil, errors.New(err.Error() + ": " + dbinfoStr)
	}
	return GenerateDB(dbInfo)
}

// LoadDatasFromDB 根据sql获取所有数据
func LoadDatasFromDB(db *sql.DB, sqlStr string, args ...interface{}) ([]string, [][]interface{}, error) {
	var (
		rows *sql.Rows
		err  error
	)
	if len(args) > 0 {
		// 有附加参数，则先进行预编译，然后再执行
		stmt, err := db.Prepare(sqlStr)
		if err != nil {
			return nil, nil, err
		}
		defer stmt.Close()
		rows, err = stmt.Query(args...)
		if err != nil {
			//log.Fatalf("Query failed: %s",err)
			return nil, nil, err
		}
		defer rows.Close()

	} else {
		// 没有附加参数，则直接执行
		rows, err = db.Query(sqlStr)
		if err != nil {
			//log.Fatalf("Query failed: %s",err)
			return nil, nil, err
		}
		defer rows.Close()
	}

	cols, err := rows.Columns()
	if err != nil {
		return nil, nil, err
	}
	colTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, nil, err
	}

	dataConfigs := [][]interface{}{}
	//dataConfigs := make([]map[string]interface{})
	rawResult := make([]sql.RawBytes, len(colTypes))

	dest := make([]interface{}, len(colTypes))
	for i := range rawResult {
		dest[i] = &rawResult[i]
	}
	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			return nil, nil, err
		}
		result := make([]interface{}, len(colTypes))
		//dataConfig := make(map[string]interface{})
		for i, colInfo := range colTypes {
			if colInfo == nil {
				result[i] = ""
				continue
			}
			// TODO 待完善
			switch t := colInfo.DatabaseTypeName(); t {
			case "VARCHAR", "CHAR":
				result[i] = string(rawResult[i])
			case "INT":
				result[i], _ = strconv.Atoi(string(rawResult[i]))
			default:
				log.Printf("un implementation data type: %s", t)
				result[i] = string(rawResult[i])
			}
		}
		dataConfigs = append(dataConfigs, result)
	}
	return cols, dataConfigs, nil
}
