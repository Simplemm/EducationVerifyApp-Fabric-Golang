package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"math/big"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// SchoolList ...
type SchoolList struct {
	SchoolAddressList []string `json:"school_address_list"`
}

//Student ...
type Student struct {
	Address     string   `json:"address"`
	Name        string   `json:"name"`
	DiplomaAddr []string `json:"educationBackground"`
}

// Diploma ...
type Diploma struct {
	Address        string   `json:"address"` //学历地址
	SchoolAddress  string   `json:"school_address"`
	StudentAddress string   `json:"student_address"`
	StudentName    string   `json:"student_name"`
	SchoolName     string   `json:"school_name"`
	StartYear      string   `json:"start_year"`
	Duration       string   `json:"duration"`
	DiplomaType    string   `json:"diploma_type"`
	Major          string   `json:"major"`
	Status         string   `json:"status"`    // 0 for 退学，1 for 在读，2 for 毕业
	Hash           [32]byte `json:"hash"`      //将结构体json 序列化后，直接sha265 hash的结果
	Signature      string   `json:"signature"` //用学校的私钥对hash进行签名
}

//School ...
type School struct {
	Name                string        `json:"name"`
	Address             string        `json:"address"`
	PublicKey           rsa.PublicKey `json:"public_key"`
	PrivateKeyPem       string        `json:"private_key"`
	StudentAddressArray []string      `json:"student_address"`
	//Description string `json:"description"`
}

//Company ...
type Company struct {
	Address string `json:"address"`
	Name    string `json:"name"`
	//CompanyType int    `json:"company_type"`
	//Description string `json:"description"`
}

//ModifyRecord ...
type ModifyRecord struct {
	DiplomaAddress string `json:"eudback_address"`
	//ChangedData    string    `json:"changed_data"`
	SchoolAddress  string   `json:"school_address"` //学校地址
	StudentAddress string   `json:"student_address"`
	Status         []string `json:"status"` //0 代表入学，1代表正常毕业，2代表退学
	ModifyDate     []string `json:"modify_date"`
	Signature      string   `json:"signature"`
	//Description []string `json:"description"`
}

// ConstructKey ...
func ConstructKey(key string, prefix string) string {
	return fmt.Sprintf("%s_%s", prefix, key)
}

// Init ...
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("edu app init.")
	return shim.Success(nil)
}

// Invoke ...
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("funtion Invoke")
	function, args := stub.GetFunctionAndParameters()
	fmt.Println(function, args)
	if function == "RegisterSchool" {
		return t.RegisterSchool(stub, args)
	} else if function == "RegisterStudent" {
		return t.RegisterStudent(stub, args)
	} else if function == "RegisterCompany" {
		return t.RegisterCompany(stub, args)
	} else if function == "UploadDiploma" {
		return t.UploadDiploma(stub, args)
	} else if function == "QuerySchoolByAddress" {
		return t.QuerySchoolByAddress(stub, args)
	} else if function == "QueryStudentByAddress" {
		return t.QueryStudentByAddress(stub, args)
	} else if function == "QueryRecordByID" {
		return t.QueryRecordByID(stub, args)
	} else if function == "QueryCompanyByAddress" {
		return t.QueryCompanyByAddress(stub, args)
	} else if function == "QueryDiplomaByAddress" {
		return t.QueryDiplomaByAddress(stub, args)
	} else if function == "VerifyDiplomaAddress" {
		return t.VerifyDiplomaAddress(stub, args)
	} else {
		return shim.Error("function name is incorrect to call")
	}
	//return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")

}

// QuerySchoolByAddress ...
func (t *SimpleChaincode) QuerySchoolByAddress(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("test school")
	if len(args) != 1 {
		return shim.Error("Incorrect number of args,expecting 1!")
	}
	address := args[0]
	schoolbytes, err := stub.GetState(ConstructKey(address, "school"))
	if err != nil {
		return shim.Error("GetState defeated!")
	}
	// school := new(School)
	// json.Unmarshal(schoolbytes, school)
	fmt.Println("query school end")
	return shim.Success(schoolbytes)
}

