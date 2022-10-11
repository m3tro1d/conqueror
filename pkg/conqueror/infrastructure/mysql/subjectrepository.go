package mysql

import (
	"context"

	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/domain"
	"github.com/jmoiron/sqlx"
)

func NewSubjectRepository(ctx context.Context, client *sqlx.Conn) domain.SubjectRepository {
	return &subjectRepository{
		ctx:    ctx,
		client: client,
	}
}

type subjectRepository struct {
	ctx    context.Context
	client *sqlx.Conn
}

func (repo *subjectRepository) NextID() domain.SubjectID {
	return domain.SubjectID(uuid.Generate())
}

func (repo *subjectRepository) Store(subject *domain.Subject) error {
	//TODO implement me
	panic("implement me")
}

func (repo *subjectRepository) GetByID(id domain.SubjectID) (*domain.Subject, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *subjectRepository) RemoveByID(id domain.SubjectID) error {
	//TODO implement me
	panic("implement me")
}
