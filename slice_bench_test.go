package just_test

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"testing"

	"github.com/kazhuravlev/just"
)

// Helper functions to create test slices
func createIntSlice(size int) []int {
	s := make([]int, size)
	for i := 0; i < size; i++ {
		s[i] = i
	}
	return s
}

func createIntSliceWithDuplicates(size int, dupFactor int) []int {
	s := make([]int, size)
	for i := 0; i < size; i++ {
		s[i] = i / dupFactor
	}
	return s
}

// BenchmarkSliceUniq benchmarks the SliceUniq function
func BenchmarkSliceUniq(b *testing.B) {
	sizes := []int{10, 1000}
	dupFactors := []int{1, 2, 10} // 1 = no dups, 2 = 50% dups, 10 = 90% dups

	for _, size := range sizes {
		for _, dupFactor := range dupFactors {
			b.Run(fmt.Sprintf("size_%d_dupFactor_%d", size, dupFactor), func(b *testing.B) {
				slice := createIntSliceWithDuplicates(size, dupFactor)

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					_ = just.SliceUniq(slice)
				}
			})
		}
	}
}

// BenchmarkSliceUniqStable benchmarks the SliceUniqStable function
func BenchmarkSliceUniqStable(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			slice := createIntSliceWithDuplicates(size, 2)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceUniqStable(slice)
			}
		})
	}
}

// BenchmarkSliceMap benchmarks the SliceMap function
func BenchmarkSliceMap(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			slice := createIntSlice(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceMap(slice, func(v int) string {
					return strconv.Itoa(v)
				})
			}
		})
	}
}

// BenchmarkSliceFlatMap benchmarks the SliceFlatMap function
func BenchmarkSliceFlatMap(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			slice := createIntSlice(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceFlatMap(slice, func(v int) []int {
					return []int{v, v * 2, v * 3}
				})
			}
		})
	}
}

// BenchmarkSliceFlatMap2 benchmarks the SliceFlatMap2 function
func BenchmarkSliceFlatMap2(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			slice := createIntSlice(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceFlatMap2(slice, func(idx int, v int) []int {
					return []int{idx, v}
				})
			}
		})
	}
}

// BenchmarkSliceApply benchmarks the SliceApply function
func BenchmarkSliceApply(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			slice := createIntSlice(size)
			sum := 0

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				just.SliceApply(slice, func(idx int, v int) {
					sum += v
				})
			}
		})
	}
}

// BenchmarkSliceMapErr benchmarks the SliceMapErr function
func BenchmarkSliceMapErr(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d_success", size), func(b *testing.B) {
			slice := createIntSlice(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = just.SliceMapErr(slice, func(v int) (string, error) {
					return strconv.Itoa(v), nil
				})
			}
		})

		b.Run(fmt.Sprintf("size_%d_error_early", size), func(b *testing.B) {
			slice := createIntSlice(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = just.SliceMapErr(slice, func(v int) (string, error) {
					if v == 0 {
						return "", errors.New("error")
					}
					return strconv.Itoa(v), nil
				})
			}
		})
	}
}

// BenchmarkSliceFilter benchmarks the SliceFilter function
func BenchmarkSliceFilter(b *testing.B) {
	sizes := []int{10, 1000}
	filterRatios := []int{2, 10, 100} // Keep 1/2, 1/10, 1/100 of elements

	for _, size := range sizes {
		for _, ratio := range filterRatios {
			b.Run(fmt.Sprintf("size_%d_keep_1/%d", size, ratio), func(b *testing.B) {
				slice := createIntSlice(size)

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					_ = just.SliceFilter(slice, func(v int) bool {
						return v%ratio == 0
					})
				}
			})
		}
	}
}

// BenchmarkSliceReverse benchmarks the SliceReverse function
func BenchmarkSliceReverse(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			slice := createIntSlice(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceReverse(slice)
			}
		})
	}
}

// BenchmarkSliceAny benchmarks the SliceAny function
func BenchmarkSliceAny(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d_found_first", size), func(b *testing.B) {
			slice := createIntSlice(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceAny(slice, func(v int) bool {
					return v == 0
				})
			}
		})

		b.Run(fmt.Sprintf("size_%d_found_last", size), func(b *testing.B) {
			slice := createIntSlice(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceAny(slice, func(v int) bool {
					return v == size-1
				})
			}
		})

		b.Run(fmt.Sprintf("size_%d_not_found", size), func(b *testing.B) {
			slice := createIntSlice(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceAny(slice, func(v int) bool {
					return v < 0
				})
			}
		})
	}
}

