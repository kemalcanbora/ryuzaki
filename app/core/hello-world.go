package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ryuzaki/helpers/schema"
	kafka "ryuzaki/platform/streaming"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	dummyData := &schema.Comment{
		Text:   "tes",
		Bos:    "pear",
		Fruits: []string{"a", "b", "c"}}
	jsonData, _ := json.Marshal(dummyData)

	kafka.Producer("jungle", jsonData)
	fmt.Fprintf(w, `Hello world`)
}
