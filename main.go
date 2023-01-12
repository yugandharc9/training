package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("learningcode.in", "www.learningcode.in", "167.71.230.152"), //Your domain here
		Cache:      autocert.DirCache("/etc/letsencrypt/live/learningcode.in"),                         //Folder for storing certificates

	}
	s := Server{}
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/contactform", s.ContactForm)
	http.HandleFunc("/bookform", s.BookForm)
	server := &http.Server{
		Addr: ":https",
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
			NextProtos: []string{
				"h2", "http/1.1", // enable HTTP/2
				acme.ALPNProto, // enable tls-alpn ACME challenges
			},
		},
	}

        go http.ListenAndServe(":http", http.HandlerFunc(redirectToHttps))

	log.Fatal(server.ListenAndServeTLS("", "")) //Key and cert are coming from Let's Encrypt
	// http.ListenAndServe(":80", nil)
}

func redirectToHttps(w http.ResponseWriter, r *http.Request) {
	    // Redirect the incoming HTTP request. Note that "127.0.0.1:443" will only work if you are accessing the server from your local machine.
	        http.Redirect(w, r, "https://learningcode.in"+r.RequestURI, http.StatusMovedPermanently)
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
