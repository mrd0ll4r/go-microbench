package go_microbench_test

import (
	"testing"
)

func BenchmarkSelectDefaultAfterReadOne(b *testing.B) {
	c0 := make(chan struct{})
	for i := 0; i < b.N; i++ {
		select {
		case <-c0:
		default:
		}
	}
}

func BenchmarkSelectDefaultAfterReadTwo(b *testing.B) {
	c0 := make(chan struct{})
	c1 := make(chan struct{})
	for i := 0; i < b.N; i++ {
		select {
		case <-c0:
		case <-c1:
		default:
		}
	}
}

func BenchmarkSelectDefaultAfterReadThree(b *testing.B) {
	c0 := make(chan struct{})
	c1 := make(chan struct{})
	c2 := make(chan struct{})
	for i := 0; i < b.N; i++ {
		select {
		case <-c0:
		case <-c1:
		case <-c2:
		default:
		}
	}
}

func BenchmarkSelectDefaultAfterReadFour(b *testing.B) {
	c0 := make(chan struct{})
	c1 := make(chan struct{})
	c2 := make(chan struct{})
	c3 := make(chan struct{})
	for i := 0; i < b.N; i++ {
		select {
		case <-c0:
		case <-c1:
		case <-c2:
		case <-c3:
		default:
		}
	}
}

func BenchmarkSelectDefaultAfterReadFive(b *testing.B) {
	c0 := make(chan struct{})
	c1 := make(chan struct{})
	c2 := make(chan struct{})
	c3 := make(chan struct{})
	c4 := make(chan struct{})
	for i := 0; i < b.N; i++ {
		select {
		case <-c0:
		case <-c1:
		case <-c2:
		case <-c3:
		case <-c4:
		default:
		}
	}
}

func BenchmarkSelectDefaultAfterReadSix(b *testing.B) {
	c0 := make(chan struct{})
	c1 := make(chan struct{})
	c2 := make(chan struct{})
	c3 := make(chan struct{})
	c4 := make(chan struct{})
	c5 := make(chan struct{})
	for i := 0; i < b.N; i++ {
		select {
		case <-c0:
		case <-c1:
		case <-c2:
		case <-c3:
		case <-c4:
		case <-c5:
		default:
		}
	}
}

func BenchmarkSelectDefaultAfterReadSeven(b *testing.B) {
	c0 := make(chan struct{})
	c1 := make(chan struct{})
	c2 := make(chan struct{})
	c3 := make(chan struct{})
	c4 := make(chan struct{})
	c5 := make(chan struct{})
	c6 := make(chan struct{})
	for i := 0; i < b.N; i++ {
		select {
		case <-c0:
		case <-c1:
		case <-c2:
		case <-c3:
		case <-c4:
		case <-c5:
		case <-c6:
		default:
		}
	}
}

func BenchmarkSelectDefaultAfterReadEight(b *testing.B) {
	c0 := make(chan struct{})
	c1 := make(chan struct{})
	c2 := make(chan struct{})
	c3 := make(chan struct{})
	c4 := make(chan struct{})
	c5 := make(chan struct{})
	c6 := make(chan struct{})
	c7 := make(chan struct{})
	for i := 0; i < b.N; i++ {
		select {
		case <-c0:
		case <-c1:
		case <-c2:
		case <-c3:
		case <-c4:
		case <-c5:
		case <-c6:
		case <-c7:
		default:
		}
	}
}

func BenchmarkSelectDefaultAfterReadNine(b *testing.B) {
	c0 := make(chan struct{})
	c1 := make(chan struct{})
	c2 := make(chan struct{})
	c3 := make(chan struct{})
	c4 := make(chan struct{})
	c5 := make(chan struct{})
	c6 := make(chan struct{})
	c7 := make(chan struct{})
	c8 := make(chan struct{})
	for i := 0; i < b.N; i++ {
		select {
		case <-c0:
		case <-c1:
		case <-c2:
		case <-c3:
		case <-c4:
		case <-c5:
		case <-c6:
		case <-c7:
		case <-c8:
		default:
		}
	}
}

func BenchmarkSelectDefaultBeforeReadOne(b *testing.B) {
	c0 := make(chan struct{})
	for i := 0; i < b.N; i++ {
		select {
		default:
		case <-c0:
		}
	}
}

func BenchmarkSelectDefaultBeforeReadTwo(b *testing.B) {
	c0 := make(chan struct{})
	c1 := make(chan struct{})
	for i := 0; i < b.N; i++ {
		select {
		default:
		case <-c0:
		case <-c1:
		}
	}
}

