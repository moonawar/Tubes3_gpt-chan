package algorithm

import (
	"fmt"
	"testing"
)

var test_cases = map[string]int{
	"berapa 1*4+5":  8,
	"hitung 2^*4-6": 8,
	"19/05/2003?":   4,
	"tambah pertanyaan matkul apa yang paling seru di semester 4 dengan jawaban stima":                                 2,
	"tambah pertanyaan berapa nilai tubes developer gpt-chan dengan jawaban nilai tubes developer gpt-chan adalah 115": 2,
	"tambah      pertanyaan  matkul apa   yang   paling seru di semester 4  dengan   jawaban stima":                    2,
	"tambah xas matkul apa yang paling seru di semester 4? dengan jawaban stima":                                       0,
	"hapus pertanyaan matkul apa yang paling seru di semester 4?":                                                      1,
	"hapus pertanyaa n xx":                           0,
	"berapa 1*4+5? dan hari apa tanggal 19/05/2023?": 12,
	"berapa 1*4+5? dan hari apa tanggal 19/05/2023? dan tambah pertanyaan matkul apa yang paling seru di semester 4 dengan jawaban stima?": 14,
	"axolotl": 0,
	"-2":      0,
}

func TestClassify(t *testing.T) {
	a := New()
	for text, expected := range test_cases {
		res := a.Classify(text)
		if res != expected {
			fmt.Println("FAIL    \"" + text + "\"")
			t.Errorf("Classify(%s) = %d, expected %d", text, res, expected)
		} else {
			fmt.Println("PASS    \"" + text + "\"")
		}
	}
}
