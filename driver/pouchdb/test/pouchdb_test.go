package test

import (
	"testing"

	"github.com/flimzy/kivik"
	_ "github.com/flimzy/kivik/driver/pouchdb"
)

const TestServer = ""
const ExpectedVersion = "6.1.2"

func TestVersion(t *testing.T) {
	s, err := kivik.New("pouchdb", TestServer)
	if err != nil {
		t.Fatalf("Error connecting to %s: %s\n", TestServer, err)
	}
	version, err := s.Version()
	if err != nil {
		t.Fatalf("Failed to get server info: %s", err)
	}
	if ExpectedVersion != version {
		t.Errorf("Server version.\n\tExpected: %s\n\t  Actual: %s\n", ExpectedVersion, version)
	}
}
