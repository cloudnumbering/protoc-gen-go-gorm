// Code generated by protoc-gen-go-gorm. DO NOT EDIT.
// source: mysql/example.proto

package example

import (
	context "context"
	json "encoding/json"
	gorm_jsonb "github.com/dariubs/gorm-jsonb"
	pq "github.com/lib/pq"
	lo "github.com/samber/lo"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	gorm "gorm.io/gorm"
	clause "gorm.io/gorm/clause"
	time "time"
)

type UserGormModels []*UserGormModel
type UserProtos []*User
type UserGormModel struct {

	// @gotags: fake:"skip"
	Sid string `gorm:"primaryKey" json:"sid" fake:"skip"`

	// @gotags: fake:"skip"
	CreatedAt *time.Time `gorm:"type:timestamp;" json:"createdAt" fake:"skip"`

	// @gotags: fake:"skip"
	UpdatedAt *time.Time `gorm:"type:timestamp;" json:"updatedAt" fake:"skip"`

	// @gotags: fake:"{price:0.00,1000.00}"
	ADouble float64 `json:"aDouble" fake:"{price:0.00,1000.00}"`

	// @gotags: fake:"{price:0.00,1000.00}"
	AFloat float32 `json:"aFloat" fake:"{price:0.00,1000.00}"`

	// @gotags: fake:"{int32}"
	AnInt32 int32 `json:"anInt32" fake:"{int32}"`

	// @gotags: fake:"{number:9223372036854775807}"
	AnInt64 int64 `json:"anInt64" fake:"{number:9223372036854775807}"`

	// @gotags: fake:"{bool}"
	ABool bool `json:"aBool" fake:"{bool}"`

	// @gotags: fake:"{hackerphrase}"
	AString string `json:"aString" fake:"{hackerphrase}"`

	// @gotags: fake:"skip"
	ABytes []byte `json:"aBytes" fake:"skip"`

	// @gotags: fake:"{price:0.00,1000.00}"
	Doubles pq.Float64Array `gorm:"type:float[];" json:"doubles" fake:"{price:0.00,1000.00}"`

	// @gotags: fake:"{price:0.00,1000.00}"
	Floats pq.Float32Array `gorm:"type:float[];" json:"floats" fake:"{price:0.00,1000.00}"`

	// @gotags: fake:"{int32}"
	Int32S pq.Int32Array `gorm:"type:int[];" json:"int32s" fake:"{int32}"`

	// @gotags: fake:"{number:9223372036854775807}"
	Int64S pq.Int64Array `gorm:"type:int[];" json:"int64s" fake:"{number:9223372036854775807}"`

	// @gotags: fake:"{bool}"
	Bools pq.BoolArray `gorm:"type:bool[];" json:"bools" fake:"{bool}"`

	// @gotags: fake:"{hackerphrase}"
	Strings pq.StringArray `gorm:"type:string[];" json:"strings" fake:"{hackerphrase}"`

	// @gotags: fake:"skip"
	Bytess pq.ByteaArray `gorm:"type:bytes[];" json:"bytess" fake:"skip"`

	// @gotags: fake:"skip"
	OptionalScalarField *string `json:"optionalScalarField" fake:"skip"`

	// @gotags: fake:"skip"
	AStructpb gorm_jsonb.JSONB `gorm:"type:jsonb" json:"aStructpb" fake:"skip"`

	CompanyId *string ``

	// @gotags: fake:"skip"
	Company *CompanyGormModel `gorm:"foreignKey:CompanyId;references:Id;constraint:OnDelete:CASCADE;" json:"company" fake:"skip"`

	// @gotags: fake:"skip"
	CompanyTwoId *string `json:"companyTwoId" fake:"skip"`

	// @gotags: fake:"skip"
	CompanyTwo *CompanyGormModel `gorm:"foreignKey:CompanyTwoId;references:Id;constraint:OnDelete:CASCADE;" json:"companyTwo" fake:"skip"`

	// @gotags: fake:"skip"
	AnUnexpectedId *string `json:"anUnexpectedId" fake:"skip"`

	// @gotags: fake:"skip"
	CompanyThree *CompanyGormModel `gorm:"foreignKey:AnUnexpectedId;references:Id;constraint:OnDelete:CASCADE;" json:"companyThree" fake:"skip"`

	// @gotags: fake:"skip"
	Address *AddressGormModel `gorm:"foreignKey:UserId;references:Id;constraint:OnDelete:CASCADE;" json:"address" fake:"skip"`

	// @gotags: fake:"skip"
	Comments []*CommentGormModel `gorm:"foreignKey:UserId;references:Id;constraint:OnDelete:CASCADE;" json:"comments" fake:"skip"`

	// @gotags: fake:"skip"
	Profiles []*ProfileGormModel `gorm:"foreignKey:Id;references:Id;many2many:users_profiles;joinForeignKey:UserId;joinReferences:ProfileId;constraint:OnDelete:CASCADE;" json:"profiles" fake:"skip"`

	// @gotags: fake:"{number:1,9}"
	IntEnum int `json:"intEnum" fake:"{number:1,9}"`

	// @gotags: fake:"{number:1,9}"
	StringEnum string `json:"stringEnum" fake:"{number:1,9}"`

	// @gotags: fake:"{number:1,9}"
	IntEnumList pq.Int32Array `gorm:"type:int[];" json:"intEnumList" fake:"{number:1,9}"`

	// @gotags: fake:"{number:1,9}"
	StringEnumList pq.StringArray `gorm:"type:string[];" json:"stringEnumList" fake:"{number:1,9}"`

	// @gotags: fake:"{date:2006-01-02}"
	Date *time.Time `json:"date" fake:"{date:2006-01-02}"`

	// @gotags: fake:"{date:2006-01-02}"
	OptionalDate *time.Time `json:"optionalDate" fake:"{date:2006-01-02}"`

	// @gotags: fake:"skip"
	SomeTimestamp *time.Time `gorm:"type:timestamp;" json:"someTimestamp" fake:"skip"`
}

