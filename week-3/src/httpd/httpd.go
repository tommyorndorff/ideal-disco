package main

import (
	"fmt"
	"net/http"
	"strconv"

	"math/big"

	"github.com/gorilla/mux"
)

func arccot(x int64, unity *big.Int) *big.Int {
	bigx := big.NewInt(x)
	xsquared := big.NewInt(x * x)
	sum := big.NewInt(0)
	sum.Div(unity, bigx)
	xpower := big.NewInt(0)
	xpower.Set(sum)
	n := int64(3)
	zero := big.NewInt(0)
	sign := false

	term := big.NewInt(0)
	for {
		xpower.Div(xpower, xsquared)
		term.Div(xpower, big.NewInt(n))
		if term.Cmp(zero) == 0 {
			break
		}
		if sign {
			sum.Add(sum, term)
		} else {
			sum.Sub(sum, term)
		}
		sign = !sign
		n += 2
	}
	return sum
}

func Pi(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	length, ok := vars["length"]
	if !ok {
		length = "1"
	}
	i, err := strconv.ParseInt(length, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	digits := big.NewInt(i + 10)
	unity := big.NewInt(0)
	unity.Exp(big.NewInt(10), digits, nil)
	pi := big.NewInt(0)
	four := big.NewInt(4)
	pi.Mul(four, pi.Sub(pi.Mul(four, arccot(5, unity)), arccot(239, unity)))
	//val := big.Mul(4, big.Sub(big.Mul(4, arccot(5, unity)), arccot(239, unity)))
	fmt.Println("Hello, Pi:  ", pi)
	fmt.Fprintf(w, "pi is: %d\n", pi)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/pi/{length}", Pi)

	http.ListenAndServe(":8080", r)
}
