package globalObject

type GodProfile struct {
	GodId            int    `json:"god_id"`
	GodName          string `json:"god_name"`
	GodKey           string `json:"god_key"`
	GodTotalSchools  int    `json:"god_total_schools"`
	GodTotalStudents int    `json:"god_total_students"`
	GodTotalTeachers int    `json:"god_total_teachers"`
	GodTotalBus      int    `json:"god_total_bus"`
	GodLastActive    string `json:"last_active"`
}

type ManagersProfile struct {
	ManagerId                 int    `json:"manager_id"`
	ManagerName               string `json:"manager_name"`
	ManagerBirthDay           string `json:"manager_birth_day"`
	ManagerCurrentLocation    string `json:"manager_current_location"`
	ManagerImagePath          string `json:"manager_image_path"`
	ManagerDeviceLocation     string `json:"manager_device_location"`
	ManagerDeviceType         string `json:"manager_device_type"`
	ManagerDeviceXNumber      string `json:"manager_device_xnumber"`
	ManagerDeviceEmulated     string `json:"manager_device_emulated"`
	ManagerDeviceBattaryLevel string `json:"manager_device_battary_level"`
	ManagerTotalMemory        string `json:"manager_total_memory"`
	ManagerUsedMemory         string `json:"manager_used_memory"`
	ManagerKey                string `json:"manager_key"`
	ManagerDeviceCapacity     string `json:"manager_device_capacity"`
	ManagerDeviceFreeDisk     string `json:"manager_device_free_disk"`
	ManagerDeviceTotalImages  string `json:"manager_total_images"`
	ManagerDeviceTotalVideos  string `json:"manager_total_videos"`
	ManagerRegisterDate       string `json:"manager_register_date"`
	ManagerActive             string `json:"manager_active"`
	GodId                     int    `json:"god_id"`
	ManagerLastActivity       string `json:"manager_last_activity"`
	ManagerPhoneNumber        string `json:"manager_phone_number"`
}

type DisProfile struct {
	DisId                int    `json:"dis_id"`
	DisFullName          string `json:"dis_full_name"`
	DisBirthday          string `json:"dis_birthday"`
	DisImagePath         string `json:"dis_image_path"`
	DisCurrentLocation   string `json:"dis_current_location"`
	DisHomeLocation      string `json:"dis_home_location"`
	DisPhoneType         string `json:"dis_phone_type"`
	DisPhoneXnumber      string `json:"dis_phone_xnumber"`
	DisPhoneIsEmulated   bool   `json:"dis_phone_is_emulated"`
	DisPhoneBattaryLevel string `json:"dis_phone_battary_level"`
	DisTotalMemory       string `json:"dis_total_memory"`
	DisUsedMemory        string `json:"dis_used_memory"`
	DisKey               string `json:"dis_key"`
	DisPhoneCapacity     string `json:"dis_phone_capacirty"`
	DisPhoneDiskFree     string `json:"dis_phone_disk_free"`
	DisTotalImages       int    `json:"dis_total_images"`
	DisTotalVideos       int    `json:"dis_total_videos"`
	ManagerId            int    `json:"manager_id"`
	DisRegisterDate      string `json:"dis_register_date"`
	DisActive            int    `json:"dis_active"`
	DisIdNumber          string `json:"dis_id_number"`
}

type SchoolOwnerProfile struct {
	SchoolOwnerId           int    `json:"school_owner_id"`
	SchoolOwnerName         string `json:"school_owner_name"`
	SchoolOwnerLocation     string `json:"school_owner_location"`
	SchoolOwnerTotalSchools int    `json:"school_owner_total_schools"`
	SchoolOwnerKey          string `json:"school_owner_key"`
	SchoolOwnerBirthDate    string `json:"school_owner_birth_date"`
	DisId                   int    `json:"dis_id"`
	ManagerId               int    `json:"manager_id"`
	SchoolOwnerACtive       int    `json:"school_owner_active"`
	RegisterDate            string `json:"school_owner_register_date"`
	SchoolOwnerPhoneNumber  string `json:"school_owner_phone_number"`
	SchoolOwnerPersonalId   string `json:"school_owner_personal_id"`
}

