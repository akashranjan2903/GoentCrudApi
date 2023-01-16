package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gocrud/config"
	"github.com/gocrud/src/router/v1routes"
	_ "github.com/lib/pq"
)

func main() {

	// intializer.Databasecon()

	// insertUser := models.User{Name: "test"}

	// result := intializer.DB.Create(&insertUser)
	// if result.Error != nil {
	// 	println("error")
	// }
	// println("count of entered data is:", result.RowsAffected)
	// // read first data by primary key
	// var read_user models.User
	// result1 := intializer.DB.First(&read_user, 1)
	// println("count of first data is", result1.RowsAffected)
	// fmt.Printf("First Record is: %v\n", read_user)
	// //   read by column name

	// result2 := intializer.DB.First(&read_user, "name = ?", "Jinzhu")
	// println("count of First data by name:", result2.RowsAffected)
	// fmt.Printf("Read First data by name: %v\n", read_user)

	// //  update user
	// var update_user models.User

	// intializer.DB.First(&update_user)

	// update_user.Name = "Akash"
	// updated_Rows := intializer.DB.Save(&update_user)
	// println(" updated user by primary key count:", updated_Rows.RowsAffected)
	// fmt.Printf("updated user by primary key : %v\n", update_user.Name)

	// affected_row := intializer.DB.Model(&models.User{}).Where("name = ?", "Akash").Update("name", "random")
	// println("update user with gven condn count", affected_row.RowsAffected)
	// fmt.Printf("update user with gven condn is : %v\n", update_user.Name)

	//  delete operation

	// deleted_row := intializer.DB.Where("name LIKE ?", "%random%").Delete(&models.User{})
	// println("Total deleted Rows are:", deleted_row.RowsAffected)

	// 	client, err := ent.Open("postgres", "host=localhost port=8080 user=postgres dbname=DemoDatabase password=@kash123 sslmode=disable")
	// 	if err != nil {
	// 		log.Fatalf("failed opening connection to postgres: %v", err)
	// 	}
	// 	CreateUser(context.Background(), client)

	// 	r := gin.Default()
	// 	r.GET("", func(c *gin.Context) {
	// 		c.JSON(200, gin.H{
	// 			"message": "pong",
	// 		})
	// 	})
	// 	r.Use(gin.Recovery())
	// 	if err := r.Run(fmt.Sprintf(":%d", 3030)); err != nil {
	// 		panic(err)
	// 	}
	// 	// r.Run()

	r := v1routes.InitRouter()

	r.Use(gin.Recovery())

	gin.SetMode(gin.DebugMode)

	// cmd.HandleArgs()

	// Run the server
	if err := r.Run(config.GetConfig().App.Port); err != nil {
		panic(err)
	}
}

// func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
// 	u, err := client.User.
// 		Create().
// 		SetAge(30).
// 		SetName("a8m").
// 		Save(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed creating user: %w", err)
// 	}
// 	log.Println("user was created: ", u)
// 	return u, nil
// }
