package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/project/go_template/docs")
	_ = viper.ReadInConfig()
	splitOsEnv()
}

func splitOsEnv() {
	for i := range os.Environ() {
		ks := strings.Split(os.Environ()[i], "=")
		key := ks[0]
		value := ks[1]
		if strings.HasPrefix(key, "OK_") {
			cfk := strings.Replace(strings.ToLower(strings.Replace(key, "OK_", "", -1)), "_", ".", -1)
			viper.Set(cfk, value)
		}
	}
}
