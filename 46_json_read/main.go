package main

import (
	"encoding/json"
	"fmt"
)

type SPAMetadataTreeConfigurationParsed struct {
	MetadataTreeId             string                     `json:"metadataTreeId,omitempty"`
	Active                     bool                       `json:"active"`
	Hash                       string                     `json:"hash,omitempty"`
	Loaded                     bool                       `json:"loaded,omitempty"`
	GcsPath                    string                     `json:"gcsPath,omitempty"`
	MetdataTreeConfigurationId string                     `json:"metadataTreeConfigurationId"`
	KgExportGcsPath            string                     `json:"kg_export_gcs_path,omitempty"`
	InsGcsPath                 string                     `json:"ins_gcs_path,omitempty"`
	ExternalTreeId             string                     `json:"externalTreeId,omitempty"`
	Published                  bool                       `json:"published"`
	CreateDate                 string                     `json:"create_date"`
	Metadata_tree_id           int64                      `json:"metadata_tree_id,omitempty"`
	ModifiedDate               string                     `json:"modified_date"`
	RePublishStatus            string                     `json:"rePublishStatus,omitempty"`
	ConfigRoot                 *SPAMetadataTreeConfigRoot `json:"config"`
	IsDraft                    bool                       `json:"is_draft"`
	IsGenericTree              bool                       `json:"is_generic_tree"`
	IsLargeTree                bool                       `json:"is_large_tree"`
	Name                       string                     `json:"name"`
	Status                     string                     `json:"status"`
}

type SPAMetadataTreeConfigRoot struct {
	Config *SPAMetadataTreeConfig `json:"config"`
}

type SPAItemSuggestMapping struct {
	Brand                   string `json:"Brand"`
	BrandId                 int    `json:"BrandId"`
	Category                string `json:"Category"`
	CategoryId              int    `json:"CategoryId"`
	CategoryNameWithParents string `json:"CategoryNameWithParents"`
}

type SPAMetadataTree struct {
	Criteria  []SPAMetadataTreeCriteria  `json:"criteria"`
	Hierarchy []SPAMetadataTreeHierarchy `json:"hierarchy"`
}
type SPAMetadataTreeHierarchy struct {
	EntityId   string `json:"entity_id"`
	EntityName string `json:"entity_name"`
	Required   bool   `json:"required"`
}
type SPAMetadataTreeCriteria struct {
	Attribute   string `json:"attribute"`
	AttributeId string `json:"attribute_id"`
	Entity      string `json:"entity"`
	EntityId    string `json:"entity_id"`
}

type SPAMetadataTreeConfig struct {
	ItemNameSuggest            *SPAMetadataTreeItemNameSuggestion `json:"item_name_suggest"`
	ItemSuggestMapping         *SPAItemSuggestMapping             `json:"item_suggest_mapping"`
	MetadataTree               *SPAMetadataTree                   `json:"metadata_tree"`
	UniqueEntities             []SPAMetadataTreeUniqueEntity      `json:"unique_entities"`
	
}
 
type SPAMetadataTreeItemNameSuggestion struct {
	Filename        string                            `json:"filename"`
	GcsUpload       string                            `json:"gcs_upload"` // note <- this is only a string bc it's represented as "True" in db (won't implicitly convert)
	Hierarchy       []SPAMetadataTreeSuggestHierarchy `json:"hierarchy"`
	Score           int                               `json:"score"`
	TreeId          string                            `json:"tree_id"`
	TreeReference   string                            `json:"tree_reference"`
	UniversalPrefix string                            `json:"universal_prefix,omitempty"`
	UniversalSuffix string                            `json:"universal_suffix,omitempty"`
	JobId           string                            `json:"job_id,omitempty"`
}

type SPAMetadataTreeSuggestHierarchy struct {
	EntityId   string `json:"entity_id"`
	EntityName string `json:"entity_name"`
	Order      int    `json:"order"`
	Prefix     string `json:"prefix"`
	Suffix     string `json:"suffix"`
}