// RegisterSchool ...
func (t *SimpleChaincode) RegisterSchool(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//提供一个参数：schoolname
	if len(args) != 1 {
		return shim.Error("Incorrect number of args,expecting 1")
	}
	if args[0] == "" {
		return shim.Error("Args[0] for school name is null")
	}
	name := args[0]
	fmt.Println("Register school ：" + name)
	//产生RSA公私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := privateKey.PublicKey
	pemString := EncodeRsaPriKeyToPem(privateKey)
	if err != nil {
		return shim.Error("Generate privatekey error")
	}
	//通过公钥地址产生学校地址
	hash := sha256.Sum256(publicKey.N.Bytes())
	address := base64.StdEncoding.EncodeToString(hash[:])
	//赋值，序列化，存入世界状态数据库
	school := &School{
		Name:          name,
		Address:       address,
		PrivateKeyPem: pemString,
		PublicKey:     publicKey,
	}
	schoolBytes, err := json.Marshal(school)
	err2 := stub.PutState(ConstructKey(address, "school"), schoolBytes)
	if err2 != nil {
		return shim.Error("Putstate school address defeated!")
	}
	fmt.Println("register school finished：", string(schoolBytes))
	//返回结果信息，学校保存相关信息
	return shim.Success(schoolBytes)
}

//RegisterStudent ...
func (t *SimpleChaincode) RegisterStudent(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//provide two args,arg[0] is name，arg[1] is school address
	if len(args) != 2 {
		return shim.Error("Incorrect number of args,expecting 2.")
	}
	if args[0] == "" || args[1] == "" {
		return shim.Error("args must be a value,not a null string.")
	}
	name := args[0]
	schooladdress := args[1]
	fmt.Println("register student :" + name)
	//generate big random number and  base64encode the hash value
	max := new(big.Int).Lsh(big.NewInt(1), 256)
	serialNumber, _ := rand.Int(rand.Reader, max)
	randByte := serialNumber.Bytes()
	hash := sha256.Sum256(randByte)
	address := base64.StdEncoding.EncodeToString(hash[:])
	student := &Student{
		Name:    name,
		Address: address,
	}
	stuByte, err := json.Marshal(student)
	if err != nil {
		return shim.Error("json marshal defeated.")
	}
	err2 := stub.PutState(ConstructKey(address, "student"), stuByte)
	if err2 != nil {
		return shim.Error("putstate student with address defeated.")
	}
	//学校中添加学生账号,先根据学校的地址查询学校的信息，添加完成后再保存
	schBytes, err := stub.GetState(ConstructKey(schooladdress, "school"))
	if err != nil {
		return shim.Error("getstate with school id defeated.")
	}
	school := new(School)
	err5 := json.Unmarshal(schBytes, school)
	if err5 != nil {
		return shim.Error("json unmarshal the school bytes defeat:" + err5.Error())
	}
	school.StudentAddressArray = append(school.StudentAddressArray, address)
	schBytes2, err7 := json.Marshal(school)
	if err7 != nil {
		return shim.Error(err7.Error())
	}
	err8 := stub.PutState(ConstructKey(schooladdress, "school"), schBytes2)
	if err8 != nil {
		return shim.Error("update school’s student list error")
	}
	fmt.Println("register student finished:", string(schBytes2))
	fmt.Println(string(stuByte))
	return shim.Success(stuByte)
}

//QueryStudentByAddress ...
func (t *SimpleChaincode) QueryStudentByAddress(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//提供一个参数，学生的账号地址
	if len(args) != 1 {
		return shim.Error("Incorrect number of args,expecting 1")
	}
	if args[0] == "" {
		return shim.Error("args[0] must be a valid value,not a empty string.")
	}
	address := args[0]
	studentBytes, err := stub.GetState(ConstructKey(address, "student"))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(studentBytes)
}

