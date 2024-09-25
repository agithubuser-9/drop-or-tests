package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

// a struct
var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {

	// Capture connection properties.
	cfg := mysql.Config{
		// gets the value of an environment variable by its name (key).
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),

		Net:  "tcp",            // default
		Addr: "127.0.0.1:3306", // default for tcp

		DBName: "recordings",
	}

	// Get a database handle.
	var err error
	// DSN, or Data Source Name, is a string that provides the connection's parameters.
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	// test the connection by sending a simple request to the db
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	// "%v" is the respective default format.
	// there also are "%+v" for structs' field names and "%#v".
	fmt.Printf("Albums found: %v\n", albums)

	albumsss, err := albumsByArtist("LA DI")
	if err != nil {
		log.Fatal(err)
	}
	// "%v" is the respective default format.
	// there also are "%+v" for structs' field names and "%#v".
	fmt.Printf("Albums found: %v\n", albumsss)

	// Hard-code ID 2 here to test the query.
	alb, err := albumByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", alb)

	albID, err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)

}

// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string) ([]Album, error) {
	// An albums slice to hold data from returned rows.
	var albums []Album

	// '?' is a convention in SQL queries with database drivers that support parameterized queries such for mysql, but not for postgres (index-based).
	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	// reduces the risk of forgetting to call close() which leads to  resource leaks such as memory leaks (no one else can use the block of memory until the program terminates).
	// cleanup function is called on all return paths.
	// consistent cleanup regardless of the function exit is a return, panic, error, etc.
	// multiple defer calls execute in last-in, first-out order as a stack.
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	// while rows.Next() returns true...
	for rows.Next() {
		var alb Album
		// the current row prepared by Next.
		// checks for an error from scanning column values into the struct fields.
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			// format 'verbs': "%q", "%v", "%d", "%f", "%T", "%p", "%t", etc
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}

		// fmt.Println("---The album row")
		// print()
		// fmt.Println(alb)
		// fmt.Println("---End of the album row.")
		// print()

		albums = append(albums, alb)
	}

	// checks for errors such as unexpected closure of the connection, incomplete query's results, etc.
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

// albumByID queries for the album with the specified ID.
func albumByID(id int64) (Album, error) {
	// An album to hold data from the returned row.
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		// automatically handles checking for row existence
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}

// addAlbum adds the specified album to the database,
// returning the album ID of the new entry
func addAlbum(alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}
