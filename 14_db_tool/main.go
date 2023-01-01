package main

import (
	"flag"
	"fmt"

 
	"github.com/bmuschko/kubectl-server-version/cmd"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"os"
 
)

func main() {
	fmt.Printf("Welcome to the LinkedIn Learning Temperature Service!\n\n")
	beachReady := flag.Bool("beach", false, "Display only beach ready destinations")
	skiReady := flag.Bool("ski", false, "Display only ski ready destinations")
	month := flag.Int("month", 0, "Look up for destinations in a given month [1,12]")
	name := flag.String("name", "", "Look up destinations by name")
	flag.Parse()
	fmt.Print(*beachReady,*skiReady,*month,*name)
}