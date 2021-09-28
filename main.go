package main

import (
	"bufio"
	"flag"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
)

const (
	INFO             = "INFO"
	DEBUG            = "DEBUG"
	WARNING          = "WARNING"
	ERROR            = "ERROR"
	TRACE            = "TRACE"
	DEFAULT_LOG_NAME = "example.log"
)

func initLogLevelCountMap(logLevelCount map[string]int) {
	logLevelCount[INFO] = 0
	logLevelCount[DEBUG] = 0
	logLevelCount[WARNING] = 0
	logLevelCount[ERROR] = 0
	logLevelCount[TRACE] = 0
}

func updateLogLevelCount(logLevelCount map[string]int, line string) {
	for key := range logLevelCount {
		if strings.Contains(line, key) {
			count := logLevelCount[key]
			count++
			logLevelCount[key] = count
			break
		}
	}
}

func main() {
	logLevelCount := make(map[string]int)
	initLogLevelCountMap(logLevelCount)

	logFile := flag.String("logFile", DEFAULT_LOG_NAME, "The log file path to analyze")

	flag.Parse()

	file, err := os.Open(*logFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		updateLogLevelCount(logLevelCount, line)
	}

	// log output to console
	// for key, value := range logLevelCount {
	// 	fmt.Println(key, value)
	// }

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"LogLevel", "Count"})
	for key, value := range logLevelCount {
		t.AppendRow([]interface{}{key, value})
		t.AppendSeparator()
	}
	t.Render()

}
