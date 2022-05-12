package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) Get(id int) (*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, title, description, year, release_date, runtime, rating,
										mpaa_rating, created_at, updated_at
						FROM movies 
						WHERE id = $1;
	`
	row := m.DB.QueryRowContext(ctx, query, id)

	var movie Movie

	err := row.Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.Year,
		&movie.ReleaseDate,
		&movie.Runtime,
		&movie.Rating,
		&movie.MPAARating,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	query = `SELECT mg.id, mg.movie_id, mg.genre_id, g.genre_name
						FROM movies_genres mg
						INNER JOIN genres g
						ON mg.genre_id = g.id
						WHERE movie_id = $1;
	`
	rows, err := m.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	mgs := []MovieGenre{}
	for rows.Next() {
		mg := MovieGenre{}
		err = rows.Scan(
			&mg.ID,
			&mg.MovieID,
			&mg.GenreID,
			&mg.Genre.GenreName,
		)
		if err != nil {
			return nil, err
		}
		mgs = append(mgs, mg)
	}

	movie.MovieGenre = mgs

	return &movie, nil
}

func (m *DBModel) GetAll() ([]*Movie, error) {
	return nil, nil
}
