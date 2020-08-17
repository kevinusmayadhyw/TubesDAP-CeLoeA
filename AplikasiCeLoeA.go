package main

/** 	DATA KELOMPOK */
/* 		Judul			 : Aplikasi CeLoe - A
Anggota Kelompok : 	- 1301194011 Kevin Usmayadhy Wijaya
					- 1301194133 Mohammad Akbar Fauzy Al
*/
import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

const N = 20
const Nsoal = 10

type users struct {
	username, password, NIM string
	role                    int
}

type arrUsers = [N]users

type assignments struct {
	tittle, subject            string
	soal                       [Nsoal]string
	peserta, tipe, jumSoal, id int
	nimPeserta                 [100]string
}

type arrAssignment = [N]assignments

type jawabanAssignment struct {
	nimPeserta string
	idSoal     int
	jawaban    [Nsoal]string
	nilai      [Nsoal]int
	nilaiTot   int
}

type nilaiType struct {
	nomor int
	nilai float64
}

type arrJawaban = [N]jawabanAssignment

type subjects = [100]string

type arrayforum = [N]tipeforum

type tipeforum struct {
	nama         string
	text         [N]string
	banyakchat   int
	pembuatforum int
}

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func scanner() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	return line
}

func login(dataUsers arrUsers, jumUser int, name string, pass string, role int, activeUser *int, logged *bool) {
	// Menerima name dan pass dengan tipe data string kemudian
	// mengembalikan true jika name DAN pass sesusai dengan data di index tertentu
	var i int

	for i < jumUser && !*logged {
		*logged = dataUsers[i].username == name && dataUsers[i].password == pass && dataUsers[i].role == role
		if *logged {
			*activeUser = i
		}
		i++
	}
}

func logout(activeUser *int, logged *bool) {
	*logged = false
	*activeUser = -1
}

func register(dataUsers *arrUsers, jumUser *int, registered *bool) {
	//Menerima dataUsers = data pengguna, jumUser = jumlah pengguna
	//Mengembalikan dataUsers, dan jumUser yange telah ditambah datanya dan registered yang bertipe boolean yang akan mengembalikan true jika sukses

	var (
		username, password, nim string
		i                       int
		hasError                bool
	)

	fmt.Print(" |\tMasukan Username : ")
	fmt.Scan(&username)
	fmt.Print(" |\tMasukan Password : ")
	fmt.Scan(&password)
	for i < *jumUser {
		if dataUsers[i].username == username && dataUsers[i].password == password {
			clear()
			header(&hasError)
			fmt.Println(" |\tMohon maaf Username/Passwrod tidak bisa dipakai, Mohon gunakan usern+ame/password yang lain")
			fmt.Print(" |\tMasukan Username : ")
			fmt.Scan(&username)
			fmt.Print(" |\tMasukan Password : ")
			fmt.Scan(&password)
			i = -1
		}
		i++
	}

	i = 0

	fmt.Print(" |\tMasukan NIM anda :  ")
	fmt.Scan(&nim)
	for i < *jumUser {
		if dataUsers[i].NIM == nim {
			fmt.Scan(&nim)
			i = -1
		}
		i++
	}

	dataUsers[*jumUser].username = username
	dataUsers[*jumUser].password = password
	dataUsers[*jumUser].NIM = nim
	dataUsers[*jumUser].role = 2
	*jumUser++
	*registered = true
}

func header(hasError *bool) {
	fmt.Println("=========================================================================================")
	fmt.Println("===        ===       ====   =======         ==       =======         ======     =     ===")
	fmt.Println("===   ========   ========   =======   ===   ==   ===========   ===   =========== ========")
	fmt.Println("===   ========       ====   =======   ===   ==       ===  ==         =======   ===   ====")
	fmt.Println("===   ========   ========   =======   ===   ==   ===========   ===   =======   ===   ====")
	fmt.Println("===        ===        ===        ==         ==       =======   ===   ========       =====")
	fmt.Println("=========================================================================================")
}

func footer(input *string) {
	fmt.Println("==========================================================================================")
	fmt.Print(" |\tMasukan No Menu : ")
	fmt.Scan(&*input)
}

func footerInt(input *int) {
	fmt.Println("==========================================================================================")
	fmt.Print(" |\tMasukan No Menu : ")
	fmt.Scan(&*input)
}

