package helpers

import (
	"encoding/json"
	// "fmt"
	"log"
	"net/http"
)

// var magicExcelHeader string

// func SetMagicExcelHeader(oidcSettings OidcSettings) {
// 	magicExcelHeader = fmt.Sprintf(`Bearer resource="https://management.azure.com/" client_id="%s", trusted_issuers="00000001-0000-0000-c000-000000000000@*", token_types="app_asserted_user_v1 service_asserted_app_v1", authorization_uri="https://login.microsoftonline.com/%s/oauth2/v2.0/authorize",Basic Realm=""`, oidcSettings.ClientID, oidcSettings.TenantID)
// }

func constructBody(msg string) []byte {
	resp := make(map[string]string)
	resp["message"] = msg
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	return jsonResp
}

func sendCode(w http.ResponseWriter, msg string, status int) {
	jsonResp := constructBody(msg)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonResp)
}

func Send400(w http.ResponseWriter, msg string) {
	sendCode(w, msg, http.StatusBadRequest)
}

func Send401(w http.ResponseWriter, msg string) {
	jsonResp := constructBody(msg)
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("WWW-Authenticate", magicExcelHeader)
	w.WriteHeader(http.StatusUnauthorized)
	w.Write(jsonResp)
}

func Send403(w http.ResponseWriter, msg string) {
	sendCode(w, msg, http.StatusForbidden)
}

func Send404(w http.ResponseWriter, msg string) {
	sendCode(w, msg, http.StatusNotFound)
}

func Send500(w http.ResponseWriter, msg string) {
	sendCode(w, msg, http.StatusInternalServerError)
}

func Send502(w http.ResponseWriter, msg string) {
	sendCode(w, msg, http.StatusBadGateway)
}