// BenchmarkSliceAll benchmarks the SliceAll function
func BenchmarkSliceAll(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d_all_true", size), func(b *testing.B) {
			slice := createIntSlice(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceAll(slice, func(v int) bool {
					return v >= 0
				})
			}
		})

		b.Run(fmt.Sprintf("size_%d_false_at_end", size), func(b *testing.B) {
			slice := createIntSlice(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceAll(slice, func(v int) bool {
					return v < size-1
				})
			}
		})
	}
}

// BenchmarkSliceContainsElem benchmarks the SliceContainsElem function
func BenchmarkSliceContainsElem(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d_found", size), func(b *testing.B) {
			slice := createIntSlice(size)
			elem := size / 2

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceContainsElem(slice, elem)
			}
		})

		b.Run(fmt.Sprintf("size_%d_not_found", size), func(b *testing.B) {
			slice := createIntSlice(size)
			elem := size + 1

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceContainsElem(slice, elem)
			}
		})
	}
}

// BenchmarkSliceAddNotExists benchmarks the SliceAddNotExists function
func BenchmarkSliceAddNotExists(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d_exists", size), func(b *testing.B) {
			slice := createIntSlice(size)
			elem := size / 2

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceAddNotExists(slice, elem)
			}
		})

		b.Run(fmt.Sprintf("size_%d_not_exists", size), func(b *testing.B) {
			slice := createIntSlice(size)
			elem := size + 1

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceAddNotExists(slice, elem)
			}
		})
	}
}

// BenchmarkSliceUnion benchmarks the SliceUnion function
func BenchmarkSliceUnion(b *testing.B) {
	sizes := []int{10, 100, 1000}
	sliceCounts := []int{2, 5, 10}

	for _, count := range sliceCounts {
		for _, size := range sizes {
			b.Run(fmt.Sprintf("slices_%d_size_%d", count, size), func(b *testing.B) {
				slices := make([][]int, count)
				for i := 0; i < count; i++ {
					slices[i] = createIntSlice(size)
				}

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					_ = just.SliceUnion(slices...)
				}
			})
		}
	}
}

// BenchmarkSlice2Map benchmarks the Slice2Map function
func BenchmarkSlice2Map(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			slice := createIntSlice(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.Slice2Map(slice)
			}
		})
	}
}

// BenchmarkSliceDifference benchmarks the SliceDifference function
func BenchmarkSliceDifference(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d_no_overlap", size), func(b *testing.B) {
			old := createIntSlice(size)
			new := make([]int, size)
			for i := 0; i < size; i++ {
				new[i] = i + size
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceDifference(old, new)
			}
		})

		b.Run(fmt.Sprintf("size_%d_half_overlap", size), func(b *testing.B) {
			old := createIntSlice(size)
			new := make([]int, size)
			for i := 0; i < size; i++ {
				new[i] = i + size/2
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceDifference(old, new)
			}
		})
	}
}

// BenchmarkSliceIntersection benchmarks the SliceIntersection function
func BenchmarkSliceIntersection(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			slice1 := createIntSlice(size)
			slice2 := make([]int, size)
			for i := 0; i < size; i++ {
				slice2[i] = i + size/2
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceIntersection(slice1, slice2)
			}
		})
	}
}

// BenchmarkSliceWithoutElem benchmarks the SliceWithoutElem function
func BenchmarkSliceWithoutElem(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			slice := createIntSliceWithDuplicates(size, 10)
			elem := size / 20 // Will have ~10 occurrences

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceWithoutElem(slice, elem)
			}
		})
	}
}

// BenchmarkSliceWithout benchmarks the SliceWithout function
func BenchmarkSliceWithout(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			slice := createIntSlice(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceWithout(slice, func(v int) bool {
					return v%2 == 0
				})
			}
		})
	}
}

