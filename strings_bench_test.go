package stringutil_test

import (
	"testing"

	"github.com/naoina/go-stringutil"
)

var benchcaseForCamelCase = "the_quick_brown_fox_jumps_over_the_lazy_dog"

func BenchmarkToUpperCamelCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringutil.ToUpperCamelCase(benchcaseForCamelCase)
	}
}

var benchcaseForSnakeCase = "TheQuickBrownFoxJumpsOverTheLazyDog"

func BenchmarkToSnakeCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringutil.ToSnakeCase(benchcaseForSnakeCase)
	}
}

func BenchmarkToSnakeCaseASCII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringutil.ToSnakeCaseASCII(benchcaseForSnakeCase)
	}
}
