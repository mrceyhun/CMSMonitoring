package pipeline

import (
	"go/intelligence/models"
	"go/intelligence/utils"
	"log"
)

// Module     : intelligence
// Author     : Rahul Indra <indrarahul2013 AT gmail dot com>
// Created    : Wed, 1 July 2020 11:04:01 GMT
// Description: CMS MONIT infrastructure Intelligence Module

//MlBox - Machine Learning predicted Data
func MlBox(data <-chan models.AmJSON) <-chan models.AmJSON {

	if utils.ConfigJSON.Server.Verbose > 0 {
		log.Println("MlBox step", len(data), "records to prcoess")
	}
	/*
		IMPLEMENT THE LOGIC OF ML PREDICTIONS
		so far we are returning back the data without any predictions.
	*/

	predictedData := make(chan models.AmJSON)
	go func() {
		defer close(predictedData)
		for d := range data {
			predictedData <- d
		}
	}()
	return predictedData
}
