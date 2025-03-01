package storage

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"movieWishlistAPI/cfg"
	"movieWishlistAPI/models"
)

type PostgreSqlDatabase struct {
	DB *sql.DB
}

func NewPostgresDB(config *cfg.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", config.PostgresURL)
	if err != nil {
		return nil, fmt.Errorf("Ошибка подключения к базе данных: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Связь с базой данных потеряна: %w", err)
	}
	return db, nil
}

func (p *PostgreSqlDatabase) SaveMovie(movie *models.Movie) error {
	query := `INSERT INTO movies (title,year,genre,director,actors,rating_imdb) VALUES ($1, $2, $3, $4, $5,$6) RETURNING id`
	return p.DB.QueryRow(query, movie.Title, movie.Year, movie.Genre, movie.Director, movie.Actors, movie.RatingIMDB).Scan(&movie.ID)
}

func (p *PostgreSqlDatabase) GetMovies() ([]*models.Movie, error) {
	rows, err := p.DB.Query("SELECT id,title,year,genre,director,actors,rating_imdb FROM movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var movies []*models.Movie
	for rows.Next() {

		var movie models.Movie
		if err := rows.Scan(&movie.Title, &movie.Title, &movie.Year, &movie.Genre, &movie.Director, &movie.Actors, &movie.RatingIMDB); err != nil {
			return nil, err
		}
		movies = append(movies, &movie)

	}
	return movies, nil
}
func (p *PostgreSqlDatabase) CloseDB() {
	p.DB.Close()
}
