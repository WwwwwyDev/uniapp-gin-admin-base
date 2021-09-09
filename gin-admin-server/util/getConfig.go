package util

import "gin-server/global"

func GetConfig(root string, key string) interface{}{
	return global.VIPER.Get(root + "." + key)
}
