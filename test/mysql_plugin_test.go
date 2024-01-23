package test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	. "github.com/cloudnumbering/protoc-gen-go-gorm/example/mysql"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/orlangure/gnomock"
	mysql_preset "github.com/orlangure/gnomock/preset/mysql"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/testing/protocmp"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

var (
	mysqlContainer *gnomock.Container
	mysqlDb        *gorm.DB
)

type MySQLPluginSuite struct {
	suite.Suite
}

func TestMySQLPluginSuite(t *testing.T) {
	suite.Run(t, new(MySQLPluginSuite))
}

// TestList tests that the list function works as expected
func (s *MySQLPluginSuite) TestList() {
	// create profiles
	profiles := getMysqlProfiles(s.T(), 3)
	profileProtos := ProfileProtos(profiles)
	_, err := profileProtos.Upsert(context.Background(), mysqlDb)
	require.NoError(s.T(), err)

	// list profiles
	fetchedProfiles := ProfileProtos{}
	err = fetchedProfiles.List(context.Background(), mysqlDb, 10, 0, nil)
	require.NoError(s.T(), err)

	// assert equality, tests are run in parallel so filter down to the ids we know about
	idsSet := hashset.New()
	for _, profile := range profileProtos {
		idsSet.Add(profile.Sid)
	}
	actualProfiles := ProfileProtos{}
	for _, profile := range fetchedProfiles {
		if idsSet.Contains(profile.Sid) {
			actualProfiles = append(actualProfiles, profile)
		}
	}
	assertMysqlProtosEquality(s.T(), profileProtos, actualProfiles,
		protocmp.IgnoreFields(&Profile{}, "created_at", "updated_at"),
	)
}

// TestGetByIds tests that the getByIds function works as expected
func (s *MySQLPluginSuite) TestGetByIds() {
	// create profiles
	profiles := getMysqlProfiles(s.T(), 3)
	profileProtos := ProfileProtos(profiles)
	_, err := profileProtos.Upsert(context.Background(), mysqlDb)
	require.NoError(s.T(), err)

	// get profiles
	ids := lo.Map(profileProtos, func(item *Profile, index int) string {
		return item.Sid
	})

	fetchedProfiles := ProfileProtos{}
	err = fetchedProfiles.GetByIds(context.Background(), mysqlDb, ids)
	require.NoError(s.T(), err)

	// assert equality
	assertMysqlProtosEquality(s.T(), profileProtos, fetchedProfiles,
		protocmp.IgnoreFields(&Profile{}, "created_at", "updated_at"),
	)
}

// TestBase tests that scalar fields are persisted as we expect them to be
func (s *MySQLPluginSuite) TestBase() {
	// create the user
	user := getMysqlUser(s.T())
	userProtos := UserProtos{user}
	_, err := userProtos.Upsert(context.Background(), mysqlDb)
	require.NoError(s.T(), err)

	// fetch the user
	fetchedUserModel, err := getMysqlUserById(user.Sid)
	require.NoError(s.T(), err)
	fetchedUserProto, err := fetchedUserModel.ToProto()
	require.NoError(s.T(), err)

	// assert equality
	assertMysqlProtosEquality(s.T(), userProtos[0], fetchedUserProto,
		protocmp.IgnoreFields(&User{}, "created_at", "updated_at"),
	)
}

// TestHasOneByObject tests that fields related with a has one relationship are persisted as we expect them to be when saved as an object
func (s *MySQLPluginSuite) TestHasOneByObject() {
	// create the user
	user := getMysqlUser(s.T())
	userProtos := UserProtos{user}
	_, err := userProtos.Upsert(context.Background(), mysqlDb)
	require.NoError(s.T(), err)
	expectedUser := userProtos[0]

	// create the address
	address := getMysqlAddress(s.T())
	address.User = user
	addressProtos := AddressProtos{address}
	_, err = addressProtos.Upsert(context.Background(), mysqlDb)
	require.NoError(s.T(), err)

	// set the address on the expected proto for comparison
	expectedUser.Address = addressProtos[0]
	expectedUser.Address.User = nil
	expectedUser.Address.UserId = &expectedUser.Sid

	// fetch the user
	fetchedUserModel, err := getMysqlUserById(user.Sid)
	require.NoError(s.T(), err)
	fetchedUserProto, err := fetchedUserModel.ToProto()
	require.NoError(s.T(), err)

	// assert equality
	assertMysqlProtosEquality(s.T(), userProtos[0], fetchedUserProto,
		protocmp.IgnoreFields(&User{}, "created_at", "updated_at"),
		protocmp.IgnoreFields(&Address{}, "created_at", "updated_at"),
	)
}

