package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	const docId = "35PSuBUAfQVx6eAblP7Vf"

	var token string

	syncCmd := flag.NewFlagSet("sync", flag.ExitOnError)
	syncCmd.StringVar(&token, "token", os.Getenv("DROPBOX_API_TOKEN"), "Dropbox API token")

	if len(os.Args) == 1 {
		fmt.Println("improper amount of arguments")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "sync":
		syncCmd.Parse(os.Args[2:])
		fetch(token, docId)
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func fetch(token string, docId string) {
	req, err := http.NewRequest("POST", "https://api.dropboxapi.com/2/paper/docs/download", nil)
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
