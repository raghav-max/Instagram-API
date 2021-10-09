func getUser(w http.ResponseWriter, r *http.Request)
{
	var p Person
	ids, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || ids < 1 
	{
        	http.NotFound(w, r)
        	return
	}
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/instadb")
	result,err:=db.Query("SELECT id,name,password,email from users where id=ids")
	err = result.Scan(&p.Id, &p.Name, &p.Password, &p.Email)
	if err!=nul
	{
		panic(err.Error())
	}
	fmt.Fprintf(w,"Name=",p.Name)
	fmt.Fprintf(w,"Password=",p.Password)
	fmt.Fprintf(w,"Email=",p.Email)
	db.Close()
	
}