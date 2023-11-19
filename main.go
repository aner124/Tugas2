package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Struct untuk menyimpan informasi pengguna
type User struct {
	Name          string
	Score         int
	CorrectAnswer int
	WrongAnswer   int
}

// Struct untuk menyimpan informasi pertanyaan
type Question struct {
	Text    string
	Options []Option
}

// Struct untuk menyimpan informasi opsi jawaban
type Option struct {
	Text  string
	Score int
}

func main() {
	// Buat objek user
	var user User

	// Minta pengguna untuk memasukkan nama
	fmt.Print("Imput  nama : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	user.Name = scanner.Text()

	// Inisialisasi pertanyaan
	questions := []Question{
		{"Pulau komodo terletak di provinsi?", []Option{
			{"NTT", 1},
			{"Aceh", 0},
			{"NTB", 0},
			{"Jawa Tengah", 0},
		}},
		{"Di ujung langit tedapat?", []Option{
			{"Bintang", 0},
			{"awan", 0},
			{"T", 1},
			{"Matahari", 0},
		}},
		{"Siapakah penemu mesin uap ?", []Option{
			{"James Watt", 1},
			{"Ronaldo", 0},
			{"Messi", 0},
			{"Agus ", 0},
		}},
		{"Apa huruf ke empat dalam abjad?", []Option{
			{"100", 0},
			{"A", 1},
			{"Tdak tau", 0},
			{"semua jawaban benar", 0},
		}},
		{"Siapa presiden ke tiga Indonsia?", []Option{
			{"Bj.Habibie", 1},
			{"Ir.soekarno", 0},
			{"Joko Widodo", 0},
			{"Megawati", 0},
		}},
	}

	// Mulai ujian
	for _, q := range questions {
		fmt.Printf("\n%s\nPilihan Jawaban:\n", q.Text)
		for i, option := range q.Options {
			fmt.Printf("%d. %s\n", i+1, option.Text)
		}

		fmt.Print("Masukkan nomor jawaban Anda : ")
		scanner.Scan()
		userAnswer := strings.ToLower(strings.TrimSpace(scanner.Text()))

		// Validasi input
		index, err := getIndexFromKey(userAnswer)
		if err != nil || index < 0 || index >= len(q.Options) {
			fmt.Println("Input tidak valid. Mohon masukkan nomor jawaban yang benar.")
			continue
		}

		// Tambahkan skor
		user.Score += q.Options[index].Score
		if q.Options[index].Score > 0 {
			fmt.Println("Jawaban benar!")
			user.CorrectAnswer++
		} else {
			fmt.Println("Jawaban salah. Jawaban yang benar adalah:", q.Options[getCorrectOptionIndex(q.Options)].Text)
			user.WrongAnswer++
		}
	}

	// Tampilkan hasil
	fmt.Printf("\n=== Hasil Ujian ===\n")
	fmt.Printf("Nama: %s\n", user.Name)
	fmt.Printf("Skor: %d\n", user.Score)
	fmt.Printf("Jumlah Jawaban Benar: %d\n", user.CorrectAnswer)
	fmt.Printf("Jumlah Jawaban Salah: %d\n", user.WrongAnswer)
}

// getIndexFromKey mengembalikan indeks dari map berdasarkan kunci (key).
func getIndexFromKey(key string) (int, error) {
	index, err := strconv.Atoi(key)
	if err != nil {
		return -1, fmt.Errorf("invalid key")
	}
	return index - 1, nil
}

// getCorrectOptionIndex mengembalikan indeks dari opsi jawaban yang benar.
func getCorrectOptionIndex(options []Option) int {
	for i, option := range options {
		if option.Score > 0 {
			return i
		}
	}
	return -1
}
