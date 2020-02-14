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
// Description: 1.4.1 Lab - Input Part 1- Section 1


/*import "fmt"

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
} */

// Programmer name: Jaime Ann
// Date completed:  02/09/2020
// Description: 1.4.1 Lab - Input Part 1 - Section 2 (Pseudocode)
//MPG calculator (3 points)
//Directions: Have the program prompt the user for the size of the gas tank and how many miles traveled.  Determine what the average Miles Per Gallon were on that trip.

//Greet the user.  Mention that you'd be more than happy to calculate their vehicle's MPG for them. 
//Print "For starters how many gallons of gas can your vehicle's gas tank hold?"
//allow them to input an integer for the variable, maxGallons
//Print "Okay,great! How many miles are you able to travel after filling your tank before it gets back down to empty again?"
//Allow them to input an integer for the variable,milesToEmpty
//Divide the second variable by the first variable to get the MPG
// Print "Based on my calculations, your vehicle gets about ___ miles to the gallon!", printing the quotient

//Calculator (3 points)
//Have the program prompt the user for two numbers.  Output each of the math operator equations using those numbers and the answer.  For example you would output "7 + 3 = 10".

//Greet the user. 
//Print "I'd love to help you do some basic addition math problems! To start, what is the first of two numbers you'd like to add together?
//Allow the user to input an integer for the variable, addOne
//Print "Perfect! And what for your second number?"
//Allow the user to input another integer for the second variable, addTwo
//Print "Great, thank you! Those numbers added together are ", (addOne) + (addTwo) = (the sum) 

// Programmer name: Jaime Ann
// Date completed:  02/13/2020
// Description: 1.4.2 Lab - Input Part 2 -(Coding)

//MPG calculator (3 points)
//Directions: Have the program prompt the user for the size of the gas tank and how many miles traveled.  Determine what the average Miles Per Gallon were on that trip.
/*import "fmt"

func main() {

var maxGallons int
var milesToEmpty int
var MPG int 

//Greet the user.  Mention that you'd be more than happy to calculate their vehicle's MPG for them. 

fmt.Println("I'd be more than happy to calculate your vehicle's average MPG for you if you could give me just a bit of information about it!")

//Print "For starters how many gallons of gas can your vehicle's gas tank hold?"

fmt.Println("For starters, how many gallons of gas can your gas tank hold?")

//allow them to input an integer for the variable, maxGallons
fmt.Scanln(&maxGallons)

//Print "Okay,great! How many miles are you able to travel after filling your tank before it gets back down to empty again?"

fmt.Println("Okay, great! How many miles are you able to travel after filling your tank before it gets back down to empty again?")

//Allow them to input an integer for the variable,milesToEmpty

fmt.Scanln(&milesToEmpty)

//Divide the second variable by the first variable to get the MPG

MPG = milesToEmpty / maxGallons

// Print "Based on my calculations, your vehicle gets about ___ miles to the gallon!", printing the quotient

fmt.Println("Based on my calculations, your vehicle gets about",(MPG), "miles to the gallon!")

//Not able to get appropriate quotient above. Not sure what error is occuring in code. 
}
*/

//Calculator (3 points)
//Have the program prompt the user for two numbers.  Output each of the math operator equations using those numbers and the answer.  For example you would output "7 + 3 = 10".
/*import "fmt"

func main () {
//Greet the user. 
//Print "I'd love to help you do some basic addition math problems! To start, what is the first of two numbers you'd like to add together?
var addOne int
var addTwo int 
var sum int

fmt.Println("I'd love to help you do some basic addition math problems! To start, what is the first of two numbers you'd like to add together?")
//Allow the user to input an integer for the variable, addOne
fmt.Scanln(&addOne)
//Print "Perfect! And what for your second number?"
fmt.Println("Perfect! And what for your second number?")
//Allow the user to input another integer for the second variable, addTwo
fmt.Scanln(&addTwo)
//Print "Great, thank you! Those numbers added together are ", (addOne) + (addTwo) = (the sum) 
sum = addOne + addTwo
fmt.Println("Great, thank you! Those numbers added together are ", (addOne), " + ", (addTwo), " = ", (sum), "!" )
}
*/


