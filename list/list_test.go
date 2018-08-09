package list

import (
	"strings"
	"testing"
)

func TestList_BFS(t *testing.T) {

	testCase := []struct {
		list   list
		find   string
		have   bool
		result string
	}{
		{
			list: list{
				items: map[string]string{
					"4":   "1",
					"6":   "2",
					"12":  "3",
					"34":  "4",
					"56":  "5",
					"45":  "6",
					"123": "7",
				},
			},
			find:   "123456",
			have:   true,
			result: "7,6,2",
		},
		{
			list: list{
				items: map[string]string{
					"4":   "1",
					"6":   "2",
					"12":  "3",
					"34":  "4",
					"56":  "5",
					"45":  "6",
					"123": "7",
				},
			},
			find:   "1234567",
			have:   false,
			result: "",
		},
		{
			list: list{
				items: map[string]string{
					"12345": "1",
					"1234":  "2",
					"56":    "3",
					"6":     "4",
				},
			},
			find:   "123456",
			have:   true,
			result: "1,4",
		},
		{
			list: list{
				items: map[string]string{
					"12345": "1",
					"1234":  "2",
					"56":    "3",
				},
			},
			find:   "123456",
			have:   true,
			result: "2,3",
		},
		{
			list: list{
				items: map[string]string{
					"012345": "1",
					"6":      "2",
					"7":      "3",
					"8":      "4",
					"9":      "5",
					"0123":   "6",
					"456":    "7",
					"789":    "8",
				},
			},
			find:   "0123456789",
			have:   true,
			result: "1,2,8",
		},
		{
			list: list{
				items: map[string]string{
					"AB":  "1",
					"BC":  "2",
					"CD":  "3",
					"A":   "4",
					"DDC": "5",
					"DD":  "6",
				},
			},
			find:   "ABCDADDCDD",
			have:   true,
			result: "1,3,4,5,6",
		},
		{
			list: list{
				items: map[string]string{
					"0123456": "1",
					"78":      "2",
					"9a":      "3",
					"bc":      "4",
					"de":      "5",
					"01234":   "6",
					"56789":   "7",
					"abcde":   "8",
				},
			},
			find:   "0123456789abcde",
			have:   true,
			result: "6,7,8",
		},
		{
			list: list{
				items: map[string]string{
					"012345":  "1",
					"6":       "2",
					"7":       "3",
					"8":       "4",
					"9":       "5",
					"0123":    "6",
					"456":     "7",
					"789":     "8",
					"012":     "9",
					"3456789": "10",
				},
			},
			find:   "0123456789",
			have:   true,
			result: "9,10",
		},
		{
			list: list{
				items: map[string]string{},
			},
			find:   "123456",
			have:   false,
			result: "",
		},
	}

	for _, v := range testCase {

		result, ok := v.list.BFS(v.find)
		if ok != v.have {
			t.Fatalf("Expext %t, but result %t\n\rTest list: %v", v.have, ok, v.list)
		}

		if !v.have {
			continue
		}

		if strings.Join(result, ",") != v.result {
			t.Fatalf("Expext %s, but %s\n\rTest list: %v", v.result, strings.Join(result, ","), v.list)
		}

	}

}

func BenchmarkList_BFS(b *testing.B) {

	list := New()
	list.Put("4", "1")
	list.Put("6", "2")
	list.Put("12", "2")
	list.Put("34", "2")
	list.Put("56", "2")
	list.Put("45", "2")
	list.Put("123", "2")

	find := "123456"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		list.BFS(find)
	}

}
