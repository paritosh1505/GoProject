package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)
type Movie struct {
	ID        string    `json:"id"`
	Isbn      string    `json:"ISBN"`
	MovieName string    `json:"MovieName"`
	Director  *Director `json:"director"`
}
var movieentry []Movie


type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"Lastname"`
}

func getMovies(writer http.ResponseWriter,request *http.Request){
	writer.Header().Set("Content-Type","application/json")
	json.NewEncoder(writer).Encode(movieentry)
}
func getMovie(writer http.ResponseWriter,request *http.Request){
	writer.Header().Set("Content-type","application/json")
	vars:=mux.Vars(request)
	id := vars["id"]
	 var filtermovie *Movie
	 for _,entry:= range movieentry{
		if entry.ID==id{
			filtermovie=&entry
			break
		}
		
	 }

	 if filtermovie!=nil {
		json.NewEncoder(writer).Encode(filtermovie)
		return
	 }else{
		
			http.NotFound(writer,request)
		
	 }
	
	
}
func createMovie(writer http.ResponseWriter, request *http.Request){
getMovieId := movieentry[len(movieentry)-1].ID
getMovi,_ := strconv.Atoi(getMovieId);
getMovi++
movieentry = append(movieentry, Movie{ID:strconv.Itoa(getMovi),Isbn:"newISBN",MovieName:"New Created movie",Director:&Director{FirstName:"Paritosh",LastName:"Pantola"}})
writer.Header().Set("Content-type","application/json")


json.NewEncoder(writer).Encode(movieentry)
writer.WriteHeader(http.StatusAccepted)


}
func UpdateMovie(writer http.ResponseWriter,request *http.Request){
 queryparam := mux.Vars(request)
 id_update := queryparam["id"]
 fmt.Println("****************",queryparam)

 for index,entry:= range movieentry{
	if(entry.ID==id_update){
		movieentry[index].MovieName="New Update Movie"
		movieentry[index].Isbn="Hello this is new ISBN"
	}
 }
writer.Header().Set("Content-Type","aplication/json")
 json.NewEncoder(writer).Encode(movieentry)

}
func DeleteEntry(writer http.ResponseWriter,request *http.Request){
queryparma:=mux.Vars(request);
id_val := queryparma["id"]
index_val:=-1;
for index,entry := range movieentry{
	if(entry.ID==id_val){
		index_val=index
		break
	}
}
if(index_val!=-1){
	movieentry=append(movieentry[:index_val])

}
writer.Header().Set("Content-Type","aplication/json")
 json.NewEncoder(writer).Encode(movieentry)
}
func main() {
	router :=mux.NewRouter();
	movieentry = append(movieentry, Movie{ID: "1", Isbn: "testIsn", MovieName: "sherk", Director: &Director{FirstName: "Sherlock", LastName: "Homes"}})
	movieentry = append(movieentry, Movie{ID: "2", Isbn: "movie2isbn", MovieName: "second move", Director: &Director{FirstName: "George", LastName: "Bush"}})
	router.HandleFunc("/movies",getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	router.HandleFunc("/movies",createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}",UpdateMovie).Methods("PUT")
	router.HandleFunc("/movies1/{id}",DeleteEntry).Methods("DELETE")

	 http.ListenAndServe(":3000",router)
	
}