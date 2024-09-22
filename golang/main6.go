package main

type Person1 struct {
	Name    string
	Age     int
	Friends []string
}

func deepCopy(p Person1) Person1 {
	copyP := p
	copyP.Friends = make([]string, len(p.Friends))
	copy(copyP.Friends, p.Friends)
	return copyP
}
