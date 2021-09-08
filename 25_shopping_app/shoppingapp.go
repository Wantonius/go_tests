package main

import (
    "net/http"
	"fmt"
	"encoding/json"
	"strconv"
)

type Item struct {
	Id		string	`json:"id"`
	Type 	string 	`json:"type"`
	Count	string	`json:"count"`
	Price	string	`json:"price"`	
}



func main() {

	shoppingItems := make([]Item,0)
	id := 100
    
	fs := http.FileServer(http.Dir("static/"))
    http.Handle("/", fs)
	
	http.HandleFunc("/api/shopping", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
			case http.MethodGet:
				json.NewEncoder(w).Encode(shoppingItems)
			case http.MethodPost:
				var item Item
				json.NewDecoder(r.Body).Decode(&item)
				item.Id = strconv.FormatInt(int64(id),10)
				id++
				shoppingItems = append(shoppingItems,item)
				w.Write([]byte("{message:'success'}"))
			default:
				w.Write([]byte("{message:'unknown command'}"))
		}
	})
	http.HandleFunc("/api/shopping/", func(w http.ResponseWriter, r *http.Request) {
		
		temp_string := r.URL.String()
		fmt.Println("temp_string",temp_string)
		temp_id := temp_string[len(temp_string)-3:]
		switch r.Method {
			case http.MethodDelete:
				for i,item := range shoppingItems {
					if item.Id == temp_id {
						shoppingItems = append(shoppingItems[:i],shoppingItems[i+1:]...)
					}
				}
				w.Write([]byte("{message:'success'}"))
			case http.MethodPut:
				var t_item Item
				json.NewDecoder(r.Body).Decode(&t_item)
				for i,item := range shoppingItems {
					if item.Id == temp_id {
						shoppingItems[i] = t_item
					}
				}
				w.Write([]byte("{message:'success'}"))
		}
		
	})		

	fmt.Println("Server ready in port 3000")
    http.ListenAndServe(":3000", nil)
}