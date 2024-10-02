package debezium

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func config() ([]byte, error) {
	requestBody := map[string]any{
		"name": "postgres-connector",
		"config": map[string]string{
			"connector.class":      "io.debezium.connector.postgresql.PostgresConnector",
			"topic.prefix":         "dbz",
			"tasks.max":            "1",
			"database.hostname":    "postgres",
			"database.port":        "5432",
			"database.user":        "debezium_user",
			"database.password":    "password",
			"database.dbname":      "university",
			"database.server.name": "university_students",
			"plugin.name":          "pgoutput",
			"slot.name":            "debezium_university_students",
			"publication.name":     "students_publication",
		},
	}
	return json.Marshal(requestBody)
}

func ConnectToDebezium() error {
	jsonConfig, err := config()
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(
		"http://localhost:8083/connectors/",
		"application/json",
		bytes.NewBuffer(jsonConfig),
	)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

	return err
}
