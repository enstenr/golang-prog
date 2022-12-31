package connection

import (
	"fmt"
	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/spf13/viper"
)

func GetArangoDBConnection(env string) {

	viperObj := viper.New()

	viperObj.SetConfigName(env)
	viperObj.SetConfigType("env")
	viperObj.AddConfigPath(".")
	viperObj.ReadInConfig()
	host := viperObj.GetString("HOST")
	port := viperObj.GetString("PORT")
	user := viperObj.GetString("USER")
	password := viperObj.GetString("PASSWORD")
	dbname := viperObj.GetString("DBNAME")

	fmt.Print(host, user, port, password, dbname)
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{host + ":" + port + "/"},
	})

	if err != nil {
		fmt.Print(err)
	}
	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(user, password),
	})

	if err != nil {
		fmt.Print(err)
	}

	db, err := client.Database(nil, dbname)
	if err != nil {
		// Handle error
	}
	fmt.Print(db)
	// Open "books" collection
	col, err := db.CollectionExists(nil, "items")
	if err != nil {
		// Handle error
	}
	fmt.Print(col)

}
