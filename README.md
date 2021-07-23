# `select` and `time.After`

See https://github.com/golang/go/issues/44343

This is to test and see discrepancies for wait periods when using `select` with `<- time.After()` in the microseconds range.

`WAIT_MICROS` is the number of microseconds that is sent to `time.After()`

`WAIT_MICROS=500 go test -cpu 1 -bench . && egrep "wait|skip" select.prom.txt`

```
# HELP excess_wait 
# TYPE excess_wait summary
excess_wait{quantile="0"} 622
excess_wait{quantile="0.025"} 649
excess_wait{quantile="0.16"} 749
excess_wait{quantile="0.5"} 792
excess_wait{quantile="0.84"} 1004
excess_wait{quantile="0.95"} 3852
excess_wait{quantile="0.975"} 8833
excess_wait{quantile="1"} 146348
excess_wait_sum 1.787731e+06
excess_wait_count 1181
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
# HELP skip 
# TYPE skip summary
skip{quantile="0"} 0
skip{quantile="0.025"} 1
skip{quantile="0.16"} 1
skip{quantile="0.5"} 1
skip{quantile="0.84"} 5
skip{quantile="0.95"} 7
skip{quantile="0.975"} 11
skip{quantile="1"} 123
skip_sum 3048
skip_count 1181
# HELP wait 
# TYPE wait summary
wait{quantile="0"} 1122
wait{quantile="0.025"} 1149
wait{quantile="0.16"} 1249
wait{quantile="0.5"} 1292
wait{quantile="0.84"} 1504
wait{quantile="0.95"} 4352
wait{quantile="0.975"} 9333
wait{quantile="1"} 146848
wait_sum 2.378231e+06
wait_count 1181
```

`skip` is measured time for a `select` with a `default`
`wait` is measured time for a `select` with a `time.After()`
`excess_wait` is measured time minus expected wait time (i.e. parameter for `time.After()`)
