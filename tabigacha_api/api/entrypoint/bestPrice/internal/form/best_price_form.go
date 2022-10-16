package form

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"golang.org/x/xerrors"
)

type BestPriceForm struct {
	Origin      string `validate:"required"`
	Destination string `validate:"required"`
	Year        int    
	Month       int    
	Day         int    
}

func NewBestPriceForm(queryParams map[string]string) (*BestPriceForm, error) {
	origin := queryParams["origin"]
	destination := queryParams["destination"]
	year, _ := strconv.Atoi(queryParams["year"])
	month, _ := strconv.Atoi(queryParams["month"])
	day, _ := strconv.Atoi(queryParams["day"])

	bpf := &BestPriceForm{
		Origin:      origin,
		Destination: destination,
		Year:        year,
		Month:       month,
		Day:         day,
	}

	validate := validator.New()
	validate.RegisterStructValidation(dateValidation, BestPriceForm{})
	if err := validate.Struct(bpf); err != nil {
		return nil, xerrors.Errorf(err.Error())
	}

	return bpf, nil
}

func dateValidation(sl validator.StructLevel) {
	// yearもmonthもdayもなければ何もしない
	// yemonthがあってdayがない場合エラー
	// dayがあってmonthがない場合エラー
	form := sl.Current().Interface().(BestPriceForm)
	if form.Year == 0 && (form.Month != 0 || form.Day != 0) {
		sl.ReportError(form.Year, "Year", "year", "年のパラメータがないとき他のパラメータは受け付けられません。", "")
	}
	if form.Year != 0  {
		if form.Month != 0 && form.Month > 12 || form.Month < 1 {
			sl.ReportError(form.Month, "Month", "month", "月の指定が不正です。", "")
		}
		if form.Day != 0 && (form.Day > 31 || form.Day < 1) {
			sl.ReportError(form.Day, "Day", "day", "日の指定が不正です。", "")
		}
	} 
}
