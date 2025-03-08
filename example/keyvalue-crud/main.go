//go:generate go run go.bytecodealliance.org/cmd/wit-bindgen-go generate --world hello --out gen ./wit
package main

import (
	"encoding/json"
	"net/http"

	"github.com/jamesstocktonj1/wasi-lib/pkg/keyvalue"
	"github.com/julienschmidt/httprouter"
	"go.wasmcloud.dev/component/log/wasilog"
	"go.wasmcloud.dev/component/net/wasihttp"
)

var logger = wasilog.DefaultLogger

func init() {
	router := httprouter.New()
	router.POST("/document/:document", create)
	router.GET("/document/:document", read)
	router.DELETE("/document/:document", del)
	wasihttp.Handle(router)
}

func create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")

	doc := ps.ByName("document")
	logger.Info("handle create", "document", doc)

	bucket, err := keyvalue.Open("store")
	if err != nil {
		logger.Error("keyvalue.Open", "error", err)
		enc.Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	err = bucket.Set(doc, "Hello World!")
	if err != nil {
		logger.Error("bucket.Set", "error", err)
		enc.Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	enc.Encode(map[string]string{
		"message": "handle create",
	})
}

func read(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")

	doc := ps.ByName("document")
	logger.Info("handle read", "document", doc)

	bucket, err := keyvalue.Open("store")
	if err != nil {
		logger.Error("keyvalue.Open", "error", err)
		enc.Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	value, err := bucket.Get(doc)
	if err != nil {
		logger.Error("bucket.Get", "error", err)
		enc.Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	enc.Encode(map[string]string{
		"message": value,
	})
}

func del(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")

	doc := ps.ByName("document")
	logger.Info("handle delete", "document", doc)

	bucket, err := keyvalue.Open("store")
	if err != nil {
		logger.Error("keyvalue.Open", "error", err)
		enc.Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	err = bucket.Delete(doc)
	if err != nil {
		logger.Error("bucket.Delete", "error", err)
		enc.Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	enc.Encode(map[string]string{
		"message": "handle delete",
	})
}

// Since we don't run this program like a CLI, the `main` function is empty. Instead,
// we call the `handleRequest` function when an HTTP request is received.
func main() {}
