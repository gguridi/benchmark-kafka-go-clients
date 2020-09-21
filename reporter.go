package main

import (
	"encoding/json"
	"os"

	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/types"
)

type JSONReporter struct {
}

type JSON struct {
	Name         string      `json:"name"`
	Library      string      `json:"library"`
	NumMessages  int         `json:"numMessages"`
	SizeMessages int         `json:"sizeMessages`
	NumSamples   int         `json:"numberOfSamples"`
	Results      []float64   `json:"results"`
	Smallest     interface{} `json:"smallest"`
	Largest      interface{} `json:"largest"`
	Average      interface{} `json:"average"`
	StdDeviation interface{} `json:"stdDeviation"`
}

func NewJSONReporter() *JSONReporter {
	return &JSONReporter{}
}

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

func (reporter *JSONReporter) SpecSuiteWillBegin(config config.GinkgoConfigType, summary *types.SuiteSummary) {
}

func (reporter *JSONReporter) BeforeSuiteDidRun(setupSummary *types.SetupSummary) {
}

func (reporter *JSONReporter) AfterSuiteDidRun(setupSummary *types.SetupSummary) {
}

func (reporter *JSONReporter) SpecWillRun(specSummary *types.SpecSummary) {
}

func (reporter *JSONReporter) SpecSuiteDidEnd(summary *types.SuiteSummary) {
}
