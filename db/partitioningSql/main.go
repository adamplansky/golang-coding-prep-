package main

import (
	"fmt"
	"time"
)

const timeFormat string = "2006-01-02 15:04:05"

func main() {
	//var schema = "cezedd"
	//var table = "event"
	var schema = "cezedd"
	var table = "event"
	var tableName = table
	if schema != "" {
		tableName = fmt.Sprintf("%s.%s", schema, tableName)
	}

	var startingYear, endingYear int
	//fmt.Println("enter db table to partition:")
	//fmt.Scanf("%s", &tableName)
	//fmt.Println("enter starting year:")
	//fmt.Scanf("%d", &startingYear)
	//fmt.Println("enter ending year:")
	//fmt.Scanf("%d", &endingYear)
	startingYear = 2000
	endingYear = 2100

	//insertFunction := fmt.Sprintf("CREATE OR REPLACE FUNCTION %s_insert_trigger()\n RETURNS TRIGGER AS\n $$ BEGIN\n", tableName)
	//var idx int
	for year := startingYear; year <= endingYear; year++ {

		startTime := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
		endTime := startTime.AddDate(1, 0, 0)
		end := endTime.Format(timeFormat)
		start := startTime.Format(timeFormat)

		CreatePartitionTable(start, end, tableName, table, year)
		//if idx != 0 {
		//	insertFunction += fmt.Sprintf("ELSIF ( NEW.created_at >= '%s' AND NEW.created_at < '%s' )\n THEN INSERT INTO %s_%d VALUES (NEW.*);\n", start, end, tableName, year)
		//} else {
		//	insertFunction += fmt.Sprintf("IF ( NEW.created_at >= '%s' AND NEW.created_at < '%s' )\n THEN INSERT INTO %s_%d VALUES (NEW.*);\n", start, end, tableName, year)
		//}
		//idx++

	}

	//insertFunction += fmt.Sprintf("ELSE RAISE EXCEPTION 'Date out of range.  Fix the %s_insert_trigger() function!';", tableName)
	//insertFunction += fmt.Sprintf("\nEND IF;\n RETURN NULL;\nEND;\n$$\nLANGUAGE plpgsql;")
	//fmt.Println("\n", insertFunction)
	//
	//trigger := fmt.Sprintf("\nCREATE TRIGGER insert_%s_trigger\n    BEFORE INSERT ON %s\n    FOR EACH ROW EXECUTE PROCEDURE %s_insert_trigger();", tableName, tableName, tableName)
	//fmt.Println(trigger)
}

func CreatePartitionTable(startTime, endTime, tableName string, table string, year int) {
	//fmt.Printf("CREATE TABLE %s_%d PARTITION %s FOR VALUES FROM created_at >= '%s' TO created_at < '%s'\n", tableName, year, tableName, startTime, endTime)
	fmt.Printf("CREATE TABLE %s_%d PARTITION OF %s FOR VALUES FROM ('%s') TO ('%s'); \n", tableName, year, table, startTime, endTime)
	//CREATE TABLE users_a_to_i PARTITION OF users FOR VALUES FROM ('a') TO ('j’);
}
