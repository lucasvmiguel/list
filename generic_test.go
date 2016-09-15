package list

import (
	"reflect"
	"strconv"
	"testing"
)

type StructTest struct {
	a int
	b string
}

func TestParse(t *testing.T) {
	def := []StructTest{
		StructTest{1, "Lucas"},
		StructTest{2, "Diego"},
	}

	l := Generic{}
	l.Add(StructTest{1, "Lucas"})
	l.Add(StructTest{2, "Diego"})

	list := l.Me()
	correctList := []StructTest{}

	for k := range list {
		correctList = append(correctList, list[k].(StructTest))
	}

	if !reflect.DeepEqual(def, correctList) {
		t.Error("expect equal lists")
	}
}

func TestReverse(t *testing.T) {
	l := Generic{}
	l.Add("1")
	l.Add("2")
	l.Add("3")
	l.Reverse()
	ls := l.ToSliceString()

	if ls[0] != "3" && ls[1] != "2" && ls[2] != "1" {
		t.Errorf("expect 3, 2 ,1. got: %s, %s, %s", ls[0], ls[1], ls[2])
	}
}

func TestAdd(t *testing.T) {
	l := Generic{}
	l.Add("1")
	l.Add("2")
	ls := l.ToSliceString()

	if ls[0] != "1" {
		t.Errorf("expect %s to be %s", ls[0], "1")
	}

	if ls[1] != "2" {
		t.Errorf("expect %s to be %s", ls[1], "2")
	}
}

func TestAddRange(t *testing.T) {
	l := Generic{}
	var rang []interface{}

	for i := 1; i < 3; i++ {
		rang = append(rang, strconv.Itoa(i))
	}

	l.AddRange(rang)

	ls := l.ToSliceString()

	if ls[0] != "1" {
		t.Errorf("expect %s to be %s", ls[0], "1")
	}

	if ls[1] != "2" {
		t.Errorf("expect %s to be %s", ls[1], "2")
	}
}

func TestInsertRangeAt(t *testing.T) {
	l := Generic{}
	var rang []interface{}

	l.InsertLast("a")
	l.InsertLast("b")
	l.InsertLast("e")
	l.InsertLast("f")

	rang = append(rang, "c")
	rang = append(rang, "d")

	l.InsertRangeAt(2, rang)

	ls := l.ToSliceString()

	if ls[0] != "a" && ls[1] != "b" && ls[2] != "c" && ls[3] != "d" && ls[4] != "e" && ls[5] != "f" {
		t.Errorf("expect a, b, c, d, e, f. got : %s, %s, %s, %s, %s, %s", ls[0], ls[1], ls[2], ls[3], ls[4], ls[5])
	}
}

func TestInsertAt(t *testing.T) {
	l := Generic{}

	l.InsertLast("a")
	l.InsertLast("b")
	l.InsertLast("d")

	l.InsertAt(2, "c")

	ls := l.ToSliceString()

	if ls[0] != "a" && ls[1] != "b" && ls[2] != "c" && ls[3] != "d" {
		t.Errorf("expect a, b, c, d. got : %s, %s, %s, %s", ls[0], ls[1], ls[2], ls[3])
	}
}

func TestInsertLast(t *testing.T) {
	l := Generic{}
	l.InsertLast(1.5)
	l.InsertLast(2.5)
	ls := l.ToSliceFloat64()

	if ls[0] != 1.5 {
		t.Errorf("expect %f to be %s", ls[0], "1.5")
	}

	if ls[1] != 2.5 {
		t.Errorf("expect %f to be %s", ls[1], "2.5")
	}
}

func TestGet(t *testing.T) {
	l := Generic{}
	l.Add("1")
	l.Add("2")
	l.Add("3")
	l.Add("4")
	l.Add("5")

	if l.GetFirst().(string) != "1" {
		t.Errorf("expect %s to be %s", l.GetFirst().(string), "1")
	}

	if l.GetAt(1).(string) != "2" {
		t.Errorf("expect %s to be %s", l.GetAt(1).(string), "2")
	}

	if l.GetLast().(string) != "5" {
		t.Errorf("expect %s to be %s", l.GetLast().(string), "5")
	}

	rang := l.GetRange(1, 3)

	if rang[0].(string) != "2" {
		t.Errorf("expect %s to be %s", rang[0], "2")
	}

	if rang[1].(string) != "3" {
		t.Errorf("expect %s to be %s", rang[1], "3")
	}

	if rang[2].(string) != "4" {
		t.Errorf("expect %s to be %s", rang[2], "4")
	}
}

func TestRemove(t *testing.T) {
	l := Generic{}
	l.Add("1")
	l.Add("2")
	l.Add("3")

	l.RemoveFirst()
	c1 := l.Count()
	ls1 := l.ToSliceString()
	if ls1[0] != "2" && c1 != 2 {
		t.Errorf("expect %s to be %s", ls1, "2")
		t.Errorf("expect count %d to be %d", c1, 2)
	}

	l.Add("6")
	l.RemoveLast()
	c2 := l.Count()
	ls2 := l.ToSliceString()
	if ls2[1] != "3" && c2 != 2 {
		t.Errorf("expect %s to be %s", ls2, "3")
		t.Errorf("expect count %d to be %d", c2, 2)
	}

	l.Add("9")
	l.RemoveAt(1)
	c3 := l.Count()
	ls3 := l.ToSliceString()
	if ls3[1] != "9" && c3 != 2 {
		t.Errorf("expect %s to be %s", ls3, "9")
		t.Errorf("expect count %d to be %d", c3, 2)
	}
}

