package main

import (
	"strings"
	"testing"
)

type Modifier interface {
	Modify() string
}

type Original struct {
	Value string
}

func (o *Original) Modify() string {
	return o.Value
}

// Upper возвращает строку в верхнем регистре.
type Upper struct {
	modifier Modifier
}

func (u *Upper) Modify() string {
	return strings.ToUpper(u.modifier.Modify())
}

// Replace заменяет строки old на new.
type Replace struct {
	modifier Modifier
	old      string
	new      string
}

func (r *Replace) Modify() string {
	return strings.ReplaceAll(r.modifier.Modify(), r.old, r.new)
}

func TestModifier(t *testing.T) {
	original := &Original{Value: "Привет, гофер!"}
	replace := &Replace{
		modifier: original,
		old:      "гофер",
		new:      "мир",
	}
	upper := &Upper{
		modifier: replace,
	}
	if upper.Modify() != "ПРИВЕТ, МИР!" {
		t.Errorf(`get %s`, upper.Modify())
	}
}