func chooseSubject(arrSubject subjects, jumSubject int) string {
	var i, input int
	var hasError bool

	for input <= 0 || input > jumSubject {
		clear()
		header(&hasError)
		i = 0
		for i < jumSubject {
			fmt.Println(" |\t", i+1, ". ", arrSubject[i])
			i++
		}
		fmt.Println("========================================================================================")
		fmt.Print(" |\tMasukan No Subject : ")
		fmt.Scan(&input)
	}

	return arrSubject[input-1]
}

func addSubject(arrSubject *subjects, jumSubject *int) {
	var i int
	var judul string
	i = 0
	fmt.Print(" |\tNama Mata Kuliah Baru : ")
	fmt.Scan(&judul)
	judul = judul + " " + scanner()
	for i < *jumSubject {
		if arrSubject[i] == judul {
			fmt.Print(" |\tNama Mata Kuliah sudah Ada, masukan nama Matakuliah lain : ")
			fmt.Scan(&judul)
			judul = scanner()
			i = -1
		}
		i++
	}
	arrSubject[*jumSubject] = judul
	*jumSubject++
}

func viewsubject(arrSubject subjects, jumSubject int) {
	var i int
	i = 0
	for i < jumSubject {
		fmt.Println(" |\t", i+1, ")\t", arrSubject[i])
		i++
	}
}

func deleteSubject(arrSubject *subjects, jumSubject *int, hapus int) {
	var i int
	i = hapus
	for i < *jumSubject {
		arrSubject[i-1] = arrSubject[i]
		i++
	}
	*jumSubject--

}

func addData(arrData *arrAssignment, jumAssignment *int, subjectName string, tipe int) {
	var i, varId int

	varId = *jumAssignment + 1
	fmt.Print(" |\tJUDUL : ")
	fmt.Scan(&arrData[*jumAssignment].tittle)
	arrData[*jumAssignment].tittle = arrData[*jumAssignment].tittle + " " + scanner()
	fmt.Print(" |\tSoal No.1 : ")
	arrData[*jumAssignment].soal[i] = scanner()
	i++
	for arrData[*jumAssignment].soal[i-1] != "end" && i < Nsoal { // TAMBAH SOAL
		fmt.Print(" |\tSoal No.", i+1, " : ")
		arrData[*jumAssignment].soal[i] = scanner()
		i++
	}
	arrData[*jumAssignment].subject = subjectName
	arrData[*jumAssignment].tipe = tipe
	arrData[*jumAssignment].jumSoal = i
	if arrData[*jumAssignment].soal[i-1] == "end" {
		arrData[*jumAssignment].jumSoal = i - 1
	}
	i = 0
	for i < *jumAssignment {
		if arrData[i].id == varId {
			varId = varId + 1
			i = -1
		}
		i++
	}
	arrData[*jumAssignment].id = varId
	*jumAssignment++
}

func editData(data *assignments, jumAssignment int, dataSubject subjects, jumSubject int, ubah string, arrSubject subjects) {
	var idx int
	if ubah == "1" { // UBAH JUDUL
		fmt.Println("Judul Lama : ", data.tittle)
		fmt.Print("Masukan Judul Baru : ")
		data.tittle = scanner() // Menghindari BUG
		data.tittle = scanner()
	} else if ubah == "2" { // UBAH MATA KULIAH
		fmt.Println("Mata Kuliah Lama : ", data.subject)
		fmt.Println("Pilih Mata Kuliah yang Baru : ")
		viewsubject(arrSubject, jumSubject)
		fmt.Scan(&idx)
		data.subject = arrSubject[idx-1]
	}
}

func editSoal(data *assignments, noSoal int) {
	var tempSoal, input string

	fmt.Println("Soal Lama : ", data.soal[noSoal])
	fmt.Print("Soal Baru : ")
	fmt.Scan(&tempSoal)
	tempSoal = tempSoal + " " + scanner()
	fmt.Println(" Data [Y = YES, N = NO]")
	fmt.Scan(&input)
	if input == "Y" {
		data.soal[noSoal] = tempSoal
	}

}

func deleteData(arrData *arrAssignment, jumAssignment *int, idxDelete int) {
	var i int
	var temp assignments

	arrData[idxDelete].tittle = ""
	arrData[idxDelete].subject = ""
	arrData[idxDelete].peserta = 0
	arrData[idxDelete].tipe = 0
	for i < arrData[idxDelete].jumSoal {
		arrData[idxDelete].soal[i] = ""
		i++
	}
	arrData[idxDelete].jumSoal = 0
	temp = arrData[idxDelete+1]
	for idxDelete < *jumAssignment {
		arrData[idxDelete] = temp
		idxDelete++
		temp = arrData[idxDelete+1]
	}

	*jumAssignment--

}

