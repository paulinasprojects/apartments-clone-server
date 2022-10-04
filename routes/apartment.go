package routes

import (
	"apartments-clone-server/models"
	"apartments-clone-server/storage"
	"apartments-clone-server/utils"
	"time"

	"github.com/kataras/iris/v12"
)

func GetApartmentsByPropertyID(ctx iris.Context) {
	params := ctx.Params()
	id := params.Get("id")

	//quering the database for apartments where the propertyID is equial to the passed in ID
	var apartments []models.Apartment
	apartmentsExists := storage.DB.Where("property_id = ?", id).Find(&apartments)

	//If we get an error during the quering we will send the user the error
	if apartmentsExists.Error != nil {
		utils.CreateError(
			iris.StatusInternalServerError,
			"Error", apartmentsExists.Error.Error(), ctx)
		return
	}

	//otherwise we will send the user the apartments
	ctx.JSON(apartments)
}

// Update a list of apartments
func UpdateApartments(ctx iris.Context) {
	//get ID from the params
	params := ctx.Params()
	id := params.Get("id")

	//quering the database for it property and its apartments
	property := GetPropertyAndAssosiationsByPropertyID(id, ctx)
	if property == nil {
		return
	}

	// make sure that the body of the request is valid
	var updatedApartments []UpdateUnitsInput
	err := ctx.ReadJSON(&updatedApartments)
	if err != nil {
		utils.HandleValidationErrors(err, ctx)
		return
	}

	//loop through the updated apartments and then if there is a new bedroomLow
	// or BedroomHigh or bathroomLow or BathroomHigh then will will assign that to the variables

	var newApartments []models.Apartment
	bedroomLow := property.BedroomLow
	bedroomHigh := property.BedroomHigh
	var bathroomLow float32 = property.BathroomLow
	var bathroomHigh float32 = property.BathroomHigh

	for _, apartment := range updatedApartments {
		if *apartment.Bedrooms > bedroomHigh {
			bedroomHigh = *apartment.Bedrooms
		}
		if *apartment.Bedrooms < bedroomLow {
			bedroomLow = *apartment.Bedrooms
		}
		if apartment.Bathrooms > bathroomHigh {
			bathroomHigh = apartment.Bathrooms
		}
		if apartment.Bathrooms < bathroomLow {
			bathroomLow = apartment.Bathrooms
		}

		//creating an object  for the current apartment
		currApartment := models.Apartment{
			Unit:        apartment.Unit,
			Bedrooms:    *apartment.Bedrooms,
			Bathrooms:   apartment.Bathrooms,
			SqFt:        apartment.SqFt,
			Active:      apartment.Active,
			AvailableOn: apartment.AvailableOn,
			PropertyID:  property.ID,
		}

		//if the current apartment has and ID then it means it already exists so then
		//we will just update that current apartment
		//otherwise we need to append to the new apartments slice

		if apartment.ID != nil {
			currApartment.ID = *apartment.ID
			storage.DB.Model(&currApartment).Updates(currApartment)
		} else {
			newApartments = append(newApartments, currApartment)
		}

	}

	//see if the newApartment length is greater  than 0 and if it is
	//we need to create new apartments
	if len(newApartments) > 0 {
		rowsUpdated := storage.DB.Create(&newApartments)

		if rowsUpdated.Error != nil {
			utils.CreateError(
				iris.StatusInternalServerError,
				"Error", rowsUpdated.Error.Error(), ctx)
			return
		}
	}
	ctx.StatusCode(iris.StatusNoContent)
}

//Request body for the UpdateApartments

type UpdateUnitsInput struct {
	ID          *uint     `json:"ID"`
	Unit        string    `json:"unit" validate:"max=512"`
	Bedrooms    *int      `json:"bedrooms" validate:"gte=0,max=6,required"` // make int a pointer so 0 is accepted
	Bathrooms   float32   `json:"bathrooms" validate:"min=0.5,max=6.5,required"`
	SqFt        int       `json:"sqFt" validate:"max=100000000000,required"`
	Active      *bool     `json:"active" validate:"required"`
	AvailableOn time.Time `json:"availableOn" validate:"required"`
}
