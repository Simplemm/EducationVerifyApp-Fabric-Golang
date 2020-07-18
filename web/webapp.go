package webapp

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chainHero/heroes-service/web/controller"
)

func Webapp(app *controller.Application) {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	//http.HandleFunc("/post", sayhelloName) //设置访问的路由
	http.HandleFunc("/", app.HomeHandler)
	http.HandleFunc("/index", app.HomeHandler)
	http.HandleFunc("/registerSchool", app.SRegSchoolHander)
	http.HandleFunc("/registerCompany", app.SRegCompanyHander)
	http.HandleFunc("/registerStudent", app.SRegStudentHander)

	http.HandleFunc("/uplaodDiploma", app.SAddDipHandler)

	http.HandleFunc("/queryDiplomaPage", app.SQueryDiplomaHandler)
	http.HandleFunc("/queryStudentPage", app.SQueryStudentHandler)
	http.HandleFunc("/querySchoolPage", app.SQuerySchoolHandler)
	http.HandleFunc("/queryCompanyPage", app.SQueryCompanyHandler)

	http.HandleFunc("/regSchoolResult", app.RegSchoolHandler)
	http.HandleFunc("/regStudentResult", app.RegStudentHandler)
	http.HandleFunc("/regCompanyResult", app.RegCompanyHandler)

	http.HandleFunc("/addDiplomaResult", app.AddDiplomaHandler)

	http.HandleFunc("/querySchoolResult", app.QuerySchoolHander)
	http.HandleFunc("/queryStudentResult", app.QueryStudentHander)
	http.HandleFunc("/queryCompanyResult", app.QueryCompanyHandler)

	http.HandleFunc("/queryDiplomaResult", app.QueryDiplomaHandler)

	http.HandleFunc("/verifyPage", app.SVerifyDiplomaHandler)
	http.HandleFunc("/verifyResult", app.VerifyDiplomaHandler)

	http.HandleFunc("/queryRecordPage", app.SQueryRecordHander)
	http.HandleFunc("/queryRecordResult", app.QueryRecordHander)

	fmt.Println("监听开始...")
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// func main() {
// 	app := new(controller.Application)
// 	webapp(app)
// }
