package utils

import (
	"testing"
)

func TestCreateJwtToken(t *testing.T) {
	jwtToken, err := CreateJwtToken("xulei", 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(jwtToken)
	jwtToken="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJOYW1lIjoieHVsZWkiLCJEY0lkIjoxLCJleHAiOjE2MDcxNzQzMzAsImlhdCI6MTYwNzA2NjMzMCwiaXNzIjoia2l0X3Y0IiwibmJmIjoxNjA3MDY2MzMwLCJzdWIiOiJsb2dpbiJ9.NYAT1ZYZ-rrSw8wBFE8_chCWcg6fYomxbCZcHPqq7tw"
	jwtInfo, err := ParseToken(jwtToken)
	if err != nil {
		t.Error(err)
	}
	user_id := jwtInfo["DcId"].(int32)
	name := jwtInfo["Name"].(string)
	t.Log(user_id,name)
}
func TestMd5(t *testing.T) {
	password:= Md5Password("abc123")
	t.Log(password)
}