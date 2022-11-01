package gcs

import (
	"context"
	"fmt"
	"time"


	"cloud.google.com/go/storage" 
	"google.golang.org/api/option"
)

func CopyFile(dstFolder, srcBucket, srcObject string)  error {
	
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("/home/bigthinker/mercari/creds/573445696111.json"))

	 
	if err != nil {
			return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()
	 
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	 
	dstObject := dstFolder+"/"+srcObject
	src := client.Bucket(srcBucket).Object(srcObject)
	dst := client.Bucket(srcBucket).Object(dstObject)
	fmt.Print(" Source :")
	fmt.Print(src)
	fmt.Println(" Destiniation ")
	fmt.Print(dst)
	//dst = dst.If(storage.Conditions{DoesNotExist: true})
	 
	if _, err := dst.CopierFrom(src).Run(ctx); err != nil {
		 fmt.Print(err)
	 
}
fmt.Print( "Blob %v in bucket %v copied to blob %v in bucket %v.\n", srcObject, srcBucket, dstObject, dstFolder)
return nil


}