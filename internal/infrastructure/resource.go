package infrastructure

import (
	"fmt"

	"github.com/gofiber/storage/redis"
	"github.com/golang-jwt/jwt/v4"
	"github.com/valyala/fasthttp"
	"github.com/zercle/gofiber-skelton/internal/datasources"
	"github.com/zercle/gofiber-skelton/pkg/models"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func InitResources(fasthttpClient *fasthttp.Client, mainDbConn *gorm.DB, logDbConn *mongo.Database, redisStorage *redis.Storage, jwtResources models.JwtResources) models.Resources {
	return models.Resources{
		FastHttpClient: fasthttpClient,
		MainDbConn:     mainDbConn,
		LogDbConn:      logDbConn,
		RedisStorage:   redisStorage,
		JwtResources:   jwtResources,
	}
}

func InitJwt(privateKeyPath, publicKeyPath string) (jwtResources models.JwtResources, err error) {
	jwtResources.JwtSignKey, jwtResources.JwtVerifyKey, jwtResources.JwtSigningMethod, err = datasources.InitJwtLocalKey(privateKeyPath, publicKeyPath)
	jwtResources.JwtKeyfunc = func(token *jwt.Token) (publicKey interface{}, err error) {
		if jwtResources.JwtVerifyKey == nil {
			err = fmt.Errorf("JWTVerifyKey not init yet")
		}
		return jwtResources.JwtVerifyKey, err
	}
	jwtResources.JwtParser = jwt.NewParser()
	return
}
