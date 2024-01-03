package broadcast

import (
	"testing"
	"time"
)

const DTFORMAT = "2006-01-02"

func TestBCYearDT(t *testing.T) {

	tst := make(map[string]Bcdint)
	tst["2017-01-01"] = 2017
	tst["2016-12-31"] = 2017
	tst["2015-12-30"] = 2016
	tst["2016-06-29"] = 2016
	tst["2017-11-30"] = 2017
	tst["2018-05-01"] = 2018
	tst["2018-11-29"] = 2018
	tst["2017-12-31"] = 2017
	tst["2016-12-26"] = 2017


	for input, expected := range tst {

		t.Logf("TestBCYearDT - input \"%s\" - expected \"%d\"", input, expected)

		dt, errdt := time.Parse(DTFORMAT, input)
		if errdt != nil {

			t.Errorf("TestBCYearDT - error in input data: \"%s\"", input)
		}

		got := BCYearDT(dt)

		if got != expected {

			t.Errorf("TestBCYearDT test failure.  Input \"%s\", got %d, expected %d", input, got, expected)
		}
	}
}

func TestBCMonthDT(t *testing.T) {

	tst := make(map[string]Bcdint)

	tst["2017-01-01"] = 1
	tst["2016-12-31"] = 1
	tst["2015-12-30"] = 1
	tst["2016-06-29"] = 7
	tst["2017-11-30"] = 12
	tst["2018-05-01"] = 5
	tst["2018-11-29"] = 12
	tst["2017-12-31"] = 12

	for input, expected := range tst {

		t.Logf("TestBCMonthDT - input \"%s\" - expected \"%d\"", input, expected)

		dt, errdt := time.Parse(DTFORMAT, input)
		if errdt != nil {
			t.Errorf("TestBCMonthDT error in test data \"%s\"", input)
		}

		got := BCMonthDT(dt)

		if got != expected {

			t.Errorf("TestBCMonthDT test failure: input\"%s\", expected \"%d\", got \"%d\"", input, expected, got)
		}
	}
}

func TestBCWeekDT(t *testing.T) {

	tst := make(map[string]Bcdint)
	tst["2017-01-01"] = 1
	tst["2016-12-31"] = 1
	tst["2015-12-30"] = 1
	tst["2016-06-29"] = 27
	tst["2017-11-30"] = 49
	tst["2018-05-01"] = 18
	tst["2018-11-29"] = 48
	tst["2017-12-31"] = 53
	tst["2017-01-30"] = 6

	for input, expected := range tst {

		t.Logf("TestBCWeekDT - input \"%s\" - expected \"%d\"", input, expected)

		dt, errdt := time.Parse(DTFORMAT, input)
		if errdt != nil {
			t.Errorf("TestBCWeekDT error in test data \"%s\"", input)
		}

		got := BCWeekDT(dt)

		if got != expected {

			t.Errorf("TestBCWeekDT test failure: input\"%s\", expected \"%d\", got \"%d\"", input, expected, got)
		}
	}

}

func TestStartOfBCWeekDT(t *testing.T) {

	dtmap := make(map[string]string)
	dtmap["2020-03-27"] = "2020-03-23"
	dtmap["2020-03-23"] = "2020-03-23"
	dtmap["2019-01-03"] = "2018-12-31"
	dtmap["2017-01-01"] = "2016-12-26"
	dtmap["2016-01-01"] = "2015-12-28"
	dtmap["2018-01-01"] = "2018-01-01"
	dtmap["2017-07-01"] = "2017-06-26"
	dtmap["2017-12-31"] = "2017-12-25"
	dtmap["2016-12-26"] = "2016-12-26"

	for input, expected := range dtmap {

		t.Logf("TestStartOfBCWeekDT - input \"%s\" - expected \"%s\"", input, expected)

		dt, err := time.Parse(DTFORMAT, input)
		if err != nil {

			t.Errorf("TestStartOfBCWeekDT - Error in test data: \"%s\"", err)
		}

		dt = StartOfBCWeekDT(dt)

		if expected != dt.Format(DTFORMAT) {

			t.Errorf("TestStartOfBCWeekDT test failed.  Input \"%s\" Expected\"%s\", received \"%s\".", input, expected, dt.Format(DTFORMAT))
		}
	}
}

func doTestStartOfWeekYRWK(t *testing.T, YR, WK int, expected string) {

	t.Logf("TestStartOfBCWeekDT - input \"YR: %d WK: %d\" - expected \"%s\"", YR, WK, expected)

	dt := StartOfBCWeekYRWK(uint32(YR), uint32(WK))

	if dt.Format(DTFORMAT) != expected {

		t.Errorf("Test TestStartOfBCWeekYRWK failed.  Input: \"%d-%d\" expected: \"%s\", got: \"%s\".",  YR, WK, expected, dt.Format(DTFORMAT))
	}
}

