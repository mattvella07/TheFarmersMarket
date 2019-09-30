package market

import (
	"fmt"
	"io"
	"strings"

	"github.com/mattvella07/farmersMarket/utils"
)

type item struct {
	Code  string
	Name  string
	Price float32
}

type special struct {
	UniqueID int
	Name     string
	Amount   float32
}

type basket struct {
	Items           []item
	Total           float32
	SpecialsApplied []special
}

var (
	chai    = item{Code: "CH1", Name: "Chai", Price: 3.11}
	apples  = item{Code: "AP1", Name: "Apples", Price: 6}
	coffee  = item{Code: "CF1", Name: "Coffee", Price: 11.23}
	milk    = item{Code: "MK1", Name: "Milk", Price: 4.75}
	oatmeal = item{Code: "OM1", Name: "Oatmeal", Price: 3.69}
	items   = []item{
		chai,
		apples,
		coffee,
		milk,
		oatmeal,
	}

	bogo         = special{Name: "BOGO", Amount: coffee.Price}
	appl         = special{Name: "APPL", Amount: apples.Price - 4.5}
	chmk         = special{Name: "CHMK", Amount: milk.Price}
	apom         = special{Name: "APOM", Amount: apples.Price / 2}
	apomWithAppl = special{Name: "APOM", Amount: 4.5 / 2}
	specials     = []special{
		bogo,
		appl,
		chmk,
		apom,
	}

	shoppingBasket basket
)

// AddItem allows the user to add an item to the shopping basket
func AddItem(readFrom io.Reader) {
	fmt.Println("\nEnter the number of the item you would like to add to your basket")

	for {
		// Display all available items
		for idx, i := range items {
			fmt.Printf("%d. %-10s|   %-10s|  $%.2f\n", idx+1, i.Code, i.Name, i.Price)
		}

		// Get user's choice
		choice, err := utils.GetIndexChosen(readFrom)
		if err != nil {
			fmt.Printf("%s\n\n", err)
			continue
		}

		// Validate user's choice
		if !utils.ChoiceValid(choice, len(items)) {
			fmt.Println("Invalid choice, please try again")
			continue
		}

		// Add item to shoppingBasket
		shoppingBasket.Items = append(shoppingBasket.Items, items[choice])
		fmt.Printf("%s has been added to your basket\n\n", items[choice].Name)

		addSpecials()
		calculateTotal()

		return
	}
}

// RemoveItem allows the user to remove an item from the shopping basket
func RemoveItem(readFrom io.Reader) {
	if len(shoppingBasket.Items) == 0 {
		fmt.Println("\nYour basket is currently empty, nothing to remove.")
		return
	}

	fmt.Println("\nEnter the number of the item you would like to remove from your basket")

	for {
		// Display all items currently in the shopping basket
		for idx, i := range shoppingBasket.Items {
			fmt.Printf("%d. %-10s|   %-10s|  $%.2f\n", idx+1, i.Code, i.Name, i.Price)
		}

		// Get user's choice
		choice, err := utils.GetIndexChosen(readFrom)
		if err != nil {
			fmt.Printf("%s\n\n", err)
			continue
		}

		// Validate user's choice
		if !utils.ChoiceValid(choice, len(shoppingBasket.Items)) {
			fmt.Println("Invalid choice, please try again")
			continue
		}

		// Remove item from shoppingBasket
		fmt.Printf("%s has been removed from your basket\n\n", shoppingBasket.Items[choice].Name)
		shoppingBasket.Items = append(shoppingBasket.Items[:choice], shoppingBasket.Items[choice+1:]...)

		addSpecials()
		calculateTotal()

		return
	}
}

// ViewBasket displays the current shopping basket and returns false if
// it is empty, true otherwise
func ViewBasket() bool {
	if len(shoppingBasket.Items) == 0 {
		fmt.Println("\nShopping basket is currently empty")
		return false
	}

	// Create map to store false if the special has been displayed, true otherwise
	displayedSpecials := make(map[special]bool)
	for _, s := range shoppingBasket.SpecialsApplied {
		displayedSpecials[s] = false
	}

	// Display header
	fmt.Println("\nItem                         Price")
	fmt.Println("----                         -----")

	// Display items and their associated specials
	for _, i := range shoppingBasket.Items {
		fmt.Printf("%s                          $%.2f\n", i.Code, i.Price)

		displayApplicableSpecials(i.Code, displayedSpecials)
	}

	// Display total
	fmt.Println("-----------------------------------")
	fmt.Printf("                             $%.2f\n", shoppingBasket.Total)

	return true
}

// Checkout displays the current shopping basket and then empties it
func Checkout() {
	// Only display messages if the basket is not empty
	if ViewBasket() {
		fmt.Println("Thanks for your purchase!")
		shoppingBasket = basket{}
		fmt.Println("Your shopping basket is now empty")
	}
}

