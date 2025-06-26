package domain

type Product struct {
	Id          string `json: "-"`
	Name        string `json: "name"`
	Description string `json: "desciption"`
	Price       int64  `json: "price"`
}
