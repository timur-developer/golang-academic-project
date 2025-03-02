package handlers

import (
	"academyProject/internal/userService"
	"academyProject/internal/web/users"
	"context"
)

type userHandler struct {
	userService *userService.UserService
}

func (h *userHandler) DeleteUsers(_ context.Context, request users.DeleteUsersRequestObject) (users.DeleteUsersResponseObject, error) {
	userRequest := request.Body

	deletedUser, err := h.userService.DeleteUserByID(*userRequest.Id)
	if err != nil {
		return nil, err
	}
	response := users.DeleteUsers200JSONResponse{
		Id:        &deletedUser.ID,
		Email:     &deletedUser.Email,
		Password:  &deletedUser.Password,
		CreatedAt: &deletedUser.CreatedAt,
		UpdatedAt: &deletedUser.UpdatedAt,
	}
	return response, nil
}

func (h *userHandler) PatchUsers(_ context.Context, request users.PatchUsersRequestObject) (users.PatchUsersResponseObject, error) {
	userRequest := request.Body

	updates := make(map[string]interface{})
	if userRequest.Email != nil {
		updates["email"] = *userRequest.Email
	}
	if userRequest.Password != nil {
		updates["password"] = *userRequest.Password
	}

	updatedUser, err := h.userService.UpdateUserByID(*userRequest.Id, updates)
	if err != nil {
		return nil, err
	}
	response := users.PatchUsers200JSONResponse{
		Id:        &updatedUser.ID,
		Email:     &updatedUser.Email,
		Password:  &updatedUser.Password,
		CreatedAt: &updatedUser.CreatedAt,
		UpdatedAt: &updatedUser.UpdatedAt,
	}

	return response, nil
}

func (h *userHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.userService.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, tsk := range allUsers {
		user := users.User{
			Id:        &tsk.ID,
			Email:     &tsk.Email,
			Password:  &tsk.Password,
			CreatedAt: &tsk.CreatedAt,
			UpdatedAt: &tsk.UpdatedAt,
		}
		response = append(response, user)
	}

	return response, nil
}

func (h *userHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body

	userToCreate := userService.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}
	createdUser, err := h.userService.CreateUser(userToCreate)
	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:        &createdUser.ID,
		Email:     &createdUser.Email,
		Password:  &createdUser.Password,
		CreatedAt: &createdUser.CreatedAt,
		UpdatedAt: &createdUser.UpdatedAt,
	}

	return response, nil
}

func NewUserHandler(service *userService.UserService) *userHandler {
	return &userHandler{
		userService: service,
	}
}
