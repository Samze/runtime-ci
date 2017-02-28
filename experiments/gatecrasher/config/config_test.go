package config_test

import (
	"os"

	"github.com/cloudfoundry/runtime-ci/experiments/gatecrasher/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	Context("when no relevant env vars are set", func() {
		BeforeEach(func() {
			Expect(os.Unsetenv("TARGET")).NotTo(HaveOccurred())
			Expect(os.Unsetenv("POLL_INTERVAL_IN_SECONDS")).NotTo(HaveOccurred())
			Expect(os.Unsetenv("TOTAL_NUMBER_OF_REQUESTS")).NotTo(HaveOccurred())
		})
		It("generates a config struct from defaults", func() {
			expectedConfig := config.Config{
				Target: "example.com",
				Poll_interval_in_seconds: 1,
				Total_number_of_requests: 10,
			}
			Expect(config.Load()).To(Equal(expectedConfig))
		})
	})
	Context("when all relevant env vars are set", func() {
		BeforeEach(func() {
			os.Setenv("TARGET", "myurl")
			os.Setenv("POLL_INTERVAL_IN_SECONDS", "10000")
			os.Setenv("TOTAL_NUMBER_OF_REQUESTS", "10000")
		})
		It("generates a config struct from env vars", func() {
			expectedConfig := config.Config{
				Target: "myurl",
				Poll_interval_in_seconds: 10000,
				Total_number_of_requests: 10000,
			}
			Expect(config.Load()).To(Equal(expectedConfig))
		})
	})
})