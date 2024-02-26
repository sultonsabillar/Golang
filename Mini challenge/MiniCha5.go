package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Struct untuk menyimpan data biodata pengguna
type Biodata struct {
	Nama   string
	Umur   int
	Alamat string
	Email  string
}

// Data statis untuk biodata pengguna
var dataBiodata = map[string]Biodata{
	"user1@example.com": {"User1", 25, "Alamat User1", "user1@example.com"},
	"user2@example.com": {"User2", 30, "Alamat User2", "user2@example.com"},
}

// // Handler untuk halaman index (halaman biodata)
// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	// Mendapatkan email dari cookie atau form data (setelah login)
// 	var email string
// 	if cookie, err := r.Cookie("email"); err == nil {
// 		email = cookie.Value
// 	} else {
// 		email = r.FormValue("email")
// 	}

	// // Memeriksa apakah email ada dalam data biodata
	// biodata, ok := dataBiodata[email]
	// if !ok {
	// 	http.Error(w, "Biodata tidak ditemukan", http.StatusNotFound)
	// 	return
	// }

	// Menampilkan biodata pengguna pada halaman index
	tmpl := template.Must(template.New("index").Parse(indexTemplate))
	tmpl.Execute(w, biodata)
}

// Handler untuk halaman login
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.ServeFile(w, r, "login.html")
		return
	}

	email := r.FormValue("email")
	_, ok := dataBiodata[email]
	if !ok {
		http.Error(w, "Email tidak valid", http.StatusBadRequest)
		return
	}

	// Set cookie email
	http.SetCookie(w, &http.Cookie{
		Name:  "email",
		Value: email,
	})

	// Redirect ke halaman index setelah login berhasil
	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)

	fmt.Println("Server started at localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// Template untuk halaman index
var indexTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Biodata</title>
</head>
<body>
    <h1>Biodata Pengguna</h1>
    <ul>
        <li><strong>Nama:</strong> {{ .Nama }}</li>
        <li><strong>Umur:</strong> {{ .Umur }}</li>
        <li><strong>Alamat:</strong> {{ .Alamat }}</li>
        <li><strong>Email:</strong> {{ .Email }}</li>
    </ul>
</body>
</html>
`