// TestHasOneByObject tests that fields related with a has one relationship are persisted as we expect them to be when saved as an id
func (s *MySQLPluginSuite) TestHasOneById() {
	// create the user
	user := getMysqlUser(s.T())
	userProtos := UserProtos{user}
	_, err := userProtos.Upsert(context.Background(), mysqlDb)
	require.NoError(s.T(), err)
	expectedUser := userProtos[0]

	// create the address
	address := getMysqlAddress(s.T())
	address.UserId = &user.Sid
	addressProtos := AddressProtos{address}
	_, err = addressProtos.Upsert(context.Background(), mysqlDb)
	require.NoError(s.T(), err)

	// set the address on the expected proto for comparison
	expectedUser.Address = addressProtos[0]
	expectedUser.Address.User = nil
	expectedUser.Address.UserId = &expectedUser.Sid

	// fetch the user
	fetchedUserModel, err := getMysqlUserById(user.Sid)
	require.NoError(s.T(), err)
	fetchedUserProto, err := fetchedUserModel.ToProto()
	require.NoError(s.T(), err)

	// assert equality
	assertMysqlProtosEquality(s.T(), userProtos[0], fetchedUserProto,
		protocmp.IgnoreFields(&User{}, "created_at", "updated_at"),
		protocmp.IgnoreFields(&Address{}, "created_at", "updated_at"),
	)
}

// TestHasMany tests that fields related with a has many relationship are persisted as we expect them to be
func (s *MySQLPluginSuite) TestHasMany() {
	// create the user
	user := getMysqlUser(s.T())
	userProtos := UserProtos{user}
	_, err := userProtos.Upsert(context.Background(), mysqlDb)
	require.NoError(s.T(), err)
	expectedUser := userProtos[0]

	// create comments
	comments := getMysqlComments(s.T(), 3)
	for _, comment := range comments {
		comment.User = user
	}
	commentProtos := CommentProtos(comments)
	_, err = commentProtos.Upsert(context.Background(), mysqlDb)
	require.NoError(s.T(), err)

	// set the comments on the expected proto for comparison
	expectedUser.Comments = commentProtos
	for _, comment := range expectedUser.Comments {
		// nil user to avoid stack overflow
		comment.User = nil
	}

	// fetch the user
	fetchedUserModel, err := getMysqlUserById(user.Sid)
	require.NoError(s.T(), err)
	fetchedUserProto, err := fetchedUserModel.ToProto()
	require.NoError(s.T(), err)

	// assert equality
	assertMysqlProtosEquality(s.T(), userProtos[0], fetchedUserProto,
		protocmp.IgnoreFields(&User{}, "created_at", "updated_at"),
		protocmp.IgnoreFields(&Comment{}, "created_at", "updated_at"),
	)
}

// TestManyToMany tests that fields related with a many-to-many relationship are persisted as we expect them to be
func (s *MySQLPluginSuite) TestManyToMany() {
	// create the user
	user := getMysqlUser(s.T())
	userProtos := UserProtos{user}
	_, err := userProtos.Upsert(context.Background(), mysqlDb)
	require.NoError(s.T(), err)
	expectedUser := userProtos[0]

	// create profiles
	profiles := getMysqlProfiles(s.T(), 3)
	profileProtos := ProfileProtos(profiles)
	_, err = profileProtos.Upsert(context.Background(), mysqlDb)
	require.NoError(s.T(), err)

	// associate profiles
	session := mysqlDb.Session(&gorm.Session{})
	userModel, err := user.ToModel()
	require.NoError(s.T(), err)
	profileModels, err := profileProtos.ToModels()
	require.NoError(s.T(), err)

	err = session.Model(userModel).Association("Profiles").Replace(profileModels)
	require.NoError(s.T(), err)

	// set the profiles on the expected proto for comparison
	expectedUser.Profiles = profiles

	// fetch the user
	fetchedUserModel, err := getMysqlUserById(user.Sid)
	require.NoError(s.T(), err)
	fetchedUserProto, err := fetchedUserModel.ToProto()
	require.NoError(s.T(), err)

	// assert equality
	assertMysqlProtosEquality(s.T(), userProtos[0], fetchedUserProto,
		protocmp.IgnoreFields(&User{}, "created_at", "updated_at"),
		protocmp.IgnoreFields(&Profile{}, "created_at", "updated_at"),
	)
}

