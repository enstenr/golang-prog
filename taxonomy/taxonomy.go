package main

import (
	 
	"github.com/enstenr/common/dao"
	"fmt"
	"os"

	"github.com/enstenr/customtypes"
)
/** Step 1 - Load L2 and the system category from tree config **/

type Item struct {
	ID    int
	Attrs MetadataTreeConfiguration
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
/**
This funtion is used to populate the table l2_category_mappings
system category id , system category & L2 category L2 category name from 
tree config is fetched and inserted into this table 
**/
func main() {
	env, flag := os.LookupEnv("env")
	if !flag {
		env = "stage"
	}
	 
	
	itemArray:=make([]customtypes.Item,0)
	itemArray=dao.FetchTreeConfig(env)
	fmt.Println(len(itemArray))
	for _,itemObj:=range (itemArray){
		fmt.Print(itemObj.MetadataTreeConfigurationId)
	for _,value:=range(itemObj.Attrs.Config.TreeMappings){
		fmt.Println(value.Category,"\t",value.CategoryID,"\t",value.CategoryNameWithParents)
		}
	}
	//fmt.Println(" going to insert ",len(itemArray))
	//fmt.Println()
	dao.SaveOrUpdate(itemArray,env)
	//Below code is for bulk inserting. But for this scenario the dupliate error may come up and bulk insert will fail. 
	// so going with normal single insert
	//connection.BulkInsert(itemArray,env)
	
}
