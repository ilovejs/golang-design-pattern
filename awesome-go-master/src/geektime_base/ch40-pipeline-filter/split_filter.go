package ch40

import (
	"errors"
	"strings"
)

var SplitFilterWrongFormatError = errors.New("input data should be string")

type SplitFilter struct {
	delimiter string
}

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter}
}

func (sf *SplitFilter) Process(data Request) (Response, error) {
	// to type that can be processed by this filter
	str, ok := data.(string)
	if !ok {
		// give up
		return nil, SplitFilterWrongFormatError
	}

	parts := strings.Split(str, sf.delimiter)
	return parts, nil
}
