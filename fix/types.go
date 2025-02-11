package fix

import (
	"fmt"
	"strconv"
	"time"
)

// Value basic methods for work with FIX-value
type Value interface {
	// ToBytes converts value to bytes
	ToBytes() []byte

	// FromBytes parse value from bytes
	FromBytes([]byte) error

	// Value returns value
	Value() interface{}

	// String convert Value to string
	String() string

	// IsNull check is empty value
	IsNull() bool

	// Set replace value with same type
	Set(d interface{}) error
}

// Raw data represented by bytes
type Raw struct {
	value []byte
}

// NewRaw
func NewRaw(v []byte) *Raw {
	return &Raw{
		value: v,
	}
}

func (v *Raw) ToBytes() []byte {
	return v.value
}

func (v *Raw) FromBytes(d []byte) (err error) {
	v.value = d
	return nil
}

func (v *Raw) IsNull() bool {
	return v.value == nil
}

func (v *Raw) Value() interface{} {
	return v.value
}

// Set set value from []byte
func (v *Raw) Set(d interface{}) error {
	if res, ok := d.([]byte); ok {
		v.value = res
		return nil
	}

	return fmt.Errorf("could not use %s as type %s", d, "Raw")
}

func (v *Raw) String() string {
	return string(v.value)
}

// String
type String struct {
	value string
	valid bool
}

func NewString(v string) *String {
	return &String{value: v, valid: true}
}

// Set set value from string
func (v *String) Set(d interface{}) error {
	if d == nil {
		v.valid = false
		return nil
	}

	if res, ok := d.(string); ok {
		v.value = res
		v.valid = true
		return nil
	}

	return fmt.Errorf("could not use %s as type %s", d, "String")
}

func (v *String) ToBytes() []byte {
	if !v.valid || v.value == "" {
		return nil
	}
	return []byte(v.value)
}

func (v *String) IsNull() bool {
	return !v.valid
}

func (v *String) Value() interface{} {
	return v.value
}

func (v *String) FromBytes(d []byte) (err error) {
	if d == nil {
		v.valid = false
		return nil
	}

	v.valid = true
	v.value = string(d)

	return nil
}

func (v *String) String() string {
	return v.value
}

type Int struct {
	value int
	valid bool
}

func NewInt(value int) *Int {
	return &Int{value: value, valid: true}
}

func (v *Int) IsNull() bool {
	return !v.valid
}

// Set set value from int
func (v *Int) Set(d interface{}) error {
	if d == nil {
		v.valid = false
		return nil
	}

	if res, ok := d.(int); ok {
		v.value = res
		v.valid = true
		return nil
	}

	return fmt.Errorf("could not use %s as type %s", d, "Int")
}

func (v *Int) Value() interface{} {
	return v.value
}
func (v *Int) String() string {
	return strconv.Itoa(v.value)
}

func (v *Int) FromBytes(d []byte) (err error) {
	if d == nil {
		v.valid = false
		return nil
	}

	v.valid = true
	v.value, err = strconv.Atoi(string(d))

	return err
}

func (v *Int) ToBytes() []byte {
	if !v.valid {
		return nil
	}
	return []byte(strconv.Itoa(v.value))
}

// Uint
type Uint struct {
	value uint64
	valid bool
}

// NewUint
func NewUint(value uint64) *Uint {
	return &Uint{value: value}
}

// Set set value from uint64
func (v *Uint) Set(d interface{}) error {
	if d == nil {
		v.valid = false
		return nil
	}

	if res, ok := d.(uint64); ok {
		v.value = res
		v.valid = true
		return nil
	}

	return fmt.Errorf("could not use %s as type %s", d, "Uint")
}

func (v *Uint) IsNull() bool {
	return !v.valid
}

func (v *Uint) FromBytes(d []byte) (err error) {
	if d == nil {
		v.valid = false
		return nil
	}

	v.valid = true
	v.value, err = strconv.ParseUint(string(d), 10, 64)

	return err
}

func (v *Uint) Value() interface{} {
	return v.value
}

func (v *Uint) String() string {
	return fmt.Sprintf("%d", v.value)
}

func (v *Uint) ToBytes() []byte {
	if !v.valid {
		return nil
	}
	return []byte(strconv.FormatUint(v.value, 10))
}

// Float
type Float struct {
	value float64
	valid bool
}

// NewFloat
func NewFloat(value float64) *Float {
	return &Float{value: value}
}

func (v *Float) IsNull() bool {
	return !v.valid
}

func (v *Float) Value() interface{} {
	return v.value
}

func (v *Float) FromBytes(d []byte) (err error) {
	if d == nil {
		v.valid = false
		return nil
	}

	v.valid = true
	v.value, err = strconv.ParseFloat(string(d), 64)

	return err
}

func (v *Float) ToBytes() []byte {
	if !v.valid {
		return nil
	}
	return []byte(strconv.FormatFloat(v.value, 'f', -1, 64))
}

func (v *Float) String() string {
	return fmt.Sprintf("%f", v.value)
}

// Set set value from float64
func (v *Float) Set(d interface{}) error {
	if d == nil {
		v.valid = false
		return nil
	}

	if res, ok := d.(float64); ok {
		v.value = res
		v.valid = true
		return nil
	}

	return fmt.Errorf("could not use %s as type %s", d, "Float")
}

// Time
type Time struct {
	value time.Time
	valid bool
}

// NewTime
func NewTime(value time.Time) *Time {
	return &Time{value: value, valid: true}
}

// Set set value from time.Time
func (v *Time) Set(d interface{}) error {
	if d == nil {
		v.valid = false
		return nil
	}

	if res, ok := d.(time.Time); ok {
		v.value = res
		v.valid = true
		return nil
	}

	return fmt.Errorf("could not use %s as type %s", d, "Time")
}

func (v *Time) IsNull() bool {
	return !v.valid
}

func (v *Time) Value() interface{} {
	return v.value
}

func (v *Time) ToBytes() []byte {
	if !v.valid {
		return nil
	}
	return []byte(v.value.Format(TimeLayout)) // TODO set layout outside
}

func (v *Time) FromBytes(d []byte) (err error) {
	if d == nil {
		v.valid = false
		return nil
	}

	v.valid = true
	v.value, err = time.Parse(TimeLayout, string(d))

	return err
}

func (v *Time) String() string {
	return v.value.Format(TimeLayout)
}

const (
	True  = "Y"
	False = "N"
)

// Bool
type Bool struct {
	value bool
	valid bool
}

func (v *Bool) ToBytes() []byte {
	if !v.valid {
		return nil
	}

	if v.value {
		return []byte(True)
	}
	return []byte(False)
}

func (v *Bool) FromBytes(d []byte) error {
	if d == nil {
		v.valid = false
		return nil
	}

	v.valid = true
	v.value = string(d) == True

	return nil
}

func (v *Bool) Value() interface{} {
	return v.value
}

func (v *Bool) String() string {
	if !v.valid {
		return ""
	}

	if v.value {
		return True
	}
	return False
}

func (v *Bool) IsNull() bool {
	return !v.valid
}

// Set set value from bool
func (v *Bool) Set(d interface{}) error {
	if d == nil {
		v.valid = false
		return nil
	}

	if res, ok := d.(bool); ok {
		v.value = res
		v.valid = true
		return nil
	}

	return fmt.Errorf("could not use %s as type %s", d, "Bool")
}
