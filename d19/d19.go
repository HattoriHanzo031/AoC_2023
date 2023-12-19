package main

import (
	"fmt"
	"maps"
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
	workflowsP2 := make(map[string]workflowP2)
	//px{a<2006:qkq,m>2090:A,rfg}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		wfName, wfDefinition, _ := strings.Cut(strings.TrimSuffix(line, "}"), "{")
		rules := []rule{}
		rulesP2 := []ruleP2{}
		for _, rr := range strings.Split(wfDefinition, ",") {
			r, next, isCompare := strings.Cut(rr, ":")
			if !isCompare {
				rules = append(rules, unconditionalFn(r))
				rulesP2 = append(rulesP2, unconditionalP2Fn(r))
				continue
			}
			rules = append(rules, ruleFn(r[0], utils.Must(strconv.Atoi(r[2:])), fns[r[1]], next))
			rulesP2 = append(rulesP2, ruleP2Fn(r[0], utils.Must(strconv.Atoi(r[2:])), fnsP2[r[1]], next))
		}
		workflows[wfName] = workflowFn(rules, workflows)
		workflowsP2[wfName] = workflowP2Fn(rulesP2, workflowsP2)
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

	totalP2 := workflowsP2["in"](partP2{'x': rng{1, 4000}, 'm': rng{1, 4000}, 'a': rng{1, 4000}, 's': rng{1, 4000}})
	fmt.Println(totalP2)
}

type rng struct{ from, to int }
type partP2 map[byte]rng

type ruleP2 func(p partP2) (string, *partP2, *partP2)

type splitRangeP2 func(r rng, val int) (*rng, *rng)

var fnsP2 = map[byte]splitRangeP2{
	'>': func(r rng, val int) (*rng, *rng) {
		switch {
		case val >= r.to:
			return nil, &r
		case val < r.from:
			return &r, nil
		default:
			return &rng{val + 1, r.to}, &rng{r.from, val}
		}
	},
	'<': func(r rng, val int) (*rng, *rng) {
		switch {
		case val > r.to:
			return &r, nil
		case val <= r.from:
			return nil, &r
		default:
			return &rng{r.from, val - 1}, &rng{val, r.to}
		}
	},
}

func ruleP2Fn(category byte, val int, fn splitRangeP2, next string) ruleP2 {
	return func(p partP2) (string, *partP2, *partP2) {
		pass, fail := fn(p[category], val)
		var newPassP, newFailP *partP2
		if pass != nil {
			newPass := maps.Clone(p)
			newPass[category] = *pass
			newPassP = &newPass
		}
		if fail != nil {
			newFail := maps.Clone(p)
			newFail[category] = *fail
			newFailP = &newFail
		}
		return next, newPassP, newFailP
	}
}

func unconditionalP2Fn(next string) ruleP2 {
	return func(p partP2) (string, *partP2, *partP2) {
		return next, &p, nil
	}
}

type workflowP2 func(p partP2) int

func workflowP2Fn(rules []ruleP2, workflows map[string]workflowP2) workflowP2 {
	return func(p partP2) int {
		total := 0
		for _, rule := range rules {
			next, pass, fail := rule(p)
			switch next {
			case "A":
				if pass != nil {
					pss := *pass
					total += ((pss['x'].to - pss['x'].from + 1) * (pss['m'].to - pss['m'].from + 1) * (pss['a'].to - pss['a'].from + 1) * (pss['s'].to - pss['s'].from + 1))
				}
			case "R":
			default:
				if pass != nil {
					total += workflows[next](*pass)
				}
			}
			if fail == nil {
				break
			}
			p = *fail
		}
		return total
	}
}
