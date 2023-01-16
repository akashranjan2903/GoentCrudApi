package controller

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gocrud/ent"
	"github.com/gocrud/ent/user"
	"github.com/gocrud/src/utilities"
	"github.com/gocrud/src/utilities/response"
	"github.com/google/uuid"
)

type userController struct {
	db      *ent.Client
	success response.ApiResSuccess
	error   response.ApiResError
}

// NewuserController ...
func NewUserController(db *ent.Client) *userController {
	return &userController{
		db:      db,
		success: response.ApiResSuccess{},
		error:   response.ApiResError{},
	}
}

type userInputSchema struct {
	ent.User
}

func (u *userController) Create(c *gin.Context) {
	ctx := context.Background()

	var userInput userInputSchema
	if err := c.ShouldBindJSON(&userInput); err != nil {
		u.error.BadRequest400(c, err, err.Error())
		return
	}
	log.Println(userInput)

	user, err := u.db.User.Create().
		SetFirstName(userInput.FirstName).
		SetLastName(userInput.LastName).
		SetEmail(userInput.Email).
		SetPassword(userInput.Password).
		SetIsActive(*userInput.IsActive).
		Save(ctx)

	if err != nil {
		u.error.BadRequest400(c, err, err.Error())
		return
	}

	u.success.Created201(c, "User created successfully", user)
}

func (u *userController) Delete(c *gin.Context) {
	ctx := context.Background()
	id := uuid.MustParse(c.Param("id"))

	err := u.db.User.DeleteOneID(id).Exec(ctx)
	if err != nil {
		u.error.BadRequest400(c, err, err.Error())
		return
	}

	u.success.NoContent204(c, "User deleted successfully")
}

// fetch by id
func (u *userController) Get(c *gin.Context) {
	ctx := context.Background()
	id := uuid.MustParse(c.Param("id"))

	user, err := u.db.User.Query().Where(user.IDEQ(id)).
		Only(ctx)
	if err != nil {
		u.error.BadRequest400(c, err, err.Error())
		return
	}

	u.success.Ok200(c, "user fetched successfully", user)
}

func (u *userController) GetAll(c *gin.Context) {
	ctx := context.Background()
	limit, offset := utilities.GetPagination(c)

	alluser, err := u.db.User.Query().
		Order(ent.Asc(user.FieldFirstName)).
		Limit(limit).
		Offset(offset).
		All(ctx)
	if err != nil {
		u.error.BadRequest400(c, err, err.Error())
		return
	}

	u.success.Ok200(c, "Users fetched successfully", alluser)
}

func (u *userController) Update(c *gin.Context) {
	ctx := context.Background()
	id := uuid.MustParse(c.Param("id"))
	var userInput userInputSchema
	if err := c.ShouldBindJSON(&userInput); err != nil {
		u.error.BadRequest400(c, err, err.Error())
		return
	}

	userUpdate := u.db.User.UpdateOneID(id)
	if userInput.FirstName != "" {
		userUpdate.SetFirstName(userInput.FirstName)
	}
	if userInput.LastName != "" {

		userUpdate.SetLastName(userInput.LastName)
	}
	// if userInput.CarID != uuid.Nil {
	// 	userUpdate.SetCarsID(userInput.CarID)
	// }

	user, err := userUpdate.Save(ctx)
	if err != nil {
		u.error.BadRequest400(c, err, err.Error())
		return
	}

	u.success.Ok200(c, "User updated successfully", user)
}
