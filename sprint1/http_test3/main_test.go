package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFamilyAddNew(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)
	family := Family{
		Members: map[Relationship]Person{
			Child: {FirstName: "Misha", LastName: "Popov", Age: 20},
		},
	}

	tests := []struct {
		name      string
		relation  Relationship
		person    Person
		wantError bool
	}{
		{
			name:     "Test name #1",
			relation: Mother,
			person:   Person{FirstName: "Olga", LastName: "Ivanova", Age: 40},
		},
		{
			name:      "Test name #2",
			relation:  Child,
			person:    Person{FirstName: "Ivan", LastName: "Ivanov", Age: 16},
			wantError: true,
		},
	}
	for _, test := range tests { // цикл по всем тестам
		t.Run(test.name, func(t *testing.T) {
			err := family.AddNew(test.relation, test.person)
			if !test.wantError {
				require.NoError(err)
				assert.Contains(family.Members, test.relation)
				return
			}
			assert.Error(err)
		})
	}
}
