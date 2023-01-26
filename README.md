# belajar-golang-dasar

## Content
- [Package](#Package)

# Package
- [Initilization](#Initilization)
- [OS](#OS)
- [Flag](#Flag)
- [Strings](#Strings)
- [Strconv](#Strconv)
- [Math](#Math)
- [List](#List)
- [Ring](#Ring)
- [Sort](#Sort)
- [Time](#Time)
- [Reflect](#Reflect)
- [RegExp](#RegExp)
# Initilization
- Saat kita membuat package, kita bisa membuat sebuah function yang akan diakses ketika package kita diakses
-Ini sangat cocok ketika contohnya, jika package kita berisi function-function untuk berkomunikasi dengan database, kita membuat function inisialisasi untuk membuka koneksi ke database
- Untuk membuat function yang diakses secara otomatis ketika package diakses, kita cukup membuat function dengan nama init

# OS
- Package os berisikan fungsionalitas untuk mengakses fitur sistem operasi secara independen (bisa digunakan  disemua sistem operasi)
```go
https://golang.org/pkg/os/
```

# Flag
- Package flag berisikan fungsionalitas untuk memparsing command line argument
```go
https://golang.org/pkg/flag/
```

# Strings
- Package strings adalah package yang berisikan function-function untuk memanipulasi tipe data String
```go
https://golang.org/pkg/strings/
```

# Strconv
- Sebelumnya kita sudah belajar cara konversi tipe data, misal dari int32 ke int34
- Bagaimana jika kita butuh melakukan konversi yang tipe datanya berbeda? Misal dari int ke string, atau sebaliknya
- Hal tersebut bisa kita lakukan dengan bantuan package strconv (string conversion)
```go
https://golang.org/pkg/strconv/
```

# Math
- Package math merupakan package yang berisikan constant dan fungsi matematika
```go
https://golang.org/pkg/math/
```

# List
- Package container/list adalah implementasi struktur data double linked list di Go-Lang
```go
https://golang.org/pkg/container/list/
```

# Ring
- Package container/ring adalah implementasi struktur data circular list
- Circular list adalah struktur data ring, dimana diakhir element akan kembali ke element awal (HEAD)
```go
https://golang.org/pkg/container/ring/
```

# Sort
- Package sort adalah package yang berisikan utilitas untuk proses pengurutan
- Agar data kita bisa diurutkan, kita harus mengimplementasikan kontrak di interface sort.Interface
```go
https://golang.org/pkg/sort/
```

# Time
- Package time adalah package yang berisikan fungsionalitas untuk management waktu di Go-Lang
```go
https://golang.org/pkg/time/
```

# Reflect
- Dalam bahasa pemrograman, biasanya ada fitur Reflection, dimana kita bisa melihat struktur kode kita pada saat aplikasi sedang berjalan
- Hal ini bisa dilakukan di Go-Lang dengan menggunakan package reflect
- Fitur ini mungkin tidak bisa dibahas secara lengkap dalam satu video, Anda bisa eksplorasi package reflec ini secara otodidak
- Reflection sangat berguna ketika kita ingin membuat library yang general sehingga mudah digunakan
```go
https://golang.org/pkg/reflect/
```

# RegExp
- Package regexp adalah utilitas di Go-Lang untuk melakukan pencarian regular expression
- Regular expression di Go-Lang menggunakan library C yang dibuat Google bernama RE2
```go
https://github.com/google/re2/wiki/Syntax
https://golang.org/pkg/regexp/
```