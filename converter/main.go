package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"encoding/json"
	"strings"
)

var typeOrder = [5]string{"object", "string", "number", "integer", "boolean"}

func main() {
	flag.Parse()

	files, _ := ioutil.ReadDir("./schema")
	for _, f := range files {

		if !strings.HasSuffix(f.Name(), ".json") {
			continue
		}

		b, err := ioutil.ReadFile("./schema/" + f.Name())
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to read the input file with error ", err)
			return
		}

		fmt.Println("File: " + f.Name())
		newObject := convertObject(b)

		ioutil.WriteFile("./schema-processed/" + f.Name(), newObject, 0644)
	}
}

func convertObject(objectMessage json.RawMessage) json.RawMessage {
	var object map[string]*json.RawMessage
	json.Unmarshal(objectMessage, &object)

	// Create Copy of object
	newObject := make(map[string]*json.RawMessage)
	for k2,v2 := range object {
		newObject[k2] = v2
	}

	// START PROCESSING

	var required []string
	var properties map[string]*json.RawMessage

	if _, ok := object["type"]; ok {
		mostImportantType := json.RawMessage{}
		mostImportantType = getMostImportantType(*object["type"])
		newObject["type"] = &mostImportantType
	}

	if _, ok := object["properties"]; !ok {
		return objectMessage // if properties is missing, exit
	}
	err := json.Unmarshal(*object["properties"], &properties)
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

		// RECURSION CHECK
		if strings.Contains(string(*property["type"]), "array") { // ARRAYS ---
			if _, ok := property["items"]; !ok {
				panic("no items in array type!")
			}

			marshalledNewItems := json.RawMessage{}
			marshalledNewItems = convertObject(*property["items"])
			property["items"] = &marshalledNewItems
		}

		if strings.Contains(string(*property["type"]), "object") { // Nested Objects
			marshalledNewProperty := json.RawMessage{}
			marshalledNewProperty = convertObject(*propertyData)
			newProperties[propertyName] = &marshalledNewProperty

			continue
		}
		// RECURSION CHECK END

		if _, ok := property["required"]; ok {
			requiredParameter := true
			err = json.Unmarshal(*property["required"], &requiredParameter)
			if err != nil {
				fmt.Println(err)
				continue
			}

			if requiredParameter {
				required = append(required, propertyName)
			}

			delete(property, "required")
		}

		if _, ok := property["type"]; ok {
			mostImportantType := json.RawMessage{}
			mostImportantType = getMostImportantType(*property["type"])
			property["type"] = &mostImportantType
		}

		newpropertyData := json.RawMessage{}
		newpropertyData, _ = json.Marshal(property)

		newProperties[propertyName] = &newpropertyData
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

	marshalledNewObject := json.RawMessage{}
	marshalledNewObject, err = json.Marshal(newObject)

	return marshalledNewObject
}

func getMostImportantType(typeParameter json.RawMessage) json.RawMessage {
	if strings.Contains(string(typeParameter), "[") {
		marshalledType := json.RawMessage{}

		var typeParameterArray []string
		json.Unmarshal(typeParameter, &typeParameterArray)

		var typeParameterArrayWithoutNull []string
		for _, currentType := range typeParameterArray {
			if currentType != "null" {
				typeParameterArrayWithoutNull = append(typeParameterArrayWithoutNull, currentType)
			}
		}

		if len(typeParameterArrayWithoutNull) == 1 {
			marshalledType, _ = json.Marshal(typeParameterArrayWithoutNull[0])
			return marshalledType
		}

		// Multiple types like "string" and "number" in one type field
		var bestMatch string
		for _, to := range typeOrder {
			if containsType(typeParameterArrayWithoutNull, to) {
				bestMatch = to
				break
			}
		}

		if len(bestMatch) == 0 {
			panic("couldnt pick a type")
		}

		marshalledType, _ = json.Marshal(bestMatch)
		return marshalledType
	}

	return typeParameter
}

func containsType(types []string, search string) bool {
	for _, currentType := range types {
		if currentType == search {
			return true
		}
	}
	return false
}