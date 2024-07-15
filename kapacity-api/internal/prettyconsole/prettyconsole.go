package prettyconsole

import (
  "github.com/fatih/color"
	"fmt"
	"time"
)

// FgRed
// FgGreen
// FgYellow
// FgBlue
// FgMagenta
// FgCyan
// FgWhite

func WARN(inputStr string) {
	yellow := color.New(color.FgYellow).SprintFunc()
	fmt.Println(yellow("[WARN]:") + CurrentTime() + "[",inputStr,"]")
}

func DEBUG(inputStr string) {
	magenta := color.New(color.FgMagenta).SprintFunc()
	fmt.Println(magenta("[DEBUG]:") + CurrentTime() + "[",inputStr,"]")
}

func INFO(inputStr string) {
	green := color.New(color.FgGreen).SprintFunc()
	fmt.Println(green("[INFO]:") + CurrentTime() + "[",inputStr,"]")
}

func ERROR(inputStr string) {
	red := color.New(color.FgRed).SprintFunc()
	fmt.Println(red("[ERROR]:") + CurrentTime() + "[",inputStr,"]")
}

func LINT(inputStr string) {
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Println(cyan("[LINT]:") + CurrentTime() + "[",inputStr,"]")
}

func CurrentTime() string {
	currentTime := time.Now().UTC().Format(time.RFC3339)
	currentTime = "[" + currentTime + "]"
	return currentTime
}