func displayApplicableSpecials(itemCode string, displayedSpecials map[special]bool) {
	specialsAlreadyDisplayed := ""
	applicableSpecials := make(map[string]struct{})

	// Sets the applicable specials based on the item code
	switch itemCode {
	case "CF1":
		applicableSpecials["BOGO"] = struct{}{}
	case "AP1":
		applicableSpecials["APPL"] = struct{}{}
		applicableSpecials["APOM"] = struct{}{}
	case "MK1":
		applicableSpecials["CHMK"] = struct{}{}
	}

	// Display applicable specials
	for s, isDisplayed := range displayedSpecials {
		if _, ok := applicableSpecials[s.Name]; ok && !isDisplayed {
			// The special can only be applied once to an individual item
			if !strings.Contains(specialsAlreadyDisplayed, s.Name) {
				fmt.Printf("           %s             -$%.2f\n", s.Name, s.Amount)

				// Since special was displayed, set it to true in the map and
				// add it to the specialsAlreadyDisplayed string
				// displayedSpecials map prevents multiple items from showing the same
				// special, while specialsAlreadyDisplayed string prevents an individual
				// item from displaying the same promotion more than once
				displayedSpecials[s] = true
				specialsAlreadyDisplayed += fmt.Sprintf("%s,", s.Name)
			}
		}
	}
}

func addSpecials() {
	shoppingBasket.SpecialsApplied = []special{}
	numCoffeePurchased := 0
	numApplesPurchased := 0
	chaiPurchased := false
	milkPurchased := false
	oatmealPurchased := false

	// Determine whcih specials apply
	for _, i := range shoppingBasket.Items {
		switch i.Name {
		case "Coffee":
			// 1. BOGO -- Buy-One-Get-One-Free Special on Coffee. (Unlimited)
			// if number of coffee purchased is divisible by 2 then append the
			// BOGO speical to the SpecialsApplied property of shoppingBasket
			numCoffeePurchased++

			if numCoffeePurchased%2 == 0 {
				newSpecial := special{
					// UniqueID helps differentiate speicals of the same name
					UniqueID: len(shoppingBasket.SpecialsApplied),
					Name:     bogo.Name,
					Amount:   bogo.Amount,
				}
				shoppingBasket.SpecialsApplied = append(shoppingBasket.SpecialsApplied, newSpecial)
			}
		case "Chai":
			chaiPurchased = true
		case "Milk":
			milkPurchased = true
		case "Apples":
			numApplesPurchased++
		case "Oatmeal":
			oatmealPurchased = true
		}
	}

	// 2. APPL -- If you buy 3 or more bags of Apples, the price drops to $4.50.
	// If 3 or more apples were purchased, then add the promotion for disocunted
	// apples for every bag of apples purchased
	if numApplesPurchased >= 3 {
		for x := 0; x < numApplesPurchased; x++ {
			newSpecial := special{
				UniqueID: len(shoppingBasket.SpecialsApplied),
				Name:     appl.Name,
				Amount:   appl.Amount,
			}
			shoppingBasket.SpecialsApplied = append(shoppingBasket.SpecialsApplied, newSpecial)
		}
	}

	// 3. CHMK -- Purchase a box of Chai and get milk free. (Limit 1)
	// If both Chai and Milk were purchased, then add the promotion for free milk
	if chaiPurchased && milkPurchased {
		newSpecial := special{
			UniqueID: len(shoppingBasket.SpecialsApplied),
			Name:     chmk.Name,
			Amount:   chmk.Amount,
		}
		shoppingBasket.SpecialsApplied = append(shoppingBasket.SpecialsApplied, newSpecial)
	}

	// 4. APOM -- Purchase a bag of Oatmeal and get 50% off a bag of Apples
	// If both Oatmeal and Apples were purchased, then add the promotion for
	// 50% off a bag of Apples. If less than 3 bags were purchased then it is
	// 50% off the original price, if 3 or more bags were purchased then it
	// is 50% off the discounted apple price of $4.50
	if oatmealPurchased && numApplesPurchased > 0 {
		newSpecial := special{
			UniqueID: len(shoppingBasket.SpecialsApplied),
			Name:     apom.Name,
		}
		if numApplesPurchased < 3 {
			newSpecial.Amount = apom.Amount
			shoppingBasket.SpecialsApplied = append(shoppingBasket.SpecialsApplied, newSpecial)
		} else {
			newSpecial.Amount = apomWithAppl.Amount
			shoppingBasket.SpecialsApplied = append(shoppingBasket.SpecialsApplied, newSpecial)
		}
	}
}

func calculateTotal() {
	shoppingBasket.Total = 0

	// Calculate total by adding the price of all items and then subtracting
	// the amount of each speical
	for _, i := range shoppingBasket.Items {
		shoppingBasket.Total += i.Price
	}

	for _, s := range shoppingBasket.SpecialsApplied {
		shoppingBasket.Total -= s.Amount
	}
}
