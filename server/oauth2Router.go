package server

import (
	_oauth2Controller "github.com/jakkaphatminthana/go-gin/pkg/auth/controller"
	_providerRepository "github.com/jakkaphatminthana/go-gin/pkg/provider/repository"
	_userRepository "github.com/jakkaphatminthana/go-gin/pkg/user/repository"
	_userService "github.com/jakkaphatminthana/go-gin/pkg/user/service"
	"github.com/jakkaphatminthana/go-gin/utils"
)

func (s *ginServer) initOAuth2Router() {
	routerLogin := s.engine.Group("/login")
	routerCallback := s.engine.Group("/callback")

	userRepository := _userRepository.NewUserRepositoryImpl(s.db, utils.GetLogger())
	providerRepository := _providerRepository.NewProviderRepositoryImpl(s.db, utils.GetLogger())
	userService := _userService.NewUserServiceImpl(userRepository, providerRepository)

	oauth2Controller := _oauth2Controller.NewGoogleOAuth2(s.conf, userService)

	routerLogin.GET("/google", oauth2Controller.GoogleLogin)

	routerCallback.GET("/google", oauth2Controller.GoogleCallback)
}
