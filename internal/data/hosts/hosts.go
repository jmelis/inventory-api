package hosts

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/go-kratos/kratos/v2/log"

	authzapi "github.com/project-kessel/inventory-api/internal/authz/api"
	models "github.com/project-kessel/inventory-api/internal/biz/hosts"
	eventingapi "github.com/project-kessel/inventory-api/internal/eventing/api"
	"github.com/project-kessel/inventory-api/internal/middleware"
)

type hostsRepo struct {
	Db      *gorm.DB
	Authz   authzapi.Authorizer
	Eventer eventingapi.Manager

	Log *log.Helper
}

func New(g *gorm.DB, a authzapi.Authorizer, e eventingapi.Manager, l *log.Helper) *hostsRepo {
	return &hostsRepo{
		Db:      g,
		Authz:   a,
		Eventer: e,

		Log: l,
	}
}

func (r *hostsRepo) Save(ctx context.Context, model *models.Host) (*models.Host, error) {
	identity, err := middleware.GetIdentity(ctx)
	if err != nil {
		return nil, nil
	}

	if err := r.Db.Session(&gorm.Session{FullSaveAssociations: true}).Create(model).Error; err != nil {
		return nil, err
	}

	if r.Eventer != nil {
		// TODO: handle eventing errors
		// TODO: Update the Object that's sent.  This is going to be what we actually emit.
		producer, _ := r.Eventer.Lookup(identity, "rhelhost", model.ID)
		evt := &eventingapi.Event{
			EventType:    "Create",
			ResourceType: "rhelhost",
			Object:       model,
		}
		producer.Produce(ctx, evt)
	}
	return model, nil
}

func (r *hostsRepo) Update(context.Context, *models.Host, string) (*models.Host, error) {
	return nil, nil
}

func (r *hostsRepo) Delete(context.Context, string) error {
	return nil
}

func (r *hostsRepo) FindByID(context.Context, string) (*models.Host, error) {
	return nil, nil
}

func (r *hostsRepo) ListAll(context.Context) ([]*models.Host, error) {
	// var model models.Host
	// var count int64
	// if err := r.Db.Model(&model).Count(&count).Error; err != nil {
	// 	return nil, err
	// }

	var results []*models.Host
	if err := r.Db.Preload(clause.Associations).Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}