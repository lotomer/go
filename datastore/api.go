package datastore

import (
	"bytes"
	"fmt"
	"log"

	"../common"
)

// GetAPIDatas 获取api对应的数据
func GetAPIDatas(dataID string, inputParams map[string]string) ([]map[string]interface{}, error) {
	if config, ok := DataConfigPool[dataID]; ok {
		sql, db := config.Options.SQL, config.DB
		var sqlBuf bytes.Buffer
		// 1、校验查询参数
		var queryParam = []interface{}{} //make(map[string]string)
		if inputParams != nil {
			sqlBuf.WriteString("select * from (")
			sqlBuf.WriteString(sql)
			sqlBuf.WriteString(") a where 1=1")
			for _, param := range config.QueryParam {
				if paramValue, ok := inputParams[param.Name]; ok {
					queryParam = append(queryParam, paramValue)
					sqlBuf.WriteString(" AND `")
					sqlBuf.WriteString(param.Name)
					sqlBuf.WriteString("` = ?")
				} else {
					if param.Required == "true" {
						return nil, fmt.Errorf("缺少必备参数： %s", param.Name)
					}
				}
			}
			sql = sqlBuf.String()
		}
		// 2、执行查询
		log.Print(sql)
		cols, datas, err := common.LoadDatasFromDB(db, sql, queryParam...)
		if err != nil {
			return nil, fmt.Errorf("查询失败：  %s", err.Error())
		}
		log.Printf("Execute %s's SQL: %s", dataID, sql)

		// 3、校验返回结果
		returnCols := make(map[string]int)
		for _, v := range config.Returns {
			returnCols[v.Name] = -1
		}
		for i := range cols {
			if _, ok := returnCols[cols[i]]; ok {
				// 需要返回，则提取字段值，设置为索引号
				returnCols[cols[i]] = i
			}
		}
		resultDatas := []map[string]interface{}{}
		for i := range datas {
			data := make(map[string]interface{})
			for k, v := range returnCols {
				if v != -1 {
					data[k] = datas[i][v]
				}
			}
			resultDatas = append(resultDatas, data)
		}
		// 4、返回结果
		return resultDatas, nil
	} else {
		return nil, fmt.Errorf("API不存在：  %s", dataID)
	}
}
