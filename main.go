package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

	"PetVet/controllers"
	"PetVet/data"
	"PetVet/middlewares"
	"PetVet/repositories"
	"PetVet/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Advertencia: no se pudo cargar .env (probablemente en producción)")
	}

	if err := data.RunPostgresDB(); err != nil {
		log.Fatal("Error DB:", err)
	}
	db := data.DB
	fmt.Println("✅ Conectado a PostgreSQL con GORM")

	userRepo := repositories.NewUserRepositoryGORM(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	adminRepo := repositories.NewAdminRepositoryGORM(db)
	adminService := services.NewAdminService(adminRepo)
	adminController := controllers.NewAdminController(adminService)

	appointmentRepo := repositories.NewAppointmentRepositoryGORM(db)
	appointmentService := services.NewAppointmentService(appointmentRepo)
	appointmentController := controllers.NewAppointmentController(appointmentService)

	petRepo := repositories.NewPetRepositoryGORM(db)
	petService := services.NewPetService(petRepo)
	petController := controllers.NewPetController(petService)

	adminTypeRepo := repositories.NewAdminTypeRepositoryGORM(db)
	adminTypeService := services.NewAdminTypeService(adminTypeRepo)
	adminTypeController := controllers.NewAdminTypeController(adminTypeService)

	userRoleRepo := repositories.NewUserRoleRepositoryGORM(db)
	userRoleService := services.NewUserRoleService(userRoleRepo)
	userRoleController := controllers.NewUserRoleController(userRoleService)

	speciesRepo := repositories.NewSpeciesRepositoryGORM(db)
	speciesService := services.NewSpeciesService(speciesRepo)
	speciesController := controllers.NewSpeciesController(speciesService)

	r := mux.NewRouter()
	userController.RegisterRoutes(r, middlewares.JWTAuthMiddleware)
	adminController.RegisterPublicRoutes(r, middlewares.AdminRegisterMiddleware)
	adminController.RegisterProtectedRoutes(r, middlewares.AdminProtected)
	appointmentController.RegisterRoutes(r, middlewares.JWTAuthMiddleware)
	petController.RegisterRoutes(r, middlewares.JWTAuthMiddleware)
	adminTypeController.RegisterRoutes(r, middlewares.AdminProtected)
	userRoleController.RegisterRoutes(r, middlewares.JWTAuthMiddleware)
	speciesController.RegisterRoutes(r, middlewares.JWTAuthMiddleware)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("🚀 Servidor escuchando en http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
