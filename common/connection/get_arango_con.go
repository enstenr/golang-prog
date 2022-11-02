package connection
import (
	"github.com/arangodb/go-driver/http"
	driver "github.com/arangodb/go-driver"
)
func GetArangoDBConnection(env string) (*driver.Client){
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://localhost:8529"},
	})

	if err != nil {
		// Handle error
	}
	client	nt, err := driver.NewClient(driver.ClientConfig{
		Connection: conn,
	})
	if err != nil {
		// Handle error
	}
}