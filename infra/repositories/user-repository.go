package repositories

import (
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rafaLino/couple-wishes-api/entities"
	"github.com/rafaLino/couple-wishes-api/infra/db"
	dbclient "github.com/rafaLino/couple-wishes-api/infra/db-client"
	"github.com/rafaLino/couple-wishes-api/ports"
	valueObjects "github.com/rafaLino/couple-wishes-api/value-objects"
)

type UserRepository struct {
	context *dbclient.DbContext
	ports.IUserRepository
}

func NewUserRepository(c *dbclient.DbContext) (ports.IUserRepository, error) {
	return &UserRepository{context: c}, nil
}

func (r *UserRepository) GetAll() ([]entities.User, error) {
	client, err := r.context.GetClient()
	usersRow, err := client.GetUsers(r.context.GetContext())

	users := entities.MapToUsers(usersRow)

	return users, err
}

func (r *UserRepository) Get(id int64) (*entities.User, error) {
	client, err := r.context.GetClient()
	userRow, err := client.GetUser(r.context.GetContext(), id)

	user := entities.MapGetUserRowToUser(userRow)

	return &user, err
}

func (r *UserRepository) CheckUsername(username string) (bool, error) {
	client, err := r.context.GetClient()
	count, err := client.CheckUserName(r.context.GetContext(), username)

	exists := count > 0

	return exists, err
}

func (r *UserRepository) CheckPassword(username valueObjects.Username, password valueObjects.Password) (*entities.User, error) {
	client, err := r.context.GetClient()
	user, err := client.GetUserWithPassword(r.context.GetContext(), username.String())

	if err != nil {
		return nil, err
	}

	match := password.Verify(user.Password)

	if !match {
		return nil, errors.New("password does not match")
	}

	var coupleID int64
	if user.CoupleID.Valid {
		coupleID = user.CoupleID.Int64
	} else {
		coupleID = 0
	}

	return &entities.User{
		ID:       user.ID,
		Name:     user.Name,
		Username: *valueObjects.NewUsername(user.Username),
		CoupleID: coupleID,
	}, nil
}

func (r *UserRepository) Create(user *entities.User) (int64, error) {
	client, err := r.context.GetClient()
	userId, err := client.CreateUser(r.context.GetContext(), db.CreateUserParams{
		Name:     user.Name,
		Username: user.Username.String(),
		Password: user.Password.Value(),
	})

	return userId, err
}

func (r *UserRepository) Update(user *entities.User) error {
	client, err := r.context.GetClient()
	client.UpdateUser(r.context.GetContext(), db.UpdateUserParams{
		ID:   user.ID,
		Name: user.Name,
	})

	return err
}

func (r *UserRepository) ChangePassword(id int64, password valueObjects.Password) error {
	client, err := r.context.GetClient()
	client.ChangePassword(r.context.GetContext(), db.ChangePasswordParams{
		ID:       id,
		Password: password.Value(),
	})
	return err
}

func (r *UserRepository) Delete(id int64) error {
	client, err := r.context.GetClient()
	client.DeleteUser(r.context.GetContext(), id)

	return err
}

func (r *UserRepository) GetPartnerUsername(id, userId int64) (string, error) {
	client, err := r.context.GetClient()
	partnerName, err := client.GetPartnerName(r.context.GetContext(), db.GetPartnerNameParams{
		ID:   id,
		ID_2: userId,
	})

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return "", nil
	}

	return partnerName, err
}

func (r *UserRepository) CreateCouple(userId, partnerId int64) (int64, error) {
	client, err := r.context.GetClient()

	id, err := client.CreateCouple(r.context.GetContext(), db.CreateCoupleParams{
		UserID:    pgtype.Int8{Int64: userId, Valid: true},
		PartnerID: pgtype.Int8{Int64: partnerId, Valid: true},
	})

	return id, err
}

func (r *UserRepository) DeleteCouple(coupleId int64) error {
	client, err := r.context.GetClient()
	client.DeleteCouple(r.context.GetContext(), coupleId)

	return err
}

func (r *UserRepository) GetByUsername(username valueObjects.Username) (*entities.User, error) {
	client, err := r.context.GetClient()
	userRow, err := client.GetUserByUsername(r.context.GetContext(), username.String())

	var coupleID int64
	if userRow.CoupleID.Valid {
		coupleID = userRow.CoupleID.Int64
	} else {
		coupleID = 0
	}

	user := &entities.User{
		ID:       userRow.ID,
		Name:     userRow.Name,
		Username: *valueObjects.NewUsername(userRow.Username),
		CoupleID: coupleID,
	}

	return user, err
}
