package register

import (
	"context"
	"fmt"

	"github.com/teq-quocbang/store/codetype"
	"github.com/teq-quocbang/store/model"
	"github.com/teq-quocbang/store/util/myerror"
	"gorm.io/gorm"
)

type pgRepository struct {
	getDB func(context.Context) *gorm.DB
}

func NewClassPG(getDB func(context.Context) *gorm.DB) Repository {
	return &pgRepository{getDB: getDB}
}

func (p *pgRepository) Create(ctx context.Context, r *model.Register) error {
	db := p.getDB(ctx)
	tx := db.Begin()
	//  create register
	err := tx.Create(&r).Error
	if err != nil {
		return myerror.ErrRegisterCreate(err)
	}

	// increase member
	class := model.Class{}
	err = tx.Model(&class).Where(`id = ?`, r.ClassID).Update("current_slot", gorm.Expr("current_slot + 1")).Take(&class).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//  check member after increase
	if class.CurrentSlot > class.MaxSlot {
		tx.Rollback()
		return fmt.Errorf("class was max slot")
	}
	tx.Commit()
	return nil
}

func (p *pgRepository) GetListBySemesterID(
	ctx context.Context,
	accountID uint,
	semesterID string,
	order []string,
	paginator codetype.Paginator,
) ([]model.Register, int64, error) {
	var (
		register []model.Register
		db       = p.getDB(ctx)
		total    int64
		offset   int
	)

	for _, o := range order {
		db = db.Order(o)
	}

	if paginator.Limit != -1 {
		if err := db.Model(&model.Register{}).Count(&total).Error; err != nil {
			return nil, 0, err
		}
	}

	if paginator.Page != 0 {
		offset = paginator.Limit * (paginator.Page - 1)
	}

	if err := db.Model(&model.Register{}).Where(`semester_id = ? and account_id = ? and is_canceled = false`, semesterID, accountID).Offset(offset).Limit(paginator.Limit).Find(&register).Error; err != nil {
		return nil, 0, err
	}
	return register, total, nil
}

func (p *pgRepository) Get(ctx context.Context, req *model.Register) (*model.Register, error) {
	var register *model.Register
	err := p.getDB(ctx).Where(`account_id = ? and semester_id = ? and class_id = ? and course_id = ?`,
		req.AccountID,
		req.SemesterID,
		req.ClassID,
		req.CourseID).Take(&register).Error
	return register, err
}

// swap the state of the is_canceled field
// false -> true and true -> false
func (p *pgRepository) BatchUpdateSwapIsCanCeledStatus(ctx context.Context, req *model.Register) error {
	db := p.getDB(ctx)
	tx := db.Begin()
	if req.IsCanceled {
		//  update is_canceled == false
		err := tx.Model(&model.Register{}).Where(`account_id = ? and class_id = ? and semester_id = ?`, req.AccountID, req.ClassID, req.SemesterID).
			Update("is_canceled", gorm.Expr(" !is_canceled")).Error
		if err != nil {
			tx.Rollback()
			return myerror.ErrRegisterUpdate(err)
		}
		// update increase class member
		class := model.Class{}
		err = tx.Model(&class).Where(`id = ?`, req.ClassID).
			Update("current_slot", gorm.Expr("current_slot + 1")).Take(&class).Error
		if err != nil {
			tx.Rollback()
			return err
		}
		//  check member after increase
		if class.CurrentSlot > class.MaxSlot {
			tx.Rollback()
			return fmt.Errorf("class was max slot")
		}
		tx.Commit()
	} else {
		//  update is_canceled == true
		err := tx.Model(&model.Register{}).Where(`account_id = ? and class_id = ? and semester_id = ?`, req.AccountID, req.ClassID, req.SemesterID).
			Update("is_canceled", gorm.Expr(" !is_canceled")).Error
		if err != nil {
			tx.Rollback()
			return myerror.ErrRegisterUpdate(err)
		}
		// update decrease class member
		err = p.getDB(ctx).Model(&model.Class{}).Where(`id = ?`, req.ClassID).
			Update("current_slot", gorm.Expr(" current_slot - 1")).Error
		if err != nil {
			tx.Rollback()
			return myerror.ErrClassUpdate(err)
		}
		tx.Commit()
	}
	return nil
}

// GetListByFirstCourseChar is get list all the course that student registered
// use the first character of course_id
// ex:
//
//	student registered S0001, T0001, M0001
//
// with S in param so the result is:
//
//	[S0001]
func (p *pgRepository) GetListByFirstCourseChar(ctx context.Context, firstChar string, accountID uint, semesterID string) ([]model.Register, error) {
	var registers []model.Register
	err := p.getDB(ctx).Where(`account_id = ? and substring(course_id, 1, 1) = ? and semester_id = ?`, accountID, firstChar, semesterID).Find(&registers).Error
	return registers, err
}

func (p *pgRepository) GetListRegistered(
	ctx context.Context,
	accountID uint,
	semesterID string,
	order []string,
	paginator codetype.Paginator,
) ([]model.Register, int64, error) {
	var (
		registers = []model.Register{}
		db        = p.getDB(ctx)
		total     int64
		offset    int
	)

	for i := range order {
		db = db.Order(order[i])
	}

	if semesterID == "" {
		db = db.Where(`account_id = ?`, accountID)
	} else {
		db = db.Where(`account_id = ? and semester_id = ?`, accountID, semesterID)
	}

	if paginator.Limit != -1 {
		err := db.Model(&model.Register{}).Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
	}

	if paginator.Limit != 0 {
		offset = paginator.Limit * (paginator.Page - 1)
	}

	if offset != 0 {
		db = db.Offset(offset).Limit(paginator.Limit)
	}

	if err := db.Model(&model.Register{}).Find(&registers).Error; err != nil {
		return nil, 0, err
	}

	return registers, total, nil
}
