package utils

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"sync"
)

func Http_(api string, req_type string, body []byte, headerKeyValue map[string]string, wg *sync.WaitGroup) error {
	r, err := http.NewRequest(req_type, api, bytes.NewBuffer(body))
	if err != nil {
		log.Println("Error while trying to create the requets")
		return err
	}
	for k, v := range headerKeyValue {
		r.Header.Add(k, v)
	}
	r.Header.Add("Content-Type", "application/json")

	clt := &http.Client{}
	res, err := clt.Do(r)

	if err != nil {
		log.Println("Error while performing the request: ", err)
		return err
	}

	defer res.Body.Close()
	rByte, err := io.ReadAll(res.Body)

	if err != nil {
		log.Println("Error reading the response to byte array: ", err)
	}

	log.Println("Request successfully completed")
	log.Println("Response: ", string(rByte))
	if wg != nil {
		wg.Done()
	}
	return nil
}
