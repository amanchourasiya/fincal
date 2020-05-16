package forcaster

// TaxDetails holds details to be used for calculating tax
type TaxDetails struct {
	GrossIncome float64
	Investment  float64
}

// CalculateTax calculates tax based on gorss income and investment made.
func CalculateTax(taxdetails TaxDetails) (float64, float64) {
	var taxold, taxnew float64
	taxold = calculateOldRegime(taxdetails)
	taxnew = calculateNewRegime(taxdetails)
	return taxold, taxnew
}

func calculateOldRegime(taxdetails TaxDetails) float64 {
	const taxunder5 = 250000.00 * 0.05
	const taxunder10 = 500000.00 * 0.20

	var taxToBePaid float64

	switch taxableIncome := taxdetails.GrossIncome - taxdetails.Investment; {
	case taxableIncome < 250000:
		taxToBePaid = 0.00

	case taxableIncome < 500000:
		taxToBePaid = (taxableIncome - 250000) * 0.05

	case taxableIncome < 1000000:
		taxToBePaid = taxunder5 + (taxableIncome-500000)*0.20

	case taxableIncome >= 1000000:
		taxToBePaid = taxunder5 + taxunder10 + (taxableIncome-1000000)*0.30

	default:
		taxToBePaid = -1.00
	}
	return taxToBePaid + calculateCess(taxToBePaid)
}

func calculateNewRegime(taxdetails TaxDetails) float64 {
	const taxunder5 = 250000.00 * 0.05
	const taxunder75 = 250000.00 * 0.10
	const taxunder10 = 250000.00 * 0.15
	const taxunder125 = 250000.00 * 0.20
	const taxunder15 = 250000.00 * 0.25

	var taxToBePaid float64

	switch taxableIncome := taxdetails.GrossIncome; {
	case taxableIncome < 500000:
		taxToBePaid = 0.00

	case taxableIncome < 750000:
		taxToBePaid = taxunder5 + (taxableIncome-500000)*0.10

	case taxableIncome < 1000000:
		taxToBePaid = taxunder5 + taxunder75 + (taxableIncome-750000)*0.15

	case taxableIncome < 1250000:
		taxToBePaid = taxunder5 + taxunder75 + taxunder10 + (taxableIncome-1000000)*0.20

	case taxableIncome < 1500000:
		taxToBePaid = taxunder5 + taxunder75 + taxunder10 + taxunder125 + (taxableIncome-1250000)*0.25

	case taxableIncome > 1500000:
		taxToBePaid = taxunder5 + taxunder75 + taxunder10 + taxunder125 + taxunder15 + (taxableIncome-1500000)*0.30

	default:
		taxToBePaid = -1.00
	}
	return taxToBePaid + calculateCess(taxToBePaid)
}

func calculateCess(tax float64) float64 {
	cess := tax * 0.04
	return cess
}
