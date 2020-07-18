package jmdict

import (
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	dictF, err := os.Open("JMdict_e")
	if err != nil {
		t.Error(err)
	}
	_, err = Parse(dictF)
	if err != nil {
		t.Error(err)
	}
}
