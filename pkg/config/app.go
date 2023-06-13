package config

import  (
	"github.com/joho/godotenv"
	"github.com/nedpals/supabase-go"
	"github.com/joho/godotenv"
	"os"
)


func connect() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	
	supabaseUrl := os.Getenv("supabaseUrl")
	supabaseKey := os.Getenv("supabaseKey")
	supabase := supabase.CreateClient(supabaseUrl, supabaseKey)
	if(supabase == nil){
		fmt.Println("Failed to create Supabase client:")
	}else {
		fmt.Println("Supabase-DB Connected")
	}
}