package cli

import (
	"fmt"
	"io"
	"os"

	"github.com/mattvella07/farmersMarket/market"
	"github.com/mattvella07/farmersMarket/utils"
)

type action struct {
	Name        string
	Description string
}

var (
	actions = []action{
		action{Name: "addItem", Description: "Add item to basket"},
		action{Name: "removeItem", Description: "Remove item from basket"},
		action{Name: "viewBasket", Description: "View basket"},
		action{Name: "checkout", Description: "Checkout"},
		action{Name: "quit", Description: "Leave"},
	}
)

// Start starts the CLI
func Start() {
	var selectedAction string
	fmt.Println("\nWelcome to the Farmer's Market")

	for selectedAction != "quit" {
		selectedAction = displayMainMenu(os.Stdin)

		switch selectedAction {
		case "addItem":
			market.AddItem(os.Stdin)
		case "removeItem":
			market.RemoveItem(os.Stdin)
		case "viewBasket":
			market.ViewBasket()
		case "checkout":
			market.Checkout()
		case "quit":
			fmt.Println("\nThanks for shopping with us today!\n")
		}
	}
}

func displayMainMenu(readFrom io.Reader) string {
	fmt.Println("Enter the number of the action you would like to perform")

	for {
		// Display main menu
		for idx, a := range actions {
			fmt.Printf("%d. %s\n", idx+1, a.Description)
		}

		// Get user's choice
		choice, err := utils.GetIndexChosen(readFrom)
		if err != nil {
			fmt.Printf("%s\n\n", err)
			continue
		}

		// Validate user's choice
		if !utils.ChoiceValid(choice, len(actions)) {
			fmt.Println("Invalid choice, please try again")
			continue
		}

		// Return action to perform
		return actions[choice].Name
	}
}
