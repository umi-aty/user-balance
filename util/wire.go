package util

import (
	"log"
	"net"
	"net/http"
	"time"
	configdb "userbalance/config"
	"userbalance/controllers"
	"userbalance/repositories"
	"userbalance/services"

	"gorm.io/gorm"
)

var (
	httpClient *http.Client
	db         *gorm.DB = configdb.ConfigDatabase()
)

func ProvideUserAuthRepo() repositories.UserRepository {
	return repositories.NewUserRepo(db, httpClient)
}

func ProvideUserAuthService() services.AuthService {
	return services.NewAuthService(ProvideUserAuthRepo())
}

func ProvideUserJwtService() services.JwtService {
	return services.NewJWTService()
}

func ProvideUserService() controllers.AuthController {
	return controllers.NewAuthController(ProvideUserAuthService(), ProvideUserJwtService())
}

func ProvideHttpClient() *http.Client {
	transport, ok := http.DefaultTransport.(*http.Transport) // get default roundtripper transport
	if !ok {
		log.Fatal("infra: defaulTransport is not *http.Transport")
	}

	transport.DisableKeepAlives = false
	transport.MaxIdleConns = 256
	transport.MaxIdleConnsPerHost = 256
	transport.MaxConnsPerHost = 0
	transport.ResponseHeaderTimeout = 60 * time.Second
	transport.IdleConnTimeout = 60 * time.Second
	transport.TLSHandshakeTimeout = time.Duration(30) * time.Second
	transport.DialContext = (&net.Dialer{
		Timeout:   time.Duration(60) * time.Second,
		KeepAlive: time.Duration(60) * time.Second,
		DualStack: true,
	}).DialContext

	httpClient = &http.Client{
		Timeout:   time.Duration(60) * time.Second,
		Transport: transport,
	}

	return httpClient
}
