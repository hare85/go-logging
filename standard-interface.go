package logging

// Logger mimics golang's standard Logger as an interface.

// Fatal is equivalent to Print() followed by a call to os.Exit() with a non-zero exit code.
func (l *Logger) Fatal(args ...interface{}) {
	l.Logger.Fatal(args...)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit() with a non-zero exit code.
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.Logger.Fatalf(format, args...)
}

// Fatalln is equivalent to Println() followed by a call to os.Exit()) with a non-zero exit code.
func (l *Logger) Fatalln(args ...interface{}) {
	l.Logger.Fatalln(args...)
}

// Print prints to the logger. Arguments are handled in the manner of fmt.Print.
func (l *Logger) Print(args ...interface{}) {
	l.Logger.Print(args...)
}

// Printf prints to the logger. Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Printf(format string, args ...interface{}) {
	l.Logger.Printf(format, args...)
}

// Println prints to the logger. Arguments are handled in the manner of fmt.Println.
func (l *Logger) Println(args ...interface{}) {
	l.Logger.Println(args...)
}
