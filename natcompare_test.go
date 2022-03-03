package natcompare_test

import (
	"github.com/simon-engledew/natcompare"
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

func TestCompare(t *testing.T) {
	require.True(t, natcompare.Less("100E", "200E"))
	require.True(t, natcompare.Less("200E", "1000E"))
}

func less(values []string) func(i, j int) bool {
	return func(i, j int) bool {
		return natcompare.Less(values[i], values[j])
	}
}

func TestSort(t *testing.T) {
	values := []string{"a", "a0", "a1", "a1a", "a1b", "a2", "a10", "a20"}
	require.True(t, sort.SliceIsSorted(values, less(values)))

	values = []string{"1000 A", "A", "B"}
	require.True(t, sort.SliceIsSorted(values, less(values)))

	values = []string{"a1.txt", "a2.txt", "a20.txt"}
	require.True(t, sort.SliceIsSorted(values, less(values)))

	values = []string{
		"1-02",
		"1-2",
		"1-20",
		"10-20",
		"fred",
		"jane",
		"pic01",
		"pic02",
		"pic2",
		"pic02a",
		"pic3",
		"pic4",
		"pic05",
		"pic100",
		"pic100a",
		"pic120",
		"pic121",
		"pic02000",
		"pic 4 else",
		"pic 5",
		"pic 5",
		"pic 5 something",
		"pic 6",
		"pic   7",
		"tom",
		"x2-g8",
		"x2-y7",
		"x2-y08",
		"x8-y8",
	}
	require.True(t, sort.SliceIsSorted(values, less(values)))

	values = []string{
		"1999-3-3",
		"1999-12-25",
		"2000-1-2",
		"2000-1-10",
		"2000-3-23",
	}
	require.True(t, sort.SliceIsSorted(values, less(values)))
}
