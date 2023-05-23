package test

import (
	"os"
	"testing"
)

func TestRemove(t *testing.T) {
	path := "../downloadFile/202206/000012/"
	err := os.RemoveAll(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

}
