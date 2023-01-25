package tests

type TestCase struct {
	Name          string
	PreRequisites func()
	Assert        func()

	SubTests []TestCase
}

type Suite interface {
	Run(name string, subtest func()) bool
}

func RunTestCases(s Suite, testCases []TestCase) {
	if len(testCases) == 0 {
		return
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.Name, func() {
			if tc.PreRequisites != nil {
				tc.PreRequisites()
			}

			if tc.Assert != nil {
				tc.Assert()
			}

			RunTestCases(s, tc.SubTests)
		})
	}
}
