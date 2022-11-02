package main

import (
	"os"

	"github.com/enstenr/common/connection"
	"github.com/spf13/viper"
)
func main() {
	env, flag := os.LookupEnv("dev")
	if !flag {
		env = "dev"
	}
	LoadProperties(env);
	connection.GetArangoDBConnection(env)
	
}
var globalViperObj *viper.Viper;
func LoadProperties(env string)( *viper.Viper){
	viperObj1 :=viper.New();
	viperObj1.SetConfigName(env)
	viperObj1.SetConfigType("env")
	viperObj1.AddConfigPath(".")
	viperObj1.ReadInConfig()
	globalViperObj=viperObj1
	return viperObj1
}