package main

import "fmt"
/*func swap(x  ,y string) (string,string) {
	return y,x
}*/
type Vertex struct
func main() {

	fmt.Print("Enter number :")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
	
	for number :=1;number <=100;number++ {
		if number%15 ==0{
			fmt.Println(number,"FizzBuzz")
		} else if number%3==0{
			fmt.Println(number,"Fizz")
		} else if number%5==0{
			fmt.Println(number,"Buzz")
		}else{
			fmt.Println(number)
		}
	}
     i :=5 */
	

   
   /*switch i :=5;i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	default:
	    fmt.Println("AAAA")

	}
	result :=true
	if result ==1{
		fmt.Println("AAAA")
	}*/

	
		//fmt.Println(add(42,13))
	//	a,b :=swap("hello","World")
	//	fmt.Println(a,b)
	/*
	i,j :=42,2701
	p :=&i
	fmt.Println(*p)
	*p =21
	fmt.Println(i)
	p =&j
	*p =*p/37
	fmt.Println(j)
	*/

	
}
