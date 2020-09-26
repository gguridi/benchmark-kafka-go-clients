package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func TestBenchmark(t *testing.T) {
	InitConfig()
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t, "Benchmark suite", []Reporter{NewJSONReporter()})
}

// Initialises the configuration for this benchmark.
func InitConfig() {
	log.Debug("Initialising viper config")
	viper.Set("kafka.brokers", Brokers)
	viper.Set("kafka.topic", Topic)
}
