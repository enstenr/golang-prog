package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type EntityResponse struct {
	Count    int      `json:"count"`
	Response []string `json:"response"`
	Success  bool     `json:"success"`
}
type EntityRequest struct {
	TreeID string `json:"tree_id"`
}

func loadEntities(metadataTreeConfigId string, l2_categories_map map[int64]map[string]string, categoryName string, category int64) map[int64]map[string]string {
	url := "https://kg-stage.endpoints.mercari-us-de.cloud.goog/get_entities_for_tree"

	entities_map := make(map[string]string)
	request_body := EntityRequest{}
	request_body.TreeID = metadataTreeConfigId
	req_obj, _ := json.Marshal(request_body)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(req_obj))
	req.Header.Set("Authorization", "Bearer <<TOKEN>>")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var result EntityResponse

	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON obj")
	}

	//fmt.Print(" Entities : ")
	//fmt.Print(result.Response)
	for _, value := range result.Response {

		entities_map[value] = categoryName

	}
	l2_categories_map[category] = entities_map
	return l2_categories_map

}
