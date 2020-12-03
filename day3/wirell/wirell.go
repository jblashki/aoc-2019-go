package wirell

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Direction is an enumeration describing the directions WireUp, WireDown WireLeft and WireRight
type Direction int

const (
	// WireUp is the direction up in the WireLL
	WireUp = iota
	// WireDown is the direction down in the WireLL
	WireDown
	// WireLeft is the direction left in the WireLL
	WireLeft
	// WireRight is the direction right in the WireLL
	WireRight
)

// WirePos Struct for x,y position of wire
type WirePos struct {
	XPos int
	YPos int
}

// Move is a struct to define a wire move used in WireLL linked list
type Move struct {
	Dir  Direction
	Dist int
}

type wireNode struct {
	pos  WirePos
	next *wireNode
}

// WireLL Linked-list for wire path
type WireLL struct {
	root *wireNode
	tail *wireNode
}

// CreateWireLL Creates a new WireLL linked list
func CreateWireLL() (*WireLL, error) {
	retValue := new(WireLL)

	initNode := new(wireNode)
	initNode.pos.XPos = 0
	initNode.pos.YPos = 0
	initNode.next = nil

	retValue.root = initNode
	retValue.tail = initNode
	return retValue, nil
}

// AddNode Adds new node to WireLL with direction dir
func AddNode(ll *WireLL, dir Direction) error {
	newNode := new(wireNode)

	newNode.pos.XPos = ll.tail.pos.XPos
	newNode.pos.YPos = ll.tail.pos.YPos

	switch dir {
	case WireUp:
		newNode.pos.YPos++

	case WireDown:
		newNode.pos.YPos--

	case WireLeft:
		newNode.pos.XPos--

	case WireRight:
		newNode.pos.XPos++
	}

	newNode.next = nil

	ll.tail.next = newNode
	ll.tail = newNode

	return nil
}

// AddMove will add a series of nodes to the WireLL based on Move
func AddMove(ll *WireLL, m Move) error {
	for i := 0; i < m.Dist; i++ {
		err := AddNode(ll, m.Dir)
		if err != nil {
			return err
		}
	}

	return nil
}

// FindStepsToPos returns the number of steps along a wire to a given x,y position
func FindStepsToPos(ll *WireLL, p WirePos) (int, error) {
	for node, count := ll.root, 0; node != nil; node, count = node.next, count+1 {
		if node.pos.XPos == p.XPos && node.pos.YPos == p.YPos {
			return count, nil
		}
	}

	errormsg := fmt.Sprintf("Pos %v not found in %v", p, ll)
	return 0, errors.New(errormsg)
}

// FindCrossovers returns a slice of WirePos for every postion where ll1 and ll2 cross over
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

// ParseMove converts a string in the format 'U12' to a Move variable
func ParseMove(s string) (Move, error) {
	m := Move{WireUp, 0}
	if s == "" {
		errormsg := fmt.Sprintf("Parsing error. Invalid value %q", s)
		return m, errors.New(errormsg)
	}

	switch s[0] {
	case 'u':
		fallthrough
	case 'U':
		m.Dir = WireUp

	case 'd':
		fallthrough
	case 'D':
		m.Dir = WireDown

	case 'l':
		fallthrough
	case 'L':
		m.Dir = WireLeft

	case 'r':
		fallthrough
	case 'R':
		m.Dir = WireRight

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
