package datastore

// DataSources 数据源配置信息。key：数据源编号
var DataSources map[int]dataSource

// 数据源配置信息
type dataSource struct {
	ID      int
	Name    string
	Options string
	Comment string
}