func (m *UserGormModel) TableName() string {
	return "users"
}

func (m UserGormModels) ToProtos() (protos UserProtos, err error) {
	protos = UserProtos{}
	for _, model := range m {
		var proto *User
		if proto, err = model.ToProto(); err != nil {
			return
		}
		protos = append(protos, proto)
	}
	return
}

func (p UserProtos) ToModels() (models UserGormModels, err error) {
	models = UserGormModels{}
	for _, proto := range p {
		var model *UserGormModel
		if model, err = proto.ToModel(); err != nil {
			return
		}
		models = append(models, model)
	}
	return
}

func (m *UserGormModel) ToProto() (theProto *User, err error) {
	if m == nil {
		return
	}
	theProto = &User{}

	theProto.Sid = m.Sid

	if m.CreatedAt != nil {
		theProto.CreatedAt = m.CreatedAt.Format(time.RFC3339Nano)
	}

	if m.UpdatedAt != nil {
		theProto.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}

	theProto.ADouble = m.ADouble

	theProto.AFloat = m.AFloat

	theProto.AnInt32 = m.AnInt32

	theProto.AnInt64 = m.AnInt64

	theProto.ABool = m.ABool

	theProto.AString = m.AString

	theProto.ABytes = m.ABytes

	theProto.Doubles = m.Doubles

	theProto.Floats = m.Floats

	theProto.Int32S = m.Int32S

	theProto.Int64S = m.Int64S

	theProto.Bools = m.Bools

	theProto.Strings = m.Strings

	theProto.Bytess = m.Bytess

	theProto.OptionalScalarField = m.OptionalScalarField

	if m.AStructpb != nil {
		var jsonBytes []byte
		if jsonBytes, err = json.Marshal(m.AStructpb); err != nil {
			return
		}
		if err = json.Unmarshal(jsonBytes, &theProto.AStructpb); err != nil {
			return
		}
	}

	if theProto.Company, err = m.Company.ToProto(); err != nil {
		return
	}

	theProto.CompanyTwoId = m.CompanyTwoId

	if theProto.CompanyTwo, err = m.CompanyTwo.ToProto(); err != nil {
		return
	}

	theProto.AnUnexpectedId = m.AnUnexpectedId

	if theProto.CompanyThree, err = m.CompanyThree.ToProto(); err != nil {
		return
	}

	if theProto.Address, err = m.Address.ToProto(); err != nil {
		return
	}

	if len(m.Comments) > 0 {
		theProto.Comments = []*Comment{}
		for _, item := range m.Comments {
			var CommentsProto *Comment
			if CommentsProto, err = item.ToProto(); err != nil {
				return
			} else {
				theProto.Comments = append(theProto.Comments, CommentsProto)
			}
		}
	}

	if len(m.Profiles) > 0 {
		theProto.Profiles = []*Profile{}
		for _, item := range m.Profiles {
			var ProfilesProto *Profile
			if ProfilesProto, err = item.ToProto(); err != nil {
				return
			} else {
				theProto.Profiles = append(theProto.Profiles, ProfilesProto)
			}
		}
	}

	theProto.IntEnum = EnumOne(m.IntEnum)

	theProto.StringEnum = EnumOne(EnumOne_value[m.StringEnum])

	if len(m.IntEnumList) > 0 {
		theProto.IntEnumList = []EnumOne{}
		for _, val := range m.IntEnumList {
			theProto.IntEnumList = append(theProto.IntEnumList, EnumOne(val))
		}
	}

	if len(m.StringEnumList) > 0 {
		theProto.StringEnumList = []EnumOne{}
		for _, val := range m.StringEnumList {
			theProto.StringEnumList = append(theProto.StringEnumList, EnumOne(EnumOne_value[val]))
		}
	}

	if m.Date != nil {
		theProto.Date = m.Date.UTC().Format("2006-01-02")
	}

	if m.OptionalDate != nil {
		theProto.OptionalDate = lo.ToPtr(m.OptionalDate.UTC().Format("2006-01-02"))
	}

	if m.SomeTimestamp != nil {
		theProto.SomeTimestamp = timestamppb.New(*m.SomeTimestamp)
	}

	return
}

