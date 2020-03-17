package main

import (
	"testing"
)

func TestValidMapping(t *testing.T) {
	s1 := "123"
	s2 := "321"
	result := isMappingExists(s1, s2)
	if !result {
		t.Errorf("isMappingExists(%s,%s) = false; want true", s1, s2)
	}
}

func TestInvalidMapping(t *testing.T) {
	s1 := "foo"
	s2 := "bar"
	result := isMappingExists(s1, s2)
	if result {
		t.Errorf("isMappingExists(%s,%s) = true; want false", s1, s2)
	}
}
