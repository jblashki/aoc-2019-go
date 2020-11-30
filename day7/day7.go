package day7

import (
	"fmt"
	"sync"

	"github.com/jblashki/aoc-intcode-go/v2"
)

const name = "Day 7"
const INPUT_FILE = "./day7/program"

func RunDay(verbose bool) {
	var maxSignalA int
	var maxSignalSeqA []int
	var maxSignalB int
	var maxSignalSeqB []int
	var err error

	if verbose {
		fmt.Printf("\n%v Output:\n", name)
	}

	maxSignalA, maxSignalSeqA, err = a()
	if err != nil {
		fmt.Printf("%va: **** Error: %q ****\n", name, err)
	} else {
		fmt.Printf("%va: Max thruster signal = %v (%v)\n", name, maxSignalA, maxSignalSeqA)
	}

	maxSignalB, maxSignalSeqB, err = b()
	if err != nil {
		fmt.Printf("%vb: **** Error: %q ****\n", name, err)
	} else {
		fmt.Printf("%vb: Max thruster signal = %v (%v)\n", name, maxSignalB, maxSignalSeqB)
	}
}

func a() (int, []int, error) {
	ic := intcode.Create()

	err := intcode.Load(ic, INPUT_FILE)
	if err != nil {
		return 0, nil, err
	}

	perms := generatePermutations([]int{0, 1, 2, 3, 4})

	maxSignal := -1
	var maxSignalPerm []int = nil
	for i := 0; i < len(perms); i++ {
		perm := perms[i]
		output, err := amplifierOutput(ic, perm)
		if err != nil {
			return 0, nil, err
		}
		if maxSignal == -1 || output > maxSignal {
			maxSignal = output
			maxSignalPerm = perm
		}
	}

	return maxSignal, maxSignalPerm, nil
}

func b() (int, []int, error) {
	ic := intcode.Create()

	err := intcode.Load(ic, INPUT_FILE)
	if err != nil {
		return 0, nil, err
	}

	perms := generatePermutations([]int{5, 6, 7, 8, 9})

	maxSignal := -1
	var maxSignalPerm []int = nil
	for i := 0; i < len(perms); i++ {
		perm := perms[i]
		output, err := amplifierOutput(ic, perm)
		if err != nil {
			return 0, nil, err
		}
		if maxSignal == -1 || output > maxSignal {
			maxSignal = output
			maxSignalPerm = perm
		}
	}

	return maxSignal, maxSignalPerm, nil
}

func amplifierOutput(ic *intcode.IntCode, phaseSetting []int) (int, error) {
	icA := intcode.Copy(ic)
	icB := intcode.Copy(ic)
	icC := intcode.Copy(ic)
	icD := intcode.Copy(ic)
	icE := intcode.Copy(ic)

	wg := new(sync.WaitGroup)
	wg.Add(5)

	chanAIn := make(chan int)
	chanAOut := make(chan int)
	chanAHlt := make(chan int, 1)

	chanBIn := make(chan int)
	chanBOut := make(chan int)
	chanBHlt := make(chan int, 1)

	chanCIn := make(chan int)
	chanCOut := make(chan int)
	chanCHlt := make(chan int, 1)

	chanDIn := make(chan int)
	chanDOut := make(chan int)
	chanDHlt := make(chan int, 1)

	chanEIn := make(chan int)
	chanEOut := make(chan int)
	chanEHlt := make(chan int, 1)

	defer func() {
		wg.Wait()
		close(chanAIn)
		close(chanAOut)
		close(chanAHlt)
		close(chanBIn)
		close(chanBOut)
		close(chanBHlt)
		close(chanCIn)
		close(chanCOut)
		close(chanCHlt)
		close(chanDIn)
		close(chanDOut)
		close(chanDHlt)
		close(chanEIn)
		close(chanEOut)
		close(chanEHlt)
	}()

	go intcode.Run(icA, chanAIn, chanAOut, chanAHlt, wg)
	go intcode.Run(icB, chanBIn, chanBOut, chanBHlt, wg)
	go intcode.Run(icC, chanCIn, chanCOut, chanCHlt, wg)
	go intcode.Run(icD, chanDIn, chanDOut, chanDHlt, wg)
	go intcode.Run(icE, chanEIn, chanEOut, chanEHlt, wg)

	chanAIn <- phaseSetting[0]
	chanBIn <- phaseSetting[1]
	chanCIn <- phaseSetting[2]
	chanDIn <- phaseSetting[3]
	chanEIn <- phaseSetting[4]

	output := 0
	for {
		// If Amplifier A has halted we should get
		// a signal on the chanAHlt channel
		select {
		case chanAIn <- output:
		case _ = <-chanAHlt:
			return output, nil
		}
		output = <-chanAOut

		chanBIn <- output
		output = <-chanBOut

		chanCIn <- output
		output = <-chanCOut

		chanDIn <- output
		output = <-chanDOut

		chanEIn <- output
		output = <-chanEOut
	}

	return output, nil
}

func generatePermutations(input []int) [][]int {
	returnPerms := make([][]int, 0)
	a := make([]int, len(input))
	copy(a, input)

	returnPerms = append(returnPerms, a)

	N := len(input)
	p := make([]int, N)

	i := 1
	j := 0
	for i < N {
		if p[i] < i {
			if i%2 != 0 {
				j = p[i]
			} else {
				j = 0
			}
			newA := make([]int, len(input))
			copy(newA, a)
			tmp := newA[j]
			newA[j] = newA[i]
			newA[i] = tmp
			returnPerms = append(returnPerms, newA)
			a = newA
			p[i] += 1
			i = 1
		} else {
			p[i] = 0
			i++
		}
	}

	return returnPerms
}