func TestCount(t *testing.T) {
	l := Generic{}
	l.Add("1")
	l.Add("2")

	if l.Count() != 2 {
		t.Errorf("expect %d to be %d", l.Count(), 2)
	}
}

func TestClear(t *testing.T) {
	l := Generic{}
	l.Add("1")
	l.Add("2")
	l.Clear()

	emptySlice := make([]interface{}, 0)

	if !reflect.DeepEqual(l.Me(), emptySlice) {
		t.Error("expect equal lists")
	}
}

func filter(val interface{}) bool {
	num := val.(int)
	return num%10 == 0
}

func TestFilter(t *testing.T) {
	l := Generic{}
	l.Add(5)
	l.Add(10)
	l.Add(15)
	l.Add(20)

	list := l.Filter(filter)
	l.Clear()
	l.AddRange(list)

	listInt := l.ToSliceInt()

	if listInt[0] != 10 {
		t.Errorf("expect %d to be %d", listInt[0], 10)
	}

	if listInt[1] != 20 {
		t.Errorf("expect %d to be %d", listInt[1], 20)
	}
}

func TestFindAll(t *testing.T) {
	l := Generic{}
	l.Add(5)
	l.Add(10)
	l.Add(15)
	l.Add(20)

	list := l.FindAll(filter)
	l.Clear()
	l.AddRange(list)

	listInt := l.ToSliceInt()

	if listInt[0] != 10 {
		t.Errorf("expect %d to be %d", listInt[0], 10)
	}

	if listInt[1] != 20 {
		t.Errorf("expect %d to be %d", listInt[1], 20)
	}
}

func TestExists(t *testing.T) {
	l1 := Generic{}
	l1.Add(5)
	l1.Add(10)
	l1.Add(15)

	b1 := l1.Exists(filter)

	l2 := Generic{}
	l2.Add(1)
	l2.Add(2)
	l2.Add(3)

	b2 := l2.Exists(filter)

	if b1 != true {
		t.Errorf("expect %s to be true", b1)
	}

	if b2 != false {
		t.Errorf("expect %s to be false", b2)
	}
}

func TestFind(t *testing.T) {
	l1 := Generic{}
	l1.Add(5)
	l1.Add(10)
	l1.Add(15)

	b1 := l1.Find(filter)

	l2 := Generic{}
	l2.Add(80)
	l2.Add(2)
	l2.Add(3)

	b2 := l2.Find(filter)

	l3 := Generic{}
	l3.Add(1)
	l3.Add(2)
	l3.Add(3)

	b3 := l3.Find(filter)

	if b1.(int) != 10 {
		t.Errorf("expect %s to be 10", b1)
	}

	if b2.(int) != 80 {
		t.Errorf("expect %s to be 80", b2)
	}

	if b3 != nil {
		t.Errorf("expect %s to be nil", b3)
	}
}

func TestFindLast(t *testing.T) {
	l1 := Generic{}
	l1.Add(5)
	l1.Add(10)
	l1.Add(15)
	l1.Add(20)

	b1 := l1.FindLast(filter)

	l2 := Generic{}
	l2.Add(1)
	l2.Add(2)
	l2.Add(3)

	b2 := l2.FindLast(filter)

	if b1.(int) != 20 {
		t.Errorf("expect %s to be 10", b1)
	}

	if b2 != nil {
		t.Errorf("expect %s to be nil", b2)
	}
}

func TestFindIndex(t *testing.T) {
	l1 := Generic{}
	l1.Add(5)
	l1.Add(10)
	l1.Add(15)

	b1 := l1.FindIndex(filter)

	l2 := Generic{}
	l2.Add(80)
	l2.Add(2)
	l2.Add(3)

	b2 := l2.FindIndex(filter)

	if b1 != 1 {
		t.Errorf("expect %s to be 1", b1)
	}

	if b2 != 0 {
		t.Errorf("expect %s to be 0", b2)
	}
}

func TestFindLastIndex(t *testing.T) {
	l1 := Generic{}
	l1.Add(5)
	l1.Add(10)
	l1.Add(15)
	l1.Add(20)

	b1 := l1.FindLastIndex(filter)

	if b1 != 3 {
		t.Errorf("expect %s to be 3", b1)
	}
}

func TestFindTrueForAll(t *testing.T) {
	l1 := Generic{}
	l1.Add(10)
	l1.Add(20)
	l1.Add(30)

	b1 := l1.TrueForAll(filter)

	l2 := Generic{}
	l2.Add(10)
	l2.Add(20)
	l2.Add(30)
	l2.Add(35)

	b2 := l2.TrueForAll(filter)

	if b1 != true {
		t.Errorf("expect %s to be true", b1)
	}

	if b2 != false {
		t.Errorf("expect %s to be true", b2)
	}
}
