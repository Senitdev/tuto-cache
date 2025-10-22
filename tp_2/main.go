package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

func main() {
	//cache avec TTL de 10s , nettoyage toutes les 30s
	c := cache.New(1*time.Minute, 2*time.Minute)
	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/users/"):]
		//On vérifie dans la cache
		if val, found := c.Get(id); found {
			fmt.Printf("cache it :\n", id)
			json.NewEncoder(w).Encode(val)
			return
		}
		//Sinon on cherche dans le BD
		fmt.Println("Cache Miss:\n", id)
		user, _ := FetchUserBD(id)
		//on stock dans la cache
		c.Set(id, user, cache.DefaultExpiration)
		json.NewEncoder(w).Encode(user)
	})
	fmt.Printf("Serveur démarrer sur localhost")
	http.ListenAndServe(":8081", nil)
}
func FetchUserBD(id string) (map[string]string, error) {
	time.Sleep(10 * time.Second)
	return map[string]string{
		"id":   id,
		"name": fmt.Sprintf("User %s", id),
	}, nil
}
