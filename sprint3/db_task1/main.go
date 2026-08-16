package main

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "video.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// row := db.QueryRowContext(context.Background(),
	// 	"SELECT title, likes, comments_disabled "+
	// 		"FROM videos ORDER BY likes DESC LIMIT 1")
	// var (
	// 	title  string
	// 	likes  int
	// 	comdis bool
	// )
	// // порядок переменных должен соответствовать порядку колонок в запросе
	// err = row.Scan(&title, &likes, &comdis)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s | %d | %t \r\n", title, likes, comdis)v
	// value, err := getDesc(context.Background(), db, "0EbFotkXOiA")
	// fmt.Printf("value='%s', err='%s'\n", value, err)

	// err = printTopVideo(context.Background(), db)
	// if err != nil {
	// 	panic(err)
	// }

	trends, err := QueryTrends(context.Background(), db)
	if err != nil {
		panic(err)
	}
	for _, t := range trends {
		fmt.Printf("%+v\n", t)
	}
}

func getDesc(ctx context.Context, db *sql.DB, id string) (string, error) {
	row := db.QueryRowContext(ctx,
		"SELECT description FROM videos WHERE video_id = ?", id)
	var desc sql.NullString

	err := row.Scan(&desc)
	if err != nil {
		return "", err
	}
	if desc.Valid {
		return desc.String, nil
	}
	return "-----", nil
}

func printTopVideo(ctx context.Context, db *sql.DB) error {
	row := db.QueryRowContext(ctx,
		"select title, views, channel_title from videos order by views desc limit 1")
	var (
		title         string
		views         int
		channel_title string
	)

	err := row.Scan(&title, &views, &channel_title)
	if err != nil {
		return err
	}

	fmt.Printf("description: %s, views: %d, channel_title: %s", title, views, channel_title)
	return nil
}

type Trend struct {
	T     TrendDate
	Count int
}

type TrendDate struct {
	time.Time
}

// Value — функция, реализующая интерфейс driver.Valuer
func (td TrendDate) Value() (driver.Value, error) {
	// преобразуем time.Time в string
	if td.IsZero() {
		return "", nil
	}
	return td.Format("06.02.01"), nil
}

// Scan — функция, реализующая интерфейс sql.Scanner
func (td *TrendDate) Scan(value interface{}) error {
	// если `value` равен `nil`, будет возвращена пустая дата
	if value == nil || value == "" {
		*td = TrendDate{time.Time{}}
		return nil
	}

	sv, err := driver.String.ConvertValue(value)
	if err != nil {
		return fmt.Errorf("cannot scan value. %w", err)
	}

	v, ok := sv.(string)
	if !ok {
		return errors.New("cannot scan value. cannot convert value to string")
	}

	d, err := time.Parse("06.02.01", v)
	if err != nil {
		return fmt.Errorf("cannot scan value. cannot parse date: %w", err)
	}

	*td = TrendDate{d}

	return nil
}

func QueryTrends(ctx context.Context, db *sql.DB) ([]Trend, error) {
	trends := make([]Trend, 0)

	rows, err := db.QueryContext(ctx,
		"SELECT trending_date, COUNT(trending_date) FROM videos GROUP BY trending_date")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t Trend
		err = rows.Scan(&t.T, &t.Count)
		if err != nil {
			return nil, err
		}
		trends = append(trends, t)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return trends, nil
}
