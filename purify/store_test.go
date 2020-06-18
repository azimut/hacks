package main

import (
	"fmt"
	"testing"
)

func TestGetAtom(t *testing.T) {
	data := &Domain{name: "google.com"}
	data.subDomains = append(data.subDomains, &Domain{name: "www"})
	data.subDomains = append(data.subDomains, &Domain{name: "ftp"})
	data.subDomains = append(data.subDomains, &Domain{name: "dev"})
	data.subDomains[2].subDomains =
		append(data.subDomains[2].subDomains, &Domain{name: "www"})
	data.subDomains[2].subDomains[0].subDomains =
		append(data.subDomains[2].subDomains[0].subDomains, &Domain{name: "www"})
	//
	if getAtom("www", data) == nil {
		t.Errorf("www not found")
	}
	if getAtom("www.dev", data) != data.subDomains[2].subDomains[0] {
		t.Errorf("www.dev not found")
	}
	if getAtom("www.www.dev", data) != data.subDomains[2].subDomains[0].subDomains[0] {
		t.Errorf("www.www.dev not found")
	}
}

func TestAddAtom(t *testing.T) {
	data := &Domain{name: "google.com"}
	addAtom("www", data)
	addAtom("ftp", data)
	addAtom("dev", data)
	addAtom("www.dev", data)
	addAtom("www.www.dev", data)
	fmt.Println("---------")
	fmt.Println(data)
	fmt.Println("---------")
	//
	if getAtom("www", data) == nil {
		t.Errorf("www not found")
	}
	if getAtom("www.dev", data) != data.subDomains[2].subDomains[0] {
		t.Errorf("www.dev not found, %v", data.subDomains[2].subDomains[0])
	}
	if getAtom("www.www.dev", data) != data.subDomains[2].subDomains[0].subDomains[0] {
		t.Errorf("www.www.dev not found")
	}
}
