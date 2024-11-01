package resp

import (
    "colab/api/grpc/user"
    "time"
)

type TokenInfo struct {
    Token     string     `json:"token"`
    ExpiresAt *time.Time `json:"expiresAt"`
}

type SigninResult struct {
    UserInfo         *user.UserInfo           `json:"userInfo"`
    OrganizationList []*user.OrganizationInfo `json:"organizationList"`
    TokenInfo        *TokenInfo               `json:"tokenInfo"`
}
