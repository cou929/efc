package efc_test

import (
	"testing"

	"github.com/cou929/efc"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, efc.Analyzer, "a")
}