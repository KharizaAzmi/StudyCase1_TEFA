package routes

import (
	"GO/database"
	"GO/models"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func FindBootcamp(DBname string, CollectionName string, JudulBootcamp string) bool{
	ctx, client := database.HandleDBConnection()

	collection := client.Database(DBname).Collection(CollectionName)

	if JudulBootcamp != "" {
		_, errGet := collection.Find(ctx, bson.M{		
			"JudulBootcamp": JudulBootcamp,
		})
	
		if errGet != nil {
			return false
		}
	} 
	defer client.Disconnect(ctx)
	return true
}

func GetBootcamp(response http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
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
	
	getErr := FindBootcamp("admin", "user", bootcamp.JudulBootcamp)

	if getErr {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(&result)
	} else {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("{\"message\": \"Duplicate Data\"}"))
		return
	}
}