func BenchmarkSelectDefaultBeforeReadThree(b *testing.B) {
	c0 := make(chan struct{})
	c1 := make(chan struct{})
	c2 := make(chan struct{})
	for i := 0; i < b.N; i++ {
		select {
		default:
		case <-c0:
		case <-c1:
		case <-c2:
		}
	}
}

func BenchmarkSelectDefaultAfterReadOneBuffered(b *testing.B) {
	c0 := make(chan struct{}, 1)
	for i := 0; i < b.N; i++ {
		select {
		case <-c0:
		default:
		}
	}
}

func BenchmarkSelectDefaultAfterReadTwoBuffered(b *testing.B) {
	c0 := make(chan struct{}, 1)
	c1 := make(chan struct{}, 1)
	for i := 0; i < b.N; i++ {
		select {
		case <-c0:
		case <-c1:
		default:
		}
	}
}

func BenchmarkSelectDefaultAfterReadThreeBuffered(b *testing.B) {
	c0 := make(chan struct{}, 1)
	c1 := make(chan struct{}, 1)
	c2 := make(chan struct{}, 1)
	for i := 0; i < b.N; i++ {
		select {
		case <-c0:
		case <-c1:
		case <-c2:
		default:
		}
	}
}

// ====================================================================
//  Writing
// ====================================================================

func BenchmarkSelectDefaultAfterWriteOne(b *testing.B) {
	c0 := make(chan struct{})
	s := struct{}{}
	for i := 0; i < b.N; i++ {
		select {
		case c0 <- s:
		default:
		}
	}
}

func BenchmarkSelectDefaultAfterWriteTwo(b *testing.B) {
	c0 := make(chan struct{})
	c1 := make(chan struct{})
	s := struct{}{}
	for i := 0; i < b.N; i++ {
		select {
		case c0 <- s:
		case c1 <- s:
		default:
		}
	}
}

func BenchmarkSelectDefaultAfterWriteThree(b *testing.B) {
	c0 := make(chan struct{})
	c1 := make(chan struct{})
	c2 := make(chan struct{})
	s := struct{}{}
	for i := 0; i < b.N; i++ {
		select {
		case c0 <- s:
		case c1 <- s:
		case c2 <- s:
		default:
		}
	}
}

func BenchmarkSelectDefaultBeforeWriteOne(b *testing.B) {
	c0 := make(chan struct{})
	s := struct{}{}
	for i := 0; i < b.N; i++ {
		select {
		default:
		case c0 <- s:
		}
	}
}

func BenchmarkSelectDefaultBeforeWriteTwo(b *testing.B) {
	c0 := make(chan struct{})
	c1 := make(chan struct{})
	s := struct{}{}
	for i := 0; i < b.N; i++ {
		select {
		default:
		case c0 <- s:
		case c1 <- s:
		}
	}
}

func BenchmarkSelectDefaultBeforeWriteThree(b *testing.B) {
	c0 := make(chan struct{})
	c1 := make(chan struct{})
	c2 := make(chan struct{})
	s := struct{}{}
	for i := 0; i < b.N; i++ {
		select {
		default:
		case c0 <- s:
		case c1 <- s:
		case c2 <- s:
		}
	}
}

func BenchmarkSelectDefaultAfterWriteOneBuffered(b *testing.B) {
	c0 := make(chan struct{}, 1)
	s := struct{}{}
	c0 <- s
	for i := 0; i < b.N; i++ {
		select {
		case c0 <- s:
		default:
		}
	}
}

func BenchmarkSelectDefaultAfterWriteTwoBuffered(b *testing.B) {
	c0 := make(chan struct{}, 1)
	c1 := make(chan struct{}, 1)
	s := struct{}{}
	c0 <- s
	c1 <- s
	for i := 0; i < b.N; i++ {
		select {
		case c0 <- s:
		case c1 <- s:
		default:
		}
	}
}

func BenchmarkSelectDefaultAfterWriteThreeBuffered(b *testing.B) {
	c0 := make(chan struct{}, 1)
	c1 := make(chan struct{}, 1)
	c2 := make(chan struct{}, 1)
	s := struct{}{}
	c0 <- s
	c1 <- s
	c2 <- s
	for i := 0; i < b.N; i++ {
		select {
		case c0 <- s:
		case c1 <- s:
		case c2 <- s:
		default:
		}
	}
}
