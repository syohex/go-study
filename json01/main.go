package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	m := make(map[string]interface{})
	m["name"] = "Tom"
	m["age"] = 9281
	m["hobbies"] = []string{"apple", "guitar", "java"}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "\t")
	if err := encoder.Encode(&m); err != nil {
		log.Fatalln(err)
	}
}
