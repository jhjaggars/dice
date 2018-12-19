package dice

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func sum(nums []int) (total int) {
	for _, n := range nums {
		total += n
	}
	return
}

func single(sides int) int {
	return 1 + rand.Intn(sides)
}

type Output struct {
	Rolls []int `json:"rolls"`
	Total int   `json:"total"`
}

type Dice struct {
	Number int
	Sides  int
}

func (d *Dice) Roll() (out Output) {
	rolls := make([]int, d.Number)
	for i := range rolls {
		rolls[i] = single(d.Sides)
	}
	return Output{rolls, sum(rolls)}
}

func convertString(s string, defaultValue int) (r int) {
	if s == "" {
		r = defaultValue
	} else {
		parsed, err := strconv.Atoi(s)
		if err != nil {
			r = defaultValue
		} else {
			r = parsed
		}
	}
	return
}

func ParseDie(s string) (d Dice, err error) {
	num := 1
	sides := 6
	hasd := false
	for _, ch := range s {
		if ch == 'd' {
			hasd = true
		}
	}
	if !hasd {
		err = errors.New(fmt.Sprintf("die '%s' has no 'd'", s))
		return
	}
	parts := strings.Split(s, "d")
	num = convertString(parts[0], 1)
	sides = convertString(parts[1], 6)
	d = Dice{num, sides}
	return
}
