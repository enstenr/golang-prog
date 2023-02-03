package connection

import (
	"bytes"
	"database/sql"
	
	"fmt"

	"strings"

	"github.com/enstenr/common/utils"
	"github.com/enstenr/customtypes"
	//"github.com/google/uuid"
	//
	_ "github.com/lib/pq"
)

type Item struct {
	ID                          int
	Attrs                       MetadataTreeConfiguration
	MetadataTreeConfigurationId string
}
type MetadataTreeConfiguration struct {
	Config struct {
		ItemNameSuggest struct {
			Filename  string `json:"filename"`
			GcsUpload string `json:"gcs_upload"`
			Hierarchy []struct {
				EntityID   string `json:"entity_id"`
				EntityName string `json:"entity_name"`
				Order      int64  `json:"order"`
				Prefix     string `json:"prefix"`
				Suffix     string `json:"suffix"`
			} `json:"hierarchy"`
			Score         int64  `json:"score"`
			TreeID        string `json:"tree_id"`
			TreeReference string `json:"tree_reference"`
		} `json:"item_name_suggest"`
		ItemSuggestMapping struct {
			Brand                   string `json:"Brand"`
			BrandID                 int64  `json:"BrandId"`
			Category                string `json:"Category"`
			CategoryID              int64  `json:"CategoryId"`
			CategoryNameWithParents string `json:"CategoryNameWithParents"`
		} `json:"item_suggest_mapping"`
		MetadataTree struct {
			Criteria []struct {
				Attribute   string `json:"attribute"`
				AttributeID string `json:"attribute_id"`
				Entity      string `json:"entity"`
				EntityID    string `json:"entity_id"`
			} `json:"criteria"`
			Hierarchy []struct {
				EntityID   string `json:"entity_id"`
				EntityName string `json:"entity_name"`
				Required   bool   `json:"required"`
			} `json:"hierarchy"`
		} `json:"metadata_tree"`
		TreeMappings []struct {
			Brand                   string `json:"Brand"`
			BrandID                 int64  `json:"BrandId"`
			Category                string `json:"Category"`
			CategoryID              int64  `json:"CategoryId"`
			CategoryNameWithParents string `json:"CategoryNameWithParents"`
			Keyword                 string `json:"Keyword"`
		} `json:"tree_mappings"`
	} `json:"config"`
}


func InitConnection(env string) (db *sql.DB) {
	return GetConnection(env)
}

func ProcessData(duplicateSkuReportObjArray []customtypes.DupliateSkuReport, env string) ([]customtypes.MetadataTree, error) {

	db := InitConnection(env)
	defer db.Close()

	buf := bytes.NewBufferString(`select metadata_tree."metadataTreeId" ,metadata_tree_configuration."metadataTreeConfigurationId", metadata_tree_configuration.name ,metadata_tree."gcsPath" from metadata_tree,metadata_tree_configuration	where metadata_tree.published=true and metadata_tree."metadataTreeConfigurationId" =metadata_tree_configuration."metadataTreeConfigurationId" and metadata_tree.active = true	 and metadata_tree_configuration.name in (`)

	for i, v := range duplicateSkuReportObjArray {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString("'")
		buf.WriteString(v.Tree_name)
		buf.WriteString("'")

	}
	buf.WriteString(")")
	fmt.Print(buf.String())
	rows, err := db.Query(buf.String())

	if err != nil {
		fmt.Print(err)
	}

	metadatatreeArray := make([]customtypes.MetadataTree, 0)

	treeMap := formTreeMap(duplicateSkuReportObjArray)
	defer rows.Close()
	for rows.Next() {
		metadataTreeObj := customtypes.MetadataTree{}

		_ = rows.Scan(&metadataTreeObj.MetadataTreeId, &metadataTreeObj.MetadataTreeConfigurationId, &metadataTreeObj.Name, &metadataTreeObj.GcsPath)

		duplicateSkuReportObj := treeMap[metadataTreeObj.Name]
		message := utils.CleanUpMessage(duplicateSkuReportObj.Message, metadataTreeObj.Name)

		skuIdCount := len(strings.Split(message, " "))
		metadataTreeObj.Message = message
		metadataTreeObj.Count = skuIdCount
		metadatatreeArray = append(metadatatreeArray, metadataTreeObj)

		writeToCSV(metadatatreeArray)

	}
	fmt.Println("Total Trees published with duplicate SKU ", len(metadatatreeArray))
	return metadatatreeArray, nil
}


func getTreeByName(inputTreeName string, duplicateSkuReportObjArray []customtypes.DupliateSkuReport) customtypes.DupliateSkuReport {
	dupliateSkuReportObj := customtypes.DupliateSkuReport{}

	for _, duplicateSkuReportObj := range duplicateSkuReportObjArray {
		if duplicateSkuReportObj.Tree_name == inputTreeName {
			return duplicateSkuReportObj
		}
	}
	return dupliateSkuReportObj
}



/* func getSystemCategory( critArray []Item.Attrs.Config.MetadataTree.Criteria)(attribute string,attributeId string){
	for _,value := range (critArray){
		if value.Entity=="system_category" {
			return value.Attribute, value. AttributeID
		}
	}

 }

*/