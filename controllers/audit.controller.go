package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/devrodriguez/audit-api/db"
	"github.com/devrodriguez/audit-api/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAudit(gCtx *gin.Context) {
	var response models.Response
	var audits []*models.Audit

	mgClient := db.GetClient()
	findOptions := options.Find()

	auditRef := mgClient.Database("audit").Collection("audits")
	auditsCur, err := auditRef.Find(context.TODO(), bson.D{{}}, findOptions)

	defer auditsCur.Close(context.TODO())

	if err != nil {
		response.Message = "Error reading data"
		response.Error = err.Error()
		gCtx.JSON(http.StatusInternalServerError, response)
		return
	}

	for auditsCur.Next(context.TODO()) {
		var audit models.Audit

		if err := auditsCur.Decode(&audit); err != nil {
			log.Fatal(err)
		}

		audits = append(audits, &audit)
	}

	if err := auditsCur.Err(); err != nil {
		response.Message = "Error reading cursor data"
		response.Error = err.Error()
		gCtx.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Message = "Data readed success"
	response.Data = audits

	gCtx.JSON(http.StatusOK, response)
}

func CreateAudit(gCtx *gin.Context) {
	var audit models.Audit
	var response models.Response
	mgClient := db.GetClient()

	if err := gCtx.BindJSON(&audit); err != nil {
		response.Error = err.Error()
		response.Message = "Error unbind JSON data"
		gCtx.JSON(http.StatusInternalServerError, response)
		return
	}

	auditRef := mgClient.Database("audit").Collection("audits")
	insertRes, err := auditRef.InsertOne(context.TODO(), audit)

	if err != nil {
		response.Error = err.Error()
		response.Message = "Error setting data"
		gCtx.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Message = "Document created"
	response.Data = gin.H{"docID": insertRes.InsertedID}

	gCtx.JSON(http.StatusOK, response)
}
