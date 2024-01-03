package period

import (
	"errors"
	"fmt"
	bd "github.com/RedmonkeyDF/mkf.golib.radiodates/broadcast"
	"strconv"
	"time"
)

const (
	PERIOD_YEAR_MIN = 1990
	PERIOD_YEAR_MAX = 2050
)

type Periodint uint32

func decodeValidPeriod(p Periodint) (Periodint, Periodint) {

	pstr := strconv.Itoa(int(p))
	yrint, _ := strconv.Atoi(pstr[0:4])
	mnint, _ := strconv.Atoi(pstr[4:])

	return Periodint(yrint), Periodint(mnint)
}

func MNPeriodValid(aperiod Periodint) bool {

	pstr := fmt.Sprintf("%d", aperiod)
	if len(pstr) != 6 {

		return false
	}

	yr, _ := strconv.Atoi(pstr[0:4])
	mn, _ := strconv.Atoi(pstr[4:])

	if mn > 12 || mn < 1 {

		return false
	}

	if yr > PERIOD_YEAR_MAX || yr < PERIOD_YEAR_MIN {

		return false
	}

	return true
}

func MNPeriodYRMN(ayr, amn Periodint) (Periodint, error) {

	cand, _ := strconv.Atoi(fmt.Sprintf("%d%.2d", ayr, amn))

	if !MNPeriodValid(Periodint(cand)) {

		return 0, errors.New(fmt.Sprintf("MNPeriodYRMN recieved invalid input: \"YR: %d MN: %d\"", ayr, amn))
	}

	return Periodint(cand), nil
}

func MNPeriodDT(t time.Time) (Periodint, error) {

	return MNPeriodYRMN(Periodint(bd.BCYearDT(t)), Periodint(bd.BCMonthDT(t)))
}

func MNPeriodDecode(p Periodint) (Periodint, Periodint, error) {

	if !MNPeriodValid(p) {

		return 0, 0, errors.New(fmt.Sprintf("MNPeriodDecode - invalid period input - \"%d\"", p))
	}

	yr, mn := decodeValidPeriod(p)

	return yr, mn, nil
}

func MNPeriodCurrent() Periodint {

	per, _ := MNPeriodDT(time.Now())

	return per
}

func MNPeriodstrYRMN(ayr, amn Periodint) (string, error) {

	p, errp := MNPeriodYRMN(ayr, amn)
	if errp != nil {
		return "", errp
	}

	return strconv.Itoa(int(p)), nil
}

func MNPeriodstrDT(t time.Time) (string, error) {

	yr := bd.BCYearDT(t)
	mn := bd.BCMonthDT(t)

	return MNPeriodstrYRMN(Periodint(yr), Periodint(mn))
}

func WKPeriodValid(p Periodint) bool {

	pstr := strconv.Itoa(int(p))

	if len(pstr) != 6 {
		return false
	}

	yr, _ := strconv.Atoi(pstr[0:4])
	wk, _ := strconv.Atoi(pstr[4:])

	if wk < 1 {

		return false
	}

	if yr > PERIOD_YEAR_MAX || yr < PERIOD_YEAR_MIN {

		return false
	}

	wksinyr := bd.WeeksInBCYearYR(bd.Bcdint(yr))

	if wk > int(wksinyr) {

		return false
	}

	return true
}

func WKPeriodYRWK(ayr, awk Periodint) (Periodint, error) {

	cand, _ := strconv.Atoi(fmt.Sprintf("%d%.2d", ayr, awk))

	if !WKPeriodValid(Periodint(cand)) {

		return 0, errors.New(fmt.Sprintf("WKPeriodYRWK recieved invalid input: \"YR: %d WK: %d\"", ayr, awk))
	}

	return Periodint(cand), nil
}

func WKPeriodDT(t time.Time) (Periodint, error) {

	return WKPeriodYRWK(Periodint(bd.BCYearDT(t)), Periodint(bd.BCWeekDT(t)))
}

func WKPeriodDecode(p Periodint) (Periodint, Periodint, error) {

	if !WKPeriodValid(p) {

		return 0, 0, errors.New(fmt.Sprintf("WKPeriodDecode - invalid period input - \"%d\"", p))
	}

	yr, wk := decodeValidPeriod(p)

	return yr, wk, nil
}

func WKPeriodstrYRWK(ayr, awk Periodint) (string, error) {

	p, errp := WKPeriodYRWK(ayr, awk)
	if errp != nil {
		return "", errp
	}

	return strconv.Itoa(int(p)), nil
}

func WKPeriodstrDT(t time.Time) (string, error) {

	ayr := bd.BCYearDT(t)
	awk := bd.BCWeekDT(t)

	return WKPeriodstrYRWK(Periodint(ayr), Periodint(awk))
}

func WKPeriodCurrent() Periodint {

	per, _ := WKPeriodDT(time.Now())

	return per
}

func QTRPeriodValid(p Periodint) bool {

	pstr := fmt.Sprintf("%d", p)
	if len(pstr) != 6 {

		return false
	}

	yr, _ := strconv.Atoi(pstr[0:4])
	qtr, _ := strconv.Atoi(pstr[4:])

	if qtr > 4 || qtr < 1 {

		return false
	}

	if yr > PERIOD_YEAR_MAX || yr < PERIOD_YEAR_MIN {

		return false
	}

	return true
}

