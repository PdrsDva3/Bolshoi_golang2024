package storage

import (
	"strconv"
	"testing"
)

type testSG struct {
	name string
	key  string
	val  string
}

func TestSG(t *testing.T) {
	tests := []testSG{
		{"test1", "string", "string"},
		{"test2", "111", "222"},
		{"test3", "222", "111"},
		{"test4", "hahaha", "hhoooho"},
		{"test5", "hello", "world"},
	}
	st, err := InitStorage()
	if err != nil {
		t.Error(err)
		return
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			st.Set(test.key, test.val)

			val := st.Get(test.key)
			if *val != test.val {
				t.Errorf("expect %v, but got %v", test.val, val)
			}
		})
	}

}

type testGetKind struct {
	name string
	key  string
	val  string
	kind string
}

func TestGetKind(t *testing.T) {
	tests := []testGetKind{
		{"test1", "string", "string", StringType},
		{"test2", "111", "222", DigitalType},
		{"test3", "222", "111", DigitalType},
		{"test4", "hahaha", "hhoooho", StringType},
		{"test5", "hellooooooooooooooooooooo", "world12121", StringType},
	}
	st, err := InitStorage()
	if err != nil {
		t.Error(err)
		return
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			st.Set(test.key, test.val)
			val := st.Get(test.key)
			kind := st.GetKind(test.key)

			if *val != test.val {
				t.Errorf("expect %v, but got %v", test.val, val)
			}

			if kind != test.kind {
				t.Errorf("expect %v, but got %v", test.kind, kind)
			}
		})
	}
}

type testGetNil struct {
	name string
	key  string
	val  *string
	kind string
}

func TestGetNil(t *testing.T) {
	tests := []testGetNil{
		{"test1", "nil", nil, UnknownType},
	}
	st, err := InitStorage()
	if err != nil {
		t.Error(err)
		return
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			val := st.Get(test.key)
			kind := st.GetKind(test.key)

			if val != test.val {
				t.Errorf("expect %v, but got %v", test.val, val)
			}

			if kind != test.kind {
				t.Errorf("expect %v, but got %v", test.kind, kind)
			}
		})
	}
}

type benchmarks struct {
	name  string
	count int
}

var btesting = []benchmarks{
	{"5", 5},
	{"100", 100},
	{"500", 500},
	{"2500", 2500},
	{"10000", 10000},
	{"100000", 100000},
	{"1000000", 1000000},
}

func BenchmarkInit(b *testing.B) {
	for _, tCase := range btesting {
		b.Run(tCase.name, func(b *testing.B) {
			for i := 0; i < tCase.count; i++ {
				_, err := InitStorage()
				if err != nil {
					b.Error(err)
				}
			}
		})
	}
}

func BenchmarkSG(b *testing.B) {
	for _, testCase := range btesting {
		b.Run(testCase.name, func(bb *testing.B) {
			var test []testSG
			for i := 0; i < testCase.count; i++ {
				test = append(test, testSG{strconv.Itoa(i + 1), "key" + strconv.Itoa(i), "value" + strconv.Itoa(i)})
			}

			st, err := InitStorage()
			if err != nil {
				b.Error(err)
			}

			bb.ResetTimer()

			for _, t := range test {

				st.Set(t.key, t.val)

				val := st.Get(t.key)
				if *val != t.val {
					b.Fatal("ahhh")
				}
			}
		})
	}
}

func BenchmarkS(b *testing.B) {
	for _, testCase := range btesting {
		b.Run(testCase.name, func(bb *testing.B) {
			var test []testSG
			for i := 0; i < testCase.count; i++ {
				test = append(test, testSG{strconv.Itoa(i + 1), "key" + strconv.Itoa(i), "value" + strconv.Itoa(i)})
			}

			st, err := InitStorage()
			if err != nil {
				b.Error(err)
			}

			bb.ResetTimer()

			for _, t := range test {
				st.Set(t.key, t.val)
			}
		})
	}
}

func BenchmarkGK(b *testing.B) {
	for _, testCase := range btesting {
		b.Run(testCase.name, func(bb *testing.B) {
			var test []testGetKind
			for i := 0; i < testCase.count; i++ {
				test = append(test, testGetKind{strconv.Itoa(i + 1), "key" + strconv.Itoa(i), "value" + strconv.Itoa(i), StringType})
			}

			st, err := InitStorage()
			if err != nil {
				b.Error(err)
			}

			bb.ResetTimer()

			for _, t := range test {
				st.GetKind(t.key)
			}
		})
	}
}
