package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	fmt.Println(res.Message)
	if res.Status != shim.OK {
		fmt.Println("Init failed", res.String())
		t.FailNow()
	}
}

func TestEduApp(t *testing.T) {
	scc := new(SimpleChaincode)
	stub := shim.NewMockStub("edu_app", scc)
	// init
	checkInit(t, stub, [][]byte{[]byte("init")})
	// //注册学校
	res0 := stub.MockInvoke("1", [][]byte{[]byte("RegisterSchool"), []byte("cqu")})
	// fmt.Println("res0", res0)
	fmt.Println("payload", res0.Payload)
	school := new(School)
	json.Unmarshal(res0.Payload, school)
	fmt.Println("school", school)
	//注册学生
	res := stub.MockInvoke("1", [][]byte{[]byte("RegisterStudent"), []byte("jianjun"), []byte(school.Address)})
	fmt.Println(res.Payload)
	s := new(Student)
	json.Unmarshal(res.Payload, s)
	fmt.Println(s)

	res5 := stub.MockInvoke("1", [][]byte{[]byte("UploadDiploma"), []byte(school.Address), []byte(s.Address), []byte("2016"), []byte("4"), []byte("computer"), []byte("0")})
	fmt.Println(string(res5.Payload))

	res6 := stub.MockInvoke("1", [][]byte{[]byte("QueryStudentByAddress"), []byte(s.Address)})
	fmt.Println(string(res6.Payload))
	s1 := new(Student)
	json.Unmarshal(res6.Payload, s1)
	fmt.Println(s1)

	res3 := stub.MockInvoke("1", [][]byte{[]byte("RegisterCompany"), []byte("baidu")})
	fmt.Println(string(res3.Payload))
	comp := new(Company)
	json.Unmarshal(res3.Payload, comp)
	fmt.Println(comp)

	res7 := stub.MockInvoke("1", [][]byte{[]byte("VerifyDiplomaAddress"), []byte(comp.Address), []byte(s1.DiplomaAddr[0])})
	fmt.Println(res7.Payload)

	res8 := stub.MockInvoke("1", [][]byte{[]byte("QueryRecordByID"), []byte(s1.DiplomaAddr[0])})
	fmt.Println(string(res8.Payload))
	// res4 := stub.MockInvoke("1", [][]byte{[]byte("QuerySchoolByAddress"), []byte(school.Address)})
	// fmt.Println(string(res4.Payload))
	// sc := new(School)
	// json.Unmarshal(res4.Payload, sc)
	// fmt.Println(sc)

	t.FailNow()
}
