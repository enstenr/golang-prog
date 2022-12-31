package connection

import (
	"encoding/csv"
	"github.com/enstenr/customtypes"
	"strconv"

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
