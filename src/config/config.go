package config

import (
	"encoding/json"
	"os"
)

const (
	defaultConfigPath = "conf/config.json"
)

var (
	raw    map[string]interface{}
	config = make(map[string]interface{})
)

func loadConfigs(prefix string, o interface{}) {
	if m, ok := o.(map[string]interface{}); ok {
		if len(prefix) > 0 {
			prefix = prefix + ":"
		}

		for k, v := range m {
			key := prefix + k
			loadConfigs(key, v)
		}
	} else {
		config[prefix] = o
	}
}

func LoadConfigs() (err error) {
	file, err := os.Open(defaultConfigPath)
	if err != nil {
		panic(err)
		return
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&raw)
	loadConfigs("", raw)
	return
}

func String(key string, defaultValue string) (value string) {
	value, _ = getString(key, defaultValue)
	return
}

func getString(key string, defaultValue string) (value string, found bool) {
	v, found := config[key]
	if !found {
		return defaultValue, false
	}
	if value, ok := v.(string); ok {
		return value, ok
	} else {
		return defaultValue, false
	}
}

func Int(key string, defaultValue int) (value int) {
	value, _ = getInt(key, defaultValue)
	return
}

func getInt(key string, defaultValue int) (value int, found bool) {
	v, found := config[key]
	if !found {
		value = defaultValue
		return
	}

	if v64, found := v.(float64); found {
		return int(v64), found
	} else {
		return defaultValue, false
	}
}

func Bool(key string, defaultValue bool) (value bool) {
	value, _ = getBool(key, defaultValue)
	return
}

func getBool(key string, defaultValue bool) (value bool, found bool) {
	v, found := config[key]
	if !found {
		value = defaultValue
		return
	}

	if b, ok := v.(bool); ok {
		return b, ok
	} else if i, ok := v.(int); ok {
		if i == 1 {
			return true, true
		} else if i == 0 {
			return false, true
		} else {
			return defaultValue, false
		}
	} else {
		return defaultValue, false
	}
}
