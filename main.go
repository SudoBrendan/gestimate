package main

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
)

var (
	// command line args
	best   string
	likely string
	worst  string
)

const dateFormatStr string = "2006/01/02"

// 3-point estimation (PERT)
func pertEstimate(best, likely, worst time.Time) float64 {
	// Convert times to durations
	a := best.Unix()
	m := likely.Unix()
	b := worst.Unix()

	// PERT formula (a + 4m + b) / 6
	// Ref https://en.wikipedia.org/wiki/Three-point_estimation
	mean := (float64(a) + 4*float64(m) + float64(b)) / 6
	return mean
}

// Standard deviation for PERT
func pertStdDev(best, worst time.Time) float64 {
	a := best.Unix()
	b := worst.Unix()

	// Standard deviation formula: (b - a) / 6
	// Ref https://en.wikipedia.org/wiki/Three-point_estimation
	return (float64(b) - float64(a)) / 6
}

// Generate confidence intervals
func confidenceIntervals(mean, stddev float64) map[string]string {
	intervals := make(map[string]string)

	// Confidence intervals based on standard deviation
	// Ref https://en.wikipedia.org/wiki/Three-point_estimation
	conf68 := mean + stddev
	conf90 := mean + 1.645*stddev
	conf95 := mean + 2*stddev

	// Convert back to date format
	// NOTE: we effectively truncate hours, so you need to always assume
	//   end-of-day delivery
	intervals["68%"] = time.Unix(int64(conf68), 0).Format(dateFormatStr)
	intervals["90%"] = time.Unix(int64(conf90), 0).Format(dateFormatStr)
	intervals["95%"] = time.Unix(int64(conf95), 0).Format(dateFormatStr)
	return intervals
}

// Parse date string in the format YYYY/MM/DD
func parseDate(dateStr string) (time.Time, error) {
	return time.Parse(dateFormatStr, dateStr)
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "gestimate",
		Short: "Gestimate is a CLI tool for 3-point estimation with confidence intervals",
		Run: func(cmd *cobra.Command, args []string) {
			// Parse args
			bestDate, err := parseDate(best)
			if err != nil {
				log.Fatalf("Invalid best date format: %v", err)
			}
			likelyDate, err := parseDate(likely)
			if err != nil {
				log.Fatalf("Invalid likely date format: %v", err)
			}
			worstDate, err := parseDate(worst)
			if err != nil {
				log.Fatalf("Invalid worst date format: %v", err)
			}

			// Calculate PERT mean and standard deviation
			mean := pertEstimate(bestDate, likelyDate, worstDate)
			stddev := pertStdDev(bestDate, worstDate)

			// Generate confidence intervals
			intervals := confidenceIntervals(mean, stddev)

			// Display the results in a table
			var reset = "\033[0m"
			var green = "\033[32m"
			fmt.Println("Confidence Interval Table")
			fmt.Println("--------------------------")
			fmt.Printf("68%% Confidence: %s\n", intervals["68%"])
			fmt.Printf("90%% Confidence: %s\n", intervals["90%"])
			fmt.Printf(green+"95%% Confidence: %s\n"+reset, intervals["95%"])
		},
	}

	// Define the flags
	rootCmd.Flags().StringVarP(&best, "best", "b", "", "Best case date (YYYY/MM/DD)")
	rootCmd.MarkFlagRequired("best")
	rootCmd.Flags().StringVarP(&likely, "likely", "l", "", "Most likely case date (YYYY/MM/DD)")
	rootCmd.MarkFlagRequired("likely")
	rootCmd.Flags().StringVarP(&worst, "worst", "w", "", "Worst case date (YYYY/MM/DD)")
	rootCmd.MarkFlagRequired("worst")

	// Run our command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
