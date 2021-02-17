import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func connectTest() (db *sql.DB) {
	str := os.Getenv("MYSQLCONNSTR_localdb")
	var user, password, host string
	fields := strings.Split(str, ";")
	for _, v := range fields {
		v = strings.TrimSpace(v)
		if len(v) > 12 && v[:12] == "Data Source=" {
			host = v[12:]
		}
		if len(v) > 8 && v[:8] == "User Id=" {
			user = v[8:]
		}
		if len(v) > 9 && v[:9] == "Password=" {
			password = v[9:]
		}
	}
	// "username:password@tcp(127.0.0.1:3306)/test"
	db, err := sql.Open("mysql", user+":"+password+"@tcp("+host+")/localdb")
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func testDB(w http.ResponseWriter, req *http.Request) {
	t := time.Now()
	fmt.Fprintln(w, "starting...")
	db := connectTest()
	defer db.Close()
	_, err := db.Exec("insert into test (field1) values (?)", "test")
	fmt.Fprintln(w, "done. Took (ms):", time.Since(t).Milliseconds())
	if err != nil {
		fmt.Fprintln(w, "got error:", err)
	}
}
