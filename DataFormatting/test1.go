package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type SalaryData struct {
	Löneart string
	From    string
	Antal   string
	Belopp  string
}

func sumBelopp(data []SalaryData) float64 {
	sum := 0.0
	for _, item := range data {
		normalizedBelopp := strings.ReplaceAll(item.Belopp, ",", ".")
		normalizedBelopp = strings.ReplaceAll(normalizedBelopp, " ", "")
		belopp, _ := strconv.ParseFloat(normalizedBelopp, 64)
		sum += belopp
	}
	return sum
}

func main() {
	start := time.Now()
	tableData := []SalaryData{
		{
			"4200  Sem.ers övriga",
			"24-03-27",
			"4,00",
			"106,56",
		},
		{
		Löneart: "0204  Timlön",
		From: "24-03-01",
		Antal: "4,00",
		Belopp: "888,00",
	},
	{
		Löneart: "4200  Sem.ers övriga",
		From: "24-03-01",
		Antal: "4,00",
		Belopp: "106,56",
	},
	{
		Löneart: "0204  Timlön",
		From: "24-03-04",
		Antal: "3,00",
		Belopp: "666,00",
	},
	{
		Löneart: "4200  Sem.ers övriga",
		From: "24-03-04",
		Antal: "3,00",
		Belopp: "79,92",
	},
	{
		Löneart: "0204  Timlön",
		From: "24-03-05",
		Antal: "5,00",
		Belopp: "1 110,00",
	},
	{
		Löneart: "4200  Sem.ers övriga",
		From: "24-03-05",
		Antal: "5,00",
		Belopp: "133,20",
	},
	{
		Löneart: "0204  Timlön",
		From: "24-03-07",
		Antal: "4,00",
		Belopp: "888,00",
	},
	{
		Löneart: "4200  Sem.ers övriga",
		From: "24-03-07",
		Antal: "4,00",
		Belopp: "106,56",
	},
	{
		Löneart: "0204  Timlön",
		From: "24-03-06",
		Antal: "5,00",
		Belopp: "1 110,00",
	},
	{
		Löneart: "4200  Sem.ers övriga",
		From: "24-03-06",
		Antal: "5,00",
		Belopp: "133,20",
	},
	{
		Löneart: "0204  Timlön",
		From: "24-03-08",
		Antal: "4,00",
		Belopp: "888,00",
	},
	{
		Löneart: "4200  Sem.ers övriga",
		From: "24-03-08",
		Antal: "4,00",
		Belopp: "106,56",
	},
	{
		Löneart: "0204  Timlön",
		From: "24-03-12",
		Antal: "4,00",
		Belopp: "888,00",
	},
	{
		Löneart: "4200  Sem.ers övriga",
		From: "24-03-12",
		Antal: "4,00",
		Belopp: "106,56",
	},
	{
		Löneart: "0204  Timlön",
		From: "24-03-13",
		Antal: "4,00",
		Belopp: "888,00",
	},
	{
		Löneart: "4200  Sem.ers övriga",
		From: "24-03-13",
		Antal: "4,00",
		Belopp: "106,56",
	},
	{
		Löneart: "0204  Timlön",
		From: "24-03-14",
		Antal: "4,00",
		Belopp: "888,00",
	},
	{
		Löneart: "4200  Sem.ers övriga",
		From: "24-03-14",
		Antal: "4,00",
		Belopp: "106,56",
	},
	{
		Löneart: "0204  Timlön",
		From: "24-03-19",
		Antal: "5,00",
		Belopp: "1 110,00",
	},
	{
		Löneart: "4200  Sem.ers övriga",
		From: "24-03-19",
		Antal: "5,00",
		Belopp: "133,20",
	},
	{
		Löneart: "0204  Timlön",
		From: "24-03-21",
		Antal: "4,00",
		Belopp: "888,00",
	},
	{
		Löneart: "4200  Sem.ers övriga",
		From: "24-03-21",
		Antal: "4,00",
		Belopp: "106,56",
	},
	{
		Löneart: "0204  Timlön",
		From: "24-03-22",
		Antal: "4,00",
		Belopp: "888,00",
	},
	{
		Löneart: "4200  Sem.ers övriga",
		From: "24-03-22",
		Antal: "4,00",
		Belopp: "106,56",
	},
	{
		Löneart: "0204  Timlön",
		From: "24-03-25",
		Antal: "5,00",
		Belopp: "1 110,00",
	},
	{
		Löneart: "4200  Sem.ers övriga",
		From: "24-03-25",
		Antal: "5,00",
		Belopp: "133,20",
	},
	{
		Löneart: "0204  Timlön",
		From: "24-03-26",
		Antal: "4,00",
		Belopp: "888,00",
	},
	{
		Löneart: "4200  Sem.ers övriga",
		From: "24-03-26",
		Antal: "4,00",
		Belopp: "106,56",
	},
	{
		Löneart: "0204  Timlön",
		From: "24-03-27",
		Antal: "4,00",
		Belopp: "888,00",
	},
	}

	var timlonData []SalaryData
	var semersOvrigaData []SalaryData

	for _, item := range tableData {
		if item.Löneart == "0204  Timlön" {
			timlonData = append(timlonData, item)
		} else if item.Löneart == "4200  Sem.ers övriga" {
			semersOvrigaData = append(semersOvrigaData, item)
		}
	}

	totalBeloppTimlon := sumBelopp(timlonData)
	totalBeloppSemersOvriga := sumBelopp(semersOvrigaData)

	beloppToPay := totalBeloppTimlon + totalBeloppSemersOvriga
	fmt.Printf("Total Belopp for 0204 Timlön: %.2f\n", totalBeloppTimlon)
	fmt.Printf("Total Belopp for 4200 Sem.ers övriga: %.2f\n", totalBeloppSemersOvriga)
	fmt.Printf("Belopp to Pay: %.2f\n", beloppToPay)

	// ACCOUNT FOR THEFT (TAXES)
	tax := 0.7
	beloppToPayAfterTax := beloppToPay * tax
	fmt.Printf("Belopp to Pay (after tax): %.2f\n", beloppToPayAfterTax)
	duration := time.Since(start)
	fmt.Println(duration)
}
