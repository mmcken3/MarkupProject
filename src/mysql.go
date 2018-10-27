package src

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Mysql is the interface for interacting with a mysql db.
type Mysql interface {
	StoreScore(string, int)
	ScoresForID(string)
	FirstAndLastID(int)
	ScoresInRange(string, string)
	AvgScore()
}

// LocalMysql implements Mysql.
type LocalMysql struct {
}

// Returns a new isntance of a LocalMysql struct.
func NewLocalMysql() Mysql {
	return &LocalMysql{}
}

// StoreScore stores the given score and id into the db.
func (l LocalMysql) StoreScore(id string, score int) {
	// Split the unique id into the id, and date.
	splitID := strings.SplitAfterN(id, "_", 2)
	onlyId := strings.Split(splitID[0], "_")[0]
	idDate := splitID[1]

	// Connect to and test connection to the db.
	db, err := sql.Open("mysql", "root@/markup_scores")
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Run the insert query on the database to insert or update the row.
	_, err = db.Exec("insert into scores (unq_id, id, id_date, score) values (?, ?, ?, ?) ON DUPLICATE KEY UPDATE unq_id=?, id=?, id_date=?, score=?",
		id,
		onlyId,
		idDate,
		score,
		id,
		onlyId,
		idDate,
		score,
	)

	if err != nil {
		log.Fatal(err)
	}

	db.Close()
}

// AvgScore gets the average score of all runs in the database.
func (l LocalMysql) AvgScore() {
	// Connect to and test the connection to the db.
	db, err := sql.Open("mysql", "root@/markup_scores")
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	var (
		id    string
		score string
	)

	// This query gets the average of each id's scores in the db.
	rows, err := db.Query("select id, avg(score) from scores group by id")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	count := 0

	// Sort through the rows of data.
	for rows.Next() {
		err := rows.Scan(&id, &score)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %v  \tAverage Score: %v\n", id, score)
		count += 1
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// If there was no count, then there is no data yet.
	if count == 0 {
		fmt.Println("No data has been scored yet!")
	}
}

// ScoresInRange prints all of the scores between date d1, and date d2.
func (l LocalMysql) ScoresInRange(d1 string, d2 string) {
	// Connect to and test connection to the db.
	db, err := sql.Open("mysql", "root@/markup_scores")
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	var (
		unq_id string
		date   string
		score  string
	)

	// This query gets the scores between the passed dates.
	rows, err := db.Query("select unq_id, run_time, score from scores where run_time > ? and run_time < ?", d1, d2)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	count := 0

	// Sort through each row of data.
	for rows.Next() {
		err := rows.Scan(&unq_id, &date, &score)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Date: %v\tScore: %v\tID: %v\n", date, score, unq_id)
		count += 1
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// If the count is 0, then no scores are in range so inform the user.
	if count == 0 {
		fmt.Printf("There were no scores in range %v to %v.\n", d1, d2)
	}
}

// ScoresForID returns all of the scores for a unique id.
func (l LocalMysql) ScoresForID(id string) {
	// Connect to and test the connection to the db.
	db, err := sql.Open("mysql", "root@/markup_scores")
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	var (
		date  string
		score string
	)

	// This query gets the scores where the unq_id is equal to the param.
	rows, err := db.Query("select run_time, score from scores where unq_id=?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	count := 0

	// Sort through the rows of data.
	for rows.Next() {
		err := rows.Scan(&date, &score)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Date: %v\tScore: %v\n", date, score)
		count += 1
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// If no rows found, inform the user about this.
	if count == 0 {
		fmt.Printf("There were no rows matching that id: %v\n", id)
	}
}

// FirstAndLastID returns the id of the highest score if n is greater than 0.
// It returns the id of the lowest score if n is 0.
func (l LocalMysql) FirstAndLastID(n int) {
	// Connect to and test connection to the db.
	db, err := sql.Open("mysql", "root@/markup_scores")
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	var (
		unid string
		rows *sql.Rows
	)
	var highLow string
	// Query for highest or lowest score based on the n.
	if n > 0 {
		highLow = "Highest"
		rows, err = db.Query("select unq_id from scores order by score desc limit 1")
	} else {
		highLow = "Lowest"
		rows, err = db.Query("select unq_id from scores order by score limit 1")
	}

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	count := 0

	// Sort through the rows of data.
	for rows.Next() {
		err := rows.Scan(&unid)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v Scored ID: %v\n", highLow, unid)
		count += 1
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// If there was no count, then there is no data yet.
	if count == 0 {
		fmt.Println("No data has been scored yet!")
	}
}
