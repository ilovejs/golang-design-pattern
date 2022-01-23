package error_handling

import (
	"encoding/binary"
	"io"
)

// Demo: Fluent Interface
// https://martinfowler.com/bliki/FluentInterface.html

type Reader struct {
	r   io.Reader
	err error // embed in struct rather than main
}

func (r *Reader) read(data interface{}) {
	if r.err == nil {
		r.err = binary.Read(r.r, binary.BigEndian, data)
	}
}

// but few ppl use io.Reader as param... i think
func parseVersionBetter(input io.Reader) (*Point, error) {
	var p Point
	r := Reader{r: input}

	r.read(&p.Longitude)
	r.read(&p.Latitude)
	r.read(&p.Distance)
	r.read(&p.ElevationGain)
	r.read(&p.ElevationLoss)

	if r.err != nil {
		return nil, r.err
	}

	return &p, nil
}

// TODO: read Scanner source code
// extract error handling into Err function
