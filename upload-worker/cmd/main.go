package main

import (
	"net/http"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/infrastructure"
)

func main() {
	r := infrastructure.InitApp()

	http.ListenAndServe(":3009", r)
}
