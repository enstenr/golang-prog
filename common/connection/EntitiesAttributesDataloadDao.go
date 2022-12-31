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

func FetchEntitiesAttributes(env string, entityId string, attributeId string, db *sql.DB, threadCount int,entity_script *os.File,attribute_script *os.File,entity_attribute_script *os.File) {




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

	fmt.Print(threadCount)
	fmt.Print(" ")
		entityObj := customtypes.Entity{}
		attributeObj := customtypes.Attribute{}

		_ = rows.Scan(&entityObj.EntityID, &entityObj.Name, &entityObj.Alias, &attributeObj.AttributeID, &attributeObj.Attribute, &attributeObj.ProcessedAttribute, &attributeObj.EntityAttributeID)
		//fmt.Print(entityObj.EntityID,attributeObj.AttributeID)
		
		InserIntoEntity(entityObj,entity_script)
		InserIntoAttribute(attributeObj,attribute_script)
		InserIntoAttributeEntity(attributeObj, entityObj,entity_attribute_script)
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
	jsonFile, err := os.Open("entity_attribute.json")
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

	entity_script, err := os.OpenFile("entity_script.log",os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	attribute_script, err := os.OpenFile("attribute_script.log",os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	entity_attribute_script, err := os.OpenFile("entity_attribute_script.log",os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
 
defer entity_script.Close()
defer attribute_script.Close()
defer entity_attribute_script.Close()

	for s.Scan() {
		var v ValueTypeToUnmarshalTo
		if err := json.Unmarshal(s.Bytes(), &v); err != nil {
			//handle error
		}

		if !strings.Contains(v.Key, "000_") { //TODO STARTS WITH
			entAttrArray := strings.Split(v.Key, "-ha-")
			//fmt.Print(loop, " ",entAttrArray)

			loop++

			go FetchEntitiesAttributes(env, entAttrArray[0], entAttrArray[1], db, loop,entity_script,attribute_script,entity_attribute_script)
			//fmt.Println(entAttrArray[0], entAttrArray[1],loop)
			if loop >= threadCount {
				threadCount = threadCount + 100
				time.Sleep(10 * time.Second)
			}

		}
	}
	if s.Err() != nil {
		db = GetDBConnection(env, "entitydb")
	}
}
func InserIntoAttribute(attributeObj customtypes.Attribute,attribute_script *os.File) {

	QUERY_STR := ` Insert into Attribute ("attributeId","attribute","processed_attribute") values( 'ATTRIBUTE_ID','ATTRIBUTE','PROCESSED_ATTRIBUTE')  on conflict do nothing;`
	QUERY_STR = strings.Replace(QUERY_STR, "ATTRIBUTE_ID", attributeObj.AttributeID, -1)
	QUERY_STR = strings.Replace(QUERY_STR, "ATTRIBUTE", attributeObj.Attribute, -1)
	QUERY_STR = strings.Replace(QUERY_STR, "PROCESSED_ATTRIBUTE", attributeObj.ProcessedAttribute, -1)

	//fmt.Println(QUERY_STR)
	if _, err := attribute_script.WriteString(QUERY_STR); err != nil {
	 
	}
}

func InserIntoEntity(entityObj customtypes.Entity,entity_script *os.File) {

	QUERY_STR := ` Insert into Entity ("entityId","name","alias") values( 'ENTITY_ID','NAME','ALIAS')  on conflict do nothing;`
	QUERY_STR = strings.Replace(QUERY_STR, "ENTITY_ID", entityObj.EntityID, -1)
	QUERY_STR = strings.Replace(QUERY_STR, "NAME", entityObj.Name, -1)
	QUERY_STR = strings.Replace(QUERY_STR, "ALIAS", entityObj.Alias, -1)

	//fmt.Println(QUERY_STR)
	if _, err := entity_script.WriteString(QUERY_STR); err != nil {
	 
	}
}

func InserIntoAttributeEntity(attributeObj customtypes.Attribute, entityObj customtypes.Entity,entity_attribute_script *os.File) {

	QUERY_STR := ` Insert into attribute_entity ("attributeEntityId","entityId","attributeId","measureIndex") values( 'ATTRIBUTE_ENTITY_ID','ENTITY_ID','ATTRIBUTE_ID',0)  on conflict do nothing;`
	QUERY_STR = strings.Replace(QUERY_STR, "ATTRIBUTE_ENTITY_ID", attributeObj.EntityAttributeID, -1)
	QUERY_STR = strings.Replace(QUERY_STR, "ENTITY_ID", entityObj.EntityID, -1)
	QUERY_STR = strings.Replace(QUERY_STR, "ATTRIBUTE_ID", attributeObj.AttributeID, -1)

	//fmt.Println(QUERY_STR)
	if _, err := entity_attribute_script.WriteString(QUERY_STR); err != nil {
	 
	}
}
