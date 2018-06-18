package routing

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewBLRouter(t *testing.T) {
	r := NewBLRouter()
	srv := httptest.NewServer(r)
	defer srv.Close()

	{
		resp, error := http.Get(srv.URL + "/home")
		if error != nil {
			t.Fatal(error)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Status code is %v, but %v expected", resp.StatusCode, http.StatusOK)
		}
	}

	{
		resp, error := http.Get(srv.URL + "/abc")
		if error != nil {
			t.Fatal(error)
		}

		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("Status code is %v, but %v expected", resp.StatusCode, http.StatusNotFound)
		}
	}

}