func TestStartOfBCWeekYRWK(t *testing.T) {

	doTestStartOfWeekYRWK(t,2017,1,"2016-12-26")
	doTestStartOfWeekYRWK(t,2016,1,"2015-12-28")
	doTestStartOfWeekYRWK(t,2018,1,"2018-01-01")
	doTestStartOfWeekYRWK(t,2017,27,"2017-06-26")
	doTestStartOfWeekYRWK(t,2017,53,"2017-12-25")
}

func TestStartOfBCYearDT(t *testing.T) {

	tst := make(map[string]string)
	tst["2017-05-05"] = "2016-12-26"
	tst["2016-12-30"] = "2016-12-26"
	tst["2018-01-07"] = "2018-01-01"
	tst["2016-12-25"] = "2015-12-28"

	for input, expected := range tst {

		t.Logf("TestStartOfBCYearDT - input \"%s\" - expected \"%s\"", input, expected)

		indt, errin := time.Parse(DTFORMAT, input)
		if errin != nil {
			t.Errorf("TestStartOfBCYearDT - error in test data.  Input:\"%s\"", input)
		}

		dt := StartOfBCYearDT(indt)

		if dt.Format(DTFORMAT) != expected {

			t.Errorf("TestStartOfBCYearDT failed: input %s expected %s got %s", input, expected, dt.Format(DTFORMAT))
		}
	}
}

func TestStartOfBCYearYR(t *testing.T) {

	tst := make(map[int]string)
	tst[2017] = "2016-12-26"
	tst[2014] = "2013-12-30"
	tst[2018] = "2018-01-01"
	tst[2016] = "2015-12-28"

	for input, expected := range tst {

		t.Logf("TestStartOfBCYearYR - input \"%d\" - expected \"%s\"", input, expected)

		dt := StartOfBCYearYR(uint32(input))

		if dt.Format(DTFORMAT) != expected {

			t.Errorf("TestStartOfBCYearYR failed: input %d expected %s got %s", input, expected, dt.Format(DTFORMAT))
		}
	}

}

func TestBCQuarterDT(t *testing.T) {

	tst := make(map[string]Bcdint)
	tst["2017-01-01"] = 1
	tst["2016-12-28"] = 1
	tst["2017-12-31"] = 4
	tst["2016-03-30"] = 2
	tst["2017-07-20"] = 3

	for input, expected := range tst {

		t.Logf("TestBCQuarterDT - input \"%s\" - expected \"%d\"", input, expected)

		dt, errdt := time.Parse(DTFORMAT, input)
		if errdt != nil {
			t.Errorf("TestBCQuarterDT - Error in test data.  Input:\"%s\"", input)
		}

		got := BCQuarterDT(dt)

		if got != expected {

			t.Errorf("TestBCQuarterDT failed: input \"%s\" expected \"%d\" got \"%d\".", input, expected, got)
		}
	}
}

func TestEndOfBCYearDT(t *testing.T) {

	tst := make(map[string]string)

	tst["2017-05-05"] = "2017-12-31"
	tst["2016-12-31"] = "2017-12-31"
	tst["2016-05-05"] = "2016-12-25"
	tst["2018-01-01"] = "2018-12-30"

	for input, expected := range tst {

		t.Logf("TestEndOfBCYearDT - input \"%s\" - expected \"%s\"", input, expected)

		indt, errindt := time.Parse(DTFORMAT, input)
		if errindt != nil {

			t.Errorf("TestEndOfBCYearDT - error in input data \"%s\"", input)
		}

		dtgot := EndOfBCYearDT(indt)

		if dtgot.Format(DTFORMAT) != expected {

			t.Errorf("TestEndOfBCYearDT failed: input \"%s\" expected \"%s\" got \"%s\".", input, expected, dtgot.Format(DTFORMAT))
		}
	}
}

func TestEndOfBCYearYR(t *testing.T) {

	tst := make(map[Bcdint]string)

	tst[2017] = "2017-12-31"
	tst[2016] = "2016-12-25"
	tst[2015] = "2015-12-27"
	tst[2018] = "2018-12-30"

	for input, expected := range tst {

		t.Logf("TestEndOfBCYearYR - input \"%d\" - expected \"%s\"", input, expected)

		got := EndOfBCYearYR(input)

		if got.Format(DTFORMAT) != expected {

			t.Errorf("TestEndOfBCYearYR failed - input \"%d\" - expected \"%s\" - got \"%s\"", input, expected, got.Format(DTFORMAT))
		}
	}
}

