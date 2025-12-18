# READ ME FIRST
Loyalty reward system for a coffee shop-alike business. 
Simulates the purchase transactions to reward the customer with points that they can later claim for offers.
customer user - can view and use points, rewards, history, purchases, and profile.
employee user - can view and accept claims, reward offers, transactions history, search profiles, purchases.

## running the program
```
  go run cmd/web/main.go
```

## Folder structure

cmd/ contains application specific 
cmd/web/ contains the web application 

internal/ contains the non application specific code. validations helpers, sql database models.

ui/ contains the use interface assets.
ui/html/ contains html templates.
ui/static contains static files like css and images.


