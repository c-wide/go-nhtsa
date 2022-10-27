package nhtsa

type VinRequest struct {
	Vin  string
	Year string
}

type PartsRequest struct {
	Type         string
	FromDate     string
	ToDate       string
	Manufacturer string
	Page         string
}

type ManufacturerType struct {
	Name string `json:"Name"`
}

type ManufacturerVehicleType struct {
	IsPrimary bool   `json:"IsPrimary"`
	Name      string `json:"Name"`
}

type BaseApiResponse[T any] struct {
	Count          int64       `json:"Count"`
	Message        string      `json:"Message"`
	SearchCriteria interface{} `json:"SearchCriteria"`
	Results        []T         `json:"Results"`
}

type DecodeVinResult struct {
	Value      string `json:"Value"`
	ValueID    string `json:"ValueId"`
	Variable   string `json:"Variable"`
	VariableID int64  `json:"VariableId"`
}

type DecodeWmiResult struct {
	CommonName            string      `json:"CommonName"`
	CreatedOn             string      `json:"CreatedOn"`
	DateAvailableToPublic string      `json:"DateAvailableToPublic"`
	Make                  string      `json:"Make"`
	ManufacturerName      string      `json:"ManufacturerName"`
	ParentCompanyName     string      `json:"ParentCompanyName"`
	URL                   string      `json:"URL"`
	UpdatedOn             interface{} `json:"UpdatedOn"`
	VehicleType           string      `json:"VehicleType"`
}

