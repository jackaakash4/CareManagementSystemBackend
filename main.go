package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"

    "go.mongodb.org/mongo-driver/mongo"
    
    "go.mongodb.org/mongo-driver/mongo/options"
)


var db *mongo.Database

func main(){
    //load environment variable
    if err := godotenv.Load(); err != nil{
        log.Println("No .env file found")
    }

    //Initialize database
    initDB()

    //setup router
    r := gin.Default()

    //cors middleware
    r.Use(cors.New(cors.Config{
        AllowOrigins: []string{"http://localhost:3000"},
        AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
        ExposeHeaders: []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))

    //API routes
    api := r.Group("/api/v1")
    {
        //User routes
        api.POST("/users", createUser)
        api.GET("/users", getUser)
        api.GET("/users/:id", getUserByID)
        api.PUT("/users/:id", updateUser)
        api.DELETE("/users/:id", deleteUser)

        //Client routes
        api.POST("/clients", createClient)
        api.GET("/clients", getClient)
        api.GET("/clients/:id", getClientByID)
        api.PUT("/clients/:id", updateClient)
        api.DELETE("/CLIENTS/:id", deleteClient)

        //Staff routes
        api.POST("/staffs", createStaff)
        api.GET("/staffs", getStaff)
        api.GET("/staffs/:id", getStaffByID)
        api.PUT("/staffs/:id", updateStaff)
        api.DELETE("/staffs/:id", deleteStaff)
    }

    port := os.Getenv("PORT")
    if port == ""{
        port = "8080"
    }

    log.Printf("Server starting on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, r)) //passing router r

}

func initDB() {
    mongoURI := os.Getenv("MONGODB_URI")
    
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
    if err != nil {
        log.Fatal("Failed to connect to MONGODB: ", err);
    }

    //Test connection
    
