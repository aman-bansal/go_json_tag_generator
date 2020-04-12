package test

import "github.com/aman-bansal/go_get_set_generator/test/sample"

type SampleObject3 struct {
	Id          string                 `tag:"Unnecessarytag";json:"id"`
	Name        string                 `json:"name"`
	Age         int64                  `json:"age"`
	IsMale      bool                   `json:"is_male"`
	AnotherUser []sample.SampleObject2 `json:"another_user"`
}

type SampleObject4 struct {
	Id     string              `json:"id"`
	Name   string              `json:"name"`
	IsMale sample.SampleObject `json:"is_male"`
	Age    SampleObject3       `json:"age"`
}
