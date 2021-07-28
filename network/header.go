package network

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func GetHeader(w http.ResponseWriter, req *http.Request) (map[string]string, error) {
	if req.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return nil, fmt.Errorf("bad request")
	}

	if req.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return nil, fmt.Errorf("bad request")
	}

	//To allocate slice for request body
	length, err := strconv.Atoi(req.Header.Get("Content-Length"))
	if err != nil {
		ErrorStatus(w)
		return nil, fmt.Errorf("content-length flags is not found")
	}

	//Read body data to parse json
	body := make([]byte, length)
	length, err = req.Body.Read(body)
	if err != nil && err != io.EOF {
		ErrorStatus(w)
		return nil, fmt.Errorf("parse error")
	}

	//parse json
	var jsonBody map[string]string
	err = json.Unmarshal(body[:length], &jsonBody)
	if err != nil {
		ErrorStatus(w)
		fmt.Println(err)
		return nil, fmt.Errorf("parse error")
	}
	return jsonBody, nil
}

func GetData(key string, w http.ResponseWriter, r *http.Request) (string, error) {
	jsonBody, err := GetHeader(w, r)
	if err != nil {
		return "", err
	}

	if value, ok := jsonBody[key]; !ok {
		return value, nil
	}

	ErrorStatus(w)
	return "", fmt.Errorf("key is not found")
}
