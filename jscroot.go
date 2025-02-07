package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Definisi struktur folder
	dirs := []string{
		"myApp/www",
		"myApp/android/app",
	}

	// Membuat folder
	for _, dir := range dirs {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			fmt.Printf("Gagal membuat folder %s: %v\n", dir, err)
			return
		}
		fmt.Println("Folder dibuat:", dir)
	}

	// Membuat file capacitor.config.json
	configContent := `{
  "appId": "com.example.myapp",
  "appName": "My Capacitor App",
  "webDir": "www",
  "bundledWebRuntime": false
}`
	err := ioutil.WriteFile("myApp/capacitor.config.json", []byte(configContent), 0644)
	if err != nil {
		fmt.Println("Gagal membuat capacitor.config.json:", err)
		return
	}
	fmt.Println("File capacitor.config.json dibuat.")

	// Membuat file HTML awal
	htmlContent := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Capacitor App</title>
</head>
<body>
    <h1>Welcome to Capacitor with Go</h1>
</body>
</html>`
	err = ioutil.WriteFile("myApp/www/index.html", []byte(htmlContent), 0644)
	if err != nil {
		fmt.Println("Gagal membuat index.html:", err)
		return
	}
	fmt.Println("File index.html dibuat.")

	fmt.Println("Proyek Capacitor siap digunakan!")
}
