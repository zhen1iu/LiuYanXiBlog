package convert

import "strconv"

type StrTo string

func (s StrTo) String() string {
	return string(s)
}

func (s StrTo) Int() (int, error) {
	num, _ := strconv.Atoi(s.String())
	return num, nil
}

func (s StrTo) MustInt() int {
	num, _ := s.Int()
	return num
}

func (s StrTo) Uint32() (uint32, error) {
	num, _ := s.Uint32()
	return num, nil
}

func (s StrTo) MustUint32() uint32 {
	num, _ := s.Uint32()
	return num
}
