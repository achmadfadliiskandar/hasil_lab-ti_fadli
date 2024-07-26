package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type tbl_mhs struct {
	npm     int
	nama    string
	kelas   string
	jurusan string
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/mahasiswa")
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

	//insert data
	_, err = db.Exec("insert into tbl_mhs values (?,?,?,?)", 51420175, "Bara", "4IA06", "Informatika")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Data berhasil ditambah\n")

	//Update data
	// _, err = db.Exec("update tbl_mhs set nama= ? where npm = ?", "Satya", 51420175)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Printf("Data berhasil diubah\n")

	// Hapus Data
	// _, err = db.Exec("delete from tbl_mhs where npm=?", 51420175)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Printf("Data berhasil dihapus\n")

	//select data
	rows, err := db.Query("select * from tbl_mhs")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []tbl_mhs

	for rows.Next() {
		var each = tbl_mhs{}
		var err = rows.Scan(&each.npm, &each.nama, &each.kelas, &each.jurusan)

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
		fmt.Print(each.npm)
		fmt.Print(" | ")
		fmt.Print(each.nama)
		fmt.Print(" | ")
		fmt.Print(each.kelas)
		fmt.Print(" | ")
		fmt.Print(each.jurusan, "\n")
	}
}

func main() {
	sqlQuery()
}