// BenchmarkSliceZip benchmarks the SliceZip function
func BenchmarkSliceZip(b *testing.B) {
	sizes := []int{10, 1000}
	sliceCounts := []int{2, 5, 10}

	for _, count := range sliceCounts {
		for _, size := range sizes {
			b.Run(fmt.Sprintf("slices_%d_size_%d", count, size), func(b *testing.B) {
				slices := make([][]int, count)
				for i := 0; i < count; i++ {
					slices[i] = createIntSlice(size)
				}

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					_ = just.SliceZip(slices...)
				}
			})
		}
	}
}

// BenchmarkSliceFillElem benchmarks the SliceFillElem function
func BenchmarkSliceFillElem(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			elem := 42

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceFillElem(size, elem)
			}
		})
	}
}

// BenchmarkSliceNotNil benchmarks the SliceNotNil function
func BenchmarkSliceNotNil(b *testing.B) {
	b.Run("nil_slice", func(b *testing.B) {
		var slice []int

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = just.SliceNotNil(slice)
		}
	})

	b.Run("non_nil_slice", func(b *testing.B) {
		slice := createIntSlice(100)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = just.SliceNotNil(slice)
		}
	})
}

// BenchmarkSliceChunk benchmarks the SliceChunk function
func BenchmarkSliceChunk(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			slice := createIntSlice(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceChunk(slice, func(idx int, v int) bool {
					return v%10 == 0
				})
			}
		})
	}
}

// BenchmarkSliceChunkEvery benchmarks the SliceChunkEvery function
func BenchmarkSliceChunkEvery(b *testing.B) {
	sizes := []int{10, 1000}
	chunkSizes := []int{1, 10, 100}

	for _, size := range sizes {
		for _, chunkSize := range chunkSizes {
			if chunkSize > size {
				continue
			}
			b.Run(fmt.Sprintf("size_%d_chunk_%d", size, chunkSize), func(b *testing.B) {
				slice := createIntSlice(size)

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					_ = just.SliceChunkEvery(slice, chunkSize)
				}
			})
		}
	}
}

// BenchmarkSliceFindFirst benchmarks the SliceFindFirst function
func BenchmarkSliceFindFirst(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d_found_early", size), func(b *testing.B) {
			slice := createIntSlice(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceFindFirst(slice, func(idx int, v int) bool {
					return v == 5
				})
			}
		})

		b.Run(fmt.Sprintf("size_%d_found_late", size), func(b *testing.B) {
			slice := createIntSlice(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceFindFirst(slice, func(idx int, v int) bool {
					return v == size-5
				})
			}
		})
	}
}

// BenchmarkSliceFindFirstElem benchmarks the SliceFindFirstElem function
func BenchmarkSliceFindFirstElem(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			slice := createIntSlice(size)
			elem := size / 2

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceFindFirstElem(slice, elem)
			}
		})
	}
}

// BenchmarkSliceFindLast benchmarks the SliceFindLast function
func BenchmarkSliceFindLast(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			slice := createIntSliceWithDuplicates(size, 10)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceFindLast(slice, func(idx int, v int) bool {
					return v == size/20
				})
			}
		})
	}
}

// BenchmarkSliceFindLastElem benchmarks the SliceFindLastElem function
func BenchmarkSliceFindLastElem(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			slice := createIntSliceWithDuplicates(size, 10)
			elem := size / 20

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceFindLastElem(slice, elem)
			}
		})
	}
}

// BenchmarkSliceFindAll benchmarks the SliceFindAll function
func BenchmarkSliceFindAll(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			slice := createIntSlice(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceFindAll(slice, func(idx int, v int) bool {
					return v%10 == 0
				})
			}
		})
	}
}

// BenchmarkSliceFindAllElements benchmarks the SliceFindAllElements function
func BenchmarkSliceFindAllElements(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			slice := createIntSlice(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceFindAllElements(slice, func(idx int, v int) bool {
					return v%10 == 0
				})
			}
		})
	}
}

// BenchmarkSliceFindAllIndexes benchmarks the SliceFindAllIndexes function
func BenchmarkSliceFindAllIndexes(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			slice := createIntSlice(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceFindAllIndexes(slice, func(idx int, v int) bool {
					return v%10 == 0
				})
			}
		})
	}
}

