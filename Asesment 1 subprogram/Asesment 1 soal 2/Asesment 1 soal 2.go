package main

import "fmt"

func tanggunganHari(jumlahHari int, tujuan string) int {
	if tujuan == "domestik" && jumlahHari > 3 {
		return 3
	} else if tujuan == "mancanegara" && jumlahHari > 8 {
		return 8
	}
	return jumlahHari
}
func biayaPerHari(jumlahMhs int) int {
	return 620000 * jumlahMhs
}
func perhitunganBiaya(jumlahMhs, lamaPerjalanan int, tujuan string, totalBiaya *float64) {
	hariEfektif := tanggunganHari(lamaPerjalanan, tujuan)
	biayaDomestik := biayaPerHari(jumlahMhs)
	*totalBiaya = float64(hariEfektif * biayaDomestik)
	if tujuan == "mancanegara" {
		*totalBiaya *= 1.5
	}
}
func main() {
	var jumlah, lama int
	var tujuan string
	var biaya float64
	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&jumlah)
	fmt.Print("Masukkan lama hari study tour: ")
	fmt.Scan(&lama)
	fmt.Print("Masukkan tujuan (domestik/mancanegara): ")
	fmt.Scan(&tujuan)
	perhitunganBiaya(jumlah, lama, tujuan, &biaya)
	fmt.Printf("\nBiaya perjalanan yang harus dikeluarkan Tel-U : Rp. %.0f\n", biaya)
}