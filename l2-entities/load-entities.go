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
	req.Header.Set("Authorization", "Bearer eyJ0eXAiOiAiSldUIiwgImFsZyI6ICJSUzI1NiIsICJraWQiOiAiMzk4OTViZjhjZTA5MWI2MWNlMjg2Yzc2OTYxYjdmMjAxYWNlMjNhNCJ9.eyJpYXQiOiAxNjg0MjI5NDAyLCAiZXhwIjogMTY4NDIzMzAwMiwgImlzcyI6ICI1NzM0NDU2OTYxMTEtY29tcHV0ZUBkZXZlbG9wZXIuZ3NlcnZpY2VhY2NvdW50LmNvbSIsICJhdWQiOiAiaHR0cHM6Ly9rZy1zdGFnZS5lbmRwb2ludHMubWVyY2FyaS11cy1kZS5jbG91ZC5nb29nIiwgInN1YiI6ICI1NzM0NDU2OTYxMTEtY29tcHV0ZUBkZXZlbG9wZXIuZ3NlcnZpY2VhY2NvdW50LmNvbSIsICJlbWFpbCI6ICI1NzM0NDU2OTYxMTEtY29tcHV0ZUBkZXZlbG9wZXIuZ3NlcnZpY2VhY2NvdW50LmNvbSJ9.DVr1uTh3ntmHJmX8qAaiSe4YwX5Cpn3Hsc5fdtj5npRd5-hu5ujEoPq-FUospz_G5kRsF906zlJWDe34iFwKL6-Fq9jycSze15Z1WHC6SBU_9n6-9EDSS2xEE5vAgcJnu4fHLwL2FHULhzjwZd5ojXQgYpng2E6P8dHAaoJ71_IihKJBsWsVRfXK4tmV0OoSCO7JF78IoSmOk3OZxQQE3jVQcwgAo677K6l68YhWYaChGNWKqjC8-Kj8qTkPCEhZvYmr2EffwDmNwzWieMiKpU4ixMFkWKX9paS6zrQFEhfYmaf-SttuWt2-GGrhlLA5Gp2S1zfubT1RqzROlslFBg")
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
