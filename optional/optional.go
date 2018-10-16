package optional

import "time"

type String struct {
	set   bool
	value string
}

func (o *String) IsSet() bool {
	return o.set
}

func (o *String) Value() string {
	return o.value
}

func (o *String) Set(v string) {
	o.set = true
	o.value = v
}

type Bool struct {
	set   bool
	value bool
}

func (o *Bool) IsSet() bool {
	return o.set
}

func (o *Bool) Value() bool {
	return o.value
}

func (o *Bool) Set(v bool) {
	o.set = true
	o.value = v
}

type Int struct {
	set   bool
	value int
}

func (o *Int) IsSet() bool {
	return o.set
}

func (o *Int) Value() int {
	return o.value
}

func (o *Int) Set(v int) {
	o.set = true
	o.value = v
}

type Float64 struct {
	set   bool
	value float64
}

func (o *Float64) IsSet() bool {
	return o.set
}

func (o *Float64) Value() float64 {
	return o.value
}

func (o *Float64) Set(v float64) {
	o.set = true
	o.value = v
}

type Time struct {
	set   bool
	value time.Time
}

func (o *Time) IsSet() bool {
	return o.set
}

func (o *Time) Value() time.Time {
	return o.value
}

func (o *Time) Set(v time.Time) {
	o.set = true
	o.value = v
}
