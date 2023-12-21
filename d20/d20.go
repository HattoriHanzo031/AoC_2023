package main

import (
	"fmt"
	"strings"
	"time"
	"utils"
)

type sender func(pulse bool, from string) (bool, []string)

func flipFlopFn(outputs []string) sender {
	var state bool
	return func(pulse bool, from string) (bool, []string) {
		if !pulse {
			state = !state
			return state, outputs
		}
		return state, nil
	}
}

func conjunctionFn(inputs, outputs []string) sender {
	ins := make(map[string]bool, len(inputs))
	for _, in := range inputs {
		ins[in] = false
	}
	return func(pulse bool, from string) (bool, []string) {
		ins[from] = pulse
		for _, v := range ins {
			if !v {
				return true, outputs
			}
		}
		return false, outputs
	}
}

func broadcasterFn(outputs []string) sender {
	return func(pulse bool, _ string) (bool, []string) {
		return pulse, outputs
	}
}

type scannedModule struct {
	t       byte
	inputs  []string
	outputs []string
}

func main() {
	defer utils.Profile(time.Now())
	scanner, cleanup := utils.FileScaner("d20/input.txt")
	defer cleanup()

	// %bm -> pm, vf
	// &xn -> kc, jb, cb, tg, ks, tx
	scanned := make(map[string]scannedModule)
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

	printGraph(scanned)

	modules := make(map[string]sender, len(scanned))
	for k, v := range scanned {
		switch v.t {
		case 'b':
			modules[k] = broadcasterFn(v.outputs)
		case '%':
			modules[k] = flipFlopFn(v.outputs)
		case '&':
			modules[k] = conjunctionFn(v.inputs, v.outputs)
		}
	}

	type out struct {
		from    string
		pulse   bool
		outputs []string
	}

	counts := make(map[bool]int)

	rxInputs := make(map[string]bool)
	for _, s := range scanned["hf"].inputs { // "hf" is the only input to "rx"
		rxInputs[s] = true
	}
	loops := make(map[string]int)
	for i := 0; len(loops) < len(rxInputs); i++ {
		counts[false]++
		var initial out
		initial.pulse, initial.outputs = modules["broadcaster"](false, "")
		for outputs := []out{initial}; len(outputs) > 0; {
			newOutputs := []out{}
			for _, oo := range outputs {
				for _, o := range oo.outputs {
					counts[oo.pulse]++
					newOut := out{}
					module, ok := modules[o]
					if !ok {
						continue
					}

					newOut.pulse, newOut.outputs = module(oo.pulse, oo.from)
					newOut.from = o
					newOutputs = append(newOutputs, newOut)

					if newOut.pulse {
						if _, ok := loops[newOut.from]; rxInputs[newOut.from] && !ok {
							loops[newOut.from] = i + 1
						}
					}
				}
			}
			outputs = newOutputs
		}
		if i+1 == 1000 {
			fmt.Println("P1:", counts[true]*counts[false])
		}
	}
	lcmInput := []int{}
	for _, l := range loops {
		lcmInput = append(lcmInput, l)
	}
	fmt.Println("P2:", utils.LCM(lcmInput[0], lcmInput[1], lcmInput[2:]...))
}

func printGraph(scanned map[string]scannedModule) {
	fmt.Print("digraph G {")
	for k, v := range scanned {
		fmt.Print(k, "-> {", strings.Join(v.outputs, ","), "};")
		if k == "rx" {
			fmt.Print(k, "[shape=Msquare];")
			continue
		}
		switch v.t {
		case 'b':
			fmt.Print(k, "[shape=Mdiamond];")
		case '%':
			fmt.Print(k, "[shape=ellipse];")
		case '&':
			fmt.Print(k, "[shape=box];")
		}
	}
	fmt.Println("}")
}
