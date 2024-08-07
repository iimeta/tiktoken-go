package main

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/iimeta/tiktoken-go"
)

// go test -benchmem -run=^$ -bench ^BenchmarkEncodingInFullLanguage$ -benchtime=100000x github.com/iimeta/tiktoken-go/test

func BenchmarkEncodingInFullLanguage(b *testing.B) {
	data, err := os.ReadFile("/tmp/udhr.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")
	tkm, err := tiktoken.EncodingForModel("gpt-4o")
	lineCount := len(lines)
	if err != nil {
		log.Fatal(err)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tkm.EncodeOrdinary(lines[n%lineCount])
	}
}
