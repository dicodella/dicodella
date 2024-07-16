package main

import (
	"bufio"
	"fmt"
	"os"
)

// Batas maksimal jumlah data
const NMAX int = 100

// Tipe data bentukan untuk menyimpan data pasien
type dataPasien struct {
	namaLengkap string
	username    string
	password    string
}

// Tipe data bentukan untuk menyimpan data dokter
type dataDokter struct {
	namaLengkap      string
	username         string
	password         string
	kategori         string
	jumlahKonsultasi int
}

// Tipe data bentukan untuk menyimpan data pertanyaan
type forum struct {
	pertanyaan string
	response   [NMAX]string
	replyBy    [NMAX]string
	tag        string
}

// Tipe data bentukan untuk menyimpan data konsultasi
type konsul struct {
	dokter   string
	question string
	answer   string
	nomor    int
	dok      string
}

// Tipe data bentukan untuk menyimpan data jumlah tag
type count struct {
	topik      string
	countTopik int
}

// Tipe alias untuk array of dataPasien, dataDokter, forum, konsultasi dan countTag dengan batas NMAX
type pasien [NMAX]dataPasien
type dokter [NMAX]dataDokter
type pertanyaan [NMAX]forum
type konsultasi [NMAX]konsul
type countTag [NMAX]count

// Variabel global yang akan sering digunakan
var input int
var input2 string
var dataPas pasien
var dataDok dokter
var dataPer pertanyaan
var dataCount countTag
var dataKonsul konsultasi
var nPas int = 0
var nDok int = 0
var nPer int = 0
var nCount int = 5
var nKonsul int = 0
var nReply int = 0
var role string
var recentUser string

func main() {
	startPage()
}

func startPage() {
	// Tampilan awal aplikasi
	fmt.Println(" --------------------------------------")
	fmt.Println("|           Selamat datang di          |")
	fmt.Println("|                Heyodoc               |")
	fmt.Println(" --------------------------------------")
	fmt.Println("1. Login                               ")
	fmt.Println("2. Belum memiliki akun? Daftar di sini ")
	fmt.Println("3. Lewati dan lihat forum              ")
	fmt.Println(" --------------------------------------")

	fmt.Print("Silahkan masukkan pilihan anda: ")
	fmt.Scan(&input)
	if input == 1 {
		loginPage()
	} else if input == 2 {
		fmt.Print("Anda akan mendaftar sebagai? (Dokter/Pasien): ")
		fmt.Scan(&role)
		if role == "Dokter" {
			registPageDokter(&dataDok, &nDok)
		} else if role == "Pasien" {
			registPagePasien(&dataPas, &nPas)
		} else {
			fmt.Println("⚠ Pilihan tidak valid. Silahkan daftar ulang.")
			startPage()
		}
	} else if input == 3 {
		viewForum(dataPer, nPer, nReply)
	} else {
		fmt.Println("⚠ Pilihan tidak valid.")
		startPage()
	}
}

func loginPage() {
	/* 	IS: Input username dan password terdefinisi untuk Log in
	FS: Log in berhasil jika username dan password sesuai dengan data pengguna */
	var usn, pw string
	var pas, dok bool
	fmt.Print("Masukkan username Anda: ")
	fmt.Scan(&usn)
	fmt.Print("Masukkan password Anda: ")
	fmt.Scan(&pw)
	pas = cekUsnP(dataPas, nPas, usn) && cekPwP(dataPas, nPas, pw)
	dok = cekUsnD(dataDok, nDok, usn) && cekPwD(dataDok, nDok, pw)
	if pas && !dok {
		fmt.Println(".·:¨༺     Login Berhasil     ༻¨:·.")
		menuPasien()
	} else if !pas && dok {
		fmt.Println(".·:¨༺     Login Berhasil     ༻¨:·.")
		menuDokter()
	} else {
		fmt.Println("⚠    Login Gagal    ⚠")
		loginPage()
	}
}

