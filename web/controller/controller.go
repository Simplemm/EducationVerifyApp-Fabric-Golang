package controller

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *Application) sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "上传成功")
}

//HomeHandler ...
func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	ShowView(w, r, "index.html", nil)
}
func (app *Application) SRegSchoolHander(w http.ResponseWriter, r *http.Request) {
	ShowView(w, r, "RegSchool.html", nil)
}
func (app *Application) SRegStudentHander(w http.ResponseWriter, r *http.Request) {
	ShowView(w, r, "RegStudent.html", nil)
}
func (app *Application) SRegCompanyHander(w http.ResponseWriter, r *http.Request) {
	ShowView(w, r, "RegCompany.html", nil)
}
func (app *Application) SAddDipHandler(w http.ResponseWriter, r *http.Request) {
	ShowView(w, r, "addDip.html", nil)
}
func (app *Application) RegStudentHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	schoolAddr := r.FormValue("address")
	fmt.Println(name, schoolAddr)
	//区块链交互代码
	studentString, transID, err := app.Fabric.RegisterStudent(name, schoolAddr)
	if err != nil {
		fmt.Fprintf(w, "注册失败")
	}
	student := new(Student)
	json.Unmarshal([]byte(studentString), student)

	data := &struct {
		TransactionID string
		Name          string
		Address       string
	}{
		TransactionID: transID[0:10],
		Name:          student.Name,
		Address:       student.Address,
	}
	ShowView(w, r, "regStudentResult.html", data)
	//返回注册成功界面
}

func (app *Application) RegSchoolHandler(w http.ResponseWriter, r *http.Request) {
	schoolName := r.FormValue("name")
	schoolLocation := r.FormValue("location")
	fmt.Println(schoolName, schoolLocation)

	schoolString, transID, err := app.Fabric.RegisterSchool(schoolName)
	if err != nil {
		fmt.Fprintf(w, "注册失败")
	}
	school := new(School)
	json.Unmarshal([]byte(schoolString), school)

	data := &struct {
		TransactionID string
		Name          string
		SchoolAddr    string
		PublicKey     rsa.PublicKey
		PrivateKey    string
	}{
		TransactionID: transID[0:10],
		Name:          school.Name,
		SchoolAddr:    school.Address,
		PublicKey:     school.PublicKey,
		PrivateKey:    school.PrivateKeyPem,
	}

	ShowView(w, r, "regSchoolResult.html", data)
	//返回注册成功界面
}

func (app *Application) RegCompanyHandler(w http.ResponseWriter, r *http.Request) {
	companyName := r.FormValue("name")
	companyLocation := r.FormValue("location")
	fmt.Println(companyName, companyLocation)

	companyString, transID, err := app.Fabric.RegisterCompany(companyName)
	if err != nil {
		fmt.Fprintf(w, "注册失败")
	}
	company := new(Company)
	json.Unmarshal([]byte(companyString), company)

	data := &struct {
		TransactionID string
		Name          string
		Address       string
	}{
		TransactionID: transID[0:10],
		Name:          company.Name,
		Address:       company.Address,
	}
	ShowView(w, r, "regCompanyResult.html", data)
	//返回注册成功界面
}

func (app *Application) AddDiplomaHandler(w http.ResponseWriter, r *http.Request) {
	schooladdr := r.FormValue("schooladdr")
	studentName := r.FormValue("studentname")
	schoolName := r.FormValue("schoolname")
	major := r.FormValue("major")
	startTime := r.FormValue("starttime")
	studentAddr := r.FormValue("studentaddr")
	gender := r.FormValue("gender")
	status := r.FormValue("status")
	diplomaType := r.FormValue("diplomatype")
	duration := r.FormValue("duration")
	fmt.Println(schooladdr, studentName, schoolName, major, startTime, studentAddr, gender, status, diplomaType, duration)

	diplomaString, transID, err := app.Fabric.UploadDiploma(schooladdr, studentAddr, startTime, duration, major, status, studentName, schoolName, diplomaType)
	if err != nil {
		fmt.Fprintf(w, "注册失败")
	}
	diploma := new(Diploma)
	json.Unmarshal([]byte(diplomaString), diploma)

	data := &struct {
		TransactionID string
		Name          string
		Address       string
	}{
		TransactionID: transID[0:10],
		Name:          diploma.StudentName, //学生姓名
		Address:       diploma.Address,
	}

	ShowView(w, r, "addDiplomaResult.html", data)
	//返回上传成功页面
}

