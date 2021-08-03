package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// controller parent type of Controllers with generic methods
type controller struct {}

// render tries to Marshal data as JSON and wrtie it in the ResponseWriter
func (c *controller) render(data interface{}, w http.ResponseWriter) {
	jsonData, err := json.Marshal(data)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	fmt.Fprint(w, string(jsonData))
}
