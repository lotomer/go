package datastore

import (
	"database/sql"
)

// DataConfigs 数据服务配置信息。key：数据服务编码
var DataConfigs map[string]dataConfig

// DataConfigPool 解析后的数据服务配置信息。key：数据服务编码
var DataConfigPool map[string]DataConfig

type sqlConfig struct {
	SQL string `json:"sql"`
}

// DataConfig 解析后的数据库服务配置
type DataConfig struct {
	DB         *sql.DB
	QueryParam []QueryParameter
	Returns    []Return
	Options    sqlConfig
}

// 数据服务配置信息
type dataConfig struct {
	ID         string
	Name       string
	QueryParam string
	Returns    string
	Options    string
	DsID       int
}

type nameText struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

// Return 返回值
type Return struct {
	nameText
}

// QueryParameter 查询参数配置
type QueryParameter struct {
	nameText
	DefaultValue string `json:"defaultValue"`
	// text|number|date|list
	Type     string `json:"type"`
	Required string `json:"required"`
	// range|like|eq
	Mode string                 `json:"mode"`
	Ext  map[string]interface{} `json:"ext"`
}
