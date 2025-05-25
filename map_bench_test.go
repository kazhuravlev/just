package just_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/kazhuravlev/just"
)

// Helper functions to create test maps of various sizes
func createIntMap(size int) map[int]int {
	m := make(map[int]int, size)
	for i := 0; i < size; i++ {
		m[i] = i * 2
	}
	return m
}

func createStringMap(size int) map[string]string {
	m := make(map[string]string, size)
	for i := 0; i < size; i++ {
		m[strconv.Itoa(i)] = fmt.Sprintf("value_%d", i)
	}
	return m
}

// BenchmarkMapMerge benchmarks the MapMerge function
func BenchmarkMapMerge(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			m1 := createIntMap(size)
			m2 := createIntMap(size)
			// Half overlapping keys
			for i := size / 2; i < size+size/2; i++ {
				m2[i] = i * 3
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.MapMerge(m1, m2, func(k, v1, v2 int) int {
					return v1 + v2
				})
			}
		})
	}
}

// BenchmarkMapFilter benchmarks the MapFilter function
func BenchmarkMapFilter(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			m := createIntMap(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.MapFilter(m, func(k, v int) bool {
					return v%2 == 0 && k < size/2
				})
			}
		})
	}
}

// BenchmarkMapFilterKeys benchmarks the MapFilterKeys function
func BenchmarkMapFilterKeys(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			m := createIntMap(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.MapFilterKeys(m, func(k int) bool {
					return k%2 == 0
				})
			}
		})
	}
}

// BenchmarkMapFilterValues benchmarks the MapFilterValues function
func BenchmarkMapFilterValues(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			m := createIntMap(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.MapFilterValues(m, func(v int) bool {
					return v > size/2
				})
			}
		})
	}
}

// BenchmarkMapGetKeys benchmarks the MapGetKeys function
func BenchmarkMapGetKeys(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			m := createIntMap(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.MapGetKeys(m)
			}
		})
	}
}

// BenchmarkMapGetValues benchmarks the MapGetValues function
func BenchmarkMapGetValues(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			m := createIntMap(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.MapGetValues(m)
			}
		})
	}
}

// BenchmarkMapPairs benchmarks the MapPairs function
func BenchmarkMapPairs(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			m := createIntMap(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.MapPairs(m)
			}
		})
	}
}

// BenchmarkMapDefaults benchmarks the MapDefaults function
func BenchmarkMapDefaults(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			m := createIntMap(size / 2) // Half filled
			defaults := createIntMap(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.MapDefaults(m, defaults)
			}
		})
	}
}

// BenchmarkMapCopy benchmarks the MapCopy function
func BenchmarkMapCopy(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			m := createIntMap(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.MapCopy(m)
			}
		})
	}
}

// BenchmarkMapMap benchmarks the MapMap function
func BenchmarkMapMap(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			m := createIntMap(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.MapMap(m, func(k, v int) (string, string) {
					return strconv.Itoa(k), strconv.Itoa(v)
				})
			}
		})
	}
}

// BenchmarkMapMapErr benchmarks the MapMapErr function
func BenchmarkMapMapErr(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d_success", size), func(b *testing.B) {
			m := createIntMap(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = just.MapMapErr(m, func(k, v int) (string, string, error) {
					return strconv.Itoa(k), strconv.Itoa(v), nil
				})
			}
		})

		b.Run(fmt.Sprintf("size_%d_error_early", size), func(b *testing.B) {
			m := createIntMap(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = just.MapMapErr(m, func(k, v int) (string, string, error) {
					if k == 0 {
						return "", "", errors.New("error")
					}
					return strconv.Itoa(k), strconv.Itoa(v), nil
				})
			}
		})
	}
}

// BenchmarkMapContainsKey benchmarks the MapContainsKey function
func BenchmarkMapContainsKey(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d_exists", size), func(b *testing.B) {
			m := createIntMap(size)
			key := size / 2

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.MapContainsKey(m, key)
			}
		})

		b.Run(fmt.Sprintf("size_%d_not_exists", size), func(b *testing.B) {
			m := createIntMap(size)
			key := size + 1

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.MapContainsKey(m, key)
			}
		})
	}
}

// BenchmarkMapContainsKeysAny benchmarks the MapContainsKeysAny function
func BenchmarkMapContainsKeysAny(b *testing.B) {
	sizes := []int{10, 1000}
	keyCounts := []int{1, 10, 100}

	for _, size := range sizes {
		for _, keyCount := range keyCounts {
			if keyCount > size {
				continue
			}

			b.Run(fmt.Sprintf("size_%d_keys_%d_found", size, keyCount), func(b *testing.B) {
				m := createIntMap(size)
				keys := make([]int, keyCount)
				for i := 0; i < keyCount; i++ {
					keys[i] = i
				}

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					_ = just.MapContainsKeysAny(m, keys)
				}
			})

			b.Run(fmt.Sprintf("size_%d_keys_%d_not_found", size, keyCount), func(b *testing.B) {
				m := createIntMap(size)
				keys := make([]int, keyCount)
				for i := 0; i < keyCount; i++ {
					keys[i] = size + i + 1
				}

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					_ = just.MapContainsKeysAny(m, keys)
				}
			})
		}
	}
}

