package connection

import (
	"database/sql"
	"fmt"	
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)
func GetConnection(env string)(db *sql.DB)  {
	viperObj:=viper.New()
 
	viperObj.SetConfigName(env)
	viperObj.SetConfigType("env")
	viperObj.AddConfigPath(".")
	viperObj.ReadInConfig()
	host     := viperObj.Get("HOST")
	port:=viperObj.Get("PORT")
	user     := viperObj.Get("USER")
	password:=viperObj.Get("PASSWORD")
	dbname:=viperObj.Get("DBNAME")
	fmt.Print(host,port,user,password,dbname)
	postgresqlDbInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", postgresqlDbInfo)

	if err != nil {
		panic(err)
	}
	

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Print("opened connection")
	 return db;
}


func GetDBConnection(env string,dbname string)(db *sql.DB)  {
	viperObj:=viper.New()
 
	viperObj.SetConfigName(env)
	viperObj.SetConfigType("env")
	viperObj.AddConfigPath(".")
	viperObj.ReadInConfig()
	host     := viperObj.Get("HOST")
	port:=viperObj.Get("PORT")
	user     := viperObj.Get("USER")
	password:=viperObj.Get("PASSWORD")
	//fmt.Print(host,port,user,password,dbname)
	postgresqlDbInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", postgresqlDbInfo)

	if err != nil {
		panic(err)
	}
	

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	//fmt.Print("opened connection")
	 return db;
}