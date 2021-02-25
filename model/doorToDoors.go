package model

// DoorToDoors modelo padrão para o DoorToDoors check
type DoorToDoors struct {
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
	Notes            string `json:"Notes" db:"notes"`
	CreatedAt        int64  `json:"CreatedAt,omitempty" db:"created_at" swaggerignore:"true"`
	UpdatedAt        int64  `json:"UpdatedAt,omitempty" db:"updated_at" swaggerignore:"true"`
	DeletedAt        int64  `json:"DeletedAt,omitempty" db:"deleted_at" swaggerignore:"true"`
}
