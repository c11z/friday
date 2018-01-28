package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	paperDocsDownloadURL = "https://api.dropboxapi.com/2/paper/docs/download"
)

var (
	token string
	docId string
)

func init() {
	token = os.Getenv("DROPBOX_API_TOKEN")
	docId = os.Getenv("JOURNAL_DOC_ID")
}

func main() {
	req, err := http.NewRequest("POST", paperDocsDownloadURL, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Dropbox-API-Arg", fmt.Sprintf(`{"doc_id":"%s","export_format": "markdown"}`, docId))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	fmt.Println(resp.Header)
}