func registPageDokter(A *dokter, n *int) {
	/* 	IS: Input data nama lengkap, username, password dan kategori dokter spesialis
	  				terdefinisi untuk daftar akun sebagai dokter
				FS: Proses daftar berhasil dan data disimpan di dalam array dengan tipe bentukan dokter */
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	fmt.Println(" -------------------------------")
	fmt.Println("|          Daftar Akun          |")
	fmt.Println(" -------------------------------")

	fmt.Print("Nama Lengkap: ")
	scanner.Scan()
	A[*n].namaLengkap = scanner.Text()

	fmt.Print("Username: ")
	scanner.Scan()
	A[*n].username = scanner.Text()

	fmt.Print("Password: ")
	scanner.Scan()
	A[*n].password = scanner.Text()

	fmt.Print("Kategori (Umum/Anak/Mata/Psikiater/Gigi/Saraf): ")
	scanner.Scan()
	A[*n].kategori = scanner.Text()

	A[*n].jumlahKonsultasi = 0

	*n++
	fmt.Println(".·:¨༺   Akun berhasil dibuat   ༻¨:·.")
	startPage()
}

func registPagePasien(A *pasien, n *int) {
	/* 	IS: Input data nama lengkap, username dan password terdefinisi untuk daftar akun sebagai pasien
	FS: Proses daftar berhasil dan data disimpan di dalam array dengan tipe bentukan pasien */
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	fmt.Println(" -------------------------------")
	fmt.Println("|          Daftar Akun          |")
	fmt.Println(" -------------------------------")

	fmt.Print("Nama Lengkap: ")
	scanner.Scan()
	A[*n].namaLengkap = scanner.Text()

	fmt.Print("Username: ")
	scanner.Scan()
	A[*n].username = scanner.Text()

	fmt.Print("Password: ")
	scanner.Scan()
	A[*n].password = scanner.Text()

	*n++
	fmt.Println(".·:¨༺   Akun berhasil dibuat   ༻¨:·.")
	startPage()
}

func cekUsnP(P pasien, n int, x string) bool {
	// function untuk memeriksa username pasien untuk login
	for i := 0; i < n; i++ {
		if P[i].username == x {
			recentUser = P[i].namaLengkap
			return true
		}
	}
	return false
}

func cekUsnD(D dokter, n int, x string) bool {
	// function untuk memeriksa username dokter untuk login
	for i := 0; i < n; i++ {
		if D[i].username == x {
			recentUser = D[i].namaLengkap
			return true
		}
	}
	return false
}

func cekPwP(P pasien, n int, x string) bool {
	// function untuk memeriksa password pasien untuk login
	for i := 0; i < n; i++ {
		if P[i].password == x {
			recentUser = P[i].namaLengkap
			return true
		}
	}
	return false
}

func cekPwD(D dokter, n int, x string) bool {
	// function untuk memeriksa password dokter untuk login
	for i := 0; i < n; i++ {
		if D[i].password == x {
			recentUser = D[i].namaLengkap
			return true
		}
	}
	return false
}

func menuPasien() {
	/* 	IS: Tampilan main menu untuk user sebagai pasien dan input pilihan menu
	FS: memanggil function sesuai dengan input pilihan pada main menu
	*/
	fmt.Println(" --------------------------------------")
	fmt.Println("|              MAIN MENU               |")
	fmt.Println(" --------------------------------------")
	fmt.Println("1. Lihat Forum                         ")
	fmt.Println("2. Konsultasi Pribadi                  ")
	fmt.Println("3. Log out                             ")
	fmt.Println(" --------------------------------------")

	fmt.Print("Silahkan masukkan pilihan anda: ")
	fmt.Scan(&input)
	if input == 1 {
		viewForumPasien(dataPer, nPer, nReply)
	} else if input == 2 {
		konsultasiPribadiPasien(&dataKonsul, &nKonsul)
	} else if input == 3 {
		startPage()
	} else {
		fmt.Println("⚠ input tidak valid.")
	}
}

