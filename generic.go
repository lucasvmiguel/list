// Package list provides handle with slice easily
package list

// Generic struct to create a generic slice
type Generic struct {
	list []interface{}
}

//AddRange Adds the elements of the specified collection to the end of the List
func (l *Generic) AddRange(vals []interface{}) {
	l.list = append(l.list, vals...)
}

// Add Adds an object to the end of the List
func (l *Generic) Add(val interface{}) {
	l.list = append(l.list, val)
}

//InsertAt Inserts an element into the List at the specified index
func (l *Generic) InsertAt(i int32, val interface{}) {
	l.list = append(l.list, 0)
	copy(l.list[i+1:], l.list[i:])
	l.list[i] = val
}

//InsertRangeAt Inserts the elements of a collection into the List at the specified index
func (l *Generic) InsertRangeAt(i int32, vals []interface{}) {
	l.list = append(l.list[:i], append(vals, l.list[i:]...)...)
}

// InsertLast Insert element at the end of the List (equal do add)
func (l *Generic) InsertLast(val interface{}) {
	l.Add(val)
}

// GetAt Get element by index
func (l *Generic) GetAt(i int32) interface{} {
	return l.list[i]
}

// GetRange Get the elements by range
func (l *Generic) GetRange(begin int32, end int32) []interface{} {
	return l.list[begin : end+1]
}

// GetFirst Get the first element in List
func (l *Generic) GetFirst() interface{} {
	return l.list[0]
}

// GetLast Get the last element in List
func (l *Generic) GetLast() interface{} {
	return l.list[len(l.list)-1]
}

// RemoveAt Removes the element at the specified index of the List
func (l *Generic) RemoveAt(i int32) {
	l.list = append(l.list[:i], l.list[i+1:]...)
}

// RemoveFirst Removes the first element in List
func (l *Generic) RemoveFirst() {
	l.list = l.list[1:]
}

// RemoveLast Removes the last element in List
func (l *Generic) RemoveLast() {
	l.list = l.list[:len(l.list)-1]
}

// RemoveRange Removes a range of elements from the List
// func (l *Generic) RemoveRange(begin int32, end int32) {
// 	//TODO
// }

// Count Count number of elements in List
func (l *Generic) Count() int {
	return len(l.list)
}

// Clear Removes all elements from the List
func (l *Generic) Clear() {
	l.list = make([]interface{}, 0)
}

// Values Get all values in List
func (l *Generic) Values() []interface{} {
	return l.list
}

// Filter Get List filtered
func (l *Generic) Filter(cond func(interface{}) bool) []interface{} {
	c := l.list[:0]
	for _, x := range l.list {
		if cond(x) {
			c = append(c, x)
		}
	}
	return c
}

// Exists Determines whether an element is in the List
func (l *Generic) Exists(cond func(interface{}) bool) bool {
	for _, x := range l.list {
		if cond(x) {
			return true
		}
	}
	return false
}

// Find Get element from List or nil
func (l *Generic) Find(cond func(interface{}) bool) interface{} {
	for _, x := range l.list {
		if cond(x) {
			return x
		}
	}
	return nil
}

// FindAll Get all elements from List (equal to Filter)
func (l *Generic) FindAll(cond func(interface{}) bool) []interface{} {
	return l.Filter(cond)
}

// FindLast Get last element from List by condition
func (l *Generic) FindLast(cond func(interface{}) bool) interface{} {
	index := -1

	for k, x := range l.list {
		if cond(x) {
			index = k
		}
	}

	if index != -1 {
		return l.list[index]
	}
	return nil
}

// FindIndex Get index of element from List by condition
func (l *Generic) FindIndex(cond func(interface{}) bool) int {
	index := -1

	for k, x := range l.list {
		if cond(x) {
			index = k
		}
	}
	return index
}

// FindLastIndex Get last index of element from List by condition
func (l *Generic) FindLastIndex(cond func(interface{}) bool) int {
	index := -1

	for k, x := range l.list {
		if cond(x) {
			index = k
		}
	}
	return index
}

// TrueForAll Determines whether every element in the List<T> matches the conditions defined by the specified predicate
func (l *Generic) TrueForAll(cond func(interface{}) bool) bool {
	for _, x := range l.list {
		if !cond(x) {
			return false
		}
	}
	return true
}

// Reverse Reverses the order of the elements in the specified range
func (l *Generic) Reverse() {
	for i := len(l.list)/2 - 1; i >= 0; i-- {
		opp := len(l.list) - 1 - i
		l.list[i], l.list[opp] = l.list[opp], l.list[i]
	}
}

// ToSliceString Return Slice String List
func (l *Generic) ToSliceString() []string {
	strList := make([]string, len(l.list))
	for k := range l.list {
		val, ok := l.list[k].(string)
		if ok {
			strList[k] = val
		}
	}
	return strList
}

// ToSliceInt Return Slice Int List
func (l *Generic) ToSliceInt() []int {
	intList := make([]int, len(l.list))
	for k := range l.list {
		val, ok := l.list[k].(int)
		if ok {
			intList[k] = val
		}
	}
	return intList
}

// ToSliceFloat64 Return Slice Float64 List
func (l *Generic) ToSliceFloat64() []float64 {
	float64List := make([]float64, len(l.list))
	for k := range l.list {
		val, ok := l.list[k].(float64)
		if ok {
			float64List[k] = val
		}
	}
	return float64List
}
