# go-microbench
Microbenchmarks for the go programming language

## Benchmarks
Benchmarks are found in their respective `_test` file, e.g. `select` benchmarks are in `select_test.go`


### `main` benchmarks
Benchmarks for reference, or stuff that didn't fit somewhere else:

- `BenchmarkEmpty`: Benchmarks the empty for loop used in all other benchmarks

### `select` benchmarks
Benchmarks for the `select` statement:

- `BenchmarkSelectDefaultAfterRead[One-Nine]`: Selects the default case after trying to read from `[One-Nine]` unbuffered channels
- `BenchmarkSelectDefaultBeforeRead[One-Three]`: Selects the default case after trying to read from `[One-Three]` unbuffered channels, but the default case is put before the other cases in the source
- `BenchmarkSelectDefaultAfterRead[One-Three]Buffered`: Selects the default case after trying to read from `[One-Three]` buffered (1) channels
- `BenchmarkSelectDefaultAfterWrite[One-Three]`: Selects the default case after trying to write to `[One-Three]` unbuffered channels
- `BenchmarkSelectDefaultBeforeWrite[One-Three]`: Selects the default case after trying to write to `[One-Three]` unbuffered channels, but the default case is put before the other cases in the source
- `BenchmarkSelectDefaultAfterWrite[One-Three]Buffered`: Selects the default case after trying to write to `[One-Three]` buffered (1) channels

## Results
For results, see `RESULTS.md`.

The results may change significantly (e.g. when I run them on a PC vs. when I run them on a laptop), but the interesting relations between results should remain the same.