// BenchmarkMapContainsKeysAll benchmarks the MapContainsKeysAll function
func BenchmarkMapContainsKeysAll(b *testing.B) {
	sizes := []int{10, 1000}
	keyCounts := []int{1, 10, 100}

	for _, size := range sizes {
		for _, keyCount := range keyCounts {
			if keyCount > size {
				continue
			}

			b.Run(fmt.Sprintf("size_%d_keys_%d_all_found", size, keyCount), func(b *testing.B) {
				m := createIntMap(size)
				keys := make([]int, keyCount)
				for i := 0; i < keyCount; i++ {
					keys[i] = i
				}

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					_ = just.MapContainsKeysAll(m, keys)
				}
			})

			b.Run(fmt.Sprintf("size_%d_keys_%d_one_missing", size, keyCount), func(b *testing.B) {
				m := createIntMap(size)
				keys := make([]int, keyCount)
				for i := 0; i < keyCount-1; i++ {
					keys[i] = i
				}
				keys[keyCount-1] = size + 1 // One key that doesn't exist

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					_ = just.MapContainsKeysAll(m, keys)
				}
			})
		}
	}
}

// BenchmarkMapApply benchmarks the MapApply function
func BenchmarkMapApply(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			m := createIntMap(size)
			sum := 0

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				just.MapApply(m, func(k, v int) {
					sum += v // Simple operation
				})
			}
		})
	}
}

// BenchmarkMapJoin benchmarks the MapJoin function
func BenchmarkMapJoin(b *testing.B) {
	mapCounts := []int{2, 5, 10}
	sizes := []int{10, 100, 1000}

	for _, mapCount := range mapCounts {
		for _, size := range sizes {
			b.Run(fmt.Sprintf("maps_%d_size_%d", mapCount, size), func(b *testing.B) {
				maps := make([]map[int]int, mapCount)
				for i := 0; i < mapCount; i++ {
					m := make(map[int]int, size)
					for j := 0; j < size; j++ {
						m[i*size+j] = j
					}
					maps[i] = m
				}

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					_ = just.MapJoin(maps...)
				}
			})
		}
	}
}

// BenchmarkMapGetDefault benchmarks the MapGetDefault function
func BenchmarkMapGetDefault(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d_exists", size), func(b *testing.B) {
			m := createIntMap(size)
			key := size / 2
			defaultVal := -1

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.MapGetDefault(m, key, defaultVal)
			}
		})

		b.Run(fmt.Sprintf("size_%d_not_exists", size), func(b *testing.B) {
			m := createIntMap(size)
			key := size + 1
			defaultVal := -1

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.MapGetDefault(m, key, defaultVal)
			}
		})
	}
}

// BenchmarkMapNotNil benchmarks the MapNotNil function
func BenchmarkMapNotNil(b *testing.B) {
	b.Run("nil_map", func(b *testing.B) {
		var m map[int]int

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = just.MapNotNil(m)
		}
	})

	b.Run("non_nil_map", func(b *testing.B) {
		m := createIntMap(100)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = just.MapNotNil(m)
		}
	})
}

// BenchmarkMapDropKeys benchmarks the MapDropKeys function
func BenchmarkMapDropKeys(b *testing.B) {
	sizes := []int{10, 1000}
	dropCounts := []int{1, 10, 100}

	for _, size := range sizes {
		for _, dropCount := range dropCounts {
			if dropCount > size {
				continue
			}

			b.Run(fmt.Sprintf("size_%d_drop_%d", size, dropCount), func(b *testing.B) {
				keys := make([]int, dropCount)
				for i := 0; i < dropCount; i++ {
					keys[i] = i
				}

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					m := createIntMap(size)
					just.MapDropKeys(m, keys...)
				}
			})
		}
	}
}

// BenchmarkMapPopKeyDefault benchmarks the MapPopKeyDefault function
func BenchmarkMapPopKeyDefault(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d_exists", size), func(b *testing.B) {
			defaultVal := -1

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				m := createIntMap(size)
				key := size / 2
				_ = just.MapPopKeyDefault(m, key, defaultVal)
			}
		})

		b.Run(fmt.Sprintf("size_%d_not_exists", size), func(b *testing.B) {
			m := createIntMap(size)
			key := size + 1
			defaultVal := -1

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.MapPopKeyDefault(m, key, defaultVal)
			}
		})
	}
}

// BenchmarkMapSetVal benchmarks the MapSetVal function
func BenchmarkMapSetVal(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d_update", size), func(b *testing.B) {
			key := size / 2
			val := 999

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				m := createIntMap(size)
				_ = just.MapSetVal(m, key, val)
			}
		})

		b.Run(fmt.Sprintf("size_%d_add", size), func(b *testing.B) {
			key := size + 1
			val := 999

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				m := createIntMap(size)
				_ = just.MapSetVal(m, key, val)
			}
		})
	}
}
