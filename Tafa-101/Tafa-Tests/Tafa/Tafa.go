package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type userData struct {
	UserName    string `json:"userName"`
	PhoneNumber string `json:"phoneNumber"`
	UserTickets uint   `json:"userTickets"`
}

type volunteerData struct {
	VolunteerName  string `json:"volunteerName"`
	VolunteerEmail string `json:"volunteerEmail"`
}

type InfoResponse struct {
	ConferenceName      string `json:"conferenceName"`
	TotalTickets        uint   `json:"totalTickets"`
	RemainingTickets    uint   `json:"remainingTickets"`
	RemainingVolunteers uint   `json:"remainingVolunteers"`
}

type BookingRequest struct {
	UserName    string `json:"userName"`
	PhoneNumber string `json:"phoneNumber"`
	UserTickets uint   `json:"userTickets"`
}

type VolunteerRequest struct {
	VolunteerName  string `json:"volunteerName"`
	VolunteerEmail string `json:"volunteerEmail"`
}

type APIResponse struct {
	Message          string      `json:"message"`
	RemainingTickets uint        `json:"remainingTickets,omitempty"`
	Data             interface{} `json:"data,omitempty"`
}

var (
	mutex             sync.Mutex
	conferenceName         = "Tafa Business Conference"
	conferenceTickets uint = 30
	remainingTickets  uint = 30
	bookings          []*userData

	volunteersNeeded    uint = 7
	remainingVolunteers uint = 7
	onboarding          []*volunteerData
)

func createUserData(name string, phone string, tickets uint) *userData {
	return &userData{
		UserName:    name,
		PhoneNumber: phone,
		UserTickets: tickets,
	}
}

func bookingProcess(remainingTickets *uint, userTickets uint) {
	*remainingTickets = *remainingTickets - userTickets
}

func createvolunteerData(fullName string, email string) *volunteerData {
	return &volunteerData{
		VolunteerName:  fullName,
		VolunteerEmail: email,
	}
}

func volunteerOnboarding(volunteersNeeded uint, remainingVolunteers *uint, registeredVolunteers uint) {
	*remainingVolunteers = volunteersNeeded - registeredVolunteers
}

func handleInformation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	mutex.Lock()
	resp := InfoResponse{
		ConferenceName:      conferenceName,
		TotalTickets:        conferenceTickets,
		RemainingTickets:    remainingTickets,
		RemainingVolunteers: remainingVolunteers,
	}
	mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func handleBookings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodPost:
		var req BookingRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(APIResponse{Message: "Invalid JSON request body"})
			return
		}
		if len(req.UserName) >= 50 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(APIResponse{Message: "Name is to long"})
			return

		}

		if len(req.PhoneNumber) < 10 || len(req.PhoneNumber) > 13 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(APIResponse{Message: "Invalid phone number"})
			return
		}
		if req.UserTickets == 0 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(APIResponse{Message: "you must book atleast one ticket"})
			return
		}

		mutex.Lock()
		defer mutex.Unlock()

		if req.UserTickets > remainingTickets {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(APIResponse{
				Message:          fmt.Sprintf("Sorry but, we only have %d Tickets left", remainingTickets),
				RemainingTickets: remainingTickets,
			})
			return
		}

		bookingProcess(&remainingTickets, req.UserTickets)
		userPointer := createUserData(req.UserName, req.PhoneNumber, req.UserTickets)
		bookings = append(bookings, userPointer)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(APIResponse{
			Message:          fmt.Sprintf("Thank you %s for purchasing %d tickets.", userPointer.UserName, userPointer.UserTickets),
			RemainingTickets: remainingTickets,
			Data:             userPointer,
		})

	case http.MethodGet:
		mutex.Lock()
		defer mutex.Unlock()
		json.NewEncoder(w).Encode(bookings)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(APIResponse{Message: "Method not allowed"})
	}
}

func handleVolunteers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodPost:
		var req VolunteerRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(APIResponse{Message: "Invalid JSON request body"})
			return
		}

		mutex.Lock()
		defer mutex.Unlock()

		if remainingVolunteers == 0 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(APIResponse{Message: "No volunteer spots remaining"})
			return
		}

		volunteerPointer := createvolunteerData(req.VolunteerName, req.VolunteerEmail)
		onboarding = append(onboarding, volunteerPointer)

		registeredCount := uint(len(onboarding))
		volunteerOnboarding(volunteersNeeded, &remainingVolunteers, registeredCount)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(APIResponse{
			Message: "Thank you for Registering we look forward to working with you",
			Data:    volunteerPointer,
		})

	case http.MethodGet:
		mutex.Lock()
		defer mutex.Unlock()
		json.NewEncoder(w).Encode(onboarding)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(APIResponse{Message: "Method not allowed"})
	}
}

func main() {
	http.HandleFunc("/", handleInformation)
	http.HandleFunc("/bookings", handleBookings)
	http.HandleFunc("/volunteers", handleVolunteers)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
