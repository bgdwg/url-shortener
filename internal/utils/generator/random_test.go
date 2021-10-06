package generator

import (
	"regexp"
	"testing"
)

var (
	validKey = regexp.MustCompile("^[a-zA-Z0-9_]*$")
)

func TestGetRandomKey(t *testing.T) {
	randomKey := GetRandomKey()
	if !validKey.MatchString(randomKey) {
		t.Errorf("key must contain a-z, A-Z, 0-9 and _")
	}
}
