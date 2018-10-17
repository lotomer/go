package config

import (
	"encoding/json"
)

// Config 配置文件
var Config = config{}

type config struct {
	configs map[string]interface{}
}

// Use 初始化
func (c *config) Use(str []byte) error {
	c.configs = make(map[string]interface{})
	if err := json.Unmarshal(str, &c.configs); err != nil {
		return err
	} else {
		return nil
	}
}

// Get 根据名字获取配置值
func (c *config) Get(name string) interface{} {
	return c.configs[name]
}

// GetAll 获取所有配置
func (c *config) GetAll() map[string]interface{} {
	return c.configs
}

// GetString 根据名字获取配置值
func (c *config) GetString(name string) string {
	if data, ok := c.configs[name].(string); ok {
		return data
	} else {
		return ""
	}
}

// // GetInt 根据名字获取配置值
// func (c *config) GetInt(name string) (int, error) {
// 	if data, ok := c.configs[name].(int); !ok {
// 		return data, nil
// 	} else {
// 		return 0, fmt.Errorf("%v is not int", c.configs[name])
// 	}
// }

// // GetFloat64 根据名字获取配置值
// func (c *config) GetFloat64(name string) (float64, error) {
// 	if data, ok := c.configs[name].(float64); !ok {
// 		return data, nil
// 	} else {
// 		return 0, fmt.Errorf("%v is not float64", c.configs[name])
// 	}
// }
