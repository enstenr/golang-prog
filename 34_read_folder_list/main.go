package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"strings"
)

func main(){
	fs,err:=ioutil.ReadDir(".")
	if err!= nil {
        fmt.Println(err)
	}
	for _,f:=range fs{
	fmt.Println(f.Name())
	}

	filepath.WalkDir(".", 
	func(path string, info os.DirEntry, err error)error{
		fmt.Print(path)
		return nil
	})

	f1,_:=os.Open("/home/bigthinker/mercari/repo/mercari-us-de-rdbms-utils/snapshots/envs/canary/metadatadb/schema")
	fs1,_:=f1.ReadDir(-1)
	for _,f2:=range fs1{

		tempfullPath:=[]string{f1.Name(),string(os.PathSeparator), f2.Name()}
		fullPath:=strings.Join(tempfullPath,"")
		info,_:=	os.Stat(fullPath)
		fmt.Println(getFileSize(info.Size()),"\t",f2.Name())
    if !strings.HasPrefix(f2.Name(),"backup_"){
		//fmt.Println(f2.Name())
	
	}else{
		
		fmt.Println("Deleting : ",f2.Name())
		os.Remove(fullPath)
		 
	}
	}
}

func getFileSize(sizeInBytes int64) string  {
	 
    var kilobyte int64= 1024
    var megabyte int64= kilobyte * kilobyte
    var gigabyte int64= megabyte * megabyte
	//var terabyte int64= gigabyte * gigabyte
	//var petabyte int64= terabyte * terabyte
	//var exabyte int64= petabyte * petabyte
	//var zettabyte int64= exabyte * exabyte
	//var yottabyte int64= zettabyte * zettabyte
	 
	returnString:=[]string{}
	if sizeInBytes< kilobyte{ 
		returnString=append(returnString,strconv.Itoa(int(sizeInBytes))," bytes")
		 
	}
	if sizeInBytes>=kilobyte && sizeInBytes< megabyte{
		 
		returnString=append(returnString,strconv.Itoa(int(sizeInBytes)/1000)," KB")
		 
	}
	 if sizeInBytes>=megabyte && sizeInBytes< gigabyte{
		returnString=append(returnString,strconv.Itoa(int(sizeInBytes)/1000)," MB")
	}
	if sizeInBytes>gigabyte{
		returnString=append(returnString,strconv.Itoa(int(sizeInBytes)/1000)," GB")
	}
	/* int64 needs to be converted to bigint
	if sizeInBytes>terabyte{
		fmt.Print(terabyte,"here",sizeInBytes)
		returnString=append(returnString,strconv.Itoa(int(sizeInBytes)/1000))
	}
	if sizeInBytes>petabyte{
		
		returnString=append(returnString,strconv.Itoa(int(sizeInBytes)/1000))
	}
	if sizeInBytes>exabyte{
		returnString=append(returnString,strconv.Itoa(int(sizeInBytes)/1000))
	}
	if sizeInBytes>zettabyte{
		returnString=append(returnString,strconv.Itoa(int(sizeInBytes)/1000))
	}
	if sizeInBytes>yottabyte{
		returnString=append(returnString,strconv.Itoa(int(sizeInBytes)/1000))
	}*/
	return strings.Join(returnString,"")
}