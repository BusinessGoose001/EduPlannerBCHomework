package digitcount

import (
	"fmt"
	"testing"
)

func TestCountDigitOccurrencesInSeries(test *testing.T) {

	for _, testcase := range testcases {
		test.Run(fmt.Sprintf("start: %d end: %d increase: %d seriesCheckType:%d searchKey:%d",
			testcase.start, testcase.end, testcase.increase, testcase.seriesCheckType, testcase.searchKey), func(test *testing.T) {
			count, err := CountDigitOccurrencesInSeries(
				testcase.start, testcase.end, testcase.increase, testcase.seriesCheckType, testcase.searchKey)
			if err != testcase.err {
				test.Errorf("unexpected error: %s \nexpected: %s", err, testcase.err)
			} else { // we caught the error as expected
				return
			}

			if count != testcase.result {
				test.Errorf("incorrect count: %d", count)
			}
		},
		)
	}
}
