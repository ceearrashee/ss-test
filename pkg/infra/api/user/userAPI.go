package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"

	"solid-software.test-task/pkg/domain/user"
	"solid-software.test-task/pkg/framework/webservice/route"
	"solid-software.test-task/pkg/infra/api"
	"solid-software.test-task/pkg/infra/api/user/di"
	"solid-software.test-task/pkg/infra/db"
)

type (
	userAPI struct{}
)

// NewUserAPI creates a new user API.
func NewUserAPI() route.Route {
	return &userAPI{}
}

// IsProtected returns true if the route is protected by authentication.
func (*userAPI) IsProtected() bool {
	return true
}

func (*userAPI) InitRoutes(party router.Party) {
	party.Party("/user").ConfigureContainer(
		func(container *router.APIContainer) {
			container.RegisterDependency(di.InitializeUserService)

			container.Post("/", handleCreateUser)
			container.Get("s", handelGetUsers)
			singleUserRoute := container.Party("/{id:uint}")
			singleUserRoute.Get("", handleGetUser)
			singleUserRoute.Put("", handleUpdateUser)
			singleUserRoute.Delete("", handleDeleteUser)
		},
	)
}

func handleRequest(irisCtx iris.Context, action func() (any, int, error)) {
	response, status, err := action()
	if err != nil {
		api.HandleError(irisCtx, status, err)
	} else if err = irisCtx.JSON(response); err != nil {
		api.HandleError(irisCtx, iris.StatusInternalServerError, err)
	}
}

func readJSONObject(irisCtx iris.Context) (*user.Entity, error) {
	var userRQ user.Entity

	if err := irisCtx.ReadJSON(&userRQ); err != nil {
		return nil, fmt.Errorf("parse JSON: %w", err)
	}

	if userRQ.Name == "" {
		return nil, fmt.Errorf("username is empty")
	}

	return &userRQ, nil
}

func saveUser(ctx context.Context, userService user.Service, userRQ *user.Entity) error {
	if err := userService.Save(ctx, userRQ); err != nil {
		return fmt.Errorf("saving user: %w", err)
	}

	return nil
}

func handleCreateUser(irisCtx iris.Context, ctx context.Context, userService user.Service) {
	executeCreateOrUpdateUser := func() (any, int, error) {
		userRQ, err := readJSONObject(irisCtx)
		if err != nil {
			return nil, iris.StatusBadRequest, err
		}

		err = saveUser(ctx, userService, userRQ)
		if err != nil {
			return nil, iris.StatusInternalServerError, err
		}

		return userRQ, iris.StatusOK, nil
	}
	handleRequest(irisCtx, executeCreateOrUpdateUser)
}

func handleUpdateUser(irisCtx iris.Context, ctx context.Context, userService user.Service) {
	executeCreateOrUpdateUser := func() (any, int, error) {
		userRQ, err := readJSONObject(irisCtx)
		if err != nil {
			return nil, iris.StatusBadRequest, err
		}

		userID, err := irisCtx.Params().GetUint("id")
		if err != nil {
			return nil, iris.StatusBadRequest, fmt.Errorf("get user ID: %w", err)
		}

		if userID != userRQ.ID {
			return nil, iris.StatusBadRequest, fmt.Errorf("user ID in path and in body are not equal")
		}

		err = saveUser(ctx, userService, userRQ)
		if err != nil {
			return nil, iris.StatusInternalServerError, err
		}

		return &userRQ, iris.StatusOK, nil
	}
	handleRequest(irisCtx, executeCreateOrUpdateUser)
}

func handelGetUsers(irisCtx iris.Context, ctx context.Context, userService user.Service) {
	executeGetUsers := func() (any, int, error) {
		userResponses, err := userService.GetWithFilter(ctx)
		if err != nil {
			return nil, iris.StatusInternalServerError, fmt.Errorf("getting users with filter: %w", err)
		}

		return userResponses, iris.StatusOK, nil
	}
	handleRequest(irisCtx, executeGetUsers)
}

func handleGetUser(irisCtx iris.Context, ctx context.Context, userService user.Service) {
	executeGetUser := func() (any, int, error) {
		userID, err := irisCtx.Params().GetUint("id")
		if err != nil {
			return nil, iris.StatusBadRequest, fmt.Errorf("get user ID: %w", err)
		}

		userResp, err := userService.GetByID(ctx, userID)
		if err != nil {
			if errors.Is(err, db.ErrRecordNotFound) {
				return nil, iris.StatusNotFound, fmt.Errorf("getting users with filter: %w", err)
			}

			return nil, iris.StatusInternalServerError, fmt.Errorf("getting user by ID: %w", err)
		}

		return userResp, iris.StatusOK, nil
	}
	handleRequest(irisCtx, executeGetUser)
}

func handleDeleteUser(irisCtx iris.Context, ctx context.Context, userService user.Service) {
	executeDeleteUser := func() (any, int, error) {
		userID, err := irisCtx.Params().GetUint("id")
		if err != nil {
			return nil, iris.StatusBadRequest, fmt.Errorf("get user ID: %w", err)
		}

		err = userService.DeleteByID(ctx, userID)
		if err != nil {
			return nil, iris.StatusInternalServerError, fmt.Errorf("deleting user by ID: %w", err)
		}

		return nil, iris.StatusOK, nil
	}
	handleRequest(irisCtx, executeDeleteUser)
}