func (p *User) ToModel() (theModel *UserGormModel, err error) {
	if p == nil {
		return
	}
	theModel = &UserGormModel{}

	theModel.Sid = p.Sid

	if p.CreatedAt != "" {
		var timestamp time.Time
		if timestamp, err = time.Parse(time.RFC3339Nano, p.CreatedAt); err != nil {
			return
		}
		theModel.CreatedAt = &timestamp
	}

	if p.UpdatedAt != nil {
		theModel.UpdatedAt = lo.ToPtr(p.UpdatedAt.AsTime())
	}

	theModel.ADouble = p.ADouble

	theModel.AFloat = p.AFloat

	theModel.AnInt32 = p.AnInt32

	theModel.AnInt64 = p.AnInt64

	theModel.ABool = p.ABool

	theModel.AString = p.AString

	theModel.ABytes = p.ABytes

	theModel.Doubles = p.Doubles

	theModel.Floats = p.Floats

	theModel.Int32S = p.Int32S

	theModel.Int64S = p.Int64S

	theModel.Bools = p.Bools

	theModel.Strings = p.Strings

	theModel.Bytess = p.Bytess

	theModel.OptionalScalarField = p.OptionalScalarField

	if p.AStructpb != nil {
		var jsonBytes []byte
		if jsonBytes, err = json.Marshal(p.AStructpb); err != nil {
			return
		}
		if err = json.Unmarshal(jsonBytes, &theModel.AStructpb); err != nil {
			return
		}
	}

	if theModel.Company, err = p.Company.ToModel(); err != nil {
		return
	}

	// if the object is present, the object's id overrides the existing id field value
	if p.Company != nil {
		theModel.CompanyId = &p.Company.Sid
	}

	theModel.CompanyTwoId = p.CompanyTwoId

	if theModel.CompanyTwo, err = p.CompanyTwo.ToModel(); err != nil {
		return
	}

	// if the object is present, the object's id overrides the existing id field value
	if p.CompanyTwo != nil {
		theModel.CompanyTwoId = &p.CompanyTwo.Sid
	}

	theModel.AnUnexpectedId = p.AnUnexpectedId

	if theModel.CompanyThree, err = p.CompanyThree.ToModel(); err != nil {
		return
	}

	// if the object is present, the object's id overrides the existing id field value
	if p.CompanyThree != nil {
		theModel.AnUnexpectedId = &p.CompanyThree.Sid
	}

	if theModel.Address, err = p.Address.ToModel(); err != nil {
		return
	}

	if len(p.Comments) > 0 {
		theModel.Comments = []*CommentGormModel{}
		for _, item := range p.Comments {
			var CommentsModel *CommentGormModel
			if CommentsModel, err = item.ToModel(); err != nil {
				return
			} else {
				theModel.Comments = append(theModel.Comments, CommentsModel)
			}
		}
	}

	if len(p.Profiles) > 0 {
		theModel.Profiles = []*ProfileGormModel{}
		for _, item := range p.Profiles {
			var ProfilesModel *ProfileGormModel
			if ProfilesModel, err = item.ToModel(); err != nil {
				return
			} else {
				theModel.Profiles = append(theModel.Profiles, ProfilesModel)
			}
		}
	}

	theModel.IntEnum = int(p.IntEnum)

	theModel.StringEnum = p.StringEnum.String()

	if len(p.IntEnumList) > 0 {
		theModel.IntEnumList = pq.Int32Array{}
		for _, val := range p.IntEnumList {
			theModel.IntEnumList = append(theModel.IntEnumList, int32(val))
		}
	}

	if len(p.StringEnumList) > 0 {
		theModel.StringEnumList = pq.StringArray{}
		for _, val := range p.StringEnumList {
			theModel.StringEnumList = append(theModel.StringEnumList, val.String())
		}
	}

	if p.Date != "" {
		var date time.Time
		if date, err = time.Parse("2006-01-02", p.Date); err != nil {
			return
		}
		dateUTC := date.UTC()
		theModel.Date = &dateUTC
	}

	if p.OptionalDate != nil {
		var date time.Time
		if date, err = time.Parse("2006-01-02", *p.OptionalDate); err != nil {
			return
		}
		dateUTC := date.UTC()
		theModel.OptionalDate = &dateUTC
	}

	if p.SomeTimestamp != nil {
		theModel.SomeTimestamp = lo.ToPtr(p.SomeTimestamp.AsTime())
	}

	return
}

