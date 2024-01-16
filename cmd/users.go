package cmd

import (
	"database/sql"
	"github.com/dev-hyunsang/clone-stackbuck-backend/auth"
	"github.com/dev-hyunsang/clone-stackbuck-backend/db"
	"github.com/dev-hyunsang/clone-stackbuck-backend/dto"
	"github.com/dev-hyunsang/clone-stackbuck-backend/models"
	"github.com/dev-hyunsang/clone-stackbuck-backend/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

const (
	BadRequestMsg            string = "올바르지 않은 요청입니다. 확인 후 다시 시도하세요,"
	BadRequestLoginMsg       string = "이메일 혹은 패스워드를 확인 후 다시 시도해 주세요."
	ErrTotalMsg              string = "요청을 처리하던 도중 오류가 발생했어요. 잠시후 다시 시도해 주세요."
	ErrTimeParseMsg          string = "사용자님의 생일을 변환하던 도중 오류가 발생했어요. 잠시후 다시 시도해 주세요."
	ErrFailedConnectMySQLMsg string = "사용자님의 정보를 조호히ㅏ던 도중 데이터베이스와 연결이 되지 않았어요. 잠시후 시도해 주세요."
	ErrFailedAuthToken       string = "사용자님의 정보를 통해서 로그인에 필요한 정보를 만들던 도중 오류가 발생했어요. 잠시후 다시 시도해 주세요."
)

func SignUpUserHandler(ctx *fiber.Ctx) error {
	req := new(dto.RequestSignUp)
	if err := ctx.BodyParser(req); err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrResponse{
			Stats: dto.Status{
				Code:    fiber.StatusBadRequest,
				Stats:   "bad request",
				Message: BadRequestMsg,
			},
			RespondedAt: time.Now(),
		})
	}

	convertDate, err := time.Parse("2006-01-02", req.Birthday)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrResponse{
			Stats: dto.Status{
				Stats:   "type error",
				Code:    fiber.StatusInternalServerError,
				Message: ErrTimeParseMsg,
			},
			RespondedAt: time.Now(),
		})
	}

	hasedPw, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&dto.ErrResponse{})
	}

	client, err := db.ConnectMySQL()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrResponse{
			Stats: dto.Status{
				Stats:   "error",
				Code:    fiber.StatusInternalServerError,
				Message: ErrTotalMsg,
			},
			RespondedAt: time.Now(),
		})
	}

	user := models.Users{
		Id:          uuid.New(),
		Email:       req.Email,
		Password:    string(hasedPw),
		Name:        req.Name,
		DisplayName: req.DisplayName,
		Birthday:    convertDate,
		PhoneNumber: req.PhoneNumber,
		AllowMarketing: sql.NullBool{
			Bool:  req.AllowMarketing,
			Valid: false,
		},
	}

	result, err := repositories.CreateUser(client, &user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON("")
	}

	log.Println(result)

	return ctx.Status(fiber.StatusOK).JSON(dto.ResponseSignUp{
		Status: dto.Status{
			Code:    fiber.StatusOK,
			Stats:   "ok",
			Message: "성공적으로 새로운 사용자의 정보를 만들었습니다.",
		},
		Data:        user,
		RespondedAt: time.Now(),
	})
}

func LoginUserHandler(ctx *fiber.Ctx) error {
	req := new(dto.RequestLogin)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrResponse{})
	}

	if len(req.Email) == 0 && len(req.Password) == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrResponse{})
	}

	client, err := db.ConnectMySQL()
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrResponse{
			Stats: dto.Status{
				Stats:   "error - failed to connect MySQL",
				Code:    fiber.StatusBadRequest,
				Message: ErrFailedConnectMySQLMsg,
			},
			RespondedAt: time.Now(),
		})
	}

	// 모든 사용자의 정보의 조회는 Email → Pk → UUID로 조회함.
	result, err := repositories.GetuserByEmail(client, req.Email)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrResponse{})
	}

	// 데이터이스 상에 없는 메일
	if req.Email != result.Email {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrResponse{})
	}

	// 데이터베이스 상에 저장된 패스워스와 입력한 패스워드 비교
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(req.Password))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrResponse{})
	}

	ts, err := auth.CreateToken(result.Id)
	if err != nil {
		log.Print(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrResponse{
			Stats: dto.Status{
				Stats:   "error - create auth token",
				Code:    fiber.StatusInternalServerError,
				Message: ErrFailedAuthToken,
			},
			RespondedAt: time.Now(),
		})
	}

	err = auth.CreateAuth(result.Id, ts)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrResponse{
			Stats: dto.Status{
				Stats:   "error - saved to Redis",
				Code:    fiber.StatusInternalServerError,
				Message: ErrFailedAuthToken,
			},
			RespondedAt: time.Now(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.ResponseLogin{
		Status: dto.Status{
			Stats:   "ok",
			Code:    fiber.StatusOK,
			Message: "회원님 어서오세요!~",
		},
		Data: dto.LoginData{
			AccessToken:  ts.AccessToken,
			RefreshToken: ts.RefreshToken,
		},
		RespondedAt: time.Now(),
	})
}