func TestWeeksInBCYearDT(t *testing.T) {

	tst := make(map[string]Bcdint)

	tst["2017-05-05"] = 53
	tst["2018-05-05"] = 52
	tst["2016-05-05"] = 52

	for input, expected := range tst {

		t.Logf("TestWeeksInBCYearDT - input \"%s\" - expected \"%d\"", input, expected)

		dt, errdt := time.Parse(DTFORMAT, input)
		if errdt != nil {

			t.Errorf("TestWeeksInBCYearDT - error in test data \"%s\"", input)
		}

		got := WeeksInBCYearDT(dt)

		if got != expected {

			t.Errorf("TestWeeksInBCYearDT failed - input \"%s\" - expected \"%d\" - got \"%d\"", input, expected, got)
		}
	}
}

func TestWeeksinBCYearYR(t *testing.T) {

	tst := make(map[Bcdint]Bcdint)

	tst[2017] = 53
	tst[2018] = 52
	tst[2016] = 52

	for input, expected := range tst {

		t.Logf("TestWeeksinBCYearYR - input \"%d\" - expected \"%d\"", input, expected)

		got := WeeksInBCYearYR(input)

		if got != expected {

			t.Errorf("TestWeeksinBCYearYR failed - input \"%d\" - expected \"%d\" - got \"%d\"", input, expected, got)
		}
	}
}

func TestStartOfBCMonthDT(t *testing.T) {

	tst := make(map[string]string)

	tst["2017-05-05"] = "2017-05-01"
	tst["2016-01-08"] = "2015-12-28"
	tst["2018-12-29"] = "2018-11-26"
	tst["2018-04-20"] = "2018-03-26"
	tst["2016-12-26"] = "2016-12-26"

	for input, expected := range tst {

		t.Logf("TestStartOfBCMonthDT - input \"%s\" - expected \"%s\"", input, expected)

		dtin, errdtin := time.Parse(DTFORMAT, input)
		if errdtin != nil {
			t.Errorf("TestStartOfBCMonthDT - error in test data \"%s\"", input)
		}

		got := StartOfBCMonthDT(dtin)

		if got.Format(DTFORMAT) != expected {

			t.Errorf("TestStartOfBCMonthDT test fail - input \"%s\" - expected \"%s\" -  got \"%s\"", input, expected, got.Format(DTFORMAT))
		}
	}

}

func doTestStartOfBCMonthYRMN(t *testing.T, YR, MN Bcdint, expected string) {

	t.Logf("TestStartOfBCMonthYRMN - input \"YR: %d MN: %d\", expected \"%s\"", YR, MN, expected)

	got := StartOfBCMonthYRMN(YR, MN)

	if got.Format(DTFORMAT) != expected {

		t.Errorf("TestStartOfBCMonthYRMN - input \"YR: %d MN: %d\" - expected \"%s\" - got \"%s\"", YR, MN, expected, got.Format(DTFORMAT))
	}
}

func TestStartOfBCMonthYRMN(t *testing.T) {

	doTestStartOfBCMonthYRMN(t, 2017,05,"2017-05-01")
	doTestStartOfBCMonthYRMN(t, 2016,01,"2015-12-28")
	doTestStartOfBCMonthYRMN(t, 2018,12,"2018-11-26")
	doTestStartOfBCMonthYRMN(t, 2018,04,"2018-03-26")
	doTestStartOfBCMonthYRMN(t, 2016,12,"2016-11-28")
}

func TestEndOfBCMonthDT(t *testing.T) {

	tst := make(map[string]string)

	tst["2017-10-11"] = "2017-10-29"
	tst["2017-12-31"] = "2017-12-31"
	tst["2016-05-30"] = "2016-06-26"

	for input, expected := range tst {

		t.Logf("TestEndOfBCMonthDT - input \"%s\" - expected \"%s\"", input, expected)

		dtin, errdtin := time.Parse(DTFORMAT, input)
		if errdtin != nil {

			t.Errorf("TestEndOfBCMonthDT error in test data \"%s\"", input)
		}

		got := EndOfBCMonthDT(dtin)

		if got.Format(DTFORMAT) != expected {

			t.Errorf("TestEndOfBCMonthDT failed - input \"%s\" - expected \"%s\" - got \"%s\"", input, expected, got.Format(DTFORMAT))
		}
	}
}

