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

func loadEntities(metadataTreeConfigId string,l2_categories_map map[int64]map[string]string,categoryName string,category int64)(map[int64]map[string]string){
	url:="https://kg-dev.endpoints.mercari-us-de.cloud.goog/get_entities_for_tree"

	entities_map:=make(map[string]string)
	request_body:=EntityRequest{}
	request_body.TreeID=metadataTreeConfigId
	req_obj, _ := json.Marshal(request_body)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(req_obj))
    req.Header.Set("Authorization", "Bearer eyJ0eXAiOiAiSldUIiwgImFsZyI6ICJSUzI1NiIsICJraWQiOiAiZDkwYzViNGU3OGI0OTE1YzEyYmU0YmJlOTk2MTdmZGU0MGY5NDU5MyJ9.eyJpYXQiOiAxNjcwMjMwNzU3LCAiZXhwIjogMTY3MDIzNDM1NywgImlzcyI6ICI1NzM0NDU2OTYxMTEtY29tcHV0ZUBkZXZlbG9wZXIuZ3NlcnZpY2VhY2NvdW50LmNvbSIsICJhdWQiOiAiaHR0cHM6Ly9rZy1kZXYuZW5kcG9pbnRzLm1lcmNhcmktdXMtZGUuY2xvdWQuZ29vZyIsICJzdWIiOiAiNTczNDQ1Njk2MTExLWNvbXB1dGVAZGV2ZWxvcGVyLmdzZXJ2aWNlYWNjb3VudC5jb20iLCAiZW1haWwiOiAiNTczNDQ1Njk2MTExLWNvbXB1dGVAZGV2ZWxvcGVyLmdzZXJ2aWNlYWNjb3VudC5jb20ifQ.QZLoTnQGDnI0VzUuhYVa0Npi9TMIxF_HU849xIkkeq2z0-HlIvYBtQwf0kY2wAGZ8lEpWA25MGO_Sw7TTGHNI3K30hW5R81li8bnScO_k1DQy2mj4KP4sO9sgddDLnPSPzCg1ncxjhHQQ8u5BRHjTVGL2mLgFW2RwyHkwTo7HuKJKSl1xyc3vIGNpPYira2HCLycGuJ_hU1PLsW4glQhKNDvm4IFaB2pKpC0lGjbtDV3zW4tnG1iXqCEbDpCcFuJkzZlWW6yZPVe_TRC2cZ-SdFjVRifipyOODQ8xjDw7IRqGZn91jFGdYlo3SzbAvoy9cCWQz25JMbevOOln9cVUw")
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
		entities_map[value]=categoryName
		
	}
	l2_categories_map[category]=entities_map
	return l2_categories_map

}