package datastore

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"strings"
)

// DBInfo 数据库配置信息
type DBInfo struct {
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

func generatURL(dbinfo *DBInfo) string {
	url := dbinfo.Urltemplate
	url = strings.Replace(url, "${port}", strconv.Itoa(int(dbinfo.Port)), -1)
	url = strings.Replace(url, "${host}", dbinfo.Host, -1)
	url = strings.Replace(url, "${dbname}", dbinfo.DBname, -1)
	url = strings.Replace(url, "${username}", dbinfo.Username, -1)
	url = strings.Replace(url, "${password}", dbinfo.Password, -1)
	return url
}

// GenerateDB 根据数据库配置信息获取数据库操作指针
func GenerateDB(dbInfo *DBInfo) (*sql.DB, error) {
	url := generatURL(dbInfo)
	log.Printf("db url: %s", url)
	return sql.Open(dbInfo.Type, url)
}

// GenerateDBWithJSONStr 根据数据库配置信息(json字符串)获取数据库操作指针
func GenerateDBWithJSONStr(dbinfoStr string) (*sql.DB, error) {
	dbInfo := &DBInfo{}
	err := json.Unmarshal([]byte(dbinfoStr), dbInfo)
	if err != nil {
		return nil, errors.New(err.Error() + ": " + dbinfoStr)
	}
	return GenerateDB(dbInfo)
}
func InitDataSource(db *sql.DB) error {
	DataSources, err := loadDataSourceFromDB(db)
	if err != nil {
		return err
	}
	dataSourcePool := make(map[int]*sql.DB)
	dataSourcePool[ThisDataSourceID] = db
	var dbTemp *sql.DB
	for id := range DataSources {
		dbTemp, err = GenerateDBWithJSONStr(DataSources[id].Options)
		if err != nil {
			log.Fatalf("Create DB failed: %s", err)
			continue
		}
		dataSourcePool[id] = dbTemp
	}

	// 最后再切换
	DataSourcePool = dataSourcePool
	return nil
}

// Use 使用数据库方式
func Use(db *sql.DB) {
	err := InitDataSource(db)
	if err != nil {
		panic(err)
	}

	err = InitDataConfig(db)
	if err != nil {
		panic(err)
	}
}

// InitDataConfig 主动初始化
func InitDataConfig(db *sql.DB) error {
	var dsIDs bytes.Buffer
	dsIDs.WriteString(strconv.Itoa(ThisDataSourceID))
	for id := range DataSources {
		dsIDs.WriteByte(',')
		dsIDs.WriteString(strconv.Itoa(id))

	}

	DataConfigs, err := loadDataConfigFromDB(db, dsIDs.String())
	if err != nil {
		return err
	}

	dataConfigPool := make(map[string]DataConfig)
	for id, config := range DataConfigs {
		dc := DataConfig{}
		dc.DB = DataSourcePool[config.DsID]
		json.Unmarshal([]byte(config.Options), &dc.Options)
		json.Unmarshal([]byte(config.QueryParam), &dc.QueryParam)
		json.Unmarshal([]byte(config.Returns), &dc.Returns)
		dataConfigPool[id] = dc
	}

	// 最后再切换
	DataConfigPool = dataConfigPool
	return nil
}

// 从数据库加载数据源配置信息
func loadDataSourceFromDB(db *sql.DB) (map[int]dataSource, error) {
	rows, err := db.Query("select ds_id,name,options,comment from T_OF_DS_DATASOURCE where DS_TYPE_ID=6 and is_valid=1 and OPTIONS <>'' and OPTIONS is not null")
	if err != nil {
		//log.Fatalf("Query failed: %s",err)
		return nil, err
	}
	defer rows.Close()
	datasources := make(map[int]dataSource)
	for rows.Next() {
		ds := dataSource{}
		err = rows.Scan(&ds.ID, &ds.Name, &ds.Options, &ds.Comment)
		if err != nil {
			return nil, err
		}
		datasources[ds.ID] = ds
	}
	return datasources, nil
}

// 从数据库加载数据源配置信息
func loadDataConfigFromDB(db *sql.DB, dsIDs string) (map[string]dataConfig, error) {
	rows, err := db.Query("select ID,NAME,QUERY_PARAMS,RETURNS,OPTIONS,DS_ID from T_OF_DS_SERVICE_CONFIG where is_valid=1 and status <>0 and DS_ID in (" + dsIDs + ")")
	if err != nil {
		//log.Fatalf("Query failed: %s",err)
		return nil, err
	}
	defer rows.Close()
	var queryParam sql.NullString

	dataConfigs := make(map[string]dataConfig)
	for rows.Next() {
		dc := dataConfig{}
		err = rows.Scan(&dc.ID, &dc.Name, &queryParam, &dc.Returns, &dc.Options, &dc.DsID)
		if err != nil {
			return nil, err
		}
		dc.QueryParam = queryParam.String
		dataConfigs[dc.ID] = dc
	}
	return dataConfigs, nil
}