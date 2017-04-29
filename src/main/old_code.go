package main

/*
const rowsAtATime = 25
const timeout = 600
url := constants.InsiderTransactionsUrl+"&issuetickers=AMD&transactiondates=20160221~20170221"

	result := getResult(url)
	totalRows := result.TotalRows
	numberOfCalls := totalRows / rowsAtATime
	transactionCollection := make([]jsons.Transaction, totalRows)
	for i, t := range jsons.ToTransactions(result.RowList) {
		transactionCollection[i] = t
	}

	for i := 1; i < numberOfCalls; i++ {
		amt := i * 25
		offset := fmt.Sprintf("&offset=%d", amt)
		result := getResult(url + offset)
		for j, t := range jsons.ToTransactions(result.RowList) {
			index := amt+j
			transactionCollection[index] = t
		}
		time.Sleep(time.Duration(timeout)*time.Millisecond)
	}
	amtAlreadyCalled := numberOfCalls * rowsAtATime
	remainder := totalRows - amtAlreadyCalled
	if remainder > 0 {
		offset := fmt.Sprintf("&offset=%d", amtAlreadyCalled)
		limit := fmt.Sprintf("&limit=%d", remainder)
		result := getResult(url + offset + limit)
		for i, t := range jsons.ToTransactions(result.RowList) {
			index := amtAlreadyCalled+i
			transactionCollection[index] = t
		}
	}
	fmt.Printf("%+v\n", transactionCollection)
	session := db.GetSession()
	db.DeleteAllDocuments(session, constants.TransactionsCollection)
	db.InsertTransactionCollection(transactionCollection[:], session)
	docs := db.GetAllDocuments(session, constants.TransactionsCollection)
	fmt.Println("docs")
	fmt.Println(docs)
	session.Close()

	router := mux.NewRouter()
	routes.People = append(routes.People, routes.Person{ID: "1", Firstname: "Nic",
		Lastname: "Raboy", Address: &routes.Address{City: "Dublin", State: "CA"}})
	routes.People = append(routes.People, routes.Person{ID: "2", Firstname: "Maria", Lastname: "Robby"})
	router.HandleFunc("/transactions", routes.GetTransactions).Methods("GET")
	router.HandleFunc("/transactions/{tickers}", routes.GetTransactions).Methods("GET")
	router.HandleFunc("/transactions/{tickers}/{sort:(?:asc|desc)}", routes.GetTransactions).Methods("GET")
	log.Fatal(http.ListenAndServe(":12345", router))


func getResult(url string) jsons.Result {
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		println("reached")
		log.Fatal("New Request: ", err)
		return jsons.Result{}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		println("reached")
		log.Fatal("Do: ", err)
		return jsons.Result{}
	}
	defer resp.Body.Close()
	var jsonResult jsons.Response
	if err := json.NewDecoder(resp.Body).Decode(&jsonResult); err != nil {
		println("reached")
		log.Println(err)
	}
	fmt.Println(jsonResult)
	return jsonResult.Result
}

	*/
