// Programmer name: Jaime Ann
// Date completed:  01/30/2020
// Description: 1.2.1 Lab - Section 1 

package main

//import "fmt" 
/*
func main() {
	fmt.Println("Hello world!")
	fmt.Println("My name is Jaime Ann, and my major is Computer Science.")
	fmt.Println("This is my sixth semester at Butte College.")
	fmt.Println("One of my favorite pastimes is hanging out with my wife and getting outdoors.")
}*/




// Programmer name: Jaime Ann
// Date completed:  02/04/2020
// Description: 1.2.2 Lab - Section 1
//1.2.2 Section 1 - Output Part 2 : (Coding Activity)


//ASCII Art - 1 point
/*
package main

import "fmt"

func main() {
  fmt.Println("            W")
fmt.Println("          ~*(._.)*~")
fmt.Println("           J(   )P")  
fmt.Println("            MMMMM  ^")
fmt.Println("            |   |_ |   ")                
fmt.Println("            | ")
fmt.Println("            V  ")  
} 
*/


//"Today" - 2 points: 
//Directions: Create two string variables that hold the current date of the week and month.  Create two integer variables that hold the current date and year.   Output an appropriate message of today's date (appropriate includes verbiage to explain what we are seeing).

// Want to say: Today is Thursday, January 30th in the year 2020
//"Today is (day of the week in string data type), (current month in string data type) (current day of the month in integer data type)th in the year (current year in integer data type)"

//Want to say: Today is the 30th of January, which is a Thursday in the year 2020. 
//"Today is the (current day of the month in integer data type)th of (current month in string data type), which is a (current day of the week in string data type) in the year (current year in integer form).""

//Want to say: 01 30 2020
//"(Integer value of current month with 0 preceding integers <10) (integer value of current date with 0 preceding integers <10) (Integer value of current year)

//Want to say: I want to do this another way but havent figured out how without commas or slash marks.
// 30 01 2020 ??


//"Temperature" - 2  points
//Directions: Create a string variable that holds your home town name and a float variable that holds a temperature.  Output the "current" temperature.  Update your temperature with tomorrow's high temperature and output tomorrow's temperature.  Make sure you use appropriate verbiage to explain what we are seeing.

//Want to say: Oroville, California  64.0 degrees in Oroville, California  currently
// "(current temperature in float data type) in (City name in string data type) , (State name in string data type)" 

// Programmer name: Jaime Ann
// Date completed:  02/06/2020
// Description: 1.4.1 Lab - Input Part 1


import "fmt"

func main() {
    //declare variable for favorite food and store your favorite food.
    //var favFood string
    var favFood= "flan"

    //declare variables for name and age (make sure they are appropriate data types)

    var nameOf string
    var ageOf int
   
    //ask the user to enter their answer for name and age.
    fmt.Println("Hi there! What's your name?")
    fmt.Scanln(&nameOf)
    fmt.Println("Awesome,",(nameOf),"And how old are you?")
    fmt.Scanln(&ageOf)

    //output a welcome statement using their name
    fmt.Println ("We're so glad to have you here", (nameOf))
    
    //output a statement that says At their age you enjoyed the favorite food
    fmt.Println("When I was", (ageOf),", my favorite food was ",(favFood))
}