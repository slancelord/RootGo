package main

import (
	"fmt"
	"math/cmplx"
	"strconv"
)

func input() (a, b, c float64, err error) {
	var a1, b1, c1 string

	fmt.Println("ax²+bx+c=0")

	fmt.Print("a = ")
	fmt.Scan(&a1)
	a, err = strconv.ParseFloat(a1, 64)
	if err != nil {
		return
	}

	fmt.Print("b = ")
	fmt.Scan(&b1)
	b, err = strconv.ParseFloat(b1, 64)
	if err != nil {
		return
	}

	fmt.Print("c = ")
	fmt.Scan(&c1)
	c, err = strconv.ParseFloat(c1, 64)
	if err != nil {
		return
	}

	return
}

func rootEq(a, b, c float64) (x [5]float64, err error) {

	if a == 0 && b == 0 {
		if c == 0 {
			err = fmt.Errorf("The equation has infinitely many solutions")
		} else {
			err = fmt.Errorf("%f ≠ 0", c)
		}
	} else {
		d := complex(b*b-4*a*c, 0)
		if a == 0 {
			x[0] = (c * -1) / b
		} else {
			d = cmplx.Sqrt(d)
			x[0] = real(((-1 * complex(b, 0)) + d) / (2 * complex(a, 0)))
			x[1] = imag(((-1 * complex(b, 0)) + d) / (2 * complex(a, 0)))
			x[2] = real(((-1 * complex(b, 0)) - d) / (2 * complex(a, 0)))
			x[3] = imag(((-1 * complex(b, 0)) - d) / (2 * complex(a, 0)))
			x[4] = b*b - 4*a*c
		}
	}

	return
}

func main() {
	var x [5]float64
	exit := false
	a, b, c, err := input()

	for !exit {
		exit = true
		if err != nil {
			fmt.Println("ERROR:", err)
			fmt.Printf("\nTry it again\n\n")
			a, b, c, err = input()
			exit = false
		}
	}

	x, err = rootEq(a, b, c)

	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		if x[4] > 0 {
			fmt.Printf("The equation has two real roots x = %f, %f", x[0], x[2])
		} else if x[4] == 0 {
			fmt.Printf("The equation has one real repeated root x = %f", x[0])
		} else if x[4] < 0 {
			fmt.Printf("The equation has two complex roots x = %g, %g", complex(x[0], x[1]), complex(x[2], x[3]))
		}
	}
}