type DecodeFlatResult struct {
	ABS                                 string `json:"ABS"`
	ActiveSafetySysNote                 string `json:"ActiveSafetySysNote"`
	AdaptiveCruiseControl               string `json:"AdaptiveCruiseControl"`
	AdaptiveDrivingBeam                 string `json:"AdaptiveDrivingBeam"`
	AdaptiveHeadlights                  string `json:"AdaptiveHeadlights"`
	AdditionalErrorText                 string `json:"AdditionalErrorText"`
	AirBagLocCurtain                    string `json:"AirBagLocCurtain"`
	AirBagLocFront                      string `json:"AirBagLocFront"`
	AirBagLocKnee                       string `json:"AirBagLocKnee"`
	AirBagLocSeatCushion                string `json:"AirBagLocSeatCushion"`
	AirBagLocSide                       string `json:"AirBagLocSide"`
	AutoReverseSystem                   string `json:"AutoReverseSystem"`
	AutomaticPedestrianAlertingSound    string `json:"AutomaticPedestrianAlertingSound"`
	AxleConfiguration                   string `json:"AxleConfiguration"`
	Axles                               string `json:"Axles"`
	BasePrice                           string `json:"BasePrice"`
	BatteryA                            string `json:"BatteryA"`
	BatteryATo                          string `json:"BatteryA_to"`
	BatteryCells                        string `json:"BatteryCells"`
	BatteryInfo                         string `json:"BatteryInfo"`
	BatteryKWh                          string `json:"BatteryKWh"`
	BatteryKWhTo                        string `json:"BatteryKWh_to"`
	BatteryModules                      string `json:"BatteryModules"`
	BatteryPacks                        string `json:"BatteryPacks"`
	BatteryType                         string `json:"BatteryType"`
	BatteryV                            string `json:"BatteryV"`
	BatteryVTo                          string `json:"BatteryV_to"`
	BedLengthIN                         string `json:"BedLengthIN"`
	BedType                             string `json:"BedType"`
	BlindSpotIntervention               string `json:"BlindSpotIntervention"`
	BlindSpotMon                        string `json:"BlindSpotMon"`
	BodyCabType                         string `json:"BodyCabType"`
	BodyClass                           string `json:"BodyClass"`
	BrakeSystemDesc                     string `json:"BrakeSystemDesc"`
	BrakeSystemType                     string `json:"BrakeSystemType"`
	BusFloorConfigType                  string `json:"BusFloorConfigType"`
	BusLength                           string `json:"BusLength"`
	BusType                             string `json:"BusType"`
	CanAacn                             string `json:"CAN_AACN"`
	CIB                                 string `json:"CIB"`
	CashForClunkers                     string `json:"CashForClunkers"`
	ChargerLevel                        string `json:"ChargerLevel"`
	ChargerPowerKW                      string `json:"ChargerPowerKW"`
	CoolingType                         string `json:"CoolingType"`
	CurbWeightLB                        string `json:"CurbWeightLB"`
	CustomMotorcycleType                string `json:"CustomMotorcycleType"`
	DaytimeRunningLight                 string `json:"DaytimeRunningLight"`
	DestinationMarket                   string `json:"DestinationMarket"`
	DisplacementCC                      string `json:"DisplacementCC"`
	DisplacementCI                      string `json:"DisplacementCI"`
	DisplacementL                       string `json:"DisplacementL"`
	Doors                               string `json:"Doors"`
	DriveType                           string `json:"DriveType"`
	DriverAssist                        string `json:"DriverAssist"`
	DynamicBrakeSupport                 string `json:"DynamicBrakeSupport"`
	EDR                                 string `json:"EDR"`
	ESC                                 string `json:"ESC"`
	EVDriveUnit                         string `json:"EVDriveUnit"`
	ElectrificationLevel                string `json:"ElectrificationLevel"`
	EngineConfiguration                 string `json:"EngineConfiguration"`
	EngineCycles                        string `json:"EngineCycles"`
	EngineCylinders                     string `json:"EngineCylinders"`
	EngineHP                            string `json:"EngineHP"`
	EngineHPTo                          string `json:"EngineHP_to"`
	EngineKW                            string `json:"EngineKW"`
	EngineManufacturer                  string `json:"EngineManufacturer"`
	EngineModel                         string `json:"EngineModel"`
	EntertainmentSystem                 string `json:"EntertainmentSystem"`
	ErrorCode                           string `json:"ErrorCode"`
	ErrorText                           string `json:"ErrorText"`
	ForwardCollisionWarning             string `json:"ForwardCollisionWarning"`
	FuelInjectionType                   string `json:"FuelInjectionType"`
	FuelTypePrimary                     string `json:"FuelTypePrimary"`
	FuelTypeSecondary                   string `json:"FuelTypeSecondary"`
	GCWR                                string `json:"GCWR"`
	GCWRTo                              string `json:"GCWR_to"`
	GVWR                                string `json:"GVWR"`
	GVWRTo                              string `json:"GVWR_to"`
	KeylessIgnition                     string `json:"KeylessIgnition"`
	LaneCenteringAssistance             string `json:"LaneCenteringAssistance"`
	LaneDepartureWarning                string `json:"LaneDepartureWarning"`
	LaneKeepSystem                      string `json:"LaneKeepSystem"`
	LowerBeamHeadlampLightSource        string `json:"LowerBeamHeadlampLightSource"`
	Make                                string `json:"Make"`
	MakeID                              string `json:"MakeID"`
	Manufacturer                        string `json:"Manufacturer"`
	ManufacturerID                      string `json:"ManufacturerId"`
	Model                               string `json:"Model"`
	ModelID                             string `json:"ModelID"`
	ModelYear                           string `json:"ModelYear"`
	MotorcycleChassisType               string `json:"MotorcycleChassisType"`
	MotorcycleSuspensionType            string `json:"MotorcycleSuspensionType"`
	NCSABodyType                        string `json:"NCSABodyType"`
	NCSAMake                            string `json:"NCSAMake"`
	NCSAMapExcApprovedBy                string `json:"NCSAMapExcApprovedBy"`
	NCSAMapExcApprovedOn                string `json:"NCSAMapExcApprovedOn"`
	NCSAMappingException                string `json:"NCSAMappingException"`
	NCSAModel                           string `json:"NCSAModel"`
	NCSANote                            string `json:"NCSANote"`
	NonLandUse                          string `json:"NonLandUse"`
	Note                                string `json:"Note"`
	OtherBusInfo                        string `json:"OtherBusInfo"`
	OtherEngineInfo                     string `json:"OtherEngineInfo"`
	OtherMotorcycleInfo                 string `json:"OtherMotorcycleInfo"`
	OtherRestraintSystemInfo            string `json:"OtherRestraintSystemInfo"`
	OtherTrailerInfo                    string `json:"OtherTrailerInfo"`
	ParkAssist                          string `json:"ParkAssist"`
	PedestrianAutomaticEmergencyBraking string `json:"PedestrianAutomaticEmergencyBraking"`
	PlantCity                           string `json:"PlantCity"`
	PlantCompanyName                    string `json:"PlantCompanyName"`
	PlantCountry                        string `json:"PlantCountry"`
	PlantState                          string `json:"PlantState"`
	PossibleValues                      string `json:"PossibleValues"`
	Pretensioner                        string `json:"Pretensioner"`
	RearAutomaticEmergencyBraking       string `json:"RearAutomaticEmergencyBraking"`
	RearCrossTrafficAlert               string `json:"RearCrossTrafficAlert"`
	RearVisibilitySystem                string `json:"RearVisibilitySystem"`
	SAEAutomationLevel                  string `json:"SAEAutomationLevel"`
	SAEAutomationLevelTo                string `json:"SAEAutomationLevel_to"`
	SeatBeltsAll                        string `json:"SeatBeltsAll"`
	SeatRows                            string `json:"SeatRows"`
	Seats                               string `json:"Seats"`
	SemiautomaticHeadlampBeamSwitching  string `json:"SemiautomaticHeadlampBeamSwitching"`
	Series                              string `json:"Series"`
	Series2                             string `json:"Series2"`
	SteeringLocation                    string `json:"SteeringLocation"`
	SuggestedVIN                        string `json:"SuggestedVIN"`
	TPMS                                string `json:"TPMS"`
	TopSpeedMPH                         string `json:"TopSpeedMPH"`
	TrackWidth                          string `json:"TrackWidth"`
	TractionControl                     string `json:"TractionControl"`
	TrailerBodyType                     string `json:"TrailerBodyType"`
	TrailerLength                       string `json:"TrailerLength"`
	TrailerType                         string `json:"TrailerType"`
	TransmissionSpeeds                  string `json:"TransmissionSpeeds"`
	TransmissionStyle                   string `json:"TransmissionStyle"`
	Trim                                string `json:"Trim"`
	Trim2                               string `json:"Trim2"`
	Turbo                               string `json:"Turbo"`
	VIN                                 string `json:"VIN"`
	ValveTrainDesign                    string `json:"ValveTrainDesign"`
	VehicleDescriptor                   string `json:"VehicleDescriptor"`
	VehicleType                         string `json:"VehicleType"`
	WheelBaseLong                       string `json:"WheelBaseLong"`
	WheelBaseShort                      string `json:"WheelBaseShort"`
	WheelBaseType                       string `json:"WheelBaseType"`
	WheelSizeFront                      string `json:"WheelSizeFront"`
	WheelSizeRear                       string `json:"WheelSizeRear"`
	Wheels                              string `json:"Wheels"`
	Windows                             string `json:"Windows"`
}

