package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"utils"
)

// {x=787,m=2655,a=1222,s=2876}
type part map[byte]int

// rfg{s<537:gd,x>2440:R,A}
type rule func(p part) string

type condition func(val1, val2 int) bool

var fns = map[byte]condition{
	'>': func(val1, val2 int) bool { return val1 > val2 },
	'<': func(val1, val2 int) bool { return val1 < val2 },
	'T': func(val1, val2 int) bool { return true },
}

func ruleFn(category byte, val int, fn condition, next string) rule {
	return func(p part) string {
		if fn(p[category], val) {
			return next
		}
		return ""
	}
}

func unconditionalFn(next string) rule {
	return func(p part) string {
		return next
	}
}

type workflow func(p part) bool

func workflowFn(rules []rule, workflows map[string]workflow) workflow {
	return func(p part) bool {
		for _, rule := range rules {
			switch next := rule(p); next {
			case "":
			case "A":
				return true
			case "R":
				return false
			default:
				return workflows[next](p)
			}
		}
		fmt.Println("SHOULD NOT BE HERE!")
		return false
	}
}

func main() {
	defer utils.Profile(time.Now())
	scanner, cleanup := utils.FileScaner("d19/input.txt")
	defer cleanup()

	workflows := make(map[string]workflow)
	//px{a<2006:qkq,m>2090:A,rfg}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		wfName, wfDefinition, _ := strings.Cut(strings.TrimSuffix(line, "}"), "{")
		rules := []rule{}
		for _, rr := range strings.Split(wfDefinition, ",") {
			r, next, isCompare := strings.Cut(rr, ":")
			if !isCompare {
				rules = append(rules, unconditionalFn(r))
				continue
			}
			rules = append(rules, ruleFn(r[0], utils.Must(strconv.Atoi(r[2:])), fns[r[1]], next))
		}
		workflows[wfName] = workflowFn(rules, workflows)
	}

	// {x=787,m=2655,a=1222,s=2876}
	parts := []part{}
	for scanner.Scan() {
		var x, m, a, s int
		fmt.Sscanf(scanner.Text(), "{x=%d,m=%d,a=%d,s=%d}", &x, &m, &a, &s)
		parts = append(parts, part{'x': x, 'm': m, 'a': a, 's': s})
	}

	total := 0
	for _, part := range parts {
		if workflows["in"](part) {
			total += part['x'] + part['m'] + part['a'] + part['s']
		}
	}
	fmt.Println("TOTAL:", total)
}
