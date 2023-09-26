package domain

import (
	"testing"
)

type Entity struct {
	ID   int
	Name string
}

func TestEmptyPage(t *testing.T) {
	page := EmptyPage[Entity]()
	if len(page.Items) != 0 {
		t.Errorf("Expected 0 items, but got %d", len(page.Items))
	}
	if page.Total != 0 || page.PageSize != 0 || page.PageNumber != 0 || page.TotalPages != 0 {
		t.Errorf("Expected all values to be 0")
	}
}

func TestPageBuilder(t *testing.T) {
	builder := &PageBuilder[Entity]{}
	items := []Entity{
		{ID: 1, Name: "Item 1"},
		{ID: 2, Name: "Item 2"},
	}

	page := builder.SetItems(items).
		SetTotal(10).
		SetPageSize(5).
		SetPageNumber(2).
		SetTotalPages(2).
		Build()

	if len(page.Items) != 2 {
		t.Errorf("Expected 2 items, but got %d", len(page.Items))
	}
	if page.Items[0].Name != "Item 1" {
		t.Errorf("Expected name of first item to be 'Item 1', but got %s", page.Items[0].Name)
	}
	if page.Total != 10 || page.PageSize != 5 || page.PageNumber != 2 || page.TotalPages != 2 {
		t.Errorf("Values do not match the builder settings")
	}
}
