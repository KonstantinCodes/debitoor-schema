package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"encoding/json"
)

func main() {
	flag.Parse()

	files, _ := ioutil.ReadDir("./schema")
	for _, f := range files {
		b, err := ioutil.ReadFile("./schema/" + f.Name())

		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to read the input file with error ", err)
			return
		}

		var object map[string]*json.RawMessage
		json.Unmarshal(b, &object)

		var required []string
		var properties map[string]*json.RawMessage

		if _, ok := object["properties"]; !ok {
			continue // if properties is missing
		}
		err = json.Unmarshal(*object["properties"], &properties)
		if err != nil {
			fmt.Println(err)
		}


		// Create Copy
		newProperties := make(map[string]*json.RawMessage)
		for k2,v2 := range properties {
			newProperties[k2] = v2
		}

		for propertyName, propertyData := range properties {
			var property map[string]*json.RawMessage
			err = json.Unmarshal(*propertyData, &property)

			var requiredParameter bool
			requiredParameter = true
			if _, ok := property["required"]; !ok {
				continue // if required is missing
			}
			err = json.Unmarshal(*property["required"], &requiredParameter)
			if err != nil {
				fmt.Println(err)
				continue
			}

			if requiredParameter {
				required = append(required, propertyName)
			}

			delete(property, "required")

			newpropertyData := json.RawMessage{}
			newpropertyData, _ = json.Marshal(property)

			newProperties[propertyName] = &newpropertyData
		}

		// Create Copy of object
		newObject := make(map[string]*json.RawMessage)
		for k2,v2 := range object {
			newObject[k2] = v2
		}

		marshalledSchema := json.RawMessage{}
		marshalledSchema, err = json.Marshal("http://json-schema.org/draft-04/schema")

		marshalledNewProperties := json.RawMessage{}
		marshalledNewProperties, err = json.Marshal(newProperties)

		marshalledNewRequired := json.RawMessage{}
		marshalledNewRequired, err = json.Marshal(required)

		newObject["$schema"] = &marshalledSchema
		newObject["required"] = &marshalledNewRequired
		newObject["properties"] = &marshalledNewProperties

		marshalledNewObject, err := json.Marshal(newObject)

		fmt.Println(string(marshalledNewObject))

		ioutil.WriteFile("./schema-processed/" + f.Name(), marshalledNewObject, 0644)
	}
}