// BenchmarkSliceRange benchmarks the SliceRange function
func BenchmarkSliceRange(b *testing.B) {
	ranges := []struct {
		start, stop, step int
	}{
		{0, 10, 1},
		{0, 100, 1},
		{0, 1000, 1},
		{0, 10000, 1},
		{0, 10000, 10},
		{0, 10000, 100},
	}

	for _, r := range ranges {
		b.Run(fmt.Sprintf("range_%d_%d_%d", r.start, r.stop, r.step), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceRange(r.start, r.stop, r.step)
			}
		})
	}
}

// BenchmarkSliceEqualUnordered benchmarks the SliceEqualUnordered function
func BenchmarkSliceEqualUnordered(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d_equal", size), func(b *testing.B) {
			slice1 := createIntSlice(size)
			slice2 := make([]int, size)
			copy(slice2, slice1)
			// Shuffle slice2
			rand.Shuffle(len(slice2), func(i, j int) {
				slice2[i], slice2[j] = slice2[j], slice2[i]
			})

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceEqualUnordered(slice1, slice2)
			}
		})

		b.Run(fmt.Sprintf("size_%d_not_equal", size), func(b *testing.B) {
			slice1 := createIntSlice(size)
			slice2 := createIntSlice(size)
			slice2[0] = size + 1

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceEqualUnordered(slice1, slice2)
			}
		})
	}
}

// BenchmarkSliceChain benchmarks the SliceChain function
func BenchmarkSliceChain(b *testing.B) {
	sliceCounts := []int{2, 5, 10}
	sizes := []int{10, 100, 1000}

	for _, count := range sliceCounts {
		for _, size := range sizes {
			b.Run(fmt.Sprintf("slices_%d_size_%d", count, size), func(b *testing.B) {
				slices := make([][]int, count)
				for i := 0; i < count; i++ {
					slices[i] = createIntSlice(size)
				}

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					_ = just.SliceChain(slices...)
				}
			})
		}
	}
}

// BenchmarkSliceSort benchmarks the SliceSort function
func BenchmarkSliceSort(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			less := func(a, b int) bool { return a < b }

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				slice := make([]int, size)
				for j := 0; j < size; j++ {
					slice[j] = rand.Intn(size)
				}
				just.SliceSort(slice, less)
			}
		})
	}
}

// BenchmarkSliceSortCopy benchmarks the SliceSortCopy function
func BenchmarkSliceSortCopy(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			slice := make([]int, size)
			for i := 0; i < size; i++ {
				slice[i] = rand.Intn(size)
			}
			less := func(a, b int) bool { return a < b }

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceSortCopy(slice, less)
			}
		})
	}
}

// BenchmarkSliceGroupBy benchmarks the SliceGroupBy function
func BenchmarkSliceGroupBy(b *testing.B) {
	sizes := []int{10, 1000}
	groupCounts := []int{2, 10, 100}

	for _, size := range sizes {
		for _, groups := range groupCounts {
			if groups > size {
				continue
			}
			b.Run(fmt.Sprintf("size_%d_groups_%d", size, groups), func(b *testing.B) {
				slice := createIntSlice(size)

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					_ = just.SliceGroupBy(slice, func(v int) int {
						return v % groups
					})
				}
			})
		}
	}
}

// BenchmarkSlice2MapFn benchmarks the Slice2MapFn function
func BenchmarkSlice2MapFn(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			slice := createIntSlice(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.Slice2MapFn(slice, func(idx int, v int) (string, int) {
					return strconv.Itoa(idx), v
				})
			}
		})
	}
}

// BenchmarkSlice2MapFnErr benchmarks the Slice2MapFnErr function
func BenchmarkSlice2MapFnErr(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			slice := createIntSlice(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = just.Slice2MapFnErr(slice, func(idx int, v int) (string, int, error) {
					return strconv.Itoa(idx), v, nil
				})
			}
		})
	}
}

// BenchmarkSliceFromElem benchmarks the SliceFromElem function
func BenchmarkSliceFromElem(b *testing.B) {
	b.Run("int", func(b *testing.B) {
		elem := 42

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = just.SliceFromElem(elem)
		}
	})

	b.Run("string", func(b *testing.B) {
		elem := "test"

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = just.SliceFromElem(elem)
		}
	})
}

