package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	"aoc2019/day1"
	"aoc2019/day10"
	"aoc2019/day11"
	"aoc2019/day12"
	"aoc2019/day13"
	"aoc2019/day14"
	"aoc2019/day15"
	"aoc2019/day16"
	"aoc2019/day17"
	"aoc2019/day18"
	"aoc2019/day19"
	"aoc2019/day2"
	"aoc2019/day20"
	"aoc2019/day21"
	"aoc2019/day22"
	"aoc2019/day23"
	"aoc2019/day24"
	"aoc2019/day25"
	"aoc2019/day3"
	"aoc2019/day4"
	"aoc2019/day5"
	"aoc2019/day6"
	"aoc2019/day7"
	"aoc2019/day8"
	"aoc2019/day9"
)

var argCalls = map[string]func(bool){
	"DAY1":  day1.RunDay,
	"DAY2":  day2.RunDay,
	"DAY3":  day3.RunDay,
	"DAY4":  day4.RunDay,
	"DAY5":  day5.RunDay,
	"DAY6":  day6.RunDay,
	"DAY7":  day7.RunDay,
	"DAY8":  day8.RunDay,
	"DAY9":  day9.RunDay,
	"DAY10": day10.RunDay,
	"DAY11": day11.RunDay,
	"DAY12": day12.RunDay,
	"DAY13": day13.RunDay,
	"DAY14": day14.RunDay,
	"DAY15": day15.RunDay,
	"DAY16": day16.RunDay,
	"DAY17": day17.RunDay,
	"DAY18": day18.RunDay,
	"DAY19": day19.RunDay,
	"DAY20": day20.RunDay,
	"DAY21": day21.RunDay,
	"DAY22": day22.RunDay,
	"DAY23": day23.RunDay,
	"DAY24": day24.RunDay,
	"DAY25": day25.RunDay,
}

var functionPointers = []func(bool){
	day1.RunDay,
	day2.RunDay,
	day3.RunDay,
	day4.RunDay,
	day5.RunDay,
	day6.RunDay,
	day7.RunDay,
	day8.RunDay,
	day9.RunDay,
	day10.RunDay,
	day11.RunDay,
	day12.RunDay,
	day13.RunDay,
	day14.RunDay,
	day15.RunDay,
	day16.RunDay,
	day17.RunDay,
	day18.RunDay,
	day19.RunDay,
	day20.RunDay,
	day21.RunDay,
	day22.RunDay,
	day23.RunDay,
	day24.RunDay,
	day25.RunDay,
}

var programName = ""

func usage() {
	fmt.Printf("\nUsage: %v [-v] [Day?]...\n", path.Base(os.Args[0]))
	flag.PrintDefaults()
}

func main() {
	programName = path.Base(os.Args[0])
	var verboseFlag = flag.Bool("v", false, "verbose mode")

	flag.Usage = func() {
		usage()
	}

	flag.Parse()

	args := flag.Args()

	if len(args) > 0 {
		// Validate args first
		idx := -1
		var err error = nil
		for i := 0; i < len(args); i++ {
			if len(args[i]) < len("day") {
				idx, err = strconv.Atoi(args[i])
				if err != nil {
					fmt.Printf("Invalid argument %q\n", args[i])
					usage()
					return
				}
				idx--
			} else {
				if strings.ToUpper(args[i][:3]) != "DAY" {
					idx, err = strconv.Atoi(args[i])
					if err != nil {
						fmt.Printf("Invalid argument %q\n", args[i])
						usage()
						return
					}
					idx--
				} else {
					idx, err = strconv.Atoi(args[i][3:])
					if err != nil {
						fmt.Printf("Invalid argument %q\n", args[i])
						usage()
						return
					}
					idx--
				}
			}

			if idx < 0 || idx >= len(functionPointers) {
				fmt.Printf("InXalid argument %q\n", args[i])
				usage()
				return
			}
		}
	}

	fmt.Println("Advent of Code (Go) 2019")
	fmt.Println("========================")

	if len(args) > 0 {
		for i := 0; i < len(args); i++ {
			idx := -1
			if len(args[i]) < len("day") {
				idx, _ = strconv.Atoi(args[i])
				idx--
			} else {
				if strings.ToUpper(args[i][:3]) != "DAY" {
					idx, _ = strconv.Atoi(args[i])
					idx--
				} else {
					idx, _ = strconv.Atoi(args[i][3:])
					idx--
				}
			}
			call := functionPointers[idx]

			call(*verboseFlag)
		}
	} else {
		// Call all options
		for i := 0; i < len(functionPointers); i++ {
			call := functionPointers[i]
			call(*verboseFlag)
		}
	}
}
