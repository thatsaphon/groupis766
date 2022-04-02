package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var mongoConnection *mongo.Client

func ConnectMongo() (*mongo.Client, context.Context) {
	var ctx context.Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://is766:HTZIetTwHD4tkQjn@is766cluster0.dpa1z.mongodb.net/is766db?retryWrites=true&w=majority"))
	if err != nil {
		fmt.Println(err)
		return nil, ctx
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	return client, ctx
}

func GetMenu() []map[string]interface{} {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://is766:HTZIetTwHD4tkQjn@is766cluster0.dpa1z.mongodb.net/is766db?retryWrites=true&w=majority"))
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// client, ctx := ConnectMongo()
	collection := client.Database("is766db").Collection("menu")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		fmt.Println(err)
		// fmt.Println(err)
	}
	defer cur.Close(ctx)
	var menues []map[string]interface{}
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			fmt.Println(err)
		}
		// menues = append(menues, result)
		// do something with result....
		menu := make(map[string]interface{})
		for _, m := range result {
			menu[m.Key] = m.Value
		}
		menues = append(menues, menu)
	}
	if err := cur.Err(); err != nil {
		fmt.Println(err)
	}
	return menues
}

func GetRecipe() []map[string]interface{} {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://is766:HTZIetTwHD4tkQjn@is766cluster0.dpa1z.mongodb.net/is766db?retryWrites=true&w=majority"))
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// client, ctx := ConnectMongo()
	collection := client.Database("is766db").Collection("recipe")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		fmt.Println(err)
		// fmt.Println(err)
	}
	defer cur.Close(ctx)
	var recipes []map[string]interface{}
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			fmt.Println(err)
		}
		// menues = append(menues, result)
		// do something with result....
		recipe := make(map[string]interface{})
		for _, m := range result {
			recipe[m.Key] = m.Value
		}
		recipes = append(recipes, recipe)
	}
	if err := cur.Err(); err != nil {
		fmt.Println(err)
	}
	return recipes
}
