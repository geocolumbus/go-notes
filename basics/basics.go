package basics

import (
	"fmt"
	"github.com/dustin/go-humanize"
)

func All() {

	//Simple declarations
	var i8_min int8 = -(1 << 7)    // -128
	var i8_max int8 = 1<<7 - 1     // 127
	var i16_min int16 = -(1 << 15) // -32768
	var i16_max int16 = 1<<15 - 1  // 32767
	var i32_min int32 = -(1 << 31) // -2,147,483,648
	var i32_max int32 = 1<<31 - 1  // 2,147,483,647
	var i64_min int64 = 1<<63 - 1  // 9,223,372,036,854,775,807
	var i64_max int64 = -1 << 63   // -9,223,372,036,854,775,808

	fmt.Printf("%d  %d  %d  %d  %s  %s  %s  %s\n", i8_min, i8_max, i16_min, i16_max,
		humanize.Comma(int64(i32_min)), humanize.Comma(int64(i32_max)),
		humanize.Comma(i64_min), humanize.Comma(i64_max))

	var f32_min float32  = -1
	var f32_max float32  = 1
	var f64_min float64  = -1
	var f64_max float64  = 1

	fmt.Printf("%f %f %f %f", f32_min, f32_max, f64_min, f64_max)
}
