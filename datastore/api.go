package datastore

import (
	"bytes"
	"fmt"
	"log"
	"regexp"
	"strings"

	"../common"
)

var regexp4param = regexp.MustCompile(`\$\{([^}]+)\}`)

func test(inputStr string, flag string) []string {
	findStrs := regexp4param.FindAllStringSubmatch(inputStr, -1)
	var result []string
	for _, str := range findStrs {
		result = append(result, str[1])
	}
	return result
}

// GetAPIDatas 获取api对应的数据
func GetAPIDatas(dataID string, inputParams map[string][]string) ([]map[string]interface{}, error) {
	if config, ok := DataConfigPool[dataID]; ok {
		sql, db := config.Options.SQL, config.DB
		sql, queryParam, err := generateSQL(sql, inputParams, config.QueryParam)
		if err != nil {
			return nil, err
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

func generateSQL(sql string, inputParams map[string][]string, configQueryParam []QueryParameter) (string, []interface{}, error) {
	// 1、校验查询参数
	var queryParam = []interface{}{} //make(map[string]string)
	// sql样例： select * from a where ${paramA} ${paramB}
	if inputParams != nil {
		findStrs := regexp4param.FindAllStringSubmatch(sql, -1)

		// 参数名对应的参数配置项
		configParams := make(map[string]QueryParameter)
		for _, config := range configQueryParam {
			if config.Code == "" {
				config.Code = config.Name // 没有特别指定参数的代码，则直接用参数的名称代替
			}
			configParams[config.Name] = config
		}

		var sqlBuf bytes.Buffer
		for _, strs := range findStrs {
			paramName, paramVar := strs[1], strs[0]
			if param, ok := configParams[paramName]; ok {
				if paramValues, ok := inputParams[paramName]; ok {
					if param.Mode == "like" && param.Type == "text" {
						// 模糊匹配
						queryParam = append(queryParam, "%"+paramValues[0]+"%")
						sqlBuf.WriteString(" AND ")
						sqlBuf.WriteString(param.Code)
						sqlBuf.WriteString(" like ?")
					} else if param.Mode == "range" {
						// 范围匹配：起始值,终止值
						vs := strings.SplitN(paramValues[0], ",", 2)
						if vStart := strings.TrimSpace(vs[0]); vStart != "" {
							// 设置起始值
							queryParam = append(queryParam, vStart)
							sqlBuf.WriteString(" AND ")
							sqlBuf.WriteString(param.Code)
							sqlBuf.WriteString(" >= ?")
						}
						if len(vs) == 2 {
							if vEnd := strings.TrimSpace(vs[1]); vEnd != "" {
								// 设置起始值
								queryParam = append(queryParam, vEnd)
								sqlBuf.WriteString(" AND ")
								sqlBuf.WriteString(param.Code)
								sqlBuf.WriteString(" <= ?")
							}
						}
					} else if param.Text == "list" {
						// 值是数组
						sqlBuf.WriteString(" AND ")
						sqlBuf.WriteString(param.Code)
						sqlBuf.WriteString(" in (")
						for i, v := range paramValues {
							queryParam = append(queryParam, v)
							if i != 0 {
								sqlBuf.WriteString(",")
							}
							sqlBuf.WriteString("?")
						}

						sqlBuf.WriteString(")")
					} else {
						// 等值匹配eq
						queryParam = append(queryParam, paramValues[0])
						sqlBuf.WriteString(" AND ")
						sqlBuf.WriteString(param.Code)
						sqlBuf.WriteString(" = ?")
					}
				} else {
					// 查询参数中没有该参数
					if param.Required == "true" {
						return sql, nil, fmt.Errorf("缺少必备参数： %s", param.Name)
					}
				}
			}
			sql = strings.Replace(sql, paramVar, sqlBuf.String(), 1) // 一次只替换一个
			sqlBuf.Reset()
		}
	}
	return sql, queryParam, nil
}
