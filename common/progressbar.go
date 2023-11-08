// Package common provides utilities that can be shared across different applications.
package common

import (
	"github.com/schollz/progressbar/v3"
)

// ProgressBar creates and returns a new progress bar with the specified size and custom options.
// It enables color codes, displays the progress in bytes, shows the count,
// predicts the remaining time, sets a description, and defines a custom theme for the progress bar.
func ProgressBar(size int) *progressbar.ProgressBar {
	bar := progressbar.NewOptions(size,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionShowCount(),
		progressbar.OptionSetPredictTime(true),
		progressbar.OptionSetDescription("[cyan]Parsing JSON Data...[reset]"),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))

	return bar
}
