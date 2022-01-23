package error_handling

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"strings"
)

type Point struct {
	Longitude     string
	Latitude      string
	Distance      string
	ElevationGain string
	ElevationLoss string
}

func parse(r io.Reader) (*Point, error) {
	var p Point

	if err := binary.Read(r, binary.BigEndian, &p.Longitude); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.Latitude); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.Distance); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.ElevationGain); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.ElevationLoss); err != nil {
		return nil, err
	}

	return &Point{}, nil
}

// 用函数式编程的方式
// 但是会带来一个问题，那就是有一个 err 变量和一个内部的函数，感觉不是很干净
func parseFunctionalStyle(r io.Reader) (*Point, error) {
	var p Point
	// global error
	var err error

	read := func(data interface{}) {
		if err != nil {
			// first error first out
			return
		}
		err = binary.Read(r, binary.BigEndian, data)
	}

	read(&p.Longitude)
	read(&p.Latitude)
	read(&p.Distance)
	read(&p.ElevationGain)
	read(&p.ElevationLoss)

	if err != nil {
		return &p, err
	}
	return &p, nil
}

// Learn from Scanner in core go module
// https://pkg.go.dev/bufio
func Learn() {

	const input = "1234 5678 1234567901234567890"
	r := strings.NewReader(input)

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		token := scanner.Text()
		// process token
		fmt.Println("token:", token)
	}

	if err := scanner.Err(); err != nil {
		// process the error
	}
}
