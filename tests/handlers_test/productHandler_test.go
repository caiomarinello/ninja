package handlers_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"testing"

	comp "ninja/caio/api/components"
	hdl "ninja/caio/api/handlers"

	"github.com/gin-gonic/gin"
)

// MockFetcher is a mock implementation of the Fetcher interface.
type MockFetcher struct {
	products []comp.Product
	err     error
}

func (m *MockFetcher) FetchById(id int) (comp.Product, error) {
	return comp.Product{}, m.err
}

func (m *MockFetcher) FetchAll() ([]comp.Product, error) {
	return m.products, m.err
}

func TestHandleGetAllProducts(t *testing.T) {
	// Set Gin mode to test to suppress unnecessary logs
	gin.SetMode(gin.TestMode)
	
	t.Run("Successful fetch all products", func(t *testing.T) {
		c, rec, err := setupTestContextRec("GET", "/products", struct{}{})
		if err != nil {
			t.Errorf("error on setup test context: %v", err)
		}

		mockFetcher := &MockFetcher{
			products: []comp.Product{
				{ProductId: 1, Name: "Product 1", Description: "Description 1", Price: 10.0, Category: "Category 1", Stock: 10},
				{ProductId: 2, Name: "Product 2", Description: "Description 2", Price: 20.0, Category: "Category 2", Stock: 20},
			},
			err: nil,
		}

		hdl.HandleGetAllProducts(mockFetcher)(c)

		if rec.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rec.Code)
		}

		var actualProducts []comp.Product
        err = json.NewDecoder(rec.Body).Decode(&actualProducts)
        if err != nil {
            t.Errorf("error decoding response body: %v", err)
            return
        }

        if !reflect.DeepEqual(mockFetcher.products, actualProducts) {
            t.Errorf("expected response body: %v, got: %v", mockFetcher.products, actualProducts)
        }
	})

	t.Run("Error fetching all products", func(t *testing.T) {
		c, rec, err := setupTestContextRec("GET", "/products", struct{}{})
		if err != nil {
			t.Errorf("error on setup test context: %v", err)
		}

		mockFetcher := &MockFetcher{
			products: nil,
			err:      errors.New("failed to fetch all products"),
		}

		hdl.HandleGetAllProducts(mockFetcher)(c)

		if rec.Code != http.StatusInternalServerError {
			t.Errorf("expected status code %d, got %d", http.StatusInternalServerError, rec.Code)
		}
	})
}