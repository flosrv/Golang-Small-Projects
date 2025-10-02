// models.go
package main

type Movie struct {
	ID       uint     `gorm:"primaryKey"` // clé primaire auto-incrément
	Isbn     string   `json:"isbn"`
	Title    string   `json:"title"`
	Director Director `gorm:"embedded"` // embedded struct pour relations simples
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
