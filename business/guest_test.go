package business

import (
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	"fiber1/config"
	"fiber1/model"

	"github.com/stretchr/testify/assert"

	"github.com/hexops/autogold"
	"github.com/jmoiron/sqlx"
	"github.com/ory/dockertest/v3"
)

var testDb *sqlx.DB

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	resource, err := pool.Run("mariadb", "latest", []string{"MYSQL_ROOT_PASSWORD=secret"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	var connStr string
	if err := pool.Retry(func() error {
		log.Println("Checking whether mariadb is up..")
		connStr = fmt.Sprintf("root:secret@(127.0.0.1:%s)/mysql", resource.GetPort("3306/tcp"))
		testDb = config.ConnectMysql(connStr)
		if testDb != nil {
			return nil
		}
		return errors.New(`testDb is nil`)
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	log.Println("Mariadb is up, start testing..")
	log.Println(connStr)
	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestGuest(t *testing.T) {
	user := model.NewUser(testDb)
	err := user.Migrate()
	if err != nil {
		t.Errorf("error migrating: %v", err)
	}

	guest := Guest{
		Db: testDb,
	}

	t.Run(`registerWithEmptyEmailMustFail`, func(t *testing.T) {
		in := Guest_RegisterIn{}
		out := guest.Register(&in)

		want := autogold.Want(`registerWithEmptyEmailMustFail1`, Guest_RegisterOut{CommonOut: CommonOut{
			StatusCode: 400,
			ErrorMsg:   "invalid email",
		}})
		want.Equal(t, out)
	})

	t.Run(`registerWithEmptyPasswordMustFail`, func(t *testing.T) {
		in := Guest_RegisterIn{
			Email: `a@gmail.com`,
		}
		out := guest.Register(&in)

		want := autogold.Want(`registerWithEmptyPasswordMustFail1`, Guest_RegisterOut{CommonOut: CommonOut{
			StatusCode: 400,
			ErrorMsg:   "invalid password",
		}})
		want.Equal(t, out)
	})

	t.Run(`registerWithEmptyNameMustFail`, func(t *testing.T) {
		in := Guest_RegisterIn{
			Email:    `a@gmail.com`,
			Password: `123456`,
		}
		out := guest.Register(&in)

		want := autogold.Want(`registerWithEmptyNameMustFail1`, Guest_RegisterOut{CommonOut: CommonOut{
			StatusCode: 400,
			ErrorMsg:   "invalid name",
		}})
		want.Equal(t, out)
	})

	t.Run(`registerMustSucceed`, func(t *testing.T) {
		in := Guest_RegisterIn{
			Email:    `a@gmail.com`,
			Password: `123456`,
			Name:     `Ayaya`,
		}
		out := guest.Register(&in)

		if out.User != nil {
			assert.NotEmpty(t, out.User.Password)
			out.User.Password = ``
		}
		want := autogold.Want(`registerMustSucceed1`, Guest_RegisterOut{User: &model.User{
			Id:    1,
			Email: "a@gmail.com",
			Name:  "Ayaya",
		}})
		want.Equal(t, out)
	})

	t.Run(`registerAgainWithSameEmailMustFail`, func(t *testing.T) {
		in := Guest_RegisterIn{
			Email:    `a@gmail.com`,
			Password: `231241`,
			Name:     `Ayaya`,
		}
		out := guest.Register(&in)

		want := autogold.Want(`registerAgainWithSameEmailMustFail1`, Guest_RegisterOut{CommonOut: CommonOut{
			StatusCode: 400,
			ErrorMsg:   "email already used",
		}})
		want.Equal(t, out)
	})

	t.Run(`loginWithEmptyPassMustFail`, func(t *testing.T) {
		in := Guest_LoginIn{
			Email: `a@gmail.com`,
		}
		out := guest.Login(&in)

		want := autogold.Want(`loginWithEmptyPassMustFail1`, Guest_LoginOut{CommonOut: CommonOut{
			StatusCode: 400,
			ErrorMsg:   "invalid password",
		}})
		want.Equal(t, out)
	})

	t.Run(`loginWithWrongPasswordMustFail`, func(t *testing.T) {
		in := Guest_LoginIn{
			Email:    `a@gmail.com`,
			Password: `123`,
		}
		out := guest.Login(&in)

		want := autogold.Want(`loginWithWrongPasswordMustFail1`, Guest_LoginOut{CommonOut: CommonOut{
			StatusCode: 400,
			ErrorMsg:   "user or password does not match",
		}})
		want.Equal(t, out)
	})

	t.Run(`loginWithUnknownEmailMustFail`, func(t *testing.T) {
		in := Guest_LoginIn{
			Email:    `b@gmail.com`,
			Password: `123`,
		}
		out := guest.Login(&in)

		want := autogold.Want(`loginWithUnknownEmailMustFail1`, Guest_LoginOut{CommonOut: CommonOut{
			StatusCode: 400,
			ErrorMsg:   "user not found",
		}})
		want.Equal(t, out)
	})

	t.Run(`loginMustSucceed`, func(t *testing.T) {
		in := Guest_LoginIn{
			Email:    `a@gmail.com`,
			Password: `123456`,
		}
		out := guest.Login(&in)

		assert.NotEmpty(t, out.SetCookie)
		out.SetCookie = ``
		if out.User != nil {
			assert.NotEmpty(t, out.User.Password)
			out.User.Password = ``
		}
		want := autogold.Want(`loginMustSucceed1`, Guest_LoginOut{User: &model.User{
			Id:    1,
			Email: "a@gmail.com",
			Name:  "Ayaya",
		}})
		want.Equal(t, out)
	})

	//t.Run(`failedUploadProfilePicture`, func(t *testing.T) {
	//	in := Guest_UploadProfilePictureIn{}
	//	guest.UploadToS3 = func(filename string, content string) error {
	//		return errors.New(`failed upload`)
	//	}
	//	out := guest.UploadProfilePicture(&in)
	//})
	//
	//t.Run(`failedUploadProfilePictureMustSuccess`, func(t *testing.T) {
	//	in := Guest_UploadProfilePictureIn{}
	//	guest.UploadToS3 = func(filename string, content string) error {
	//		return nil
	//	}
	//	out := guest.UploadProfilePicture(&in)
	//})
}
