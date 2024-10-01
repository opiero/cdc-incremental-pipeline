package debezium

func config() map[string]any {
	requestBody := map[string]any{
		"name": "postgres-connector",
		"config": map[string]string{
			"connector.class":      "io.debezium.connector.postgresql.PostgresConnector",
			"tasks.max":            "1",
			"database.hostname":    "postgres",
			"database.port":        "5432",
			"database.user":        "debezium_user",
			"database.password":    "password",
			"database.dbname":      "university",
			"database.server.name": "university_students",
			"table.include.list":   "public.customers",
			"plugin.name":          "pgoutput",
			"slot.name":            "debezium_university_students",
			"publication.name":     "students_publication ",
		},
	}
	return requestBody
}
