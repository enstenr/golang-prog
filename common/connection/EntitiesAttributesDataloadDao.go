package connection

import (
	"bufio"
	"bytes"
	"database/sql"
	"encoding/json"

	"fmt"
	"os"
	"strings"
	"time"

	"github.com/enstenr/customtypes"

	_ "github.com/lib/pq"
)

func FetchEntitiesAttributes(env string, entityId string, attributeId string, db *sql.DB, threadCount int) {

	fmt.Print(threadCount)
	buf := bytes.NewBufferString(`
	select e."entityId",e."name" ,e."alias",a."attributeId",a."attribute",a."processed_attribute",aet."attributeEntityId" from entity e, attribute a, attribute_temp aet  
	where a."attributeId"=aet."attributeId" 
	and e."entityId"=aet."entityId"  and 
	e."entityId"=$1
	and a."attributeId"=$2 `)

	rows, err := db.Query(buf.String(), entityId, attributeId)

	if err != nil {
		fmt.Print(err)
	}
	defer rows.Close()
	for rows.Next() {
		entityObj := customtypes.Entity{}
		attributeObj := customtypes.Attribute{}

		_ = rows.Scan(&entityObj.EntityID, &entityObj.Name, &entityObj.Alias, &attributeObj.AttributeID, &attributeObj.Attribute, &attributeObj.ProcessedAttribute, &attributeObj.EntityAttributeID)
		//fmt.Print(entityObj.EntityID,attributeObj.AttributeID)

		InserIntoEntity(entityObj)
		InserIntoAttribute(attributeObj)
		InserIntoAttributeEntity(attributeObj, entityObj)
	}

}

type ValueTypeToUnmarshalTo struct {
	Key     string `json:"_key"`
	From    string `json:"_from"`
	To      string `json:"_to"`
	Created string `json:"created"`
}

func ParseJson(env string) {
	db := GetDBConnection(env, "entitydb")
	defer db.Close()
	jsonFile, err := os.Open("framework_dev_data_apps_bowker.com_romance_fiction_books_bowker_kg_transformation_2022_12_06-00_01_09_collections_entity-attributes-col.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	if err != nil {
		// handle error
	}
	s := bufio.NewScanner(jsonFile)
	var loop int = 0
	var threadCount int = 0
	for s.Scan() {
		var v ValueTypeToUnmarshalTo
		if err := json.Unmarshal(s.Bytes(), &v); err != nil {
			//handle error
		}

		if !strings.Contains(v.Key, "000_") { //TODO STARTS WITH
			entAttrArray := strings.Split(v.Key, "-ha-")
			//fmt.Print(loop, " ")

			loop++

			go FetchEntitiesAttributes(env, entAttrArray[0], entAttrArray[1], db, loop)
			//fmt.Println(entAttrArray[0], entAttrArray[1],loop)
			if loop >= threadCount {
				threadCount = threadCount + 500
				time.Sleep(20 * time.Second)
			}

		}
	}
	if s.Err() != nil {
		// handle scan error
	}
}
func InserIntoAttribute(attributeObj customtypes.Attribute) {

	QUERY_STR := ` Insert into Attribute ("attributeId","attribute","processed_attribute") values( 'ATTRIBUTE_ID','ATTRIBUTE','PROCESSED_ATTRIBUTE')  on conflict do nothing;`
	QUERY_STR = strings.Replace(QUERY_STR, "ATTRIBUTE_ID", attributeObj.AttributeID, -1)
	QUERY_STR = strings.Replace(QUERY_STR, "ATTRIBUTE", attributeObj.Attribute, -1)
	QUERY_STR = strings.Replace(QUERY_STR, "PROCESSED_ATTRIBUTE", attributeObj.ProcessedAttribute, -1)

	fmt.Println(QUERY_STR)
}

func InserIntoEntity(entityObj customtypes.Entity) {

	QUERY_STR := ` Insert into Entity ("entityId","name","alias") values( 'ENTITY_ID','NAME','ALIAS')  on conflict do nothing;`
	QUERY_STR = strings.Replace(QUERY_STR, "ENTITY_ID", entityObj.EntityID, -1)
	QUERY_STR = strings.Replace(QUERY_STR, "NAME", entityObj.Name, -1)
	QUERY_STR = strings.Replace(QUERY_STR, "ALIAS", entityObj.Alias, -1)

	fmt.Println(QUERY_STR)
}

func InserIntoAttributeEntity(attributeObj customtypes.Attribute, entityObj customtypes.Entity) {

	QUERY_STR := ` Insert into attribute_entity ("attributeEntityId","entityId","attributeId","measureIndex") values( 'ATTRIBUTE_ENTITY_ID','ENTITY_ID','ATTRIBUTE_ID',0)  on conflict do nothing;`
	QUERY_STR = strings.Replace(QUERY_STR, "ATTRIBUTE_ENTITY_ID", attributeObj.EntityAttributeID, -1)
	QUERY_STR = strings.Replace(QUERY_STR, "ENTITY_ID", entityObj.EntityID, -1)
	QUERY_STR = strings.Replace(QUERY_STR, "ATTRIBUTE_ID", attributeObj.AttributeID, -1)

	fmt.Println(QUERY_STR)
}
