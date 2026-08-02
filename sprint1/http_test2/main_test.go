package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullName(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name      string
		firstName string
		lastName  string
		want      string
	}{
		{
			name:      "Test #1",
			firstName: "Alex",
			lastName:  "Alexandrov",
			want:      "Alex Alexandrov",
		},
	}
	for _, test := range tests { // цикл по всем тестам
		t.Run(test.name, func(t *testing.T) {
			u := User{FirstName: test.firstName, LastName: test.lastName}
			assert.Equal(test.want, u.FullName())
		})
	}
}
