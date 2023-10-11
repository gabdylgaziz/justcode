package main

import (
	"encoding/json"
	"fmt"
)

// Пример JSON-структуры
var jsonInput = `
{
	"Person": {
		"Name": "John",
		"Age": 30,
		"Address": {
			"Street": "123 Main St",
			"City": "New York"
		}
	}
}
`

func main() {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonInput), &data); err != nil {
		fmt.Println(err)
		return
	}

	generateGoCode(data, "MyGeneratedStruct")
}

func generateGoCode(data map[string]interface{}, structName string) {
	fmt.Printf("type %s struct {\n", structName)
	generateFields(data, 1)
	fmt.Printf("}\n")
}

func generateFields(data map[string]interface{}, level int) {
	for key, value := range data {
		indent := ""
		for i := 0; i < level; i++ {
			indent += "\t"
		}

		switch v := value.(type) {
		case map[string]interface{}:
			fmt.Printf("%s%s %s struct {\n", indent, key, "GeneratedStruct")
			generateFields(v, level+1)
			fmt.Printf("%s}\n", indent)
		default:
			fmt.Printf("%s%s %T `json:\"%s\"`\n", indent, key, value, key)
		}
	}
}
