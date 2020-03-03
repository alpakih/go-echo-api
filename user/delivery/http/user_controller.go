package http

import (
	"github.com/labstack/echo"
	"github.com/ulule/paging"
	"go-echo-api/infrastructure/database"
	"go-echo-api/infrastructure/response"
	"go-echo-api/user"
	"go-echo-api/utils"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type userController struct {
	userRepository user.Repository
	userMapper     *user.Mapper
}

func NewUserController(s user.Repository) *userController {
	return &userController{userRepository: s,
		userMapper: user.NewUserMapper(),
	}
}


func (c *userController) FindById(ctx echo.Context) error {
	id := ctx.Param("id")
	result, err := c.userRepository.FindById(id)
	if err != nil {
		return response.InternalServerError(ctx, utils.InternalServerError, nil, err.Error())
	}
	return response.SingleData(ctx, utils.OK, c.userMapper.Map(result), nil)
}

func (c *userController) Store(ctx echo.Context) error {
	var dto user.Dto
	if err := ctx.Bind(&dto); err != nil {
		return response.InternalServerError(ctx, utils.InternalServerError, nil, err.Error())
	}
	if err := ctx.Validate(dto); err != nil {
		errorData := make(echo.Map)
		for _, v := range err.(validator.ValidationErrors) {
			errorData[v.Field()] = v.Tag()
		}
		return response.ValidationError(ctx, utils.ValidationError, nil, err.Error())
	}
	result, err := c.userRepository.Save(dto)
	if err != nil {
		return response.InternalServerError(ctx, utils.InternalServerError, nil, err.Error())
	}
	return response.SingleData(ctx, utils.OK, c.userMapper.Map(result), nil)
}

func (c *userController) FindAll(ctx echo.Context) error {
	result, err := c.userRepository.FindAll()
	if err != nil {
		return response.InternalServerError(ctx, utils.InternalServerError, nil, err.Error())
	}
	store, err := paging.NewGORMStore(database.GetLinkDb(), &result)
	if err != nil {
		return response.InternalServerError(ctx, utils.InternalServerError, nil, err.Error())
	}
	options := paging.NewOptions()
	request, _ := http.NewRequest(ctx.Request().Method, ctx.Request().URL.String(), nil)
	paginator, _ := paging.NewOffsetPaginator(store, request, options)
	err = paginator.Page()
	if err != nil {
		return response.InternalServerError(ctx, utils.InternalServerError, nil, err.Error())
	}
	return response.Paginate(ctx, utils.OK, paginator, c.userMapper.MapList(result), nil)
}

func (c *userController) Update(ctx echo.Context) error {
	id := ctx.Param("id")
	var dto user.Dto
	if err := ctx.Bind(&dto); err != nil {
		return response.InternalServerError(ctx, utils.InternalServerError, nil, err.Error())
	}
	if err := ctx.Validate(dto); err != nil {
		errorData := make(echo.Map)
		for _, v := range err.(validator.ValidationErrors) {
			errorData[v.Field()] = v.Tag()
		}
		return response.ValidationError(ctx, utils.ValidationError, nil, err.Error())
	}
	result, err := c.userRepository.Update(id, dto)
	if err != nil {
		return response.InternalServerError(ctx, utils.InternalServerError, nil, err.Error())
	}
	return response.SingleData(ctx, utils.OK, c.userMapper.Map(result), nil)
}

func (c *userController) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	err:=c.userRepository.Delete(id)
	if err !=nil {
		return response.InternalServerError(ctx, utils.InternalServerError, nil, err.Error())
	}
	return response.SingleData(ctx, utils.OK,nil, nil)
}


