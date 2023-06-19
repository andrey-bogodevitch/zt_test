package service

import "testing"

func TestUser_ComputeHmac(t *testing.T) {
	var s Service

	expected := "2362731ec02471ab810f74117f2bbad9756e2c2a2335d058456673f128dd6afc846dd3805b0a5fdf3e0cee17cdb94cf478c4ceb6f254dcf1ae31c3591d63a4d4"
	result, err := s.ComputeHmac("text", "key")
	if err != nil {
		t.Fatal(err)
	}
	if result != expected {
		t.Fatal("wrong result")
	}

}
