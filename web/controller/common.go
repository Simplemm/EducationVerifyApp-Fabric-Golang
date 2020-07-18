package controller

import "crypto/rsa"

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
