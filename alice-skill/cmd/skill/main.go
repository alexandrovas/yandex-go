// пакеты исполняемых приложений должны называться main
package main

import (
	"alice-skill/internal/logger"
	"alice-skill/internal/store/pg"
	"database/sql"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
)

// функция main вызывается автоматически при запуске приложения
func main() {
	// обрабатываем аргументы командной строки
	parseFlags()

	if err := run(); err != nil {
		panic(err)
	}
}

// функция run будет полезна при инициализации зависимостей сервера перед запуском
func run() error {
	if err := logger.Initialize(flagLogLevel); err != nil {
		return err
	}

	// создаём соединение с СУБД PostgreSQL с помощью аргумента командной строки
	conn, err := sql.Open("pgx", flagDatabaseURI)
	if err != nil {
		return err
	}

	// создаём экземпляр приложения, передавая реализацию хранилища pg в качестве внешней зависимости
	appInstance := newApp(pg.NewStore(conn))

	logger.Log.Info("Running server", zap.String("address", flagRunAddr))
	// обернём хендлер webhook в middleware с логированием и поддержкой gzip
	return http.ListenAndServe(flagRunAddr, logger.RequestLogger(gzipMiddleware(appInstance.webhook)))
}
