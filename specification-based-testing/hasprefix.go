package main

// extracted from Go's source code at [https://cs.opensource.google/go/go/+/refs/tags/go1.24.2:src/internal/stringslite/strings.go;l=16]
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}
