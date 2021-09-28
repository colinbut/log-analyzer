# Log Analyzer

This is a tool written in Go that takes in a log file and outputs a table with count of each log level occurences.

e.g.

Here is an example of an output table:


```bash
+----------+-------+
| LOGLEVEL | COUNT |
+----------+-------+
| INFO     |    41 |
+----------+-------+
| DEBUG    |     0 |
+----------+-------+
| WARNING  |     4 |
+----------+-------+
| ERROR    |     0 |
+----------+-------+
| TRACE    |     7 |
+----------+-------+
```

## Usage

```bash
go run main.go
```

## Author

Colin But
