Error handler
=============

This package is a try-catch implementation for Go.

## Usage
```go
TryCatch{
	Try: func() {
		// some code that falls to panic ...
	},
	Catch(e Exception) {
		// Error handling.
		// Info about panic will available in variable 'e'.
	}
}
```
