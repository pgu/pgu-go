package api

import (
	"appengine/datastore"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type badRequest struct{ error }

type BasicAuth struct {
	Username string
	Password string
}

func requireAuth(w http.ResponseWriter, err error) {
	w.Header().Set("WWW-Authenticate", "Basic realm=\"Pgu-go-API\"")
	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)

	if err != nil {
		log.Println("Auth err: ", err)
	}
}

func handleBasicAuth(w http.ResponseWriter, r *http.Request) bool {

	auth := r.Header.Get("Authorization")
	if len(auth) == 0 {
		requireAuth(w, errors.New("No Authorization header"))
		return false
	}

	auth_parts := strings.SplitN(auth, " ", 2)
	if len(auth_parts) != 2 {
		requireAuth(w, errors.New("auth parts != 2"))
		return false
	}
	if auth_parts[0] != "Basic" {
		requireAuth(w, errors.New("auth != Basic"))
		return false
	}

	b, err := base64.StdEncoding.DecodeString(auth_parts[1])
	if err != nil {
		requireAuth(w, errors.New("decode crashed"))
		return false
	}

	credentials_parts := strings.Split(string(b), ":")
	if len(credentials_parts) != 2 {
		requireAuth(w, errors.New("credentials malformed"))
		return false
	}

	credentials := &BasicAuth{Username: credentials_parts[0], Password: credentials_parts[1]}
	log.Println("Credentials ", credentials)

	if credentials.Username != "mickey" || credentials.Password != "mouse" {
		requireAuth(w, errors.New("wrong credentials"))
		return false
	}

	return true
}

func wrapJsonHandler(f func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		/*
			referer := r.Referer()
			front_origin := ""

			if strings.Contains(referer, "localhost:9000") || strings.Contains(referer, "127.0.0.1:9000") {
				front_origin = strings.TrimSuffix(referer, "/")
			}

			w.Header().Set("Access-Control-Allow-Origin", front_origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Authorization")

			if r.Method == "OPTIONS" {
				fmt.Fprint(w, "")
				return
			}
		*/

		//
		// PRE OPS:

		// - basic authentication
		/*
			ok := handleBasicAuth(w, r)
			if !ok {
				return
			}
		*/

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")
		w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, DELETE")
		w.Header().Set("Access-Control-Expose-Headers", "Location")

		if r.Method == "OPTIONS" {
			//			fmt.Fprint(w, "")
			return
		}

		// - setup response
		w.Header().Set("Content-Type", "application/json")

		//
		// EXEC handler
		err := f(w, r)

		//
		// POST OPS
		// - handle an error
		if err == nil {
			return
		}

		if err == datastore.ErrNoSuchEntity {
			http.Error(w, "entity not found", http.StatusNotFound)
			return
		}

		switch err.(type) {
		case badRequest:
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			log.Println(err)
			http.Error(w, "oops", http.StatusInternalServerError)
		}
	}
}

func parseID(r *http.Request) (int64, error) {
	txt, ok := mux.Vars(r)["id"]
	if !ok {
		return 0, fmt.Errorf("task id not found")
	}
	return strconv.ParseInt(txt, 10, 0)
}