func menuDokter() {
	/* 	IS: Tampilan main menu untuk user sebagai dokter dan input pilihan menu
	FS: memanggil function sesuai dengan input pilihan pada main menu
	*/
	fmt.Println(" --------------------------------------")
	fmt.Println("|              MAIN MENU               |")
	fmt.Println(" --------------------------------------")
	fmt.Println("1. Lihat Forum                         ")
	fmt.Println("2. Konsultasi Pribadi                  ")
	fmt.Println("3. Log out                             ")
	fmt.Println(" --------------------------------------")

	fmt.Print("Silahkan masukkan pilihan anda: ")
	fmt.Scan(&input)
	if input == 1 {
		viewForumDokter(dataPer, nPer, nReply)
	} else if input == 2 {
		konsultasiPribadiDokter(&dataKonsul, nKonsul, recentUser)
	} else if input == 3 {
		startPage()
	} else {
		fmt.Println("⚠ input tidak valid.")
	}
}

func viewForum(A pertanyaan, n int, m int) {
	/* 	IS: Tampilan forum untuk user yang tidak log in serta input pilihan menu
				FS: memanggil function sesuai dengan input pilihan pada halaman fitur forum. User yang tidak log in
	      		hanya bisa melihat forum antara pasien dan dokter yang mendaftar. User harus log in
	          terlebih dahulu jika ingin menambahkan response
	*/
	fmt.Println(" --------------------------------------")
	fmt.Println("|                 FORUM                |")
	fmt.Println(" --------------------------------------")
	fmt.Println("1. Tampilkan seluruh pertanyaan        ")
	fmt.Println("2. Ajukan Pertanyaan pada Forum        ")
	fmt.Println("3. Cari pertanyaan berdasarkan tag     ")
	fmt.Println("4. Kembali ke Main Menu   					    ")
	fmt.Println(" --------------------------------------")

	fmt.Print("Silahkan masukkan pilihan anda: ")
	fmt.Scan(&input)

	if input == 1 {
		for i := 0; i < n; i++ {
			fmt.Printf("%d. Pertanyaan: %s", i+1, A[i].pertanyaan)
			fmt.Println()
			for j := 0; j < m; j++ {
				fmt.Printf("reply by %s : %s", A[i].replyBy[j], A[i].response[j])
				fmt.Println()
			}
		}
		fmt.Print("Tambahkan response? (Ya/Tidak): ")
		fmt.Scan(&input2)
		if input2 == "Ya" {
			fmt.Println("Silahkan login terlebih dahulu.")
			startPage()
		} else if input2 == "Tidak" {
			fmt.Print("Kembali ke menu forum? Ya untuk melanjutkan: ")
			fmt.Scan(&input2)
			if input2 == "Ya" {
				viewForum(dataPer, nPer, nReply)
			}
		}
	} else if input == 2 {
		askQuestion(&dataPer, &nPer)
	} else if input == 3 {
		searchTag(dataPer, nPer)
	} else if input == 4 {
		menuPasien()
	} else {
		fmt.Println("⚠ input tidak valid.")
	}
}

func viewForumPasien(A pertanyaan, n int, m int) {
	/* 	IS: Tampilan forum untuk user sebagai pasien serta input pilihan menu
				FS: memanggil function sesuai dengan input pilihan pada halaman fitur forum. User pasien dapat
	      		menampilkan seluruh pertanyaan forum, bertanya pada forum, meresponse serta mencari pertanyaan
	          berdasarkan tag
	*/
	fmt.Println(" --------------------------------------")
	fmt.Println("|                 FORUM                |")
	fmt.Println(" --------------------------------------")
	fmt.Println("1. Tampilkan seluruh pertanyaan        ")
	fmt.Println("2. Ajukan Pertanyaan pada Forum        ")
	fmt.Println("3. Cari pertanyaan berdasarkan tag     ")
	fmt.Println("4. Kembali ke Main Menu   		        ")
	fmt.Println(" --------------------------------------")

	fmt.Print("Silahkan masukkan pilihan anda: ")
	fmt.Scan(&input)

	if input == 1 {
		for i := 0; i < n; i++ {
			fmt.Printf("%d. Pertanyaan: %s", i+1, A[i].pertanyaan)
			fmt.Println()
			for j := 0; j < m; j++ {
				fmt.Printf("reply by %s : %s", A[i].replyBy[j], A[i].response[j])
				fmt.Println()
			}
		}
		fmt.Print("Tambahkan response? (Ya/Tidak): ")
		fmt.Scan(&input2)
		if input2 == "Ya" {
			responsePasien(&dataPer, &nReply)
		} else if input2 == "Tidak" {
			fmt.Print("Kembali ke menu forum? Ya untuk melanjutkan: ")
			fmt.Scan(&input2)
			if input2 == "Ya" {
				viewForumPasien(dataPer, nPer, nReply)
			} else {
				fmt.Println("⚠ input tidak valid.")
			}
		}
	} else if input == 2 {
		askQuestion(&dataPer, &nPer)
	} else if input == 3 {
		searchTag(dataPer, nPer)
	} else if input == 4 {
		menuPasien()
	} else {
		fmt.Println("⚠ input tidak valid.")
	}
}

