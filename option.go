package goption

// Float64Option describes Option[float64]
type Float64Option struct {
	value float64
	isSet bool
}

// NewFloat64Some returns an Option with existing value val
func NewFloat64Some(val float64) Float64Option {
	return Float64Option{
		value: val,
		isSet: true,
	}
}

// NewFloat64None returns an Option with no value set
func NewFloat64None() Float64Option {
	return Float64Option{}
}

// IsSet returns whether the option is set
func (o Float64Option) IsSet() bool {
	return o.isSet
}

// Get returns a value and a flag whether the value is set
func (o Float64Option) Get() (float64, bool) {
	return o.value, o.isSet
}

// GetOrElse returns value if set and `def` if not
func (o Float64Option) GetOrElse(def float64) float64 {
	if o.isSet {
		return o.value
	}
	return def
}

// ToSlice converts Option to a slice
// it returns an empty slice if option has not value set
// it returns a slice with 1 element otherwise
func (o Float64Option) ToSlice() []float64 {
	if o.isSet {
		return []float64{o.value}
	}
	return []float64{}
}

// Reset resets the Option to "no value" state
func (o *Float64Option) Reset() {
	var zero float64
	o.isSet = false
	o.value = zero
}

// Set sets the Option with value `val`
func (o *Float64Option) Set(val float64) {
	o.isSet = true
	o.value = val
}