func (m UserGormModels) GetByModelIds(ctx context.Context, tx *gorm.DB, preloads ...string) (err error) {
	ids := []string{}
	for _, model := range m {
		ids = append(ids, model.Sid)
	}
	if len(ids) > 0 {
		statement := tx.Preload(clause.Associations)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		err = statement.Where("sid in ?", ids).Find(&m).Error
	}
	return
}

// Upsert creates the protos using an on conflict clause to do updates. This function does not update *any* associations
// use gorm's association mode functions to update associations as you see fit after calling upsert. See https://gorm.io/docs/associations.html#Replace-Associations
func (p *UserProtos) Upsert(ctx context.Context, tx *gorm.DB) (models UserGormModels, err error) {
	if p != nil {
		models, err = p.ToModels()
		if err != nil {
			return
		}
		// create new session so the tx isn't modified
		session := tx.Session(&gorm.Session{})
		err = session.
			// on conflict, update all fields
			Clauses(clause.OnConflict{
				UpdateAll: true,
			}).
			// exclude associations from upsert
			Omit(clause.Associations).
			Create(&models).Error
	}
	return
}

func (p *UserProtos) List(ctx context.Context, tx *gorm.DB, limit, offset int, order interface{}, preloads ...string) (err error) {
	if p != nil {
		var models UserGormModels
		statement := tx.Preload(clause.Associations).Limit(limit).Offset(offset)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		if order != nil {
			statement = statement.Order(order)
		}
		if err = statement.Find(&models).Error; err != nil {
			return
		}
		if len(models) > 0 {
			*p, err = models.ToProtos()
		} else {
			*p = UserProtos{}
		}
	}
	return
}

func (p *UserProtos) GetByIds(ctx context.Context, tx *gorm.DB, ids []string, preloads ...string) (err error) {
	if p != nil {
		var models UserGormModels
		statement := tx.Preload(clause.Associations)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		if err = statement.Where("sid in ?", ids).Find(&models).Error; err != nil {
			return
		}
		if len(models) > 0 {
			*p, err = models.ToProtos()
		} else {
			*p = UserProtos{}
		}
	}
	return
}

func DeleteUserGormModels(ctx context.Context, tx *gorm.DB, ids []string) error {
	statement := tx.Where("sid in ?", ids)
	return statement.Delete(&UserGormModel{}).Error
}

type CompanyGormModels []*CompanyGormModel
type CompanyProtos []*Company
type CompanyGormModel struct {

	// @gotags: fake:"skip"
	Sid string `gorm:"primaryKey" json:"sid" fake:"skip"`

	// @gotags: fake:"skip"
	CreatedAt *time.Time `gorm:"type:timestamp;" json:"createdAt" fake:"skip"`

	// @gotags: fake:"skip"
	UpdatedAt *time.Time `gorm:"type:timestamp;" json:"updatedAt" fake:"skip"`

	// @gotags: fake:"{name}"
	Name string `json:"name" fake:"{name}"`
}

func (m *CompanyGormModel) TableName() string {
	return "companies"
}

func (m CompanyGormModels) ToProtos() (protos CompanyProtos, err error) {
	protos = CompanyProtos{}
	for _, model := range m {
		var proto *Company
		if proto, err = model.ToProto(); err != nil {
			return
		}
		protos = append(protos, proto)
	}
	return
}

func (p CompanyProtos) ToModels() (models CompanyGormModels, err error) {
	models = CompanyGormModels{}
	for _, proto := range p {
		var model *CompanyGormModel
		if model, err = proto.ToModel(); err != nil {
			return
		}
		models = append(models, model)
	}
	return
}

func (m *CompanyGormModel) ToProto() (theProto *Company, err error) {
	if m == nil {
		return
	}
	theProto = &Company{}

	theProto.Sid = m.Sid

	if m.CreatedAt != nil {
		theProto.CreatedAt = timestamppb.New(*m.CreatedAt)
	}

	if m.UpdatedAt != nil {
		theProto.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}

	theProto.Name = m.Name

	return
}

