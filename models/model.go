package models

type Tache struct {
	ID          int    
	Title       string
	Description string	
}

type Ms struct {
	Type  string      
	Size  int         
	Tasks any
	nbBlock int
}