func viewForumDokter(A pertanyaan, n int, m int) {
	/* 	IS: Tampilan forum untuk user sebagai dokter serta input pilihan menu
				FS: memanggil function sesuai dengan input pilihan pada halaman fitur forum. User dokter dapat
	      		menampilkan seluruh pertanyaan forum dan melihat tag terurut descending sesuai frekuensi tag
	*/
	fmt.Println(" --------------------------------------")
	fmt.Println("|                 FORUM                |")
	fmt.Println(" --------------------------------------")
	fmt.Println("1. Tampilkan seluruh pertanyaan        ")
	fmt.Println("2. Lihat tag                           ")
	fmt.Println("3. Kembali ke Main Menu   			    ")
	fmt.Println(" --------------------------------------")

	fmt.Print("Silahkan masukkan pilihan anda: ")
	fmt.Scan(&input)

	if input == 1 {
		for i := 0; i < n; i++ {
			fmt.Printf("%d. Pertanyaan: %s", i+1, A[i].pertanyaan)
			fmt.Println()
			for j := 0; j < m; j++ {
				fmt.Printf("reply by %s : %s", A[i].replyBy[j], A[i].response[j])
			}
		}
		fmt.Print("Tambahkan response? (Ya/Tidak): ")
		fmt.Scan(&input2)
		if input2 == "Ya" {
			responseDokter(&dataPer, &nReply)
		} else if input2 == "Tidak" {
			fmt.Print("Kembali ke menu forum? Ya untuk melanjutkan: ")
			fmt.Scan(&input2)
			if input2 == "Ya" {
				viewForumDokter(dataPer, nPer, nReply)
			} else {
				fmt.Println("⚠ input tidak valid.")
			}
		}
	} else if input == 2 {
		topik(dataPer, &dataCount, nPer)
		searchTag(dataPer, nPer)
	} else if input == 3 {
		menuDokter()
	} else {
		fmt.Println("⚠ input tidak valid.")
	}
}

func searchTag(A pertanyaan, n int) {
	/* 	IS: input kategori tag Umum/Anak/Mata/Psikiater/Gigi/Saraf terdefinisi
				FS: menampilkan array pertanyaan dan jawaban dengan kategori tag berdasarkan input
	      		kategori tag
	*/
	var x string

	fmt.Println("Silahkan masukkan tag dari pertanyaan yang ingin Anda tampilkan ")
	fmt.Println("(Umum/Anak/Mata/Psikiater/Gigi/Saraf): ")
	fmt.Scan(&x)
	for i := 0; i < n; i++ {
		if A[i].tag == x {
			fmt.Printf("%d. Pertanyaan: %s", i+1, A[i].pertanyaan)
			fmt.Println()
			fmt.Printf("Jawaban: %s", A[i].response)
			fmt.Println()
		}
	}
	fmt.Print("Apakah Anda ingin mencari pertanyaan yang lain? (Ya/Tidak): ")
	fmt.Scan(&input2)
	if input2 == "Ya" {
		searchTag(dataPer, nPer)
	} else if input2 == "Tidak" {
		if role == "Pasien" {
			viewForumPasien(dataPer, nPer, nReply)
		} else if role == "Dokter" {
			viewForumDokter(dataPer, nPer, nReply)
		}
	}
}

