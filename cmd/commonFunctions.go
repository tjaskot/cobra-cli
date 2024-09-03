package cmd

import (
  "log"
  
  "github.com/spf13/viper"
)

func SetViperConfigs() string {
  viper.SetConfigName(".bo.env")
  viper.SetConfigType("env")
  viper.AddConfigPath("$HOME/")
  viperErr := viper.ReadInConfig()
  if viperErr != nil {
    log.Fatal("Viper could not read file. Please verify: ", viperErr)
  }
  localJwtViperVar := viper.GetString("jwtBoEnvVar")
  if localJwtViperVar == "" {
    log.Fatal("No local jwt token found. Please use getToken (gt) to obtain a jwt.")
  }
}