func viewData(arrData arrAssignment, jumAssignment int, keySubject string, tipe int) {
	var i int
	if jumAssignment == 0 {
		fmt.Println("Tidak ada Tugas/Quiz")
	} else {
		fmt.Println("  | Subject : ", keySubject)
		for i < jumAssignment {
			if (keySubject == "All Subject" || arrData[i].subject == keySubject) && arrData[i].tipe == tipe {
				fmt.Println("  |    ", i+1, ". Judul : ", arrData[i].tittle, "(Peserta : ", arrData[i].peserta, ")")
			}
			i++
		}
	}
}

func viewDetilData(data assignments) {
	var numSoal int

	fmt.Println(" |\tJudul : ", data.tittle)
	fmt.Println(" |\tSubject : ", data.subject)
	fmt.Println(" |\tPeserta : ", data.peserta)
	fmt.Println(" |\tSoal : ")
	for numSoal < data.jumSoal {
		fmt.Println(" |\t", numSoal+1, ". ", data.soal[numSoal])
		numSoal++
	}
}

func searchMhs(dataUsers arrUsers, jumUser int, key string) int {
	var middle, top, bottom int
	var found bool

	top = jumUser - 1
	bottom = 0

	for bottom <= top && !found {
		middle = (bottom + top) / 2
		if dataUsers[middle].NIM > key {
			bottom = middle + 1
		} else if dataUsers[middle].NIM < key {
			top = middle - 1
		} else {
			found = dataUsers[middle].NIM == key
		}
	}

	if found {
		return middle
	} else {
		fmt.Println("Data tidak Ditemukan")
		return -1
	}
}

func searchJawaban(dataJawaban arrJawaban, jumDataJawaban, idSoal int, key string) int {
	var i int
	var found bool

	for i < jumDataJawaban && !found {
		found = dataJawaban[i].nimPeserta == key && dataJawaban[i].idSoal == idSoal
		i++
	}
	if found {
		return i - 1
	} else {
		return -1
	}
}

func getJawaban(dataAssignment assignments, dataJawaban arrJawaban, hDataJawaban *arrJawaban, jumDataJawaban, tipe int) {
	var i, j, k int

	for i < dataAssignment.peserta { //PERULANGAN UNTUK MENYIMPAN JAWABAN SESUAI DENGAN ID SOAL
		for j < jumDataJawaban {
			if dataAssignment.id == dataJawaban[j].idSoal && dataAssignment.tipe == tipe {
				hDataJawaban[k] = dataJawaban[j]
				k++
			}
			j++
		}
		i++
	} // END PERULANGAN
}

func viewJawaban(jwb jawabanAssignment, jumSoal int) {
	var i int

	fmt.Println(" |\tJawaban : ")
	for i < jumSoal {
		fmt.Println(" |\t", i+1, ". ", jwb.jawaban[i], "(", jwb.nilai[i], ")")
		i++
	}
}

func ssort(dataJawaban *arrJawaban, n int, tugaske int) {
	/* IS. terdefinisi array data jawaban yang memiliki n buah bilangan bulat
	   FS. array data jawaban terurut ascending dengan cara selection sort berdasarkan nilaitotal*/
	var pass, idx_min, i int
	var tampung jawabanAssignment
	pass = 0
	i = 0
	for pass <= n-1 {
		idx_min = pass
		i = pass + 1
		for i <= n-1 {
			if dataJawaban[idx_min].nilaiTot > dataJawaban[i].nilaiTot {
				idx_min = i
			}
			i++
		}
		tampung = dataJawaban[pass]
		dataJawaban[pass] = dataJawaban[idx_min]
		dataJawaban[idx_min] = tampung
		pass++
	}
}

func isort(dataJawaban *arrJawaban, n int) {
	/* IS. terdefinisi array data Jawaban yang berisi n Bilangan bulat
	   FS. array dataJawaban terurut secara descending dengan menggunakan insertion sort berdasar nilai total*/
	var pass, i int
	var temp jawabanAssignment
	pass = 0
	i = 0
	for pass < n {
		i = pass + 1
		temp = dataJawaban[i]
		for i > 0 && temp.nilaiTot > dataJawaban[i-1].nilaiTot {
			dataJawaban[i] = dataJawaban[i-1]
			i--
		}
		dataJawaban[i] = temp
		pass++
	}
}

