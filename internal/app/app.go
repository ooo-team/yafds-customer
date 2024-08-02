package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
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

func (a *App) getCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	customerIdStr := r.URL.Query().Get("customer_id")
	if customerIdStr == "" {
		http.Error(w, "Missing customer ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(customerIdStr)
	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	customer, err := a.serviceProvider.CustomerRepo().Get(ctx, uint32(id), false)
	if err != nil {
		http.Error(w, "Failed to find customer", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(customer)
	if err != nil {
		http.Error(w, "Error converting results to json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

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

	id_str, err := a.serviceProvider.CustomerService().Create(ctx, &requestData)
	if err != nil {
		msg := "Failed to create customer: " + err.Error()
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	response := fmt.Sprintf("Created: Phone = %s, Email = %s, Address = %s, ID = %d", requestData.Phone, requestData.Email, requestData.Address, id_str)

	fmt.Fprintln(w, response)
}

func (a *App) initServiceProvider() error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initHttpServer() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/customer/register", a.addCustomer)
	mux.HandleFunc("/customer/get", a.getCustomer)
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
