package schoolhttpcomp

import (
	"io"
	genKey "iqdev/ss/libs/key"
	"iqdev/ss/libs/sql"
	managerwscomp "iqdev/ss/websocket/manager/managerWSComp"
	"net/http"
	"os"
	"os/exec"
)

type addManagerPostStruct struct {
	StudentFullName          string `json:"student_full_name"`
	StudentBirthDate         string `json:"student_birth_date"`
	StudentParentFullName    string `json:"student_parent_full_name"`
	StudentPhoneNumber       string `json:"student_phone_number"`
	StudentParentPhoneNumber string `json:"student_parent_phone_number"`
	StudentLocation          string `json:"student_location"`
	StudentClass             string `json:"student_class"`
	StudentStudyGroupId      int    `json:"student_study_group_id"`
	StudentIdBack            string `json:"student_id_back"`
	StudentIdFront           string `json:"student_id_front"`
	SchoolId                 int    `json:"school_id"`
	DisId                    int    `json:"dis_id"`
	ManagerId                int    `json:"manager_id"`
	BusId                    int    `json:"bus_id"`
}

func AddStudent() {
	http.HandleFunc("POST /http/school/addStudent", func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("Expect") == "100-continue" {
			w.WriteHeader(http.StatusContinue)
		}
		r.ParseMultipartForm(100000000)

		r.ParseForm()

		file, header, err := r.FormFile("disImage")
		if err != nil {
			http.Error(w, "faild to get upload file ", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		destenation := "./images/dis/"
		key := genKey.RandomKey(20)
		cfile, cfileErr := os.OpenFile(
			key+header.Filename,
			os.O_WRONLY|os.O_CREATE,
			0666,
		)
		if cfileErr != nil {
			http.Error(w, "faild to read upload file ", http.StatusInternalServerError)
			return
		}
		_, cpyErr := io.Copy(cfile, file)
		if cpyErr != nil {
			http.Error(w, "faild to read upload file ", http.StatusInternalServerError)
			return
		}
		exec.Command("mv", cfile.Name(), destenation+cfile.Name())

		ManagerKey := r.FormValue("manKey")
		DisFullName := r.FormValue("dis_full_name")
		DisBirthday := r.FormValue("dis_birthday")
		DisCurrentLocation := r.FormValue("dis_current_location")
		DisHomeLocation := r.FormValue("dis_home_location")
		DisPhoneType := r.FormValue("dis_phone_type")
		DisPhoneXnumber := r.FormValue("dis_phone_xnumber")
		//DisPhoneIsEmulated := r.FormValue("dis_phone_is_emulated")
		DisPhoneBattaryLevel := r.FormValue("dis_phone_battary_level")
		DisTotalMemory := r.FormValue("dis_total_memory")
		DisUsedMemory := r.FormValue("dis_used_memory")
		DisPhoneCapacity := r.FormValue("dis_phone_capacirty")
		DisPhoneDiskFree := r.FormValue("dis_phone_disk_free")
		DisTotalImages := r.FormValue("dis_total_images")
		DisTotalVideos := r.FormValue("dis_total_videos")
		//DisActive := r.FormValue("dis_active")
		DisIdNumber := r.FormValue("dis_id_number")

		managerInfo := managerwscomp.GetManagerProfile(ManagerKey, nil)

		if managerInfo.ManagerExsist == false {
			http.Error(w, "faild to read upload file ", http.StatusInternalServerError)
			return
		}

		addDisSqlQuery := `INSERT INTO dis(
      dis_full_name,
      dis_birthday,
      dis_image_path,
      dis_current_location,
      dis_home_location,
      dis_phone_type,
      dis_phone_xnumber,
      dis_phone_is_emulated,
      dis_phone_battary_level,
      dis_total_memory,
      dis_used_memory,
      dis_key,
      dis_phone_capacity,
      dis_phone_disk_free,
      dis_total_images,
      dis_total_videos,
      manager_id,
      dis_register_date,
      dis_active,
      dis_id_number
    ) values (
      ?,
      ?,
      ?,
      ?,
      ?,
      ?,
      ?,
      ?,
      ?,
      ?,
      ?,
      ?,
      ?,
      ?,
      ?,
      ?,
      ?,
      ?,
      ?,
      ?);`

		sqlConn := sql.InitConnection().Connections

		_, sqlErr := sqlConn.Exec(addDisSqlQuery,
			DisFullName,
			DisBirthday,
			destenation+cfile.Name(),
			DisCurrentLocation,
			DisHomeLocation,
			DisPhoneType,
			DisPhoneXnumber,
			true,
			DisPhoneBattaryLevel,
			DisTotalMemory,
			DisUsedMemory,
			genKey.RandomKey(120),
			DisPhoneCapacity,
			DisPhoneDiskFree,
			DisTotalImages,
			DisTotalVideos,
			managerInfo.ManagerInfo.ManagerId,
			"2024",
			true,
			DisIdNumber)

		if sqlErr != nil {
			http.Error(w, "faild to insert dis into database", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("file uploaded yay"))
	})

}
