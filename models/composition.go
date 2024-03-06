package models

type Composition struct {
	Id   int    `json:"id"`   // id Склада
	Name string `json:"name"` // Название склада
	Type int    `json:"type"` // тип склада
}
