package auth

type AuthService interface {
}

func authservice() {

	//key := "Secret-session-key" // Replace with your SESSION_SECRET or similar
	//maxAge := 86400 * 30        // 30 days
	//isProd := false             // Set to true when serving over https
	//
	//store := sessions.NewCookieStore([]byte(key))
	//store.MaxAge(maxAge)
	//store.Options.Path = "/"
	//store.Options.HttpOnly = true // HttpOnly should always be enabled
	//store.Options.Secure = isProd
	//
	//gothic.Store = store
	//
	//goth.UseProviders(
	//	google.New("1045320931526-md4072oj2sbv1922msjhf70m6ob3ldtl.apps.googleusercontent.com",
	//		"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",
	//		"http://localhost:3000/auth/google/callback",
	//		"email", "profile"),
	//)
	//
	//p := pat.New()
	//p.Get("/auth/{provider}/callback", func(res http.ResponseWriter, req *http.Request) {
	//
	//	user, err := gothic.CompleteUserAuth(res, req)
	//	if err != nil {
	//		fmt.Fprintln(res, err)
	//		return
	//	}
	//	t, _ := template.ParseFiles("templates/success.html")
	//	t.Execute(res, user)
	//})
	//
	//p.Get("/auth/{provider}", func(res http.ResponseWriter, req *http.Request) {
	//	gothic.BeginAuthHandler(res, req)
	//})
	//
	//p.Get("/", func(res http.ResponseWriter, req *http.Request) {
	//	t, _ := template.ParseFiles("templates/index.html")
	//	t.Execute(res, false)
	//})
	//log.Println("listening on localhost:3000")
	//log.Fatal(http.ListenAndServe(":3000", p))

}