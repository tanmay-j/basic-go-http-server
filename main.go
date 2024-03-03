package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var BASE_PATH = "./public/"
var USER_FILE_PATH = "./users.txt"

var ErrUserNotFound = errors.New("User Not Found")

func readFile(path string) (contentStr string, errStr error) {
	path = BASE_PATH + path

	content, err := os.ReadFile(path)

	return string(content), err
}

func login(username string, password string) (role string, err error) {
	content, err := os.ReadFile(USER_FILE_PATH)

    users := strings.Split(string(content),"\r\n") 
    //users = users[:len(users)-1]
	
	for _,user := range users {
		fmt.Printf("%q\n",user)
        user_split := strings.Split(user,",")
        user = user_split[0]
        role:= "user"
        if len(user_split) > 1 {
            role = user_split[1]
        }
		username_password := fmt.Sprintf("%s:%s",username,password)	
		fmt.Printf("Comparing: '%s' '%s'\n\n", user, username_password)
		if username_password == user {
			return role, nil
		}
	}
	//fmt.Printf("%q\n",users) 

	return "default",ErrUserNotFound
}

func getRoot(w http.ResponseWriter, r *http.Request) {

	//ctx := r.Context()

	username, password, isBasicAuth := r.BasicAuth()
	
	fmt.Printf("Got Request at /, context: %v %v %v \n", username, password, isBasicAuth)

	if !isBasicAuth {
		http.Error(w,"Bad Request: Basic Authentication not used.", 400)
		return
	}

	role, err := login(username,password)

	if err != nil {
		errCode := 500
		errStr := "Internal Server Error while auth"
		if errors.Is(err, ErrUserNotFound) {
			errCode = 401
            w.Header().Set("WWW-Authenticate",`Basic realm="Secure Server"`)
			errStr = "Unauthorized"
		}
		http.Error(w,errStr,errCode)
		return

	}
    
    if r.RequestURI[len(r.RequestURI)-3:] == ".go" {
        if role != "admin" {
            errCode := 403
            errStr := "Forbidden"
		    
		    http.Error(w,errStr,errCode)
		    return      
        }
    }

	content, err := readFile(r.RequestURI)

	if err != nil {
		errStr := fmt.Sprintf("Error: %v", err)
		// io.WriteString(w, errStr)
		errCode := 500
		if errors.Is(err, os.ErrNotExist) {
			errCode = 404
		}
		http.Error(w,errStr,errCode)
		return
	}


	fmt.Printf("Content: %s\n", content)

	responseContent := fmt.Sprintf("%s",content)

	io.WriteString(w, responseContent)
}

func main() {

	//login("applebanana")
	//return
	
	http.HandleFunc("/", getRoot)

	err := http.ListenAndServe(":8080",nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server Closed\n")
	} else if err != nil {
		fmt.Println("error starting server: %s\n", err)
		os.Exit(1)
	}
}
