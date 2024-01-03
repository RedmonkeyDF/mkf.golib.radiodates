package period

import (
	"fmt"
	"testing"
	"time"
)

const DTFORMAT = "2006-01-02"

func TestMNPeriodValid(t *testing.T) {

	t.Log("MNPeriodValid acceptance - testing \"func MNPeriodValid(aperiod Periodint) (bool, error)\"")

	tst := make(map[Periodint]bool)

	tst[201601] = true
	tst[201602] = true
	tst[20160] = false
	tst[2016031423] = false
	tst[201613] = false
	tst[201600] = false
	tst[190001] = false
	tst[281001] = false

	for input, expected := range tst {

		t.Logf("MNPeriodValid test - input \"%d\" - expected \"%t\"", input, expected)

		got := MNPeriodValid(input)

		if got != expected {

			t.Errorf("MNPeriodValid test failed - input \"%d\" - expected \"%t\" - got \"%t\"", input, expected, got)
		}
	}
}

func doTestMNPeriodYRMNAcceptance(t *testing.T, ayr, amn, expected Periodint) {

	t.Logf("MNPeriodYRMN accept test - input \"YR: %d MN: %d\" - expected \"%d\"", ayr, amn, expected)

	got, errmnp := MNPeriodYRMN(ayr, amn)
	if errmnp != nil {

		t.Errorf("MNPeriodYRMN accept test error in input \"YR: %d MN: %d\"", ayr, amn)
	}

	if got != expected {

		t.Errorf("MNPeriodYRMN accept test failed - input \"YR: %d MN: %d\" - expected \"%d\" - got \"%d\"", ayr, amn, expected, got)
	}
}

func TestMNPeriodYRMNAcceptance(t *testing.T) {

	t.Logf("MNPeriodYRMN acceptance testing \"MNPeriodYRMN(ayr, amn Periodint) (Periodint, error)\"")

	doTestMNPeriodYRMNAcceptance(t, 2017, 01, 201701)
	doTestMNPeriodYRMNAcceptance(t, 2017, 12, 201712)
	doTestMNPeriodYRMNAcceptance(t, 2016, 11, 201611)
	doTestMNPeriodYRMNAcceptance(t, 2018, 1, 201801)
}

func doTestMNPeriodYRMNRejection(t *testing.T, ayr, amn Periodint) {

	expected := fmt.Sprintf("MNPeriodYRMN recieved invalid input: \"YR: %d MN: %d\"", ayr, amn)

	t.Logf("MNPeriodYRMN reject test - input \"YR: %d MN: %d\" - expected \"%s\"", ayr, amn, expected)

	_, errgot := MNPeriodYRMN(ayr, amn)

	if errgot != nil {

		if errgot.Error() != expected {

			t.Errorf("MNPeriodYRMN reject test failed - input \"YR: %d MN: %d\" - expected \"%s\" - got \"%s\"", ayr, amn, expected, errgot)
		}
	} else {

		t.Errorf("MNPeriodYRMN reject test failed - input \"YR: %d MN: %d\" - expected \"%s\" - no error thrown", ayr, amn, expected)
	}
}

func TestMNPeriodYRMNRejection(t *testing.T) {

	t.Log("MNPeriodYRMN rejection testing \"MNPeriodYRMN(ayr, amn Periodint) (Periodint, error)\"")

	doTestMNPeriodYRMNRejection(t, 2018, 13)
	doTestMNPeriodYRMNRejection(t, 2018, 0)
	doTestMNPeriodYRMNRejection(t, 1899, 12)
	doTestMNPeriodYRMNRejection(t, 2500, 12)
}

func TestMNPeriodDTAccept(t *testing.T) {

	tst := make(map[string]Periodint)

	tst["2017-01-01"] = 201701
	tst["2017-01-30"] = 201702
	tst["2017-12-31"] = 201712
	tst["2015-12-30"] = 201601

	t.Log("MNPeriodDT acceptance testing \"func MNPeriodDT(t time.Time) (Periodint, error)\"")

	for input, expected := range tst {

		t.Logf("MNPeriodDT acceptance case - input \"%s\" - expected \"%d\"", input, expected)

		dtin, errdtin := time.Parse(DTFORMAT, input)
		if errdtin != nil {

			t.Errorf("MNPeriodDT acceptance case error in input \"%s\" - expected \"%d\"", input, expected)
		}

		got, _ := MNPeriodDT(dtin)

		if got != expected {

			t.Errorf("MNPeriodDT test failed - input \"%s\" - expected \"%d\" - got \"%d\"", input, expected, got)
		}
	}
}

