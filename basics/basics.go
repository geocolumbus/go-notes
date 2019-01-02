package basics

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"math"
	"math/cmplx"
	"unicode/utf8"
	"unsafe"
)

func All() {
	// integersAndFloatsAndComplex()
	// stringsAndRunes()
	loopsAndIf()
}

func integersAndFloatsAndComplex() {

	fmt.Println("\n----- Number types ----------------")
	fmt.Println("")

	var i = math.MaxInt64
	var i8_min int8 = -1 << 7     // -128
	var i8_max int8 = 1<<7 - 1    // 127
	var i16_min int16 = -1 << 15  // -32768
	var i16_max int16 = 1<<15 - 1 // 32767
	var i32_min int32 = -1 << 31  // -2,147,483,648
	var i32_max int32 = 1<<31 - 1 // 2,147,483,647
	var i64_min int64 = 1<<63 - 1 // 9,223,372,036,854,775,803
	var i64_max int64 = -1 << 63  // -9,223,372,036,854,775,808

	fmt.Printf("%d  %d  %d  %d  %d  %s  %s  %s  %s\n",
		i,
		i8_min, i8_max,
		i16_min, i16_max,
		humanize.Comma(int64(i32_min)), humanize.Comma(int64(i32_max)),
		humanize.Comma(i64_min), humanize.Comma(i64_max))

	var ui8_min uint8 = 0
	var ui8_max uint8 = 1<<8 - 1
	var ui16_min uint16 = 0
	var ui16_max uint16 = 1<<16 - 1
	var ui32_min uint32 = 0
	var ui32_max uint32 = 1<<32 - 1
	var ui64_min uint64 = 0
	var ui64_max uint64 = 1<<64 - 1

	fmt.Printf("%d  %d  %d  %d  %d  %d  %d  %d\n",
		ui8_min, ui8_max,
		ui16_min, ui16_max,
		ui32_min, ui32_max,
		ui64_min, ui64_max)

	var ptr uintptr = uintptr(unsafe.Pointer(new(string)))

	fmt.Printf("%v\n", ptr)

	var f = math.MaxFloat64
	var f32_max float32 = math.MaxFloat32
	var f64_max float64 = math.MaxFloat64

	fmt.Printf("%.4e  %.4e  %.4e\n", f, f32_max, f64_max)

	c := complex(3, 4)
	var c64 complex64 = complex(math.MaxFloat32, math.MaxFloat32)
	var c128 complex128 = complex(math.MaxFloat64, math.MaxFloat64)

	fmt.Printf("%v  %v  %v  %v  %v  %v\n", c, cmplx.Abs(c), c64, cmplx.Abs(complex128(c64)), c128, cmplx.Abs(c128))
}

func stringsAndRunes() {
	fmt.Println("\n----- Strings, runes & bytes ------")
	fmt.Println("")

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
	for ; k < 5; {
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
	switch (animal) {
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

	fmt.Println()

}
