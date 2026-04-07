package main

import "fmt"

func tanggunganHari(jumlahHari int, tujuan string) int {
	
	if tujuan == "domestik" && jumlahHari > 7 {
		return 7
	} else if tujuan == "mancanegara" && jumlahHari > 3 {
		return 3
	}
	return jumlahHari
}
func biayaPerHari(jumlahMhs int) int {
	return jumlahMhs * 100000
}
func perhitunganBiaya(jumlahMhs, lamaPerjalanan int, tujuan string, totalBiaya *float64) {
	hariDitanggung := tanggunganHari(lamaPerjalanan, tujuan)
	biayaDomestik := biayaPerHari(jumlahMhs)
	if tujuan == "domestik" {
		*totalBiaya = float64(hariDitanggung * biayaDomestik)
	} else if tujuan == "mancanegara" {
		*totalBiaya = float64(hariDitanggung) * (float64(biayaDomestik) * 1.5)
	}
}
func main() {
	var jumlah, lama int
	var tujuan string
	var biaya float64
	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&jumlah)
	fmt.Print("Masukkan lama perjalanan (hari): ")
	fmt.Scan(&lama)
	fmt.Print("Masukkan tujuan (domestik/mancanegara): ")
	fmt.Scan(&tujuan)
	perhitunganBiaya(jumlah, lama, tujuan, &biaya)
	fmt.Printf("Total biaya yang ditanggung Tel-U: %.0f\n", biaya)
}