package main

import (
	"strings"
	"strconv"
	"fmt"
	"math"
)

type Program struct {
	tokens []string
	next   int
}

type Token struct {
	str    string
	tokens []Token
}

func (t Token) String() string {
	if t.tokens == nil {
		return t.str
	} else {
		return fmt.Sprint(t.tokens)
	}
}

type Env struct {
	dict map[string] interface {}
	outer *Env
}

func NewEnv(outer *Env) *Env {
	return &Env{dict: make(map[string]interface {}), outer: outer }
}

func (e *Env) find(key string) *Env {
	if _, present := e.dict[key]; present {
		return e
	}  else if (e.outer != nil) {
		return e.outer.find(key)
	} else {
		panic(fmt.Sprintf("%v not found", key))
	}
}

func add_globals(env *Env) *Env {
	g := map[string]interface {}{
		"+": func(params... interface {}) interface {} {
			sum := 0.0
			for _, n := range params {
				switch n := n.(type){
				case int64:
					sum += float64(n)
				case float64:
					sum += n
				case string:
					v, err := strconv.ParseFloat(n, 64)
					if err == nil {
						sum += v
					}
				}
			}
			return sum
		},
		">": func(params... interface {}) interface {} {
			before := math.MaxFloat64
			for _, n := range params {
				switch n := n.(type){
				case int64:
					if before > float64(n) {
						before = float64(n)
					} else {
						return false
					}
				case float64:
					if before > n {
						before = n
					} else {
						return false
					}
				case string:
					v, err := strconv.ParseFloat(n, 64)
					if err == nil {
						if before > v {
							before = v
						} else {
							return false
						}
					}
				}
			}
			return true
		},
		"-": func(params... interface {}) interface {} {
			sum := params[0].(float64)
			for _, n := range params[1:] {
				switch n := n.(type){
				case int64:
					sum -= float64(n)
				case float64:
					sum -= n
				case string:
					v, err := strconv.ParseFloat(n, 64)
					if err == nil {
						sum -= v
					}
				}
			}
			return sum
		},
		"*": func(params... interface {}) interface {} {
			result := 1.0
			for _, n := range params {
				switch n := n.(type){
				case int64:
					result *= float64(n)
				case float64:
					result *= n
				case string:
					v, err := strconv.ParseFloat(n, 64)
					if err == nil {
						result *= v
					}
				}
			}
			return result
		},
	}

	for k, v := range g {
		env.dict[k] = v
	}
	return env
}

func tokenize(p string) *Program {
	r := strings.NewReplacer("(", " ( ", ")", " ) ")
	return &Program{tokens: strings.Fields(r.Replace(p)), next:0 }
}

func read(p string) Token {
	return readFrom(tokenize(p))
}

func readFrom(t *Program) Token {
	if t.next >= len(t.tokens) {
		panic("unexpected EOF while reading")
	}
	token := t.tokens[t.next]
	t.next += 1
	if "(" == token {
		L := make([]Token, 0)
		for ; t.tokens[t.next] != ")"; {
			token = t.tokens[t.next]
			if token == "(" {
				L = append(L, readFrom(t))
			} else {
				L = append(L, Token{str: token})
				t.next += 1
			}
		}
		t.next += 1 // remove )
		return Token{tokens: L}
	} else if ")" == token {
		panic("unexpected )")
	}
	return Token{}
}

func eval(x Token, env *Env) interface {} {
	if x.str != "" {
		if v, e := strconv.ParseFloat(x.str, 64); e == nil {
			return v
		}
		return env.find(x.str).dict[x.str]
	}

	tokens := x.tokens

	switch tokens[0].str{
	case "quote":
		return tokens[1:]
	case "if":
		test := eval(tokens[1], env)
		if t, ok := test.(bool); ok && !t {
			return eval(tokens[3], env)
		} else {
			return eval(tokens[2], env)
		}
	case "set!":
		v := tokens[1].str
		env.find(v).dict[v] = eval(tokens[2], env)
	case "define":
		v := tokens[1].str
		env.dict[v] = eval(tokens[2], env)
	case "lambda":
		v := tokens[1].tokens
		expr := tokens[2]
		return func(params...interface {}) interface {} {
			e := NewEnv(env)
			for i, name := range v {
				e.dict[name.str] = params[i]
			}
			return eval(expr, e)
		}
	case "begin":
		var val interface {}
		for _, e := range tokens[1:] {
			val = eval(e, env)
		}
		return val
	default:
		exps := make([]interface {}, len(tokens))
		for i, e := range (tokens) {
			exps[i] = eval(e, env)
		}

		if proc, ok := exps[0].(func (...interface {}) interface {}); ok {
			return proc(exps[1:]...)
		} else {
			panic(fmt.Sprintf("%s Not a function", exps[0]))
		}
	}
	return nil
}

//func

func main() {
	exprs := []string{
		"(begin (define f (lambda (n) (* n n))) (f 10))",
		"(- 1 1)",
		`
(begin
  (define f (lambda (n)
	(if (> n 0)
	  (* n (f (- n 1)))
	1)))
  (f 6))`,
	}

	for _, p := range exprs {
		globals := add_globals(NewEnv(nil))
		fmt.Println(p, "=>", eval(read(p), globals))

	}
}
