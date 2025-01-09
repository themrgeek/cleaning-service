package model

import (
	"context"
	"errors"
)

type Cleaner struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	// Add other fields as necessary
}

var cleaners = []Cleaner{
	{ID: 1, Name: "John Doe"},
	{ID: 2, Name: "Jane Smith"},
	// Add more cleaners as necessary
}

func GetCleaners(ctx context.Context) ([]Cleaner, error) {
	if len(cleaners) == 0 {
		return nil, errors.New("no cleaners found")
	}
	return cleaners, nil
}

func GetCleanerByID(ctx context.Context, id int) (Cleaner, error) {
	for _, cleaner := range cleaners {
		if cleaner.ID == id {
			return cleaner, nil
		}
	}
	return Cleaner{}, errors.New("cleaner not found")
}
func CreateCleaner(ctx context.Context, cleaner Cleaner) (Cleaner, error) {
	// Check if cleaner with the same ID already exists
	for _, c := range cleaners {
		if c.ID == cleaner.ID {
			return Cleaner{}, errors.New("cleaner with this ID already exists")
		}
	}
	cleaners = append(cleaners, cleaner)
	return cleaner, nil
}

func UpdateCleaner(ctx context.Context, id int, updatedCleaner Cleaner) (Cleaner, error) {
	for i, cleaner := range cleaners {
		if cleaner.ID == id {
			cleaners[i] = updatedCleaner
			return updatedCleaner, nil
		}
	}
	return Cleaner{}, errors.New("cleaner not found")
}

func DeleteCleaner(ctx context.Context, id int) error {
	for i, cleaner := range cleaners {
		if cleaner.ID == id {
			cleaners = append(cleaners[:i], cleaners[i+1:]...)
			return nil
		}
	}
	return errors.New("cleaner not found")
}
