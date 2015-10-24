package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
    var ss []string = strings.Fields(s)
	for _, v := range ss {
		cnt := m[v]
		m[v] = cnt+1
	}
    return m
}

func main() {
	wc.Test(WordCount)
}
