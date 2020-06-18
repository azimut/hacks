package main

import "testing"

func TestProcessSubdomains(t *testing.T) {
	root := "starbucks.com.cn"
	data := &Domain{name: root}
	addAtom("www", data)
	addAtom("email", data)
	addAtom("asdojjaioxi", data)
	addAtom("nnnn.www", data)
	processSubdomains(data.subDomains, root)
}
