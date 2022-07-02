package fetching_test

import (
	"testing"

	"github.com/Sraik25/golang-introduction/08-automated_tests/internal/fetching"
	"github.com/Sraik25/golang-introduction/08-automated_tests/internal/storage/inmen"
)

func TestFetchByID(t *testing.T) {
	repo := inmen.NewRepository()

	service := fetching.NewService(repo)

	expected := 127

	b, err := service.FetchByID(expected)

	if err != nil {
		t.Fatalf("expected %d, got an error %v", expected, err)
	}

	if b.ProductID != expected {
		t.Fatalf("expedted %d, got: %d ", expected, b.BeerID)
	}
}
