package calendar

import (
	"errors"
)

type Date struct {
	year  int
	month int
	day   int
}

// SetYear D *Date is like this in python
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("year cannot be less than 1")
	}
	d.year = year
	return nil
}
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("month number must be between 1 and 12")
	}
	d.month = month
	return nil
}
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("day number must be between 1 and 31")
	}
	d.day = day
	return nil
}

// Year getter example. stylish you're not allowed to name it like GetYear
func (d *Date) Year() int {
	return d.year
}
func (d *Date) Month() int {
	return d.month
}
func (d *Date) Day() int {
	return d.day
}
