package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EgorSkurihin/url-shortener/store"
	"github.com/gorilla/mux"
)

func (srv *Server) index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func (srv *Server) showURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sURL := vars["URL"]
	urlObj := &store.URLModel{
		ShortURL: sURL,
	}
	err := srv.store.GetLongURL(urlObj)
	if err != nil || urlObj.LongURL == "" {
		fmt.Fprintf(w, "Could not find saved long url that corresponds to the short url %s \n", sURL)
	}
	fmt.Println(urlObj)
	http.Redirect(w, r, urlObj.LongURL, http.StatusFound)

}

func (srv *Server) createURL(w http.ResponseWriter, r *http.Request) {
	urlObj := &store.URLModel{}
	if err := json.NewDecoder(r.Body).Decode(urlObj); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Could not decode json object in req body\n"))
	}
	r.Body.Close()
	if err := srv.store.CreateRoute(urlObj); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("This short url already taken\n"))
	}
	fmt.Fprint(w, urlObj.ShortURL)
}
