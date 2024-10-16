package apicore

import (
	"testing"
)

func TestApiCore(t *testing.T) {
	result := ApiCore("works")
	if result != "ApiCore works" {
		t.Error("Expected ApiCore to append 'works'")
	}
}
