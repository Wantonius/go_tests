package main

import (
    "net/http"
	"fmt"
	"encoding/json"
	"strconv"
	"math/rand"
	"time"
)

type Item struct {
	Id		string	`json:"id"`
	Type 	string 	`json:"type"`
	Count	string	`json:"count"`
	Price	string	`json:"price"`	
}

type User struct {
	Username	string	`json:"username"`
	Password 	string	`json:"password"`
}

type Session struct {
	TTL			int64		`json:"ttl"`
	token		string		`json:"token"`
}

type MyToken struct {
	Token		string		`json:"token"`
}

type BackendMessage struct {
	Message 	string 		`json:"message"`
}

const time_to_live = 3600
var shoppingItems []Item
var registeredUsers []User
var loggedSessions []Session
var id int
type Middleware func(http.HandlerFunc) http.HandlerFunc
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func handleGetAndPost(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
			case http.MethodGet:
				json.NewEncoder(w).Encode(shoppingItems)
			case http.MethodPost:
				var item Item
				json.NewDecoder(r.Body).Decode(&item)
				item.Id = strconv.FormatInt(int64(id),10)
				id++
				shoppingItems = append(shoppingItems,item)
				message := BackendMessage{Message:"success"}
				json.NewEncoder(w).Encode(message)
			default:
				message := BackendMessage{Message:"unknown command"}
				json.NewEncoder(w).Encode(message)
		}
}
func handleDeleteAndPut(w http.ResponseWriter, r *http.Request) {
		
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
				message := BackendMessage{Message:"success"}
				json.NewEncoder(w).Encode(message)
		}
		
}



func createToken() string {
	rand.Seed(time.Now().UnixNano())
    b := make([]rune, 128)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func register(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
			case http.MethodPost:
				var user User
				json.NewDecoder(r.Body).Decode(&user)
				registeredUsers = append(registeredUsers,user)
				message := BackendMessage{Message:"success"}
				json.NewEncoder(w).Encode(message)
			default:
				message := BackendMessage{Message:"unknown command"}
				json.NewEncoder(w).Encode(message)
		}
}


func login(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
			case http.MethodPost:
				var user User
				json.NewDecoder(r.Body).Decode(&user)
				for _,u := range registeredUsers {
					if(u.Username == user.Username) {
						if(u.Password == u.Password) {
							now := time.Now().Unix() + time_to_live
							t := createToken()
							loggedSessions = append(loggedSessions,Session{TTL:now,token:t})
							data := MyToken{Token:t}
							json.NewEncoder(w).Encode(data)
							return
						} else {
							w.WriteHeader(http.StatusForbidden)
							message := BackendMessage{Message:"forbidden"}
							json.NewEncoder(w).Encode(message)
							return
						}
					}
				}
				w.WriteHeader(http.StatusForbidden)
				message := BackendMessage{Message:"forbidden"}
				json.NewEncoder(w).Encode(message)				
			default:
				message := BackendMessage{Message:"unknown command"}
				json.NewEncoder(w).Encode(message)
		}
}
func isUserLogged() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("token")
			if(token == "") {
				w.WriteHeader(http.StatusForbidden)
				message := BackendMessage{Message:"forbidden"}
				json.NewEncoder(w).Encode(message)
				return
			}
			fmt.Printf("token:%s\n",token)
			for i,session := range loggedSessions {
				fmt.Println(session)
				if(token == session.token) {
					now := time.Now().Unix()
					if(now > session.TTL) {
						loggedSessions = append(loggedSessions[:i],loggedSessions[i+1:]...)
						w.WriteHeader(http.StatusForbidden)
						message := BackendMessage{Message:"forbidden"}
						json.NewEncoder(w).Encode(message)
						return
					} else {
						session.TTL = now + time_to_live
						f(w, r)
						return 
					}
				}
			}
			w.WriteHeader(http.StatusForbidden)
			message := BackendMessage{Message:"forbidden"}
			json.NewEncoder(w).Encode(message)
		}
	}
}

func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
    for _, m := range middlewares {
        f = m(f)
    }
    return f
}

func main() {

	shoppingItems = make([]Item,0)
	registeredUsers = make([]User,0)
	loggedSessions = make([]Session,0)
	id = 100
    
	fs := http.FileServer(http.Dir("static/"))
    http.Handle("/", fs)
	
	http.HandleFunc("/api/shopping", Chain(handleGetAndPost, isUserLogged()))
	http.HandleFunc("/api/shopping/", Chain(handleDeleteAndPut, isUserLogged()))	
	http.HandleFunc("/register",register)
	http.HandleFunc("/login",login)

	fmt.Println("Server ready in port 3000")
    http.ListenAndServe(":3000", nil)
}