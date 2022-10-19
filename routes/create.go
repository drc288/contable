package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/drc288/contable/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var expense models.Gasto

	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), 400)
		return
	}

	if len(strconv.Itoa(expense.Year)) == 0 {
		http.Error(w, "Error validating the year", 400)
		return
	}

	if len(expense.Month) == 0 {
		http.Error(w, "Error validating the month", 400)
		return
	}

	if len(expense.ExpenseName) == 0 {
		http.Error(w, "Error validating the expense name", 400)
		return
	}

	if len(strconv.Itoa(expense.Expense)) == 0 {
		http.Error(w, "Error validating the expense", 400)
		return
	}

	obj := fmt.Sprintf("%v", expense)

	fmt.Fprintln(w, obj)
}
