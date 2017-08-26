package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type calcRequest struct {
	Operand1 int    `json:"operand1"`
	Operand2 int    `json:"operand2"`
	Operator string `json:"operator"`
}

type calcResult struct {
	Result int `json:"result"`
}

func calcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST!", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var req calcRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request: %s", err), http.StatusBadRequest)
		return
	}

	op, err := getOperation(req.Operator)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting operation: %s", err), http.StatusBadGateway)
		return
	}

	result, err := calculate(req.Operand1, req.Operand2, op)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error during calculation: %s", err), http.StatusBadGateway)
		return
	}

	calcResult := calcResult{
		Result: result,
	}

	// Write HTTP header
	w.WriteHeader(http.StatusOK)

	// Write content (JSON)
	if err := json.NewEncoder(w).Encode(calcResult); err != nil {
		http.Error(w, fmt.Sprintf("Error writing result: %s", err), http.StatusInternalServerError)
	}
}

func getOperation(op string) (operation, error) {
	switch op {
	case "+":
		return opPlus, nil
	case "-":
		return opMinus, nil
	case "*":
		return opMultiply, nil
	case "/":
		return opDivide, nil
	default:
		return opNone, fmt.Errorf("invalid operator: %s", op)
	}
}
