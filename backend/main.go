package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/team07/app/controllers"
	"github.com/team07/app/ent"
	"github.com/team07/app/ent/jobposition"
)

type Users struct {
	User []User
}

type User struct {
	name          string
	email         string
	password      string
	jobpositionID int
}

type JobPositions struct {
	JobPosition []JobPosition
}

type JobPosition struct {
	name string
}

type InspectionResults struct {
	InspectionResult []InspectionResult
}

type InspectionResult struct {
	resultname    string
	jobpositionID int
}

type Purposes struct {
	Purpose []Purpose
}

type Purpose struct {
	objective string
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewUserController(v1, client)
	controllers.NewInspectionResultController(v1, client)

	//ลงข้อมูล User
	jobpositions := []string{"เจ้าหน้าที่ตรวจสภาพรถ", "เจ้าหน้าที่รถพยาบาล", "เจ้าหน้าที่โอเปอร์เรเตอร์", "เจ้าหน้าที่ซ่อมบำรุงรถ"}
	for _, jp := range jobpositions {
		client.JobPosition.
			Create().
			SetPositionName(jp).
			Save(context.Background())
	}

	users := Users{
		User: []User{
			User{"พี่โต", "to@email.com", "1234", 1},
			User{"พี่หลาม", "lam@email.com", "1234", 1},
			User{"อัลฟ่าน", "alfan@email.com", "1234", 2},
			User{"ใจเกเร", "jaikere@email.com", "1234", 3},
			User{"แดงกีต้าร์", "dang@email.com", "1234", 3},
			User{"อินเดีย", "india@email.com", "1234", 4},
		},
	}

	for _, u := range users.User {
		jp, err := client.JobPosition.
			Query().
			Where(jobposition.IDEQ(int(u.jobpositionID))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		client.User.
			Create().
			SetName(u.name).
			SetEmail(u.email).
			SetPassword(u.password).
			SetJobposition(jp).
			Save(context.Background())
	}

	inspectionresults := InspectionResults{
		InspectionResult: []InspectionResult{
			InspectionResult{"พร้อมปฏิบัติหน้าที่", 1},
			InspectionResult{"ส่งซ่อมแซม", 1},
			InspectionResult{"ไม่ได้รับการตรวจสภาพ", 1},
			InspectionResult{"รอส่งตรวจสภาพรถ", 2},
			InspectionResult{"ส่งตรวจสภาพรถ", 2},
			InspectionResult{"ปลดประจำการ", 2},
			InspectionResult{"พร้อมใช้งาน", 2},
		},
	}

	for _, ir := range inspectionresults.InspectionResult {
		jp, err := client.JobPosition.
			Query().
			Where(jobposition.IDEQ(int(ir.jobpositionID))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		client.InspectionResult.
			Create().
			SetResultName(ir.resultname).
			SetJobposition(jp).
			Save(context.Background())
	}

	// Set Purposes Data
	purposes := Purposes{
		Purpose: []Purpose{
			Purpose{"รับแจ้งอุบัติเหตุ"},
			Purpose{"ขนส่งย้ายผู้ป่วย"},
			Purpose{"อื่นๆ"},
		},
	}

	for _, pp := range purposes.Purpose {
		client.Purpose.
			Create().
			SetStatus(st.Purpose).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
