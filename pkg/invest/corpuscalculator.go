package invest

import (
	"github.com/amanchourasiya/fincal/pkg/utils"
)

// InvestmentDetails will hold details about the investment
type InvestmentDetails struct {
	Sum   float64
	IsSip bool
	Time  float64
	Rate  float64
}

// Lumpsum calculates corpus based on lumpsum invested over a period of time
func Lumpsum(investment InvestmentDetails) float64 {
	ci := utils.InvestmentDetails{}
	ci.Principal = int(investment.Sum)
	ci.Rate = investment.Rate
	ci.Time = investment.Time
	return utils.CompoundInterest(ci)
}

// Sip caculates corpus based on SIP amount invested over a period of time
func Sip(investment InvestmentDetails) float64 {
	ci := utils.InvestmentDetails{}
	ci.Principal = int(investment.Sum)
	ci.Rate = investment.Rate
	ci.Time = investment.Time
	return utils.SipAmount(ci)
}

// CalculateCorpus calculates total corpus
func CalculateCorpus(investment InvestmentDetails) float64 {
	var res float64
	if investment.IsSip == true {
		res = Sip(investment)
	} else {
		res = Lumpsum(investment)
	}
	return res
}