func (p *Company) ToModel() (theModel *CompanyGormModel, err error) {
	if p == nil {
		return
	}
	theModel = &CompanyGormModel{}

	theModel.Sid = p.Sid

	if p.CreatedAt != nil {
		theModel.CreatedAt = lo.ToPtr(p.CreatedAt.AsTime())
	}

	if p.UpdatedAt != nil {
		theModel.UpdatedAt = lo.ToPtr(p.UpdatedAt.AsTime())
	}

	theModel.Name = p.Name

	return
}

func (m CompanyGormModels) GetByModelIds(ctx context.Context, tx *gorm.DB, preloads ...string) (err error) {
	ids := []string{}
	for _, model := range m {
		ids = append(ids, model.Sid)
	}
	if len(ids) > 0 {
		statement := tx.Preload(clause.Associations)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		err = statement.Where("sid in ?", ids).Find(&m).Error
	}
	return
}

// Upsert creates the protos using an on conflict clause to do updates. This function does not update *any* associations
// use gorm's association mode functions to update associations as you see fit after calling upsert. See https://gorm.io/docs/associations.html#Replace-Associations
func (p *CompanyProtos) Upsert(ctx context.Context, tx *gorm.DB) (models CompanyGormModels, err error) {
	if p != nil {
		models, err = p.ToModels()
		if err != nil {
			return
		}
		// create new session so the tx isn't modified
		session := tx.Session(&gorm.Session{})
		err = session.
			// on conflict, update all fields
			Clauses(clause.OnConflict{
				UpdateAll: true,
			}).
			// exclude associations from upsert
			Omit(clause.Associations).
			Create(&models).Error
	}
	return
}

func (p *CompanyProtos) List(ctx context.Context, tx *gorm.DB, limit, offset int, order interface{}, preloads ...string) (err error) {
	if p != nil {
		var models CompanyGormModels
		statement := tx.Preload(clause.Associations).Limit(limit).Offset(offset)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		if order != nil {
			statement = statement.Order(order)
		}
		if err = statement.Find(&models).Error; err != nil {
			return
		}
		if len(models) > 0 {
			*p, err = models.ToProtos()
		} else {
			*p = CompanyProtos{}
		}
	}
	return
}

func (p *CompanyProtos) GetByIds(ctx context.Context, tx *gorm.DB, ids []string, preloads ...string) (err error) {
	if p != nil {
		var models CompanyGormModels
		statement := tx.Preload(clause.Associations)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		if err = statement.Where("sid in ?", ids).Find(&models).Error; err != nil {
			return
		}
		if len(models) > 0 {
			*p, err = models.ToProtos()
		} else {
			*p = CompanyProtos{}
		}
	}
	return
}

func DeleteCompanyGormModels(ctx context.Context, tx *gorm.DB, ids []string) error {
	statement := tx.Where("sid in ?", ids)
	return statement.Delete(&CompanyGormModel{}).Error
}

type AddressGormModels []*AddressGormModel
type AddressProtos []*Address
type AddressGormModel struct {

	// @gotags: fake:"skip"
	Sid string `gorm:"primaryKey" json:"sid" fake:"skip"`

	// @gotags: fake:"skip"
	CreatedAt *time.Time `gorm:"type:timestamp;" json:"createdAt" fake:"skip"`

	// @gotags: fake:"skip"
	UpdatedAt *time.Time `gorm:"type:timestamp;" json:"updatedAt" fake:"skip"`

	// @gotags: fake:"{name}"
	Name string `json:"name" fake:"{name}"`

	// @gotags: fake:"skip"
	UserId *string `json:"userId" fake:"skip"`

	// @gotags: fake:"skip"
	User *UserGormModel `gorm:"foreignKey:UserId;references:Id;constraint:OnDelete:CASCADE;" json:"user" fake:"skip"`

	// @gotags: fake:"skip"
	CompanyBlob gorm_jsonb.JSONB `gorm:"type:jsonb" json:"companyBlob" fake:"skip"`
}

func (m *AddressGormModel) TableName() string {
	return "addresses"
}

func (m AddressGormModels) ToProtos() (protos AddressProtos, err error) {
	protos = AddressProtos{}
	for _, model := range m {
		var proto *Address
		if proto, err = model.ToProto(); err != nil {
			return
		}
		protos = append(protos, proto)
	}
	return
}

func (p AddressProtos) ToModels() (models AddressGormModels, err error) {
	models = AddressGormModels{}
	for _, proto := range p {
		var model *AddressGormModel
		if model, err = proto.ToModel(); err != nil {
			return
		}
		models = append(models, model)
	}
	return
}

