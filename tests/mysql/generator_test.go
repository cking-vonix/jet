package mysql

import (
	"github.com/go-jet/jet/generator/mysql"
	"github.com/go-jet/jet/internal/testutils"
	"github.com/go-jet/jet/tests/dbconfig"
	"gotest.tools/assert"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

const genTestDirRoot = "./.gentestdata3"
const genTestDir3 = "./.gentestdata3/mysql"

func TestGenerator(t *testing.T) {

	for i := 0; i < 3; i++ {
		err := mysql.Generate(genTestDir3, mysql.DBConnection{
			Host:     dbconfig.MySqLHost,
			Port:     dbconfig.MySQLPort,
			User:     dbconfig.MySQLUser,
			Password: dbconfig.MySQLPassword,
			DBName:   "dvds",
		})

		assert.NilError(t, err)

		assertGeneratedFiles(t)
	}

	err := os.RemoveAll(genTestDirRoot)
	assert.NilError(t, err)
}

func TestCmdGenerator(t *testing.T) {
	goInstallJet := exec.Command("sh", "-c", "go install github.com/go-jet/jet/cmd/jet")
	goInstallJet.Stderr = os.Stderr
	err := goInstallJet.Run()
	assert.NilError(t, err)

	err = os.RemoveAll(genTestDir3)
	assert.NilError(t, err)

	cmd := exec.Command("jet", "-source=MySQL", "-dbname=dvds", "-host=localhost", "-port=3306",
		"-user=jet", "-password=jet", "-path="+genTestDir3)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err = cmd.Run()
	assert.NilError(t, err)

	assertGeneratedFiles(t)

	err = os.RemoveAll(genTestDirRoot)
	assert.NilError(t, err)
}

func assertGeneratedFiles(t *testing.T) {
	// Table SQL Builder files
	tableSQLBuilderFiles, err := ioutil.ReadDir(genTestDir3 + "/dvds/table")
	assert.NilError(t, err)

	testutils.AssertFileNamesEqual(t, tableSQLBuilderFiles, "actor.go", "address.go", "category.go", "city.go", "country.go",
		"customer.go", "film.go", "film_actor.go", "film_category.go", "film_text.go", "inventory.go", "language.go",
		"payment.go", "rental.go", "staff.go", "store.go")

	testutils.AssertFileContent(t, genTestDir3+"/dvds/table/actor.go", "\npackage table", actorSQLBuilderFile)

	// View SQL Builder files
	viewSQLBuilderFiles, err := ioutil.ReadDir(genTestDir3 + "/dvds/view")
	assert.NilError(t, err)

	testutils.AssertFileNamesEqual(t, viewSQLBuilderFiles, "actor_info.go", "film_list.go", "nicer_but_slower_film_list.go",
		"sales_by_film_category.go", "customer_list.go", "sales_by_store.go", "staff_list.go")

	testutils.AssertFileContent(t, genTestDir3+"/dvds/view/actor_info.go", "\npackage view", actorInfoSQLBuilerFile)

	// Enums SQL Builder files
	enumFiles, err := ioutil.ReadDir(genTestDir3 + "/dvds/enum")
	assert.NilError(t, err)

	testutils.AssertFileNamesEqual(t, enumFiles, "film_rating.go", "film_list_rating.go", "nicer_but_slower_film_list_rating.go")
	testutils.AssertFileContent(t, genTestDir3+"/dvds/enum/film_rating.go", "\npackage enum", mpaaRatingEnumFile)

	// Model files
	modelFiles, err := ioutil.ReadDir(genTestDir3 + "/dvds/model")
	assert.NilError(t, err)

	testutils.AssertFileNamesEqual(t, modelFiles, "actor.go", "address.go", "category.go", "city.go", "country.go",
		"customer.go", "film.go", "film_actor.go", "film_category.go", "film_text.go", "inventory.go", "language.go",
		"payment.go", "rental.go", "staff.go", "store.go",
		"film_rating.go", "film_list_rating.go", "nicer_but_slower_film_list_rating.go",
		"actor_info.go", "film_list.go", "nicer_but_slower_film_list.go", "sales_by_film_category.go",
		"customer_list.go", "sales_by_store.go", "staff_list.go")

	testutils.AssertFileContent(t, genTestDir3+"/dvds/model/actor.go", "\npackage model", actorModelFile)
}

var mpaaRatingEnumFile = `
package enum

import "github.com/go-jet/jet/mysql"

var FilmRating = &struct {
	G    mysql.StringExpression
	Pg   mysql.StringExpression
	Pg13 mysql.StringExpression
	R    mysql.StringExpression
	Nc17 mysql.StringExpression
}{
	G:    mysql.NewEnumValue("G"),
	Pg:   mysql.NewEnumValue("PG"),
	Pg13: mysql.NewEnumValue("PG-13"),
	R:    mysql.NewEnumValue("R"),
	Nc17: mysql.NewEnumValue("NC-17"),
}
`

var actorSQLBuilderFile = `
package table

import (
	"github.com/go-jet/jet/mysql"
)

var Actor = newActorTable()

type ActorTable struct {
	mysql.Table

	//Columns
	ActorID    mysql.ColumnInteger
	FirstName  mysql.ColumnString
	LastName   mysql.ColumnString
	LastUpdate mysql.ColumnTimestamp

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

// creates new ActorTable with assigned alias
func (a *ActorTable) AS(alias string) *ActorTable {
	aliasTable := newActorTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newActorTable() *ActorTable {
	var (
		ActorIDColumn    = mysql.IntegerColumn("actor_id")
		FirstNameColumn  = mysql.StringColumn("first_name")
		LastNameColumn   = mysql.StringColumn("last_name")
		LastUpdateColumn = mysql.TimestampColumn("last_update")
	)

	return &ActorTable{
		Table: mysql.NewTable("dvds", "actor", ActorIDColumn, FirstNameColumn, LastNameColumn, LastUpdateColumn),

		//Columns
		ActorID:    ActorIDColumn,
		FirstName:  FirstNameColumn,
		LastName:   LastNameColumn,
		LastUpdate: LastUpdateColumn,

		AllColumns:     mysql.ColumnList{ActorIDColumn, FirstNameColumn, LastNameColumn, LastUpdateColumn},
		MutableColumns: mysql.ColumnList{FirstNameColumn, LastNameColumn, LastUpdateColumn},
	}
}
`

var actorModelFile = `
package model

import (
	"time"
)

type Actor struct {
	ActorID    uint16 ` + "`sql:\"primary_key\"`" + `
	FirstName  string
	LastName   string
	LastUpdate time.Time
}
`

var actorInfoSQLBuilerFile = `
package view

import (
	"github.com/go-jet/jet/mysql"
)

var ActorInfo = newActorInfoTable()

type ActorInfoTable struct {
	mysql.Table

	//Columns
	ActorID   mysql.ColumnInteger
	FirstName mysql.ColumnString
	LastName  mysql.ColumnString
	FilmInfo  mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

// creates new ActorInfoTable with assigned alias
func (a *ActorInfoTable) AS(alias string) *ActorInfoTable {
	aliasTable := newActorInfoTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newActorInfoTable() *ActorInfoTable {
	var (
		ActorIDColumn   = mysql.IntegerColumn("actor_id")
		FirstNameColumn = mysql.StringColumn("first_name")
		LastNameColumn  = mysql.StringColumn("last_name")
		FilmInfoColumn  = mysql.StringColumn("film_info")
	)

	return &ActorInfoTable{
		Table: mysql.NewTable("dvds", "actor_info", ActorIDColumn, FirstNameColumn, LastNameColumn, FilmInfoColumn),

		//Columns
		ActorID:   ActorIDColumn,
		FirstName: FirstNameColumn,
		LastName:  LastNameColumn,
		FilmInfo:  FilmInfoColumn,

		AllColumns:     mysql.ColumnList{ActorIDColumn, FirstNameColumn, LastNameColumn, FilmInfoColumn},
		MutableColumns: mysql.ColumnList{ActorIDColumn, FirstNameColumn, LastNameColumn, FilmInfoColumn},
	}
}
`
