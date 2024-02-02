package utils

import (
	"net/http"
	"io"
	"encoding/json"
)

func ParseBody(r *http.Request, X interface{}){
	if body, err := io.ReadAll(r.Body); err == nil{
		if err := json.Unmarshal([]byte(body), &X); err != nil{
			return 
		}
	}
}