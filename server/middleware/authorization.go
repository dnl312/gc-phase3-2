package middleware

import (
	"context"
	"log"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func UnaryAuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx, err := AuthInterceptor(ctx)
	if err != nil {
		return nil, err
	}
	return handler(ctx, req)
}

func CustomJWTMiddleware(ctx context.Context) (context.Context, error) {	// return func(c echo.Context) error {
	// 	authHeader := c.Request().Header.Get("Authorization")
	// 	if authHeader == "" {
	// 		return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization Header")
	// 	}

	// 	parts := strings.Split(authHeader, " ")
	// 	if len(parts) != 2 || parts[0] != "Bearer" {
	// 		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization Header Format")
	// 	}
	// 	tokenString := parts[1]

	// 	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
	// 		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
	// 			return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unexpected signing method")
	// 		}
	// 		return []byte(os.Getenv("JWT_SECRET")), nil
	// 	})
	// 	if err != nil || !token.Valid {
	// 		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
	// 	}

	// 	c.Set("user", token)
	// 	return next(c)
	// }

	return ctx, nil
}

func AuthInterceptor(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		log.Println("No metadata found")
		return nil, status.Errorf(codes.Unauthenticated, "Unauthorized")
	}

	token := md["authorization"]
	if len(token) == 0 || strings.Contains(token[0],"Bearer"){
		log.Println("Invalid or Missing token")
		return nil, status.Errorf(codes.Unauthenticated, "Unauthorized")
	}
	
	return ctx, nil
	// userToken := c.Get("user")
	// if userToken == nil {
	// 	return 0, errors.New("user token is missing from context")
	// }

	// token, ok := userToken.(*jwt.Token)
	// if !ok {
	// 	return 0, errors.New("invalid token format")
	// }

	// claims, ok := token.Claims.(jwt.MapClaims)
	// if !ok || !token.Valid {
	// 	return 0, errors.New("invalid or malformed token claims")
	// }

	// // Extract userID
	// userID, ok := claims["user_id"].(float64)
	// if !ok {
	// 	return 0, errors.New("User ID not found or invalid in token claims")
	// }

	// return int(userID), nil
}