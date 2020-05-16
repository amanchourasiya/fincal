package main

import (
	"fmt"

	"github.com/amanchourasiya/fincal/pkg/forcaster"
	"github.com/amanchourasiya/fincal/pkg/invest"
)

func main() {
	displayMainOptions()
	var serviceType int
	fmt.Scanf("%d", &serviceType)

	switch serviceType {

	case 1:
		displayTaxOptions()
		var grossIncome, investment, taxold, taxnew float64
		fmt.Printf("Enter your Gross income")
		fmt.Scanf("%f", &grossIncome)
		fmt.Printf("Enter your investment amount")
		fmt.Scanf("%f", &investment)
		taxcalc := forcaster.TaxDetails{}
		taxcalc.GrossIncome = grossIncome
		taxcalc.Investment = investment
		taxold, taxnew = forcaster.CalculateTax(taxcalc)
		fmt.Printf("Your tax under old regime is %.2f and under new regime is %.2f", taxold, taxnew)
	case 2:
		displayInvestmentOptions()
		var investmentType int
		var investedAmount, time, rate, futureValue float64
		fmt.Scanf("%d", &investmentType)
		investcalc := invest.InvestmentDetails{}
		switch investmentType {
		case 1:
			fmt.Println("Enter invested amount")
			fmt.Scanf("%f", &investedAmount)
			fmt.Println("Enter expected rate of return")
			fmt.Scanf("%f", &rate)
			fmt.Println("Enter no. years of investment")
			fmt.Scanf("%f", &time)
			investcalc.Sum = investedAmount
			investcalc.Time = time
			investcalc.Rate = rate
			investcalc.IsSip = false
			futureValue = invest.CalculateCorpus(investcalc)
			fmt.Printf("Your invested amount %.2f will become %.2f", investedAmount, futureValue)

		case 2:
			fmt.Println("Enter SIP amount")
			fmt.Scanf("%f", &investedAmount)
			fmt.Println("Enter expected rate of return")
			fmt.Scanf("%f", &rate)
			fmt.Println("Enter no. years of investment")
			fmt.Scanf("%f", &time)
			investcalc.Sum = investedAmount
			investcalc.Time = time
			investcalc.Rate = rate
			investcalc.IsSip = true
			futureValue = invest.CalculateCorpus(investcalc)
			investedAmount = investedAmount * float64(12) * time
			fmt.Printf("Your invested amount %.2f will become %.2f", investedAmount, futureValue)
		}
	}
}

func displayMainOptions() {
	fmt.Println("Welcome to financial calculator.")
	fmt.Println("We have following services-\n1.Income tax calculator\n2.Investment corpus calculator")
}

func displayTaxOptions() {
	fmt.Println("Please enter your choice for tax calculation")
	fmt.Println("1.Tax calcultion with old regime.\n2.Tax calculation with new regime.")
}

func displayInvestmentOptions() {
	fmt.Println("Please enter your choice for investment calculation")
	fmt.Println("1.Lumpsum investment.\n2.SIP investment")
}
