package startup

import (
	auth "github.com/unusualcodeorg/gomicro/auth-service/api/auth/model"
	user "github.com/unusualcodeorg/gomicro/auth-service/api/user/model"
	"github.com/unusualcodeorg/goserve/arch/mongo"
)

func EnsureDbIndexes(db mongo.Database) {
	go mongo.Document[auth.Keystore](&auth.Keystore{}).EnsureIndexes(db)
	go mongo.Document[auth.ApiKey](&auth.ApiKey{}).EnsureIndexes(db)
	go mongo.Document[user.User](&user.User{}).EnsureIndexes(db)
	go mongo.Document[user.Role](&user.Role{}).EnsureIndexes(db)
}
