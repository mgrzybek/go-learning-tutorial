package data

//Post ...
type Post struct {
	UserID int    `json:"user_id,omitempty"`
	ID     int    `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Body   string `json:"body,omitempty"`
}

//User ...
type User struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	UserName string `json:"user_name,omitempty"`
	Email    string `json:"email,omitempty"`
}
