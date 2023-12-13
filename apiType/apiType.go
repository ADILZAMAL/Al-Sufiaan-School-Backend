package apiType

type ListClassResponse struct {
	Name     string `json:"name"`
	SchoolId uint   `json:"school_id"`
}

type GetSchoolResponse struct {
	Id            uint   `json:"id"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	Mobile        string `json:"mobile"`
	UdiceNo       string `json:"udice_no"`
	Email         string `json:"email"`
	ActiveSession string `json:"active_session"`
}
