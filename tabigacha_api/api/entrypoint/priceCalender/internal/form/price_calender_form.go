package form

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/xerrors"
)

type PriceCalenderForm struct {
	Origin      string `validate:"required"`
	Destination string `validate:"required"`
}

func NewPriceCalenderForm(queryParams map[string]string) (*PriceCalenderForm, error) {
	origin := queryParams["origin"]
	destination := queryParams["destination"]

	bpf := &PriceCalenderForm{
		Origin:      origin,
		Destination: destination,
	}

	validate := validator.New()
	if err := validate.Struct(bpf); err != nil {
		return nil, xerrors.Errorf(err.Error())
	}

	return bpf, nil
}
