package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	authpb "app/src/proto/auth"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var authClient authpb.AuthServiceClient

func InitAuthGRPC() {
	var conn *grpc.ClientConn
	var err error

	// Intentos de reconexión
	for i := 0; i < 10; i++ {
		conn, err = grpc.Dial(
			"auth-service:50051",
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithBlock(), // esperar conexión
		)

		if err == nil {
			break
		}

		fmt.Println("Reintentando conexión gRPC...", err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		panic("No se pudo conectar al servicio PuntoVenta: " + err.Error())
	}
	authClient = authpb.NewAuthServiceClient(conn)
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