func doTestMNPeriodDecodeAccept(t *testing.T, p, expectedyr, expectedmn Periodint) {

	t.Logf("MNPeriodDecode accept test - input \"%d\" - expected \"YR: %d MN: %d\"", p, expectedyr, expectedmn)

	gotyr, gotmn, _ := MNPeriodDecode(p)

	if gotyr != expectedyr || gotmn != expectedmn {

		t.Errorf("MNPeriodDecode accept test failed - input \"%d\" - expected \"YR: %d MN: %d\" - got \"YR: %d MN: %d\"", p, expectedyr, expectedmn, gotyr, gotmn)
	}
}

func TestMNPeriodDecodeAccept(t *testing.T) {

	t.Log("MNPeriodDecode acceptance testing \"func MNPeriodDecode(p Periodint) (Periodint, Periodint, error)\"")

	doTestMNPeriodDecodeAccept(t, 201212, 2012, 12)
	doTestMNPeriodDecodeAccept(t, 201801, 2018, 1)
	doTestMNPeriodDecodeAccept(t, 201508, 2015, 8)
}

func TestMNPeriodDecodeReject(t *testing.T) {

	t.Log("MNPeriodDecode rejection testing \"func MNPeriodDecode(p Periodint) (Periodint, Periodint, error)\"")

	tst := []Periodint{2012, 201313, 201200, 205001, 180005}

	for input := range tst {

		_, _, errpd := MNPeriodDecode(Periodint(input))

		if errpd != nil {

			expected := fmt.Sprintf("MNPeriodDecode - invalid period input - \"%d\"", input)
			if errpd.Error() != expected {

				t.Errorf("MNPeriodDecode rejection case failed - input \"%d\" - expected \"%s\" - got \"%s\"", input, expected, errpd.Error())
			}
		} else {

			t.Errorf("MNPeriodDecode rejection case failed - input \"%d\" - no error thrown", input)
		}
	}

}

func TestMNPeriodstrYRMNAccept(t *testing.T) {

	t.Log("MNPeriodstrYRMN acceptance testing - \"func MNPeriodstrYRMN(ayr, amn string) (Periodint,  error)\"")

	tst := make(map[string][2]Periodint)

	tst["201601"] = [2]Periodint{2016, 1}
	tst["201712"] = [2]Periodint{2017, 12}
	tst["201102"] = [2]Periodint{2011, 2}

	for expected, input := range tst {

		t.Logf("MNPeriodstrYRMN accept test case - input \"YR: %d MN: %d\" - expect \"%s\"", input[0], input[1], expected)

		got, errgot := MNPeriodstrYRMN(input[0], input[1])

		if errgot != nil {

			t.Errorf("MNPeriodstrYRMN test failed with unexpected error - input \"YR: %d MN: %d\" - error \"%s\"", input[0], input[1], errgot)
		}

		if got != expected {

			t.Errorf("MNPeriodstrYRMN test failed - input \"YR: %d MN: %d\" - expected \"%s\" - got \"%s\"", input[0], input[1], expected, got)
		}
	}
}

func TestMNPeriodstrYRMNReject(t *testing.T) {

	type YRMN struct {
		YR Periodint
		MN Periodint
	}

	t.Log("MNPeriodstrYRMN rejection testing - \"func MNPeriodstrYRMN(ayr, amn string) (Periodint,  error)\"")

	tst := []YRMN{{12, 300}, {2010, 0}, {2051, 1}, {2010, 13}}

	for _, val := range tst {

		t.Logf("MNPeriodstrYRMN reject test case - input \"YR: %d MN: %d\"", val.YR, val.MN)

		_, errgot := MNPeriodstrYRMN(val.YR, val.MN)

		if errgot != nil {

			expected := fmt.Sprintf("MNPeriodYRMN recieved invalid input: \"YR: %d MN: %d\"", val.YR, val.MN)
			if errgot.Error() != expected {

				t.Errorf("MNPeriodstrYRMN reject test case failed - input \"YR: %d MN: %d\" - expected \"%s\" - got \"%s\"", val.YR, val.MN, expected, errgot.Error())
			}

		} else {

			t.Errorf("MNPeriodstrYRMN reject test case failed - input \"YR: %d MN: %d\" - no error thrown", val.YR, val.MN)
		}
	}
}