func isortUser(dataUsers *arrUsers, n int) {
	/* IS. terdefinisi array data Users yang berisi n Bilangan bulat
	   FS. array dataUsers terurut secara descending dengan menggunakan insertion sort berdasar NIM*/
	var pass, i int
	var temp users
	pass = 0
	i = 0
	for pass < n {
		i = pass + 1
		temp = dataUsers[i]
		for i > 0 && temp.NIM > dataUsers[i-1].NIM {
			dataUsers[i] = dataUsers[i-1]
			i--
		}
		dataUsers[i] = temp
		pass++
	}
}

func forumforum(forum *arrayforum, banyakforum int) {
	// IS. Menerima array forum sebanyak banyak forum //
	// FS. Menampilkan forum forum yang telah terdaftar //
	var iterasi int
	iterasi = 0
	for iterasi < banyakforum {
		fmt.Println(iterasi+1, ")", forum[iterasi].nama)
		iterasi++
	}
}

func chatforum(forum *arrayforum, idx *int, active int, dataUsers arrUsers, banyakforum int) {
	// IS. Membuka array forum dalam index ke idx //
	// FS. Menambah chat dalam array forum yang telah dibuka //
	var iterasi int
	var keluar bool
	fmt.Println(" |\tMasukan nomor forum :")
	fmt.Scan(&*idx)
	keluar = false
	for *idx != -1 && !keluar {
		for banyakforum < *idx {
			fmt.Print(" |\tMohon maaf forum belum terdaftar masukan ulang forum :")
			fmt.Scan(&*idx)
		}
		if *idx > 0 {
			iterasi = 0
			fmt.Println("=========", forum[*idx-1].nama, "===========")
			fmt.Println(" |\tKetik esc untuk keluar")

			for forum[*idx].banyakchat != 0 && forum[*idx].banyakchat > iterasi {
				fmt.Println(forum[*idx].text[iterasi])
				iterasi++
			}

			forum[*idx].text[forum[*idx].banyakchat] = scanner() // Untuk menghidari bug
			if forum[*idx].text[forum[*idx].banyakchat] == "esc" {
				keluar = true
			}
			fmt.Print(dataUsers[active].username, " : ")
			forum[*idx].text[forum[*idx].banyakchat] = scanner()
			if forum[*idx].text[forum[*idx].banyakchat] == "esc" {
				keluar = true
			}

			for forum[*idx].text[forum[*idx].banyakchat] != "esc" {
				forum[*idx].text[forum[*idx].banyakchat] = dataUsers[active].username + " : " + forum[*idx].text[forum[*idx].banyakchat]
				forum[*idx].banyakchat++
				fmt.Print(dataUsers[active].username, " : ")
				forum[*idx].text[forum[*idx].banyakchat] = scanner()
				if forum[*idx].text[forum[*idx].banyakchat] == "esc" {
					keluar = true
				}
			}
		}
	}
}

func editforum(forum *arrayforum, idx int, active int, input *string) {
	// IS. Mengecek orang yang berwenang mengedit forum //
	// FS. Mengganti nama forum //
	if forum[idx-1].pembuatforum == active || active == 0 {
		fmt.Print(" |\tMasukan Nama forum :")
		forum[idx-1].nama = scanner() // Untuk menghindari bug
		forum[idx-1].nama = scanner()
	} else {
		fmt.Println(" |\tAnda Tidak dapat mengubah nama")
		*input = scanner() // Untuk menghindari bug
		*input = scanner()
	}
}

func hapusforum(forum *arrayforum, banyakforum *int, active int, idx int, input *string) {
	// IS. Mengecek orang yang berwenang menghapus forum //
	// FS. Menghapus array forum yang dipilih //
	var iterasi int
	iterasi = 0
	if forum[idx-1].pembuatforum == active || active == 0 {
		if idx == *banyakforum {
			forum[*banyakforum].banyakchat = 0
			*banyakforum--
		} else {
			for iterasi < *banyakforum {
				forum[idx-1] = forum[idx]
				idx++
				iterasi++
			}
			forum[*banyakforum].banyakchat = 0
			*banyakforum--
		}
	} else {
		fmt.Println(" |\tAnda Tidak dapat mengubah nama")
		*input = scanner() // Untuk menghindari bug
		*input = scanner()
	}
}

