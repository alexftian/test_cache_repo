package lru

import (
	"testing"
)

type TestValue struct {
	context string
}
func TestGet(t *testing.T) {
	lru := New(20, nil)
	lru.Add("a", "1234")
	lru.Add("b", "5678")
}