type SchoolsProfile struct {
	SchoolId                     int    `json:"school_id"`
	SchoolName                   string `json:"school_name"`
	SchoolManagerName            string `json:"school_manager_name"`
	SchoolOwnerName              string `json:"school_owner_name"`
	SchoolManagerBirthDate       string `json:"school_manager_birth_date"`
	SchoolOwnerBirthDate         string `json:"school_owner_birth_date"`
	SchoolLocation               string `json:"school_location"`
	SchoolManagerLocation        string `json:"school_manager_location"`
	SchoolManagerId              string `json:"school_manager_id"`
	SchoolManagerCurrentLocation string `json:"school_manager_current_location"`
	SchoolTotalStudent           int    `json:"school_total_student"`
	SchoolTotalTeachers          int    `json:"school_total_teachers"`
	SchoolTotalBus               int    `json:"school_total_bus"`
	SchoolTotalAcc               int    `json:"school_total_acc"`
	SchoolOwnerId                int    `json:"school_owner_id"`
	ManagerId                    int    `json:"manager_id"`
	DisId                        int    `json:"dis_id"`
	SchoolKey                    string `json:"school_key"`
	SchoolActive                 bool   `json:"school_active"`
	SchoolStoreActive            bool   `json:"school_store_active"`
}

type StudentProfile struct {
	StudentId                int    `json:"student_id"`
	StudentFullName          string `json:"student_full_name"`
	StudentBirthDate         string `json:"student_birth_date"`
	StudentParentFullName    string `json:"student_parent_full_name"`
	StudentPhoneNumber       string `json:"student_phone_number"`
	StudentParentPhoneNumber string `json:"student_parent_phone_number"`
	StudentLocation          string `json:"student_location"`
	StudentCurrentLocation   string `json:"student_current_location"`
	StudentClass             string `json:"student_class"`
	StudentStudyGroupId      int    `json:"student_study_group_id"`
	StudentIdBack            string `json:"student_id_back"`
	StudentIdFront           string `json:"student_id_front"`
	StudentDeviceType        string `json:"student_device_type"`
	StudentDeviceOsNum       string `json:"student_device_osnum"`
	StudentActive            bool   `json:"student_actve"`
	StudentKey               string `json:"student_key"`
	StudentOverAllNum        int    `json:"student_overall_num"`
	SchoolId                 int    `json:"school_id"`
	DisId                    int    `json:"dis_id"`
	ManagerId                int    `json:"manager_id"`
	BusId                    int    `json:"bus_id"`
}

type BusProfile struct {
	BusId         int    `json:"bus_id"`
	BusName       string `json:"bus_name"`
	BusDocumentId string `json:"bus_document_id"`
	BusKey        string `json:"bus_key"`
	SchoolId      int    `json:"school_id"`
}

type TeacherProfile struct {
	TeacherId               int    `json:"teacher_id"`
	TeacherFullName         string `json:"teacher_name"`
	TeacherBirthDay         string `json:"teacher_birth_day"`
	TeacherLocation         string `json:"teacher_location"`
	TeacherDeviceLocation   string `json:"teacher_device_location"`
	TeacherTotalStudent     int    `json:"teacher_total_student"`
	TeacherTotalScore       int    `json:"teacher_total_score"`
	TeacherTotalPresent     int    `json:"teacher_present"`
	TeacherImagePath        string `json:"teacher_image_path"`
	TeacherIdXNumber        string `json:"teacher_id_xnumber"`
	TeacherIdImagePathFront string `json:"teacher_id_image_path_front"`
	TeacherIdImagePathBack  string `json:"teacher_id_imgae_path_back"`
	TeacherDegree           string `json:"teacher_degree"`
	TeacherMajor            string `json:"teacher_major"`
	TeacherKey              string `json:"teacher_key"`
	TeacherTotalGroups      int    `json:"teacher_total_groups"`
	SchoolId                int    `json:"school_id"`
	OwnerId                 int    `json:"owner_id"`
	TeacherActive           bool   `json:"teacher_active"`
}

