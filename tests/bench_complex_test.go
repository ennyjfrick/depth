package tests

import (
	"strings"
	"testing"

	"github.com/KyleBanks/depth"
)

type testArgs struct {
	PackageName     string
	ResolveInternal bool
	ResolveTest     bool
}

func (t testArgs) Name() string {
	split := strings.Split(t.PackageName, "/")
	name := split[len(split)-1]
	if t.ResolveInternal {
		name += "/internal"
	}

	if t.ResolveTest {
		name += "/test"
	}

	return name
}

func BenchmarkTree_ResolveComplex(b *testing.B) {
	tt := []testArgs{
		{
			PackageName:     "gorm.io/gorm",
			ResolveInternal: false,
			ResolveTest:     false,
		},
		{
			PackageName:     "gorm.io/gorm",
			ResolveInternal: true,
			ResolveTest:     false,
		},
		{
			PackageName:     "gorm.io/gorm",
			ResolveInternal: false,
			ResolveTest:     true,
		},
		{
			PackageName:     "gorm.io/gorm",
			ResolveInternal: true,
			ResolveTest:     true,
		},
		{
			PackageName:     "github.com/glebarez/sqlite",
			ResolveInternal: false,
			ResolveTest:     false,
		},
		{
			PackageName:     "github.com/glebarez/sqlite",
			ResolveInternal: true,
			ResolveTest:     false,
		},
		{
			PackageName:     "github.com/glebarez/sqlite",
			ResolveInternal: false,
			ResolveTest:     true,
		},
		{
			PackageName:     "github.com/glebarez/sqlite",
			ResolveInternal: true,
			ResolveTest:     true,
		},
	}

	for _, tc := range tt {
		tc := tc
		b.Run(tc.Name(), func(b *testing.B) {
			benchmarkTree(tc.PackageName, &depth.Tree{
				ResolveInternal: tc.ResolveInternal,
				ResolveTest:     tc.ResolveTest,
			}, b)
		})
	}
}

func benchmarkTree(pkg string, t *depth.Tree, b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := t.Resolve(pkg); err != nil {
			b.Fatal(err)
		}
	}
}
