package repo

import (
	"context"
	"database/sql"
	"fmt"
	"svc-insani-go/app"
)

// TODO: get all pengaturan
func GetPengaturan(a *app.App, ctx context.Context, atribut string) (string, error) {
	sqlQuery := getPengaturanQuery(atribut)
	var nilai string
	err := a.DB.QueryRowContext(ctx, sqlQuery).Scan(&nilai)
	if err == sql.ErrNoRows {
		return "", nil
	}
	if err != nil {
		return "", fmt.Errorf("error querying get pengaturan: %w", err)
	}
	return nilai, nil

}

func InsertPengaturan(a *app.App, ctx context.Context, atribut, nilai, userUpdate string) error {
	sqlQuery := insertPengaturanQuery()
	stmt, err := a.DB.PrepareContext(ctx, sqlQuery)
	defer stmt.Close()
	if err != nil {
		return fmt.Errorf("error prepare context: %w", err)
	}
	res, err := stmt.ExecContext(ctx, atribut, nilai, userUpdate, userUpdate)
	if err != nil {
		return fmt.Errorf("error exec context: %w", err)
	}
	affectedRows, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error get rows affected: %w", err)
	}
	if affectedRows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func UpdatePengaturan(a *app.App, ctx context.Context, atribut, nilai, userUpdate string) error {
	sqlQuery := updatePengaturanQuery()
	stmt, err := a.DB.PrepareContext(ctx, sqlQuery)
	defer stmt.Close()
	if err != nil {
		return fmt.Errorf("error prepare context: %w", err)
	}
	fmt.Printf("[DEBUG] nilai repo: %s\n", nilai)
	fmt.Printf("[DEBUG] userUpdate repo: %s\n", userUpdate)
	fmt.Printf("[DEBUG] atribut repo: %s\n", atribut)
	_, err = stmt.ExecContext(ctx, nilai, userUpdate, atribut)
	if err != nil {
		return fmt.Errorf("error exec context: %w", err)
	}
	return nil
}
