// Copyright 2015,2016,2017,2018,2019 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"

	"github.com/kasworld/prettystring"
)

type TestStruct struct {
	fieldInt       int
	fieldComplex   complex128
	fieldUint      uint
	fieldFloat     float64
	fieldString    string
	fieldBool      bool
	fieldChan      chan interface{}
	fieldArray     [14]string
	fieldSlice     []int
	fieldMap       map[int]string
	fieldPtr       *TestStruct
	fieldInterface interface{}
	fieldHidden    int        `prettystring:"hide"`
	fieldSimple    [14]string `prettystring:"simple"`
	fieldFn        func()
}

func newStruct() *TestStruct {
	return &TestStruct{
		fieldInt:     -34,
		fieldComplex: 14 + 34i,
		fieldUint:    12,
		fieldFloat:   35.2342,
		fieldString:  "hello world",
		fieldBool:    true,
		fieldChan:    make(chan interface{}, 10),
		// fieldArray : ,
		fieldSlice: make([]int, 3, 7),
		fieldMap: map[int]string{
			3: "asdf",
			6: "dsqer",
		},
		// fieldPtr : ,
		fieldFn: func() {},
	}
}

func main() {
	tw := newStruct()
	tw.fieldPtr = newStruct()
	tw.fieldInterface = newStruct()
	fmt.Println(prettystring.PrettyString(tw, 4))
}
