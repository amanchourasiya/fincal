package utils

import "math"

// InvestmentDetails holds details for calculating interest
type InvestmentDetails struct {
	Principal int
	Rate      float64
	Time      float64
}

// CompoundInterest calculates total amount after compound interest by formula C.I. = P * [1 + (R/100)]^T
func CompoundInterest(investment InvestmentDetails) float64 {
	rate := investment.Rate
	principal := investment.Principal
	time := investment.Time
	res := (1 + (rate / 100.00))
	res = math.Pow(res, time)
	res = res * float64(principal)
	return res
}

// SimpleInterest calculates amount after adding simple interest by formula S.I. = (P * R * T) / 100
func SimpleInterest(investment InvestmentDetails) float64 {
	rate := investment.Rate
	principal := investment.Principal
	time := investment.Time
	res := float64(principal) * rate * time * 0.01
	return res
}

// SipAmount calculates SIP amount accumulated over a period
func SipAmount(investment InvestmentDetails) float64 {
	rate := investment.Rate / 100.00 / 12.00
	time := investment.Time * 12
	principal := investment.Principal
	res := (math.Pow((1+rate), time) - 1) / rate
	res = res * float64(principal) * (1 + rate)
	return res

}
