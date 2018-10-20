package db

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/lotomer/go/config"
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

// GetDefaultDB 获取默认数据库对象
func GetDefaultDB() (*sql.DB, error) {
	return GetDB(ThisDataSourceID)
}

// GetDB 获取数据库对象
func GetDB(dsID int) (*sql.DB, error) {
	var db *sql.DB
	if dsID < 0 {
		db = dataSourcePool[ThisDataSourceID]
	} else {
		db = dataSourcePool[dsID]
	}
	if db == nil {
		return nil, fmt.Errorf("Get DB failed for dsID: %d", dsID)
	}
	return db, nil
}

// dbInfo 数据库配置信息
type dbInfo struct {
	Port        uint16 `json:"port"`
	Host        string `json:"host"`
	DBname      string `json:"dbname"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Type        string `json:"type"`
	MaxPoolSize int    `json:"maxPoolSize" `
	MaxIdleSize int    `json:"maxIdleSize" `
	URLTemplate string `json:"urlTemplate" `
}

func generatURL(dbinfo *dbInfo) string {
	url := dbinfo.URLTemplate
	url = strings.Replace(url, "${port}", strconv.Itoa(int(dbinfo.Port)), -1)
	url = strings.Replace(url, "${host}", dbinfo.Host, -1)
	url = strings.Replace(url, "${dbname}", dbinfo.DBname, -1)
	url = strings.Replace(url, "${username}", dbinfo.Username, -1)
	url = strings.Replace(url, "${password}", dbinfo.Password, -1)
	return url
}

// GenerateDB 根据数据库配置信息获取数据库操作指针
func GenerateDB() (*sql.DB, error) {
	// dbInfo := dbInfo{}
	// dbInfo.Urltemplate = config.Config.Get("urlTemplate").(string)
	// dbInfo.Port = uint16(config.Config.Get("port"))
	// dbInfo.Host = config.Config.Get("host").(string)
	// dbInfo.DBname = config.Config.Get("dbname").(string)
	// dbInfo.Username = config.Config.Get("username").(string)
	// dbInfo.Password = config.Config.Get("password").(string)
	// return generateDB(&dbInfo)
	data, err := json.Marshal(config.Config.GetAll())
	if err != nil {
		return nil, err
	}

	return GenerateDBWithJSONStr(data)
}
func generateDB(dbinfo *dbInfo) (*sql.DB, error) {
	url := generatURL(dbinfo)
	log.Printf("db url: %s", url)
	db, err := sql.Open(dbinfo.Type, url)
	if err != nil {
		return nil, err
	}
	if dbinfo.MaxPoolSize > 0 {
		db.SetMaxOpenConns(dbinfo.MaxPoolSize)
	}
	if dbinfo.MaxIdleSize > 0 {
		db.SetMaxIdleConns(dbinfo.MaxIdleSize)
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

// GenerateDBWithJSONStr 根据数据库配置信息(json字符串)获取数据库操作指针
func GenerateDBWithJSONStr(dbinfoStr []byte) (*sql.DB, error) {
	dbInfo := &dbInfo{}
	err := json.Unmarshal(dbinfoStr, dbInfo)
	if err != nil {
		return nil, errors.New(err.Error() + ": " + string(dbinfoStr))
	}

	return generateDB(dbInfo)
}

// Insert 插入一条数据，并返回新增的id
func Insert(db *sql.DB, sqlStr string, args ...interface{}) (int64, error) {
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	ret, err := stmt.Exec(args...)
	if err != nil {
		return 0, err
	}
	id, err := ret.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Exec 执行SQL，并返回影响的记录数
func Exec(db *sql.DB, sqlStr string, args ...interface{}) (int64, error) {
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	ret, err := stmt.Exec(args...)
	if err != nil {
		return 0, err
	}
	num, err := ret.RowsAffected()
	if err != nil {
		return 0, err
	}
	return num, nil
}

// Query 查询
func Query(db *sql.DB, sqlStr string, args ...interface{}) (*sql.Rows, error) {
	var (
		rows *sql.Rows
		err  error
	)
	if len(args) > 0 {
		// 有附加参数，则先进行预编译，然后再执行
		stmt, err := db.Prepare(sqlStr)
		if err != nil {
			return nil, err
		}
		defer stmt.Close()
		rows, err = stmt.Query(args...)
		if err != nil {
			//log.Fatalf("Query failed: %s",err)
			return nil, err
		}
	} else {
		// 没有附加参数，则直接执行
		rows, err = db.Query(sqlStr)
		if err != nil {
			//log.Fatalf("Query failed: %s",err)
			return nil, err
		}
	}
	return rows, nil
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
