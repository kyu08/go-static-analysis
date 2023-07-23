package unused_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/kyu08/unused"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, unused.Analyzer, "a")
}
