package prototype

import "testing"

func TestClone(t *testing.T) {
	shirtCache := GetShirtCloner()
	if shirtCache == nil {
		t.Fatal("Shirtcache was nil")
	}
	item1, err := shirtCache.GetClone(White)
	if err != nil {
		t.Error(err)
	}

	shirt1, ok := item1.(*Shirt)
	if !ok {
		t.Fatal("Type Assertion form Shirt could be done successfully for item1")
	}

	shirt1.SKU = "abbcc"

	item2, err := shirtCache.GetClone(White)
	if err != nil {
		t.Error(err)
	}

	shirt2, ok := item2.(*Shirt)
	if !ok {
		t.Fatal("Type Assertion form Shirt could be done successfully for item2")
	}

	if shirt1.SKU == shirt2.SKU {
		t.Error("Sku for shirt1 and shirt 2 must be different")
	}

	if shirt1 == shirt2 {
		t.Error("shirt1 cannot be equal to shirt2")
	}

}
