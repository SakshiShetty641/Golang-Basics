# Golang-Basics

-> import "fmt" imports the fmt package, which provides I/O functions.

-> In statically typed languages, the type of a variable is known at compile time.
-> In dynamically typed languages, the type of a variable is known at runtime. 

->  The go.mod file will be used to track dependencies for your project.
 When you import packages from other modules and run your code or use commands like go build or go run, Go will automatically add these dependencies to the go.mod file and create a go.sum file that ensures the versions of dependencies are consistent.

 ->  Go deduces automatically the type of the variable by looking at the initial value (bool, int, string etc)

 -> Go is a strong typed language (cannot assign float to int)

 -> In Go, iota is a predeclared identifier that simplifies the creation of enumerated constants. 
 It's used within a constant declaration block to automatically generate sequential values.
 iota is reset to 0 whenever the const keyword appears in a block.
It increments by 1 after each line in the block.
-There are ONLY boolean constants, rune constants, integer constants, floating-point constants, complex constants, and string constants.

-> rune: alias for int32, represents a Unicode code point.

->fmt.Sprintf is a function from the fmt package that formats strings similarly to fmt.Printf, but instead of printing the formatted string to standard output, it returns the formatted string itself as a string value. 

-> In Go, an alias is a way to create an alternative name for an existing type. 

-> In Go, command-line arguments are accessible via the os package, specifically through the os.Args variable. This variable is a slice of strings that contains all the arguments passed to the program from the command line.


