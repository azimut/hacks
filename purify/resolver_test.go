package main

import "testing"

func TestHasWildcard(t *testing.T) {
	domain := "starbucks.com.cn"
	wildcard, wildrecords, err := hasWildcard(domain)
	if err != nil {
		t.Error(err)
	}
	if !wildcard {
		t.Error("faild to identify wildcard")
	}
	if len(wildrecords) == 0 {
		t.Error("zero record returned for wildcards")
	}
}
