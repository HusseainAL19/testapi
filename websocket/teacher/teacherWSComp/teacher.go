package teacherwscomp

import (
	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/globalObject"
	"iqdev/ss/libs/sql"

	"github.com/gorilla/websocket"
)

type TeacherInfoStruct struct {
	TeacherExsist bool
	TeacherInfo   globalObject.TeacherProfile
}

func GetTeacherProfile(
	teacherKey string,
	conn *websocket.Conn) TeacherInfoStruct {

	teacherProfiles := TeacherInfoStruct{}

	// sql connection

	// query
	getManagerQuery := "SELECT * FROM teacher WHERE teacher_key = ?;"
	// sql connection
	connObj := sql.InitConnection()
	// query manager info
	sqlError := connObj.Connections.QueryRow(getManagerQuery, teacherKey).Scan(
		&teacherProfiles.TeacherInfo.TeacherId,
		&teacherProfiles.TeacherInfo.TeacherFullName,
		&teacherProfiles.TeacherInfo.TeacherBirthDay,
		&teacherProfiles.TeacherInfo.TeacherLocation,
		&teacherProfiles.TeacherInfo.TeacherDeviceLocation,
		&teacherProfiles.TeacherInfo.TeacherTotalStudent,
		&teacherProfiles.TeacherInfo.TeacherTotalScore,
		&teacherProfiles.TeacherInfo.TeacherTotalPresent,
		&teacherProfiles.TeacherInfo.TeacherImagePath,
		&teacherProfiles.TeacherInfo.TeacherIdXNumber,
		&teacherProfiles.TeacherInfo.TeacherIdImagePathFront,
		&teacherProfiles.TeacherInfo.TeacherIdImagePathBack,
		&teacherProfiles.TeacherInfo.TeacherDegree,
		&teacherProfiles.TeacherInfo.TeacherMajor,
		&teacherProfiles.TeacherInfo.TeacherKey,
		&teacherProfiles.TeacherInfo.TeacherTotalGroups,
		&teacherProfiles.TeacherInfo.SchoolId,
		&teacherProfiles.TeacherInfo.SchoolId,
		&teacherProfiles.TeacherInfo.OwnerId,
		&teacherProfiles.TeacherInfo.TeacherActive)

	if sqlError != nil {
		teacherProfiles.TeacherExsist = false
		libErrors.ReturnError(conn)
	}

	//defer connObj.Connections.Close()

	if sqlError != nil {
		teacherProfiles.TeacherExsist = false
		libErrors.ReturnError(conn)
	}

	teacherProfiles.TeacherExsist = true
	return teacherProfiles
}
