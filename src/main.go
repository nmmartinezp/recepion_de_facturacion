// @title Microservicio de Recepción de Facturación
// @version 1.0
// @description Este es un microservicio para la recepción de facturas generadas oara el servicio de SIN.
// @termsOfService http://swagger.io/terms/

// @contact.name Martinez Pardo Nisse Maximiliano
// @contact.url http://www.tusitio.com
// @contact.email nikimartin56@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Ingrese el token en el formato: **Bearer &lt;token&gt;**
package main

import (
	conf "app/src/configs"
	"fmt"
	"log"
)

func main() {
	app := SetupApp()
	config := conf.VarConfig()

	log.Println("Servidor corriendo en " + fmt.Sprintf("http://localhost:%s", config.App.Port))
	app.Run(":" + config.App.Port)
}
