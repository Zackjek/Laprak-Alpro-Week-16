package main

import (
    "fmt"
)

func main() {
    var jumlah, hitungan float64
    var masukan float64

    fmt.Println("Masukkan bilangan riil (akhiri dengan 9999):")
    for {
        fmt.Scan(&masukan)
        if masukan == 9999 {
            break
        }
        jumlah += masukan
        hitungan++
    }

    if hitungan > 0 {
        fmt.Printf("Rata-rata: %.2f\n", jumlah/hitungan)
    } else {
        fmt.Println("Tidak ada bilangan yang dimasukkan.")
    }
}
