package digitcount

type digitcountTest struct {
	start           int
	end             int
	increase        int
	seriesCheckType int
	searchKey       int
	result          int
	err             error
}

var testcases = []digitcountTest{
	{ // all
		start:           568,
		end:             1247,
		increase:        32,
		seriesCheckType: 1,
		searchKey:       4,
		result:          7,
		err:             nil,
	}, { //positive only
		start:           568,
		end:             1247,
		increase:        32,
		seriesCheckType: 2,
		searchKey:       4,
		result:          7,
		err:             nil,
	}, { // negative only
		start:           568,
		end:             1247,
		increase:        32,
		seriesCheckType: 3,
		searchKey:       4,
		result:          0,
		err:             nil,
	},
	{ // invalid series setup
		start:           13,
		end:             1,
		increase:        4,
		seriesCheckType: 1,
		searchKey:       4,
		result:          -1,
		err:             errors["invalid_series"],
	},
	{ //invalid series number
		start:           13,
		end:             1,
		increase:        4,
		seriesCheckType: 6,
		searchKey:       4,
		result:          -1,
		err:             errors["invalid_series_check"],
	},
}
