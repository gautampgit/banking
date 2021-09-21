package dto

type CustomerResponse struct {
	Id          string `json:"id" xml:"id"`
	Name        string `json:"full_name" xml:"name"`
	City        string `json:"city" xml:"city"`
	ZipCode     string `json:"zip" xml:"zip"`
	DateOfBirth string `json:"dob" xml:"dob"`
	Status      string `json:"status" xml:"status"`
}
