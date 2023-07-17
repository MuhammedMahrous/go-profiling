package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

const NUM_USERS = 1000

type Response struct {
	RequestID string
	Messages  []string
}

type HelloHandler struct {
	version int
}

func NewHelloHandler(version int) (*HelloHandler, error) {
	if version != 1 && version != 2 && version != 3 && version != 4 {
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
	case 3:
		h.sayHelloV3(rw, r)
	case 4:
		h.sayHelloV4(rw, r)
	default:
		rw.WriteHeader(500)
		rw.Write([]byte("Invalid App Version!"))
	}

}

// TODO: refactor
func (h *HelloHandler) sayHelloV1(rw http.ResponseWriter, r *http.Request) {
	response := Response{RequestID: uuid.NewString()}

	for i := 0; i < NUM_USERS; i++ {
		response.Messages = append(response.Messages, "hello-"+uuid.NewString())
	}

	res, _ := json.Marshal(response)

	rw.Write(res)
}

func (h *HelloHandler) sayHelloV2(rw http.ResponseWriter, r *http.Request) {
	response := Response{RequestID: uuid.NewString()}

	for i := 0; i < NUM_USERS; i++ {
		response.Messages = append(response.Messages, "hello-"+strconv.Itoa(i))
	}
	res, _ := json.Marshal(response)

	rw.Write(res)
}

func (h *HelloHandler) sayHelloV3(rw http.ResponseWriter, r *http.Request) {
	t := time.NewTicker(time.Microsecond)
	response := Response{RequestID: uuid.NewString()}

	for i := 0; i < NUM_USERS; i++ {
		select {
		case <-t.C:
			response.Messages = append(response.Messages, "hello-lucky-"+strconv.Itoa(i))
		default:
			response.Messages = append(response.Messages, "hello-"+strconv.Itoa(i))
		}
	}
	t.Stop()

	res, _ := json.Marshal(response)

	rw.Write(res)
}

func (h *HelloHandler) sayHelloV4(rw http.ResponseWriter, r *http.Request) {
	t := time.NewTicker(time.Microsecond)
	response := Response{RequestID: uuid.NewString()}

	for i := 0; i < NUM_USERS; i++ {
		select {
		case <-t.C:
			response.Messages = append(response.Messages, "hello-lucky-"+strconv.Itoa(i))
		default:
			response.Messages = append(response.Messages, "hello-"+strconv.Itoa(i))
		}
	}

	t.Stop()

	res, _ := json.Marshal(response)

	rw.Write(res)
}
