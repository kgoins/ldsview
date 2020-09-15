package internal

import "strings"

var placeholder = struct{}{}

type HashSetStr struct {
	values map[string]struct{}
}

func NewHashSetStr(initVals ...string) HashSetStr {
	numVals := len(initVals)
	set := HashSetStr{
		values: make(map[string]struct{}, len(initVals)),
	}

	if numVals > 0 {
		set.Add(initVals...)
	}

	return set
}

func (set HashSetStr) GetSingleValue() string {
	retVal := ""
	for val := range set.values {
		retVal = val
		break
	}

	return retVal
}

func (set *HashSetStr) Add(items ...string) {
	for _, item := range items {
		set.values[item] = placeholder
	}
}

func (set HashSetStr) Size() int {
	return len(set.values)
}

func (set HashSetStr) IsEmpty() bool {
	return len(set.values) == 0
}

func (set HashSetStr) Values() []string {
	values := make([]string, set.Size())

	count := 0
	for item := range set.values {
		values[count] = item
		count++
	}

	return values
}

func (set HashSetStr) Equals(set2 HashSetStr) bool {
	if set.Size() != set2.Size() {
		return false
	}

	for val := range set.values {
		if !set2.Contains(val) {
			return false
		}
	}

	return true
}

func (set HashSetStr) Contains(val string) bool {
	_, found := set.values[val]
	return found
}

func (set HashSetStr) ContainsSubstr(substr string) bool {
	for val := range set.values {
		if strings.Contains(val, substr) {
			return true
		}
	}

	return false
}

func (set HashSetStr) ContainsSubstrIgnoreCase(substr string) bool {
	substr = strings.ToLower(substr)

	for val := range set.values {
		if strings.Contains(strings.ToLower(val), substr) {
			return true
		}
	}

	return false
}

func (set *HashSetStr) Clear() {
	set.values = make(map[string]struct{})
}
