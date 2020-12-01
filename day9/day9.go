package day9

import (
	"errors"
	"fmt"
	"sync"

	"github.com/jblashki/aoc-intcode-go/v3"
)

const name = "Day 9"

const input_file = "./day9/program"

func RunDay(verbose bool) {
	var aResult int
	var bResult int
	var err error

	if verbose {
		fmt.Printf("\n%v Output:\n", name)
	}

	aResult, err = a(verbose)
	if err != nil {
		fmt.Printf("%va: **** Error: %q ****\n", name, err)
	} else {
		fmt.Printf("%va: Program Result = %v\n", name, aResult)
	}

	bResult, err = b(verbose)
	if err != nil {
		fmt.Printf("%vb: **** Error: %q ****\n", name, err)
	} else {
		fmt.Printf("%vb: Program Result = %v\n", name, bResult)
	}
}

func a(verbose bool) (int, error) {
	value, err := boost(1, verbose)
	return value, err
}

func b(verbose bool) (int, error) {
	value, err := boost(2, verbose)
	return value, err
}

func boost(input int, verbose bool) (int, error) {
	ic := intcode.Create()

	err := intcode.Load(ic, input_file)
	if err != nil {
		return 0, err
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)

	inputChan := make(chan int, 1)
	outputChan := make(chan int)
	haltSignalChan := make(chan int, 1)
	errorMsgChan := make(chan string, 1)

	defer func() {
		close(inputChan)
		close(outputChan)
		close(haltSignalChan)
		close(errorMsgChan)
	}()

	inputChan <- input

	go intcode.Run(ic, inputChan, outputChan, haltSignalChan, errorMsgChan, wg, "")

	value := -1
	exit := false
	for !exit {
		select {
		case value = <-outputChan:
			if verbose {
				fmt.Printf("> %v\n", value)
			}
		case hltValue := <-haltSignalChan:
			if hltValue == 1 {
				err := <-errorMsgChan
				errormsg := fmt.Sprintf("Program halted with error: %v", err)
				return 0, errors.New(errormsg)
			}
			exit = true
		}
	}

	wg.Wait()

	return value, nil
}
