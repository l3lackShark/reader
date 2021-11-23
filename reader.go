package reader

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
)

type Reader struct {
	reader *bufio.Reader
}

func New() *Reader {
	return &Reader{}
}

func (r *Reader) read(field reflect.Value) {
	for i := 0; i < field.NumField(); i++ {
		structField := field.Type().Field(i)
		setField := field.Field(i)
		switch structField.Type.Kind() { //TODO: Figure out a way to simplify this (interface{}?)
		case reflect.Bool:
			var val byte
			binary.Read(r.reader, binary.LittleEndian, &val)
			out := val == 1
			actVal := reflect.ValueOf(out)
			setField.Set(actVal)
		case reflect.Int8:
			var val int8
			binary.Read(r.reader, binary.LittleEndian, &val)
			actVal := reflect.ValueOf(val)
			setField.Set(actVal)
		case reflect.Float32:
			var val float32
			binary.Read(r.reader, binary.LittleEndian, &val)
			actVal := reflect.ValueOf(val)
			setField.Set(actVal)
		case reflect.Float64:
			var val float64
			binary.Read(r.reader, binary.LittleEndian, &val)
			actVal := reflect.ValueOf(val)
			setField.Set(actVal)
		case reflect.Uint8:
			var val uint8
			binary.Read(r.reader, binary.LittleEndian, &val)
			actVal := reflect.ValueOf(val)
			setField.Set(actVal)
		case reflect.Int16:
			var val int16
			binary.Read(r.reader, binary.LittleEndian, &val)
			actVal := reflect.ValueOf(val)
			setField.Set(actVal)
		case reflect.Uint16:
			var val uint16
			binary.Read(r.reader, binary.LittleEndian, &val)
			actVal := reflect.ValueOf(val)
			setField.Set(actVal)
		case reflect.Int32:
			var val int32
			binary.Read(r.reader, binary.LittleEndian, &val)
			actVal := reflect.ValueOf(val)
			setField.Set(actVal)
		case reflect.Int64:
			var val int64
			binary.Read(r.reader, binary.LittleEndian, &val)
			actVal := reflect.ValueOf(val)
			setField.Set(actVal)
		case reflect.String:
			var startingByte byte
			binary.Read(r.reader, binary.LittleEndian, &startingByte)
			switch startingByte {
			case 0x0:
				continue
			case 0x0B:
				strlen, err := readVarUint(r.reader, 32)
				check(err)
				if strlen > 1000000 {
					panic("malformed/old database")
				}
				stringBytes := make([]byte, strlen)
				_, err = io.ReadFull(r.reader, stringBytes)
				check(err)
				actVal := reflect.ValueOf(string(stringBytes))
				setField.Set(actVal)
			}
		case reflect.Slice:
			var sliceLen int32
			binary.Read(r.reader, binary.LittleEndian, &sliceLen)
			if sliceLen > 1000000 {
				panic("malformed/old database")
			}
			if sliceLen == -1 { //no data
				continue
			}
			newSlice := reflect.MakeSlice(structField.Type, int(sliceLen), int(sliceLen))
			setField.Set(newSlice)
			//recursively populate slice till the end of all embedded structs
			for i := 0; i < newSlice.Len(); i++ {
				r.read(newSlice.Index(i))
			}
		default:
			panic(fmt.Sprintf("Unsupported type: %s", structField.Type.Kind()))
		}

	}
}

func (r *Reader) Read(file string, input interface{}) error {
	pval := reflect.ValueOf(input)
	val := reflect.Indirect(pval)

	if pval.Kind() != reflect.Ptr || val.Kind() != reflect.Struct {
		return fmt.Errorf("input must be a pointer to struct")
	}
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	r.reader = bufio.NewReader(f)
	r.read(val)
	return nil
}

func check(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}

func readVarUint(r io.Reader, n uint) (uint64, error) {
	if n > 64 {
		panic(errors.New("leb128: n must <= 64"))
	}
	p := make([]byte, 1)
	var res uint64
	var shift uint
	for {
		_, err := io.ReadFull(r, p)
		if err != nil {
			return 0, err
		}
		b := uint64(p[0])
		switch {
		// note: can not use b < 1<<n, when n == 64, 1<<n will overflow to 0
		case b < 1<<7 && b <= 1<<n-1:
			res += (1 << shift) * b
			return res, nil
		case b >= 1<<7 && n > 7:
			res += (1 << shift) * (b - 1<<7)
			shift += 7
			n -= 7
		default:
			return 0, errors.New("leb128: invalid uint")
		}
	}
}
