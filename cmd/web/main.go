package main

import (
	"flag"
	"log"
	"os"
	"net/http"
)

type application struct {
	errLog *log.Logger
	infoLog *log.Logger
}

func main() {

	addr := flag.String("addr", ":8080", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errLog: errLog,
		infoLog: infoLog,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Let's define a custom http server with its
	// address, logger and handler
	srv := &http.Server {
		Addr: *addr,
		ErrorLog: errLog,
		Handler: mux,
	}

	infoLog.Println("Starting server on ", *addr)
	err := srv.ListenAndServe()
	
	errLog.Fatal(err.Error())
}