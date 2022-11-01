package main

import (
	"github.com/enstenr/common/connection"
	"github.com/enstenr/common/gcs"
	"fmt"
	"log"
	"os"
	"github.com/spf13/viper"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type Article struct {
    Id      string `json:"Id"`
    Title   string `json:"Title"`
    Desc    string `json:"desc"`
    Content string `json:"content"`
}

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var Articles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}


func returnSingleArticle(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    key := vars["id"]

    fmt.Fprintf(w, "Key: " + key)
	for _, article := range Articles {
        if article.Id == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}


func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
   // unmarshal this into a new Article struct
   // append this to our Articles array.    
   reqBody, _ := ioutil.ReadAll(r.Body)
   var article Article 
   json.Unmarshal(reqBody, &article)
   // update our global Articles array to include
   // our new Article
   Articles = append(Articles, article)

   json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
   // once again, we will need to parse the path parameters
   vars := mux.Vars(r)
   // we will need to extract the `id` of the article we
   // wish to delete
   id := vars["id"]

   // we then need to loop through all our articles
   for index, article := range Articles {
	   // if our id path parameter matches one of our
	   // articles
	   if article.Id == id {
		   // updates our Articles array to remove the 
		   // article
		   Articles = append(Articles[:index], Articles[index+1:]...)
	   }
   }

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
 
func main() {
	env, flag := os.LookupEnv("stage")
	if !flag {
		env = "stage"
	}
	LoadProperties(env);
	bucket_name:=globalViperObj.GetString("BUCKET_NAME")
	treeNameArray:=gcs.ReadFromGCSPath("metadata_tree/reports/2022_10_27-12_21_07/reports.csv",bucket_name)
	connection.ProcessData(treeNameArray, env)
	fmt.Println(" Server Started in Port 9090")
	fmt.Print("http://localhost:9090/")

	Articles = []Article{
        Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }
	handleRequests()
	 
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)

	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	
	log.Fatal(http.ListenAndServe(":9090", myRouter))
	
}