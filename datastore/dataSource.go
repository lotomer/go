package datastore

import "database/sql"

// DataSources 数据源配置信息。key：数据源编号
var DataSources map[int]dataSource

// DataSourcePool 数据源连接池
var DataSourcePool map[int]*sql.DB

// ThisDataSourceID 默认自身数据源编号
const ThisDataSourceID = -999

// 数据源配置信息
type dataSource struct {
	ID      int
	Name    string
	Options string
	Comment string
}
