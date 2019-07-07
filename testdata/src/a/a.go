package a

import "fmt"

func main() {
	var err error
	o := MyObj{}

	// Okay
	fmt.Printf("%d %s", 100, "abc")
	fmt.Println("abc", 1, 2, 3)
	myPrint("aa int %d, error %+v", 10, err)
	o.Log("unexpected error. id = %d, err = %+v", 999, err)
	myPrint2(999, "unexpected error. id = %d, err = %+v", 1, err)
	format := "%+v %s"
	fmt.Printf(format, err, "test")

	// Not okay
	fmt.Printf("%v", err)                                           // want `should use %\+v format for error type`
	myPrint("int %d, error %#v", 10, err)                           // want `should use %\+v format for error type`
	o.Log("unexpected error occurred. id = %d, err = %s", 999, err) // want `should use %\+v format for error type`
	invalidFormat := "%s %s"
	fmt.Printf(invalidFormat, err, "test")

	// Handling %%
	fmt.Printf("%% %+v", err)

	// Invalid directive or args.
	// These are not valid for Printf but
	// this checker will address only matching between`%+v` and error arg.
	fmt.Printf("%+v", err, 1)
	fmt.Printf("%+v %d", err)
	fmt.Printf("%", err)         // want `should use %\+v format for error type`
	fmt.Printf("abc % def", err) // want `should use %\+v format for error type`
	fmt.Printf("%d, %", 1, err)  // want `should use %\+v format for error type`
	fmt.Printf("%, %d", err, 1)  // want `should use %\+v format for error type`
}

func myPrint(s string, i int, e error) {
	fmt.Printf(s, i, e)
}

type MyObj struct{}

func (o *MyObj) Log(format string, args ...interface{}) {
	fmt.Printf(format, args)
}

func myPrint2(i int, format string, args ...interface{}) {
	fmt.Printf(format, args)
}