func (m *AddressGormModel) ToProto() (theProto *Address, err error) {
	if m == nil {
		return
	}
	theProto = &Address{}

	theProto.Sid = m.Sid

	if m.CreatedAt != nil {
		theProto.CreatedAt = timestamppb.New(*m.CreatedAt)
	}

	if m.UpdatedAt != nil {
		theProto.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}

	theProto.Name = m.Name

	theProto.UserId = m.UserId

	if theProto.User, err = m.User.ToProto(); err != nil {
		return
	}

	if m.CompanyBlob != nil {
		var jsonBytes []byte
		if jsonBytes, err = json.Marshal(m.CompanyBlob); err != nil {
			return
		}
		if err = json.Unmarshal(jsonBytes, &theProto.CompanyBlob); err != nil {
			return
		}
	}

	return
}

func (p *Address) ToModel() (theModel *AddressGormModel, err error) {
	if p == nil {
		return
	}
	theModel = &AddressGormModel{}

	theModel.Sid = p.Sid

	if p.CreatedAt != nil {
		theModel.CreatedAt = lo.ToPtr(p.CreatedAt.AsTime())
	}

	if p.UpdatedAt != nil {
		theModel.UpdatedAt = lo.ToPtr(p.UpdatedAt.AsTime())
	}

	theModel.Name = p.Name

	theModel.UserId = p.UserId

	if theModel.User, err = p.User.ToModel(); err != nil {
		return
	}

	// if the object is present, the object's id overrides the existing id field value
	if p.User != nil {
		theModel.UserId = &p.User.Sid
	}

	if p.CompanyBlob != nil {
		var jsonBytes []byte
		if jsonBytes, err = json.Marshal(p.CompanyBlob); err != nil {
			return
		}
		if err = json.Unmarshal(jsonBytes, &theModel.CompanyBlob); err != nil {
			return
		}
	}

	return
}

func (m AddressGormModels) GetByModelIds(ctx context.Context, tx *gorm.DB, preloads ...string) (err error) {
	ids := []string{}
	for _, model := range m {
		ids = append(ids, model.Sid)
	}
	if len(ids) > 0 {
		statement := tx.Preload(clause.Associations)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		err = statement.Where("sid in ?", ids).Find(&m).Error
	}
	return
}

// Upsert creates the protos using an on conflict clause to do updates. This function does not update *any* associations
// use gorm's association mode functions to update associations as you see fit after calling upsert. See https://gorm.io/docs/associations.html#Replace-Associations
func (p *AddressProtos) Upsert(ctx context.Context, tx *gorm.DB) (models AddressGormModels, err error) {
	if p != nil {
		models, err = p.ToModels()
		if err != nil {
			return
		}
		// create new session so the tx isn't modified
		session := tx.Session(&gorm.Session{})
		err = session.
			// on conflict, update all fields
			Clauses(clause.OnConflict{
				UpdateAll: true,
			}).
			// exclude associations from upsert
			Omit(clause.Associations).
			Create(&models).Error
	}
	return
}

func (p *AddressProtos) List(ctx context.Context, tx *gorm.DB, limit, offset int, order interface{}, preloads ...string) (err error) {
	if p != nil {
		var models AddressGormModels
		statement := tx.Preload(clause.Associations).Limit(limit).Offset(offset)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		if order != nil {
			statement = statement.Order(order)
		}
		if err = statement.Find(&models).Error; err != nil {
			return
		}
		if len(models) > 0 {
			*p, err = models.ToProtos()
		} else {
			*p = AddressProtos{}
		}
	}
	return
}

func (p *AddressProtos) GetByIds(ctx context.Context, tx *gorm.DB, ids []string, preloads ...string) (err error) {
	if p != nil {
		var models AddressGormModels
		statement := tx.Preload(clause.Associations)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		if err = statement.Where("sid in ?", ids).Find(&models).Error; err != nil {
			return
		}
		if len(models) > 0 {
			*p, err = models.ToProtos()
		} else {
			*p = AddressProtos{}
		}
	}
	return
}

func DeleteAddressGormModels(ctx context.Context, tx *gorm.DB, ids []string) error {
	statement := tx.Where("sid in ?", ids)
	return statement.Delete(&AddressGormModel{}).Error
}

