package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/goark/mt/mt19937"
	"github.com/goark/mt/v2"
)

func main() {
	flag.Parse()
	inp := bufio.NewReader(os.Stdin)
	switch flag.Arg(0) {
	case "encode":
		fmt.Println("Input string:")
		in, err := inp.ReadString('\n')
		if err != nil {
			panic(err)
		}
		in = in[:len(in)-1]

		fmt.Println("Input seed:")
		_seed, err := inp.ReadString('\n')
		if err != nil {
			panic(err)
		}

		_seed = _seed[:len(_seed)-1]
		seed, err := strconv.ParseInt(_seed, 10, 64)
		if err != nil {
			panic(err)
		}

		fmt.Println("Input step:")
		_step, err := inp.ReadString('\n')
		if err != nil {
			panic(err)
		}

		_step = _step[:len(_step)-1]
		step, err := strconv.ParseFloat(_step, 64)
		if err != nil {
			panic(err)
		}

		rng := mt.New(mt19937.New(seed))

		n := make([]rune, len(in))

		for i, r := range in {
			n[i] = r + rune(rng.Real(1)*step)
			println(r, n[i])
		}

		fmt.Println(string(n))
		os.WriteFile("out.txt", []byte(string(n)), 0644)
	case "decode":
		data, err := os.ReadFile("out.txt")
		if err != nil {
			panic(err)
		}

		in := string(data)
		fmt.Println("Input seed:")
		_seed, err := inp.ReadString('\n')
		if err != nil {
			panic(err)
		}

		_seed = _seed[:len(_seed)-1]
		seed, err := strconv.ParseInt(_seed, 10, 64)
		if err != nil {
			panic(err)
		}

		fmt.Println("Input step:")
		_step, err := inp.ReadString('\n')
		if err != nil {
			panic(err)
		}

		_step = _step[:len(_step)-1]
		step, err := strconv.ParseFloat(_step, 64)
		if err != nil {
			panic(err)
		}

		rng := mt.New(mt19937.New(seed))

		n := make([]rune, len(in))

		for i, r := range in {
			n[i] = r - rune(rng.Real(1)*step)
		}

		fmt.Println(string(n))
	}
}