func TestWKPeriodValid(t *testing.T) {

	tst := make(map[Periodint]bool)

	tst[201701] = true
	tst[201702] = true
	tst[20171] = false
	tst[201700] = false
	tst[201653] = false
	tst[201753] = true
	tst[189901] = false
	tst[205552] = false
	tst[2053245552] = false

	t.Log("WKPeriodValid acceptance testing. \"func WKPeriodValid(p Periodint) bool\"")

	for input, expected := range tst {

		t.Logf("WKPeriodValid test case.  Input \"%d\".  Expected \"%t\"", input, expected)

		got := WKPeriodValid(input)

		if got != expected {

			t.Logf("WKPeriodValid test case failed.  Input \"%d\".  Expected \"%t\".  Got \"%t\"", input, expected, got)
		}
	}
}

func TestWKPeriodYRWK(t *testing.T) {

	tst := make(map[Periodint][]Periodint)

	tst[201601] = []Periodint{2016, 01}
	tst[201722] = []Periodint{2017, 22}
	tst[201811] = []Periodint{2018, 11}
	tst[201801] = []Periodint{2018, 1}

	t.Log("WKPeriodYRWK acceptance testing.  \"func WKPeriodYRWK(ayr, awk Periodint) (Periodint, error)\"")

	for expected, input := range tst {

		t.Logf("WKPeriodYRWK test case.  Input \"YR: %d WK: %d\".  Expected: \"%d\".", input[0], input[1], expected)

		got, errgot := WKPeriodYRWK(input[0], input[1])

		if errgot != nil {

			t.Logf("WKPeriodYRWK failed.  Input \"YR: %d WK: %d\".  Expected: \"%d\".  Unexpected error \"%s\"", input[0], input[1], expected, errgot)
		}

		if got != expected {

			t.Logf("WKPeriodYRWK failed.  Input \"YR: %d WK: %d\".  Expected: \"%d\".  Got: \"%d\"", input[0], input[1], expected, got)
		}
	}
}

func TestWKPeriodDT(t *testing.T) {

	tst := make(map[string]Periodint)

	tst["2017-02-02"] = 201706
	tst["2017-01-30"] = 201706
	tst["2017-12-31"] = 201753
	tst["2015-12-30"] = 201601

	t.Log("WKPeriodDT testing.  \"func WKPeriodDT(t time.Time) (Periodint, error)\"")

	for input, expected := range tst {

		t.Logf("WKPeriodDT test case.  Input \"%s\".  Expected \"%d\".", input, expected)

		dtin, errdtin := time.Parse(DTFORMAT, input)
		if errdtin != nil {

			t.Errorf("WKPeriodDT test failed - error in input data.  Input \"%s\".  Expected \"%d\".  Error:\"%s\"", input, expected, errdtin)
		}

		got, errgot := WKPeriodDT(dtin)
		if errgot != nil {

			t.Errorf("WKPeriodDT test failed - Unexpected error.  Input \"%s\".  Expected \"%d\".  Error:\"%s\"", input, expected, errgot)
		}

		if got != expected {

			t.Errorf("WKPeriodDT test failed.  Input \"%s\".  Expected \"%d\".  Got:\"%d\"", input, expected, got)
		}
	}
}

