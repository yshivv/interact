package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type UserSession struct {
	URL        string
	LastAccess time.Time
}

var (
	sessions map[string]*UserSession
	mu       sync.Mutex
)

func executeInteractshClient(command string) (string, error) {
	cmd := exec.Command("interactsh-client", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func getURLHandler(c *gin.Context) {
	token := c.GetHeader("Authorization")
	mu.Lock()
	defer mu.Unlock()
	session, ok := sessions[token]
	if !ok {
		session = &UserSession{}
		sessions[token] = session
	}
	session.LastAccess = time.Now()

	output, err := executeInteractshClient("")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	url := extractURL(output)
	session.URL = url

	// Return the URL with formatting
	c.String(http.StatusOK, "[INF] **%s**\n", url)
}

func extractURL(output string) string {
	parts := strings.Split(output, "\n")
	for _, part := range parts {
		if strings.Contains(part, "Listing 1 payload") {
			return strings.Fields(part)[2]
		}
	}
	return ""
}

func getInteractionsHandler(c *gin.Context) {
	token := c.GetHeader("Authorization")
	mu.Lock()
	defer mu.Unlock()
	session, ok := sessions[token]
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	output, err := executeInteractshClient(fmt.Sprintf("-server %s", session.URL))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	interactions := parseInteractions(output)

	// Format interactions for response
	formattedInteractions := "[INF] Interactions:\n"
	for _, interaction := range interactions {
		formattedInteractions += fmt.Sprintf("[c23b2la0kl1krjcrdj10cndmnioyyyyyn] %s\n", interaction)
	}

	// Return interactions with formatting
	c.String(http.StatusOK, formattedInteractions)
}

func parseInteractions(output string) []map[string]string {
	lines := strings.Split(output, "\n")
	var interactions []map[string]string
	for _, line := range lines {
		if strings.Contains(line, "Received HTTP interaction") {
			fields := strings.Fields(line)
			interaction := make(map[string]string)
			interaction["interaction_type"] = "HTTP"
			interaction["caller_ip"] = fields[6]
			interaction["timestamp"] = fields[len(fields)-3] + " " + fields[len(fields)-2]
			interactions = append(interactions, interaction)
		}
	}
	return interactions
}

func cleanupSessions() {
	for {
		time.Sleep(5 * time.Minute)
		mu.Lock()
		for token, session := range sessions {
			if time.Since(session.LastAccess) > 30*time.Minute {
				delete(sessions, token)
			}
		}
		mu.Unlock()
	}
}

func main() {
	sessions = make(map[string]*UserSession)
	go cleanupSessions()

	router := gin.Default()

	router.GET("/api/getURL", getURLHandler)
	router.GET("/api/getInteractions", getInteractionsHandler)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