//UploadDiploma ...
func (t *SimpleChaincode) UploadDiploma(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//提供9个参数，学校地址，学生地址，开始年份，学制，专业，状态（入学，退学，毕业),学生名，学校名，学历类型
	if len(args) != 9 {
		return shim.Error("Incorrect number of args,expecting 8.")
	}
	for i := 0; i < len(args); i++ {
		if args[i] == "" {
			return shim.Error("args" + string(i) + "is empty string.")
		}
	}
	fmt.Println("upload start...")
	//产生地址
	max := new(big.Int).Lsh(big.NewInt(1), 256)
	serialNumber, _ := rand.Int(rand.Reader, max)
	randByte := serialNumber.Bytes()
	hash := sha256.Sum256(randByte)
	address := base64.StdEncoding.EncodeToString(hash[:])
	//赋值参数
	eduAddr := address
	schoolAddr := args[0]
	studentAddr := args[1]
	startYear := args[2]
	duration := args[3]
	major := args[4]
	status := args[5]
	studentName := args[6]
	schoolName := args[7]
	diplomaType := args[8]

	diploma := &Diploma{
		Address:        eduAddr,
		StudentAddress: studentAddr,
		SchoolAddress:  schoolAddr,
		StartYear:      startYear,
		Duration:       duration,
		Major:          major,
		Status:         status,
		StudentName:    studentName,
		SchoolName:     schoolName,
		DiplomaType:    diplomaType,
	}
	//序列化用于hash
	diplomaByte1, _ := json.Marshal(diploma)
	diplomaHash := sha256.Sum256(diplomaByte1)
	diploma.Hash = diplomaHash
	schoolBytes, err := stub.GetState(ConstructKey(schoolAddr, "school"))
	if err != nil {
		return shim.Error("Getstate school with address wrong.")
	}
	school := new(School)
	err = json.Unmarshal(schoolBytes, school)
	if err != nil {
		return shim.Error("unmarshal schoolbytes wrong.")
	}
	//学校签名
	//schoolPriKeyPem := school.PrivateKey
	//schoolPriKey := DecodePemToPriKey(schoolPriKeyPem)
	schoolPriKeyPem := school.PrivateKeyPem
	schoolPriKey := DecodePemToRsaPriKey(schoolPriKeyPem)

	signature, err1 := rsa.SignPKCS1v15(rand.Reader, schoolPriKey, crypto.SHA256, diplomaHash[:])
	if err1 != nil {
		return shim.Error(err1.Error())
	}
	diploma.Signature = hex.EncodeToString(signature)
	diplomaBytes2, err2 := json.Marshal(diploma)
	if err2 != nil {
		return shim.Error("diploma marshal error.")
	}
	//将学历信息存入世界状态数据库
	err = stub.PutState(ConstructKey(diploma.Address, "diploma"), diplomaBytes2)
	fmt.Println(string(diplomaBytes2))
	if err != nil {
		return shim.Error("putstate diploma error")
	}
	//为学生添加学历
	studentBytes, err4 := stub.GetState(ConstructKey(studentAddr, "student"))
	if err4 != nil {
		return shim.Error("getstate student error.")
	}
	student := new(Student)
	err = json.Unmarshal(studentBytes, student)
	if err != nil {
		return shim.Error("json unmarshal student bytes error")
	}
	student.DiplomaAddr = append(student.DiplomaAddr, diploma.Address)
	studentBytes2, err := json.Marshal(student)
	if err != nil {
		return shim.Error("json marshal student bytes error.")
	}
	stub.PutState(ConstructKey(studentAddr, "student"), studentBytes2)
	//添加修改记录,先查询record是否存在,不存在就初始化.
	modifyDate := time.Now().String()
	//var record *ModifyRecord
	// if recordBytes, err := stub.GetState(ConstructKey(address, "record")); err != nil {
	// record = &ModifyRecord{
	// 	DiplomaAddress: address,
	// 	SchoolAddress:  schoolAddr,
	// 	StudentAddress: studentAddr,
	// 	}
	// } else {
	// 	json.Unmarshal(recordBytes, *record)
	// }
	record := &ModifyRecord{
		DiplomaAddress: address,
		SchoolAddress:  schoolAddr,
		StudentAddress: studentAddr,
	}
	record.ModifyDate = append(record.ModifyDate, modifyDate) //append日期
	record.Status = append(record.Status, status)             //append状态
	recordBytes, _ := json.Marshal(record)
	recordHash := sha256.Sum256(recordBytes)
	// 签名
	recordSign, err11 := rsa.SignPKCS1v15(rand.Reader, schoolPriKey, crypto.SHA256, recordHash[:])
	if err11 != nil {
		return shim.Error("sign recoed error" + err11.Error())
	}
	record.Signature = hex.EncodeToString(recordSign)
	recordsBytes2, _ := json.Marshal(record)
	if err := stub.PutState(ConstructKey(address, "record"), recordsBytes2); err != nil {
		return shim.Error(fmt.Sprintf("putstate record error :%s", err.Error()))
	}
	fmt.Println("upload diploma finished:", diplomaBytes2)
	return shim.Success(diplomaBytes2)
}

//QueryDiplomaByAddress
func (t *SimpleChaincode) QueryDiplomaByAddress(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error(fmt.Sprintf("Incorrect number of args,expecting 1."))
	}
	address := args[0]
	DiplomaByte, err := stub.GetState(ConstructKey(address, "diploma"))
	if err != nil {
		return shim.Error(fmt.Sprintf("getstate diploma error:%s", err.Error()))
	}
	return shim.Success(DiplomaByte)
}

