package modifier

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

const testHeader = "X-Correlation-ID"

func TestCorrelationHeaderModifierGenerateNewId(t *testing.T) {
	cfg := fmt.Sprintf(`{"header_name":"%s"}`, testHeader)
	modifier, err := FromJSON([]byte(cfg))
	if err != nil {
		t.Error(err)
		return
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Error("wrong method:", r.Method)
		}

		if r.Header.Get(testHeader) == "" {
			t.Errorf("Test header '%s' cannot be empty", testHeader)
		}
	}))

	req, err := http.NewRequest("POST", ts.URL, ioutil.NopCloser(bytes.NewReader([]byte{})))
	if err != nil {
		t.Error(err)
		return
	}

	if err := modifier.RequestModifier().ModifyRequest(req); err != nil {
		t.Error(err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	if resp.StatusCode != 200 {
		t.Error("unexpected status code:", resp.StatusCode)
		return
	}
}

func TestCorrelationHeaderModifierSettedId(t *testing.T) {
	cfg := fmt.Sprintf(`{"header_name":"%s"}`, testHeader)
	modifier, err := FromJSON([]byte(cfg))
	if err != nil {
		t.Error(err)
		return
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Error("wrong method:", r.Method)
		}

		if r.Header.Get(testHeader) != "xxx" {
			t.Errorf("Test header '%s' is setted before as 'xxx'", testHeader)
		}
	}))

	req, err := http.NewRequest("POST", ts.URL, ioutil.NopCloser(bytes.NewReader([]byte{})))
	if err != nil {
		t.Error(err)
		return
	}

	req.Header.Set(testHeader, "xxx")

	if err := modifier.RequestModifier().ModifyRequest(req); err != nil {
		t.Error(err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	if resp.StatusCode != 200 {
		t.Error("unexpected status code:", resp.StatusCode)
		return
	}
}

func TestHeaderCorrelationModifier_badDSL(t *testing.T) {
	if _, err := FromJSON([]byte(`"x"]}`)); err == nil {
		t.Errorf("error expected")
	}
}