func TestWKPeriodYRWKReject(t *testing.T) {

	tst := [][]Periodint{
		[]Periodint{2017, 0},
		[]Periodint{2017, 54},
		[]Periodint{2016, 53},
		[]Periodint{1905, 12},
		[]Periodint{2500, 10},
	}

	t.Log("WKPeriodYRWK rejection tests. \"func WKPeriodYRWK(ayr, awk Periodint) (Periodint, error)\"")

	for _, input := range tst {

		t.Logf("WKPeriodYRWK rejection test case.  Input\"YR %d WK: %d\"", input[0], input[1])
		_, errp := WKPeriodYRWK(input[0], input[1])

		if errp == nil {

			t.Errorf("WKPeriodYRWK rejection test failed - no error thrown.  Input\"YR %d WK: %d\"", input[0], input[1])
		} else {

			expected := fmt.Sprintf("WKPeriodYRWK recieved invalid input: \"YR: %d WK: %d\"", input[0], input[1])
			if errp.Error() != expected {

				t.Errorf("WKPeriodYRWK rejection test failed.  Input\"YR %d WK: %d\".  Expected \"%s\".  Got \"%s\"", input[0], input[1], expected, errp.Error())
			}
		}
	}
}

func TestWKPeriodDTReject(t *testing.T) {

	tst := []string{"1900-01-01", "2700-01-01"}

	t.Log("WKPeriodDT(t time.Time) rejection tests. \"func WKPeriodDT(t time.Time) (Periodint, error)\"")

	for _, input := range tst {

		t.Logf("WKPeriodDT rejection test case.  Input\"%s\"", input)

		dtin, errdtin := time.Parse(DTFORMAT, input)
		if errdtin != nil {

			t.Errorf("WKPeriodDT rejection test case failed with input data error.  Input\"%s\"", input)
		}

		_, errgot := WKPeriodDT(dtin)

		if errgot == nil {

			t.Errorf("WKPeriodDT rejection test failed - no error thrown.  Input\"%s\"", input)
		} else {

			expected := "WKPeriodYRWK recieved invalid input"

			if len(errgot.Error()) < 36 && errgot.Error()[0:36] != expected {

				t.Errorf("WKPeriodYRWK rejection test failed.  Input\"%s\".  Expected \"%s\".  Got \"%s\"", input, expected, errgot.Error())
			}
		}
	}
}

func TestWKPeriodDecodeAccept(t *testing.T) {

	tst := make(map[Periodint][]Periodint)

	tst[201718] = []Periodint{2017, 18}
	tst[201001] = []Periodint{2010, 1}

	t.Log("WKPeriodDecode accept tests.  \"func WKPeriodDecode(p Periodint) (Periodint, Periodint, error)\"")

	for input, expected := range tst {

		t.Logf("WKPeriodDecode accept test case.  Input \"%d\"", input)

		gotyr, gotwk, errgot := WKPeriodDecode(input)

		if errgot != nil {

			t.Errorf("WKPeriodDecode accept test case failed.  Input \"%d\", unexpected error \"%s\"", input, errgot)
		}

		if gotyr != expected[0] && gotwk != expected[1] {

			t.Errorf("WKPeriodDecode accept test case failed.  Input \"%d\", Expected \"YR: %d WK: %d\", Got \"TR: %d WK: %d\"", input, expected[0], expected[1], gotyr, gotwk)
		}
	}
}

func TestWKPeriodDecodeRejection(t *testing.T) {

	tst := []Periodint{2000, 201754, 205101, 185422, 201653}

	t.Log("WKPeriodDecode rejection tests.  \"func WKPeriodDecode(p Periodint) (Periodint, Periodint, error) {\"")

	for _, input := range tst {

		t.Logf("WKPeriodDecode rejection test case.  Input \"%d\"", input)

		_, _, errgot := WKPeriodDecode(input)

		if errgot == nil {

			t.Errorf("WKPeriodDecode rejection test failed, no error thrown.  Input \"%d\".", input)
		} else {

			expected := fmt.Sprintf("WKPeriodDecode - invalid period input - \"%d\"", input)
			if errgot.Error() != expected {

				t.Errorf("WKPeriodDecode rejection test failed.  Input \"%d\".  Expected \"%s\".  Got \"%s\"", input, expected, errgot.Error())
			}
		}
	}
}

