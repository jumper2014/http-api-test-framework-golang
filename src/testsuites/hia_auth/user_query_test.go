package hia_auth

import (
	"testing"
	"libs/constant"
	"libs/api/hia_auth"
)



func TestUserQuery(t *testing.T) {

	indexType := ""
	username := ""
	indexValue := ""



	resp, err := hia_auth.UserQuery(indexType, username, indexValue)

	if err != nil {
		t.Fatal("\t\tShould be able to make the Get call",
			constant.CheckMark, err)

	}
	t.Log("\t\tShould be able to make the Get call", constant.CheckMark)

	defer resp.Body.Close()

	if resp.StatusCode == constant.StatusCode200 {
		t.Logf("\t\tShould receive a \"%d\" status. %v", constant.StatusCode200, constant.CheckMark)
	} else {
		t.Errorf("\t\tShould receive a \"%d\" status. %v %v", constant.StatusCode200, constant.BallotX, resp.StatusCode)
	}


}
