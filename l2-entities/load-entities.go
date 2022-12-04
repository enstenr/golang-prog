package main
import (
    "fmt"
    "bytes"
    "net/http"
	"io/ioutil"
    "encoding/json"
)

type EntityResponse struct {
	Count    int      `json:"count"`
	Response []string `json:"response"`
	Success  bool     `json:"success"`
}
type EntityRequest struct {
	TreeID string `json:"tree_id"`
}

func loadEntities(metadataTreeConfigId string,l2_categories_map map[string]map[string]struct {},category string)(map[string]map[string]struct {}){
	url:="https://kg-dev.endpoints.mercari-us-de.cloud.goog/get_entities_for_tree"
	type void struct {}
	var member void
	entities_map:=make(map[string]struct{})
	request_body:=EntityRequest{}
	request_body.TreeID=metadataTreeConfigId
	req_obj, _ := json.Marshal(request_body)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(req_obj))
    req.Header.Set("Authorization", "Bearer eyJ0eXAiOiAiSldUIiwgImFsZyI6ICJSUzI1NiIsICJraWQiOiAiZDkwYzViNGU3OGI0OTE1YzEyYmU0YmJlOTk2MTdmZGU0MGY5NDU5MyJ9.eyJpYXQiOiAxNjcwMDc4Mjc1LCAiZXhwIjogMTY3MDA4MTg3NSwgImlzcyI6ICI1NzM0NDU2OTYxMTEtY29tcHV0ZUBkZXZlbG9wZXIuZ3NlcnZpY2VhY2NvdW50LmNvbSIsICJhdWQiOiAiaHR0cHM6Ly9rZy1kZXYuZW5kcG9pbnRzLm1lcmNhcmktdXMtZGUuY2xvdWQuZ29vZyIsICJzdWIiOiAiNTczNDQ1Njk2MTExLWNvbXB1dGVAZGV2ZWxvcGVyLmdzZXJ2aWNlYWNjb3VudC5jb20iLCAiZW1haWwiOiAiNTczNDQ1Njk2MTExLWNvbXB1dGVAZGV2ZWxvcGVyLmdzZXJ2aWNlYWNjb3VudC5jb20ifQ.Q__qbJPsNDxxCglYaNApite37vLTWyX9rzdBQ2_Vg4EEG0uVVvkS5CkYi_VubM7W6uNHEPXf_NTvgNV3vt4A8bKepzixVFf75tLofUbIBvgvB6-2-RGM9yw3VYUyle7tgKsPupG1v_abz8m_XIl0Qh3AuOZMTuBk5OrmyOKKNPYPiTDRLdnjr2K_kuWHDpxmsX5BCt19BsBns6BEJeCax16DgENjKhSZ9rGsAgvwH_BTNEM3fXZBhcq-sTArq2tBu8-j5AUv9kNh5Tz3Uef2hcgfVJziBs7RdQYUJdG1BdiU8XeAXVMzFkzVRHgYzY1DNyoPRKHROooKy1jtlYTVhQ")
    req.Header.Set("Content-Type", "application/json; charset=UTF-8")


    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    
	body, _ := ioutil.ReadAll(resp.Body)
	var result EntityResponse
	if err := json.Unmarshal(body, &result); err != nil {   // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	//fmt.Print(" Entities : ")
	//fmt.Print(result.Response)
	for _,value:=range(result.Response){
		//fmt.Print(value)
		entities_map[value]=member
		
	}
	l2_categories_map[category]=entities_map
	return l2_categories_map

}