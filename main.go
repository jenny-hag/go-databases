package main

  import (
     "time"
      "database/sql"
      "fmt"
      _ "github.com/mattn/go-sqlite3"
      _ "github.com/go-sql-driver/mysql"
      _ "github.com/lib/pq"
  )

  func main() {
    //insertToSQLite()
    //querySQLLite()
    //insertToMySQL()
    //queryMySQL()
    insertToPostgreSQL()
    queryPostgreSQL()
  }

  func insertToSQLite() {
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
    db.Close()
  }

  func querySQLLite() {
    db, err := sql.Open("sqlite3", "./travel-map.db")
    checkErr(err)

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

  func insertToMySQL() {
    db, err := sql.Open("mysql", "bucket_list_user:Rep1ace_with_real_password!@tcp(34.145.40.45:3306)/jenny-db")
  	if err != nil {
  		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
  	}
  	defer db.Close()

  	// Prepare statement for inserting data
  	stmtIns, err := db.Prepare("INSERT INTO bucket_list(place, country, photo, visited, created_at) values(?,?,?,?,?)") // ? = placeholder
  	if err != nil {
  		panic(err.Error()) // proper error handling instead of panic in your app
  	}
    res, err := stmtIns.Exec("chichen itza", "Mexico", "https://upload.wikimedia.org/wikipedia/commons/thumb/5/51/Chichen_Itza_3.jpg/440px-Chichen_Itza_3.jpg", false, "2022-11-04")
    checkErr(err)

    lastId, err := res.LastInsertId()
    checkErr(err)

    rowCnt, err := res.RowsAffected()
    checkErr(err)

    fmt.Println("ID = %d, affected = %d\n", lastId, rowCnt)

  	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates
  }

  func queryMySQL() {

  }

  func insertToPostgreSQL() {



    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "bucket_list_user", "Rep1ace_with_real_password!", "jenny-db")

    db, err := sql.Open("postgres", psqlconn)
    checkErr(err)

    defer db.Close()

    // dynamic
    stmtIns := `insert into "bucket_list"("place", "country", "photo", "visited", "created_at") values($1, $2, $3, $4, $5)`
    _, err = db.Exec(stmtIns, "chichen itza", "Mexico", "https://upload.wikimedia.org/wikipedia/commons/thumb/5/51/Chichen_Itza_3.jpg/440px-Chichen_Itza_3.jpg", false, "2022-11-04")
    checkErr(err)
  }

  func queryPostgreSQL() {

  }

  func checkErr(err error) {
      if err != nil {
          panic(err)
      }
  }
