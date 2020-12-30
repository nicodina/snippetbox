package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/nicodina/snippetbox/pkg/models/mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golangcollege/sessions"
)

type application struct {
	errLog   *log.Logger
	infoLog  *log.Logger
	session *sessions.Session
	snippets *mysql.SnippetService
	temaplateCache map[string]*template.Template
}

func main() {

	addr := flag.String("addr", ":8080", "HTTP network address")
	dsn := flag.String("dsn", "web:mysql-password@/snippetbox?parseTime=true", "DSN string to connect to the database")
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Connection to the database
	db, err := openDB(*dsn)
	if err != nil {
		errLog.Fatal(err.Error())
	}
	defer db.Close()

	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(10)

	// Load templates into the cache
	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errLog.Fatal(err)
	}

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour
	session.Secure = true

	app := &application{
		errLog:   errLog,
		infoLog:  infoLog,
		session: session,
		snippets: &mysql.SnippetService{DB: db},
		temaplateCache: templateCache,
	}

	// Let's define a custom http server with its
	// address, logger and handler
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errLog,
		Handler:  app.routes(),
	}

	infoLog.Println("Starting server on ", *addr)
	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")

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
