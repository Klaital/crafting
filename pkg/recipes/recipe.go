package recipes

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Component struct {
	MaterialID ItemId
	Material   Item
	Quantity   uint
}

type Recipe struct {
	Name        string
	Ingredients []Component
	Results     []Component
	Seconds     float64
}

func LoadFromJsonFile(path string) (library []Recipe, err error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	recipes := make([]Recipe, 0)
	err = json.Unmarshal(jsonBytes, &recipes)
	return recipes, err
}
