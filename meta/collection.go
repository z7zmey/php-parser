package meta

// Collection of Meta objects
type Collection []*Data

// SetTokenName sets TokenName for the all elements in the collection
func (mc *Collection) SetTokenName(tn TokenName) *Collection {
	for _, m := range *mc {
		m.TokenName = tn
	}

	return mc
}

// Push adds elements to the end of an Collection
func (mc *Collection) Push(mm ...*Data) *Collection {
	*mc = append(*mc, mm...)
	return mc
}

// Unshift prepends elements to the beginning of an Collection
func (mc *Collection) Unshift(mm ...*Data) *Collection {
	*mc = append(mm, *mc...)
	return mc
}

// AppendTo - appends elements of the collection to the end of the target collection
func (mc *Collection) AppendTo(target *Collection) *Collection {
	if len(*mc) == 0 {
		return mc
	}
	*target = append(*target, *mc...)
	return mc
}

// PrependTo - prepends elements of the collection to the start of the target collection
func (mc *Collection) PrependTo(target *Collection) *Collection {
	if len(*mc) == 0 {
		return mc
	}
	*target = append(*mc, *target...)
	return mc
}

// Cut elements by TokenName
func (mc *Collection) Cut(f Filter) *Collection {
	collection := (*mc)[:0]
	cutted := Collection{}

	for _, m := range *mc {
		if fr := f(m); fr {
			cutted = append(cutted, m)
		} else {
			collection = append(collection, m)
		}
	}

	*mc = collection

	return &cutted
}

// Filter function signature
type Filter func(d *Data) bool

// TokenNameFilter generates filter function that returns true
// if data.TokenName exactly same as given
func TokenNameFilter(tn TokenName) Filter {
	return func(d *Data) bool {
		return d.TokenName == tn
	}
}

// TypeFilter generates filter function that returns true
// if data.Type exactly same as given
func TypeFilter(t Type) Filter {
	return func(d *Data) bool {
		return d.Type == t
	}
}

// AndFilter generates filter function that returns true
// if all given filters return true
func AndFilter(filters ...Filter) Filter {
	return func(d *Data) bool {
		for _, filter := range filters {
			if result := filter(d); !result {
				return false
			}
		}

		return true
	}
}

// OrFilter generates filter function that returns true
// if one of given filters return true
func OrFilter(filters ...Filter) Filter {
	return func(d *Data) bool {
		for _, filter := range filters {
			if result := filter(d); result {
				return true
			}
		}

		return false
	}
}

// NotFilter negates given filter
func NotFilter(f Filter) Filter {
	return func(d *Data) bool {
		return !f(d)
	}
}