func TestWKPeriodstrYRMN(t *testing.T) {

	tst := make(map[string][]Periodint)

	tst["201652"] = []Periodint{2016, 52}
	tst["201701"] = []Periodint{2017, 1}

	t.Log("WKPeriodstrYRWK acceptance tests. \"func WKPeriodstrYRWK(ayr, awk Periodint) (string, error)\"")

	for expected, input := range tst {

		t.Logf("WKPeriodstrYRWK acceptance test case. Input \"YR: %d WK: %d\".  Expected \"%s\"", input[0], input[1], expected)

		got, errgot := WKPeriodstrYRWK(input[0], input[1])

		if errgot != nil {

			t.Errorf("WKPeriodstrYRWK acceptance test failed. Input \"YR: %d WK: %d\".  Unexpected error \"%s\" encountered", input[0], input[1], errgot)
		}

		if got != expected {

			t.Errorf("WKPeriodstrYRWK acceptance test failed. Input \"YR: %d WK: %d\".  Expected \"%s\".  Got \"%s\"", input[0], input[1], expected, got)
		}
	}
}

func TestWKPeriodstrDT(t *testing.T) {

	tst := make(map[string]string)

	tst["2016-12-29"] = "201701"
	tst["2017-12-28"] = "201753"
	tst["2018-05-26"] = "201821"

	t.Log("WKPeriodstrDT acceptance tests. \"func WKPeriodstrDT(t time.Time) (string, error)\"")

	for input, expected := range tst {

		t.Logf("WKPeriodstrDT acceptance test case. Input \"%s\".  Expected \"%s\"", input, expected)

		dtin, errdtin := time.Parse(DTFORMAT, input)
		if errdtin != nil {

			t.Errorf("WKPeriodstrDT acceptance test case failed with bad input data - error \"%s\". Input \"%s\".  Expected \"%s\"", errdtin, input, expected)
		}

		got, errgot := WKPeriodstrDT(dtin)

		if errgot != nil {

			t.Errorf("WKPeriodstrDT acceptance test case failed - Unexpected error \"%s\". Input \"%s\".  Expected \"%s\"", errdtin, input, expected)
		}

		if got != expected {

			t.Errorf("WKPeriodstrDT acceptance test failed. Input \"%s\".  Expected \"%s\".  Got \"%s\"", input, expected, got)
		}
	}
}

func TestQTRPeriodValidAccept(t *testing.T) {

	tst := make(map[Periodint]bool)

	tst[1] = false
	tst[201813] = false
	tst[201705] = false
	tst[201704] = true
	tst[195004] = false
	tst[205101] = false
	tst[201603] = true

	t.Log("TestQTRPeriodValid acceptance tests.  \"func QTRPeriodValid(p Periodint) bool\"")

	for input, expected := range tst {

		t.Logf("TestQTRPeriodValid acceptance test case.  Input \"%d\".  Expected \"%t\".", input, expected)

		got := QTRPeriodValid(input)

		if got != expected {

			t.Errorf("TestQTRPeriodValid acceptance test failed.  Input \"%d\".  Expected \"%t\".  Got \"%t\"", input, expected, got)
		}
	}
}

func TestQTRPeriodYRMNAccept(t *testing.T) {

	tst := make(map[Periodint][]Periodint)

	tst[201601] = []Periodint{2016, 1}
	tst[201504] = []Periodint{2015, 4}
	tst[202003] = []Periodint{2020, 3}
	tst[201702] = []Periodint{2017, 2}

	t.Log("TestQTRPeriodYRMNAccept acceptance tests.  \"func QTRPeriodYRQTR(ayr, amn Periodint) (Periodint, error)\"")

	for expected, input := range tst {

		t.Logf("TestQTRPeriodYRMNAccept acceptance test case.  Input \"YR %d MN %d\".  Expected \"%d\"", input[0], input[1], expected)

		got, errgot := QTRPeriodYRQTR(input[0], input[1])

		if errgot != nil {

			t.Errorf("TestQTRPeriodYRMNAccept acceptance test failed.  Input \"YR %d MN %d\".  Expected \"%d\".  Unexpected error \"%s\"", input[0], input[1], expected, errgot)
		}

		if got != expected {

			t.Errorf("TestQTRPeriodYRMNAccept acceptance test failed.  Input \"YR %d MN %d\".  Expected \"%d\".  Got \"%d\"", input[0], input[1], expected, got)
		}
	}
}

