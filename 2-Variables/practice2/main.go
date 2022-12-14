// 1 : int x, float64 y type conversion sample

/* package main
import "fmt"
func main() {
	x := 75
	var y float64
	y = float64(x) // type(value)
	fmt.Println(y)
} */

// 2 : multiple assing sample x, y = y, x

/* package main
import "fmt"
func main() {
	x := 5
	y := 10
	fmt.Println("X:", x, "Y:", y)
	x, y = y, x // x = y , y = x
	fmt.Println("X:", x, "Y:", y)
} */

// 3 : non English variable names
//utf8 uyumlu dil direk farkli alfabelerde degskenler olsablr
/* package main
import "fmt"
func main() {
	 	yaş := 40
	   	fmt.Println(yaş)
	名稱 := "Arin"
	年齡 := 40
	fmt.Println("Name:", 名稱, "Age:", 年齡)
} */

// 4 : shadowing kavramı? gölgeleme

/* package main
import "fmt"
func main() {
	x := 5
	if true {
		//x:=10 buradaki x ustteki x i golgelems oluyor
		x = 10
		x++
		fmt.Println(x)
	}
	fmt.Println(x)
} */

// 5 : 40 as a string

package main

import "fmt"

/*func main() {

	x := 65

	s := string(x)

	fmt.Printf("%v, %T\n", x, x) // 65
	fmt.Printf("%v, %T\n", s, s) // A

	y := strconv.Itoa(x)
	fmt.Printf("%v, %T\n", y, y) // "65"
}*/


func main(){
	x := 40
	var y float64

	fmt.Printf("%T", x)
	fmt.Println()
	fmt.Printf("location ", &x)
	fmt.Println()

	if true {
		x = 20
		fmt.Printf("location ", &x)

		x:=2
		fmt.Printf("location ", &x)


	}

	fmt.Printf("%T", y)
}

// 1 : int x, float64 y type conversion sample
/*

 */