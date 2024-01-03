package broadcast

import (
	"time"
)

type Bcdint = uint32

func datesYMDEqual(t1, t2 time.Time) bool {

	return t1.Year() == t2.Year() && t1.YearDay() == t2.YearDay()
}

//Broadcast year methods
func BCYearDT(t time.Time) Bcdint {

	sonextyr := StartOfBCWeekYRWK(Bcdint(t.Year() + 1), 1)

	if datesYMDEqual(t, sonextyr) || t.After(sonextyr) {

		return Bcdint(t.Year() + 1)
	}

	return Bcdint(t.Year())
}

func BCMonthDT(t time.Time) Bcdint {

	opyr := t.Year()
	opmn := t.Month()

	opmn++
	for opmn > 12 {
		opmn -= 12
		opyr++
	}

	sonextmn := StartOfBCWeekDT(time.Date(opyr, opmn, 1, 0, 0, 0, 0, time.UTC))
	if datesYMDEqual(t, sonextmn) || t.After(sonextmn) {

		return Bcdint(opmn)
	}

	return Bcdint(t.Month())
}

func BCWeekDT(t time.Time) Bcdint {

	/*soy := StartOfBCYearDT(t)

	t.Sub(soy).Hours() / 24

	(( / 7) + 1*/

	return Bcdint(((t.Sub(StartOfBCYearDT(t)).Hours() / 24) / 7) + 1)
}

func BCQuarterDT(t time.Time) Bcdint {

	switch mn := BCMonthDT(t); mn {
	case 4,5,6: return 2
	case 7,8,9: return 3
	case 10,11,12: return 4
	default: return 1
	}
}

func StartOfBCWeekDT(t time.Time) time.Time {

	dow := t.Weekday()
	dow --
	if dow < 0 {
		dow = 6
	}

	return t.AddDate(0, 0, -int(dow))
}

func StartOfBCWeekYRWK(ayear, aweek Bcdint) time.Time {

	return StartOfBCWeekDT(time.Date(int(ayear), 1, 1, 0, 0, 0, 0, time.UTC)).AddDate(0, 0, int(((aweek - 1) * 7)))
}

func StartOfBCYearDT(t time.Time) time.Time {

	thedt := t

	if thedt.Month() == 12 {

		startnext := StartOfBCWeekYRWK(Bcdint(thedt.Year()+1), 1)
		startthedt := StartOfBCWeekDT(thedt)
		if startnext.Year() == startthedt.Year() && startnext.YearDay() == startthedt.YearDay() {

			thedt = time.Date(thedt.Year()+1, 1, 1, 0, 0, 0, 0, time.UTC)
		}
	}

	return StartOfBCWeekDT(time.Date(thedt.Year(), 1, 1, 0, 0, 0, 0, time.UTC))
}

func StartOfBCYearYR(ayr Bcdint) time.Time {

	return StartOfBCYearDT(time.Date(int(ayr), 1, 1, 0, 0, 0, 0, time.UTC))
}

func EndOfBCYearDT(t time.Time) time.Time {

	yr := BCYearDT(t)

	return StartOfBCYearYR(yr + 1).AddDate(0, 0, -1)
}

func EndOfBCYearYR(ayr Bcdint) time.Time {

	return EndOfBCYearDT(time.Date(int(ayr), 1, 1, 0, 0, 0, 0, time.UTC))
}

func WeeksInBCYearDT(t time.Time) Bcdint {

	return Bcdint(((EndOfBCYearDT(t).Sub(StartOfBCYearDT(t)).Hours() / 24) / 7) + 1)
}

func WeeksInBCYearYR(ayr Bcdint) Bcdint {

	return WeeksInBCYearDT(time.Date(int(ayr), 1, 1, 0, 0, 0, 0, time.UTC))
}

func StartOfBCMonthDT(t time.Time) time.Time {

	yr := BCYearDT(t)
	mn := BCMonthDT(t)

	return StartOfBCWeekDT(time.Date(int(yr), time.Month(mn), 1, 0, 0, 0, 0, time.UTC))
}

func StartOfBCMonthYRMN(ayr, amn Bcdint) time.Time {

	return StartOfBCMonthDT(time.Date(int(ayr), time.Month(amn), 1, 0, 0, 0, 0, time.UTC))
}

func EndOfBCMonthDT(t time.Time) time.Time {

	ayr := Bcdint(BCYearDT(t))
	amn := Bcdint(BCMonthDT(t))

	amn ++

	for amn > 12 {

		amn -= 12
		ayr ++
	}

	return StartOfBCMonthYRMN(ayr, amn).AddDate(0, 0, -1)
}

func EndOfBCMonthYRMN(ayr, amn Bcdint) time.Time {

	return EndOfBCMonthDT(time.Date(int(ayr), time.Month(amn), 1, 0, 0, 0, 0, time.UTC))
}

func StartOfBCQuarterDT(t time.Time) time.Time {

	yr := BCYearDT(t)
	qtr := BCQuarterDT(t)

	mn := ((qtr * 3) - 3) + 1

	return StartOfBCMonthYRMN(yr, mn)
}

func StartOfBCQuarterYRQT(ayr, aqtr Bcdint) time.Time {

	return StartOfBCMonthYRMN(ayr, ((aqtr * 3) - 3) + 1)
}

func EndOfBCQuarterDT(t time.Time) time.Time {

	yr := BCYearDT(t)
	qtr := BCQuarterDT(t)
	mn := (qtr * 3) + 1

	for mn > 12 {

		mn -= 12
		yr ++
	}

	return StartOfBCMonthYRMN(Bcdint(yr), Bcdint(mn)).AddDate(0, 0, -1)
}

func EndOfBCQuarterYRQT(ayr, aqtr Bcdint) time.Time {

	return EndOfBCQuarterDT(StartOfBCQuarterYRQT(ayr, aqtr))
}

func WeeksInBCMonthDT(t time.Time) Bcdint {

	return Bcdint(((EndOfBCMonthDT(t).Sub(StartOfBCMonthDT(t)).Hours() / 24) / 7) + 1)
}

func WeeksInBCMonthYRMN(ayr, amn Bcdint) Bcdint {

	return WeeksInBCMonthDT(time.Date(int(ayr), time.Month(amn), 1, 0, 0, 0, 0, time.UTC))
}

func EndOfBCWeekDT(t time.Time) time.Time {

	return StartOfBCWeekDT(t).AddDate(0, 0, 6)
}

func EndOfBCWeekYRWK(ayr, awk Bcdint) time.Time {

	return StartOfBCWeekYRWK(ayr, awk).AddDate(0, 0, 6)
}


