package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	model "github.com/ooo-team/yafds/internal/model/customer"
)

type HttpResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type App struct {
	httpServer      *http.Server
	serviceProvider *serviceProvider
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page!")
}

func (a *App) addCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Чтение тела запроса
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Декодирование JSON
	var requestData model.CustomerInfo
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	_, err = a.serviceProvider.CustomerService().Create(ctx, &requestData)
	if err != nil {
		msg := "Failed to create customer: " + err.Error()
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	response := fmt.Sprintf("Created: Phone = %s, Email = %s, Address = %s", requestData.Phone, requestData.Email, requestData.Address)

	fmt.Fprintln(w, response)
}

func (a *App) initServiceProvider() error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initHttpServer() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/customer/register", a.addCustomer)
	a.httpServer = &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 МБ
	}

	return nil
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}
	a.initHttpServer()
	a.initServiceProvider()
	return a, nil
}

func (a *App) Run() {
	if err := a.httpServer.ListenAndServe(); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
