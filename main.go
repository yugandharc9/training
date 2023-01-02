package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	s := Server{}
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/contactform", s.ContactForm)
	http.HandleFunc("/bookform", s.BookForm)
	http.ListenAndServe(":8000", nil)
}

type Server struct {
}

func writeToFile(filename, txt string) {
	f, err := os.OpenFile(filename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(txt + "\n"); err != nil {
		log.Println(err)
	}
}

func (s *Server) ContactForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Method) // prints the form data
	fmt.Println(r.Form)   // prints the form data
	if r.Method == http.MethodPost {
		// handle form submission
		j, err := json.Marshal(r.Form)
		fmt.Println(string(j), err)
		writeToFile("contact.txt", string(j))
		// display the form
		w.Write([]byte(`
	      <script>
		      alert("Thanks for your interest we will reach out to you soon");
		      window.location.href = "index.html";
	      </script>
	    `))
	}
}

func (s *Server) BookForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Method) // prints the form data
	fmt.Println(r.Form)   // prints the form data
	if r.Method == http.MethodPost {
		// handle form submission
		j, err := json.Marshal(r.Form)
		fmt.Println(string(j), err)
		writeToFile("book.txt", string(j))
		w.Write([]byte(`
	      <script>
		      alert("Thanks for your interest we will reach out to you soon");
		      window.location.href = "index.html";
	      </script>
	    `))
		// display the form
	}
}
