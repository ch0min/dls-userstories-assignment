package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	_ "github.com/lib/pq"
)

type Shipment struct {
	ID       int    `json:"id"`
	Chemical string `json:"chemical"`
	Amount   int    `json:"amount"`
	Accepted bool   `json:"accepted"`
}

func main() {
	connStr := "host=localhost port=5432 user=postgres password=test dbname=chemical sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app.Post("/api/shipment", acceptShipment(db))
	app.Get("/api/audit", getAudit(db))

	log.Fatal(app.Listen(":4000"))
}

//

func acceptShipment(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		shipment := new(Shipment)
		if err := c.BodyParser(shipment); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
		}

		// Get warehouse capacity
		var warehouseCapacity int
		err := db.QueryRow("SELECT capacity FROM warehouse WHERE id = 1").Scan(&warehouseCapacity)
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString("Unable to fetch warehouse capacity")
		}

		// Check if warehouse has enough capacity
		var accepted = shipment.Amount <= warehouseCapacity

		// Insert the shipment details into the shipment_details table (which acts as a receipt for where accepted=true)
		var lastInsertID int
		err = db.QueryRow("INSERT INTO shipment_details (chemical, amount, accepted) VALUES ($1, $2, $3) RETURNING id", shipment.Chemical, shipment.Amount, accepted).Scan(&lastInsertID)
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString("Failed to accept shipment")
		}
		if !accepted {
			return c.Status(fiber.StatusBadRequest).SendString("Shipment rejected: not enough capacity")
		}

		// Update the warehouse capacity
		newCapacity := warehouseCapacity - shipment.Amount
		_, err = db.Exec("UPDATE warehouse SET capacity = $1 WHERE id = 1", newCapacity)
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString("Unable to update warehouse capacity")
		}

		// Prepare the response with shipment details
		shipment.ID = lastInsertID
		shipment.Accepted = true
		return c.Status(http.StatusCreated).JSON(shipment)
	}
}

func getAudit(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var shipments []Shipment
		rows, err := db.Query("SELECT id, chemical, amount, accepted FROM shipment_details WHERE accepted = true")
		if err != nil {
			log.Println("Error querying database:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal Server Error",
			})
		}
		defer rows.Close()

		for rows.Next() {
			var shipment Shipment
			if err := rows.Scan(&shipment.ID, &shipment.Chemical, &shipment.Amount, &shipment.Accepted); err != nil {
				log.Println("Error scanning row:", err)
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Internal Server Error",
				})
			}
			shipments = append(shipments, shipment)
		}

		if err := rows.Err(); err != nil {
			log.Println("Error iterating over rows:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal Server Error",
			})
		}

		// Return the shipments as JSON
		return c.JSON(shipments)
	}
}
