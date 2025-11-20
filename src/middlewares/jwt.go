package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	authpb "app/src/proto/auth"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var authClient authpb.AuthServiceClient
var authConn *grpc.ClientConn

func InitAuthGRPC() {
	var err error
	authConn, err = grpc.NewClient("auth:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("No se pudo conectar al servicio Auth gRPC: " + err.Error())
	}
	authClient = authpb.NewAuthServiceClient(authConn)
}

// AuthMiddleware valida el token JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token no proporcionado"})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Llamada gRPC al servicio Auth
		resp, err := authClient.VerifyToken(context.Background(), &authpb.VerifyTokenRequest{
			Token: token,
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Error al validar token: " + err.Error()})
			c.Abort()
			return
		}

		if !resp.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		// Guardamos user info en contexto
		c.Set("user_id", resp.UserId)
		c.Set("username", resp.Username)
		c.Next()
	}
}
