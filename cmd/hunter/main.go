package main

import (
	"bashlui/hunter/internal/scraper"
	"fmt"
	"os"

	"github.com/fatih/color"
)

const version = "1.0.0"

var rainbow = []*color.Color{
	color.New(color.FgHiRed),
	color.New(color.FgHiYellow),
	color.New(color.FgHiGreen),
	color.New(color.FgHiCyan),
	color.New(color.FgHiBlue),
	color.New(color.FgHiMagenta),
}

func rainbowPrint(s string) {
	chars := []rune(s)
	for i, ch := range chars {
		rainbow[i%len(rainbow)].Printf("%c", ch)
	}
	fmt.Println()
}

func banner() {
	lines := []string{
		"	.__                  __                ",
		"	|  |__  __ __  _____/  |_  ___________ ",
		"	|  |  \\|  |  \\/    \\   __\\/ __ \\_  __ \\",
		"	|   Y  \\  |  /   |  \\  | \\  ___/|  | \\/",
		"	|___|  /____/|___|  /__|  \\___  >__|   ",
		"	    \\/           \\/          \\/        ",
	}
	for _, line := range lines {
		rainbowPrint(line)
	}
	// Subtitle in white + bold
	color.New(color.FgHiWhite, color.Bold).
		Printf("\n   Image Hunter CLI v%s - scrape image URLs from a page\n\n by toño\n", version)
}

func main() {
	banner()
	if len(os.Args) < 2 {
		fmt.Println("Usage: hunter <url>")
		os.Exit(1)
	}
	url := os.Args[1]
	imgURLs, err := scraper.ScrapeImageURLs(url)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	for _, imgURL := range imgURLs {
		fmt.Println(imgURL)
	}
}
