package main

func main() {
	ctx := context.Background()
	mysqlURL := "root:hunter2@unix(/tmp/mysql.sock)/myCoolDB"
	ctx = db.OpenSQL(ctx, "main", "mysql", mysqlURL) // HL
	defer db.Close(ctx)                              // closes all DB connections // HL
	kami.Context = ctx                               // HL
	kami.Get("/hello/:name", hello)
	kami.Serve()
}

func hello(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	mainDB := db.SQL(ctx, "main") // *sql.DB // HL
	var greeting string
	mainDB.QueryRow("SELECT content FROM greetings WHERE name = ?", kami.Param(ctx, "name")).
		Scan(&greeting)
	fmt.Fprintf(w, "Hello, %s!", greeting)
}
