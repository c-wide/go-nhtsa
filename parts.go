package nhtsa

type PartsRequest struct {
	Type         string
	FromDate     string
	ToDate       string
	Manufacturer string
	Page         string
}

func GetParts() {
}
