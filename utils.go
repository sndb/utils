package utils

import (
	"io"
	"os"
	"strconv"
)

func Input() string {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func Atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func GCD(m, n int) int {
	for n != 0 {
		m, n = n, m%n
	}
	return m
}

func LCM(m, n int) int {
	return m * n / GCD(m, n)
}

func Mod(p, q int) int {
	r := p % q
	if r < 0 {
		return q + r
	}
	return r
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type Pos struct {
	r, c int
}

func (p Pos) Distance(q Pos) int {
	return Abs(p.r-q.r) + Abs(p.c-q.c)
}

func (p Pos) Inside(a, b Pos) bool {
	return p.r >= a.r && p.r < b.r && p.c >= a.c && p.c < b.c
}

func (p Pos) Add(q Pos) Pos {
	return Pos{p.r + q.r, p.c + q.c}
}

func (p Pos) Adjacent(q Pos) bool {
	return Abs(p.r-q.r) <= 1 && Abs(p.c-q.c) <= 1
}