// BenchmarkSliceGetFirstN benchmarks the SliceGetFirstN function
func BenchmarkSliceGetFirstN(b *testing.B) {
	sizes := []int{10, 1000}
	nValues := []int{1, 10, 100}

	for _, size := range sizes {
		for _, n := range nValues {
			if n > size {
				continue
			}
			b.Run(fmt.Sprintf("size_%d_n_%d", size, n), func(b *testing.B) {
				slice := createIntSlice(size)

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					_ = just.SliceGetFirstN(slice, n)
				}
			})
		}
	}
}

// BenchmarkSliceCopy benchmarks the SliceCopy function
func BenchmarkSliceCopy(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			slice := createIntSlice(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceCopy(slice)
			}
		})
	}
}

// BenchmarkSliceReplaceFirst benchmarks the SliceReplaceFirst function
func BenchmarkSliceReplaceFirst(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			newElem := -1

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				slice := createIntSlice(size)
				just.SliceReplaceFirst(slice, func(idx int, v int) bool {
					return v == size/2
				}, newElem)
			}
		})
	}
}

// BenchmarkSliceReplaceFirstOrAdd benchmarks the SliceReplaceFirstOrAdd function
func BenchmarkSliceReplaceFirstOrAdd(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d_found", size), func(b *testing.B) {
			slice := createIntSlice(size)
			newElem := -1

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceReplaceFirstOrAdd(slice, func(idx int, v int) bool {
					return v == size/2
				}, newElem)
			}
		})

		b.Run(fmt.Sprintf("size_%d_not_found", size), func(b *testing.B) {
			slice := createIntSlice(size)
			newElem := -1

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceReplaceFirstOrAdd(slice, func(idx int, v int) bool {
					return v < 0
				}, newElem)
			}
		})
	}
}

// BenchmarkSliceLastDefault benchmarks the SliceLastDefault function
func BenchmarkSliceLastDefault(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			slice := createIntSlice(size)
			defaultVal := -1

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceLastDefault(slice, defaultVal)
			}
		})
	}

	b.Run("empty_slice", func(b *testing.B) {
		var slice []int
		defaultVal := -1

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = just.SliceLastDefault(slice, defaultVal)
		}
	})
}

// BenchmarkSlice2Iter benchmarks the Slice2Iter function
func BenchmarkSlice2Iter(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d_full_iteration", size), func(b *testing.B) {
			slice := createIntSlice(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				iter := just.Slice2Iter(slice)
				sum := 0
				iter(func(idx int, v int) bool {
					sum += v
					return true
				})
			}
		})

		b.Run(fmt.Sprintf("size_%d_early_exit", size), func(b *testing.B) {
			slice := createIntSlice(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				iter := just.Slice2Iter(slice)
				count := 0
				iter(func(idx int, v int) bool {
					count++
					return count < 10
				})
			}
		})
	}
}

// BenchmarkSliceIter benchmarks the SliceIter function
func BenchmarkSliceIter(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			slice := createIntSlice(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				iter := just.SliceIter(slice)
				sum := 0
				iter(func(ctx just.IterContext, v int) bool {
					if ctx.IsFirst() || ctx.IsLast() {
						sum += v * 2
					} else {
						sum += v
					}
					return true
				})
			}
		})
	}
}

// BenchmarkSliceShuffle benchmarks the SliceShuffle function
func BenchmarkSliceShuffle(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				slice := createIntSlice(size)
				just.SliceShuffle(slice)
			}
		})
	}
}

// BenchmarkSliceShuffleCopy benchmarks the SliceShuffleCopy function
func BenchmarkSliceShuffleCopy(b *testing.B) {
	sizes := []int{10, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			slice := createIntSlice(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = just.SliceShuffleCopy(slice)
			}
		})
	}
}

// BenchmarkSliceLastN benchmarks the SliceLastN function
func BenchmarkSliceLastN(b *testing.B) {
	sizes := []int{10, 1000}
	nValues := []int{1, 10, 100}

	for _, size := range sizes {
		for _, n := range nValues {
			if n > size {
				continue
			}
			b.Run(fmt.Sprintf("size_%d_n_%d", size, n), func(b *testing.B) {
				slice := createIntSlice(size)

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					_ = just.SliceLastN(slice, n)
				}
			})
		}
	}
}
