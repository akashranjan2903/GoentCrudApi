package controller

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/gocrud/ent"
	"github.com/gocrud/ent/car"
	"github.com/gocrud/ent/user"
	"github.com/gocrud/src/utilities"
	"github.com/gocrud/src/utilities/response"
	"github.com/google/uuid"
)

type carController struct {
	db      *ent.Client
	success response.ApiResSuccess
	error   response.ApiResError
}

// NewuserController ...
func NewcarController(db *ent.Client) *carController {
	return &carController{
		db:      db,
		success: response.ApiResSuccess{},
		error:   response.ApiResError{},
	}
}

type carInputSchema struct {
	ent.Car
	UserID uuid.UUID `json:"user_id"`
}

func (ca *carController) Create(c *gin.Context) {
	ctx := context.Background()

	var carInput carInputSchema
	if err := c.ShouldBindJSON(&carInput); err != nil {
		ca.error.BadRequest400(c, err, err.Error())
		return
	}

	newCar, err := ca.db.Car.Create().
		SetModel(carInput.Model).
		SetUserID(carInput.UserID).
		Save(ctx)
	if err != nil {
		ca.error.BadRequest400(c, err, err.Error())
		return
	}

	ca.success.Created201(c, "Cars created successfully", newCar)
}

func (ca *carController) GetCarsbyOwner(c *gin.Context) {

	ctx := context.Background()
	userID := uuid.MustParse(c.Param("id"))
	limit, offset := utilities.GetPagination(c)

	cars, err := ca.db.Car.Query().
		Where(car.HasUserWith(user.IDEQ(userID))).
		Order(ent.Asc(car.FieldUpdatedAt)).
		Limit(limit).
		Offset(offset).
		All(ctx)
	if err != nil {
		ca.error.BadRequest400(c, err, err.Error())
		return
	}

	ca.success.Ok200(c, "cars fetched successfully", cars)
}

func (ca *carController) Delete(c *gin.Context) {
	ctx := context.Background()
	id := uuid.MustParse(c.Param("id"))

	err := ca.db.Car.DeleteOneID(id).Exec(ctx)
	if err != nil {
		ca.error.BadRequest400(c, err, err.Error())
		return
	}

	ca.success.NoContent204(c, "User deleted successfully")
}
