package main

import (
	"fmt"
	"strings"
	"time"
	"utils"
)

type sender func(pulse bool, from string) (bool, string, []string)

func flipFlopFn(me string, outputs []string) sender {
	var state bool
	return func(pulse bool, from string) (bool, string, []string) {
		if !pulse {
			state = !state
			return state, me, outputs
		}
		return state, me, nil
	}
}

func conjunctionFn(me string, inputs, outputs []string) sender {
	ins := make(map[string]bool, len(inputs))
	for _, in := range inputs {
		ins[in] = false
	}
	return func(pulse bool, from string) (bool, string, []string) {
		ins[from] = pulse
		for _, v := range ins {
			if !v {
				return true, me, outputs
			}
		}
		return false, me, outputs
	}
}

func broadcasterFn(me string, outputs []string) sender {
	return func(pulse bool, _ string) (bool, string, []string) {
		return pulse, me, outputs
	}
}

func main() {
	defer utils.Profile(time.Now())
	scanner, cleanup := utils.FileScaner("d20/input.txt")
	defer cleanup()

	type module struct {
		t       byte
		inputs  []string
		outputs []string
	}

	// %bm -> pm, vf
	// &xn -> kc, jb, cb, tg, ks, tx
	scanned := make(map[string]module)
	for scanner.Scan() {
		name, outputs, _ := strings.Cut(scanner.Text(), " -> ")
		t := name[0]
		if name != "broadcaster" {
			name = name[1:]
		}
		this := scanned[name]
		this.t = t
		this.outputs = strings.Split(outputs, ", ")
		for _, o := range this.outputs {
			other := scanned[o]
			other.inputs = append(other.inputs, name)
			scanned[o] = other
		}
		scanned[name] = this
	}

	modules := make(map[string]sender, len(scanned))
	for k, v := range scanned {
		fmt.Println("KV", k, v)
		switch v.t {
		case 'b':
			modules[k] = broadcasterFn(k, v.outputs)
		case '%':
			modules[k] = flipFlopFn(k, v.outputs)
		case '&':
			modules[k] = conjunctionFn(k, v.inputs, v.outputs)
		}
	}

	type out struct {
		from    string
		pulse   bool
		outputs []string
	}

	counts := make(map[bool]int)
	for i := 0; i < 1000; i++ {
		counts[false]++
		countRx := 0
		var initial out
		initial.pulse, initial.from, initial.outputs = modules["broadcaster"](false, "")
		//fmt.Println("INITIAL", initial)
		//fmt.Println("PUSH")
		for outputs := []out{initial}; len(outputs) > 0; {
			newOutputs := []out{}
			//fmt.Println("OUTPUTS", outputs)
			for _, oo := range outputs {
				for _, o := range oo.outputs {
					newOut := out{}
					counts[oo.pulse]++
					module, ok := modules[o]
					if !ok {
						if o == "rx" {
							countRx++
						}
						continue
					}
					newOut.pulse, newOut.from, newOut.outputs = module(oo.pulse, oo.from)
					newOutputs = append(newOutputs, newOut)
				}
			}
			//fmt.Println("")
			outputs = newOutputs
		}
		if countRx == 1 {
			fmt.Println("I:", i)
			return
		}
	}
	fmt.Println("COUNTS", counts)
}

// PUSH
// button -low-> broadcaster
// broadcaster -low-> a

// a -high-> inv
// a -high-> con

// inv -low-> b
// con -high-> output

// b -high-> con
// con -low-> output

// PUSH
// button -low-> broadcaster
// broadcaster -low-> a
// a -low-> inv
// a -low-> con
// inv -high-> b
// con -high-> output

// PUSH
// button -low-> broadcaster
// broadcaster -low-> a
// a -high-> inv
// a -high-> con
// inv -low-> b
// con -low-> output
// b -low-> con
// con -high-> output

// PUSH
// button -low-> broadcaster
// broadcaster -low-> a
// a -low-> inv
// a -low-> con
// inv -high-> b
// con -high-> output
