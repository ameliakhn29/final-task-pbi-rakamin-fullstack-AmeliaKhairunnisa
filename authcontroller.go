package controllers

import (
	"errors"
	_ "hash"
	"html/template"
	"net/http"

	"github.com/Asus/final-task/app"
	"github.com/Asus/final-task/database"
	"github.com/Asus/final-task/models"
	"golang.org/x/crypto/bcrypt"
)

type UserInput struct {
	Username string
	Password string
}

var UserModel = models.NewUserModel()

func Index(w http.ResponseWriter, r *http.Request) {

	session, _ := database.Store.Get(r, database.SESSION_ID)

	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {

		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			temp, _ := template.ParseFiles("views/index.html")
			temp.Execute(w, nil)
		}

	}

}

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		temp, _ := template.ParseFiles("views/login.html")
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		UserInput := &UserInput{
			Username: r.Form.Get("username"),
			Password: r.Form.Get("password"),
		}

		var user app.User
		UserModel.Where(&user, "username", UserInput.Username)

		var message error
		if user.Username == "" {
			message = errors.New("username salah")
		} else {
			// pengecekan password
			errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(UserInput.Password))
			if errPassword != nil {
				message = errors.New("password salah")
			}
		}

		if message != nil {

			data := map[string]interface{}{
				"error": message,
			}

			temp, _ := template.ParseFiles("views/login.html")
			temp.Execute(w, data)
		} else {
			// set session

			session, _ := database.Store.Get(r, database.SESSION_ID)

			session.Values["loggedIn"] = true
			session.Values["email"] = user.Email
			session.Values["username"] = user.Username

			session.Save(r, w)

			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}

}

func Register(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		temp, _ := template.ParseFiles("views/register.html")
		temp.Execute(w, nil)

	} else if r.Method == http.MethodPost {

		r.ParseForm()

		user := app.User{
			Username: r.Form.Get("username"),
			Email:    r.Form.Get("email"),
			Password: r.Form.Get("password"),
		}

		errorMessages := make(map[string]interface{})

		if user.Username == "" {
			errorMessages["Username"] = "Username harus diisi"
		}
		if user.Email == "" {
			errorMessages["Email"] = "Email harus diisi"
		}
		if user.Password == "" {
			errorMessages["Password"] = "Username harus diisi"
		}

		if len(errorMessages) > 0 {

			data := map[string]interface{}{
				"validation": errorMessages,
			}

			temp, _ := template.ParseFiles("views/register.html")
			temp.Execute(w, data)
		} else {

			hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			user.Password = string(hashPassword)

			_, err := UserModel.Create(user)

			var error string
			var message string

			if err != nil {
				message = "Registration Process Failed : " + error
			} else {
				message = "Registration is Successful, Please login"
			}

			data := map[string]interface{}{
				"pesan": message,
			}

			temp, _ := template.ParseFiles("views/register.html")
			temp.Execute(w, data)
		}

	}
}
