package config

import (
	"os"

	"github.com/rs/zerolog/log"

	"github.com/spf13/viper"
)

// LoadConfig is a read configuration from config file
func LoadConfig(fpath string) {
	logData := make(map[string]interface{})
	logData["tag"] = "configuration"
	logData["file"] = fpath
	viper.SetConfigFile(fpath)
	viper.WatchConfig()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Caller().Fields(logData).Msg("Error reading config file, " + err.Error())
	}

	host, _ := os.Hostname()
	viper.Set("hostname", host)
	// oracle.InitConnection()
	// viper.OnConfigChange(func(in fsnotify.Event) {
	// 	oracle.InitConnection()
	// })
	log.Info().Fields(logData).Msg("Load Configuration Completed.")
}
