package market

import (
	"fmt"
	"strings"
	"testing"
)

func TestAddItem(t *testing.T) {
	t.Run("Add items to shopping basket", func(t *testing.T) {
		{
			expected := 0
			if len(shoppingBasket.Items) != expected {
				t.Fatalf("Expected shoppingBasket.Items to have a length of %d, got %d", expected, len(shoppingBasket.Items))
			}
		}

		AddItem(strings.NewReader("1\n"))

		{
			expected := 1
			if len(shoppingBasket.Items) != expected {
				t.Fatalf("Expected shoppingBasket.Items to have a length of %d, got %d", expected, len(shoppingBasket.Items))
			}
		}

		{
			expected := "CH1"
			if shoppingBasket.Items[0].Code != expected {
				t.Fatalf("Expected item code to equal %s, got %s", expected, shoppingBasket.Items[0].Code)
			}
		}

		{
			expected := "Chai"
			if shoppingBasket.Items[0].Name != expected {
				t.Fatalf("Expected item name to equal %s, got %s", expected, shoppingBasket.Items[0].Name)
			}
		}

		{
			expected := float32(3.11)
			if fmt.Sprintf("%.2f", shoppingBasket.Items[0].Price) != fmt.Sprintf("%.2f", expected) {
				t.Fatalf("Expected item price to equal %.2f, got %.2f", expected, shoppingBasket.Items[0].Price)
			}
		}

		{
			expected := float32(3.11)
			if fmt.Sprintf("%.2f", shoppingBasket.Total) != fmt.Sprintf("%.2f", expected) {
				t.Fatalf("Expected total to equal %.2f, got %.2f", expected, shoppingBasket.Total)
			}
		}

		AddItem(strings.NewReader("3\n"))

		{
			expected := 2
			if len(shoppingBasket.Items) != expected {
				t.Fatalf("Expected shoppingBasket.Items to have a length of %d, got %d", expected, len(shoppingBasket.Items))
			}
		}

		{
			expected := float32(11.23)
			if fmt.Sprintf("%.2f", shoppingBasket.Items[1].Price) != fmt.Sprintf("%.2f", expected) {
				t.Fatalf("Expected item price to equal %.2f, got %.2f", expected, shoppingBasket.Items[1].Price)
			}
		}

		{
			expected := float32(14.34)
			if fmt.Sprintf("%.2f", shoppingBasket.Total) != fmt.Sprintf("%.2f", expected) {
				t.Fatalf("Expected total to equal %.2f, got %.2f", expected, shoppingBasket.Total)
			}
		}
	})
}

func TestRemoveItem(t *testing.T) {
	t.Run("Remove items from shopping basket", func(t *testing.T) {
		{
			expected := 2
			if len(shoppingBasket.Items) != expected {
				t.Fatalf("Expected shoppingBasket.Items to have a length of %d, got %d", expected, len(shoppingBasket.Items))
			}
		}

		RemoveItem(strings.NewReader("1\n"))

		{
			expected := 1
			if len(shoppingBasket.Items) != expected {
				t.Fatalf("Expected shoppingBasket.Items to have a length of %d, got %d", expected, len(shoppingBasket.Items))
			}
		}

		{
			expected := "CF1"
			if shoppingBasket.Items[0].Code != expected {
				t.Fatalf("Expected item code to equal %s, got %s", expected, shoppingBasket.Items[0].Code)
			}
		}

		{
			expected := "Coffee"
			if shoppingBasket.Items[0].Name != expected {
				t.Fatalf("Expected item name to equal %s, got %s", expected, shoppingBasket.Items[0].Name)
			}
		}

		{
			expected := float32(11.23)
			if shoppingBasket.Items[0].Price != expected {
				t.Fatalf("Expected item price to equal %f, got %f", expected, shoppingBasket.Items[0].Price)
			}
		}
	})

	RemoveItem(strings.NewReader("1\n"))

	{
		expected := 0
		if len(shoppingBasket.Items) != expected {
			t.Fatalf("Expected shoppingBasket.Items to have a length of %d, got %d", expected, len(shoppingBasket.Items))
		}
	}
}

func TestViewBasket(t *testing.T) {
	t.Run("View empty shopping basket", func(t *testing.T) {
		itemsInBasket := ViewBasket()

		{
			expected := false
			if itemsInBasket != expected {
				t.Fatalf("Expected ViewBasket to return %v, got %v", expected, itemsInBasket)
			}
		}
	})

	t.Run("View shopping basket with items", func(t *testing.T) {
		AddItem(strings.NewReader("1\n"))
		AddItem(strings.NewReader("3\n"))

		itemsInBasket := ViewBasket()

		{
			expected := true
			if itemsInBasket != expected {
				t.Fatalf("Expected ViewBasket to return %v, got %v", expected, itemsInBasket)
			}
		}
	})
}