func TestQTRPeriodYRMNReject(t *testing.T) {

	tst := [][]Periodint{
		[]Periodint{2017, 5},
		[]Periodint{1850, 3},
		[]Periodint{2051, 01},
		[]Periodint{1, 2},
	}

	t.Log("TestQTRPeriodYRMNReject rejection tests.  \"func QTRPeriodYRQTR(ayr, amn Periodint) (Periodint, error)\"")

	for _, input := range tst {

		expected := fmt.Sprintf("QTRPeriodYRQTR recieved invalid input: \"YR: %d MN: %d\"", input[0], input[1])
		t.Logf("TestQTRPeriodYRMNReject rejection test case.  Input \"YR: %d MN: %d\".  Expected \"%s\".", input[0], input[1], expected)

		_, errgot := QTRPeriodYRQTR(input[0], input[1])

		if errgot == nil || errgot.Error() != expected {

			t.Errorf("TestQTRPeriodYRMNReject rejection test failed.  Input \"YR: %d MN: %d\".  Expected \"%s\".  Got\"%s\"", input[0], input[1], expected, errgot)
		}
	}
}

func TestQtrPeriodDecodeAccept(t *testing.T) {

	tst := make(map[Periodint][2]Periodint)

	tst[201501] = [2]Periodint{2015, 1}
	tst[201702] = [2]Periodint{2017, 2}
	tst[200003] = [2]Periodint{2000, 3}
	tst[202004] = [2]Periodint{2020, 4}

	t.Log("TestQtrPeriodDecodeAccept acceptance tests.  \"func QTRPeriodDecode(p Periodint) (Periodint, Periodint, error)\"")

	for input, expected := range tst {

		t.Logf("TestQtrPeriodDecodeAccept acceptance test case.  Input \"%d\".  Expected \"YR: %d, QTR: %d\".", input, expected[0], expected[1])

		gotyr, gotqtr, errgot := QTRPeriodDecode(input)

		if errgot != nil {

			t.Errorf("TestQtrPeriodDecodeAccept acceptance test failed - unexpected error.  Input \"%d\".  Expected \"YR: %d, QTR: %d\".  Error \"%s\"", input, expected[0], expected[1], errgot)
		}

		if gotyr != expected[0] || gotqtr != expected[1] {

			t.Errorf("TestQtrPeriodDecodeAccept acceptance test failed.  Input \"%d\".  Expected \"YR: %d, QTR: %d\".  Got \"YR: %d, QTR: %d\"", input, expected[0], expected[1], gotyr, gotqtr)
		}
	}
}

func TestQtrPeriodDecodeReject(t *testing.T) {

	tst := []Periodint{2000, 201700, 201705, 205101, 176601}

	t.Log("TestQtrPeriodDecodeReject rejection tests.  \"func QTRPeriodDecode(p Periodint) (Periodint, Periodint, error)\"")

	for _, input := range tst {

		t.Logf("TestQtrPeriodDecodeReject rejection test case.  Input \"%d\"", input)

		expected := fmt.Sprintf("QTRPeriodDecode - invalid period input - \"%d\"", input)
		_, _, errgot := QTRPeriodDecode(input)

		if errgot == nil {

			t.Errorf("TestQtrPeriodDecodeReject rejection test case failed.  Input \"%d\".  No error produced", input)
		} else {

			if errgot.Error() != expected {

				t.Errorf("TestQtrPeriodDecodeReject rejection test case failed.  Input \"%d\".  Expected \"%s\".  Got \"%s\".", input, expected, errgot)
			}
		}
	}
}

func TestQtrPeriodstrYRQTR(t *testing.T) {

	tst := make(map[string][2]Periodint)

	tst["201701"] = [2]Periodint{2017, 1}
	tst["201704"] = [2]Periodint{2017, 4}

	t.Log("TestQtrPeriodstrYRQTR acceptance tests.  \"func QTRPeriodstrYRQTR(ayr, aqtr Periodint) (string, error)\"")

	for expected, input := range tst {

		t.Logf("TestQtrPeriodstrYRQTR acceptance test case.  Input \"YR: %d, QTR: %d\".  Expected \"%s\".", input[0], input[1], expected)

		got, errgot := QTRPeriodstrYRQTR(input[0], input[1])

		if errgot != nil {

			t.Errorf("TestQtrPeriodstrYRQTR acceptance test case failed - unexpected error.  Input \"YR: %d, QTR: %d\".  Expected \"%s\".  Error \"%s\"", input[0], input[1], expected, errgot)
		}

		if got != expected {

			t.Errorf("TestQtrPeriodstrYRQTR acceptance test case failed.  Input \"YR: %d, QTR: %d\".  Expected \"%s\".  Got \"%s\"", input[0], input[1], expected, got)
		}
	}
}

