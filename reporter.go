package main

import (
	"encoding/json"
	"os"

	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/types"
)

// JSONReporter is a custom reporter for Ginkgo to output the test results
// into a JSON file with custom fields.
type JSONReporter struct {
}

// JSON is a struct containing the information that we will output into
// the report JSON file.
type JSON struct {
	Name         string      `json:"name"`
	Library      string      `json:"library"`
	NumMessages  int         `json:"numMessages"`
	SizeMessages int         `json:"sizeMessages"`
	NumSamples   int         `json:"numberOfSamples"`
	Results      []float64   `json:"results"`
	Smallest     interface{} `json:"smallest"`
	Largest      interface{} `json:"largest"`
	Average      interface{} `json:"average"`
	StdDeviation interface{} `json:"stdDeviation"`
}

// NewJSONReporter returns a new JSONReporter instance.
func NewJSONReporter() *JSONReporter {
	return &JSONReporter{}
}

// SpecDidComplete implements the ginkgo.Reporter interface. It runs when the tests have passed,
// allowing us to output the information into an output JSON file.
func (reporter *JSONReporter) SpecDidComplete(specSummary *types.SpecSummary) {
	switch specSummary.State {
	case types.SpecStatePassed:
		report := JSON{
			Library:      Library,
			NumMessages:  NumMessages,
			SizeMessages: MessageSize,
			NumSamples:   specSummary.NumberOfSamples,
		}
		for _, value := range specSummary.Measurements {
			report.Name = value.Name
			report.Results = value.Results
			report.Smallest = value.Smallest
			report.Largest = value.Largest
			report.Average = value.Average
			report.StdDeviation = value.StdDeviation
		}

		data, _ := json.Marshal(report)
		f, _ := os.Create("results.json")
		f.WriteString(string(data))
	}
}

// SpecSuiteWillBegin implements the ginkgo.Reporter interface.
func (reporter *JSONReporter) SpecSuiteWillBegin(config config.GinkgoConfigType, summary *types.SuiteSummary) {
}

// BeforeSuiteDidRun implements the ginkgo.Reporter interface.
func (reporter *JSONReporter) BeforeSuiteDidRun(setupSummary *types.SetupSummary) {
}

// AfterSuiteDidRun implements the ginkgo.Reporter interface.
func (reporter *JSONReporter) AfterSuiteDidRun(setupSummary *types.SetupSummary) {
}

// SpecWillRun implements the ginkgo.Reporter interface.
func (reporter *JSONReporter) SpecWillRun(specSummary *types.SpecSummary) {
}

// SpecSuiteDidEnd implements the ginkgo.Reporter interface.
func (reporter *JSONReporter) SpecSuiteDidEnd(summary *types.SuiteSummary) {
}