func (app *Application) VerifyHandler(w http.ResponseWriter, r *http.Request) {

	//返回验证界面
}

func (app *Application) SQueryDiplomaHandler(w http.ResponseWriter, r *http.Request) {

	//返回查询学历界面
	ShowView(w, r, "quDiploma.html", nil)
}
func (app *Application) SQuerySchoolHandler(w http.ResponseWriter, r *http.Request) {
	//返回查询学校结果
	ShowView(w, r, "quSchool.html", nil)
}
func (app *Application) SQueryStudentHandler(w http.ResponseWriter, r *http.Request) {
	//返回查询学生结果

	ShowView(w, r, "quStudent.html", nil)
}
func (app *Application) SQueryCompanyHandler(w http.ResponseWriter, r *http.Request) {
	//返回查询学生的结果

	ShowView(w, r, "quCompany.html", nil)
}

func (app *Application) QuerySchoolHander(w http.ResponseWriter, r *http.Request) {
	fmt.Println("查询学校：", r.Form)
	schooladdr := r.FormValue("schooladdr")
	fmt.Println(schooladdr)

	schoolString, transID, err := app.Fabric.QuerySchool(schooladdr)
	if err != nil {
		fmt.Fprintf(w, "注册失败")
	}
	school := new(School)
	json.Unmarshal([]byte(schoolString), school)

	data := &struct {
		TransactionID string
		Name          string
		Address       string
		PublicKey     rsa.PublicKey
	}{
		TransactionID: transID[0:10],
		Name:          school.Name,
		Address:       school.Address,
		PublicKey:     school.PublicKey,
	}

	ShowView(w, r, "quSchoolResult.html", data)
	//返回查询结果界面
}
func (app *Application) QueryStudentHander(w http.ResponseWriter, r *http.Request) {
	fmt.Println("查询学生：", r.Form)
	studentAddr := r.FormValue("studentAddr")
	fmt.Println(studentAddr)
	//区块链交互代码
	studentString, transID, err := app.Fabric.QueryStudent(studentAddr)
	if err != nil {
		fmt.Fprintf(w, "查询失败")
	}
	student := new(Student)
	json.Unmarshal([]byte(studentString), student)

	var list []string
	list = append(list, "adasdfasfdgsafdgasf")
	data := &struct {
		TransactionID string
		Name          string
		Address       string
		DiplomaList   []string
	}{
		TransactionID: transID[0:10],
		Name:          student.Name,
		Address:       student.Address,
		DiplomaList:   student.DiplomaAddr,
	}
	ShowView(w, r, "quStudentResult.html", data)
	//返回查询结果界面
}
func (app *Application) QueryCompanyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("查询企业：", r.Form)
	companyAddr := r.FormValue("companyAddr")
	fmt.Println(companyAddr)

	companyString, transID, err := app.Fabric.QueryCompany(companyAddr)
	if err != nil {
		fmt.Fprintf(w, "注册失败")
	}
	company := new(Company)
	json.Unmarshal([]byte(companyString), company)

	data := &struct {
		TransactionID string
		Name          string
		Address       string
	}{
		TransactionID: transID[0:10],
		Name:          company.Name,
		Address:       company.Address,
	}
	ShowView(w, r, "quCompanyResult.html", data)
	//返回查询结果界面
}
func (app *Application) QueryDiplomaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("查询学历：", r.Form)
	diplomaAddr := r.FormValue("diplomaaddr")
	studentName := r.FormValue("name")

	diplomaString, transID, err := app.Fabric.QueryDiploma(diplomaAddr)
	if err != nil {
		fmt.Fprintf(w, "注册失败")
	}
	diploma := new(Diploma)
	json.Unmarshal([]byte(diplomaString), diploma)

	data := &struct {
		TransactionID  string
		StudentName    string
		SchoolName     string
		Major          string
		Duratiom       string
		StartTime      string
		Hash           [32]byte
		Sinature       string
		Status         string
		DiplomaType    string
		DiplomaAddress string
		SchoolAddr     string
		StudentAddr    string
	}{
		TransactionID:  transID[0:10],
		StudentName:    diploma.StudentName,
		SchoolName:     diploma.SchoolName,
		Major:          diploma.Major,
		Duratiom:       diploma.Duration,
		StartTime:      diploma.StartYear,
		Hash:           diploma.Hash,
		Status:         diploma.Status,
		DiplomaType:    diploma.DiplomaType,
		DiplomaAddress: diploma.Address,
		SchoolAddr:     diploma.SchoolAddress,
		StudentAddr:    diploma.StudentAddress,
		Sinature:       diploma.Signature,
	}
	fmt.Println(diplomaAddr, studentName)
	ShowView(w, r, "quDiplomaResult.html", data)
}
func (app *Application) SVerifyDiplomaHandler(w http.ResponseWriter, r *http.Request) {
	ShowView(w, r, "verify.html", nil)
}

