package updatewscomp

import "iqdev/ss/libs/globalObject"

type StudentUpdatePush struct {
	StudentStore   int
	StudentHistory int
	StudentBus     int
	StudentEvents  int
	StudentTest    int
	StudentMsg     int
	StudentSchool  int
}

func ResetUpdateLoopStudent() {}

func UpdateLoopStudent(studentProfile globalObject.StudentProfile) {
	getStudentInfoQ := `select * from student where student_id = ?;`

}
