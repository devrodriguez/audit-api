package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/devrodriguez/audit-api/db"
	"github.com/devrodriguez/audit-api/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetEnterprises(gCtx *gin.Context) {
	var enterprises []*models.Enterprise
	mgClient := *db.GetClient()
	findOptions := options.Find()

	// Get entitie reference
	enterpriseRef := mgClient.Database("audit").Collection("enterprises")
	enterpriseCur, err := enterpriseRef.Find(context.TODO(), bson.D{{}}, findOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Read data cursor
	for enterpriseCur.Next(context.TODO()) {
		var enterprise models.Enterprise

		err := enterpriseCur.Decode(&enterprise)
		if err != nil {
			log.Fatal(err)
		}

		enterprises = append(enterprises, &enterprise)
	}

	// Validate cursor status
	if err := enterpriseCur.Err(); err != nil {
		log.Fatal(err)
	}

	enterpriseCur.Close(context.TODO())

	gCtx.JSON(http.StatusOK, enterprises)
}

func AddEnterprise(gCtx *gin.Context) {
	var enterprise models.Enterprise
	var response models.Response
	mgClient := *db.GetClient()

	if err := gCtx.BindJSON(&enterprise); err != nil {
		response.Error = err.Error()
		response.Message = "Fail binding JSON data"
		gCtx.JSON(http.StatusInternalServerError, response)
		return
	}

	// Get reference
	enterpriseRef := mgClient.Database("audit").Collection("enterprises")

	insertResult, err := enterpriseRef.InsertOne(context.TODO(), enterprise)
	if err != nil {
		response.Error = err.Error()
		response.Message = "Fail insert data"
		gCtx.JSON(http.StatusInternalServerError, response)
		return
	}

	// Build response
	response.Message = "Document created"
	response.Data = gin.H{"docId": insertResult.InsertedID}

	gCtx.JSON(http.StatusOK, response)
}

func UpdateEnterprise(gCtx *gin.Context) {
	var response models.Response
	var enterprise models.Enterprise
	mgClient := *db.GetClient()

	// Get request data
	id := gCtx.Param("id")
	docId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": docId}

	if err := gCtx.BindJSON(&enterprise); err != nil {
		response.Error = err.Error()
		response.Message = "Fail binding JSON data"
		gCtx.JSON(http.StatusInternalServerError, response)
		return
	}

	// Get reference
	enterpriseRef := mgClient.Database("audit").Collection("enterprises")
	// Update data
	updateResult, err := enterpriseRef.UpdateOne(context.TODO(), filter, bson.M{"$set": enterprise})

	if err != nil {
		response.Error = err.Error()
		response.Message = "Fail updating data"
		gCtx.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Message = "Data updated"
	response.Data = gin.H{"updated": updateResult.ModifiedCount}

	gCtx.JSON(http.StatusOK, response)
}
