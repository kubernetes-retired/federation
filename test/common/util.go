package common

// TestLogger defines operations common across different types of testing
type TestLogger interface {
	Fatalf(format string, args ...interface{})
	Fatal(msg string)
	Logf(format string, args ...interface{})
}
