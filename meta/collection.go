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

// FindBy filter
func (mc *Collection) FindBy(f Filter) Collection {
	found := Collection{}

	for _, m := range *mc {
		if fr := f(m); fr {
			found = append(found, m)
		}
	}

	return found
}

// Filter function signature
type Filter func(d *Data) bool

// TokenNameFilter generates filter function that returns true
// if data.TokenName has in the arguments list
func TokenNameFilter(tokenNames ...TokenName) Filter {
	return func(d *Data) bool {
		for _, tn := range tokenNames {
			if d.TokenName == tn {
				return true
			}
		}
		return false
	}
}

// TypeFilter generates filter function that returns true
// if data.Type has in the arguments list
func TypeFilter(types ...Type) Filter {
	return func(d *Data) bool {
		for _, t := range types {
			if d.Type == t {
				return true
			}
		}
		return false
	}
}

// ValueFilter generates filter function that returns true
// if data.Value has in the arguments list
func ValueFilter(values ...string) Filter {
	return func(d *Data) bool {
		for _, v := range values {
			if d.Value == v {
				return true
			}
		}
		return false
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

// StopOnFailureFilter always returns false after first failure
func StopOnFailureFilter(f Filter) Filter {
	stopFlag := false
	return func(d *Data) bool {
		if stopFlag == true {
			return false
		}

		if !f(d) {
			stopFlag = true
			return false
		}

		return true
	}
}
