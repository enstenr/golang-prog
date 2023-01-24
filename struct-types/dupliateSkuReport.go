package customtypes
 

type DupliateSkuReport struct {
	Tree_name string
	Message   string
	State     string
	Status    string
}

type MetadataTree struct {
	MetadataTreeId              string
	MetadataTreeConfigurationId string
	Name                        string
	Message	                    string
	Count                        int
	Hash						string
	Modified_date				string
	RePublishStatus				string
	GcsPath  string
	System_category string
	system_category_id string
	L2_category string
	l2_category_id string
	
}

type MetadataTreeDifference struct {
	MetadataTreeId              string
	MetadataTreeConfigurationId1 string
	MetadataTreeConfigurationId2 string
	Name                        string
	Message	                    string
	Count                        int
	Hash1						string
	Modified_date1				string
	RePublishStatus1			string
	Hash2						string
	Modified_date2				string
	RePublishStatus2			string
	
}

type Entity struct {
EntityID string
CreatedDate string
ModifiedDate string
Active bool
Name string
Alias string
Published bool
}

type Attribute struct {
AttributeID string
CreateDate string
ModifiedDate string
Active bool
Attribute string
Published bool
ProcessedAttribute string
EntityAttributeID string
}


type Item struct {
	ID    int
	Attrs MetadataTreeConfiguration
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

