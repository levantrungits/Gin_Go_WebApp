package main

import "time"

// Employee
type Employee struct {
	ID        uint
	FirstName string
	LastName  string
	StartDate time.Time
	Position  string
	TotalPTO  float32
	Status    string
	TimesOff  []TimeOff
}

// TimeOff
type TimeOff struct {
	Type      string
	Amount    float32
	StartDate time.Time
	Status    string
}

// Employees
var Employees = map[string]Employee{
	"962134": Employee{
		ID:        962134,
		FirstName: "Jennifer",
		LastName:  "Watson",
		StartDate: time.Now().Add(-13 * time.Hour * 24 * 365),
		Position:  "CEO",
		TotalPTO:  30,
		Status:    "Active",
	},
	"176158": Employee{
		ID:        176158,
		FirstName: "Allison",
		LastName:  "Jane",
		StartDate: time.Now().Add(-4 * time.Hour * 24 * 365),
		Position:  "COO",
		TotalPTO:  30,
		Status:    "Active",
	},
	"297365": Employee{
		ID:        297365,
		FirstName: "Jonathon",
		LastName:  "Anderson",
		StartDate: time.Now().Add(-12 * time.Hour * 24 * 365),
		Position:  "Worker Bee",
		TotalPTO:  30,
		Status:    "Active",
	},
}

// TimesOff
var TimesOff = map[string][]TimeOff{
	"962134": []TimeOff{
		{
			Type:      "Holiday",
			Amount:    8.,
			StartDate: time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC),
			Status:    "Taken",
		},
		{
			Type:      "PTO",
			Amount:    16.,
			StartDate: time.Date(2016, 8, 16, 0, 0, 0, 0, time.UTC),
			Status:    "Scheduled",
		},
		{
			Type:      "PTO",
			Amount:    16.,
			StartDate: time.Date(2016, 12, 8, 0, 0, 0, 0, time.UTC),
			Status:    "Requested",
		},
	},
}
