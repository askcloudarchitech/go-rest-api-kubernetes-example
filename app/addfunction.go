package main

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// myDescarptiveAddFunction is the handler for the /add endpoint
func myDescriptiveAddFunction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	result, err := addStringsAsNumbers(ps.ByName("num1"), ps.ByName("num2"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte(strconv.Itoa(result)))
}

//addStringsAsNumbers will take two values that should be numbers in string format
// and return a sum of the two as an integer.
// this is necessary because the values provided in the URL will always be strings
// initially and will need to be converted
func addStringsAsNumbers(num1 string, num2 string) (int, error) {
	num1AsInt, err := strconv.Atoi(num1)
	if err != nil {
		return 0, err
	}
	num2AsInt, err := strconv.Atoi(num2)
	if err != nil {
		return 0, err
	}
	return num1AsInt + num2AsInt, nil
}
