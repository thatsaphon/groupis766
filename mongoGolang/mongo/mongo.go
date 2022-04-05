package mongo

import (
	"context"
	"goMongo/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var mongoConnection *mongo.Client

// func ConnectMongo() (*mongo.Client, context.Context) {
// 	var ctx context.Context
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://is766:HTZIetTwHD4tkQjn@is766cluster0.dpa1z.mongodb.net/is766db?retryWrites=true&w=majority"))
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, ctx
// 	}
// 	// defer func() {
// 	// 	if err = client.Disconnect(ctx); err != nil {
// 	// 		panic(err)
// 	// 	}
// 	// }()
// 	return client, ctx
// }

func GetMenu() ([]map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://is766:HTZIetTwHD4tkQjn@is766cluster0.dpa1z.mongodb.net/is766db?retryWrites=true&w=majority"))
	if err != nil {
		return []map[string]interface{}{}, err
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("is766db").Collection("menu")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return []map[string]interface{}{}, err
	}
	defer cur.Close(ctx)
	var menus []map[string]interface{}
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			return []map[string]interface{}{}, err
		}
		// do something with result....
		menu := make(map[string]interface{})
		for _, m := range result {
			menu[m.Key] = m.Value
		}
		menus = append(menus, menu)
	}
	if err := cur.Err(); err != nil {
		return []map[string]interface{}{}, err
	}
	return menus, nil
}

func GetRecipe() ([]map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://is766:HTZIetTwHD4tkQjn@is766cluster0.dpa1z.mongodb.net/is766db?retryWrites=true&w=majority"))
	if err != nil {
		return []map[string]interface{}{}, err
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
		return []map[string]interface{}{}, err
	}
	defer cur.Close(ctx)
	var recipes []map[string]interface{}
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			return []map[string]interface{}{}, err
		}
		// do something with result....
		recipe := make(map[string]interface{})
		for _, m := range result {
			recipe[m.Key] = m.Value
		}
		recipes = append(recipes, recipe)
	}
	if err := cur.Err(); err != nil {
		return []map[string]interface{}{}, err
	}
	return recipes, nil
}

func CreateMenu(menu model.CreateMenuRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://is766:HTZIetTwHD4tkQjn@is766cluster0.dpa1z.mongodb.net/is766db?retryWrites=true&w=majority"))
	if err != nil {
		return err
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	collection := client.Database("is766db").Collection("menu")
	menuBson, err := bson.Marshal(menu)

	_, err = collection.InsertOne(ctx, menuBson)

	return nil
}
