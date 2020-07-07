package main

func inSlice(stack []string, needle string, glass func(string, string) bool) (found bool) {
	for _, hay := range stack {
		if glass(hay, needle) {
			found = true
		}
	}
	return found
}
