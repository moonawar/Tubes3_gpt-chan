package algorithm

import (
	"fmt"
	"testing"
)

var test_cases = map[string]int {
	"berapa 1*4+5": 1,
	"hitung 2^*4-6": 1,
	"19/05/2003?": 2,
	"tambah pertanyaan matkul apa yang paling seru di semester 4 dengan jawaban stima": 4,
	"tambah      pertanyaan  matkul apa   yang   paling seru di semester 4  dengan   jawaban stima": 4,
	"tambah xas matkul apa yang paling seru di semester 4? dengan jawaban stima": 0,
	"hapus pertanyaan matkul apa yang paling seru di semester 4?": 8,
	"hapus pertanyaa n xx": 0,
	"berapa 1*4+5? dan hari apa tanggal 19/05/2023?": 3,
	"berapa 1*4+5? dan hari apa tanggal 19/05/2023? dan tambah pertanyaan matkul apa yang paling seru di semester 4 dengan jawaban stima?": 7,
	"axolotl" : 0,
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