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
/*
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
*/


// Programmer name: Jaime Ann
// Date completed:  02/25/2020
// Description: 2.3.2 Lab- If Part 2

//Paystub- 3 points
//Have a user enter their hourly wages, how many hours they work, their name.  Determine their wages including overtime and deduct their taxes (using a flat tax rate of 12%).  Print a paystub.
/*
import 
 "fmt"


func main() {
//declare variables for hourly wage, hourly wage during overtime, hours worked, name, total wage earned, taxes, and total after tax deduction

var userName string
var wage float64 
var wageOvertime float64
var hoursWorked float64
var hoursOvertime float64
var hoursWorkedMinusOt float64
var overtimeEarnings float64
var totalEarned float64
var taxTotal float64
var totalAwarded float64

//print statement requesting user's name
fmt.Println("Hi there! Before we get started, please confirm your name.")
//scan in userName
fmt.Scanln(&userName)
//print statement letting the user know what we're attempting to accomplish
fmt.Println("Great, ", userName,"! Now we'll move into calculating your projected total earnings and take home income for this week's paystub.")
//print statement requesting user's hourly wage
fmt.Println("What is your current hourly wage?")
//scan in wage
fmt.Scanln(&wage)
//create equation to calculate wageOvertime 
wageOvertime = wage * 1.5 
//print statement requesting how many hours the user has worked 
fmt.Println("How many hours did you work this week?")
//scan in hoursWorked
fmt.Scanln(&hoursWorked)
//create if statement for hoursWorked being greater than 40, hoursWorked - 40 = hoursOvertime, hoursOvertime * wageOvertime = overtimeEarnings 
if hoursWorked > 40 {
  hoursOvertime= hoursWorked - 40
  hoursWorkedMinusOt= hoursWorked-hoursOvertime
  overtimeEarnings= hoursOvertime*wageOvertime
  totalEarned= (hoursWorked*wage)+overtimeEarnings
//create else statement for hoursWorked<= 40, hoursWorked * wage = totalEarned
}else{ 
  totalEarned= hoursWorked* wage
  hoursOvertime= 0
  wageOvertime= 0
  overtimeEarnings= 0
  hoursWorkedMinusOt = hoursWorked
}
//print statement telling user their total earnings before taxes
fmt.Println("Your total earnings before tax deductions are",totalEarned,".")
//create equation to find taxTotal by multiplying totalEarned by .12
taxTotal=totalEarned*.12
//subtract taxTotal from totalEarned to find totalAwarded
totalAwarded=totalEarned-taxTotal
//print statement telling user their earnings after deductions
fmt.Println("Your earnings after tax deductions are", totalAwarded,".")
//print statement including user weekly income report
fmt.Println("Below I have printed for you your weekly paystub for the current week.")
fmt.Println("Name:",userName)
fmt.Println("Hourly Wage:",wage)
fmt.Println("Hourly Overtime Wage:",wageOvertime)
fmt.Println("Hours Worked:",hoursWorkedMinusOt)
fmt.Println("Overtime Hours Worked:",hoursOvertime)
fmt.Println("Overtime Earnings:",overtimeEarnings)
fmt.Println("Total Earning Prior to Deductions:",totalEarned)
fmt.Println("Taxes Deducted:", taxTotal)
fmt.Println("Total Wages Awarded:",totalAwarded)
}
*/

/*
// Programmer name: Jaime Ann
// Date completed:  03/03/2020
// Description: Lab 2.5.1. - Section 1 - You Code It

import "fmt"

func main() {
    //Create a string array of at least 5 values.  It should hold 5 products to buy
    //Create a float array of at least 5 values.  It should hold a price for each of the products
    //Iterate through the array and output the product and it's price (similar to a menu)
}
*/

// Programmer name: Jaime Ann
// Date completed:  03/10/2020
// Description: Project #2 - "The Game of Nim" Suggestion

//basics and objective of game: two players remove stones from three differently valued piles; the ultimate goal is to have all three piles entirely diminished. The player who removes the final stone wins. Players can only remove stones from one pile per turn. There are no limits on how many stones can be taken from the piles, though. 

import "fmt"

