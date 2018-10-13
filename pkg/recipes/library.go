package recipes

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Library struct {
	Items map[ItemId]Item
}

func LoadLibraryFromFile(path string) (lib *Library, err error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	items := make([]Item, 0)
	err = json.Unmarshal(jsonBytes, &items)
	if err != nil {
		return nil, err
	}

	// Reformat into an array keyed by the item ID
	newLibraryData := make(map[ItemId]Item)
	for _, item := range items {
		newLibraryData[item.ID] = item
	}

	return &Library{Items: newLibraryData}, nil
}
