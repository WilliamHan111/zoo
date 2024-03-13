package mongo

import (
	"context"

	"github.com/WilliamHan111/zoo/pkg/conf"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoDB
type MongoDataServer struct {
	client     *mongo.Client
	recordColl *mongo.Collection
}

// 创建MongoDB数据管理器
func NewMongoDataServer(mongoConfig conf.MongoConfig) (*MongoDataServer, error) {
	//连接
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoConfig.Uri).SetAuth(options.Credential{
		AuthSource:  mongoConfig.DB,
		Username:    mongoConfig.User,
		Password:    mongoConfig.Password,
		PasswordSet: mongoConfig.Password != "",
	}))
	if err != nil {
		panic(err)
	}
	//测试
	err = mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		panic(err)
	}
	db := mongoClient.Database(mongoConfig.DB)
	collection := db.Collection(mongoConfig.Collection)
	mod := mongo.IndexModel{
		Keys: bson.M{
			"guid": 1,
		}, Options: nil,
	}

	_, err = collection.Indexes().CreateOne(context.Background(), mod)
	if err != nil {
		panic(err)
	}
	//创建数据管理器
	server := &MongoDataServer{
		client:     mongoClient,
		recordColl: collection,
	}
	return server, nil
}
