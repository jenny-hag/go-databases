package main

  import (
     "time"
      "database/sql"
      "fmt"
      _ "github.com/mattn/go-sqlite3"
  )

  func main() {
      db, err := sql.Open("sqlite3", "./travel-map.db")
      checkErr(err)

      // insert
      stmt, err := db.Prepare("INSERT INTO bucket_list(place, country, photo, visited, created_at) values(?,?,?,?,?)")
      checkErr(err)

      res, err := stmt.Exec("chichen itza", "Mexico", "https://upload.wikimedia.org/wikipedia/commons/thumb/5/51/Chichen_Itza_3.jpg/440px-Chichen_Itza_3.jpg", false, "2022-11-04");
      checkErr(err)

      id, err := res.LastInsertId()
      checkErr(err)

      fmt.Println(id)

      // query
      rows, err := db.Query("SELECT * FROM bucket_list")
      checkErr(err)
      var uid int
      var place string
      var country string
      var photo string
      var visited bool
      var created_at time.Time

      for rows.Next() {
          err = rows.Scan(&uid, &place, &country, &photo, &visited, &created_at)
          checkErr(err)
          fmt.Println(uid)
          fmt.Println(place)
          fmt.Println(country)
          fmt.Println(photo)
          fmt.Println(visited)
          fmt.Println(created_at)
      }

      rows.Close()
      db.Close()
  }

  func checkErr(err error) {
      if err != nil {
          panic(err)
      }
  }
