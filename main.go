package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type City struct {
	Name        string `json:"name"`
	ImageURL    string `json:"imageUrl"`
	Description string `json:"description"`
}

type Country struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Cities []City `json:"cities"`
}

type Temple struct {
	Description string `json:"description"`
	ID          int    `json:"id"`
	ImageURL    string `json:"imageUrl"`
	Name        string `json:"name"`
}

type Beach struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ImageURL    string `json:"imageUrl"`
	Description string `json:"description"`
}

type Data struct {
	Countries []Country `json:"countries"`
	Temples   []Temple  `json:"temples"`
	Beaches   []Beach   `json:"beaches"`
}

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins, or specify your frontend's URL
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	data := Data{
		Countries: []Country{
			{
				ID:   1,
				Name: "Australia",
				Cities: []City{
					{Name: "Sydney, Australia", ImageURL: "sydneyaus.jpeg", Description: "A vibrant city known for its iconic landmarks like the Sydney Opera House and Sydney Harbour Bridge."},
					{Name: "Melbourne, Australia", ImageURL: "melbourneaus.jpeg", Description: "A cultural hub famous for its art, food, and diverse neighborhoods."},
				},
			},
			{
				ID:   2,
				Name: "Japan",
				Cities: []City{
					{Name: "Tokyo, Japan", ImageURL: "tokyojapan.jpeg", Description: "A bustling metropolis blending tradition and modernity, famous for its cherry blossoms and rich culture."},
					{Name: "Kyoto, Japan", ImageURL: "kyotojapan.jpeg", Description: "Known for its historic temples, gardens, and traditional tea houses."},
				},
			},
			{
				ID:   3,
				Name: "Brazil",
				Cities: []City{
					{Name: "Rio de Janeiro, Brazil", ImageURL: "riodejaneiro.jpeg", Description: "A lively city known for its stunning beaches, vibrant carnival celebrations, and iconic landmarks."},
					{Name: "São Paulo, Brazil", ImageURL: "saopaulo.jpeg", Description: "The financial hub with diverse culture, arts, and a vibrant nightlife."},
				},
			},
		},
		Temples: []Temple{
			{ID: 1, Name: "Angkor Wat, Cambodia", ImageURL: "angkorwat.jpeg", Description: "A UNESCO World Heritage site and the largest religious monument in the world."},
			{ID: 2, Name: "Taj Mahal, India", ImageURL: "tajmahal.jpeg", Description: "An iconic symbol of love and a masterpiece of Mughal architecture."},
		},
		Beaches: []Beach{
			{ID: 1, Name: "Bora Bora, French Polynesia", ImageURL: "borabora.jpeg", Description: "An island known for its stunning turquoise waters and luxurious overwater bungalows."},
			{ID: 2, Name: "Copacabana Beach, Brazil", ImageURL: "copabeach.jpeg", Description: "A famous beach in Rio de Janeiro, Brazil, with a vibrant atmosphere and scenic views."},
		},
	}
	//Set the response headers
	w.Header().Set("Content-Type", "application/json")
	//Encode and return json data
	json.NewEncoder(w).Encode(data)

}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/api/info", getInfo)
	fmt.Println("Server running at" + port + "...")
	http.ListenAndServe(":"+port, nil)
}
