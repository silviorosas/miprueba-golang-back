package models

type Persona struct{


	Id int64 `json:"id"`
	Nombre string `json:"nombre"`
	Apellido string `json:"apellido"`
	Direccion string `json:"direccion"`
	Telefono string `json:"telefono"`

}