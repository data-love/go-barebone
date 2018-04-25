package main //HealthCheckHandler returns OK if reachable
import (
	"io"
	"log"
	"net/http"
)

//HealthCheckHandler returns OK if reachable
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func main() {
	//port := os.Getenv("PORT")
	port := "3000"
	http.HandleFunc("/health-check", HealthCheckHandler)
	log.Print("Go Server runs on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
