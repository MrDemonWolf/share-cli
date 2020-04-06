package main

// Important neeeded modules
import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/mitchellh/go-homedir"
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
	config, err := homedir.Expand("~/.config/share-cli")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	f, err := os.Open(config + "/config.yml")
	if err != nil {
		err = os.MkdirAll(config, os.ModePerm)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		c := Config{}
		d, err := yaml.Marshal(&c)

		if err != nil {
			log.Fatalf("error: %v", err)
		}
		// Write bytes to file
		err = ioutil.WriteFile(config+"/config.yml", []byte(d), 0644)
		log.Print("Created config file please open ~/.config/share-cli/config.yml")
		log.Print("URL: The server API URL (https://example.com/api/v1/upload")
		log.Print("APIKEY: Your API key for using for auth")
		os.Exit(1)
	}

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Println(err)

		os.Exit(1)
	}
	defer f.Close()

	if len(cfg.Server.URL) <= 0 && len(cfg.Creds.APIKEY) <= 0 {
		fmt.Println("You MUST put URL of the server you want to upload the file to.")
		fmt.Println("Token must be provided for the server to know who you are.")

		os.Exit(1)
	}

	if len(cfg.Server.URL) <= 0 {
		fmt.Println("You MUST put URL of the server you want to upload the file to.")
		os.Exit(1)
	}
	if len(cfg.Creds.APIKEY) <= 0 {
		fmt.Println("Token must be provided for the server to know who you are.")
		os.Exit(1)
	}

	filePtr := flag.String("f", "", "File to upload to your share server (Required)")
	pathPtr := flag.Bool("p", true, "If this is false it will not add the path to the file.")
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

		fmt.Print("Your file has been uploaded")
		fmt.Printf("URL: %s \n", data.File.URL)
		fmt.Printf("Delete URL: %s \n", data.File.Delete)
	}
}
