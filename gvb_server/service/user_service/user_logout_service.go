package user_service

import (
	"gvb_server/service/redis_service"
	"gvb_server/untils/jwts"
	"time"
)

func (UserService) Logout(claims *jwts.CustomClaims, token string) error {
	//claims.ExpiresAt token的过期时间
	// fmt.Println(claims.ExpiresAt)
	//计算距离过期的剩余时间
	exp := claims.ExpiresAt
	now := time.Now()

	diff := exp.Time.Sub(now)

	// fmt.Println(diff)

	return redis_service.Logout(token, diff)
}