func doTestEndOfBCMonthYRMN(t *testing.T, YR, MN Bcdint, expected string) {

	t.Logf("TestEndOfBCMonthYRMN - input \"YR: %d MN: %d\" - expected \"%s\"", YR, MN, expected)

	got := EndOfBCMonthYRMN(YR, MN)

	if got.Format(DTFORMAT) != expected {

		t.Errorf("TestEndOfBCMonthYRMN failed - input \"YR: %d MN: %d\" - expected \"%s\" - got \"%s\"", YR, MN, expected, got.Format(DTFORMAT))
	}
}

func TestEndOfBCMonthYRMN(t *testing.T) {

	doTestEndOfBCMonthYRMN(t, 2017,10,"2017-10-29")
	doTestEndOfBCMonthYRMN(t, 2017,12,"2017-12-31")
	doTestEndOfBCMonthYRMN(t, 2016,05,"2016-05-29")
}

func TestStartOfBCQuarterDT(t *testing.T) {

	tst := make(map[string]string)

	tst["2017-01-01"] = "2016-12-26"
	tst["2018-01-05"] = "2018-01-01"
	tst["2040-08-15"] = "2040-06-25"
	tst["2016-06-05"] = "2016-03-28"
	tst["2015-10-15"] = "2015-09-28"

	for input, expected := range tst {

		t.Logf("TestStartOfBCQuarterDT - input \"%s\" - expected \"%s\"", input, expected)

		dtin, errdtin := time.Parse(DTFORMAT, input)
		if errdtin != nil {

			t.Errorf("TestStartOfBCQuarterDT error in input data - input \"%s\"", input)
		}

		got := StartOfBCQuarterDT(dtin)

		if got.Format(DTFORMAT) != expected {

			t.Errorf("TestStartOfBCQuarterDT failed - input \"%s\" - expected \"%s\" - got \"%s\"", input, expected, got.Format(DTFORMAT))
		}
	}
}

func doTestStartOfBCQuarterYRQT(t *testing.T, YR, QTR Bcdint, expected string) {

	t.Logf("TestStartOfBCQuarterYRQT - input \"YR: %d QTR: %d\" - expected \"%s\"", YR, QTR, expected)

	got := StartOfBCQuarterYRQT(YR, QTR)

	if got.Format(DTFORMAT) != expected {

		t.Errorf("TestStartOfBCQuarterYRQT failed- input \"YR: %d QTR: %d\" - expected \"%s\" - got \"%s\"", YR, QTR, expected, got.Format(DTFORMAT))
	}
}

func TestStartOfBCQuarterYRQT(t *testing.T) {

	doTestStartOfBCQuarterYRQT(t, 2018, 1, "2018-01-01")
	doTestStartOfBCQuarterYRQT(t, 2017, 1, "2016-12-26")
	doTestStartOfBCQuarterYRQT(t, 2019, 1, "2018-12-31")
	doTestStartOfBCQuarterYRQT(t, 2020, 2, "2020-03-30")
	doTestStartOfBCQuarterYRQT(t, 2021, 4, "2021-09-27")
}

func TestEndOfBCQuarterDT(t *testing.T) {

	tst := make(map[string]string)
	tst["2021-12-01"] = "2021-12-26"
	tst["2020-09-29"] = "2020-12-27"
	tst["2017-05-12"] = "2017-06-25"

	for input, expected := range tst {

		t.Logf("TestEndOfBCQuarterDT - input \"%s\" - expected \"%s\"", input, expected)

		dtin, errdtin := time.Parse(DTFORMAT, input)
		if errdtin != nil {

			t.Errorf("TestEndOfBCQuarterDT - error in input data \"%s\"", input)
		}

		got := EndOfBCQuarterDT(dtin)

		if got.Format(DTFORMAT) != expected {

			t.Errorf("TestEndOfBCQuarterDT - input \"%s\" - expected \"%s\" - got \"%s\"", input, expected, got.Format(DTFORMAT))
		}
	}

}

func doTestEndOfBCQuarterYRQT(t *testing.T, YR, QTR Bcdint, expected string) {

	t.Logf("TestEndOfBCQuarterYRQT - input \"YR: %d QTR: %d\" - expected \"%s\"", YR, QTR, expected)

	got := EndOfBCQuarterYRQT(YR, QTR)

	if got.Format(DTFORMAT) != expected {

		t.Errorf("TestEndOfBCQuarterYRQT - input \"YR: %d QTR: %d\" - expected \"%s\" - got \"%s\"", YR, QTR, expected, got.Format(DTFORMAT))
	}
}

func TestEndOfBCQuarterYRQT(t *testing.T) {

	doTestEndOfBCQuarterYRQT(t, 2021, 4, "2021-12-26")
	doTestEndOfBCQuarterYRQT(t, 2020, 4, "2020-12-27")
	doTestEndOfBCQuarterYRQT(t,	2017, 2, "2017-06-25")
}

