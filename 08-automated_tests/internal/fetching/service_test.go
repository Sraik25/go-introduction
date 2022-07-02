package fetching_test

import (
	"errors"
	"testing"

	"github.com/Sraik25/golang-introduction/08-automated_tests/internal/fetching"
	"github.com/Sraik25/golang-introduction/08-automated_tests/internal/storage/inmen"
)

func TestFetchByID(t *testing.T) {

	tests := map[string]struct {
		input int
		want  int
		err   error
	}{
		"valid beer":     {input: 127, want: 127, err: nil},
		"not found beer": {input: 999999, err: errors.New("error")},
	}

	repo := inmen.NewRepository()
	service := fetching.NewService(repo)

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			b, err := service.FetchByID(tc.input)

			if err != nil && tc.err == nil {
				t.Fatalf("not expected any errors and got %v", err)
			}

			if err == nil && tc.err != nil {
				t.Fatalf("expected an error and got nil")
			}

			if b.ProductID != tc.want {
				t.Fatalf("expedted %d, got: %d ", tc.want, b.ProductID)
			}
		})
	}

}
