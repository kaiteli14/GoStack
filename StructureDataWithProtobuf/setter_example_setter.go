// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: setter_example.proto

package setter_example

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/travisjeffery/protoc-gen-setter"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (t *Person) Set(v string) {
	t.id = v
}

func (t *Person) Set(v string) {
	t.name = v
}

func (t *Animal) Set(v string) {
	t.name = v
}
