package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) Get(id int) (*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, title, description, year, release_date, runtime, rating,
										mpaa_rating, created_at, updated_at, coalesce(poster, '') 
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
		&movie.Poster,
	)
	if err != nil {
		return nil, err
	}

	genres, err := m.getMovieGenres(movie.ID)
	if err != nil {
		return nil, err
	}

	movie.MovieGenre = genres

	return &movie, nil
}

func (m *DBModel) GetAll(genres ...int) ([]*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	where := ""
	if len(genres) > 0 {
		where = fmt.Sprintf(`WHERE id IN 
														(SELECT movie_id 
														FROM movies_genres 
														WHERE genre_id = %d)`, genres[0])
	}

	query := fmt.Sprintf(`SELECT id, title, description, year, release_date, runtime, rating,
										mpaa_rating, created_at, updated_at, coalesce(poster, '') 
						FROM movies %s 
						ORDER BY title;`, where)
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*Movie
	for rows.Next() {
		var movie Movie
		err = rows.Scan(
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
			&movie.Poster,
		)
		if err != nil {
			return nil, err
		}

		genres, err := m.getMovieGenres(movie.ID)
		if err != nil {
			return nil, err
		}

		movie.MovieGenre = genres
		movies = append(movies, &movie)
	}

	return movies, nil
}

func (m *DBModel) AllGenres() ([]*Genre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, genre_name, created_at, updated_at 
						FROM genres;`
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var genres []*Genre
	for rows.Next() {
		var genre Genre
		err = rows.Scan(
			&genre.ID,
			&genre.GenreName,
			&genre.CreatedAt,
			&genre.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		genres = append(genres, &genre)
	}

	return genres, nil
}

func (m *DBModel) getMovieGenres(movieId int) (map[int]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT mg.genre_id, g.genre_name
						FROM movies_genres mg
							INNER JOIN genres g
							ON mg.genre_id = g.id
						WHERE movie_id = $1;
	`
	rows, err := m.DB.QueryContext(ctx, query, movieId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	genres := make(map[int]string)
	for rows.Next() {
		mg := MovieGenre{}
		err = rows.Scan(
			&mg.GenreID,
			&mg.Genre.GenreName,
		)
		if err != nil {
			return nil, err
		}
		genres[mg.GenreID] = mg.Genre.GenreName
	}

	return genres, nil
}

func (m *DBModel) InsertMovie(movie Movie) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := `INSERT INTO movies (title, description, year, release_date, 
									runtime, rating, mpaa_rating, created_at, updated_at, poster) 
								VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`
	_, err := m.DB.ExecContext(ctx, statement,
		movie.Title,
		movie.Description,
		movie.Year,
		movie.ReleaseDate,
		movie.Runtime,
		movie.Rating,
		movie.MPAARating,
		movie.CreatedAt,
		movie.UpdatedAt,
		movie.Poster,
	)
	if err != nil {
		return err
	}

	return nil
}

func (m *DBModel) UpdateMovie(movie Movie) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := `UPDATE movies 
								SET title = $1, 
									description = $2, 
									year = $3, 
									release_date = $4, 
									runtime = $5, 
									rating = $6, 
									mpaa_rating = $7, 
									updated_at = $8,
									poster = $9 
								WHERE id = $10`
	_, err := m.DB.ExecContext(ctx, statement,
		movie.Title,
		movie.Description,
		movie.Year,
		movie.ReleaseDate,
		movie.Runtime,
		movie.Rating,
		movie.MPAARating,
		movie.UpdatedAt,
		movie.Poster,
		movie.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (m *DBModel) DeleteMovie(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := `DELETE FROM movies 
								WHERE id = $1`
	_, err := m.DB.ExecContext(ctx, statement, id)
	if err != nil {
		return err
	}

	return nil
}
