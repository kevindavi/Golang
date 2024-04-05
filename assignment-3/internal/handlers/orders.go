package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"assignment-2/internal/models"

	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/orders", createOrder(db)).Methods("POST")
	router.HandleFunc("/orders", getOrders(db)).Methods("GET")
	router.HandleFunc("/orders/{orderId}", updateOrder(db)).Methods("PUT")
	router.HandleFunc("/orders/{orderId}", deleteOrder(db)).Methods("DELETE")

	return router
}

func RegisterOrderHandlers(router *mux.Router, db *sql.DB) {
	router.HandleFunc("/orders", createOrder(db)).Methods("POST")
	router.HandleFunc("/orders", getOrders(db)).Methods("GET")
	router.HandleFunc("/orders/{orderId}", updateOrder(db)).Methods("PUT")
	router.HandleFunc("/orders/{orderId}", deleteOrder(db)).Methods("DELETE")
}

func createOrder(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var order models.Order
		err := json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Insert order into the database
		stmt := `INSERT INTO orders (customer_name, ordered_at) VALUES ($1, $2) RETURNING order_id`
		err = db.QueryRow(stmt, order.CustomerName, order.OrderedAt).Scan(&order.OrderID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Insert items into the database
		for _, item := range order.Items {
			stmt := `INSERT INTO items (item_code, description, quantity, order_id) VALUES ($1, $2, $3, $4)`
			_, err := db.Exec(stmt, item.ItemCode, item.Description, item.Quantity, order.OrderID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(order)
	}
}

func getOrders(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve orders from the database
		rows, err := db.Query(`SELECT order_id, customer_name, ordered_at FROM orders`)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var orders []models.Order
		for rows.Next() {
			var order models.Order
			err := rows.Scan(&order.OrderID, &order.CustomerName, &order.OrderedAt)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			orders = append(orders, order)
		}

		json.NewEncoder(w).Encode(orders)
	}
}

func updateOrder(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		orderID := params["orderId"]

		var order models.Order
		err := json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Update order in the database
		stmt := `UPDATE orders SET customer_name = $1, ordered_at = $2 WHERE order_id = $3`
		_, err = db.Exec(stmt, order.CustomerName, order.OrderedAt, orderID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Update items in the database
		for _, item := range order.Items {
			stmt := `UPDATE items SET item_code = $1, description = $2, quantity = $3 WHERE item_id = $4`
			_, err := db.Exec(stmt, item.ItemCode, item.Description, item.Quantity, item.ItemID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(order)
	}
}

func deleteOrder(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		orderID := params["orderId"]

		// Delete items associated with the order from the database
		stmt := `DELETE FROM items WHERE order_id = $1`
		_, err := db.Exec(stmt, orderID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Delete the order from the database
		stmt = `DELETE FROM orders WHERE order_id = $1`
		_, err = db.Exec(stmt, orderID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
