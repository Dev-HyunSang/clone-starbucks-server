package auth

import (
	"context"
	"github.com/dev-hyunsang/clone-stackbuck-backend/config"
	"github.com/dev-hyunsang/clone-stackbuck-backend/db"
	"github.com/dev-hyunsang/clone-stackbuck-backend/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

func CreateToken(userId uuid.UUID) (*models.TokenDetails, error) {
	td := new(models.TokenDetails)
	var err error

	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.New().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshToken = uuid.New().String()

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_uuid"] = userId
	atClaims["exo"] = time.Now().Add(time.Minute * 15).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(config.GetDotEnv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userId
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(config.GetDotEnv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}

	return td, nil
}

func CreateAuth(userId uuid.UUID, td *models.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	client, err := db.ConnectRedis()
	if err != nil {
		return err
	}

	err = client.Set(context.Background(), td.AccessUuid, userId.String(), at.Sub(now)).Err()
	if err != nil {
		return err
	}

	err = client.Set(context.Background(), td.RefreshUuid, userId.String(), rt.Sub(now)).Err()
	if err != nil {
		return err
	}

	return nil
}
