package output

import tm "github.com/buger/goterm"

func Error(message string) {
	tm.Println(tm.Color(message, tm.RED))
	tm.Flush()
}
