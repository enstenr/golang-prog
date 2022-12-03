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

func loadEntities(metadataTreeConfigId string){
	url:="https://kg-dev.endpoints.mercari-us-de.cloud.goog/get_entities_for_tree"

	request_body:=EntityRequest{}
	request_body.TreeID=metadataTreeConfigId
	req_obj, _ := json.Marshal(request_body)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(req_obj))
    req.Header.Set("Authorization", "Bearer eyJ0eXAiOiAiSldUIiwgImFsZyI6ICJSUzI1NiIsICJraWQiOiAiZDkwYzViNGU3OGI0OTE1YzEyYmU0YmJlOTk2MTdmZGU0MGY5NDU5MyJ9.eyJpYXQiOiAxNjcwMDQ0Mzc5LCAiZXhwIjogMTY3MDA0Nzk3OSwgImlzcyI6ICI1NzM0NDU2OTYxMTEtY29tcHV0ZUBkZXZlbG9wZXIuZ3NlcnZpY2VhY2NvdW50LmNvbSIsICJhdWQiOiAiaHR0cHM6Ly9rZy1kZXYuZW5kcG9pbnRzLm1lcmNhcmktdXMtZGUuY2xvdWQuZ29vZyIsICJzdWIiOiAiNTczNDQ1Njk2MTExLWNvbXB1dGVAZGV2ZWxvcGVyLmdzZXJ2aWNlYWNjb3VudC5jb20iLCAiZW1haWwiOiAiNTczNDQ1Njk2MTExLWNvbXB1dGVAZGV2ZWxvcGVyLmdzZXJ2aWNlYWNjb3VudC5jb20ifQ.egdV2bwdkhkFOXniHNy2zNTQjAXRDLI-1VI5BY3BiHb8zMlR6rTz9EmovWrXJ04NNsVXRqMK0JdaPyyyUgeG-gioTP3qodMVnYmJDOlQTv32pYhzU1MXVNCB6DxYHKnEUoXoLug8jeAizgk7ZoPqI6CxZHnuLvxrkhqqEDRkdeP-uR1ITacZs27rXi1X76GKbPMljTJJtDpFGecqm25gcnsSN4MbXTuUFGUXDP6DlA5kT2Gd3uyLcGOMpjBFERhbrD2daqeZSvCd8KFFryoFCMYfgezPtw9wZA8novrLKJJHk48JAN0F6AEEXmVGydF6pQG-smq4AeycPiZ6Nl36qQ")
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
	fmt.Print(" Entities : ")
	fmt.Print(result.Response)

}