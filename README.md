# `select` and `time.After`

This is to test and see discrepancies for wait periods when using `select` with `<- time.After()` in the microseconds range.

`WAIT_MICROS` is the number of microseconds that is sent to `time.After()`

`WAIT_MICROS=500 go test -cpu 1 -bench . && egrep "wait|skip" select.prom.txt`
