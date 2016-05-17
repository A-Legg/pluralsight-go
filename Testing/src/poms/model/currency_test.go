package model

import (
	"encoding/json"
	"testing"
)

func TestGetsTheCorrectCurrencies(t *testing.T) {
	//arange

	//act
	result := getCurrencies()

	//assert
	resultJSON, _ := json.Marshal(result)
	expected, _ := jason.Marshal([]*Currency{
		{ID: 1, Name: "USD"},
		{ID: 2, Name: "EUR"},
	})

	if string(resultJSON) != string(expected) {
		t.Error("Did not get expected currencies")
	}
}