//QueryRecordByID ... 通过学历的地址查询学历的修改纪律
func (t *SimpleChaincode) QueryRecordByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error(fmt.Sprintf("Incorrect number of args,expecting 1..."))
	}
	id := args[0]
	recordBytes, err := stub.GetState(ConstructKey(id, "record"))
	if err != nil {
		return shim.Error(fmt.Sprintf("getstate company error:%s", err.Error()))
	}
	return shim.Success(recordBytes)
}

//ChangediplomaStatus 将入学状态变成毕业状态，退学状态；或将退学状态变成入学状态，毕业状态
func (t *SimpleChaincode) ChangediplomaStatus(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//提供参数:学校公钥，学历id，改变后的状态status

	return shim.Success(nil)
}

//RegisterCompany ...
func (t *SimpleChaincode) RegisterCompany(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of args,Expecting 1.")
	}
	name := args[0]
	fmt.Println("register company :" + name)
	max := new(big.Int).Lsh(big.NewInt(1), 256)
	serialNumber, _ := rand.Int(rand.Reader, max)
	randByte := serialNumber.Bytes()
	hash := sha256.Sum256(randByte)
	address := base64.StdEncoding.EncodeToString(hash[:])
	company := &Company{
		Name:    name,
		Address: address,
	}
	companyBytes, err := json.Marshal(company)
	if err != nil {
		return shim.Error("json marshal error")
	}
	err2 := stub.PutState(ConstructKey(address, "company"), companyBytes)
	if err2 != nil {
		return shim.Error("putstate company error")
	}
	fmt.Println("register company success.:", string(companyBytes))
	return shim.Success(companyBytes)
}

//QueryCompanyByAddress ...
func (t *SimpleChaincode) QueryCompanyByAddress(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error(fmt.Sprintf("Incorrect number of args,expecting 1."))
	}
	address := args[0]
	companyBytes, err := stub.GetState(ConstructKey(address, "company"))
	if err != nil {
		return shim.Error(fmt.Sprintf("getstate company error:%s", err.Error()))
	}
	return shim.Success(companyBytes)
}

//VerifyDiplomaAddress ...
func (t *SimpleChaincode) VerifyDiplomaAddress(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//提供两个参数，公司地址，学历证书
	if len(args) != 1 {
		return shim.Error("Incorrect number of args,Expecting 2.")
	}
	if args[0] == "" {
		return shim.Error("args is invalid,please check the value.")
	}
	diplomaAddr := args[0]
	//验证学历地址是否合法
	diplomaBytes, err := stub.GetState(ConstructKey(diplomaAddr, "diploma"))
	if err != nil {
		return shim.Error("diploma address is nil ")
	}

	diploma := new(Diploma)
	err1 := json.Unmarshal(diplomaBytes, diploma)
	if err1 != nil {
		return shim.Error("json unmarshal error.")
	}
	//验证签名，哈希
	hash := diploma.Hash
	signature := diploma.Signature
	schoolAddr := diploma.SchoolAddress
	var null [32]byte
	diploma.Hash = null
	diploma.Signature = ""
	diplomaText, _ := json.Marshal(diploma)
	hash2 := sha256.Sum256(diplomaText)
	for i := 0; i < len(hash); i++ {
		if hash[i] != hash2[i] {
			return shim.Error("hash is wrong value")
		}
	}
	schoolBytes, _ := stub.GetState(ConstructKey(schoolAddr, "school"))
	school := new(School)
	json.Unmarshal(schoolBytes, school)
	publicKey := school.PublicKey
	signBytes, _ := hex.DecodeString(signature)
	err3 := rsa.VerifyPKCS1v15(&publicKey, crypto.SHA256, hash[:], signBytes)
	if err != nil {
		return shim.Error("diploma verify error:" + err3.Error())
	}
	fmt.Println("verify end:yes")
	return shim.Success([]byte("Yes"))
}

//EncodeRsaPriKeyToPem ...
func EncodeRsaPriKeyToPem(priKey *rsa.PrivateKey) string {
	ecder := x509.MarshalPKCS1PrivateKey(priKey)
	b := new(pem.Block)
	b.Bytes = ecder
	bytes := pem.EncodeToMemory(b)
	return string(bytes)
}

//DecodePemToRsaPriKey ...
func DecodePemToRsaPriKey(pemString string) *rsa.PrivateKey {
	bytes := []byte(pemString)
	b, _ := pem.Decode(bytes)
	priKey, _ := x509.ParsePKCS1PrivateKey(b.Bytes)
	return priKey
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
