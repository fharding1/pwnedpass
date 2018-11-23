package pwnedpass

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientV2Count(t *testing.T) {
	mockHandler := func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
	}

	mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mockHandler(w, r)
	}))
	defer mock.Close()

	client := ClientV2{
		HTTPClient: http.DefaultClient,
		BaseURL:    mock.URL,
	}

	testCases := []struct {
		name          string
		handler       http.HandlerFunc
		password      string
		expectedCount int
		expectError   bool
	}{
		{
			name:          "empty response",
			handler:       func(w http.ResponseWriter, r *http.Request) {},
			password:      "foo",
			expectedCount: 0,
			expectError:   false,
		},
		{
			name: "password hash does not exist in response",
			handler: func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "0018A45C4D1DEF81644B54AB7F969B88D65:12")
				fmt.Fprintln(w, "00D4F6E8FA6EECAD2A3AA415EEC418D38EC:12")
			},
			password:      "foo",
			expectedCount: 0,
			expectError:   false,
		},
		{
			name: "empty password",
			handler: func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "0018A45C4D1DEF81644B54AB7F969B88D65:12")
				fmt.Fprintln(w, "00D4F6E8FA6EECAD2A3AA415EEC418D38EC:12")
			},
			password:      "",
			expectedCount: 0,
			expectError:   false,
		},
		{
			name: "password hash exists in response and has count of 12",
			handler: func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "0018A45C4D1DEF81644B54AB7F969B88D65:4")
				fmt.Fprintln(w, "7B5EA3F0FDBC95D0DD47F3C5BC275DA8A33:12")
				fmt.Fprintln(w, "00D4F6E8FA6EECAD2A3AA415EEC418D38EC:9")
			},
			password:      "foo",
			expectedCount: 12,
			expectError:   false,
		},
		{
			name: "password hash exists in lowercase results",
			handler: func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "0018a45c4d1def81644b54ab7f969b88d65:4")
				fmt.Fprintln(w, "7b5ea3f0fdbc95d0dd47f3c5bc275da8a33:12")
				fmt.Fprintln(w, "00d4f6e8fa6eecad2a3aa415eec418d38ec:9")
			},
			password:      "foo",
			expectedCount: 12,
			expectError:   false,
		},
		{
			name: "password hash exists and has count of 12",
			handler: func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "0018A45C4D1DEF81644B54AB7F969B88D65:4")
				fmt.Fprintln(w, "7B5EA3F0FDBC95D0DD47F3C5BC275DA8A33:12")
				fmt.Fprintln(w, "00D4F6E8FA6EECAD2A3AA415EEC418D38EC:9")
			},
			password:      "foo",
			expectedCount: 12,
			expectError:   false,
		},
		{
			name: "password hash exists with invalid count",
			handler: func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "0018A45C4D1DEF81644B54AB7F969B88D65:4")
				fmt.Fprintln(w, "7B5EA3F0FDBC95D0DD47F3C5BC275DA8A33:twelve")
				fmt.Fprintln(w, "00D4F6E8FA6EECAD2A3AA415EEC418D38EC:9")
			},
			password:      "foo",
			expectedCount: 0,
			expectError:   true,
		},
		{
			name: "password hash exists ignores weird rows",
			handler: func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "bla:foo:0018A45C4D1DEF81644B54AB7F969B88D65:4")
				fmt.Fprintln(w, ":4")
				fmt.Fprintln(w, "")
				fmt.Fprintln(w, "7B5EA3F0FDBC95D0DD47F3C5BC275DA8A33:12")
				fmt.Fprintln(w, "00D4F6E8FA6EECAD2A3AA415EEC418D38EC:9")
				fmt.Fprintln(w, "")
			},
			password:      "foo",
			expectedCount: 12,
			expectError:   false,
		},
		{
			name: "password hash exists in large response",
			handler: func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, testData)
			},
			password:      "password",
			expectedCount: 3533661,
			expectError:   false,
		},
	}

	for i, tt := range testCases {
		t.Logf("test case %d: %s", i, tt.name)
		mockHandler = tt.handler

		count, err := client.Count(tt.password)
		assert.Equal(t, tt.expectError, err != nil)
		assert.Equal(t, tt.expectedCount, count)
	}
}

func BenchmarkCount(b *testing.B) {
	mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, testData)
	}))
	defer mock.Close()

	client := ClientV2{
		HTTPClient: http.DefaultClient,
		BaseURL:    mock.URL,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.Count("password")
	}
}

func BenchmarkCountIntegration(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DefaultClient.Count("password")
	}
}