type GetWmiResult struct {
	Country               interface{} `json:"Country"`
	CreatedOn             string      `json:"CreatedOn"`
	DateAvailableToPublic string      `json:"DateAvailableToPublic"`
	ID                    int64       `json:"Id"`
	Name                  string      `json:"Name"`
	UpdatedOn             interface{} `json:"UpdatedOn"`
	VehicleType           string      `json:"VehicleType"`
	WMI                   string      `json:"WMI"`
}

type GetAllMakesResult struct {
	MakeId   int64  `json:"Make_ID"`
	MakeName string `json:"Make_Name"`
}

type GetPartsResult struct {
	CoverLetterURL   string      `json:"CoverLetterURL"`
	LetterDate       string      `json:"LetterDate"`
	ManufacturerID   int64       `json:"ManufacturerId"`
	ManufacturerName string      `json:"ManufacturerName"`
	ModelYearFrom    interface{} `json:"ModelYearFrom"`
	ModelYearTo      interface{} `json:"ModelYearTo"`
	Name             string      `json:"Name"`
	Type             string      `json:"Type"`
	URL              string      `json:"URL"`
}

type GetAllManufacturersResult struct {
	Country      string                    `json:"Country"`
	CommonName   string                    `json:"Mfr_CommonName"`
	ID           int64                     `json:"Mfr_ID"`
	Name         string                    `json:"Mfr_Name"`
	VehicleTypes []ManufacturerVehicleType `json:"VehicleTypes"`
}

