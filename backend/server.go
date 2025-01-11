package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func enableCors(w *http.ResponseWriter) {

	// TODO Only allow known origin
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}

func checkMethodCorrect(r *http.Request, method string) bool {

	if r.Method == method {
		return true
	} else {
		return false
	}

}

func process_request(r *http.Request) {}

func main() {
	http.HandleFunc("/api/submit", submitHandler)
	http.HandleFunc("/api/checkSubmissionFinished", checkSubmissionFinishedHandler)
	http.HandleFunc("/api/getLetterLatex", getLetterLatexHandler)
	http.HandleFunc("/api/getLetterPDF", getLetterPDFHandler)
	http.HandleFunc("/api/getKDVLatex", getKDVLatexHandler)
	http.HandleFunc("/api/getSubmission", getSubmissionHandler)
	http.HandleFunc("/api/getKDVPDF", getKDVPDFHandler)
	http.HandleFunc("/api/getUser", getUserHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var submission Submission
	jsonError := json.Unmarshal(body, &submission)

	if jsonError != nil {
		resp := submitResponse{Status: "failed", Message: "json error!"}
		json.NewEncoder(w).Encode(resp)
	}
	fmt.Printf("body : %s", string(body))
	fmt.Printf("Submission : %+v", submission)
	fmt.Printf("Error : %+v", jsonError)

	//fmt.Printf("Fisrtname: %s, Lastname: %s, id: %s", submission.user.firstName, submission.user.lastname, submission.id)

	resp := submitResponse{Status: "success", Message: ""}
	json.NewEncoder(w).Encode(resp)
}
func checkSubmissionFinishedHandler(w http.ResponseWriter, r *http.Request) {
	print("Accessed!")
	print(r.Body)
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "{\"Message\":\"Hello, World!\"}")
}
func getLetterLatexHandler(w http.ResponseWriter, r *http.Request) {
	print("Accessed!")
	print(r)
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "{\"Message\":\"Hello, World!\"}")
}
func getLetterPDFHandler(w http.ResponseWriter, r *http.Request) {
	print("Accessed!")
	print(r)
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "{\"Message\":\"Hello, World!\"}")
}
func getKDVLatexHandler(w http.ResponseWriter, r *http.Request) {
	print("Accessed!")
	print(r)
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "{\"Message\":\"Hello, World!\"}")
}
func getSubmissionHandler(w http.ResponseWriter, r *http.Request) {
	print("Accessed!")
	print(r)
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "{\"Message\":\"Hello, World!\"}")
}
func getKDVPDFHandler(w http.ResponseWriter, r *http.Request) {
	print("Accessed!")
	print(r)
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "{\"Message\":\"Hello, World!\"}")
}
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	print("Accessed!")
	print(r)
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "{\"Message\":\"Hello, World!\"}")
}
