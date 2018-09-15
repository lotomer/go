package datastore

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"strconv"

	mydb "github.com/lotomer/go/db"
)

// InitDataSource 初始化数据源配置
func InitDataSource(db *sql.DB) error {
	dataSources, err := loadDataSourceFromDB(db)
	if err != nil {
		return err
	}
	var dbTemp *sql.DB
	for id, dataSource := range dataSources {
		dbTemp, err = mydb.GenerateDBWithJSONStr(dataSource.Options)
		if err != nil {
			log.Fatalf("Create DB failed: %s", err)
			continue
		}
		mydb.AddDB(id, dbTemp)
	}

	// 最后再切换
	DataSources = dataSources
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
	dsIDs.WriteString(strconv.Itoa(mydb.ThisDataSourceID))
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
		dc.dsID = config.DsID
		if err := json.Unmarshal([]byte(config.Options), &dc.Options); err != nil {
			log.Printf("Parse json failed: %s, input: %s", err.Error(), config.Options)
			continue
		}
		if config.QueryParam != "" { // 查询参数可能为空
			if err := json.Unmarshal([]byte(config.QueryParam), &dc.QueryParam); err != nil {
				log.Printf("Parse json failed: %s, input: %s", err.Error(), config.QueryParam)
				continue
			}
		}
		if err := json.Unmarshal([]byte(config.Returns), &dc.Returns); err != nil {
			log.Printf("Parse json failed: %s, input: %s", err.Error(), config.Returns)
			continue
		}

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
