package handler

import (
	"enceremony-be/internal/auth/authorizer"
	"enceremony-be/internal/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)
import (
	"github.com/markbates/goth/gothic"
)

type AuthHandler interface {
	HandleCallback(ctx *gin.Context)
	BeginAuthHandler(c *gin.Context)
}

type impl struct {
	conf    *config.Config
	authSvc authorizer.Service
}

func (i *impl) HandleCallback(c *gin.Context) {
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		fmt.Fprintln(c.Writer, err)
		return
	}

	fmt.Println(user)
	fmt.Println("name: ", user.Name)
	fmt.Println("token: ", user.AccessToken)
	fmt.Println("email: ", user.Email)
	fmt.Println("atSecret: ", user.AccessTokenSecret)
	fmt.Println("raw: ", user.RawData)
	//t, _ := template.ParseFiles("templates/success.html")
	//t.Execute(c.Writer, user)
	//c.JSON(http.StatusOK, user)
	res, err := i.authSvc.VerifyAndCreateUser(c, &user)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, res)
}

//{
//"iss": "https://accounts.google.com",
//"azp": "1045320931526-md4072oj2sbv1922msjhf70m6ob3ldtl.apps.googleusercontent.com",
//"aud": "1045320931526-md4072oj2sbv1922msjhf70m6ob3ldtl.apps.googleusercontent.com",
//"sub": "117077591790970225951",
//"email": "enceremony23@gmail.com",
//"email_verified": true,
//"at_hash": "eoRD0_uCh9Ncp1HROvMIpQ",
//"name": "Enceremony",
//"picture": "https://lh3.googleusercontent.com/a/AAcHTtchDiz_xf7J5Ky3tQ00NH79dDknzUd5gDdESWMiT6H-=s96-c",
//"given_name": "Enceremony",
//"locale": "en-GB",
//"iat": 1690167604,
//"exp": 1690171204
//}

func (i *impl) BeginAuthHandler(c *gin.Context) {
	key := i.conf.GCP.SessionKey // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30         // 30 days
	isProd := i.conf.IsProd      // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/auth/google"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store

	q := c.Request.URL.Query()
	q.Add("provider", c.Param("provider"))
	c.Request.URL.RawQuery = q.Encode()

	goth.UseProviders(
		google.New(i.conf.GCP.ClientId, i.conf.GCP.Secret, i.conf.GCP.CallbackUrl, "email", "profile"),
	)

	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func NewAuthHandler(conf *config.Config, authSvc authorizer.Service) AuthHandler {
	return &impl{
		conf:    conf,
		authSvc: authSvc,
	}
}

func InitGoogleAuthConnection(conf *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}
