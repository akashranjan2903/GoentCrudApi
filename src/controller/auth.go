package controller

import (
	"context"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gocrud/ent"
	"github.com/gocrud/ent/user"
	"github.com/gocrud/src/utilities/auth"
	"github.com/gocrud/src/utilities/response"
	"github.com/google/uuid"
)

type authController struct {
	db      *ent.Client
	success response.ApiResSuccess
	error   response.ApiResError
}

type userTokenResponse struct {
	AccessToken auth.JWTTokenResponse `json:"access_token"`
}
type changePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

// NewAuthController ...
func NewAuthController(db *ent.Client) *authController {
	return &authController{
		db:      db,
		success: response.ApiResSuccess{},
		error:   response.ApiResError{},
	}
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ac *authController) Login(c *gin.Context) {
	ctx := context.Background()
	var loginRequest loginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		ac.error.BadRequest400(c, err, err.Error())
		return
	}

	// Check if the user exists with the same email
	userQuery := ac.db.User.Query().Where(user.Email(loginRequest.Email))
	isUserExists, err := userQuery.Exist(ctx)
	if err != nil {
		ac.error.InternalServerError500(c, err)
		return
	}

	// Guard clause -> If the user not exists with the same email
	if !isUserExists {
		ac.error.BadRequest400(c, err, "User not found")
		return
	}

	// Get the user
	userInstance, err := ac.db.User.Query().Where(user.Email(loginRequest.Email)).Only(ctx)
	if err != nil {
		ac.error.InternalServerError500(c, err)
		return
	}

	// Check if the user is Active
	if !*userInstance.IsActive {
		ac.error.BadRequest400(c, err, "Account Inactive, Please verify your Account")
		return
	}

	// Check if the password is correct
	isPasswordCorrect, err := auth.VerifyPassword(loginRequest.Password, userInstance.Password)
	if err != nil {
		ac.error.InternalServerError500(c, err)
		return
	}
	if !isPasswordCorrect {
		ac.error.BadRequest400(c, err, "Password is incorrect")
		return
	}

	// Generate the JWT token
	accessToken, err := auth.GenerateJWT(auth.Access, userInstance.ID.String())
	if err != nil {
		ac.error.InternalServerError500(c, err)
		return
	}

	// refreshToken, err := auth.GenerateJWT(auth.Refresh, userInstance.ID.String())
	// if err != nil {
	// 	ac.error.InternalServerError500(c, err)
	// 	return
	// }

	// Return the response
	ac.success.Ok200(c, "Login successful", userTokenResponse{
		AccessToken: accessToken,
	})

}

func (ac *authController) ChangePassword(c *gin.Context) {
	ctx := context.Background()
	// Get Access Token from Authorization Header
	accessTokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer")
	if accessTokenString == "" {
		ac.error.BadRequest400(c, nil, "Access token is required")
		return
	}
	fmt.Println("token is:", accessTokenString)

	// Decode the access token
	_, userId, err := auth.DecodeJWT(accessTokenString, auth.Access)
	if err != nil {
		ac.error.BadRequest400(c, err, "Invalid access token")
		return
	}

	// Parse the user id to uuid
	userUuid := uuid.MustParse(userId)

	// Check if the user exists with the same id
	userInstance, err := ac.db.User.Get(ctx, userUuid)
	if err != nil {
		ac.error.InternalServerError500(c, err)
		return
	}

	// Get the request body
	var changePasswordRequest changePasswordRequest
	err = c.ShouldBindJSON(&changePasswordRequest)
	if err != nil {
		ac.error.BadRequest400(c, err, "Invalid request body")
		return
	}

	// Validate the old password
	isOk, err := auth.VerifyPassword(changePasswordRequest.OldPassword, userInstance.Password)
	if err != nil {
		ac.error.InternalServerError500(c, err)
		return
	}

	// If the old password is not valid
	if !isOk {
		ac.error.BadRequest400(c, nil, "Invalid old password")
		return
	}

	// Validate the new password
	isOkNewPass, msg := auth.PasswordValidator(changePasswordRequest.NewPassword)
	if !isOkNewPass {
		ac.error.BadRequest400(c, nil, msg)
		return
	}

	// Hash the new password
	// hashedPassword, err := auth.HashPassword(changePasswordRequest.NewPassword)
	// if err != nil {
	// 	ac.error.InternalServerError500(c, err)
	// 	return
	// }

	// Update the password
	_, err = userInstance.Update().SetPassword(changePasswordRequest.NewPassword).Save(ctx)
	if err != nil {
		ac.error.InternalServerError500(c, err)
		return
	}

	// Return the response
	ac.success.Ok200(c, "Password changed successfully", nil)
}