type StoreProfile struct {
	StoreId             int `json:"store_id"`
	StoreTotalItems     int `json:"store_total_items"`
	StoreTotalSales     int `json:"store_total_salse"`
	StoreTotalVides     int `json:"store_total_Videos"`
	StoreTotalDocuments int `json:"store_total_documents"`
	SchoolId            int `json:"school_id"`
}
type StoreEvents struct {
	StoreEventId           int    `json:"store_event_id"`
	StoreEventTitle        string `json:"store_event_title"`
	StoreEventDesc         string `json:"store_event_desc"`
	StoreEventTime         string `json:"store_event_time"`
	StoreEventRegisterTime string `json:"store_event_register_time"`
	SchoolId               int    `json:"school_id"`
}
type StoreProducts struct {
	StoreProductsId        int    `json:"store_products_id"`
	StoreProductName       string `json:"store_product_name"`
	StoreProductPath      string `json:"store_product_path"`
	StoreProductPrice     int    `json:"store_product_price"`
	StoreProductImagePath string `json:"store_product_image_path"`
	SchoolId               int    `json:"school_id"`
	ProductListId          int    `json:"product_list_id"`
}
type StudyChatGroupProfile struct {
	StudyChatGroupId           int    `json:"study_chat_group_id"`
	StudyChatGroupName         string `json:"student_groups_name"`
	StudyChatGroupTotalStudent int    `json:"student_groups_total_student"`
	StudyGroupId               int    `json:"student_groups_id"`
	SchoolId                   int    `json:"school_id"`
	TeacherId                  int    `json:"teacher_id"`
	TeacherName                int    `json:"teacher_name"`
}
type StudyGroupProfile struct {
	StudentGroupId           int    `json:"student_groups_id"`
	StudentGroupName         string `json:"student_groups_name"`
	StudentGroupTotalStudent int    `json:"student_groups_total_student"`
	StudentGroupTeacherName  string `json:"student_groups_teacher_name"`
	SchoolId                 int    `json:"school_id"`
}
type StudentPaymentProfile struct {
	StudentPaymentId                   int    `json:"student_payments_id"`
	StudentPaymentRegisterPerson       string `json:"student_payments_register_person"`
	StudentPaymentRegisterTitle        string `json:"student_payments_title"`
	StudentPaymentRegisterDesc         string `json:"student_payments_desc"`
	StudentPaymentRegisterRegisterDate string `json:"student_payments_register_date"`
	StudentPaymentRegisterCurrentDate  string `json:"student_payments_current_date"`
	StudentPaymentTotalAmmount         int    `json:"student_payments_total_ammount"`
	StudentPaymentDiscount             int    `json:"student_payments_discount"`
	SchoolId                           int    `json:"school_id"`
	StudentId                          int    `json:"student_id"`
}
type TeacherPaymentProfile struct {
	TeacherPaymentId      int    `json:"teacher_payment_id"`
	TeacherPaymentTitle   string `json:"teacher_payment_title"`
	TeacherPaymentDesc    string `json:"teacher_payment_desc"`
	TeacherPaymentAmmount string `json:"teacher_payment_ammount"`
	TeacherPaymentDate    string `json:"teacher_payment_date"`
	TeacherRegisterDate   string `json:"teacher_payment_register_date"`
	SchoolId              int    `json:"school_id"`
	TeacherId             int    `json:"teacher_id"`
	TeacherName           string `json:"teacher_name"`
}
type TeacherStudyGroup struct {
	TeacherStudyGroupId int    `json:"teacher_study_group_id"`
	TeacherId           string `json:"teacher_id"`
	StudyGroupId        int    `json:"study_group_id"`
	SchoolId            int    `json:"school_id"`
}

type ScheduleProfile struct {
	ScheduleId          int `json:"schedule_id"`
	StudentStudyGroupId int `json:"study_group_id"`
	Day                 int `json:"day"`
	First               int `json:"first_lesston"`
	Second              int `json:"second_lesston"`
	Thrid               int `json:"thrid_lesston"`
	Fourth              int `json:"fourth_lesston"`
	Fifith              int `json:"fifth_lesston"`
	Sixth               int `json:"sixth_lesston"`
	Seven               int `json:"seven_lesston"`
	SchoolId            int `json:"school_id"`
}
type StudentHistory struct {
	StudentHistoryId       int    `json:"student_history_id"`
	StudentHistoryTitle    string `json:"student_history_title"`
	StudentHistoryMaterial string `json:"student_history_material"`
	StudentHistoryDate     string `json:"student_history_date"`
	StudentId              int    `json:"student_id"`
	SchoolId               int    `json:"school_id"`
}

type StudentExams struct {
	StudentExamId       int    `json:"student_exams_id"`
	StudentName         string `json:"student_name"`
	StudentMake         string `json:"student_mark"`
	StudentExamMaterial string `json:"student_exam_material"`
	TeacherName         string `json:"teacher_name"`
	TeacherId           int    `json:"teacher_id"`
	StudentId           int    `json:"student_id"`
	SchoolId            int    `json:"school_id"`
	Studenthistorydate  string `json:"exam_register_date"`
	StudyGroupId        string `json:"study_group_id"`
}

type SchoolEvents struct {
	SchoolEventId           int    `json:"school_events_id"`
	SchoolEventTitle        string `json:"school_events_title"`
	SchoolEventPerson       string `json:"school_events_person"`
	SchoolEventDesc         string `json:"school_events_desc"`
	SchoolId                int    `json:"school_id"`
	SchoolEventRegisterDate string `json:"school_event_register_date"`
}
