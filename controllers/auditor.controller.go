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

func GetAuditor(gCtx *gin.Context) {
	var auditors []*models.Auditor
	var response models.Response

	mgClient := *db.GetClient()
	findOptions := options.Find()

	//findOptions.SetLimit(2)

	auditorsRef := mgClient.Database("audit").Collection("auditors")
	auditorCur, err := auditorsRef.Find(context.TODO(), bson.D{{}}, findOptions)

	if err != nil {
		response.Message = "Error getting data"
		response.Error = err.Error()

		gCtx.JSON(http.StatusInternalServerError, response)
		return
	}

	for auditorCur.Next(context.TODO()) {
		var auditor models.Auditor

		err := auditorCur.Decode(&auditor)
		if err != nil {
			log.Fatal(err)
		}

		auditors = append(auditors, &auditor)
	}

	if err := auditorCur.Err(); err != nil {
		log.Fatal(err)
	}

	auditorCur.Close(context.TODO())

	gCtx.JSON(http.StatusOK, auditors)
}

func AddAuditor(gCtx *gin.Context) {
	var response models.Response
	var auditor models.Auditor
	mgClient := db.GetClient()

	if err := gCtx.BindJSON(&auditor); err != nil {
		response.Error = err.Error()
		response.Message = "Error binding JSON data"
		gCtx.JSON(http.StatusInternalServerError, response)
		return
	}

	auditorRef := mgClient.Database("audit").Collection("auditors")

	insertRes, err := auditorRef.InsertOne(context.TODO(), auditor)

	if err != nil {
		response.Error = err.Error()
		response.Message = "Error reading data"
		gCtx.JSON(http.StatusInternalServerError, response)
		return
	}

	// Build response
	response.Message = "Document created"
	response.Data = gin.H{"docId": insertRes.InsertedID}
	gCtx.JSON(http.StatusOK, response)
}

func UpdateAuditor(gCtx *gin.Context) {
	var response models.Response
	var auditor models.Auditor

	mgClient := db.GetClient()

	// Get request data
	id := gCtx.Param("id")
	docId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": docId}

	if err := gCtx.BindJSON(&auditor); err != nil {
		response.Message = "Error binding data"
		response.Error = err.Error()
		gCtx.JSON(http.StatusInternalServerError, response)
		return
	}

	auditorRef := mgClient.Database("audit").Collection("auditors")
	updateRes, err := auditorRef.UpdateOne(context.TODO(), filter, bson.M{"$set": auditor})

	if err != nil {
		response.Message = "Error updating data"
		response.Error = err.Error()
		gCtx.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Message = "Data updated"
	response.Data = gin.H{"updated": updateRes.ModifiedCount}

	gCtx.JSON(http.StatusOK, response)
}

func DeleteAuditor(gCtx *gin.Context) {
	gCtx.JSON(http.StatusOK, gin.H{"message": "Auditor deleted"})
}
