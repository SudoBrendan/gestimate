package main

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
)

var (
	best   string
	likely string
	worst  string
)

// 3-point estimation (PERT)
func pertEstimate(best, likely, worst time.Time) float64 {
	// Convert times to durations
	bestDur := best.Unix()
	likelyDur := likely.Unix()
	worstDur := worst.Unix()

	// PERT formula (B + 4L + W) / 6
	mean := (float64(bestDur) + 4*float64(likelyDur) + float64(worstDur)) / 6
	return mean
}

// Standard deviation for PERT
func pertStdDev(best, worst time.Time) float64 {
	bestDur := best.Unix()
	worstDur := worst.Unix()

	// Standard deviation formula: (W - B) / 6
	return (float64(worstDur) - float64(bestDur)) / 6
}

// Generate confidence intervals
func confidenceIntervals(mean, stddev float64) map[string]string {
	intervals := make(map[string]string)
	// Confidence intervals based on standard deviation
	conf68 := mean + stddev
	conf90 := mean + 1.645*stddev
	conf95 := mean + 1.96*stddev

	// Convert back to date format
	intervals["68%"] = time.Unix(int64(conf68), 0).Format("2006/01/02")
	intervals["90%"] = time.Unix(int64(conf90), 0).Format("2006/01/02")
	intervals["95%"] = time.Unix(int64(conf95), 0).Format("2006/01/02")
	return intervals
}

// Parse date string in the format YYYY/MM/DD
func parseDate(dateStr string) (time.Time, error) {
	layout := "2006/01/02"
	return time.Parse(layout, dateStr)
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "gestimate",
		Short: "Gestimate is a CLI tool for 3-point estimation with confidence intervals",
		Run: func(cmd *cobra.Command, args []string) {
			// Parse the dates
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
			fmt.Println("Confidence Interval Table")
			fmt.Println("------------------------")
			fmt.Printf("68%% Confidence: %s\n", intervals["68%"])
			fmt.Printf("90%% Confidence: %s\n", intervals["90%"])
			fmt.Printf("95%% Confidence: %s\n", intervals["95%"])
		},
	}

	// Define the flags
	rootCmd.Flags().StringVarP(&best, "best", "b", "", "Best case date (YYYY/MM/DD)")
	rootCmd.Flags().StringVarP(&likely, "likely", "l", "", "Most likely case date (YYYY/MM/DD)")
	rootCmd.Flags().StringVarP(&worst, "worst", "w", "", "Worst case date (YYYY/MM/DD)")

	// Make flags required
	rootCmd.MarkFlagRequired("best")
	rootCmd.MarkFlagRequired("likely")
	rootCmd.MarkFlagRequired("worst")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
