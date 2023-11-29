package html_handlers

import types "github.com/unexpectedtokens/api-tester/common_types"

type testCasePageData struct {
	Testcases []types.TestCase
}

type testReportsPageData struct {
	TestReports []types.TestReport
}

type testReportDetailpageData struct {
	TestReport      types.TestReport
	TestcaseResults []types.TestCaseResult
}
