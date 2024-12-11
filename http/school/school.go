package httpSchool

import (
	schoolhttpcomp "iqdev/ss/http/school/schoolComp"
	"net/http"
)

func HttpSchoolHanlder() {
	http.HandleFunc("/http/school/schoolCheck", schoolhttpcomp.CheckSchoolExsit)
}
