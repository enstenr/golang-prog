package dao

import (
	"bytes"

	"fmt"
	"github.com/enstenr/common/connection"

	"github.com/enstenr/customtypes"

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

func FetchTree(env string) []customtypes.MetadataTree {
	itemArray := make([]customtypes.MetadataTree, 0)
	db := connection.InitConnection(env)
	defer db.Close()

	buf := bytes.NewBufferString(`select "metadataTreeConfigurationId",hash,modified_date,"rePublishStatus" from metadata_tree `)

	rows, err := db.Query(buf.String())

	if err != nil {
		fmt.Print(err)
	}
	defer rows.Close()
	for rows.Next() {

		metadataTree := new(customtypes.MetadataTree)

		_ = rows.Scan(&metadataTree.MetadataTreeConfigurationId, &metadataTree.Hash, &metadataTree.Modified_date, &metadataTree.RePublishStatus)

		itemArray = append(itemArray, *metadataTree)
	}
	return itemArray
}

func FetchTreeConfig(env string) []customtypes.Item {
	itemArray := make([]customtypes.Item, 0)
	db := connection.InitConnection(env)
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