type CommentGormModels []*CommentGormModel
type CommentProtos []*Comment
type CommentGormModel struct {

	// @gotags: fake:"skip"
	Sid string `gorm:"primaryKey" json:"sid" fake:"skip"`

	// @gotags: fake:"skip"
	CreatedAt *time.Time `gorm:"type:timestamp;" json:"createdAt" fake:"skip"`

	// @gotags: fake:"skip"
	UpdatedAt *time.Time `gorm:"type:timestamp;" json:"updatedAt" fake:"skip"`

	// @gotags: fake:"{name}"
	Name string `json:"name" fake:"{name}"`

	UserId *string ``

	// @gotags: fake:"skip"
	User *UserGormModel `gorm:"foreignKey:UserId;references:Id;constraint:OnDelete:CASCADE;" json:"user" fake:"skip"`
}

func (m *CommentGormModel) TableName() string {
	return "comments"
}

func (m CommentGormModels) ToProtos() (protos CommentProtos, err error) {
	protos = CommentProtos{}
	for _, model := range m {
		var proto *Comment
		if proto, err = model.ToProto(); err != nil {
			return
		}
		protos = append(protos, proto)
	}
	return
}

func (p CommentProtos) ToModels() (models CommentGormModels, err error) {
	models = CommentGormModels{}
	for _, proto := range p {
		var model *CommentGormModel
		if model, err = proto.ToModel(); err != nil {
			return
		}
		models = append(models, model)
	}
	return
}

func (m *CommentGormModel) ToProto() (theProto *Comment, err error) {
	if m == nil {
		return
	}
	theProto = &Comment{}

	theProto.Sid = m.Sid

	if m.CreatedAt != nil {
		theProto.CreatedAt = timestamppb.New(*m.CreatedAt)
	}

	if m.UpdatedAt != nil {
		theProto.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}

	theProto.Name = m.Name

	if theProto.User, err = m.User.ToProto(); err != nil {
		return
	}

	return
}

func (p *Comment) ToModel() (theModel *CommentGormModel, err error) {
	if p == nil {
		return
	}
	theModel = &CommentGormModel{}

	theModel.Sid = p.Sid

	if p.CreatedAt != nil {
		theModel.CreatedAt = lo.ToPtr(p.CreatedAt.AsTime())
	}

	if p.UpdatedAt != nil {
		theModel.UpdatedAt = lo.ToPtr(p.UpdatedAt.AsTime())
	}

	theModel.Name = p.Name

	if theModel.User, err = p.User.ToModel(); err != nil {
		return
	}

	// if the object is present, the object's id overrides the existing id field value
	if p.User != nil {
		theModel.UserId = &p.User.Sid
	}

	return
}

func (m CommentGormModels) GetByModelIds(ctx context.Context, tx *gorm.DB, preloads ...string) (err error) {
	ids := []string{}
	for _, model := range m {
		ids = append(ids, model.Sid)
	}
	if len(ids) > 0 {
		statement := tx.Preload(clause.Associations)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		err = statement.Where("sid in ?", ids).Find(&m).Error
	}
	return
}

// Upsert creates the protos using an on conflict clause to do updates. This function does not update *any* associations
// use gorm's association mode functions to update associations as you see fit after calling upsert. See https://gorm.io/docs/associations.html#Replace-Associations
func (p *CommentProtos) Upsert(ctx context.Context, tx *gorm.DB) (models CommentGormModels, err error) {
	if p != nil {
		models, err = p.ToModels()
		if err != nil {
			return
		}
		// create new session so the tx isn't modified
		session := tx.Session(&gorm.Session{})
		err = session.
			// on conflict, update all fields
			Clauses(clause.OnConflict{
				UpdateAll: true,
			}).
			// exclude associations from upsert
			Omit(clause.Associations).
			Create(&models).Error
	}
	return
}

func (p *CommentProtos) List(ctx context.Context, tx *gorm.DB, limit, offset int, order interface{}, preloads ...string) (err error) {
	if p != nil {
		var models CommentGormModels
		statement := tx.Preload(clause.Associations).Limit(limit).Offset(offset)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		if order != nil {
			statement = statement.Order(order)
		}
		if err = statement.Find(&models).Error; err != nil {
			return
		}
		if len(models) > 0 {
			*p, err = models.ToProtos()
		} else {
			*p = CommentProtos{}
		}
	}
	return
}

func (p *CommentProtos) GetByIds(ctx context.Context, tx *gorm.DB, ids []string, preloads ...string) (err error) {
	if p != nil {
		var models CommentGormModels
		statement := tx.Preload(clause.Associations)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		if err = statement.Where("sid in ?", ids).Find(&models).Error; err != nil {
			return
		}
		if len(models) > 0 {
			*p, err = models.ToProtos()
		} else {
			*p = CommentProtos{}
		}
	}
	return
}

func DeleteCommentGormModels(ctx context.Context, tx *gorm.DB, ids []string) error {
	statement := tx.Where("sid in ?", ids)
	return statement.Delete(&CommentGormModel{}).Error
}

