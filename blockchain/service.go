package blockchain

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func (setup *FabricSetup) RegisterSchool(schoolname string) (string, string, error) {
	var args []string
	args = append(args, "RegisterSchool")
	args = append(args, schoolname)

	transientDataMap := make(map[string][]byte)
	transientDataMap["result"] = []byte("register school invoke")
	request := channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1])}, TransientMap: transientDataMap}
	response, err := setup.client.Execute(request)
	if err != nil {
		return "", "", fmt.Errorf("failed to register school:%v", err)
	}
	return string(response.Payload), string(response.TransactionID), nil
}
func (setup *FabricSetup) RegisterStudent(studentname, schooladdr string) (string, string, error) {
	var args []string
	args = append(args, "RegisterStudent")
	args = append(args, studentname)
	args = append(args, schooladdr)

	request := channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2])}}
	response, err := setup.client.Execute(request)
	if err != nil {
		return "", "", fmt.Errorf("failed to register student:%v", err)
	}
	return string(response.Payload), string(response.TransactionID), nil

}

func (setup *FabricSetup) RegisterCompany(companyname string) (string, string, error) {
	var args []string
	args = append(args, "RegisterCompany")
	args = append(args, companyname)

	request := channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1])}}
	response, err := setup.client.Execute(request)
	if err != nil {
		return "", "", fmt.Errorf("failed to register student:%v", err)
	}
	return string(response.Payload), string(response.TransactionID), nil
}
func (setup *FabricSetup) UploadDiploma(schooladdr, studentAddr, startYear, duration, major, status, studentName, schoolName, diplomaType string) (string, string, error) {
	var args []string
	args = append(args, "UploadDiploma")
	args = append(args, schooladdr)
	args = append(args, studentAddr)
	args = append(args, startYear)
	args = append(args, duration)
	args = append(args, major)
	args = append(args, status)
	args = append(args, studentName)
	args = append(args, schoolName)
	args = append(args, diplomaType)

	request := channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2]), []byte(args[3]), []byte(args[4]), []byte(args[5]), []byte(args[6]), []byte(args[7]), []byte(args[8]), []byte(args[9])}}
	response, err := setup.client.Execute(request)
	if err != nil {
		return "", "", fmt.Errorf("failed to upload:%v", err)
	}
	return string(response.Payload), string(response.TransactionID), nil
}

func (setup *FabricSetup) QuerySchool(schoolAddr string) (string, string, error) {
	var args []string
	args = append(args, "QuerySchoolByAddress")
	args = append(args, schoolAddr)

	request := channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1])}}
	response, err := setup.client.Execute(request)
	if err != nil {
		return "", "", fmt.Errorf("failed to query school:%v", err)
	}
	return string(response.Payload), string(response.TransactionID), nil
}
func (setup *FabricSetup) QueryStudent(address string) (string, string, error) {
	var args []string
	args = append(args, "QueryStudentByAddress")
	args = append(args, address)

	request := channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1])}}
	response, err := setup.client.Query(request)
	if err != nil {
		return "", "", fmt.Errorf("failed to student school:%v", err)
	}
	return string(response.Payload), string(response.TransactionID), nil
}
func (setup *FabricSetup) QueryCompany(address string) (string, string, error) {
	var args []string
	args = append(args, "QueryCompanyByAddress")
	args = append(args, address)

	request := channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1])}}
	response, err := setup.client.Query(request)
	if err != nil {
		return "", "", fmt.Errorf("failed to query company:%v", err)
	}
	return string(response.Payload), string(response.TransactionID), nil
}

func (setup *FabricSetup) QueryDiploma(address string) (string, string, error) {
	var args []string
	args = append(args, "QueryDiplomaByAddress")
	args = append(args, address)

	request := channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1])}}
	response, err := setup.client.Query(request)
	if err != nil {
		return "", "", fmt.Errorf("failed to query diploma:%v", err)
	}
	return string(response.Payload), string(response.TransactionID), nil
}
func (setup *FabricSetup) QueryRecord(address string) (string, string, error) {
	var args []string
	args = append(args, "QueryRecordByID")
	args = append(args, address)

	request := channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1])}}
	response, err := setup.client.Query(request)
	if err != nil {
		return "", "", fmt.Errorf("failed to query record:%v", err)
	}
	return string(response.Payload), string(response.TransactionID), nil
}
func (setup *FabricSetup) VerifyDiploma(address string) (string, string, error) {
	var args []string
	args = append(args, "VerifyDiplomaAddress")
	args = append(args, address)

	request := channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1])}}
	response, err := setup.client.Query(request)
	if err != nil {
		return "", "", fmt.Errorf("failed to query record:%v", err)
	}
	return string(response.Payload), string(response.TransactionID), nil
}
