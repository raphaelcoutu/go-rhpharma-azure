package routes

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/raphaelcoutu/go-rhpharma-azure/database"
	"github.com/raphaelcoutu/go-rhpharma-azure/logger"
)

type Constraint struct {
	Id                  uint      `json:"id"`
	StartDate           time.Time `json:"startDate"`
	EndDate             time.Time `json:"endDate"`
	Weight              bool      `json:"weight"`
	Comment             string    `json:"comment"`
	Status              uint      `json:"status"`
	NumberOfOccurrences uint      `json:"numberOfOccurrences"`
	Disposition         uint      `json:"disposition"`
	IsDayOfWeek         bool      `json:"isDayOfWeek"`
	Day                 uint      `json:"day"`
	Day1                uint      `json:"day1"`
	Discriminator       string    `json:"discriminator"`
	ConstraintTypeId    uint      `json:"constraintTypeId"`
	ConstraintTypeName  string    `json:"constraintTypeName"`
	UserId              uint      `json:"userId"`
	UserFirstName       string    `json:"userFirstName"`
	UserLastName        string    `json:"userLastName"`
}

func GetConstraints(c *fiber.Ctx) error {
	ctx := context.Background()
	db := database.DB

	err := db.PingContext(ctx)
	if err != nil {
		logger.Log("Error: PingContext. Message: " + err.Error())
		return err
	}

	layout := "2006-01-02"
	qStartDate := c.Query("startDate")
	qEndDate := c.Query("endDate")
	var startDate, endDate time.Time

	if qStartDate == "" || qEndDate == "" {
		return c.SendString("Missing parameter.")
	}

	if qStartDate != "" {
		startDate, err = time.Parse(layout, qStartDate)
		if err != nil {
			logger.Log("Error: StartDate Parse. Message: " + err.Error())
			return err
		}
	}

	if qEndDate != "" {
		endDate, err = time.Parse(layout, qEndDate)
		if err != nil {
			logger.Log("Error: EndDate Parse. Message: " + err.Error())
			return err
		}
	}

	tsql := fmt.Sprintf("SELECT C.Id as Constraint_id, C.StartDate, C.EndDate, C.Weight, C.Comment," +
		"C.Status, C.NumberOfOccurrences, C.Disposition, C.IsDayOfWeek, C.Day, C.Day1, C.Discriminator," +
		"Ct.Id as ConstraintType_id, Ct.Name," +
		"U.Id as User_id, U.FirstName, U.LastName " +
		"FROM Constraints As C " +
		"JOIN ConstraintTypes AS Ct ON Ct.Id = C.TypeId " +
		"JOIN Users As U ON U.Id = C.UserId " +
		"WHERE ((StartDate >= @StartDate AND EndDate <= @EndDate) OR StartDate <= @EndDate AND EndDate >= @StartDate) AND TypeID <> 79 " +
		"ORDER BY U.Lastname;")

	stmt, err := db.Prepare(tsql)
	if err != nil {
		logger.Log("Error: Prepare. Message: " + err.Error())
		return err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, sql.Named("StartDate", startDate), sql.Named("EndDate", endDate))
	if err != nil {
		logger.Log("Error: QueryContext. Message: " + err.Error())
		return err
	}

	defer rows.Close()

	constraints := []Constraint{}

	for rows.Next() {
		var c Constraint
		var numberOfOccurrences, disposition, day, day1 sql.NullInt16
		var isDayOfWeek sql.NullBool

		err := rows.Scan(&c.Id, &c.StartDate, &c.EndDate, &c.Weight, &c.Comment,
			&c.Status, &numberOfOccurrences, &disposition, &isDayOfWeek, &day, &day1, &c.Discriminator,
			&c.ConstraintTypeId, &c.ConstraintTypeName, &c.UserId, &c.UserFirstName, &c.UserLastName)
		if err != nil {
			return err
		}

		if numberOfOccurrences.Valid {
			c.NumberOfOccurrences = uint(numberOfOccurrences.Int16)
		}

		if disposition.Valid {
			c.Disposition = uint(disposition.Int16)
		}

		if isDayOfWeek.Valid {
			c.IsDayOfWeek = isDayOfWeek.Bool
		}

		if day.Valid {
			c.Day = uint(day.Int16)
		}

		if day1.Valid {
			c.Day1 = uint(day.Int16)
		}

		constraints = append(constraints, c)
	}

	jsonData, _ := json.Marshal(constraints)
	return c.SendString(string(jsonData))
}
