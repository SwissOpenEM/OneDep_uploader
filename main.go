package main

import (
	"fmt"
	"net/http"
	"tasks/depositions/onedep"

	"github.com/gin-gonic/gin"
)

var allowedType = []string{"vo-map", "add-map", "co-cif", "co-pdb", "fsc-xml", "half-map", "img-emdb", "mask-map"}
var needMeta = []string{"vo-map", "add-map", "half-map", "mask-map"}

func create(c *gin.Context) {
	var userInput onedep.UserInput
	if err := c.BindJSON(&userInput); err != nil {
		return
	}
	client := &http.Client{}
	deposition, err := onedep.CreateDeposition(client, userInput)
	if err != nil {
		return
	}
	fmt.Println("created deposition", deposition.Id)
	// FIXME: invoke the converter to mmCIF to produce metadata file
	// FIXME: add calls to validation API
	for f := range userInput.Files {
		// for all files provideed, upload them to onedep an dadd to struct
		fileDeposited, err := onedep.AddFileToDeposition(client, deposition, userInput.Files[f])
		if err != nil {
			return
		}
		deposition.Files = append(deposition.Files, fileDeposited)
		for i := range needMeta {
			if userInput.Files[f].Type == needMeta[i] {
				onedep.AddMetadataToFile(client, fileDeposited)
			}
		}
		// FIXME:  need to add checks if all necessary files are preaent based on deposition type
	}

	_, err = onedep.ProcesseDeposition(client, deposition)
	if err != nil {
		return
	}
}

func main() {
	router := gin.Default()
	router.POST("/onedep", create)
	router.Run("localhost:8080")

}
