package controllers__test

import (
	"testing"
)

// TESTS

func TestAnimeListCorrect(t *testing.T) {
	res := getReq(t, "/api/v1/anime/list/0", 200)
	AssertEqualInt(t, res.Success, 1)
	AssertEqualBool(t, (res.Len > 0), true)
}

func TestAnimeListNoParameter(t *testing.T) {
	res := getReq(t, "/api/v1/anime/list", 200)
	AssertEqualInt(t, res.Success, 1)
	AssertEqualBool(t, (res.Len > 0), true)
}

func TestAnimeListNegativIndex(t *testing.T) {
	getReq(t, "/api/v1/anime/list/-1", 400)
}

func TestAnimeListIncorrectParameterValue(t *testing.T) {
	getReq(t, "/api/v1/anime/list/onlyIntegerHere", 400)
}

func TestAnimeListBigIntParameter(t *testing.T) {
	getReq(t, "/api/v1/anime/list/999999999999999999999999999999999999999999999", 400)
}
