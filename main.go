package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUser     = "admin"
	dbPassword = "Password1234"
	dbHost     = "bow-hotels-rds.c1a0iooss11x.us-east-1.rds.amazonaws.com"
	dbPort     = "3306"
	dbName     = "bowhotels"
)

var db *sql.DB

func main() {
	// Initialize database connection
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Test the database connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Unable to reach database: %v", err)
	}
	fmt.Println("Connected to the database successfully!")

	// Serve static form
	// http.HandleFunc("/", formHandler)
	// Handle form submission
	http.HandleFunc("/submit", submitHandler)

	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	// Start server on port 80
	log.Println("Starting server on port 80...")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}

// Struct to store form data
type InquiryData struct {
	Sender  string
	Email   string
	Message string
}

// Handler to serve the form
func formHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form template
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Error loading form", http.StatusInternalServerError)
		return
	}

	// Serve the form
	tmpl.Execute(w, nil)
}

// Handler to process form submission
func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}

	print("Someone called this endpoint")

	// Capture submitted data
	data := InquiryData{
		Sender:  r.FormValue("name"),
		Email:   r.FormValue("email"),
		Message: r.FormValue("message"),
	}

	// Insert data into the database
	query := "INSERT INTO inquiry (sender, email, message) VALUES (?, ?, ?)"
	_, err = db.Exec(query, data.Sender, data.Email, data.Message)
	if err != nil {
		log.Printf("Error inserting data: %v", err)
		http.Error(w, "Failed to save data", http.StatusInternalServerError)
		return
	}

	// Display the registration details as response (for testing purposes)
	fmt.Fprintf(w, "Inquiry sent!\n\nDetails:\nFirst Name: %s\nEmail: %s\nMessage: %s", data.Sender, data.Email, data.Message)
}
