package teacherwscomp

import (
	"encoding/json"
	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/globalObject"

	"github.com/gorilla/websocket"
)

type AllSchoolOwnerInfoObj struct {
	TeacherProfiles    globalObject.TeacherProfile
	TeacherStudyGroups globalObject.TeacherProfile
	TeacherChatGroups  globalObject.TeacherProfile
	TeacherExams       globalObject.TeacherProfile
	TeacherScore       globalObject.TeacherProfile
}

type ConnectionsList struct {
	GodConnectionNum   int
	GodConnectionsList []*websocket.Conn
}

type connectionsProfile struct {
}

func assembleInfo(teacherProfile globalObject.TeacherProfile,
	connection *websocket.Conn) AllSchoolOwnerInfoObj {
	allSchoolOwnerInfo := AllSchoolOwnerInfoObj{}

	if connection == nil {
		return allSchoolOwnerInfo
	}

	return allSchoolOwnerInfo
}

func GetAllTeacherInfo(teacherProfile globalObject.TeacherProfile,
	connection *websocket.Conn) {

	var allInfoRow AllSchoolOwnerInfoObj
	allInfoRow = assembleInfo(teacherProfile, connection)

	allInfoJson, ctjsonError := json.Marshal(allInfoRow)
	if ctjsonError != nil {
		libErrors.ReturnError(connection)
	}

	writeError := connection.WriteMessage(1, allInfoJson)
	if writeError != nil {
		libErrors.ReturnError(connection)
	}
}
