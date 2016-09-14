package list

import (
	"strconv"
	"testing"
)

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

	i := 0
	for i < 3 {
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

func TestInsertLast(t *testing.T) {
	l := Generic{}
	l.InsertLast("1")
	l.InsertLast("2")
	ls := l.ToSliceString()

	if ls[0] != "1" {
		t.Errorf("expect %s to be %s", ls[0], "1")
	}

	if ls[1] != "2" {
		t.Errorf("expect %s to be %s", ls[1], "2")
	}
}

func TestGet(t *testing.T) {
	l := Generic{}
	l.Add("1")
	l.Add("2")
	l.Add("3")

	if l.GetFirst().(string) != "1" {
		t.Errorf("expect %s to be %s", l.GetFirst().(string), "1")
	}

	if l.GetAt(1).(string) != "2" {
		t.Errorf("expect %s to be %s", l.GetAt(1).(string), "2")
	}

	if l.GetLast().(string) != "3" {
		t.Errorf("expect %s to be %s", l.GetLast().(string), "3")
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
