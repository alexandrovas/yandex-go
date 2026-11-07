package main

import (
	"bytes"
	"strings"
	"testing"
)

const (
	maxSize = 1000 // будем выравнивать строку до 1000 знаков
	leader  = '#'
)

func naiveAlign(s string, length int, lead rune) string {
	for len(s) < length {
		s = string(lead) + s
	}
	return s
}

func bufferAlign(s string, length int, lead rune) string {
	buf := bytes.Buffer{}
	for i := 0; i < length-len(s); i++ {
		buf.WriteRune(lead)
	}
	buf.WriteString(s)
	return buf.String()
}

func deltaAlign(s string, length int, lead rune) string {
	if len(s) < length {
		return strings.Repeat(string(lead), length-len(s)) + s
	}
	return s
}

func BenchmarkNaive(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		naiveAlign("", maxSize, leader)
	}
}

func BenchmarkBuffer(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		bufferAlign("", maxSize, leader)
	}
}

func BenchmarkDelta(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		deltaAlign("", maxSize, leader)
	}
}
