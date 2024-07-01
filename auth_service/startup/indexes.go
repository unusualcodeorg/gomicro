package startup

import (
	// sampleModel "github.com/unusualcodeorg/gomicro/auth/api/sample/model"
	"github.com/unusualcodeorg/goserve/arch/mongo"
)

func EnsureDbIndexes(db mongo.Database) {
	// go mongo.Document[sampleModel.Sample](&sampleModel.Sample{}).EnsureIndexes(db)
}
