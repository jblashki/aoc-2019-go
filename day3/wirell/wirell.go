package wirell

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Direction int

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)

type WirePos struct {
	XPos int
	YPos int
}

type Move struct {
	Dir  Direction
	Dist int
}

type WireNode struct {
	pos  WirePos
	next *WireNode
}

type WireLL struct {
	root *WireNode
	tail *WireNode
}

func CreateWireLL() (*WireLL, error) {
	retValue := new(WireLL)

	initNode := new(WireNode)
	initNode.pos.XPos = 0
	initNode.pos.YPos = 0
	initNode.next = nil

	retValue.root = initNode
	retValue.tail = initNode
	return retValue, nil
}

func AddNode(ll *WireLL, dir Direction) error {
	newNode := new(WireNode)

	newNode.pos.XPos = ll.tail.pos.XPos
	newNode.pos.YPos = ll.tail.pos.YPos

	switch dir {
	case UP:
		newNode.pos.YPos++

	case DOWN:
		newNode.pos.YPos--

	case LEFT:
		newNode.pos.XPos--

	case RIGHT:
		newNode.pos.XPos++
	}

	newNode.next = nil

	ll.tail.next = newNode
	ll.tail = newNode

	return nil
}

func AddMove(ll *WireLL, m Move) error {
	for i := 0; i < m.Dist; i++ {
		err := AddNode(ll, m.Dir)
		if err != nil {
			return err
		}
	}

	return nil
}

func FindStepsToPos(ll *WireLL, p WirePos) (int, error) {
	for node, count := ll.root, 0; node != nil; node, count = node.next, count+1 {
		if node.pos.XPos == p.XPos && node.pos.YPos == p.YPos {
			return count, nil
		}
	}

	errormsg := fmt.Sprintf("Pos %v not found in %ll", p, ll)
	return 0, errors.New(errormsg)
}

func FindCrossovers(ll1 *WireLL, ll2 *WireLL) []WirePos {
	returnWirePos := make([]WirePos, 0)

	for node := ll1.root; node != nil; node = node.next {
		_, err := FindStepsToPos(ll2, node.pos)
		if err == nil {
			// FOUND -- Ignore 0,0
			if node.pos.XPos != 0 || node.pos.YPos != 0 {
				//fmt.Printf("Found (%v, %v)\n", node.pos.XPos, node.pos.YPos)
				returnWirePos = append(returnWirePos, node.pos)
			}
		}
	}

	return returnWirePos
}

func ParseMove(s string) (Move, error) {
	m := Move{UP, 0}
	if s == "" {
		errormsg := fmt.Sprintf("Parsing error. Invalid value %q", s)
		return m, errors.New(errormsg)
	}

	switch s[0] {
	case 'u':
		fallthrough
	case 'U':
		m.Dir = UP

	case 'd':
		fallthrough
	case 'D':
		m.Dir = DOWN

	case 'l':
		fallthrough
	case 'L':
		m.Dir = LEFT

	case 'r':
		fallthrough
	case 'R':
		m.Dir = RIGHT

	default:
		errormsg := fmt.Sprintf("Invalid direction '%c' in %q", s[0], s)
		return m, errors.New(errormsg)
	}

	var err error
	m.Dist, err = strconv.Atoi(strings.TrimSpace(s[1:]))
	if err != nil {
		return m, err
	}

	return m, nil
}

func (d Direction) String() string {
	return [...]string{"Up", "Down", "Left", "Right"}[d]
}
