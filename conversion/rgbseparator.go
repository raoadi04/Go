// rgbseparator.go
// description: split rgb input to red, green and blue and vice versa
// author(s) [darmiel](https://github.com/darmiel)
// see rgbseparator_test.go

package conversion

// ExtractRGB splits an RGB input (e.g. a color in hex format; 0x<color-code>)
// into the individual components: red, green and blue
func ExtractRGB(rgb uint) (red, green, blue byte) {
	// A hex code is structured like this:
	// #3498db (light blue) - converted to binary:
	// 00110100 10011000 11011011
	//  <red>   <green>   <blue>

	// To get the blue value we use the bit operation AND with the bit mask 0xFF (in binary: 11111111)
	// 00110100 10011000 <11011011> &
	// 00000000 00000000  11111111  =
	// 00000000 00000000 <11011011> =
	blue = byte(rgb & 0xFF)

	// To get the green value, we first shift the value 8 bits to the right:
	//  00110100 <10011000>  11011011  >> 8 =
	//  00000000  00110100  <10011000> &
	//  00000000  00000000   11111111  =
	//  00000000  00000000  <10011000> =
	green = byte((rgb >> 8) & 0xFF)

	// Same as green value, only this time shift 16 to the right
	// Alternatively, you can apply a bitmask first and then shift it.
	// <00110100> 10011000  11011011 &
	//  11111111  00000000  00000000 =
	// <00110100> 00000000  00000000 >> 16
	//  00000000  00000000 <00110100> =
	red = byte((rgb >> 16) & 0xFF)
	return
}

// CombineRGB does exactly the opposite of ExtractRGB:
// it combines the three components red, green and blue to an RGB value, which can be converted to e.g. Hex
func CombineRGB(red, green, blue byte) (rgb uint) {
	// Sets the bits of blue in position 1-8, green in 9-16 and red in 17-24

	// Red: 00110100
	// Green: 10011000
	// Blue: 11011011
	// RGB:
	// R << 16: [00110100]  00000000   00000000 |
	// G << 8 :  00000000  {10011000}  00000000 |
	// B      :  00000000   00000000  <11011011> =
	//          [00110100] {10011000} <11011011>
	return (uint(red) << 16) | (uint(green) << 8) | uint(blue)
}
