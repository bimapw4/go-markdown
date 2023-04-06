package entity

type DataPayload struct {
	OriginData    map[string]interface{} `json:"origin_data"`
	TargetData    map[string]interface{} `json:"target_data"`
	ParsingFormat []ParsingFormatPayload `json:"parsing_format"`
}

type OriginDataPayload struct {
	PrivyId        string               `json:"privy_id"`
	UserData       OriginUserData       `json:"user_data"`
	AdditionalData OriginAdditionalData `json:"additional_data"`
}

type OriginUserData struct {
	Fullname    string      `json:"fullname"`
	Nik         string      `json:"nik"`
	Email       string      `json:"email"`
	DateOfBirth interface{} `json:"date_of_birth"`
	Phone       string      `json:"phone"`
}

type OriginAdditionalData struct {
	Fullname           string                       `json:"blood_type"`
	Nik                string                       `json:"gender"`
	ResidentialAddress OriginResidentialAddressData `json:"residential_address"`
	MedicineAllergy    string                       `json:"medicine_allergy"`
	DoctorCode         string                       `json:"doctor_code"`
	PolyclinicCode     string                       `json:"polyclinic_code"`
	QueueNumber        string                       `json:"queue_number"`
	PickedSchdule      string                       `string:"picked_schdule"`
}

type OriginResidentialAddressData struct {
	Address1    string `json:"address_1"`
	Address2    string `json:"address_2"`
	SubDistrict string `json:"sub_district"`
	City        string `json:"city"`
	Province    string `json:"province"`
	ZipCode     string `json:"zipcode"`
}

type TargetDataPayload struct {
	Nama     string `json:"nama"`
	Nik      string `json:"nik"`
	TglLahir string `json:"tgl_lhr"`
	Email    string `json:"email"`
	Hp       string `json:"hp"`
	Alamat   string `json:"alamat"`

	GolonganDarah string `json:"golongan_darah"`
	JenisKelamin  string `json:"jenis_kelamin"`
	Alergi        string `json:"alergi"`
	KodeDokter    string `json:"kode_dokter"`
	KodeBagian    string `json:"kode_bagian"`
	NoAntrian     string `json:"no_antrian"`

	TglDaftar string `json:"tgl_daftar"`
	Privyid   string `json:"privyid"`
}

type ParsingFormatPayload struct {
	Origin string `json:"origin"`
	Target string `json:"target"`
	Format string `json:"format"`
}
