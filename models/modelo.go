package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome string `json:"nome" validate:"nonzero"`
	CPF  string `json:"cpf" validate:"len=11, regexp=^[0-9]*$"`
	RG   string `json:"rg" validate:"len=9, regexp=^[0-9]*$"`
}

func ValidaDados(aluno *Aluno) error{
	err := validator.Validate(aluno)
	if err != nil{
		return err
	}
	return nil

}
