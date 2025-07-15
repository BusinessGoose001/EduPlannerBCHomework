package digitcount

import (
	"fmt"
	"math"
)

var AcceptedTypes = map[int]bool{1: true, 2: true, 3: true}

var errors = map[string]error{
	"invalid_series":       fmt.Errorf("Invalid series. If start greater than end, the increment must be negative"),
	"invalid_series_check": fmt.Errorf("The series check is not an accepted value. Can be one of %v", AcceptedTypes),
}

func CountDigitOccurrencesInSeries(start int, end int, increase int, seriesCheckType int, searchKey int) (int, error) {

	if !AcceptedTypes[seriesCheckType] {
		return -1, errors["invalid_series_check"]
	}

	if start > end {
		if increase > 0 {
			return -1, errors["invalid_series"]
		} else {
			start, end, increase = end, start, -1*increase
		}
	}

	if seriesCheckType == 3 && increase%2 == 0 && start%2 == 0 {
		//	if looking for odd and start is even and increase is even then all series will be even and nothing can be checked so return 0.
		return 0, nil
	} else if seriesCheckType == 2 && increase%2 == 0 && start%2 != 0 {
		//	if looking for even and start is odd and increase is even then all series will be odd and nothing can be checked so return 0.
		return 0, nil
	}

	count := 0
	for index := start; index <= end; index += increase {
		if shouldAnalyse(index, seriesCheckType) {
			value := int(math.Abs(float64(index))) // pacify negatives

			if value == 0 && searchKey > 0 {
				count++
			} else {
				for value > 0 {
					digit := value % 10
					if digit == searchKey {
						count++
					}
					//println("value:", value, "digit:", digit, "searchKey:", searchKey, "count:", count)
					value = value / 10
				}
			}
		}
	}
	return count, nil
}

func shouldAnalyse(index int, seriesCheckType int) bool {
	switch seriesCheckType {
	case 1: // all elements
		return true // all elements
	case 2: // even elements
		return index%2 == 0
	case 3: // odd elements
		return index%2 != 0
	default:
		return false
	}
}