func TestWeeksInBCMonthDT(t *testing.T) {

	tst := make(map[string]Bcdint)

	tst["2017-01-01"] = 5
	tst["2017-02-01"] = 4
	tst["2017-03-01"] = 4
	tst["2017-04-01"] = 5
	tst["2017-05-01"] = 4
	tst["2017-06-01"] = 4
	tst["2017-07-01"] = 5
	tst["2017-08-01"] = 4
	tst["2017-09-01"] = 4
	tst["2017-10-01"] = 5
	tst["2017-11-01"] = 4
	tst["2017-12-01"] = 5

	for input, expected := range tst {

		t.Logf("TestWeeksInBCMonthDT - input \"%s\" - expected \"%d\"", input, expected)

		dtin, errdtin := time.Parse(DTFORMAT, input)
		if errdtin != nil {

			t.Errorf("TestWeeksInBCMonthDT - error in input data \"%s\"", input)
		}

		got := WeeksInBCMonthDT(dtin)

		if got != expected {

			t.Errorf("TestWeeksInBCMonthDT failed - input \"%s\" - expected \"%d\" - got \"%d\"", input, expected, got)
		}
	}
}

func doTestWeeksInBCMonthYRMN(t *testing.T, YR, MN, expected Bcdint) {

	t.Logf("TestWeeksInBCMonthYRMN - input \"YR: %d MN: %d\" - expected \"%d\"", YR, MN, expected)

	got := WeeksInBCMonthYRMN(YR, MN)

	if got != expected {

		t.Errorf("TestWeeksInBCMonthYRMN failed - input \"YR: %d MN: %d\" - expected \"%d\" - got \"%d\"", YR, MN, expected, got)
	}
}

func TestWeeksInBCMonthYRMN(t *testing.T) {

	doTestWeeksInBCMonthYRMN(t,2017,1,5)
	doTestWeeksInBCMonthYRMN(t,2017,2,4)
	doTestWeeksInBCMonthYRMN(t,2017,3,4)
	doTestWeeksInBCMonthYRMN(t,2017,4,5)
	doTestWeeksInBCMonthYRMN(t,2017,5,4)
	doTestWeeksInBCMonthYRMN(t,2017,6,4)
	doTestWeeksInBCMonthYRMN(t,2017,7,5)
	doTestWeeksInBCMonthYRMN(t,2017,8,4)
	doTestWeeksInBCMonthYRMN(t,2017,9,4)
	doTestWeeksInBCMonthYRMN(t,2017,10,5)
	doTestWeeksInBCMonthYRMN(t,2017,11,4)
	doTestWeeksInBCMonthYRMN(t,2017,12,5)
}

func TestEndOfBCWeekDT(t *testing.T) {

	tst := make(map[string]string)

	tst["2019-05-09"] = "2019-05-12"
	tst["2019-05-05"] = "2019-05-05"
	tst["2019-12-25"] = "2019-12-29"

	for input, expected := range tst {

		t.Logf("TestEndOfBCWeekDT - input \"%s\" - expected \"%s\"", input, expected)

		dtin, errdtin := time.Parse(DTFORMAT, input)
		if errdtin != nil {
			t.Errorf("TestEndOfBCWeekDT - error in input \"%s\"", input)
		}

		got := EndOfBCWeekDT(dtin)

		if got.Format(DTFORMAT) != expected {

			t.Errorf("TestEndOfBCWeekDT - input \"%s\" - expected \"%s\" - got \"%s\"", input, expected, got.Format(DTFORMAT))
		}
	}
}

func doTestEndOfBCWeekYRWK(t *testing.T, YR, WK Bcdint, expected string) {

	t.Logf("TestEndOfBCWeekYRMN - input \"YR: %d WK: %d\" - expected \"%s\"", YR, WK, expected)

	got := EndOfBCWeekYRWK(YR, WK)

	if got.Format(DTFORMAT) != expected {

		t.Errorf("TestEndOfBCWeekYRMN - input \"YR: %d WK: %d\" - expected \"%s\" - got \"%s\"", YR, WK, expected, got.Format(DTFORMAT))
	}
}

func TestEndOfBCWeekYRWK(t *testing.T) {

	doTestEndOfBCWeekYRWK(t, 2019, 47, "2019-11-24")
	doTestEndOfBCWeekYRWK(t, 2018, 44, "2018-11-04")
	doTestEndOfBCWeekYRWK(t, 2017, 53, "2017-12-31")
	doTestEndOfBCWeekYRWK(t, 2017, 1, "2017-01-01")

}