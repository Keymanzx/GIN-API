package user

import (
	db "api-gin/src/db/mongo"
	"api-gin/src/models/user"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetAllUsers retrieves all users from the users collection
func GetAllUsers() ([]user.Users, error) {
	var users []user.Users

	// Call the ConnectToMongoDB function
	client, err := db.ConnectToMongoDB()
	if err != nil {
		// Handle error
		fmt.Println("Error connecting to MongoDB:", err)
		return nil, err
	}

	// Access users collection
	collection := client.Database("gin_db").Collection("users")

	// Define projection to exclude password field
	projection := bson.M{"password": 0}

	// Find all users with projection
	cursor, err := collection.Find(context.Background(), bson.M{}, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Iterate through the cursor and decode each user
	for cursor.Next(context.Background()) {
		var user user.Users
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// GetByID retrieves a user by their ID from the users collection
func GetByID(userID string) (*user.Users, error) {
	// Call the ConnectToMongoDB function
	client, err := db.ConnectToMongoDB()
	if err != nil {
		// Handle error
		fmt.Println("Error connecting to MongoDB:", err)
		return nil, err
	}

	// Access users collection
	collection := client.Database("gin_db").Collection("users")

	// Convert the string ID to an ObjectID
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	// Define a filter to find the user by ID
	filter := bson.M{"_id": objectID}

	// Find the user
	var result user.Users
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func GetByUserName(userName string) (*user.Users, error) {
	// Call the ConnectToMongoDB function
	client, err := db.ConnectToMongoDB()
	if err != nil {
		// Handle error
		fmt.Println("Error connecting to MongoDB:", err)
		return nil, err
	}

	// Access users collection
	collection := client.Database("gin_db").Collection("users")

	// Define a filter to find the user by ID
	filter := bson.M{"user_name": userName}

	// Find the user
	var result user.Users
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateUser inserts a new user into the users collection
func CreateUser(user *user.Users) error {
	// Call the ConnectToMongoDB function
	client, err := db.ConnectToMongoDB()
	if err != nil {
		// Handle error
		fmt.Println("Error connecting to MongoDB:", err)
		return err
	}

	collection := client.Database("gin_db").Collection("users")

	// Insert user document
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	fmt.Println("User inserted successfully!")

	return nil
}

// UpdateByID updates a user by their ID in the users collection
func UpdateByID(userID string, updatedUser *user.Users) error {
	// Call the ConnectToMongoDB function
	client, err := db.ConnectToMongoDB()
	if err != nil {
		// Handle error
		fmt.Println("Error connecting to MongoDB:", err)
		return err
	}

	// Access users collection
	collection := client.Database("gin_db").Collection("users")

	// Convert the string ID to an ObjectID
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	// Define a filter to find the user by ID
	filter := bson.M{"_id": objectID}

	// Define an update to set the new user data
	update := bson.M{"$set": updatedUser}

	// Update the user
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	fmt.Println("User updated successfully!")

	return nil
}

// DeleteByID deletes a user by their ID from the users collection
func DeleteByID(userID string) error {
	// Call the ConnectToMongoDB function
	client, err := db.ConnectToMongoDB()
	if err != nil {
		// Handle error
		fmt.Println("Error connecting to MongoDB:", err)
		return err
	}

	// Access users collection
	collection := client.Database("gin_db").Collection("users")

	// Convert the string ID to an ObjectID
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	// Define a filter to find the user by ID
	filter := bson.M{"_id": objectID}

	// Delete the user
	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}
