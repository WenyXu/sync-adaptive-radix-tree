# Adaptive radix tree Synchronized

This repository provides the implementation of the Adaptive Radix Tree with Optimistic Lock Coupling.

The implementation is based on following papers:

- [The Adaptive Radix Tree: ARTful Indexing for Main-Memory Databases](https://db.in.tum.de/~leis/papers/ART.pdf)
- [The ART of Practical Synchronization](https://dl.acm.org/citation.cfm?id=2933352)

## Benchmark

The benchmark tests are located in [tree_concurrent_test.go](./tree_concurrent_test.go)

`frac_x` means `0.x` read fraction. frac_0 means write-only, frac_10 means read-only.

```bash
// Skip list
BenchmarkSklReadWrite
BenchmarkSklReadWrite/frac_0
BenchmarkSklReadWrite/frac_0-8          	 2643458	       605.3 ns/op
BenchmarkSklReadWrite/frac_1
BenchmarkSklReadWrite/frac_1-8          	 2917704	       475.5 ns/op
BenchmarkSklReadWrite/frac_2
BenchmarkSklReadWrite/frac_2-8          	 3183054	       450.3 ns/op
BenchmarkSklReadWrite/frac_3
BenchmarkSklReadWrite/frac_3-8          	 3397573	       409.4 ns/op
BenchmarkSklReadWrite/frac_4
BenchmarkSklReadWrite/frac_4-8          	 3857455	       388.2 ns/op
BenchmarkSklReadWrite/frac_5
BenchmarkSklReadWrite/frac_5-8          	 4338092	       384.0 ns/op
BenchmarkSklReadWrite/frac_6
BenchmarkSklReadWrite/frac_6-8          	 5014927	       322.1 ns/op
BenchmarkSklReadWrite/frac_7
BenchmarkSklReadWrite/frac_7-8          	 6169614	       306.8 ns/op
BenchmarkSklReadWrite/frac_8
BenchmarkSklReadWrite/frac_8-8          	 7852905	       256.6 ns/op
BenchmarkSklReadWrite/frac_9
BenchmarkSklReadWrite/frac_9-8          	10119721	       182.2 ns/op
BenchmarkSklReadWrite/frac_10
BenchmarkSklReadWrite/frac_10-8         	240735708	         5.294 ns/op

// Sync-ART (our implementation)
BenchmarkArtReadWrite
BenchmarkArtReadWrite/frac_0
BenchmarkArtReadWrite/frac_0-8          	 8110496	       182.7 ns/op
BenchmarkArtReadWrite/frac_1
BenchmarkArtReadWrite/frac_1-8          	 7584531	       202.1 ns/op
BenchmarkArtReadWrite/frac_2
BenchmarkArtReadWrite/frac_2-8          	 7491115	       171.5 ns/op
BenchmarkArtReadWrite/frac_3
BenchmarkArtReadWrite/frac_3-8          	 9471957	       147.1 ns/op
BenchmarkArtReadWrite/frac_4
BenchmarkArtReadWrite/frac_4-8          	10404744	       139.5 ns/op
BenchmarkArtReadWrite/frac_5
BenchmarkArtReadWrite/frac_5-8          	11536880	       149.2 ns/op
BenchmarkArtReadWrite/frac_6
BenchmarkArtReadWrite/frac_6-8          	14883390	       106.3 ns/op
BenchmarkArtReadWrite/frac_7
BenchmarkArtReadWrite/frac_7-8          	12550774	       103.6 ns/op
BenchmarkArtReadWrite/frac_8
BenchmarkArtReadWrite/frac_8-8          	18014038	        85.53 ns/op
BenchmarkArtReadWrite/frac_9
BenchmarkArtReadWrite/frac_9-8          	28922254	        62.25 ns/op
BenchmarkArtReadWrite/frac_10
BenchmarkArtReadWrite/frac_10-8         	164754861	         7.134 ns/op

// Map with sync.Mutex
BenchmarkReadWriteMap
BenchmarkReadWriteMap/frac_0
BenchmarkReadWriteMap/frac_0-8          	 2181128	       569.5 ns/op
BenchmarkReadWriteMap/frac_1
BenchmarkReadWriteMap/frac_1-8          	 2656879	       475.6 ns/op
BenchmarkReadWriteMap/frac_2
BenchmarkReadWriteMap/frac_2-8          	 3409819	       415.9 ns/op
BenchmarkReadWriteMap/frac_3
BenchmarkReadWriteMap/frac_3-8          	 3728640	       398.0 ns/op
BenchmarkReadWriteMap/frac_4
BenchmarkReadWriteMap/frac_4-8          	 3700561	       396.7 ns/op
BenchmarkReadWriteMap/frac_5
BenchmarkReadWriteMap/frac_5-8          	 3923648	       387.9 ns/op
BenchmarkReadWriteMap/frac_6
BenchmarkReadWriteMap/frac_6-8          	 4845210	       354.7 ns/op
BenchmarkReadWriteMap/frac_7
BenchmarkReadWriteMap/frac_7-8          	 4749722	       297.4 ns/op
BenchmarkReadWriteMap/frac_8
BenchmarkReadWriteMap/frac_8-8          	 6310018	       271.9 ns/op
BenchmarkReadWriteMap/frac_9
BenchmarkReadWriteMap/frac_9-8          	 8999284	       235.0 ns/op
BenchmarkReadWriteMap/frac_10
BenchmarkReadWriteMap/frac_10-8         	 8761596	       137.3 ns/op

// Sync.Map 
BenchmarkReadWriteSyncMap
BenchmarkReadWriteSyncMap/frac_0
BenchmarkReadWriteSyncMap/frac_0-8      	 1378032	       764.8 ns/op
BenchmarkReadWriteSyncMap/frac_1
BenchmarkReadWriteSyncMap/frac_1-8      	 1484005	       749.5 ns/op
BenchmarkReadWriteSyncMap/frac_2
BenchmarkReadWriteSyncMap/frac_2-8      	 1923976	       675.5 ns/op
BenchmarkReadWriteSyncMap/frac_3
BenchmarkReadWriteSyncMap/frac_3-8      	 1898367	       666.8 ns/op
BenchmarkReadWriteSyncMap/frac_4
BenchmarkReadWriteSyncMap/frac_4-8      	 2078604	       671.0 ns/op
BenchmarkReadWriteSyncMap/frac_5
BenchmarkReadWriteSyncMap/frac_5-8      	 1937666	       689.1 ns/op
BenchmarkReadWriteSyncMap/frac_6
BenchmarkReadWriteSyncMap/frac_6-8      	 1850360	       655.9 ns/op
BenchmarkReadWriteSyncMap/frac_7
BenchmarkReadWriteSyncMap/frac_7-8      	 1910458	       689.3 ns/op
BenchmarkReadWriteSyncMap/frac_8
BenchmarkReadWriteSyncMap/frac_8-8      	 1982641	       672.2 ns/op
BenchmarkReadWriteSyncMap/frac_9
BenchmarkReadWriteSyncMap/frac_9-8      	 2560873	       521.2 ns/op
BenchmarkReadWriteSyncMap/frac_10
BenchmarkReadWriteSyncMap/frac_10-8     	149435343	         7.603 ns/op

// Sync-ART (our implementation)
BenchmarkArtConcurrentInsert
BenchmarkArtConcurrentInsert-8          	 6504886	       171.3 ns/op

// Sync-ART (another implementation)
BenchmarkAnotherArtConcurrentInsert
BenchmarkAnotherArtConcurrentInsert-8   	 7657372	       195.6 ns/op

// Btree 
BenchmarkBtreeConcurrentInsert
BenchmarkBtreeConcurrentInsert-8        	 1502853	       993.1 ns/op
```