package main

import (
	"bufio"
	"errors"
	//"github.com/gotk3/gotk3/gtk"
	"fmt"
	"io/ioutil"
	"log"
	mrand "math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
	"golang.org/x/oauth2/google"
)
 
var globalViperObj *viper.Viper

func LoadProperties(env string) *viper.Viper {
	viperObj1 := viper.New()
	viperObj1.SetConfigName(env)
	viperObj1.SetConfigType("env")
	viperObj1.AddConfigPath(".")
	viperObj1.ReadInConfig()
	globalViperObj = viperObj1
	return viperObj1
}

func main() {

	 


	var errorMap = make(map[int8]string)
	errorMap[1]=" Please select an option from the list "
	errorMap[2]=" Please select a number from the list "
	errorMap[3]=" You have selected invalid option "
	errorMap[0]=" I guess you are having trouble selecting a number. Please try again. "
	
	prop_env, flag := os.LookupEnv("env")
	if !flag {
		prop_env = "env"
	}
	LoadProperties(prop_env)

	envMap := make(map[int]string)
	envMap[1] = "DEV"
	envMap[2] = "STAGE"
	envMap[3] = "CANARY"
	envMap[4] = "PROD"
	var AUDIENCE = ""
	//colorReset := "\033[0m"

	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	colorCyan := "\033[36m"
	colorBlue := "\033[34m"
	/*colorPurple := "\033[35m"

	  colorWhite := "\033[37m" */
	fmt.Println()
	fmt.Println(string(colorGreen), "Google API token Generator")
	fmt.Println()
	fmt.Println(string(colorYellow), "Place your cred file in same directory along with the program so that it will be auto picked. ")
	fmt.Println()
	
	reader := bufio.NewReader(os.Stdin)


	fileInfoArray,_:=ioutil.ReadDir(".")

	  secret_file:=""
	jsonfile_array:=make([]string,0)
	for _,value :=range fileInfoArray{
		secret_file=value.Name()
		ok,_:=validateFile(value.Name())
		if(ok){
			jsonfile_array=append(jsonfile_array,secret_file)
		}
		
	}
	/** if more than one file is present in the current directory then 
	proceed with if blok. else go to else block **/
	if(len(jsonfile_array)>0){
		 
		if(len(jsonfile_array)>1){
			//TODO
			for index,fileName := range jsonfile_array{
				fmt.Print(index,fileName)
			}
		}else{
			fileName:=jsonfile_array[0]
			fmt.Printf("Found a json file %s . \n\nCan i use it for connection ? (y/n) : ",fileName)
			choice, _ := reader.ReadString('\n')
			if(strings.Compare(choice,"y")>0){
			secret_file=fileName
			}else if(strings.Compare(choice,"n")>0){
				fmt.Print(" \nPlease enter a JSON file path : ")
				secret_file, _ = reader.ReadString('\n')
				 
					ok,_:=validateFile(secret_file)
					if(!ok){
						fmt.Println(" Not a Valid JSON File. Try again. Exiting Program. see you soon. ")
							os.Exit(600)
					}
				}
		}
		
	}else{
		fmt.Print("No JSON file found in current directory. Please enter a JSON file path")
		secret_file, _ = reader.ReadString('\n')
		
		ok,_:=validateFile(secret_file)
		if(!ok){
			fmt.Println(" Not a Valid JSON File. Try again. Exiting Program. see you soon. ")
				os.Exit(600)
		}
	}	 
	 
	 
	secret_file=strings.TrimRight(secret_file,"\n")
		secret_file=strings.Trim(secret_file, "")
	data, err := ioutil.ReadFile(secret_file)
	if err != nil {
		 
		log.Fatal(err)
	}

	
	fmt.Println(string(colorCyan), "Environment Options:  \n 1: dev \n 2: stage \n 3: canary \n 4: prod \n Ctrl-C to Quit")
	fmt.Println()
	fmt.Print(string(colorRed), " Enter an option (number) and hit Enter key :")
	var env string
	var flag1 = false
	for !flag1 {
		env, _ := reader.ReadString('\n')
		fmt.Println()

		// Trimming the whitespace and newlines
		env = strings.TrimRight(env, "\n")
		env = strings.Trim(env, "")
		intval, err := strconv.Atoi(env)
		if err != nil {
			//panic(" Please select a valid option")
			
			mrand.Seed(time.Now().Unix())
			wi:=mrand.Intn(len(errorMap))
			fmt.Print(errorMap[int8(wi)])
			flag1=false
		}else{
			flag1=true
		}
		AUDIENCE = globalViperObj.GetString(envMap[intval])
	}

	/* switch env {
	case "1":
		AUDIENCE = globalViperObj.GetString("DEV")
	case "2":
		AUDIENCE = globalViperObj.GetString("STAGE")
	} */

	/* if env == "dev" {
		AUDIENCE = DEV_AUDIENCE
	} else if env == "stage" {
		AUDIENCE = STAGE_AUDIENCE
	} */
	fmt.Println(string(colorBlue), "Environment is :", env)
	fmt.Println(string(colorBlue), "Target Audience is :", AUDIENCE)
	// Your credentials should be obtained from the Google
	// Developer Console (https://console.developers.google.com).
	// Navigate to your project, then see the "Credentials" page
	// under "APIs & Auth".
	// To create a service account client, click "Create new Client ID",
	// select "Service Account", and click "Create Client ID". A JSON
	// key file will then be downloaded to your computer.
	//"/home/bigthinker/mercari/creds/573445696111.json"
	

	conf1, err := google.JWTAccessTokenSourceFromJSON(data, AUDIENCE)

	if err != nil {
		log.Fatal(err)
	}
	token, _ := conf1.Token()
	fmt.Println(" Copy the below token for sending the API request")
	fmt.Println("------------------------------------------------")
	fmt.Println(string(colorYellow), token.AccessToken)
	// Initiate an http.Client. The following GET request will be
	// authorized and authenticated on the behalf of
	// your service account.
	//client1 := conf1.Client(oauth2.NoContext)
	//client1.Get("...")
	fmt.Println("------------------------------------------------")

}
/* func gtkCodes(){
 // Initialize GTK without parsing any command line arguments.
 gtk.Init(nil)
 // Create a new toplevel window, set its title, and connect it to the
    // "destroy" signal to exit the GTK main loop when it is destroyed.
    win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
} */

func validateFile(file string) (bool,error) {
	file=strings.TrimSpace(file)
	 
	if (!strings.HasSuffix(file,".json")){
		// panic(" Invalid JSON file")
		return false, errors.New("Invalid JSON file")
	}
	return true,nil
}