type SPAMetadataTreeUniqueEntity struct {
	EntityId      string `json:"entityId"`
	EntityName    string `json:"entityName"`
	IsSearchFacet bool   `json:"is_search_facet"`
}
func main() {
	bytearray := []byte{123, 34, 97, 99, 116, 105, 118, 101, 34, 58, 116, 114, 117, 101, 44, 34, 109, 101, 116, 97, 100, 97, 116, 97, 84, 114, 101, 101, 67, 111, 110, 102, 105, 103, 117, 114, 97, 116, 105, 111, 110, 73, 100, 34, 58, 34, 53, 53, 54, 52, 100, 102, 101, 51, 53, 99, 57, 49, 52, 101, 102, 98, 97, 52, 55, 102, 102, 51, 49, 51, 48, 56, 54, 100, 57, 54, 48, 99, 34, 44, 34, 112, 117, 98, 108, 105, 115, 104, 101, 100, 34, 58, 116, 114, 117, 101, 44, 34, 99, 114, 101, 97, 116, 101, 95, 100, 97, 116, 101, 34, 58, 34, 50, 48, 50, 48, 45, 48, 51, 45, 48, 50, 84, 49, 51, 58, 50, 54, 58, 51, 53, 46, 51, 54, 50, 54, 56, 50, 90, 34, 44, 34, 109, 111, 100, 105, 102, 105, 101, 100, 95, 100, 97, 116, 101, 34, 58, 34, 50, 48, 50, 50, 45, 49, 49, 45, 49, 53, 84, 48, 56, 58, 48, 54, 58, 52, 55, 46, 53, 48, 48, 48, 56, 53, 90, 34, 44, 34, 99, 111, 110, 102, 105, 103, 34, 58, 123, 34, 99, 111, 110, 102, 105, 103, 34, 58, 123, 34, 105, 116, 101, 109, 95, 110, 97, 109, 101, 95, 115, 117, 103, 103, 101, 115, 116, 34, 58, 123, 34, 102, 105, 108, 101, 110, 97, 109, 101, 34, 58, 34, 108, 105, 115, 116, 105, 110, 103, 45, 110, 97, 109, 101, 45, 115, 117, 103, 103, 101, 115, 116, 105, 111, 110, 115, 47, 110, 105, 110, 116, 101, 110, 100, 111, 95, 103, 97, 109, 101, 115, 46, 99, 115, 118, 34, 44, 34, 103, 99, 115, 95, 117, 112, 108, 111, 97, 100, 34, 58, 34, 84, 114, 117, 101, 34, 44, 34, 104, 105, 101, 114, 97, 114, 99, 104, 121, 34, 58, 91, 123, 34, 101, 110, 116, 105, 116, 121, 95, 105, 100, 34, 58, 34, 49, 97, 97, 48, 100, 100, 53, 101, 45, 57, 101, 48, 52, 45, 52, 50, 48, 50, 45, 97, 101, 102, 101, 45, 55, 50, 102, 101, 48, 57, 49, 51, 102, 102, 97, 101, 34, 44, 34, 101, 110, 116, 105, 116, 121, 95, 110, 97, 109, 101, 34, 58, 34, 80, 108, 97, 116, 102, 111, 114, 109, 34, 44, 34, 111, 114, 100, 101, 114, 34, 58, 50, 44, 34, 112, 114, 101, 102, 105, 120, 34, 58, 34, 102, 111, 114, 32, 34, 44, 34, 115, 117, 102, 102, 105, 120, 34, 58, 34, 34, 125, 44, 123, 34, 101, 110, 116, 105, 116, 121, 95, 105, 100, 34, 58, 34, 100, 53, 52, 57, 51, 57, 53, 100, 45, 98, 53, 98, 100, 45, 52, 97, 55, 55, 45, 57, 48, 99, 49, 45, 53, 57, 49, 51, 52, 50, 97, 54, 97, 51, 57, 51, 34, 44, 34, 101, 110, 116, 105, 116, 121, 95, 110, 97, 109, 101, 34, 58, 34, 77, 111, 100, 101, 108, 34, 44, 34, 111, 114, 100, 101, 114, 34, 58, 49, 44, 34, 112, 114, 101, 102, 105, 120, 34, 58, 34, 34, 44, 34, 115, 117, 102, 102, 105, 120, 34, 58, 34, 34, 125, 93, 44, 34, 115, 99, 111, 114, 101, 34, 58, 49, 48, 48, 48, 48, 44, 34, 116, 114, 101, 101, 95, 105, 100, 34, 58, 34, 34, 44, 34, 116, 114, 101, 101, 95, 114, 101, 102, 101, 114, 101, 110, 99, 101, 34, 58, 34, 110, 105, 110, 116, 101, 110, 100, 111, 95, 103, 97, 109, 101, 115, 34, 125, 44, 34, 105, 116, 101, 109, 95, 115, 117, 103, 103, 101, 115, 116, 95, 109, 97, 112, 112, 105, 110, 103, 34, 58, 123, 34, 66, 114, 97, 110, 100, 34, 58, 34, 78, 105, 110, 116, 101, 110, 100, 111, 34, 44, 34, 66, 114, 97, 110, 100, 73, 100, 34, 58, 52, 53, 57, 49, 44, 34, 67, 97, 116, 101, 103, 111, 114, 121, 34, 58, 34, 71, 97, 109, 101, 115, 34, 44, 34, 67, 97, 116, 101, 103, 111, 114, 121, 73, 100, 34, 58, 55, 57, 54, 44, 34, 67, 97, 116, 101, 103, 111, 114, 121, 78, 97, 109, 101, 87, 105, 116, 104, 80, 97, 114, 101, 110, 116, 115, 34, 58, 34, 69, 108, 101, 99, 116, 114, 111, 110, 105, 99, 115, 32, 92, 117, 48, 48, 51, 101, 32, 86, 105, 100, 101, 111, 32, 103, 97, 109, 101, 115, 32, 92, 117, 48, 48, 50, 54, 32, 99, 111, 110, 115, 111, 108, 101, 115, 32, 92, 117, 48, 48, 51, 101, 32, 71, 97, 109, 101, 115, 34, 125, 44, 34, 109, 101, 116, 97, 100, 97, 116, 97, 95, 116, 114, 101, 101, 34, 58, 123, 34, 99, 114, 105, 116, 101, 114, 105, 97, 34, 58, 91, 123, 34, 97, 116, 116, 114, 105, 98, 117, 116, 101, 34, 58, 34, 86, 105, 100, 101, 111, 32, 71, 97, 109, 101, 115, 34, 44, 34, 97, 116, 116, 114, 105, 98, 117, 116, 101, 95, 105, 100, 34, 58, 34, 55, 101, 50, 101, 57, 48, 48, 53, 45, 48, 51, 50, 49, 45, 52, 51, 54, 55, 45, 98, 49, 100, 50, 45, 57, 99, 98, 54, 55, 100, 97, 102, 49, 55, 53, 97, 34, 44, 34, 101, 110, 116, 105, 116, 121, 34, 58, 34, 115, 121, 115, 116, 101, 109, 95, 99, 97, 116, 101, 103, 111, 114, 121, 34, 44, 34, 101, 110, 116, 105, 116, 121, 95, 105, 100, 34, 58, 34, 98, 57, 51, 53, 54, 57, 53, 102, 45, 101, 52, 55, 57, 45, 52, 102, 55, 101, 45, 97, 97, 53, 52, 45, 98, 49, 55, 57, 54, 54, 102, 100, 51, 102, 97, 56, 34, 125, 44, 123, 34, 97, 116, 116, 114, 105, 98, 117, 116, 101, 34, 58, 34, 78, 105, 110, 116, 101, 110, 100, 111, 34, 44, 34, 97, 116, 116, 114, 105, 98, 117, 116, 101, 95, 105, 100, 34, 58, 34, 57, 48, 100, 99, 56, 55, 55, 57, 45, 52, 99, 53, 52, 45, 52, 101, 52, 49, 45, 56, 101, 100, 48, 45, 101, 101, 50, 48, 102, 54, 56, 50, 52, 54, 100, 48, 34, 44, 34, 101, 110, 116, 105, 116, 121, 34, 58, 34, 98, 114, 97, 110, 100, 34, 44, 34, 101, 110, 116, 105, 116, 121, 95, 105, 100, 34, 58, 34, 50, 97, 102, 53, 51, 52, 52, 52, 45, 48, 49, 55, 49, 45, 52, 56, 55, 100, 45, 57, 49, 51, 51, 45, 50, 48, 56, 57, 53, 48, 51, 56, 56, 101, 99, 55, 34, 125, 93, 44, 34, 104, 105, 101, 114, 97, 114, 99, 104, 121, 34, 58, 91, 123, 34, 101, 110, 116, 105, 116, 121, 95, 105, 100, 34, 58, 34, 49, 97, 97, 48, 100, 100, 53, 101, 45, 57, 101, 48, 52, 45, 52, 50, 48, 50, 45, 97, 101, 102, 101, 45, 55, 50, 102, 101, 48, 57, 49, 51, 102, 102, 97, 101, 34, 44, 34, 101, 110, 116, 105, 116, 121, 95, 110, 97, 109, 101, 34, 58, 34, 112, 108, 97, 116, 102, 111, 114, 109, 34, 44, 34, 114, 101, 113, 117, 105, 114, 101, 100, 34, 58, 116, 114, 117, 101, 125, 44, 123, 34, 101, 110, 116, 105, 116, 121, 95, 105, 100, 34, 58, 34, 100, 53, 52, 57, 51, 57, 53, 100, 45, 98, 53, 98, 100, 45, 52, 97, 55, 55, 45, 57, 48, 99, 49, 45, 53, 57, 49, 51, 52, 50, 97, 54, 97, 51, 57, 51, 34, 44, 34, 101, 110, 116, 105, 116, 121, 95, 110, 97, 109, 101, 34, 58, 34, 109, 111, 100, 101, 108, 34, 44, 34, 114, 101, 113, 117, 105, 114, 101, 100, 34, 58, 116, 114, 117, 101, 125, 93, 125, 44, 34, 117, 110, 105, 113, 117, 101, 95, 101, 110, 116, 105, 116, 105, 101, 115, 34, 58, 91, 93, 44, 34, 112, 114, 111, 100, 117, 99, 116, 95, 115, 101, 97, 114, 99, 104, 95, 99, 111, 110, 102, 105, 103, 117, 114, 97, 116, 105, 111, 110, 34, 58, 123, 34, 112, 114, 111, 100, 117, 99, 116, 95, 115, 101, 116, 34, 58, 34, 118, 105, 100, 101, 111, 95, 103, 97, 109, 101, 115, 95, 110, 105, 110, 116, 101, 110, 100, 111, 34, 44, 34, 112, 105, 112, 101, 108, 105, 110, 101, 95, 116, 121, 112, 101, 34, 58, 34, 34, 44, 34, 112, 114, 111, 100, 117, 99, 116, 95, 116, 105, 116, 108, 101, 34, 58, 110, 117, 108, 108, 44, 34, 112, 114, 111, 100, 117, 99, 116, 95, 99, 97, 116, 101, 103, 111, 114, 121, 34, 58, 34, 97, 112, 112, 97, 114, 101, 108, 45, 118, 50, 34, 44, 34, 102, 105, 108, 116, 101, 114, 105, 110, 103, 95, 99, 114, 105, 116, 101, 114, 105, 97, 34, 58, 91, 93, 44, 34, 111, 109, 97, 107, 97, 115, 101, 95, 101, 110, 116, 105, 116, 121, 95, 108, 101, 118, 101, 108, 34, 58, 123, 34, 101, 110, 116, 105, 116, 121, 95, 105, 100, 34, 58, 34, 100, 53, 52, 57, 51, 57, 53, 100, 45, 98, 53, 98, 100, 45, 52, 97, 55, 55, 45, 57, 48, 99, 49, 45, 53, 57, 49, 51, 52, 50, 97, 54, 97, 51, 57, 51, 34, 44, 34, 101, 110, 116, 105, 116, 121, 95, 110, 97, 109, 101, 34, 58, 34, 109, 111, 100, 101, 108, 34, 125, 44, 34, 112, 114, 111, 100, 117, 99, 116, 95, 100, 105, 115, 112, 108, 97, 121, 95, 116, 105, 116, 108, 101, 34, 58, 91, 123, 34, 111, 114, 100, 101, 114, 34, 58, 48, 44, 34, 101, 110, 116, 105, 116, 121, 34, 58, 34, 109, 111, 100, 101, 108, 34, 44, 34, 112, 114, 101, 102, 105, 120, 34, 58, 34, 34, 44, 34, 115, 117, 102, 102, 105, 120, 34, 58, 34, 34, 125, 93, 44, 34, 112, 114, 111, 100, 117, 99, 116, 95, 100, 105, 115, 112, 108, 97, 121, 95, 101, 110, 116, 105, 116, 105, 101, 115, 34, 58, 91, 123, 34, 111, 114, 100, 101, 114, 34, 58, 48, 44, 34, 101, 110, 116, 105, 116, 121, 34, 58, 34, 115, 121, 115, 116, 101, 109, 95, 99, 97, 116, 101, 103, 111, 114, 121, 34, 44, 34, 101, 110, 116, 105, 116, 121, 95, 97, 108, 105, 97, 115, 34, 58, 34, 34, 125, 93, 44, 34, 112, 114, 111, 100, 117, 99, 116, 95, 100, 101, 115, 99, 114, 105, 112, 116, 105, 111, 110, 95, 101, 110, 116, 105, 116, 105, 101, 115, 34, 58, 91, 123, 34, 111, 114, 100, 101, 114, 34, 58, 48, 44, 34, 101, 110, 116, 105, 116, 121, 34, 58, 34, 112, 108, 97, 116, 102, 111, 114, 109, 34, 44, 34, 101, 110, 116, 105, 116, 121, 95, 97, 108, 105, 97, 115, 34, 58, 34, 34, 125, 93, 125, 125, 125, 44, 34, 105, 115, 95, 100, 114, 97, 102, 116, 34, 58, 102, 97, 108, 115, 101, 44, 34, 105, 115, 95, 103, 101, 110, 101, 114, 105, 99, 95, 116, 114, 101, 101, 34, 58, 102, 97, 108, 115, 101, 44, 34, 105, 115, 95, 108, 97, 114, 103, 101, 95, 116, 114, 101, 101, 34, 58, 102, 97, 108, 115, 101, 44, 34, 110, 97, 109, 101, 34, 58, 34, 110, 105, 110, 116, 101, 110, 100, 111, 95, 103, 97, 109, 101, 115, 34, 44, 34, 115, 116, 97, 116, 117, 115, 34, 58, 34, 110, 101, 119, 34, 125}
	strValue := string(bytearray)
	obj:=SPAMetadataTreeConfigurationParsed{}
	//fmt.Println(strValue)
	json.Unmarshal([]byte(strValue),&obj)
	fmt.Print(obj.ConfigRoot.Config.ItemNameSuggest.Filename)
}