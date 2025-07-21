# Soal-1 CLI Validasi Umur -Tugas Pekan 2-

CLI interactive programs for input data validation from users

## How To Run

`go run main.go`
```
Nama: 
```

Input user `Nama` and press enter
```
Nama: Fajar
Umur:
```

Input user `Umur` and press enter
```
Nama: Fajar
Umur: 18 
```

## Output
```
Nama: Fajar
Umur: 18
Selamat datang, Fajar!
```

if the user `Nama` is empty
```
Nama:
Umur: 18
Error: Isi nama Anda!
```

if the user `Umur` is empty or less than 18
```
Nama: Fajar
Umur: 17
Error: Umur tidak valid (minimal 18 tahun)
```