func TestWKPeriodsContainedAccept(t *testing.T) {

	tst := make(map[int][2]Periodint)

	tst[1] = [2]Periodint{201801, 201801}
	tst[53] = [2]Periodint{201701, 201753}
	tst[61] = [2]Periodint{201649, 201804}

	t.Log("WKPeriodsContained acceptance tests.  \"func WKPeriodsContained(astartperiod, aendperiod Periodint) (Periodint, error)\"")

	for expected, input := range tst {

		t.Logf("WKPeriodsContained acceptance tests case.  Input \"PSTART: %d PEND: %d\".  Expected \"%d\"", input[0], input[1], expected)

		got, errgot := WKPeriodsContained(input[0], input[1])
		if errgot != nil {

			t.Errorf("WKPeriodsContained acceptance test failed - unexpected error.  Input \"PSTART: %d PEND: %d\".  Expected \"%d\".  Error \"%s\"", input[0], input[1], expected, errgot)
		}

		if got != Periodint(expected) {

			t.Errorf("WKPeriodsContained acceptance test failed.  Input \"PSTART: %d PEND: %d\".  Expected \"%d\".  Got \"%d\"", input[0], input[1], expected, got)
		}
	}
}

func TestWKPeriodsContainedRejectStBiggerFin(t *testing.T) {

	t.Log("TestWKPeriodsContainedRejectStBiggerFin - WKPeriodsContained - should reject startperiod > endperiod")

	_, err := WKPeriodsContained(200612, 200501)

	if err == nil {

		t.Errorf("TestWKPeriodsContainedRejectStBiggerFin failed.  Did not return an error.")
	} else {

		expected := fmt.Sprintf("Start week period \"%d\" is bigger than end week period \"%d\".", 200612, 200501)

		if err.Error() != expected {

			t.Errorf("TestWKPeriodsContainedRejectStBiggerFin failed.  Expected error \"%s\".  Got \"%s\".", expected, err)
		}
	}
}

func TestWKPeriodsContainedInvalidInput(t *testing.T) {

	t.Log("TestWKPeriodsContainedInvalidInput - WKPeriodsContained - should reject invalid input periods")

	_, err := WKPeriodsContained(202054, 202103)
	if err == nil {

		t.Errorf("%s Failed - No error returned for incorrect start period argument.", t.Name())
	} else {

		expected := fmt.Sprintf("WKPeriodDecode - invalid period input - \"%d\"", 202054)

		if err.Error() != expected {

			t.Errorf("%s failed.  Expected error \"%s\".  Got \"%s\".", t.Name(), expected, err)
		}
	}

	_, err = WKPeriodsContained(202001, 202100)
	if err == nil {

		t.Errorf("%s Failed - No error returned for incorrect end period argument.", t.Name())
	} else {

		expected := fmt.Sprintf("WKPeriodDecode - invalid period input - \"%d\"", 202100)

		if err.Error() != expected {

			t.Errorf("%s failed.  Expected error \"%s\".  Got \"%s\".", t.Name(), expected, err)
		}
	}
}

func TestWKPeriodSliceAccept(t *testing.T) {

	t.Logf("%s - WKPeriodSlice acceptance tests.", t.Name())

	type TestDat struct {
		name     string
		inst     Periodint
		infi     Periodint
		expected []Periodint
	}

	tst := []TestDat{
		{"t01", 201801, 201801, []Periodint{201801}},
		{"t02", 201801, 201805, []Periodint{201801, 201802, 201803, 201804, 201805}},
		{"t03", 201751, 201802, []Periodint{201751, 201752, 201753, 201801, 201802}},
	}

	for _, td := range tst {

		got, errgot := WKPeriodSlice(td.inst, td.infi)
		if errgot != nil {

			t.Errorf("%s failed.  Input \"%v\" Unexpected error \"%s\" encountered.", t.Name(), td, errgot)
		}

		if len(got) != len(td.expected) {

			t.Errorf("%s failed.  Input \"%v\".  Lenght of expected \"%d\".  Length of got \"%d\".", t.Name(), td, len(td.expected), len(got))
		} else {

			for idx, elem := range td.expected {

				if elem != got[idx] {

					t.Errorf("%s - %s failed. Got element \"%d\" at idx \"%d\" does not equal expected element \"%d\".  Input \"%v\"", t.Name(), td.name, got[idx], idx, elem, td)
				}
			}
		}
	}
}

