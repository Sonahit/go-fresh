package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status bool
}
type ResponseData struct {
	Data   interface{} `json:"data"`
	Status bool        `json:"status"`
}

func (r *Response) WriteResponse(data interface{}, writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(writer, fmt.Sprintf("%v", data))
}

func (r *Response) WriteJson(data interface{}, writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	encoder.Encode(ResponseData{Data: data, Status: r.Status})
}