func QTRPeriodYRQTR(ayr, aqtr Periodint) (Periodint, error) {

	cand, _ := strconv.Atoi(fmt.Sprintf("%d%.2d", ayr, aqtr))

	if !QTRPeriodValid(Periodint(cand)) {

		return 0, errors.New(fmt.Sprintf("QTRPeriodYRQTR recieved invalid input: \"YR: %d MN: %d\"", ayr, aqtr))
	}

	return Periodint(cand), nil
}

func QTRPeriodDT(t time.Time) (Periodint, error) {

	yr := bd.BCYearDT(t)
	qtr := bd.BCQuarterDT(t)

	return QTRPeriodYRQTR(Periodint(yr), Periodint(qtr))
}

func QTRPeriodDecode(p Periodint) (Periodint, Periodint, error) {

	if !QTRPeriodValid(p) {

		return 0, 0, errors.New(fmt.Sprintf("QTRPeriodDecode - invalid period input - \"%d\"", p))
	}

	yr, qtr := decodeValidPeriod(p)

	return yr, qtr, nil
}

func QTRPeriodstrYRQTR(ayr, aqtr Periodint) (string, error) {

	p, errp := QTRPeriodYRQTR(ayr, aqtr)
	if errp != nil {
		return "", errp
	}

	return strconv.Itoa(int(p)), nil
}

func QTRPeriodstrDT(t time.Time) (string, error) {

	yr := bd.BCYearDT(t)
	qtr := bd.BCQuarterDT(t)

	return QTRPeriodstrYRQTR(Periodint(yr), Periodint(qtr))
}

func WKPeriodsContained(astartperiod, afinperiod Periodint) (Periodint, error) {

	if astartperiod > afinperiod {

		return 0, errors.New(fmt.Sprintf("Start week period \"%d\" is bigger than end week period \"%d\".", astartperiod, afinperiod))
	}

	styr, stwk, errdecst := WKPeriodDecode(astartperiod)
	if errdecst != nil {

		return 0, errdecst
	}

	fiyr, fiwk, errdecfi := WKPeriodDecode(afinperiod)

	if errdecfi != nil {

		return 0, errdecfi
	}

	dtst := bd.StartOfBCWeekYRWK(bd.Bcdint(styr), bd.Bcdint(stwk))
	dtfi := bd.StartOfBCWeekYRWK(bd.Bcdint(fiyr), bd.Bcdint(fiwk))

	return Periodint(((bd.EndOfBCWeekDT(dtfi).Sub(bd.StartOfBCWeekDT(dtst)).Hours() / 24) / 7) + 1), nil
}

func WKPeriodSlice(astperiod, afiperiod Periodint) ([]Periodint, error) {

	if astperiod > afiperiod {

		return nil, errors.New(fmt.Sprintf("Start week period \"%d\" is bigger than end week period \"%d\".", astperiod, afiperiod))
	}

	if !WKPeriodValid(astperiod) {

		return nil, errors.New(fmt.Sprintf("Start period \"%d\" is invalid.", astperiod))
	}

	if !WKPeriodValid(afiperiod) {

		return nil, errors.New(fmt.Sprintf("End period \"%d\" is invalid.", afiperiod))
	}

	numperiods, errnumperiods := WKPeriodsContained(astperiod, afiperiod)
	if errnumperiods != nil {

		return nil, errnumperiods
	}

	pyr, pwk := decodeValidPeriod(afiperiod)

	ret := make([]Periodint, numperiods)

	for i := len(ret) - 1; i > -1; i-- {

		ret[i] = (pyr * 100) + pwk
		pwk--
		if pwk < 1 {

			pyr--
			pwk = Periodint(bd.WeeksInBCYearYR(bd.Bcdint(pyr)))
		}
	}

	return ret, nil
}

func WKPeriodSubtractWeeks(aperiod, weekstosub Periodint) (error, Periodint) {

	yr, wk, errdec := WKPeriodDecode(aperiod)
	if errdec != nil {

		return errdec, aperiod
	}

	for i := weekstosub; i > 0; i-- {

		wk--
		if wk < 1 {

			yr--
			wk = Periodint(bd.WeeksInBCYearYR(bd.Bcdint(yr)))
		}
	}

	per, errper := WKPeriodYRWK(yr, wk)
	if errper != nil {

		return errper, aperiod
	}

	return nil, per
}

func WKPeriodAddWeeks(aperiod, weekstoadd Periodint) (error, Periodint) {

	yr, wk, errdec := WKPeriodDecode(aperiod)
	if errdec != nil {

		return errdec, aperiod
	}

	wksinyr := bd.WeeksInBCYearYR(bd.Bcdint(yr))

	for i := 0; i < int(weekstoadd); i++ {

		wk++
		if wk > Periodint(wksinyr) {

			wk = 1
			yr++
			wksinyr = bd.WeeksInBCYearYR(bd.Bcdint(yr))
		}
	}

	per, errper := WKPeriodYRWK(yr, wk)
	if errper != nil {

		return errper, aperiod
	}

	return nil, per
}
