package connection

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"strings"

	"github.com/enstenr/common/utils"
	"github.com/enstenr/customtypes"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
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

func (a *MetadataTreeConfiguration) Scan(value interface{}) error {

	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
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

func formTreeMap(duplicateSkuReportObjArray []customtypes.DupliateSkuReport) map[string]customtypes.DupliateSkuReport {
	treeMap := make(map[string]customtypes.DupliateSkuReport, 0)
	for _, duplicateSkuReportObj := range duplicateSkuReportObjArray {
		treeMap[duplicateSkuReportObj.Tree_name] = duplicateSkuReportObj

	}
	return treeMap
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

func FetchTreeConfig(env string) []customtypes.Item {
	itemArray := make([]customtypes.Item, 0)
	db := InitConnection(env)
	defer db.Close()

	buf := bytes.NewBufferString(`select config,"metadataTreeConfigurationId" from metadata_tree_configuration where active=true`)

	rows, err := db.Query(buf.String())

	if err != nil {
		fmt.Print(err)
	}
	defer rows.Close()
	for rows.Next() {

		item := new(Item)
		itemCustomType := new(customtypes.Item)

		_ = rows.Scan(&item.Attrs, &item.MetadataTreeConfigurationId)
		copier.Copy(&itemCustomType, &item)
		itemArray = append(itemArray, *itemCustomType)
	}
	return itemArray
}

func SaveOrUpdate(itemArray []customtypes.Item, env string) {

	db := InitConnection(env)
	defer db.Close()

	errorCount := 0
	for _, itemObj := range itemArray {
		metadata_tree_configuration_id := itemObj.MetadataTreeConfigurationId
		for _, value := range itemObj.Attrs.Config.TreeMappings {
			//system_category
			var system_category_id string
			for _, criteriaObj := range itemObj.Attrs.Config.MetadataTree.Criteria {
				if criteriaObj.Entity == "system_category" {
					//system_category=criteriaObj.Attribute
					system_category_id = criteriaObj.AttributeID
				}
			}
			system_category_l2_category_id := strings.Replace(uuid.New().String(), "-", "", -1)

			//system_category_l2_category_id:=utils.GetSHA1(fmt.Sprint(system_category_id,value.CategoryID))
			//criteriaArray:=itemObj.Attrs.Config.MetadataTree.Criteria
			//system_category,system_categoryId:=getSystemCategory(criteriaArray)
			sqlStatement := `
	INSERT INTO l2_category_mappings ("l2CategoryMappingsId",system_category_id,l2_category_id,active,metadata_tree_configuration_id)
	VALUES ($1,$2,$3,$4,$5)`
			//fmt.Print(system_category_l2_category_id,system_category, system_category_id,value.CategoryID,value.Category)
			_, err := db.Exec(sqlStatement, system_category_l2_category_id, system_category_id, value.CategoryID, true, metadata_tree_configuration_id)
			if err != nil {
				fmt.Print(err)
				errorCount++
				continue
			}

		}

		fmt.Print(len(itemArray))
		fmt.Print(errorCount)
	}
}

/* func getSystemCategory( critArray []Item.Attrs.Config.MetadataTree.Criteria)(attribute string,attributeId string){
	for _,value := range (critArray){
		if value.Entity=="system_category" {
			return value.Attribute, value. AttributeID
		}
	}

 }

*/

func BulkInsert(unsavedRows []customtypes.Item, env string) error {
	db := InitConnection(env)
	defer db.Close()

	valueStrings := make([]string, 0, len(unsavedRows))
	valueArgs := make([]interface{}, 0, len(unsavedRows)*5)
	i := 0
	for _, itemObj := range unsavedRows {

		for _, value := range itemObj.Attrs.Config.TreeMappings {

			var system_category, system_category_id string
			for _, criteriaObj := range itemObj.Attrs.Config.MetadataTree.Criteria {
				if criteriaObj.Entity == "system_category" {
					system_category = criteriaObj.Attribute
					system_category_id = criteriaObj.AttributeID
				}
			}

			valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d,$%d,$%d)", i*5+1, i*5+2, i*5+3, i*5+4, i*5+5))
			valueArgs = append(valueArgs, system_category)
			valueArgs = append(valueArgs, system_category_id)
			valueArgs = append(valueArgs, value.CategoryID)
			valueArgs = append(valueArgs, value.Category)
			valueArgs = append(valueArgs, true)
			i++
		}
	}
	stmt := fmt.Sprintf("INSERT INTO l2_category_mappings (system_category,system_category_id,l2_category_id,l2_category,active) VALUES %s", strings.Join(valueStrings, ","))
	_, err := db.Exec(stmt, valueArgs...)

	if err != nil {
		fmt.Print(err)

	}

	return err
}

func SaveOrUpdateL2CategoryEntities(env string, l2_category_mappings map[int64]map[string]string) {

	db := InitConnection(env)
	defer db.Close()

	for l2_category_id, entityMap := range l2_category_mappings {

		for entity_name, categoryName := range entityMap {
			fmt.Println(l2_category_id, categoryName, entity_name)
			sqlStatement := `
	INSERT INTO l2_category_entities ("l2_category_id",l2_category_name,entity_name)
	VALUES ($1,$2,$3)`
			value, err := db.Exec(sqlStatement, l2_category_id, categoryName, entity_name)
			if err != nil {
				fmt.Print(err)

				continue
			}
			fmt.Print(value)

		}

	}
}