func (s *MySQLPluginSuite) TestSliceTransformers() {
	user := getMysqlUser(s.T())
	users := UserProtos{user}
	models, err := users.ToModels()
	require.NoError(s.T(), err)
	transformedThings, err := models.ToProtos()
	require.NoError(s.T(), err)
	assertMysqlProtosEquality(s.T(), users, transformedThings)
}

func (s *MySQLPluginSuite) SetupSuite() {
	s.T().Parallel()
	preset := mysql_preset.Preset(
		mysql_preset.WithUser("test", "test"),
		mysql_preset.WithDatabase("test"),
	)
	var err error
	portOpt := gnomock.WithCustomNamedPorts(gnomock.NamedPorts{"default": gnomock.Port{
		Protocol: "tcp",
		Port:     5432,
		HostPort: 5432,
	}})
	mysqlContainer, err = gnomock.Start(preset, portOpt)
	require.NoError(s.T(), err)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", mysqlContainer.Host, mysqlContainer.DefaultPort(), "test", "test", "test", "disable")
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
	mysqlDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger})
	require.NoError(s.T(), err)
	err = mysqlDb.AutoMigrate(&UserGormModel{}, &AddressGormModel{}, &CommentGormModel{})
	require.NoError(s.T(), err)
}

func (s *MySQLPluginSuite) TearDownSuite() {
	require.NoError(s.T(), gnomock.Stop(mysqlContainer))
}

func (s *MySQLPluginSuite) SetupTest() {
}

func assertMysqlProtosEquality(t *testing.T, expected, actual interface{}, opts ...cmp.Option) {
	// ignoring id, created_at, updated_at, user_id because the original proto doesn't have those, but the
	// one converted from the created model does
	defaultOpts := []cmp.Option{
		cmpopts.SortSlices(func(x, y *Comment) bool {
			return x.Name < y.Name
		}),
		cmpopts.SortSlices(func(x, y *Profile) bool {
			return x.Name < y.Name
		}),
		protocmp.Transform(),
		protocmp.SortRepeated(func(x, y *Comment) bool {
			return x.Name < y.Name
		}),
		protocmp.SortRepeated(func(x, y *Profile) bool {
			return x.Name < y.Name
		}),
	}
	defaultOpts = append(defaultOpts, opts...)
	diff := cmp.Diff(
		expected,
		actual,
		defaultOpts...,
	)
	require.Empty(t,
		diff,
		diff,
	)
}

func getMysqlUser(t *testing.T) (thing *User) {
	thing = &User{}
	err := gofakeit.Struct(&thing)
	require.NoError(t, err)
	theMap := gofakeit.Map()
	bytes, err := json.Marshal(theMap)
	require.NoError(t, err)
	err = json.Unmarshal(bytes, &thing.AStructpb)
	require.NoError(t, err)
	return
}

func getRandomNumMysqlComments(t *testing.T) []*Comment {
	return getMysqlComments(t, gofakeit.Number(2, 10))
}

func getMysqlComments(t *testing.T, num int) []*Comment {
	comments := []*Comment{}
	for i := 0; i < num; i++ {
		var comment *Comment
		err := gofakeit.Struct(&comment)
		require.NoError(t, err)
		comments = append(comments, comment)
	}
	return comments
}

func getRandomNumMysqlProfiles(t *testing.T) []*Profile {
	return getMysqlProfiles(t, gofakeit.Number(2, 10))
}

func getMysqlProfiles(t *testing.T, num int) []*Profile {
	profiles := []*Profile{}
	for i := 0; i < num; i++ {
		var profile *Profile
		err := gofakeit.Struct(&profile)
		require.NoError(t, err)
		profiles = append(profiles, profile)
	}
	return profiles
}

func getRandomNumMysqlCompanys(t *testing.T) []*Company {
	return getMysqlCompanys(t, gofakeit.Number(2, 10))
}

func getMysqlCompanys(t *testing.T, num int) []*Company {
	companys := []*Company{}
	for i := 0; i < num; i++ {
		companys = append(companys, getMysqlCompany(t))
	}
	return companys
}

func getMysqlCompany(t *testing.T) *Company {
	var company *Company
	err := gofakeit.Struct(&company)
	require.NoError(t, err)
	return company
}

func getMysqlAddress(t *testing.T) *Address {
	var address *Address
	err := gofakeit.Struct(&address)
	require.NoError(t, err)
	address.CompanyBlob = getMysqlCompany(t)
	return address
}

func getMysqlUserById(id string) (*UserGormModel, error) {
	session := mysqlDb.Session(&gorm.Session{})
	var user *UserGormModel
	err := session.Preload(clause.Associations).First(&user, "id = ?", id).Error
	return user, err
}