type ProfileGormModels []*ProfileGormModel
type ProfileProtos []*Profile
type ProfileGormModel struct {

	// @gotags: fake:"skip"
	Sid string `gorm:"primaryKey" json:"sid" fake:"skip"`

	// @gotags: fake:"skip"
	CreatedAt *time.Time `gorm:"type:timestamp;" json:"createdAt" fake:"skip"`

	// @gotags: fake:"skip"
	UpdatedAt *time.Time `gorm:"type:timestamp;" json:"updatedAt" fake:"skip"`

	// @gotags: fake:"{name}"
	Name string `json:"name" fake:"{name}"`
}

func (m *ProfileGormModel) TableName() string {
	return "profiles"
}

func (m ProfileGormModels) ToProtos() (protos ProfileProtos, err error) {
	protos = ProfileProtos{}
	for _, model := range m {
		var proto *Profile
		if proto, err = model.ToProto(); err != nil {
			return
		}
		protos = append(protos, proto)
	}
	return
}

func (p ProfileProtos) ToModels() (models ProfileGormModels, err error) {
	models = ProfileGormModels{}
	for _, proto := range p {
		var model *ProfileGormModel
		if model, err = proto.ToModel(); err != nil {
			return
		}
		models = append(models, model)
	}
	return
}

func (m *ProfileGormModel) ToProto() (theProto *Profile, err error) {
	if m == nil {
		return
	}
	theProto = &Profile{}

	theProto.Sid = m.Sid

	if m.CreatedAt != nil {
		theProto.CreatedAt = timestamppb.New(*m.CreatedAt)
	}

	if m.UpdatedAt != nil {
		theProto.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}

	theProto.Name = m.Name

	return
}

func (p *Profile) ToModel() (theModel *ProfileGormModel, err error) {
	if p == nil {
		return
	}
	theModel = &ProfileGormModel{}

	theModel.Sid = p.Sid

	if p.CreatedAt != nil {
		theModel.CreatedAt = lo.ToPtr(p.CreatedAt.AsTime())
	}

	if p.UpdatedAt != nil {
		theModel.UpdatedAt = lo.ToPtr(p.UpdatedAt.AsTime())
	}

	theModel.Name = p.Name

	return
}

func (m ProfileGormModels) GetByModelIds(ctx context.Context, tx *gorm.DB, preloads ...string) (err error) {
	ids := []string{}
	for _, model := range m {
		ids = append(ids, model.Sid)
	}
	if len(ids) > 0 {
		statement := tx.Preload(clause.Associations)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		err = statement.Where("sid in ?", ids).Find(&m).Error
	}
	return
}

// Upsert creates the protos using an on conflict clause to do updates. This function does not update *any* associations
// use gorm's association mode functions to update associations as you see fit after calling upsert. See https://gorm.io/docs/associations.html#Replace-Associations
func (p *ProfileProtos) Upsert(ctx context.Context, tx *gorm.DB) (models ProfileGormModels, err error) {
	if p != nil {
		models, err = p.ToModels()
		if err != nil {
			return
		}
		// create new session so the tx isn't modified
		session := tx.Session(&gorm.Session{})
		err = session.
			// on conflict, update all fields
			Clauses(clause.OnConflict{
				UpdateAll: true,
			}).
			// exclude associations from upsert
			Omit(clause.Associations).
			Create(&models).Error
	}
	return
}

func (p *ProfileProtos) List(ctx context.Context, tx *gorm.DB, limit, offset int, order interface{}, preloads ...string) (err error) {
	if p != nil {
		var models ProfileGormModels
		statement := tx.Preload(clause.Associations).Limit(limit).Offset(offset)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		if order != nil {
			statement = statement.Order(order)
		}
		if err = statement.Find(&models).Error; err != nil {
			return
		}
		if len(models) > 0 {
			*p, err = models.ToProtos()
		} else {
			*p = ProfileProtos{}
		}
	}
	return
}

func (p *ProfileProtos) GetByIds(ctx context.Context, tx *gorm.DB, ids []string, preloads ...string) (err error) {
	if p != nil {
		var models ProfileGormModels
		statement := tx.Preload(clause.Associations)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		if err = statement.Where("sid in ?", ids).Find(&models).Error; err != nil {
			return
		}
		if len(models) > 0 {
			*p, err = models.ToProtos()
		} else {
			*p = ProfileProtos{}
		}
	}
	return
}

func DeleteProfileGormModels(ctx context.Context, tx *gorm.DB, ids []string) error {
	statement := tx.Where("sid in ?", ids)
	return statement.Delete(&ProfileGormModel{}).Error
}