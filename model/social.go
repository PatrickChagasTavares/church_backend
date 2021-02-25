package model

// Social modelo padrão para o Social check
type Social struct {
	ID               int32  `json:"Id,omitempty" db:"id" gorm:"primaryKey;autoIncrement:true" swaggerignore:"true"`
	Date             int64  `json:"Date" db:"date"  gorm:"not null" validate:"required"`
	NameTribe        string `json:"NameTribe" db:"name_tribe" gorm:"size:20; not null" validate:"required"`
	Address          string `json:"Address" db:"address" gorm:"not null" validate:"required"`
	AgeApproximate   string `json:"Age" db:"age" gorm:"not null" validate:"required"`
	Bible            bool   `json:"Bible" db:"bible"`
	Evangelical      bool   `json:"Evangelical" db:"evangelical"`
	Contact          bool   `json:"Contact" db:"contact"`
	FrequentsChurch  bool   `json:"FrequentsChurch" db:"frequents_church"`
	CultHome         bool   `json:"CultHome" db:"cult_home"`
	StudyBible       bool   `json:"StudyBible" db:"study_bible"`
	StudyConfimation bool   `json:"StudyConfimation" db:"study_confimation"`
	PrayerRequest    bool   `json:"PrayerRequest" db:"prayer_request"`
	Reconciled       bool   `json:"Reconciled" db:"reconciled"`
	AcceptVisit      bool   `json:"AcceptVisit" db:"accept_visit"`
	AcceptChrist     bool   `json:"AcceptChrist" db:"accept_christ"`
	Medical          bool   `json:"Medical" db:"medical"`
	Optician         bool   `json:"Optician" db:"optician"`
	Pressure         bool   `json:"Pressure" db:"pressure"`
	Glucose          bool   `json:"Glucose" db:"glucose"`
	Aesthetics       bool   `json:"Aesthetics" db:"aesthetics"`
	CuttingHair      bool   `json:"CuttingHair" db:"cutting_hair"`
	Hairstyle        bool   `json:"Hairstyle" db:"hairstyle"`
	Nail             bool   `json:"Nail" db:"nail"`
	Eyebrow          bool   `json:"Eyebrow" db:"eyebrow"`
	Notes            string `json:"Notes" db:"notes"`
	CreatedAt        int64  `json:"CreatedAt,omitempty" db:"created_at" swaggerignore:"true"`
	UpdatedAt        int64  `json:"UpdatedAt,omitempty" db:"updated_at" swaggerignore:"true"`
	DeletedAt        int64  `json:"DeletedAt,omitempty" db:"deleted_at" swaggerignore:"true"`
}
