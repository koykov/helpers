package verbose

import (
	"fmt"
)

type VerbosityLevel int

const (
	LevelInfo   VerbosityLevel = 0
	LevelDebug1 VerbosityLevel = 1
	LevelDebug2 VerbosityLevel = 2
	LevelDebug3 VerbosityLevel = 3
	LevelOk     VerbosityLevel = 4
	LevelWarn   VerbosityLevel = 5
	LevelFail   VerbosityLevel = 6

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
func (v *Verbose) verbose(level VerbosityLevel, a ...interface{}) {
	debug := ""
	if level >= LevelDebug1 && level <= LevelDebug3 {
		if level > v.debugLevel {
			return
		}
		debug = fmt.Sprintf("debug%d: ", level)
	}
	color := v.colorMap[level]
	prefix := fmt.Sprintf("\033[0;%d;49m%s", color, debug)
	fmt.Println(prefix + fmt.Sprint(a...) + resetCode)
}

// Print simple info message.
func (v *Verbose) Info(a ...interface{}) {
	v.verbose(LevelInfo, a...)
}

// Print OK message.
func (v *Verbose) Ok(a ...interface{}) {
	v.verbose(LevelOk, a...)
}

// Print warning message.
func (v *Verbose) Warning(a ...interface{}) {
	v.verbose(LevelWarn, a...)
}

// Print fail message.
func (v *Verbose) Fail(a ...interface{}) {
	v.verbose(LevelFail, a...)
}

// Print debug message (level 1).
func (v *Verbose) Debug1(a ...interface{}) {
	v.verbose(LevelDebug1, a...)
}

// Print debug message (level 2).
func (v *Verbose) Debug2(a ...interface{}) {
	v.verbose(LevelDebug2, a...)
}

// Print debug message (level 3).
func (v *Verbose) Debug3(a ...interface{}) {
	v.verbose(LevelDebug3, a...)
}

// Prints a message according its verbosity level.
func (v *Verbose) verbosef(format string, level VerbosityLevel, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	v.verbose(level, s)
}

// Info with formatting support.
func (v *Verbose) Infof(format string, a ...interface{}) {
	v.verbosef(format, LevelInfo, a)
}

// Ok() with formatting support.
func (v *Verbose) Okf(format string, a ...interface{}) {
	v.verbosef(format, LevelOk, a)
}

// Warning() with formatting support.
func (v *Verbose) Warningf(format string, a ...interface{}) {
	v.verbosef(format, LevelWarn, a)
}

// Fail() with formatting support.
func (v *Verbose) Failf(format string, a ...interface{}) {
	v.verbosef(format, LevelFail, a)
}

// Debug1() with formatting support.
func (v *Verbose) Debug1f(format string, a ...interface{}) {
	v.verbosef(format, LevelDebug1, a)
}

// Debug2() with formatting support.
func (v *Verbose) Debug2f(format string, a ...interface{}) {
	v.verbosef(format, LevelDebug2, a)
}

// Debug3() with formatting support.
func (v *Verbose) Debug3f(format string, a ...interface{}) {
	v.verbosef(format, LevelDebug3, a)
}
