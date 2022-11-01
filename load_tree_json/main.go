package main

import (
	"connection_module"
 
	"os"
	"strings"

	"github.com/spf13/viper"
)


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

/** This funciton is used to read the gcs report generated after tree job and 
	read the tree name from the report. query the db for the json path 
	copy the json files to a folder in gcs 
**/
func main() {
	env, flag := os.LookupEnv("stage")
	if !flag {
		env = "stage"
	}
	LoadProperties(env);
	bucket_name:=globalViperObj.GetString("BUCKET_NAME")
	treeNameArray:=gcs.ReadFromGCSPath("metadata_tree/reports/2022_10_26-12_31_16/reports.csv",bucket_name)
	metadatatreeArray,_:=connection.ProcessData(treeNameArray,env)

	for _,value:= range(metadatatreeArray){
		//fmt.Print(value.GcsPath)
		GcsPath:=value.GcsPath
		GcsPath=strings.Replace(GcsPath,"gs://content_us","",-1)
		GcsPath=strings.Replace(GcsPath,"/","",1)
	 	gcs.CopyFile("metadata_tree_duplicate_sku/"+env,bucket_name,GcsPath)
	}
}