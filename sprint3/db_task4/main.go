package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "modernc.org/sqlite"
)

type TagVideo struct {
	Tags string
}

func getList(ctx context.Context, db *sql.DB) (videos []TagVideo, err error) {
	rows, err := db.QueryContext(ctx,
		"SELECT tags FROM videos WHERE tags LIKE '%best%' GROUP BY tags")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var video TagVideo

		if err = rows.Scan(&video.Tags); err != nil {
			return
		}
		videos = append(videos, video)
	}
	return
}

func main() {
	// открываем соединение с БД
	db, err := sql.Open("sqlite", "newvideo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx := context.Background()

	videos, err := getList(ctx, db)
	if err != nil {
		log.Fatal(err)
	}

	var updates int64 // подсчитаем количество изменённых записей
	for _, v := range videos {
		var tags []string
		// удаляем лишние теги
		for _, tag := range strings.Split(v.Tags, `|`) {
			if !strings.Contains(strings.ToLower(tag), "best") {
				tags = append(tags, tag)
			}
		}
		res, err := db.ExecContext(ctx,
			"UPDATE videos SET tags = ? WHERE tags = ?",
			strings.Join(tags, "|"), v.Tags)
		if err != nil {
			log.Fatal(err)
		}
		// посмотрим, сколько записей было обновлено
		if upd, err := res.RowsAffected(); err == nil {
			updates += upd
		}
	}
	fmt.Println(updates)
}
