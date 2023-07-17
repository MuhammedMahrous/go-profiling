package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

type Response struct {
	RequestID string
	Messages  []string
}

type HelloHandler struct {
	version int
}

func NewHelloHandler(version int) (*HelloHandler, error) {
	if version != 1 && version != 2 {
		return nil, fmt.Errorf("invalid version %d", version)
	}

	return &HelloHandler{version: version}, nil
}

func (h *HelloHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch h.version {
	case 1:
		h.sayHelloV1(rw, r)
	case 2:
		h.sayHelloV2(rw, r)
	default:
		rw.WriteHeader(500)
		rw.Write([]byte("Invalid App Version!"))
	}

}

// TODO: refactor
func (h *HelloHandler) sayHelloV1(rw http.ResponseWriter, r *http.Request) {
	response := Response{RequestID: uuid.NewString()}

	for i := 0; i < 100; i++ {
		response.Messages = append(response.Messages, "hello-"+uuid.NewString())
	}

	res, _ := json.Marshal(response)

	//TODO: Marshal json directly to output stream / http socket
	rw.Write(res)
}

func (h *HelloHandler) sayHelloV2(rw http.ResponseWriter, r *http.Request) {
	response := Response{RequestID: uuid.NewString()}

	for i := 0; i < 100; i++ {
		response.Messages = append(response.Messages, "hello-"+strconv.Itoa(i))
	}

	res, _ := json.Marshal(response)

	//TODO: Marshal json directly to output stream / http socket
	rw.Write(res)
}
