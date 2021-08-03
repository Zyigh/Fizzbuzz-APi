package entities

import (
	"bytes"
	"fizzbuzz-api/conf"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// entity parent type of entities with generic methods
type entity struct {}

// postData save JSON as []byte to CouchDb database
func (e *entity) postData(dbname string, data []byte) error {
	dburl := fmt.Sprintf(
		"http://%s:%s@%s:%d/%s",
		conf.DBUserName,
		conf.DBPassword,
		conf.DBHost,
		conf.DBPort,
		dbname,
		)

	req, err := http.NewRequest("POST", dburl, bytes.NewBuffer(data))

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return err
	}

	log.Println("Posting data to \"", dburl, "\" returned :", res.Status)

	return res.Body.Close()
}

// createDatabaseIfNotExists checks if specific database exists and creates it if it doesn't
func (e *entity) createDatabaseIfNotExists(dbname string) error {
	dburl := fmt.Sprintf(
		"http://%s:%s@%s:%d/%s",
		conf.DBUserName,
		conf.DBPassword,
		conf.DBHost,
		conf.DBPort,
		dbname,
	)

	reqExists, err := http.NewRequest("HEAD", dburl, nil)

	if err != nil {
		return err
	}

	client := &http.Client{}
	resExists, err := client.Do(reqExists)

	defer resExists.Body.Close()

	if err != nil {
		return err
	}

	if resExists.StatusCode == http.StatusOK {
		return nil
	}

	reqCreate, err := http.NewRequest("PUT", dburl, nil)

	if err != nil {
		return err
	}

	resCreate, err := client.Do(reqCreate)

	if err != nil {
		return err
	}

	log.Println("Created database", dbname)

	return resCreate.Body.Close()
}

// getData makes a get request on specified database with specified query
func (e *entity) getData(dbname string, query string) (*[]byte, error) {
	dburl := fmt.Sprintf(
		"http://%s:%s@%s:%d/%s%s",
		conf.DBUserName,
		conf.DBPassword,
		conf.DBHost,
		conf.DBPort,
		dbname,
		query,
	)

	request, err := http.NewRequest("GET", dburl, nil)

	client := &http.Client{}
	response, err := client.Do(request)

	defer response.Body.Close()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if response.StatusCode > http.StatusNoContent {
		log.Println(dburl, "invalid database url")
		return nil, fmt.Errorf("invalid database url")
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &bodyBytes, nil
}