func askQuestion(A *pertanyaan, n *int) {
	/* 	IS: function untuk user pasien menambahkan pertanyaan pada forum, user pasien memasukkan
	  				input pertanyaan dan tag sesuai dengan kategori pertanyaan
				FS: pertanyaan dan tag yang tersimpan pada array pertanyaan bertipe forum
	*/
	var pilih string
	var T forum

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	fmt.Println("Hai, silahkan masukkan pertanyaan Anda.")
	scanner.Scan()
	T.pertanyaan = scanner.Text()
	T.tag = addTag()
	A[*n] = T
	*n++
	fmt.Print("Apakah anda akan bertanya lagi? yes/no: ")
	fmt.Scan(&pilih)
	if pilih == "yes" {
		askQuestion(&dataPer, &nPer)
	} else if pilih == "no" {
		viewForumPasien(dataPer, nPer, nReply)
	} else {
		fmt.Println("⚠ input tidak valid.")
	}
}

func addTag() string {
	/* 	IS: input kategori tag Umum/Anak/Mata/Psikiater/Gigi/Saraf terdefinisi
	FS: mengembalikan nilai t berisi kategori tag
	*/
	var t string
	fmt.Print("Tambahkan tag (Umum/Anak/Mata/Psikiater/Gigi/Saraf): ")
	fmt.Scan(&t)
	return t
}

func topik(A pertanyaan, B *countTag, n int) {
	/* 	IS: array pertanyaan untuk menghitung frekuensi tag
				FS: menampilkan array frekuensi tag secara descending dengan memanggil
	      		function sortTag
	*/
	B[0].topik = "Umum"
	B[1].topik = "Anak"
	B[2].topik = "Mata"
	B[3].topik = "Psikiater"
	B[4].topik = "Gigi"
	B[5].topik = "Saraf"
	for i := 0; i < n; i++ {
		if A[i].tag == "Umum" {
			B[0].countTopik++
		} else if A[i].tag == "Anak" {
			B[1].countTopik++
		} else if A[i].tag == "Mata" {
			B[2].countTopik++
		} else if A[i].tag == "Psikiater" {
			B[3].countTopik++
		} else if A[i].tag == "Gigi" {
			B[4].countTopik++
		} else if A[i].tag == "Saraf" {
			B[5].countTopik++
		}
	}
	sortTag(&dataCount, nCount)
	for i := 0; i <= 5; i++ {
		fmt.Printf("Terdapat %d topik tentang %s\n", B[i].countTopik, B[i].topik)
	}
}

func sortTag(B *countTag, n int) {
	/* 	IS: array countTag terdefinisi
	FS: mengubah array countTag terurut secara descending berdasarkan frekuensi tag
	*/
	for i := 0; i < n; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if B[j].countTopik > B[maxIdx].countTopik {
				maxIdx = j
			}
		}
		B[i], B[maxIdx] = B[maxIdx], B[i]
	}
}

func responseDokter(A *pertanyaan, n *int) {
	/* 	IS: input response pada pertanyaan forum
	FS: array berisi response terhadap pertanyaan
	*/
	fmt.Print("Pertanyaan mana yang akan Anda respons?: ")
	fmt.Scan(&input)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	fmt.Println("Silahkan masukkan tanggapan Anda.")
	scanner.Scan()
	A[input-1].response[*n] = scanner.Text()
	A[input-1].replyBy[*n] = recentUser
	fmt.Println(".·:¨༺   Response berhasil ditambahkan   ༻¨:·.")
	*n++
	fmt.Print("Apakah Anda ingin menambahkan respons lagi? (Ya/Tidak): ")
	fmt.Scan(&input2)
	if input2 == "Ya" {
		responseDokter(&dataPer, &nPer)
	} else if input2 == "Tidak" {
		viewForumDokter(dataPer, nPer, nReply)
	} else {
		fmt.Println("⚠ input tidak valid.")
	}
}

