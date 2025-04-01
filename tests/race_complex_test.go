package tests

import (
	"testing"

	"github.com/KyleBanks/depth"
)

func TestTree_Race(t *testing.T) {
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
		t.Run(tc.Name(), func(t *testing.T) {
			testRaceTree(tc.PackageName, &depth.Tree{
				ResolveInternal: tc.ResolveInternal,
				ResolveTest:     tc.ResolveTest,
			}, t)
		})
	}
}

func testRaceTree(pkg string, tree *depth.Tree, t *testing.T) {
	if err := tree.Resolve(pkg); err != nil {
		t.Fatal(err)
	}
}
