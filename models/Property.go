package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Property struct {
	gorm.Model
	UnitType          string         `json:"unitType"`
	PropertyType      string         `json:"propertyType"`
	Street            string         `json:"street"`
	City              string         `json:"city"`
	State             string         `json:"state"`
	Zip               int            `json:"zip"`
	Lat               float32        `json:"lat"`
	Lng               float32        `json:"lng"`
	BedroomLow        int            `json:"bedroomLow"`
	BedroomHigh       int            `json:"bedroomHigh"`
	BathroomLow       float32        `json:"bathroomLow"`
	BathroomHigh      float32        `json:"bathroomHigh"`
	RentLow           float32        `json:"rentLow"`
	RentHigh          float32        `json:"rentHigh"`
	UserID            uint           `json:"userID"`
	Name              string         `json:"name"`
	Amenities         datatypes.JSON `json:"amenities"`
	IncludedUtilities datatypes.JSON `json:"includedUtilities"`
	Images            datatypes.JSON `json:"images"`
	Description       string         `json:"description"`
	Email             string         `json:"email"`
	FirstName         string         `json:"firstName"`
	LastName          string         `json:"lastName"`
	LaundryType       string         `json:"laundryType"`
	OnMarket          *bool          `json:"onMarket"`
	ParkingFee        float32        `json:"parkingFee"`
	PetsAllowed       string         `json:"petsAllowed"`
	CountryCode       string         `json:"countryCode"`
	CallingCode       string         `json:"callingCode"`
	PhoneNumber       string         `json:"phoneNumber"`
	Website           string         `json:"website"`
	Stars             float32        `json:"stars"`
	Apartments        []Apartment    `json:"apartments"`
	Reviews           []Review       `json:"reviews"`
}
