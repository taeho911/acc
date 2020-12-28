package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type person struct {
	id primitive.ObjectID `bson:"_id,omitempty"`
	name string `bson:"name,omitempty"`
	age int	`bson:"age,omitempty"`
	address string `bson:"address,omitempty"`
	height float64 `bson:"height,omitempty"`
	alias []string `bson:"alias,omitempty"`
}

func main() {
	kim := person{
		name: "Kim",
		age: 28,
		alias: []string{"dog", "cat", "mice"},
	}
	fmt.Println("kim:", kim)
	
	result, err := bson.Marshal(kim)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("result:", string(result))

	test := bson.M{}
	test["name"] = "Kim"
	test["age"] = 30
	test["address"] = "Tokyo"
	fmt.Println("test:", test)

	test2 := bson.M{"$set": bson.M{}}
	fmt.Println("test2[$set]:", test2["$set"])
}