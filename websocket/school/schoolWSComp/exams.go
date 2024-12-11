package schoolwscomp

import (
	libErrors "iqdev/ss/libs/errors"
	"iqdev/ss/libs/globalObject"
	"iqdev/ss/libs/sql"

	"github.com/gorilla/websocket"
)

func GetAllExamsInfo(
	schoolId int,
	conn *websocket.Conn) []globalObject.StudentExams {

	studentExams := []globalObject.StudentExams{}
	studentExamstmp := globalObject.StudentExams{}

	// sql connection

	// query
	getManagerQuery := "SELECT * FROM student_exams where school_id = ?;"
	// sql connection
	connObj := sql.InitConnection()
	// query manager info
	rows, sqlError := connObj.Connections.Query(getManagerQuery, schoolId)

	if sqlError != nil {
		libErrors.ReturnError(conn)
	}

	defer rows.Close()
	defer connObj.Connections.Close()

	for rows.Next() {
		err := rows.Scan(
			&studentExamstmp.StudentExamId,
			&studentExamstmp.StudentName,
			&studentExamstmp.StudentMake,
			&studentExamstmp.StudentExamMaterial,
			&studentExamstmp.TeacherName,
			&studentExamstmp.TeacherId,
			&studentExamstmp.StudentId,
			&studentExamstmp.SchoolId,
			&studentExamstmp.Studenthistorydate,
			&studentExamstmp.StudyGroupId)
		if err != nil {
			libErrors.ReturnError(conn)
		}

		studentExams = append(studentExams, globalObject.StudentExams{
			studentExamstmp.StudentExamId,
			studentExamstmp.StudentName,
			studentExamstmp.StudentMake,
			studentExamstmp.StudentExamMaterial,
			studentExamstmp.TeacherName,
			studentExamstmp.TeacherId,
			studentExamstmp.StudentId,
			studentExamstmp.SchoolId,
			studentExamstmp.Studenthistorydate,
			studentExamstmp.StudyGroupId})

	}
	return studentExams
}
