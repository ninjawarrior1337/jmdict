// Package jmdict implements a parser for the JMdict Japanese-Multilingual dictionary.
// The JMdict files are available from http://www.edrdg.org/jmdict/j_jmdict.html.
package jmdict

import (
	"encoding/xml"
	"io"
)

// Parse parses the JMdict file from r.
func Parse(r io.Reader) (result *JMdict, err error) {
	d := xml.NewDecoder(r)
	d.Entity = entity
	if err := d.Decode(&result); err != nil {
		return nil, err
	}
	return
}
