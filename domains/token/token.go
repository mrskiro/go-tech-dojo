package token

import (
	"crypto/rand"
)

type Token string

// 適当な仕様がないので適当に生成する
func GenToken(digit uint64) (Token, error) {
	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	var result string
	for _, v := range b {
		// 制御文字が当たらないように調整
		result += string(v%byte(94) + 33)
	}
	return Token(result), nil
}

func (t Token) String() string {
	return string(t)
}
