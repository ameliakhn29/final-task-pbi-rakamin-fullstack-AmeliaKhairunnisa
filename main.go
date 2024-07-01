package finaltask

import (
	"fmt"
	"net/http"

	authcontroller "github.com/Asus/final-task/controllers"
)

func main() {
	http.HandleFunc("/", authcontroller.Index)
	http.HandleFunc("/register", authcontroller.Register)

	fmt.Println("Server jalan di: http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
