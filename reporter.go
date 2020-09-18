package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/types"
)

type JSONReporter struct {
}

type JSON struct {
	Library      string                   `json:"library"`
	NumMessages  int                      `json:"numMessages"`
	SizeMessages int                      `json:"sizeMessages`
	NumSamples   int                      `json:"numberOfSamples"`
	Measurements []map[string]interface{} `json:"measurements"`
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
			Measurements: []map[string]interface{}{},
		}
		for _, value := range specSummary.Measurements {
			fmt.Printf("SpecDidComplete: %+v", value)
			measurement := map[string]interface{}{
				"name":         value.Name,
				"results":      value.Results,
				"smallest":     value.Smallest,
				"largest":      value.Largest,
				"average":      value.Average,
				"stdDeviation": value.StdDeviation,
			}
			report.Measurements = append(report.Measurements, measurement)
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