func main() {
	var (
		stop, checkPoint, checkPoint2, isLogin, registered, hasError, terjawab                                 bool
		tempInput, input, username, password, stringTipe, dummyVar                                             string
		dataUsers, tempUsers                                                                                   arrUsers
		jumUser, jumSubject, jumAssignment, jumDataJawaban, active, temp, idx, idxJawaban, i, banyakforum, mhs int
		dataSubject                                                                                            subjects
		dataAssignment                                                                                         arrAssignment
		dataJawaban, tempdataJawaban                                                                           arrJawaban
		tempJawaban                                                                                            jawabanAssignment
		forum                                                                                                  arrayforum
	)

	dataUsers[0].username = "admin"
	dataUsers[0].password = "admin"
	dataUsers[0].role = 1
	jumUser = 1
	dataSubject[0] = "All Subject"
	jumSubject = 1

	fmt.Print("Tekan Enter.......")
	fmt.Scanln(&input)
	clear()
	for !stop {

		for !checkPoint { // MENU PILIHAN =================================================
			header(&hasError)

			fmt.Println(" |\tPilih sebagai : ")
			fmt.Println(" |\t1. Dosen")
			fmt.Println(" |\t2. Mahasiswa")
			footer(&input)
			clear()
			if input == "1" || input == "2" {
				checkPoint = true
			}
		}

		tempInput = input
		checkPoint = false

		for !checkPoint && input != "9" {

			header(&hasError)

			fmt.Println(" |\t1. Login")

			if tempInput == "1" {
				temp = 1
			} else if tempInput == "2" {
				fmt.Println(" |\t2. Register")
				temp = 2
			}

			footer(&input)
			clear()

			header(&hasError)

			if input == "1" { // LOGIN -------------------------------------------------------------------------------------------------
				fmt.Print(" |\tMasukan Username : ")
				fmt.Scan(&username)
				fmt.Print(" |\tMasukan Password : ")
				fmt.Scan(&password)
				login(dataUsers, jumUser, username, password, temp, &active, &isLogin)
				if isLogin {
					checkPoint = true
				}

			} else if input == "2" && tempInput == "2" { // REGISTER ----------------------------------------------------------------------
				registered = false
				for !registered {
					register(&dataUsers, &jumUser, &registered)
					checkPoint = true
					clear()
				} // END REGISTERED lOOP
			}
		} // END CHECKPOINT FOR REGISTER AND LOGIN LOOP

		checkPoint = false

		for !checkPoint && isLogin { //CHECKPOINT IF LOGGED
			for !checkPoint { //CHECKPOINT FOR MAIN MENU
				clear()
				header(&hasError)

				fmt.Println(" |\t1. Menu Tugas")
				fmt.Println(" |\t2. Menu Quiz")
				fmt.Println(" |\t3. Forum")
				if dataUsers[active].role == 1 { // JIKA USER DOSEN
					fmt.Println(" |\t4. Menu Mata Kuliah")
				}
				fmt.Println(" |\t9. Logout")

				footer(&input)

				if input == "1" {
					stringTipe = "Tugas"
					temp = 1 // temp for assigmenmts type Tugas
					checkPoint = true
				} else if input == "2" {
					stringTipe = "Quiz"
					temp = 2 // temp for assigmenmts type Quiz
					checkPoint = true
				} else if input == "3" {
					checkPoint2 = false
					for !checkPoint2 {
						clear()
						header(&hasError)
						fmt.Println(" |\t1) Buat Forum")
						fmt.Println(" |\t2) Edit Nama Forum")
						fmt.Println(" |\t3) Hapus Forum")
						fmt.Println(" |\t4) Masuk Forum")
						fmt.Println(" |\t5) Kembali")
						footer(&input)
						clear()
						if input == "1" {
							forum[banyakforum].nama = scanner()
							fmt.Println(" |\tMasukan nama forum :")
							forum[banyakforum].nama = scanner()
							forum[banyakforum].pembuatforum = active
							forum[banyakforum].banyakchat = 0
							banyakforum++
						} else if input == "2" {
							forumforum(&forum, banyakforum)
							fmt.Println(" |\t(-1) untuk kembali")
							fmt.Print(" |\tForum yang ingin diedit :")
							fmt.Scan(&idx)
							for banyakforum < idx {
								fmt.Print(" |\tMohon maaf forum belum terdaftar masukan ulang forum :")
								fmt.Scan(&idx)
							}
							if banyakforum >= 1 && idx > 0 {
								editforum(&forum, idx, active, &input)
							}
						} else if input == "3" {
							forumforum(&forum, banyakforum)
							fmt.Println(" |\t(-1) untuk kembali")
							fmt.Print(" |\tForum yang ingin dihapus :")
							fmt.Scan(&idx)
							for banyakforum < idx {
								fmt.Print(" |\tMohon maaf forum belum terdaftar masukan ulang forum :")
								fmt.Scan(&idx)
							}
							if banyakforum >= 1 && idx > 0 {
								hapusforum(&forum, &banyakforum, active, idx, &input)
							}
						} else if input == "4" {
							forumforum(&forum, banyakforum)
							chatforum(&forum, &idx, active, dataUsers, banyakforum)
						} else if input == "5" {
							checkPoint2 = true
						}
					}
				} else if input == "4" && dataUsers[active].role == 1 {
					stringTipe = "Mata Kuliah"
					checkPoint = true
					temp = 3
				} else if input == "9" {
					logout(&active, &isLogin)
					checkPoint = true
					clear()
				}
			}

			checkPoint = false

			for !checkPoint && input != "9" {
				clear()
				header(&hasError)
				if dataUsers[active].role == 1 { // MENU DOSEN =================================================================
					fmt.Println(" |\t1. Tambah", stringTipe)
					fmt.Println(" |\t2. Edit", stringTipe)
					fmt.Println(" |\t3. Hapus", stringTipe)
					fmt.Println(" |\t4. Lihat", stringTipe)
					if temp != 3 {
						fmt.Println(" |\t5. Beri Nilai", stringTipe)
						fmt.Println(" |\t6. Lihat Urutan Nilai", stringTipe)
					}
					fmt.Println(" |\t9. Kembali")
					footer(&input)
					// ==============================================================================================================
					clear()
					header(&hasError)

					if input != "9" && temp != 3 { // PILIH SUBJECT/MATA KULIAH
						tempInput = chooseSubject(dataSubject, jumSubject)
						clear()
						header(&hasError)
					}

					// VARIABEL TEMPINPUT == SUBJECT YANG SUDAH DIPILIH ===============================================================

					if input == "1" { // TAMBAH DATA ===============================================================================
						if temp == 3 { // TAMBAH DATA UNTUK MATA KULIAH
							addSubject(&dataSubject, &jumSubject)
						} else { // TAMBAH DATA UNTUK TUGAS/QUIZ
							addData(&dataAssignment, &jumAssignment, tempInput, temp)
						}
					} else if input == "2" { // EDIT DATA ===============================================================================
						if temp == 3 {
							viewsubject(dataSubject, jumSubject)
							fmt.Println("\n |\tPilih Data yang mau di Edit : ")
							fmt.Scan(&idx)
							for idx > jumSubject && idx == 1 {
								fmt.Scan(&idx)
							}
							fmt.Println("Nama mata kuliah lama :", dataSubject[idx-1])
							fmt.Print("Masukan nama Mata kuliah baru :")
							dataSubject[idx-1] = scanner() // Menghindari Bug
							dataSubject[idx-1] = scanner()
						} else {
							viewData(dataAssignment, jumAssignment, tempInput, temp)
							fmt.Println("\n |\tPilih Data yang mau di Edit : ")
							fmt.Scan(&idx)
							if jumAssignment != 0 && dataAssignment[idx-1].tipe == temp && (dataAssignment[idx-1].subject == tempInput || tempInput == "All Subject") {
								viewDetilData(dataAssignment[idx-1])
								fmt.Println(" |\tPilih bagian yang mau di Ubah : ")
								fmt.Println(" |\t1. Judul")
								fmt.Println(" |\t2. Mata Kuliah")
								fmt.Println(" |\t3. Soal")
								footer(&input)
								if input == "1" || input == "2" {
									editData(&dataAssignment[idx-1], jumAssignment, dataSubject, jumSubject, input, dataSubject)
								} else if input == "3" {
									for input != "N" || input == "Y" {
										fmt.Println(" |\tPilih Soal yang mau diubah : ")
										fmt.Scan(&i)
										editSoal(&dataAssignment[idx-1], i-1)
										fmt.Print(" |\tIngin ubah Soal lain? [Y = YES, N = NO] : ")
										fmt.Scan(&input)
									}
								}
							}
						}
					} else if input == "3" { // DELETE DATA ===============================================================================
						if temp == 3 { // DELETE DATA SUBJECTS
							viewsubject(dataSubject, jumSubject)
							fmt.Println("\n |\tPilih Data yang mau di Edit : ")
							fmt.Scan(&idx)
							for idx > jumSubject && idx == 1 {
								fmt.Scan(&idx)
							}
							deleteSubject(&dataSubject, &jumSubject, idx)
						} else { // DELETE DATA TUGAS/QUIZ
							viewData(dataAssignment, jumAssignment, tempInput, temp)
							if jumAssignment != 0 { // JIKA DATA TIDAK KOSONG
								fmt.Println("\n |\tPilih Data yang mau di Hapus : ")
								fmt.Scan(&idx)
								if dataAssignment[idx-1].tipe == temp && (dataAssignment[idx-1].subject == tempInput || tempInput == "All Subject") {
									deleteData(&dataAssignment, &jumAssignment, idx-1)
								}
							} else { //JIKA DATA KOSONG
								fmt.Scan(&dummyVar)
							}
						}
					} else if input == "4" { // LIHAT DATA ===============================================================================
						if temp == 3 {
							viewsubject(dataSubject, jumSubject)
							fmt.Scan(&idx)
						} else {
							viewData(dataAssignment, jumAssignment, tempInput, temp)
							footerInt(&idx)
							if dataAssignment[idx-1].tipe == temp && (dataAssignment[idx-1].subject == tempInput || tempInput == "All Subject") {
								viewDetilData(dataAssignment[idx-1])
								fmt.Scan(&dummyVar)
							}
						}
					} else if input == "5" && temp != 3 { // BERI NILAI =======================================================================
						viewData(dataAssignment, jumAssignment, tempInput, temp)
						footerInt(&idx)
						if dataAssignment[idx-1].tipe == temp && (dataAssignment[idx-1].subject == tempInput || tempInput == "All Subject") {
							i = 0
							tempUsers = dataUsers
							isortUser(&tempUsers, jumUser)
							for i < dataAssignment[idx-1].peserta { // VIEW PESERTA
								mhs = searchMhs(tempUsers, jumUser, dataAssignment[idx-1].nimPeserta[i])
								idxJawaban = searchJawaban(dataJawaban, jumDataJawaban, dataAssignment[idx-1].id, tempUsers[mhs].NIM)
								fmt.Print(" |\tNAMA : ", tempUsers[mhs].username, "\t\t|\t\tNIM : ", dataAssignment[idx-1].nimPeserta[i])
								if dataJawaban[idxJawaban].nilaiTot == 0 {
									fmt.Print(" |\tNOT GRADED\n")
								} else {
									fmt.Println(" |\tNilai : ", dataJawaban[idxJawaban].nilaiTot, "\n")
								}
								i++
							}
							fmt.Print("\nMasukan NIM : ")
							fmt.Scan(&dummyVar)

							if searchMhs(tempUsers, jumUser, dummyVar) != -1 {
								for i != -1 {
									viewDetilData(dataAssignment[idx-1])
									idxJawaban = searchJawaban(dataJawaban, jumDataJawaban, dataAssignment[idx-1].id, dummyVar)
									if idxJawaban >= 0 {
										viewJawaban(dataJawaban[idxJawaban], dataAssignment[idx-1].jumSoal)
										fmt.Print(" |\tMasukan Nomor yang mau diberi Nilai [Masukan -1 jika sudah selesai] : ")
										fmt.Scan(&i)
										for i > dataAssignment[idx-1].jumSoal {
											fmt.Scan(&i)
										}
										if i != -1 {
											fmt.Print(" |\tNilai : ")
											fmt.Scan(&dataJawaban[idxJawaban].nilai[i-1])
										}
									}
								}
								i = 0
								dataJawaban[idxJawaban].nilaiTot = 0
								for i < dataAssignment[idx-1].jumSoal {
									dataJawaban[idxJawaban].nilaiTot = dataJawaban[idxJawaban].nilaiTot + dataJawaban[idxJawaban].nilai[i]
									i++
								}
							}
						}
					} else if input == "6" && temp != 3 {
						viewData(dataAssignment, jumAssignment, tempInput, temp)
						footerInt(&idx)
						checkPoint2 = false
						clear()
						for !checkPoint2 && dataAssignment[idx-1].tipe == temp {
							fmt.Println(" |\tPilih urutan yang mau di lihat : ")
							fmt.Println(" |\t1. Ascending")
							fmt.Println(" |\t2. Descending")
							fmt.Println(" |\t3. Kembali")
							footer(&input)
							getJawaban(dataAssignment[idx-1], dataJawaban, &tempdataJawaban, jumDataJawaban, temp)
							clear()
							if input == "1" {
								i = 0
								ssort(&tempdataJawaban, dataAssignment[idx-1].peserta, idx)
								for i < dataAssignment[idx-1].peserta {
									if dataAssignment[idx-1].tipe == temp {
										fmt.Println(" |\t", i+1, ")\t", "NIM :", tempdataJawaban[i].nimPeserta, "\t\t", "Nilai :", tempdataJawaban[i].nilaiTot)
									}
									i++
								}
							} else if input == "2" {
								i = 0
								isort(&tempdataJawaban, dataAssignment[idx-1].peserta)
								for i < dataAssignment[idx-1].peserta {
									if dataAssignment[idx-1].tipe == temp {
										fmt.Println(" |\t", i+1, ")\t", "NIM :", tempdataJawaban[i].nimPeserta, "\t\t", "Nilai :", tempdataJawaban[i].nilaiTot)
									}
									i++
								}
							} else if input == "3" {
								checkPoint2 = true
							}
						}
					}
					//==============================================================================================================================================================================================================

				} else if dataUsers[active].role == 2 { // MENU MAHASISWA ===============================================================
					fmt.Println(" |\t1.", stringTipe)
					fmt.Println(" |\t2. Lihat Nilai", stringTipe)
					fmt.Println(" |\t9. Kembali")
					footer(&input)

					if input != "9" {
						tempInput = chooseSubject(dataSubject, jumSubject)
						clear()
						header(&hasError)
					}

					if input == "1" {
						fmt.Println(" |\tPilih ", stringTipe, " yang mau di kerjakan : ")
						viewData(dataAssignment, jumAssignment, tempInput, temp)
						footerInt(&idx)
						for idx <= 0 {
							footerInt(&idx)
						}

						if dataAssignment[idx-1].tipe == temp && (dataAssignment[idx-1].subject == tempInput || tempInput == "All Subject") {
							viewDetilData(dataAssignment[idx-1])

							terjawab = false
							i = 0

							for i < jumDataJawaban && !terjawab { // MENGECEK JIKA USER SUDAH MENGERJAKAN TUGAS/QUIZ
								terjawab = dataJawaban[i].nimPeserta == dataUsers[active].NIM && dataAssignment[idx-1].id == dataJawaban[i].idSoal
								i++
							}

							if !terjawab { // JIKA USER BELUM MENGERJAKAN TUGAS/QUIZ

								tempJawaban.idSoal = dataAssignment[idx-1].id
								tempJawaban.nimPeserta = dataUsers[active].NIM
								idxJawaban = 0

								for idxJawaban != -1 {
									fmt.Print(" |\tPilih Nomor yang mau dijawab [Masukan -1 jika sudah selesai] : ")
									fmt.Scan(&idxJawaban)
									for idxJawaban < -1 || idxJawaban > Nsoal || idxJawaban == 0 {
										fmt.Scan(&idxJawaban)
									}
									if idxJawaban == -1 { // JIKA USER MAU MEN-SUBMIT JAWABAN
										fmt.Print(" |\tYakin ingin men-Submit jawaban ? [Jawaban tidak dapat diubah lagi Y = YES N = NO] : ")
										fmt.Scan(&dummyVar)
										if dummyVar == "N" {
											idxJawaban = 0
										}
									} else {
										fmt.Print(" |\t\tSoal : ", dataAssignment[idx-1].soal[idxJawaban-1], "\n |\t\tJawaban : ")
										fmt.Scan(&dummyVar)
										dummyVar = dummyVar + " " + scanner()
										tempJawaban.jawaban[idxJawaban-1] = dummyVar
									}
								}
								dataJawaban[jumDataJawaban] = tempJawaban
								jumDataJawaban++
								dataAssignment[idx-1].nimPeserta[dataAssignment[idx-1].peserta] = dataUsers[active].NIM
								dataAssignment[idx-1].peserta++
							} else { // JIKA USER SUDAH MENGERJAKAN TUGAS/QUIZ
								fmt.Println(" |\tAnda Sudah menyelesaikan Tugas/Quiz ini")
								fmt.Scan(&dummyVar)
							}
						}
					} else if input == "2" {
						fmt.Println(" |\tNilai ", stringTipe, " anda : ")
						i = 0
						terjawab = false

						for i < jumAssignment {
							idxJawaban = searchJawaban(dataJawaban, jumDataJawaban, dataAssignment[i].id, dataUsers[active].NIM)
							if dataAssignment[i].tipe == temp && idxJawaban != -1 {
								fmt.Print(" |\tJUDUL : ", dataAssignment[i].tittle)
								if dataJawaban[idxJawaban].nilaiTot == 0 {
									fmt.Print(" |\tNOT GRADED\n")
								} else {
									fmt.Println(" |\tNilai : ", dataJawaban[idxJawaban].nilaiTot, "\n")
								}
							}
							i++
						}
						footer(&dummyVar)

					}
				} //tempInput = input
			}
		} // END CHECKPOINT TO LOGOUT LOOP
	} // END PROGRAM LOOP
}
