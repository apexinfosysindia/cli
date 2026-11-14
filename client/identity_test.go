package client

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/viper"
)

func TestUserAgent(t *testing.T) {
	old := Version
	t.Cleanup(func() { Version = old })

	Version = "5.6.0"
	if got, want := userAgent(), "ApexOSCLI/5.6.0"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	// Test binaries carry no module version
	Version = ""
	if got, want := userAgent(), "ApexOSCLI/dev"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestIdentityHeadersSent(t *testing.T) {
	tests := []struct {
		name       string
		apexSource string
	}{
		{"no source", ""},
		{"source set", "my_app"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotUserAgent string
			var gotSource []string
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotUserAgent = r.Header.Get("User-Agent")
				gotSource = r.Header.Values("X-Apex-Source")
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"result":"ok","data":{}}`))
			}))
			t.Cleanup(server.Close)

			viper.SetDefault("endpoint", server.Listener.Addr().String())
			viper.SetDefault("apex-source", tt.apexSource)
			t.Cleanup(func() { viper.SetDefault("apex-source", "") })
			client = nil
			t.Cleanup(func() { client = nil })

			if _, err := GenericJSONGet("supervisor", "info"); err != nil {
				t.Fatal(err)
			}
			if want := userAgent(); gotUserAgent != want {
				t.Errorf("User-Agent: got %q, want %q", gotUserAgent, want)
			}
			if tt.apexSource == "" {
				if gotSource != nil {
					t.Errorf("X-Apex-Source: got %q, want none", gotSource)
				}
			} else if len(gotSource) != 1 || gotSource[0] != tt.apexSource {
				t.Errorf("X-Apex-Source: got %q, want %q", gotSource, tt.apexSource)
			}
		})
	}
}
