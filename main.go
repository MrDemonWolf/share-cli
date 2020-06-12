package main

// Important neeeded modules
import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mrdemonwolf/share-cli/pkg/config"
)

type ResponseFile struct {
	Name      string `json:"name"`
	Ext       string `json:"ext"`
	Size      string `json:"size"`
	URL       string `json:"url"`
	Delete    string `json:"delete"`
	DeleteKey string `json:"deleteKey"`
}

// Response Json ...
type Response struct {
	Auth    bool         `json:"auth"`
	Success bool         `json:"success"`
	Error   string       `json:"error"`
	File    ResponseFile `json:"file"`
	Status  int          `json:"status"`
}

// Creates a new file upload http request with optional extra params
func newfileUploadRequest(uri string, headers map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	// Does the request with the headers
	req, err := http.NewRequest("POST", uri+"/api/v2/upload", body)
	for key, val := range headers {
		req.Header.Set(key, val)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("User-Agent", "Share-CLI/1.0")
	return req, err
}

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Println(err)
		return
	}

	// Setup the args for functions and there flags
	// This is for uploading a file
	fileUploadCommand := flag.NewFlagSet("upload", flag.ExitOnError)
	filePtr := fileUploadCommand.String("f", "", "File to upload to your share server. (Required)")
	pathPtr := fileUploadCommand.Bool("p", true, "If this is false it will not add the path to the file.")

	// This is for shortening alink
	linkShotenCommand := flag.NewFlagSet("link", flag.ExitOnError)
	urlPtr := linkShotenCommand.String("u", "", "File to upload to your share server.")
	limitPtr := linkShotenCommand.String("l", "", "File to upload to your share server.")
	// limitPtr := fileUploadCommand.String("f", "", "File to upload to your share server.")

	_ = urlPtr
	_ = limitPtr

	if len(os.Args) < 2 {
		fmt.Println("upload or link subcommand is required")
		return
	}

	switch os.Args[1] {
	case "upload":
		fileUploadCommand.Parse(os.Args[2:])
	case "link":
		linkShotenCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		return
	}

	if fileUploadCommand.Parsed() {
		if *filePtr == "" {
			fileUploadCommand.PrintDefaults()
			return
		}

		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
			return
		}

		var file = *filePtr

		if *pathPtr == true {
			file = path + "/" + *filePtr
		}

		headers := map[string]string{
			"Authorization": "Bearer" + " " + cfg.Creds.APIKEY,
		}

		request, err := newfileUploadRequest(cfg.Server.URL, headers, "file", file)
		if err != nil {
			log.Fatal(err)
		}
		client := &http.Client{}
		resp, err := client.Do(request)
		if err != nil {
			log.Fatal(err)
		} else {
			defer resp.Body.Close()

			data := new(Response)

			json.NewDecoder(resp.Body).Decode(data)

			if data.Status != 200 {
				fmt.Printf("Error: %s \n", data.Error)

			} else {
				fmt.Print("Your file has been uploaded")
				fmt.Printf("URL: %s \n", data.File.URL)
				fmt.Printf("Delete URL: %s \n", data.File.Delete)
			}
		}
	}
	if linkShotenCommand.Parsed() {

		// headers := map[string]string{
		// 	"Authorization": "Bearer" + " " + cfg.Creds.APIKEY,
		// }

		// request, err := newfileUploadRequest(cfg.Server.URL, headers, "file", file)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// client := &http.Client{}
		// resp, err := client.Do(request)
		// if err != nil {
		// 	log.Fatal(err)
		// } else {
		// 	defer resp.Body.Close()

		// 	data := new(Response)

		// 	json.NewDecoder(resp.Body).Decode(data)

		// 	fmt.Print("Your shorted link has been created")
		// 	fmt.Printf("URL: %s \n", data.File.URL)
		// 	fmt.Printf("Delete URL: %s \n", data.File.Delete)
		// }
		fmt.Print("Your shorted link has been created")

	}

}