func (app *Application) VerifyDiplomaHandler(w http.ResponseWriter, r *http.Request) {
	diplomaAddr := r.FormValue("diplomaAddr")

	verifyString, trasnID, err := app.Fabric.VerifyDiploma(diplomaAddr)
	if err != nil {
		fmt.Fprintf(w, "验证失败")
	}
	isDiploma0 := "无效"
	diploma := new(Diploma)
	school := new(School)
	if verifyString == "Yes" {
		isDiploma0 = "有效"
		diplomaString, _, _ := app.Fabric.QueryDiploma(diplomaAddr)
		json.Unmarshal([]byte(diplomaString), diploma)
		schoolString, _, _ := app.Fabric.QuerySchool(diploma.SchoolAddress)
		json.Unmarshal([]byte(schoolString), school)
	}
	data := &struct {
		IsDiploma     string
		TransactionID string
		SchoolName    string
		Sinature      string
		DiplomaAddr   string
		SchoolAddr    string
		StudentAddr   string
		PublicKey     rsa.PublicKey
	}{
		IsDiploma:     isDiploma0,
		TransactionID: trasnID[0:10],
		SchoolName:    diploma.SchoolName,
		SchoolAddr:    diploma.SchoolAddress,
		StudentAddr:   diploma.StudentAddress,
		DiplomaAddr:   diploma.Address,
		Sinature:      diploma.Signature,
		PublicKey:     school.PublicKey,
	}

	ShowView(w, r, "verifyResult.html", data)
}

func (app *Application) SQueryRecordHander(w http.ResponseWriter, r *http.Request) {
	ShowView(w, r, "quRecord.html", nil)
}

func (app *Application) QueryRecordHander(w http.ResponseWriter, r *http.Request) {
	diplomaAddr := r.FormValue("diplomaAddr")

	modifyString, transID, err := app.Fabric.QueryRecord(diplomaAddr)
	if err != nil {
		fmt.Fprintf(w, "查询失败")
	}
	record := new(ModifyRecord)
	json.Unmarshal([]byte(modifyString), record)

	data := &struct {
		TransactionID string
		SchoolAddr    string
		StudentAddr   string
		ModifyTime    []string
		DiplomaAddr   string
		Sinature      string
		Status        []string
	}{
		TransactionID: transID,
		DiplomaAddr:   record.DiplomaAddress,
		SchoolAddr:    record.SchoolAddress,
		StudentAddr:   record.StudentAddress,
		ModifyTime:    record.ModifyDate,
		Sinature:      record.Signature,
		Status:        record.Status,
	}
	ShowView(w, r, "quRecordResult.html", data)
}
