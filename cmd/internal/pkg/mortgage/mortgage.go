package mortgage

import (
	"errors"
	"log"
	"math"

	"github.com/martijnkorbee/michael/cmd/internal/pkg/util"
)

// Mortgage reflects a single row in the mortgage tables.
type Mortgage struct {
	Type              string  `db:"type"`
	Mortgage          int     `db:"mortgage"`
	Terms             int     `db:"terms"`
	Interest          float64 `db:"interest"`
	Month             int     `db:"month"`
	Year              int     `db:"Year"`
	Remainder         float64 `db:"remainder"`
	InterestPayment   float64 `db:"interest_pmt"`
	RedemptionPayment float64 `db:"redemption_pmt"`
	TotalPayment      float64 `db:"total_pmt"`
	TaxDiscount       float64 `db:"tax_discount"`
	NettPayment       float64 `db:"nett_pmt"`
	PayedInterest     float64 `db:"payed_interest"`
	PayedRedemption   float64 `db:"payed_redemption"`
	TotalPayed        float64 `db:"total_payed"`
	NettPayed         float64 `db:"nett_payed"`
}

func New(mortgageType string, sum, duration int, interest float64) *Mortgage {
	mg := Mortgage{
		Type:      mortgageType,
		Mortgage:  sum,
		Remainder: float64(sum),
		Terms:     duration * 12,
		Interest:  interest,
		Year:      1,
	}

	switch mortgageType {
	case "linear":
		// set the fixed redemption payment for the linear model
		mg.RedemptionPayment = util.RoundFloat((float64(mg.Mortgage) / float64(mg.Terms)), 2)
	case "annuities":
		// set the fixed total monthly payment for the annuities model
		monthlyInterest := (mg.Interest / 100) / 12
		annuity := (monthlyInterest * float64(mg.Mortgage))

		mg.TotalPayment = util.RoundFloat(annuity/(1-(math.Pow((1+monthlyInterest), -360))), 2)
	default:
		log.Fatalln("unsupported mortgage type")
	}

	return &mg
}

func (m *Mortgage) CalculateNextMonth() error {
	// increment month/year
	m.incrementMonth()

	// monthly payments
	switch m.Type {
	case "linear":
		m.InterestPayment = util.RoundFloat((m.Remainder*(m.Interest/100))/12, 2)
		m.TotalPayment = util.RoundFloat(m.InterestPayment+m.RedemptionPayment, 2)
	case "annuities":
		m.InterestPayment = util.RoundFloat((m.Remainder*(m.Interest/100))/12, 2)
		m.RedemptionPayment = util.RoundFloat(m.TotalPayment-m.InterestPayment, 2)
	default:
		return errors.New("unsupported mortgage type")
	}

	// tax discount and nett payment
	m.TaxDiscount = util.RoundFloat((m.InterestPayment*0.3693)-((float64(m.Mortgage)*0.0035)/12), 2)
	if m.TaxDiscount < 0.00 {
		m.TaxDiscount = 0.00
	}
	m.NettPayment = util.RoundFloat(m.TotalPayment-m.TaxDiscount, 2)

	// totals
	m.Remainder = util.RoundFloat(m.Remainder-m.RedemptionPayment, 2)
	m.PayedInterest = util.RoundFloat(m.PayedInterest+m.InterestPayment, 2)
	m.PayedRedemption = util.RoundFloat(m.PayedRedemption+m.RedemptionPayment, 2)
	m.TotalPayed = util.RoundFloat(m.PayedInterest+m.PayedRedemption, 2)
	m.NettPayed = util.RoundFloat(m.NettPayed+m.NettPayment, 2)

	return nil
}

func (m *Mortgage) incrementMonth() {
	if m.Month+1 > 12 {
		m.Month = 1
		m.Year++
	} else {
		m.Month++
	}
}
