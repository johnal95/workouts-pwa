package httpx

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator/v10"
)

type Parser struct {
	validate *validator.Validate
}

func NewParser(v *validator.Validate) *Parser {
	return &Parser{
		validate: v,
	}
}

func (p *Parser) ParseJSON(r io.Reader, target any) error {
	d := json.NewDecoder(r)
	d.DisallowUnknownFields()
	if err := d.Decode(target); err != nil {
		return BadRequest(err, "invalid request body", nil)
	}
	if err := p.validate.Struct(target); err != nil {
		return ValidationError(err, target)
	}
	return nil
}
