package config

import "github.com/spf13/viper"

type Cfg struct {
	Db_URL			string	`mapstructure:"DB_URL"`
	Server_Port		string	`mapstructure:"SERVER_PORT"`
	Internal_URL	string	`mapstructure:"INTERNAL_URL"`
	External_URL	string	`mapstructure:"EXTERNAL_URL"`
}

func LoadEnvVariables(path string)(c Cfg, err error){
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)
	return
}