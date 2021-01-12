/*   file name : main.go
   The Initialize method is responsible for create a database connection and wire up the routes, 
   and the Run method will simply start the application.
*/
package main

func main() {
	shopalyst := Shopalyst{}
	shopalyst.Initialize("root", "qwerty", "shopalyst")
	shopalyst.Run(":8080")
}