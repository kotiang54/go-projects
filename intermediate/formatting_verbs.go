package intermediate

import "fmt"

func main() {
	/*
		--- General formating verbs ---
		%v  Prints the value in the default format
		%#v Prints the value in Go-syntax format
		%T  Prints the type of the value
		%%  Prints the % sign
	*/
	i := 11_111_555 // float64 15.5

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%v%%\n", i)

	// Strings
	str := "Hello World!"
	fmt.Printf("%v\n", str)
	fmt.Printf("%#v\n", str)
	fmt.Printf("%T\n\n", str)

	/*
		--- Integer formating ---
		%b   base 2
		%d   base 10
		%+d  base 10 and always show sign
		%o   base 8
		%O   base 8, with leading 0o
		%x   base16, lowercase
		%X   base 16, uppercase
		%#X  base 16, with leading 0x
		%4d  Pad with spaces (width 4, right justified)
		%-4d Pad with spaces (width 4, left justified)
		%04d Pad with zeros (width 4)
	*/
	num := 255
	fmt.Printf("%b\n", num)
	fmt.Printf("%d\n", num)
	fmt.Printf("%+d\n", num)
	fmt.Printf("%o\n", num)
	fmt.Printf("%O\n", num)
	fmt.Printf("%x\n", num)
	fmt.Printf("%X\n", num)
	fmt.Printf("%#X\n", num)
	fmt.Printf("%4d\n", num)
	fmt.Printf("%-4d\n", num)
	fmt.Printf("%04d\n\n", num)

	/*
		--- String formating ---
		%s   Prints the value as a plain string
		%q   Prints the value as a double-quoted string
		%8s  Prints the value as plain string (width 8, right justified)
		%-8s Prints the value as plain string (width 8, left justified)
		%x   Prints the value as hex dump of byte values
		% x  Prints the value as hex dump with spaces
	*/
	txt := "Hello World"
	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%16s\n", txt)
	fmt.Printf("%-16s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("% x\n\n", txt)

	// Boolean formating
	// %t Value of the Boolean operator in true of false format (same as usin %v)
	t := true
	f := false

	fmt.Printf("%t\n", t)
	fmt.Printf("%t\n\n", f)

	/*
		--- Float formating ---
		%e    Scientific notation with 'e' as exponent
		%f    Decimal point, no exponent
		%.2f  Default width, precision 2
		%6.2f Width 6, precision 2
		%g    Exponent as needed, only necessary digits
	*/
	flt := 900000000.182
	fmt.Printf("%e\n", flt)
	fmt.Printf("%f\n", flt)
	fmt.Printf("%.2f\n", flt)
	fmt.Printf("%6.2f\n", flt)
	fmt.Printf("%g\n", flt)
}
