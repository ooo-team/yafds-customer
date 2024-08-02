package common

import (
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"os"

	"github.com/joho/godotenv"
)

func ReadHeaderParam(w http.ResponseWriter, r *http.Request, param_name string, required bool) string {
	param_str := r.URL.Query().Get(param_name)
	if required && param_str == "" {
		msg := "Missing required request param" + param_name
		http.Error(w, msg, http.StatusBadRequest)
		return ""
	}
	return param_str
}

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

func InitEnv() {
	homedir := os.Getenv("HOME")
	if err := godotenv.Load(homedir + "/.config/go/env/.env"); err != nil {
		log.Panic("No .env file found", err.Error())
	}
}

func Valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func LoadEnvVar(var_name string) string {
	var_, exists := os.LookupEnv(var_name)

	if !exists {
		InitEnv()
		inf_msg := fmt.Sprintf("Env variable %s is not set, calling InitEnv", var_name)
		log.Println(inf_msg)
		var_, exists = os.LookupEnv(var_name)
	}

	err_msg := fmt.Sprintf("Env variable %s is not set", var_name)
	if !exists {
		panic(err_msg)
	}
	return var_
}
