package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func (cfg *Config) Login() http.Handler {
	return http.HandlerFunc(cfg.loginHandler)
}

func (cfg *Config) loginHandler(w http.ResponseWriter, r *http.Request) {
	cfg.Logger.Infof("Getting page: %s", r.URL.Path)
	//isAuth, _ := r.Cookie("is_auth")
	//if isAuth.Value == "true" {
	//	http.Redirect(w, r, "/ui", http.StatusFound)
	//	return
	//}
	defError := "Incorrect login or password."
	defErrorStr := "Perhaps you have chosen a different keyboard layout or pressed the \"Caps Lock\" key."
	error_part := `
            <link href="/static/css/theme.css" rel="stylesheet">
            <div class="alert alert-danger" role="alert">
                <strong>%s</strong><br>%s
            </div>`
	html_page := `<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <meta name="description" content="On-line tools service">
        <meta name="author" content="Konstantin Kukharuk (Konstantin@Kukharuk.ru)">
        <title>On-line TOOLS</title>
        <link href="/static/css/bootstrap.min.css" rel="stylesheet">
        <link href="/static/css/ie10-viewport-bug-workaround.css" rel="stylesheet">
        <link href="/static/css/signin.css" rel="stylesheet">
        <!--[if lt IE 9]><script src="/static/js/ie8-responsive-file-warning.js"></script><![endif]-->
        <script src="/static/js/ie-emulation-modes-warning.js"></script>
        <!--[if lt IE 9]>
            <script src="/static/js/html5shiv.min.js"></script>
            <script src="/static/js/respond.min.js"></script>
        <![endif]-->
    </head>
    <body>
        <div class="container">
            <form class="form-signin" method="POST" name="loginform" enctype="multipart/form-data">%s
                <h2 class="form-signin-heading">Please sign in</h2>
                <label for="inputText" class="sr-only">User name</label>
                <input type="text" name="username" id="inputText" class="form-control" placeholder="User name" required autofocus>
                <label for="inputPassword" class="sr-only">Password</label>
                <input type="password" name="password" id="inputPassword" class="form-control" placeholder="Password" required>
                <button class="btn btn-lg btn-primary btn-block" type="submit">Sign in</button>
            </form>
        </div>
        <script src="/static/js/ie10-viewport-bug-workaround.js"></script>
    </body>
</html>
`
	if r.Method == "GET" {
		failAuth, err := r.Cookie("fail_auth")
		if err != nil {
			failCookie := http.Cookie{
				Name:  "fail_auth",
				Value: "0",
			}
			http.SetCookie(w, &failCookie)
			okCookie := http.Cookie{
				Name:  "is_auth",
				Value: "false",
			}
			http.SetCookie(w, &okCookie)
			error_part = ""
		} else {
			countFail, err := strconv.Atoi(failAuth.Value)
			fmt.Print(err)
			if err != nil {
				error_part = fmt.Sprintf(error_part, "Unknown error", "")
			} else {
				failCookie := http.Cookie{
					Name:  "fail_auth",
					Value: string(countFail + 1),
				}
				http.SetCookie(w, &failCookie)
				okCookie := http.Cookie{
					Name:  "is_auth",
					Value: "false",
				}
				http.SetCookie(w, &okCookie)
				error_part = fmt.Sprintf(error_part, defError, defErrorStr)
			}
		}
		fmt.Fprint(w, fmt.Sprintf(html_page, error_part))
	} else if r.Method == "POST" {
		if CheckUser(r.FormValue("username"), r.FormValue("password")) {
			failCookie := http.Cookie{
				Name:  "fail_auth",
				Value: "0",
			}
			http.SetCookie(w, &failCookie)
			okCookie := http.Cookie{
				Name:  "is_auth",
				Value: "true",
			}
			http.SetCookie(w, &okCookie)
			userCookie := http.Cookie{
				Name:  "is_user",
				Value: r.FormValue("username"),
			}
			http.SetCookie(w, &userCookie)
			http.Redirect(w, r, "/ui", http.StatusFound)
			return
		} else {
			failAuth, err := r.Cookie("fail_auth")
			if err != nil {
				failCookie := http.Cookie{
					Name:  "fail_auth",
					Value: "0",
				}
				http.SetCookie(w, &failCookie)
				okCookie := http.Cookie{
					Name:  "is_auth",
					Value: "false",
				}
				http.SetCookie(w, &okCookie)
				error_part = ""
			} else {
				countFail, err := strconv.Atoi(failAuth.Value)
				if err != nil {
					error_part = fmt.Sprintf(error_part, "Unknown error", "")
				} else {
					failCookie := http.Cookie{
						Name:  "fail_auth",
						Value: string(countFail + 1),
					}
					http.SetCookie(w, &failCookie)
					okCookie := http.Cookie{
						Name:  "is_auth",
						Value: "false",
					}
					http.SetCookie(w, &okCookie)
					error_part = fmt.Sprintf(error_part, defError, defErrorStr)
				}
			}
			fmt.Fprint(w, fmt.Sprintf(html_page, error_part))
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
