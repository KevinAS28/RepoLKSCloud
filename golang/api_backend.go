package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var password string = "Admin#1234" // password PDSN

// struktur input data json yang diterima nanti
type RequestBody struct {
	Kunci string // variabel di dalam struct, karakter pertama harus kapital, namun saat di postman itu tidak peduli kapital (case insenstivie)
}

func testPassword(w http.ResponseWriter, r *http.Request) {
	// Baca input
	body, err := io.ReadAll(r.Body)

	// kalau error saat baca input
	if err != nil {
		fmt.Fprintf(w, "Error reading request body: %v", err) // respon teks
		return                                                // fungsi langsung berakhir (tidak usah lanjut ke bawah codenya)
	}

	// Kalau inputnya kosong
	if len(body) == 0 {
		fmt.Fprintf(w, "Kamu tidak memasukan data apa-apa") // respon teks
		return                                              // fungsi langsung berakhir (tidak usah lanjut ke bawah codenya)
	}

	// Parsing input data json, agar bisa diproses
	var reqBody RequestBody
	err = json.Unmarshal(body, &reqBody)

	// Kalau error saat parsing input data ke json
	if err != nil {
		fmt.Fprintf(w, "Bukan JSON itu: %v", err) // respon bentuk teks
		return
	}

	// variable tampungan respon json
	var message map[string]string = nil

	if reqBody.Kunci != password { // kalau key tidak sesuai dengan password
		// isi variable tampung dengan data
		message = map[string]string{
			"status": "400",
			"pesan":  "Password anda salah",
		}
	} else { // selain "passwordnya tidak benar", maka passwordnya benar. "else"="selainnya"
		// isi variable tampung dengan data
		message = map[string]string{
			"status": "200",
			"pesan":  "Password anda benar",
		}
	}

	w.Header().Set("Content-Type", "application/json") // http header
	err = json.NewEncoder(w).Encode(message)           // respon dengan bentuk json
}

func main() {
	http.HandleFunc("/", testPassword)           // jika alamatnya ke root ("/") maka panggil fungsi testPassword
	fmt.Println("Server listening on port 8080") // print output
	http.ListenAndServe(":8080", nil)            // mulai program sebagai server ke port 8080
}
