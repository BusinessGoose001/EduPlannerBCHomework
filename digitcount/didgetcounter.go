package digitcount

import "math"

var AcceptedTypes = map[int]bool{1: true, 2: true, 3: true}

func CountDigitOccurrencesInSeries(start int, end int, increase int, seriesCheckType int, searchKey int) int {

	if !AcceptedTypes[seriesCheckType] {
		return -1
	}

	if start > end {
		if increase > 0 {
			return -1
		} else {
			start, end, increase = end, start, -1*increase
		}
	}

	occurrences := 0
	for index := start; index <= end; index += increase {
		if shouldAnalyse(index, seriesCheckType) {
			value := int(math.Abs(float64(index))) // pacify negatives
			for value > 0 {
				digit := value % 10
				if digit == searchKey {
					occurrences++
				}
				value = value / 10
			}
		}
	}
	return occurrences
}

func shouldAnalyse(index int, Type int) bool {
	switch Type {
	case 1: // all elements
		return true
	case 2: // even elements
		return index%2 == 0
	case 3: // odd elements
		return index%2 != 0
	default:
		return false
	}
}
