package verbose

import (
	"fmt"
	"strconv"
)

type VerbosityLevel int

const (
	LevelInfo   VerbosityLevel = 0
	LevelDebug1                = 1
	LevelDebug2                = 2
	LevelDebug3                = 3
	LevelOk                    = 4
	LevelWarn                  = 5
	LevelFail                  = 6

	// ANSI escape color codes. See more colors here
	// https://en.wikipedia.org/wiki/ANSI_escape_code#Colors
	colorInfo   = 37
	colorOk     = 32
	colorWarn   = 33
	colorFail   = 31
	colorDebug1 = 250
	colorDebug2 = 2
	colorDebug3 = 90

	resetCode = "\033[0m"
)

type Verbose struct {
	debugLevel VerbosityLevel
	colorMap   []int
}

// The constructor.
func NewVerbose(debugLevel VerbosityLevel) *Verbose {
	v := Verbose{
		debugLevel,
		[]int{colorInfo, colorDebug1, colorDebug2, colorDebug3, colorOk, colorWarn, colorFail},
	}
	return &v
}

// Prints a message according its verbosity level.
func (v *Verbose) verbose(level VerbosityLevel, value string) {
	prefix := ""
	if level >= LevelDebug1 && level <= LevelDebug3 {
		if level > v.debugLevel {
			return
		}
		prefix = "debug" + strconv.Itoa(int(level)) + ": "
	}
	color := v.colorMap[level]
	fmt.Println(fmt.Sprintf("\033[0;%d;49m%s", color, prefix) + value + resetCode)
}

// Print simple info message.
func (v *Verbose) Info(value string) {
	v.verbose(LevelInfo, value)
}

// Print OK message.
func (v *Verbose) Ok(value string) {
	v.verbose(LevelOk, value)
}

// Print warning message.
func (v *Verbose) Warning(value string) {
	v.verbose(LevelWarn, value)
}

// Print fail message.
func (v *Verbose) Fail(value string) {
	v.verbose(LevelFail, value)
}

// Print debug message (level 1).
func (v *Verbose) Debug1(value string) {
	v.verbose(LevelDebug1, value)
}

// Print debug message (level 2).
func (v *Verbose) Debug2(value string) {
	v.verbose(LevelDebug2, value)
}

// Print debug message (level 3).
func (v *Verbose) Debug3(value string) {
	v.verbose(LevelDebug3, value)
}
