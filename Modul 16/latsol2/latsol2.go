package main

import "fmt"

func main() {
    var kata, s1, s2, s3, s4, s5 string
    var jumlah int

    fmt.Print("Masukkan string kata: ")
    fmt.Scanln(&kata)

    fmt.Print("Masukkan jumlah string jumlah (Max 5 ya): ")
    fmt.Scanln(&jumlah)

    fmt.Println("Masukkan string:")
    if jumlah >= 1 {
        fmt.Scanln(&s1)
    }
    if jumlah >= 2 {
        fmt.Scanln(&s2)
    }
    if jumlah >= 3 {
        fmt.Scanln(&s3)
    }
    if jumlah >= 4 {
        fmt.Scanln(&s4)
    }
    if jumlah >= 5 {
        fmt.Scanln(&s5)
    }

    ditemukan := false
    posisi := -1
    hitung := 0

    if jumlah >= 1 && s1 == kata {
        hitung++
        if !ditemukan {
            ditemukan = true
            posisi = 1
        }
    }
    if jumlah >= 2 && s2 == kata {
        hitung++
        if !ditemukan {
            ditemukan = true
            posisi = 2
        }
    }
    if jumlah >= 3 && s3 == kata {
        hitung++
        if !ditemukan {
            ditemukan = true
            posisi = 3
        }
    }
    if jumlah >= 4 && s4 == kata {
        hitung++
        if !ditemukan {
            ditemukan = true
            posisi = 4
        }
    }
    if jumlah >= 5 && s5 == kata {
        hitung++
        if !ditemukan {
            ditemukan = true
            posisi = 5
        }
    }
    fmt.Println()

    if ditemukan {
        fmt.Println("a. String kata ada dalam kumpulan jumlah data string tersebut.")
    } else {
        fmt.Println("a. String kata tidak ada dalam kumpulan jumlah data string tersebut.")
    }

    if ditemukan {
        fmt.Printf("b. String kata ditemukan di posisi ke-%d.\n", posisi)
    } else {
        fmt.Println("b. String kata tidak ditemukan.")
    }

    fmt.Printf("c. String kata muncul sebanyak %d kali dalam kumpulan jumlah data string tersebut.\n", hitung)

    if hitung >= 2 {
        fmt.Println("d. Terdapat sedikitnya dua string kata dalam jumlah data string tersebut.")
    } else {
        fmt.Println("d. Tidak terdapat sedikitnya dua string kata dalam jumlah data string tersebut.")
    }
}