type GetManufacturerDetailsResult struct {
	Address                  string                    `json:"Address"`
	Address2                 string                    `json:"Address2"`
	City                     string                    `json:"City"`
	ContactEmail             string                    `json:"ContactEmail"`
	ContactFax               interface{}               `json:"ContactFax"`
	ContactPhone             string                    `json:"ContactPhone"`
	Country                  string                    `json:"Country"`
	DBAs                     string                    `json:"DBAs"`
	EquipmentItems           []interface{}             `json:"EquipmentItems"`
	LastUpdated              string                    `json:"LastUpdated"`
	ManufacturerTypes        []ManufacturerType        `json:"ManufacturerTypes"`
	MfrCommonName            string                    `json:"Mfr_CommonName"`
	MfrID                    int64                     `json:"Mfr_ID"`
	MfrName                  string                    `json:"Mfr_Name"`
	OtherManufacturerDetails string                    `json:"OtherManufacturerDetails"`
	PostalCode               string                    `json:"PostalCode"`
	PrimaryProduct           interface{}               `json:"PrimaryProduct"`
	PrincipalFirstName       string                    `json:"PrincipalFirstName"`
	PrincipalLastName        interface{}               `json:"PrincipalLastName"`
	PrincipalPosition        string                    `json:"PrincipalPosition"`
	StateProvince            string                    `json:"StateProvince"`
	SubmittedName            string                    `json:"SubmittedName"`
	SubmittedOn              string                    `json:"SubmittedOn"`
	SubmittedPosition        string                    `json:"SubmittedPosition"`
	VehicleTypes             []ManufacturerVehicleType `json:"VehicleTypes"`
}

type GetMakeManufacturerNameResult struct {
	MakeID   int64  `json:"Make_ID"`
	MakeName string `json:"Make_Name"`
	MfrName  string `json:"Mfr_Name"`
}

type GetMakeManufacturerNameYearResult struct {
	MakeID   int64  `json:"MakeId"`
	MakeName string `json:"MakeName"`
	MfrID    int64  `json:"MfrId"`
	MfrName  string `json:"MfrName"`
}

type VehicleTypeNameResult struct {
	MakeID          int64  `json:"MakeId"`
	MakeName        string `json:"MakeName"`
	VehicleTypeID   int64  `json:"VehicleTypeId"`
	VehicleTypeName string `json:"VehicleTypeName"`
}

type VehicleTypeIdResult struct {
	VehicleTypeID   int64  `json:"VehicleTypeId"`
	VehicleTypeName string `json:"VehicleTypeName"`
}

type GetEquipmentPlantCodesResult struct {
	Address       interface{} `json:"Address"`
	City          interface{} `json:"City"`
	Country       string      `json:"Country"`
	DotCode       string      `json:"DOTCode"`
	Name          string      `json:"Name"`
	OldDotCode    string      `json:"OldDotCode"`
	PostalCode    interface{} `json:"PostalCode"`
	StateProvince interface{} `json:"StateProvince"`
	Status        string      `json:"Status"`
}

type GetModelsResult struct {
	MakeID    int64  `json:"Make_ID"`
	MakeName  string `json:"Make_Name"`
	ModelID   int64  `json:"Model_ID"`
	ModelName string `json:"Model_Name"`
}