func main() {

  //create/declare the three "piles of stones" - variables that store an integer representative of the number of stones in each piles. Piles cannot be < 0. 

  var pileA int
  var pileB int
  var pileC int
  var pileSum int
  var pileSelection string
  var stonesRemoved int 

  //set values for each pile of stones

  pileA= 13
  pileB= 45
  pileC=7
  pileSum = pileA + pileB + pileC

  //explain the game and rules to the user briefly
fmt.Println("Let's play the Game of Nim! The object of the game is simple: be the first to remove all of the stones from three piles of stones and you win! In each turn, you will be allowed to remove as many stones as you like from one pile. You may not remove more stones than are in a pile nor can you remove from more than one pile in each turn. If that all sounds alright, let's get started!")

//declare how many stones are in each pile
fmt.Println("To start, pile A has", pileA,"stones in it. pile B has", pileB,"stones in it, and pile C has", pileC, "stones.")

  //create an if/until loop that (if- all piles >0) asks the user which pile they want to remove stones from and how many stones they want to remove , and then computer removes the entire amount of stones left in a pile, (until-all piles =0) then, the game declares the winner as whomever removed the last stone to make the piles = 0. 

  //start loop of game
  for pileSum > 0 {
  //ask user which pile they want to remove stones from, A, B, or C
 fmt.Println("Would you like to remove stones from pile A, B, or C? Please enter A, B, or C below to select a pile.")
 fmt.Scanln(&pileSelection)
  //If pileSelection = "A" 
  if pileSelection == "A" {
  //If pileA=0, print ("Pile A is empty. Please select another pile") else ask user how many stones they want to remove
  if pileA==0 {
   fmt.Println("Pile A is empty. Please select another pile")
  
  }else{
    fmt.Println("How many stones would you like to remove from Pile A?")
    fmt.Scanln(&stonesRemoved)
  }
  //if stonesRemoved > pileA, print ("There are not that many stones in the pile. Please remove fewer stones") else stonesRemoved<= pileA, pileA= pileA-stonesRemoved print("Great! There are now",pileA, "stones left in Pile A!")

  if stonesRemoved > pileA {
    fmt.Println("There are not that many stones in the pile. Please remove fewer stones.")
  
  }else{
  pileA= pileA - stonesRemoved
  fmt.Println("Great! There are now", pileA,"stones left in Pile A!")
  }
  //If pileSelection = "B" 
   } else if pileSelection == "B"{
  //If pileB=0, print ("Pile B is empty. Please select another pile") else ask user how many stones they want to remove
  if pileB==0 {
    fmt.Println("Pile B is empty. Please select another pile")
   
   }else{
    fmt.Println("How many stones would you like to remove from Pile B?")
    fmt.Scanln(&stonesRemoved)
  }
  //if stonesRemoved > pileB, print ("There are not that many stones in the pile. Please remove fewer stones") else  pileB= pileB-stonesRemoved print("Great! There are now",pileB, "stones left in Pile B!")
    if stonesRemoved > pileB {
    fmt.Println("There are not that many stones in the pile. Please remove fewer stones.")
    
    }else{
  pileB= pileB - stonesRemoved
  fmt.Println("Great! There are now", pileB,"stones left in Pile B!")
  }

  //If pileSelection = "C" 
  }else if pileSelection == "C"{
  //If pileC=0, print ("Pile C is empty. Please select another pile") else ask user how many stones they want to remove and scan input
  if pileC== 0 {
   fmt.Println("Pile C is empty. Please select another pile")
   
   } else{
    fmt.Println("How many stones would you like to remove from Pile C?")
    fmt.Scanln(&stonesRemoved)
  }
  //if stonesRemoved > pileC, print ("There are not that many stones in the pile. Please remove fewer stones")else pileC= pileC-stonesRemoved  print("Great! There are now",pileC, "stones left in Pile C!")
  if stonesRemoved > pileC {
    fmt.Println("There are not that many stones in the pile. Please remove fewer stones.")
  
   } else {
  pileC= pileC - stonesRemoved
  fmt.Println("Great! There are now", pileC,"stones left in Pile C!")
  }

  //if pileSelection not = A, B, or C, print "That isn't one of our piles. Please select pile A, B, or C."
  } else {
    fmt.Println("That isn't one of our piles. Please select pile A, B, or C.")
    }
  }
  for pileSum > 0 {
  /*
  //if pileA >0, computer subtracts pileA from pileA 
  if pileA >0 { 
    pileA= pileA - 1
  //if pileA =0, computer subtracts pileB from pileB
  } else if pileA=0
  pileB= pileB - 1{
  //if pileB >0, computer subtracts pileB from pileB
  } else if pileB >0
  pileB= pileB - 1{
  //if pileB =0, computer subtracts pileC from pileC
  } else if pileB=0 
  pileC= pileC - 1 {
  //if pileC >0, computer subtracts pileC from pileC.
   } else if pileC >0
  pileC= pileC - 1{
  //if pileC =0, computer subtracts pileA from pileA
  }else if pileA >0  
    pileA= pileA - 1{
  //if pileSum =0, print "Congratulations! 
  You've removed all of the stones from the piles. You win!"
    }else if pileSum =0 
    fmt.Println("Congratulations! You've removed all of the stones from the piles. You win!")
  
  }*/
}