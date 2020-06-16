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
	"net/url"
	"os"
	"path/filepath"

	"github.com/mrdemonwolf/share-cli/pkg/config"
)

// ResponseFile ...
type ResponseFile struct {
	Name      string `json:"name"`
	Ext       string `json:"ext"`
	Size      string `json:"size"`
	URL       string `json:"url"`
	Delete    string `json:"delete"`
	DeleteKey string `json:"deleteKey"`
}

// ResponseLink ...
type ResponseLink struct {
	URL       string `json:"url"`
	Code      string `json:"code"`
	Limit     string `json:"limit"`
	NewURL    string `json:"newurl"`
	Delete    string `json:"delete"`
	DeleteKey string `json:"deleteKey"`
}

// ResponseErrors ...
type ResponseErrors struct {
	Code string `json:"code"`
	URL  string `json:"url"`
	Tags string `json:"tags"`
}

// Response Json ...
type Response struct {
	Auth    bool           `json:"auth"`
	Success bool           `json:"success"`
	Error   string         `json:"error"`
	Errors  ResponseErrors `json:"errors"`
	File    ResponseFile   `json:"file"`
	Link    ResponseLink   `json:"link"`
	Status  int            `json:"status"`
}

// LinkFormData ...
type LinkFormData struct {
	URL   string
	Code  string
	Limit string
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

func newLinkRequest(uri string, headers map[string]string, formData *LinkFormData) (*http.Request, error) {
	data := url.Values{}
	if formData.URL != "" {
		data.Set("url", formData.URL)
	}
	if formData.Code != "" {
		data.Set("code", formData.Code)
	}
	if formData.Limit != "" {
		data.Set("limit", formData.Limit)
	}

	// Does the request with the headers
	req, err := http.NewRequest("POST", uri+"/api/v2/link", bytes.NewBuffer([]byte(data.Encode())))
	for key, val := range headers {
		req.Header.Set(key, val)
	}
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
	urlPtr := linkShotenCommand.String("u", "", "File to upload to your share server. (Required)")
	limitPtr := linkShotenCommand.String("l", "", "Limit on link clicks.")
	codePtr := linkShotenCommand.String("c", "", "Custom short link code.")

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
			log.Println(err)
			return
		}
		client := &http.Client{}
		resp, err := client.Do(request)
		if err != nil {
			log.Println(err)
		} else {
			defer resp.Body.Close()

			data := new(Response)

			json.NewDecoder(resp.Body).Decode(data)

			if data.Status == 200 {
				fmt.Print("Your file has been uploaded")
				fmt.Printf("URL: %s \n", data.File.URL)
				fmt.Printf("Delete URL: %s \n", data.File.Delete)
			} else {
				fmt.Printf("Error: %s \n", data.Error)

			}
		}
	}
	if linkShotenCommand.Parsed() {
		if *urlPtr == "" {
			linkShotenCommand.PrintDefaults()
			return
		}

		headers := map[string]string{
			"Authorization": "Bearer" + " " + cfg.Creds.APIKEY,
			"Content-Type":  "application/x-www-form-urlencoded",
		}

		data := &LinkFormData{
			URL:   *urlPtr,
			Limit: *limitPtr,
			Code:  *codePtr,
		}

		request, err := newLinkRequest(cfg.Server.URL, headers, data)
		if err != nil {
			log.Println(err)
			return
		}
		client := &http.Client{}
		resp, err := client.Do(request)
		if err != nil {
			log.Println(err)
		} else {
			defer resp.Body.Close()

			data := new(Response)

			json.NewDecoder(resp.Body).Decode(data)

			if data.Status == 200 {
				// 	fmt.Print("Your shorted link has been created")
				// 	fmt.Printf("URL: %s \n", data.File.URL)
				// 	fmt.Printf("Delete URL: %s \n", data.File.Delete)
				// }
				fmt.Print("Your shorted link has been created \n")
				fmt.Printf("New URL: %s \n", data.Link.NewURL)
				if data.Link.Code != "" {
					fmt.Printf("Click Limit: %s \n", data.Link.Limit)
				}
				fmt.Printf("Delete URL: %s \n", data.Link.Delete)

			} else {
				fmt.Printf("Error: %s %s %s  \n", data.Errors.Code, data.Errors.URL, data.Errors.Tags)
			}

		}
	}

}