func TestWKPeriodSliceReject(t *testing.T) {

	t.Logf("%s - rejection test - start bigger than end.", t.Name())

	_, err := WKPeriodSlice(202012, 202001)

	if err == nil {

		t.Errorf("%s - rejection test - start bigger than end failed.  No error received.", t.Name())
	} else {

		expected := fmt.Sprintf("Start week period \"%d\" is bigger than end week period \"%d\".", 202012, 202001)
		if err.Error() != expected {

			t.Errorf("%s - rejection test - start bigger than end failed.  Expected error \"%s\".  Got error \"%s\".", t.Name(), expected, err)
		}
	}

	t.Logf("%s - rejection test - start period invalid.", t.Name())
	_, err = WKPeriodSlice(2020, 202001)

	if err == nil {

		t.Errorf("%s - rejection test - start period invalid failed.  No error received.", t.Name())
	} else {

		expected := fmt.Sprintf("Start period \"%d\" is invalid.", 2020)
		if err.Error() != expected {

			t.Errorf("%s - rejection test - start period invalid failed.  Expected error \"%s\".  Got error \"%s\".", t.Name(), expected, err)
		}
	}

	t.Logf("%s - rejection test - end period invalid.", t.Name())
	_, err = WKPeriodSlice(202012, 202100)

	if err == nil {

		t.Errorf("%s - rejection test - end period invalid failed.  No error received.", t.Name())
	} else {

		expected := fmt.Sprintf("End period \"%d\" is invalid.", 202100)
		if err.Error() != expected {

			t.Errorf("%s - rejection test - end period invalid failed.  Expected error \"%s\".  Got error \"%s\".", t.Name(), expected, err)
		}
	}
}

func TestWkPeriodSubtractWeeks(t *testing.T) {

	type TType struct {
		Input   Periodint
		Wks     Periodint
		Expeced Periodint
	}

	tests := []TType{
		{201802, 1, 201801},
		{202001, 53, 201852},
		{202001, 105, 201753},
		{202001, 104, 201801},
		{202001, 0, 202001},
	}

	for _, td := range tests {

		errper, got := WKPeriodSubtractWeeks(td.Input, td.Wks)

		if errper != nil {

			t.Errorf("Input: \"%d\", subtract %d wks.  Expected: \"%d\".  Unexpected error: \"%s\".", td.Input, td.Wks, td.Expeced, errper)
		}

		if got != td.Expeced {

			t.Errorf("Input: \"%d\", subtract %d wks.  Expected: \"%d\".  Got: \"%d\".", td.Input, td.Wks, td.Expeced, got)
		}
	}
}

func TestWKPeriodAddWeeks(t *testing.T) {

	type TTest struct {
		Input    Periodint
		Wks      Periodint
		Expected Periodint
	}

	tests := []TTest{

		{201801, 1, 201802},
		{201852, 1, 201901},
		{201852, 53, 202001},
		{201852, 104, 202052},
		{201852, 105, 202101},
		{201852, 0, 201852},
	}

	for _, tst := range tests {

		errper, got := WKPeriodAddWeeks(tst.Input, tst.Wks)

		if errper != nil {

			t.Errorf("Input: \"%d\", subtract %d wks.  Expected: \"%d\".  Unexpected error: \"%s\".", tst.Input, tst.Wks, tst.Expected, errper)
		}

		if got != tst.Expected {

			t.Errorf("Input: \"%d\", subtract %d wks.  Expected: \"%d\".  Got: \"%d\".", tst.Input, tst.Wks, tst.Expected, got)
		}
	}
}
