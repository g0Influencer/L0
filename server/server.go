package server

import (
	"L0/cache"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)


func Run(){
	//defer database.Close()

	r:=mux.NewRouter()

	r.HandleFunc("/orders", OrderHanlder)

	err:= http.ListenAndServe(":8000", r)
	if err != nil {
		fmt.Print(err)
	}

}



func OrderHanlder(w http.ResponseWriter, r *http.Request){


	id:=r.URL.Query().Get("id")
	order,ok:=cache.GetByID(id)
	if !ok{
		http.NotFound(w,r)
		return
	}

	ts, err := template.ParseFiles("./static/orders.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, order)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}