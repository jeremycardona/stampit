package main

import (
	"net/http"

	"github.com/jeremycardona/stampit/code/ui/views"
)

func (app *App) home(w http.ResponseWriter, r *http.Request) {
	views.Home().Render(r.Context(), w)
}

func (app *App) claimPost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(201)
	w.Write([]byte("Claim rewards for the customer" + r.URL.Path + "\n"))
}

func (app *App) redeemPost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("enter code or scan qr code for a redemption " + r.URL.Path + "\n"))
}

func (app *App) pointsPost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Claim rewards for the customer" + r.URL.Path + "\n"))
}

func (app *App) points(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("View points and progress towards rewards" + r.URL.Path + "\n"))
}

func (app *App) rewards(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("View rewards for the customer" + r.URL.Path + "\n"))
}

func (app *App) rewardsPost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Refund your rewards" + r.URL.Path + "\n"))
}

func (app *App) customerHistory(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("customer history for all the transactions" + r.URL.Path + "\n"))
}

func (app *App) employeeHistory(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("employee history for all the transactions" + r.URL.Path + "\n"))
}

func (app *App) offers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("list of offers created for customer to get rewards" + r.URL.Path + "\n"))
}

func (app *App) offerPost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("create offers" + r.URL.Path + "\n"))
}

func (app *App) offerPut(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("modify offer / edit" + r.URL.Path + "\n"))
}

func (app *App) offerDelete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete a offer / remove" + r.URL.Path + "\n"))
}

func (app *App) search(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("view a profile by search" + r.URL.Path + "\n"))
}

func (app *App) searchPost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("search for an email" + r.URL.Path + "\n"))
}

func (app *App) purchase(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("available store items" + r.URL.Path + "\n"))
}

func (app *App) purchasePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pay for items in cart" + r.URL.Path + "\n"))
}

func (app *App) loginPost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("log in account" + r.URL.Path + "\n"))
}

func (app *App) signupPost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("sign up for an account" + r.URL.Path + "\n"))
}

func (app *App) memberPost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("enter member email to start purchase and automatically give rewards" + r.URL.Path + "\n"))
}
