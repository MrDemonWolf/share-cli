package main

// Important neeeded modules
import (
	"bytes"
	"encoding/json"
	"flag"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Config ...
type Config struct {
	Server struct {
		URL string `yaml:"url"`
	} `yaml:"server"`
	Creds struct {
		APIKEY string `yaml:"apikey"`
	} `yaml:"creds"`
}

// Response ...
type Response struct {
	Auth    bool `json:"auth"`
	Success bool `json:"success"`
	File    struct {
		Name      string `json:"name"`
		Ext       string `json:"ext"`
		Size      string `json:"size"`
		URL       string `json:"url"`
		Delete    string `json:"delete"`
		DeleteKey string `json:"deleteKey"`
	} `json:"file"`
	Status int `json:"status"`
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

	req, err := http.NewRequest("POST", uri, body)
	for key, val := range headers {
		req.Header.Set(key, val)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}
func main() {
	f, err := os.Open("config.yml")
	if err != nil {
		log.Println(err)
		log.Print("Please create a config file.")
		os.Exit(1)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	filePtr := flag.String("file", "", "File to upload to your share server (Required)")
	flag.Parse()

	if *filePtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	headers := map[string]string{
		"Authorization": "Bearer" + " " + cfg.Creds.APIKEY,
	}
	request, err := newfileUploadRequest(cfg.Server.URL, headers, "file", path+"/"+*filePtr)
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

		log.Print("Your file has been uploaded")
		log.Printf("URL: %s \n", data.File.URL)
		log.Printf("Delete URL: %s \n", data.File.Delete)
	}
}
