package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type tbl_barang struct {
	kode_barang  int
	nama_barang  string
	harga_barang string
	jenis_barang string
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/belajardbse")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	// insert data
	_, err = db.Exec("insert into tbl_barang values (?,?,?,?)", 12345, "Jam Tangan", "500.000", "Aksesoris")
	_, err = db.Exec("insert into tbl_barang values (?,?,?,?)", 678910, "Pendeng", "200.000", "Aksesoris")
	_, err = db.Exec("insert into tbl_barang values (?,?,?,?)", 111213, "Celana Dalam", "100.000", "Aksesoris")
	_, err = db.Exec("insert into tbl_barang values (?,?,?,?)", 141516, "Jam Dinding", "700.000", "Aksesoris")
	_, err = db.Exec("insert into tbl_barang values (?,?,?,?)", 17181920, "Baju Bola", "400.000", "Aksesoris")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Data berhasil ditambah\n")

	//Update data
	// _, err = db.Exec("update tbl_barang set nama_barang= ? where kode_barang = ?", "Tas Laptop", 12345)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Printf("Data berhasil diubah\n")

	// Hapus Data
	// _, err = db.Exec("delete from tbl_barang where npm=?", 51420175)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Printf("Data berhasil dihapus\n")

	//select data
	rows, err := db.Query("select * from tbl_barang")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []tbl_barang

	for rows.Next() {
		var each = tbl_barang{}
		var err = rows.Scan(&each.kode_barang, &each.nama_barang, &each.harga_barang, &each.jenis_barang)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Print(each.kode_barang)
		fmt.Print(" | ")
		fmt.Print(each.nama_barang)
		fmt.Print(" | ")
		fmt.Print(each.harga_barang)
		fmt.Print(" | ")
		fmt.Print(each.jenis_barang, "\n")
	}
}

func main() {
	sqlQuery()
}
