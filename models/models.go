package models

/*
App ...
*/
type App struct {
	Apps []RegisteredApp `json:"apps"`
}

/*
RegisteredApp ...
*/
type RegisteredApp struct {
	ID       string    `json:"app_id"`
	Products []Product `json:"products"`
}

/*
Product ...
*/
type Product struct {
	ID        string   `json:"product_id"`
	AvgRating float32  `json:"avg_rating"`
	Reviews   []Review `json:"reviews"`
}

/*
Review ...
*/
type Review struct {
	ID         int64   `json:"review_id"`
	Rating     int32   `json:"rating"`
	Reviewer   string  `json:"reviewer"`
	ReviewerID int64   `json:"reviewer_id"`
	Comments   Comment `json:"comment"`
}

/*
Comment ...
*/
type Comment struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}
