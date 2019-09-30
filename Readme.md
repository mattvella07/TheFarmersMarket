## The Farmer's Market
<hr>

### Description

The Farmer's Market that sells the following items:

```
+--------------|--------------|---------+
| Product Code |     Name     |  Price  |
+--------------|--------------|---------+
|     CH1      |   Chai       |  $3.11  |
|     AP1      |   Apples     |  $6.00  |
|     CF1      |   Coffee     | $11.23  |
|     MK1      |   Milk       |  $4.75  |
|     OM1      |   Oatmeal    |  $3.69  |
+--------------|--------------|---------+
```

The following specials are currently available:
```
1. BOGO -- Buy-One-Get-One-Free Special on Coffee. (Unlimited)
2. APPL -- If you buy 3 or more bags of Apples, the price drops to $4.50.
3. CHMK -- Purchase a box of Chai and get milk free. (Limit 1)
4. APOM -- Purchase a bag of Oatmeal and get 50% off a bag of Apples
```

### Building and Running

- This app runs in a Docker container and can be built and ran using the command `make all`
- That will build the image from the dockerfile in the root directory of this project, and run it in a new container.

### Testing

- All tests are located in the `market` folder and can be run using `go test` from within that folder.

### Using

- This is a CLI program that the user can interact with by selecting the number of the menu item they want.  The prompts and menu choices will guide the user through add items to the basket, removing items from the basket, viewing the current basket, and checking out.
