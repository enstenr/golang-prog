package connection

import (
	"encoding/csv"
	"fmt"
	"strconv"

	"github.com/enstenr/customtypes"

	"log"
	"os"
)

func writeToCSV(metadataTreeArray []customtypes.MetadataTree) {

	csvFile, err := os.Create("employee1.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)
	value := []string{"MetadataTreeConfigurationId", "Tree Name", "Duplicate SKU ID ", "Count of Duplicate SKUs"}
	csvwriter.Write(value)
	for _, empRow := range metadataTreeArray {
		value = []string{empRow.MetadataTreeConfigurationId, empRow.Name, empRow.Message, strconv.Itoa(empRow.Count), empRow.GcsPath}

		_ = csvwriter.Write(value)
	}
	csvwriter.Flush()
	csvFile.Close()
}


func WriteTreeToCSV(metadataTreeArray []customtypes.MetadataTreeDifference,fileName string) {

	csvFile, err := os.Create(fileName)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	fmt.Print(*csvFile)
	csvwriter := csv.NewWriter(csvFile)
	value := []string{"Canary MetadataTreeConfigurationId","Stage MetadataTreeConfigurationId", "Tree Name", " Canary Hash", "Stage Hash"," Canary Modified Date","Stage Modified Date","Canary Republish Status","Stage Republish Status"}
	csvwriter.Write(value)
	for _, empRow := range metadataTreeArray {
		value = []string{empRow.MetadataTreeConfigurationId1,empRow.MetadataTreeConfigurationId2, empRow.Name,empRow.Hash1,empRow.Hash2,empRow.Modified_date1,empRow.Modified_date2,empRow.RePublishStatus1,empRow.RePublishStatus2}

		_ = csvwriter.Write(value)
	}
	csvwriter.Flush()
	csvFile.Close()
}

