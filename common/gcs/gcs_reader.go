package gcs

import (
	 "context"
	"github.com/enstenr/customtypes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
	)
func ReadFromGCSPath(path string,bucket_name string)([]customtypes.DupliateSkuReport){
	treeNameArray := make([]customtypes.DupliateSkuReport, 0)
	
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("/home/bigthinker/mercari/creds/573445696111.json"))
	if err != nil {
		// TODO: Handle error.
		fmt.Print(err)
	}
	 
	bkt := client.Bucket(bucket_name)
	obj := bkt.Object(path)

	r, err := obj.NewReader(ctx)
	csv_reader := csv.NewReader(r)
	 
	for {
		record, err := csv_reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(" error in csv reader")
		}
		dupliateSkuReportObj := customtypes.DupliateSkuReport{}
		dupliateSkuReportObj.Tree_name = record[0]
		dupliateSkuReportObj.Message = record[1]
		dupliateSkuReportObj.State = record[2]
		dupliateSkuReportObj.Status = record[3]

		if record[3] == "Contains Duplicate SKUs" {
			treeNameArray = append(treeNameArray, dupliateSkuReportObj)
		}
	}
	return treeNameArray;

}