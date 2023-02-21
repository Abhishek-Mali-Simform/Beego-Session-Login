package services

import (
	"LoginUser/models"
	"LoginUser/utils"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

var err error

func ConnectDb() {
	var err error
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", utils.EnvConfigs.DBHost, utils.EnvConfigs.DBUser, utils.EnvConfigs.DBName, utils.EnvConfigs.DBPassword, utils.EnvConfigs.DBPort)
	err = orm.RegisterDriver("postgres", orm.DRPostgres)
	ErrorCheck("Error Registering Driver.", err)
	err = orm.RegisterDataBase(utils.EnvConfigs.DBAlias, "postgres", dbURI)
	ErrorCheck("Error Registering Database.", err)
	orm.RegisterModel(new(models.User))
	orm.SetMaxIdleConns(utils.EnvConfigs.DBAlias, 100)
	orm.SetMaxOpenConns(utils.EnvConfigs.DBAlias, 100)
	err = orm.RunSyncdb(utils.EnvConfigs.DBAlias, false, true)
	ErrorCheck("Error Syncing Database.", err)
}

func GetDB() (o orm.Ormer) {
	o = orm.NewOrm()
	err := o.Using(utils.EnvConfigs.DBAlias)
	ErrorCheck("Couldn't Use Default Alias. Reason: ", err)
	return
}
