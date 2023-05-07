package routes

import (
	"GO/database"
	"GO/models"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func HandleBootcampInsert(DBname string, CollectionName string, JudulBootcamp string, Deskripsi string, Harga float64, Durasi string, Jadwal string) bool {
	ctx, client := database.HandleDBConnection()

	collection := client.Database(DBname).Collection(CollectionName)

	_, errInsert := collection.InsertOne(ctx, bson.M{		
		"JudulBootcamp": JudulBootcamp,
		"Deskripsi": Deskripsi,
		"Harga": Harga,
		"Durasi": Durasi,  
		"Jadwal": Jadwal, 
	})

	if errInsert != nil {
		return false
	}
	defer client.Disconnect(ctx)
	return true
}

func HandleAddBootcamp(response http.ResponseWriter, request *http.Request){
	if request.Method != "POST" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("{\"message\": \"Method not allowed\"}"))
		return
	}

	var bootcamp models.BootcampModel
	var result models.ResponseModel

	err := json.NewDecoder(request.Body).Decode(&bootcamp)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	
	insertErr := HandleBootcampInsert("admin", "user", bootcamp.JudulBootcamp, bootcamp.Deskripsi, bootcamp.Harga, bootcamp.Jadwal, bootcamp.Jadwal)

	if insertErr {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(&result)
	} else {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("{\"message\": \"Duplicate Data\"}"))
		return
	}
}