package model

type (
	User struct {
		Id       int     `json:"id"`
		Name     string  `json:"name"`
		Username string  `json:"username"`
		Email    string  `json:"email"`
		Address  Address `json:"address"`
		Phone    string  `json:"phone"`
		Website  string  `json:"website"`
		Company  Company `json:"company"`
	}

	Address struct {
		Street     string `json:"street"`
		Suite      string `json:"suite"`
		City       string `json:"city"`
		Zipcode    string `json:"zipcode"`
		Geographic Geo    `json:"geo"`
	}

	Geo struct {
		Latitude  string `json:"lat"`
		Longitude string `json:"lng"`
	}

	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Businesses  string `json:"bs"`
	}
)