// Programmer name: Jaime Ann
// Date completed:  02/13/2020
// Description: Project #1 (SLO A, B, & D)

//6. Create a program that will determine how much a person pays for their yogurt at the new yogurt shop based on size of yogurt (in oz), toppings, coupon, sales tax and tip. The new yogurt shop charges $.17 per oz of yogurt and the cup is ¼ of an oz.  Each topping costs $.50.  If the customer has a coupon they can take off between 10% and 50% of their costs.  Sales tax in Chico is 7.25% and a person can choose their own tip.  Output a receipt with ounces of yogurt without the cup, toppings costs, sales tax, discounts, and tip as well as the total fee to be paid.

//Ultimate goal: give customer a cost for their yogurt based on ounces of yogurt, number of toppings, coupon price reduction, sales tax price additions, and an optional tip

import 
 "fmt"


func main() {

var totalOz float64
var yogurtOz float64
var yogurtCost float64
var toppingQty float64
var toppingCost float64
var preCouponTotal float64
var couponPercentOff float64
var couponSavings float64
var preTaxTotal float64
var caSalesTax float64
var postTaxTotal float64
var tip float64
var postTipTotal float64
var customerSignedName string


//how many ounces of yogurt (cost at .17 per oz), print an interaction and then scan in the input

fmt.Println("Please place the frozen yogurt for the current order on the scale and input the number displayed on the scale's screen here.")

fmt.Scanln(&totalOz)

//subtract weight of container from overall weight, est .25 oz weight

yogurtOz = totalOz - .25

yogurtCost = yogurtOz * .17

//how many toppings (cost at .50 each), print interaction and scan input

fmt.Println("How many toppings were added to the customer's tasty treat today?")

fmt.Scanln(&toppingQty)

toppingCost = toppingQty * .50

//calculate pre-tax and pre-coupon total

preCouponTotal = yogurtCost + toppingCost

//coupon price reductions (10-50% reduction), print interaction and scan input

fmt.Println("So far, so good! If the customer brought a coupon in today, what percentage off is the coupon for? Please input this percentage as a decimal. *Note that your store currently has coupons out only for between 10% and 50% off the total price.")

fmt.Scanln(&couponPercentOff)

couponSavings = preCouponTotal * couponPercentOff

preTaxTotal = preCouponTotal - couponSavings


//sales tax based on post coupon total (multiply total by .0725 for 7.25% tax. Add product to post coupon total)

caSalesTax = preTaxTotal * .0725
postTaxTotal = preTaxTotal + caSalesTax

//Print subtotal for employee and advise screen be turned for customer interaction portion of purchase process

fmt.Println("The customer's subtotal after tax is ", (postTaxTotal), ". Please turn this screen toward the customer for optional tip and receipt options") 

//Ask customer if they would like to tip. Add tip to post-tax total

fmt.Println("Your subtotal is ", (postTaxTotal) , ". If you would you like to leave a tip for our froyo experts today, please input a tip amount below.")

fmt.Scanln(&tip)

postTipTotal = postTaxTotal + tip

//Print the grand total for the customer and request their electronic signature for purchase approval 

fmt.Println("Thank you for coming in to Cold Hard Cache FroYo today! Your grand total is ", (postTipTotal), "! Please enter your name below to approve this purchase.")

fmt.Scanln(&customerSignedName)

//Print receipt with ounces of yogurt without the cup, toppings costs, sales tax, discounts, and tip as well as the total fee to be paid.

fmt.Println("Enjoy! Your receipt is shown below.")
fmt.Println("Ounces of FroYo:", (yogurtOz))
fmt.Println("FroYo Subtotal: $",(yogurtCost))
fmt.Println("Topping Subtotal: $",(toppingCost))
fmt.Println("CA Sales Tax: $",(caSalesTax))
fmt.Println("Coupon Savings: $",(couponSavings))
fmt.Println("Tip Subtotal: $",(tip))
fmt.Println("Grand Total: $",(postTipTotal))
fmt.Println("Customer Signature:",(customerSignedName))

}