func responsePasien(A *pertanyaan, n *int) {
	/* 	IS: input response pada pertanyaan forum
	FS: array berisi response terhadap pertanyaan
	*/
	fmt.Print("Pertanyaan mana yang akan Anda respons?: ")
	fmt.Scan(&input)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	fmt.Println("Silahkan masukkan tanggapan Anda.")
	scanner.Scan()
	A[input-1].response[*n] = scanner.Text()
	A[input-1].replyBy[*n] = recentUser
	fmt.Println(".·:¨༺   Response berhasil ditambahkan   ༻¨:·.")
	*n++
	fmt.Print("Apakah Anda ingin menambahkan respons lagi? (Ya/Tidak): ")
	fmt.Scan(&input2)
	if input2 == "Ya" {
		responsePasien(&dataPer, &nPer)
	} else if input2 == "Tidak" {
		viewForumPasien(dataPer, nPer, nReply)
	} else {
		fmt.Println("⚠ input tidak valid.")
	}
}

func konsultasiPribadiDokter(A *konsultasi, n int, x string) {
	scanner := bufio.NewScanner(os.Stdin)

	j := 1
	for i := 0; i < n; i++ {
		if A[i].dokter == x {
			fmt.Println(j, ".", A[i].question)
			fmt.Println("Berikan jawaban Anda:")
			scanner.Scan()
			if scanner.Scan() {
				A[i].answer = scanner.Text()
			}
			for k := 0; k < nDok; k++ {
				if dataDok[k].namaLengkap == x {
					dataDok[k].jumlahKonsultasi++
				}
			}
			j++
		}
	}
	fmt.Println("Anda sudah menjawab semua konsultasi pribadi pasien.")
	fmt.Print("Kembali ke menu? Ya untuk melanjutkan: ")
	fmt.Scan(&input2)
	if input2 == "Ya" {
		menuDokter()
	} else {
		fmt.Println("⚠ input tidak valid.")
	}
}

func konsultasiPribadiPasien(A *konsultasi, n *int) {
	fmt.Println("Cari dokter spesialis untuk konsultasi pribadi Anda")
	fmt.Println("(Umum/Anak/Mata/Psikiater/Gigi/Saraf): ")
	fmt.Scan(&input2)
	sortDokter(&dataDok, nDok)
	spesialis(dataDok, &dataKonsul, nDok, input2)
	fmt.Println("Pilih dokter :")
	fmt.Scan(&input)
	A[*n].dokter = A[input-1].dok
	fmt.Println("Silahkan tanyakan sesuatu pada dokter pribadi Anda.")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Scan() {
		(*A)[*n].question = scanner.Text()
	}
	fmt.Println(".·:¨༺   Pertanyaan Anda telah dikirimkan   ༻¨:·.")
	*n++
	fmt.Print("Apakah Anda ingin menambahkan pertanyaan lagi? (Ya/Tidak): ")
	fmt.Scan(&input2)
	if input2 == "Ya" {
		konsultasiPribadiPasien(A, n)
	} else if input2 == "Tidak" {
		menuPasien()
	} else {
		fmt.Println("⚠ input tidak valid.")
	}
}

func spesialis(D dokter, E *konsultasi, n int, x string) {
	/* 	IS: input kategori tag Umum/Anak/Mata/Psikiater/Gigi/Saraf terdefinisi
	FS: menampilkan nama dokter spesialis sesuai dengan kategori tag yang dicari
	*/
	j := 1
	for i := 0; i < n; i++ {
		if D[i].kategori == x {
			fmt.Println(j, ".", D[i].namaLengkap)
			fmt.Printf("%d pasien telah terbantu oleh dokter %s\n", D[i].jumlahKonsultasi, D[i].namaLengkap)
			fmt.Println()
			(*E)[j-1].dok = D[i].namaLengkap
			j++
		}
	}
}

func sortDokter(D *dokter, n int) {
	//Mengurutkan dokter berdasarkan dokter yang melayani konsultasi terbanyak.
	//Array terurut secara descending.
	//Pengurutan menggunakan algoritma insertion sort.

	for i := 1; i < n; i++ {
		temp := D[i]
		j := i - 1
		for j >= 0 && D[j].jumlahKonsultasi < temp.jumlahKonsultasi {
			D[j+1] = D[j]
			j--
		}
		D[j+i] = temp
	}
}
