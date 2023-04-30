package enum

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

var ErrEnumParse = errors.New("unable to parse Enum")

/*
	type Enumer interface {
		fmt.Stringer
		fmt.GoStringer
		json.Marshaler
		json.Unmarshaler
	}
*/
var (
	typeValStrMaps = sync.Map{}
	typeStrValMaps = sync.Map{}

	ErrEnumAlreadyRegistered = errors.New("Enum already registered")
	ErrEnumNotRegistered     = errors.New("Enum not registered")
	ErrEnumBadRegistration   = errors.New("Enum bad registration")
	ErrEnumValIsEmpty        = errors.New("Enum value is empty")
	ErrEnumStrIsEmpty        = errors.New("Enum string is empty")
	ErrEnumMustUniq          = errors.New("Enum value and string must be unique")
)

/*
type SafeSet[T comparable] struct {
    Values map[T]bool
}

func NewSafeSet[T comparable]() SafeSet[T] {
    return SafeSet[T]{Values: make(map[T]bool)}
}
*/

type Enumer interface {
	comparable
}

type Enum[T comparable] struct {
	V T
}

func NewEnum[T comparable](v T) Enum[T] {
	return Enum[T]{V: v}
}

func (e Enum[T]) String() string {
	return fmt.Sprintf("%v(%T)", e.V, e.V)
}

/*
func (e Enum[T]) String() string {
	var val0 T
	typeName := reflect.TypeOf(val0).String()
	if vs, has := typeValStrMaps.Load(typeName); !has {
		panic(fmt.Errorf("%w: (%s)", ErrEnumNotRegistered, typeName))
	} else {
		if valStr, is := vs.(map[T]string); !is {
			panic(fmt.Errorf("%w: (%s)", ErrEnumBadRegistration, typeName))
		} else {
			return valStr[val0]
		}
	}
	//return ""
}
*/

/*
	func String[T comparable]() string {
		var val0 T
		typeName := reflect.TypeOf(val0).String()
		if vs, has := typeValStrMaps.Load(typeName); !has {
			panic(fmt.Errorf("%w: (%s)", ErrEnumNotRegistered, typeName))
		} else {
			if valStr, is := vs.(map[T]string); !is {
				panic(fmt.Errorf("%w: (%s)", ErrEnumBadRegistration, typeName))
			} else {
				return valStr[val0]
			}
		}
		//return ""
	}
*/
/*
func (e *Enum[T]) GoString() string {
	var val0 T
	typeName := reflect.TypeOf(val0).String()
	if vs, has := typeValStrMaps.Load(typeName); !has {
		panic(fmt.Errorf("%w: (%s)", ErrEnumNotRegistered, typeName))
	} else {
		if valStr, is := vs.(map[T]string); !is {
			panic(fmt.Errorf("%w: (%s)", ErrEnumBadRegistration, typeName))
		} else {
			return valStr[val0]
		}
	}
	//return ""
}
*/
/*
	func (e Enum[T]) MarshalJSON() ([]byte, error) {
		return String2Bytes(rs.String()), nil
	}

	func (e *Enum[T]) UnmarshalJSON(bs []byte) (err error) {
		*rs, err = EnumParse(remoteStatusStrings, Bytes2String(bs))
		return err
	}
*/

func RegisterEnum[T comparable](vs map[T]string) error {
	var val0 T
	typeName := reflect.TypeOf(val0).String()
	valStr := map[T]string{}
	strVal := map[string]T{}

	if _, has := typeValStrMaps.Load(typeName); has {
		return fmt.Errorf("%w: (%s)", ErrEnumAlreadyRegistered, typeName)
	}

	uniqVals := map[T]bool{}
	uniqStrs := map[string]bool{}
	for val, str := range vs {
		if val == val0 {
			return fmt.Errorf("%w: (%s)", ErrEnumValIsEmpty, typeName)
		}
		if str == "" {
			return fmt.Errorf("%w: (%s)", ErrEnumStrIsEmpty, typeName)
		}
		uniqVals[val] = true
		uniqStrs[str] = true
		valStr[val] = str
		strVal[str] = val
	}
	if len(vs) != len(uniqVals) || len(vs) != len(uniqStrs) || len(vs) != len(valStr) || len(vs) != len(strVal) {
		return fmt.Errorf("%w: (%s)", ErrEnumMustUniq, typeName)
	}

	typeValStrMaps.Store(typeName, valStr)
	typeStrValMaps.Store(typeName, strVal)

	return nil
}

// *****************************************

type EnumBase int

func (E EnumBase) String() string {
	return fmt.Sprintf("(%T)", E)
}

func (E *EnumBase) GoString() string {
	return fmt.Sprintf("(%T)", E)
}

type Stat EnumBase

const (
	statInvalid Stat = iota + 1
	statAvailable
	statMissing
)

const (
	eInvalid EnumBase = iota + 1
)

type status int

type Status Enum[int]

const (
	statusInvalid status = iota + 1
	statusAvailable
	statusMissing
)

/*
const (
	StatI = NewEnum(statusInvalid)
)
*/
/*
func NewEnum(v status) Enum[status] // func[T comparable](v T) Enum[T]
*/

var (
	StatusInvalid = NewEnum(statusInvalid)
)

func init() {
	fmt.Println(eInvalid.String())
	//fmt.Println(statInvalid.GoString())
	fmt.Println(StatusInvalid.String())
}

/*
func init() {
	//z := NewEnum[status](int(statusAvailable))
	if err := RegisterEnum(map[Status]string{
		statusInvalid:   "invalid",
		statusAvailable: "available",
		statusMissing:   "missing",
	}); err != nil {
		panic(err)
	}
}
*/
