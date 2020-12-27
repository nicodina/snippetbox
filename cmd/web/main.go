package main

import (
	"database/sql"
	"flag"
	"log"
	"os"
	"net/http"

	"github.com/nicodina/snippetbox/pkg/models/mysql"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errLog *log.Logger
	infoLog *log.Logger
	snippets *mysql.SnippetService
}

func main() {

	addr := flag.String("addr", ":8080", "HTTP network address")
	dsn := flag.String("dsn", "web:mysql-password@/snippetbox?parseTime=true", "DSN string to connect to the database")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Connection to the database
	db, err := openDB(*dsn)
	if err != nil {
		errLog.Fatal(err.Error())
	}
	defer db.Close()

	app := &application{
		errLog: errLog,
		infoLog: infoLog,
		snippets: &mysql.SnippetService{DB: db},
	}

	// Let's define a custom http server with its
	// address, logger and handler
	srv := &http.Server {
		Addr: *addr,
		ErrorLog: errLog,
		Handler: app.routes(),
	}

	infoLog.Println("Starting server on ", *addr)
	err = srv.ListenAndServe()
	
	errLog.Fatal(err.Error())
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
		return db, nil
}