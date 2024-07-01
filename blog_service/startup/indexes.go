package startup

import (
	blog "github.com/unusualcodeorg/gomicro/blog-service/api/blog/model"
	"github.com/unusualcodeorg/goserve/arch/mongo"
)

func EnsureDbIndexes(db mongo.Database) {
	go mongo.Document[blog.Blog](&blog.Blog{}).EnsureIndexes(db)
}
