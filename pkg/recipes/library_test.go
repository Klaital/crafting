package recipes

import "testing"

func TestLoadLibraryFromFile(t *testing.T) {
	lib, err := LoadLibraryFromFile("../../testdata/items/seablock.json")
	if err != nil {
		t.Errorf("Error loading: %v", err)
		t.Fail()
		return
	}

	if lib == nil {
		t.Error("Null pointer returned")
		t.Fail()
		return
	}

	if len(lib.Items) == 0 {
		t.Error("No data loaded")
		t.Fail()
		return
	}

	if len(lib.Items) < 12 {
		t.Errorf("Not all data loaded. Expected at least 12, got %d", len(lib.Items))
		t.Fail()
		return
	}
}
