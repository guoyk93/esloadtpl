package main

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	ExtJSON = ".json"
)

func exit(err *error) {
	if *err != nil {
		log.Printf("exited with error: %s", (*err).Error())
		os.Exit(1)
	} else {
		log.Println("exited")
	}
}

func main() {
	var err error
	defer exit(&err)

	// elasticsearch client
	var client *elastic.Client
	if client, err = elastic.NewClient(
		elastic.SetURL(os.Getenv("ES_URL")),
		elastic.SetSniff(false),
	); err != nil {
		return
	}

	// working directory
	var dir string
	if dir, err = os.Getwd(); err != nil {
		return
	}

	// read dir
	var fis []os.FileInfo
	if fis, err = ioutil.ReadDir(dir); err != nil {
		return
	}

	// find .json file and put index template
	for _, fi := range fis {
		if !strings.HasSuffix(fi.Name(), ExtJSON) {
			continue
		}
		name := strings.TrimSuffix(fi.Name(), ExtJSON)
		log.Println("updating template:", name)
		var buf []byte
		if buf, err = ioutil.ReadFile(filepath.Join(dir, fi.Name())); err != nil {
			return
		}
		if _, err = client.IndexPutTemplate(name).BodyJson(json.RawMessage(buf)).Do(context.Background()); err != nil {
			return
		}
	}
}
