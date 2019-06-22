package basics

import (
	"fmt"
	"math"
	"math/cmplx"
	"unicode/utf8"
	"unsafe"

	"github.com/dustin/go-humanize"
)

// All calls code examples
func All() {
	integersAndFloatsAndComplex()
	stringsAndRunesAndBytes()
	loopsAndIf()
	deferExample()
}

func integersAndFloatsAndComplex() {

	fmt.Printf("\n=== Integers, Floats, Complex ===\n\n")

	var i = math.MaxInt64
	var i8min int8 = -1 << 7     // -128
	var i8max int8 = 1<<7 - 1    // 127
	var i16min int16 = -1 << 15  // -32768
	var i16max int16 = 1<<15 - 1 // 32767
	var i32min int32 = -1 << 31  // -2,147,483,648
	var i32max int32 = 1<<31 - 1 // 2,147,483,647
	var i64min int64 = 1<<63 - 1 // 9,223,372,036,854,775,803
	var i64max int64 = -1 << 63  // -9,223,372,036,854,775,808

	fmt.Printf("%d  %d  %d  %d  %d  %s  %s  %s  %s\n",
		i,
		i8min, i8max,
		i16min, i16max,
		humanize.Comma(int64(i32min)), humanize.Comma(int64(i32max)),
		humanize.Comma(i64min), humanize.Comma(i64max))

	var ui8min uint8 = 0
	var ui8max uint8 = 1<<8 - 1
	var ui16min uint16 = 0
	var ui16max uint16 = 1<<16 - 1
	var ui32min uint32 = 0
	var ui32max uint32 = 1<<32 - 1
	var ui64min uint64 = 0
	var ui64max uint64 = 1<<64 - 1

	fmt.Printf("%d  %d  %d  %d  %d  %d  %d  %d\n",
		ui8min, ui8max,
		ui16min, ui16max,
		ui32min, ui32max,
		ui64min, ui64max)

	var ptr uintptr = uintptr(unsafe.Pointer(new(string)))

	fmt.Printf("%v\n", ptr)

	var f = math.MaxFloat64
	var f32max float32 = math.MaxFloat32
	var f64max float64 = math.MaxFloat64

	fmt.Printf("%.4e  %.4e  %.4e\n", f, f32max, f64max)

	c := complex(3, 4)
	var c64 complex64 = complex(math.MaxFloat32, math.MaxFloat32)
	var c128 complex128 = complex(math.MaxFloat64, math.MaxFloat64)

	fmt.Printf("%v  %v  %v  %v  %v  %v\n", c, cmplx.Abs(c), c64, cmplx.Abs(complex128(c64)), c128, cmplx.Abs(c128))
}

func stringsAndRunesAndBytes() {
	fmt.Printf("\n=== Strings, Runes and Bytes ===\n\n")

	var s string = "george"
	var v byte = 1<<8 - 1
	var runeString string = `†¥ÓÂ`

	fmt.Printf("string: \"%s\"  string bytes:  % x  byte: %v\n", s, s, v)
	fmt.Printf("runeString: %s    % x    %+q    %d\n", runeString, runeString, runeString, utf8.RuneCountInString(runeString))

	var runes = []rune(runeString)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%x - %s\n", runes[i], string(runes[i]))
	}
}

func loopsAndIf() {

	fmt.Printf("\n=== Loops and If ===\n\n")

	// simple for loop with i declared outside of loop
	var i int
	for i = 0; i < 3; i++ {
		fmt.Printf("Iteration %d\n", i)
	}

	fmt.Println()

	// more compact with inline declaration of j
	for j := 3; j > 0; j-- {
		fmt.Printf("Iteration %d\n", j)
	}

	fmt.Println()

	// init and post statements can be omitted
	k := 3
	for k < 5 {
		fmt.Printf("Optional iteration %d\n", k)
		k++
	}

	fmt.Println()

	// remove semicolons and it is a while loop
	l := 3
	for l < 5 {
		fmt.Printf("Iterate like a while loop %d\n", l)
		l++
	}

	fmt.Println()

	// Infinite loop
	// for {}

	// Basic if statement
	n := "George"
	o := "George Campbell"

	if n == o {
		fmt.Printf("%s == %s\n", n, o)
	} else {
		fmt.Printf("%s != %s\n", n, o)
	}

	fmt.Println()

	// if with a pre statement - note p is available in the else statement too.
	q := 1
	if p := 2; p < q {
		fmt.Printf("%d < %d\n", p, q)
	} else {
		fmt.Printf("%d > %d\n", p, q)
	}

	fmt.Println()

	animal := "bear"

	// The switch statement
	switch animal {
	case "deer":
		fmt.Println("Bambi")
	case "bear":
		fmt.Println("Balloo")
	}

	fmt.Println()

	// The switch statement with no variable
	w := 1
	switch {
	case w == 1:
		fmt.Println("One")
	case w == 2:
		fmt.Println("Two")
	}

}

func deferExample() {

	fmt.Printf("\n=== Deferreds ===\n\n")

	a := func() int {
		fmt.Println("Begin a()")
		b := 1 // Defer uses the value of b when the statement is executed
		defer fmt.Printf("defer: b = %d\n", b)
		b = 2 // Change the value of b after the defer
		fmt.Printf("just before return: b = %d\n", b)
		b = 3 // Change the value of b one more time
		fmt.Println("End a()")
		return b
	}

	fmt.Printf("function return: b = %d\n", a())

}
