package postgres

import (
	"todo/models"
	"todo/storage"

	"github.com/jmoiron/sqlx"
)

type userRoleRepo struct {
	db *sqlx.DB
}

func NewUserRoleRepo(db *sqlx.DB) storage.UserRoleRepoI {
	return userRoleRepo{
		db: db,
	}
}

func (r userRoleRepo) GetUserRoles(userID string) ([]models.UserRole, error) {
	var roles []models.UserRole
	query := `SELECT ur.role_id, ur.role_name
			  FROM user_roles_users uru
			  JOIN user_roles ur ON uru.role_id = ur.role_id
			  WHERE uru.user_id = $1`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var role models.UserRole
		err := rows.Scan(&role.RoleID, &role.RoleName)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

func (r userRoleRepo) CreateUserRole(userID string, roleID int) error {
	query := `INSERT INTO user_roles_users (user_id, role_id) VALUES ($1, $2)`
	_, err := r.db.Exec(query, userID, roleID)
	return err
}

func (r userRoleRepo) DeleteUserRole(userID string, roleID int) error {
	query := `DELETE FROM user_roles_users WHERE user_id = $1 AND role_id = $2`
	_, err := r.db.Exec(query, userID, roleID)
	return err
}

func (r userRoleRepo) GetByID(ID string) (models.User, error) {
	var user models.User
	err := r.db.QueryRow(
		`SELECT id, full_name,
		user_name,
		email,
		user_password from users where id = $1
		`,
		ID,
	).Scan(&user.Id, &user.FullName, &user.UserName, &user.Email, &user.UserPassword)

	if err != nil {
		return user, err
	}

	return user, nil
}
