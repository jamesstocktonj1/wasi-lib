//go:generate go run go.bytecodealliance.org/cmd/wit-bindgen-go generate --world hello --out gen ./wit
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jamesstocktonj1/wasi-lib/pkg/config"
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
	documentId := fmt.Sprintf("%s.%s", config.GetStringDefault("keyspace", "default"), doc)
	logger.Info("handle create", "document", documentId)

	value := map[string]any{}
	err := json.NewDecoder(r.Body).Decode(&value)
	if err != nil {
		enc.Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	bucket, err := keyvalue.Open("store")
	if err != nil {
		logger.Error("keyvalue.Open", "error", err)
		enc.Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	err = bucket.SetDocument(documentId, value)
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
	documentId := fmt.Sprintf("%s.%s", config.GetStringDefault("keyspace", "default"), doc)
	logger.Info("handle read", "document", documentId)

	bucket, err := keyvalue.Open("store")
	if err != nil {
		logger.Error("keyvalue.Open", "error", err)
		enc.Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	value := map[string]any{}
	err = bucket.GetDocument(documentId, &value)
	if err != nil {
		logger.Error("bucket.GetDocument", "error", err)
		enc.Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	enc.Encode(value)
}

func del(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")

	doc := ps.ByName("document")
	documentId := fmt.Sprintf("%s.%s", config.GetStringDefault("keyspace", "default"), doc)
	logger.Info("handle delete", "document", documentId)

	if !config.GetBoolDefault("delete_enabled", false) {
		enc.Encode(map[string]string{
			"message": "delete not enabled",
		})
		return
	}

	bucket, err := keyvalue.Open("store")
	if err != nil {
		logger.Error("keyvalue.Open", "error", err)
		enc.Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	err = bucket.Delete(documentId)
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
