package main

// Import our dependencies. We'll use the standard HTTP library as well as the gorilla router for this app
import (
	"fmt"
	"log"
	"net/http"

	"github.com/Parth-PDF/Visimage-App/dao"

	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

const (
	host     = "ec2-52-44-46-66.compute-1.amazonaws.com"
	port     = 5432
	user     = "daqyecxiqmzavm"
	password = "b62390bfa5f1605dd77816c233020ff4c0716ae3226de957626e466b3ac19442"
	dbname   = "d8cvoj226qv7eq"
)

func main() {
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)
	db := sqlx.MustConnect("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname))
	defer db.Close()

	imageDao := dao.NewImageDao(db)

	log.Println("Database successfully connected!")

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: VerificationKeyGetter,
		SigningMethod:       jwt.SigningMethodRS256,
	})

	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./views/")))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	r.Handle("/images", jwtMiddleware.Handler(http.HandlerFunc(ImagesHandler(imageDao)))).Methods("GET")
	r.Handle("/upload", jwtMiddleware.Handler(http.HandlerFunc(UploadHandler(imageDao)))).Methods("POST")
	r.Handle("/delete", jwtMiddleware.Handler(http.HandlerFunc(DeleteHandler(imageDao)))).Methods("DELETE")

	// For dev only - Set up CORS so React client can consume our API
	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", corsWrapper.Handler(r))
}

//sudo kill -9 `sudo lsof -t -i:8080`
