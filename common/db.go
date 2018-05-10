package common

import (
	"database/sql"
	"log"
	"strconv"
)

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
