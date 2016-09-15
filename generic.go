package list

type Generic struct {
	list []interface{}
}

func (l *Generic) AddRange(vals []interface{}) {
	l.list = append(l.list, vals...)
}

func (l *Generic) Add(val interface{}) {
	l.list = append(l.list, val)
}

func (l *Generic) InsertAt(i int32, val interface{}) {
	l.list = append(l.list, 0)
	copy(l.list[i+1:], l.list[i:])
	l.list[i] = val
}

func (l *Generic) InsertRangeAt(i int32, vals []interface{}) {
	l.list = append(l.list[:i], append(vals, l.list[i:]...)...)
}

func (l *Generic) InsertLast(val interface{}) {
	l.Add(val)
}

func (l *Generic) GetAt(i int32) interface{} {
	return l.list[i]
}

func (l *Generic) GetRange(begin int32, end int32) []interface{} {
	return l.list[begin : end+1]
}

func (l *Generic) GetFirst() interface{} {
	return l.list[0]
}

func (l *Generic) GetLast() interface{} {
	return l.list[len(l.list)-1]
}

func (l *Generic) RemoveAt(i int32) {
	l.list = append(l.list[:i], l.list[i+1:]...)
}

func (l *Generic) RemoveFirst() {
	l.list = l.list[1:]
}

func (l *Generic) RemoveLast() {
	l.list = l.list[:len(l.list)-1]
}

// func (l *Generic) RemoveRange(begin int32, end int32) {
// 	//TODO
// }

func (l *Generic) Count() int {
	return len(l.list)
}

func (l *Generic) Clear() {
	l.list = make([]interface{}, 0)
}

func (l *Generic) Me() []interface{} {
	return l.list
}

func (l *Generic) Filter(cond func(interface{}) bool) []interface{} {
	c := l.list[:0]
	for _, x := range l.list {
		if cond(x) {
			c = append(c, x)
		}
	}
	return c
}

func (l *Generic) Exists(cond func(interface{}) bool) bool {
	for _, x := range l.list {
		if cond(x) {
			return true
		}
	}
	return false
}

func (l *Generic) Find(cond func(interface{}) bool) interface{} {
	for _, x := range l.list {
		if cond(x) {
			return x
		}
	}
	return nil
}

func (l *Generic) FindAll(cond func(interface{}) bool) []interface{} {
	return l.Filter(cond)
}

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

func (l *Generic) FindIndex(cond func(interface{}) bool) int {
	index := -1

	for k, x := range l.list {
		if cond(x) {
			index = k
		}
	}
	return index
}

func (l *Generic) FindLastIndex(cond func(interface{}) bool) int {
	index := -1

	for k, x := range l.list {
		if cond(x) {
			index = k
		}
	}
	return index
}

func (l *Generic) TrueForAll(cond func(interface{}) bool) bool {
	for _, x := range l.list {
		if !cond(x) {
			return false
		}
	}
	return true
}

func (l *Generic) Reverse() {
	for i := len(l.list)/2 - 1; i >= 0; i-- {
		opp := len(l.list) - 1 - i
		l.list[i], l.list[opp] = l.list[opp], l.list[i]
	}
}

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
