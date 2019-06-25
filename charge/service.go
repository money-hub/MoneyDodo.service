package charge

import (
	"context"
	"log"
	"time"

	"github.com/money-hub/MoneyDodo.service/db"
	"github.com/money-hub/MoneyDodo.service/model"
)

// BASE_URL=http://hostname:port/api/tasks

type ChargeService interface {
	// 查询所有充值记录
	GetAll(ctx context.Context, page, offset, limit int, orderby string) (status bool, errinfo string, data []model.Charge)
	// 查询某个充值记录
	GetSpec(ctx context.Context, chargeId string) (status bool, errinfo string, data interface{})
	// 充值操作
	Post(ctx context.Context, charge interface{}) (status bool, errinfo string, data interface{})
	// 删除充值记录
	Delete(ctx context.Context, chargeId string) (status bool, errinfo string, data *model.Charge)
	// 查询用户相关记录
	GetAllOfUser(ctx context.Context, userId string, page, offset, limit int, orderby string) (status bool, errinfo string, data []model.Charge)
}

type basicChargeService struct {
	*db.DBService
}

func (b *basicChargeService) GetAll(ctx context.Context, page, offset, limit int, orderby string) (status bool, errinfo string, data []model.Charge) {
	if ctx.Value("role") != 0 {
		return b.GetAllOfUser(ctx, ctx.Value("id").(string), page, offset, limit, orderby)
	}
	data = make([]model.Charge, 0)
	var err error
	err = b.Engine().Find(&data)
	if orderby == "-id" {
		for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
			data[i], data[j] = data[j], data[i]
		}
	}

	if limit > 0 {
		res := make([]model.Charge, 0)
		for i, charge := range data {
			if i >= offset && i < offset+page*limit {
				res = append(res, charge)
			}
		}

		status = err == nil
		if err != nil {
			errinfo = err.Error()
		}
		return status, errinfo, res
	}

	status = err == nil
	if err != nil {
		errinfo = err.Error()
	}
	return status, errinfo, data
}

func (b *basicChargeService) GetAllOfUser(ctx context.Context, userId string, page, offset, limit int, orderby string) (status bool, errinfo string, data []model.Charge) {
	if ctx.Value("role") != 0 {
		return false, "You are not administrator", nil
	}
	dataAll := make([]model.Charge, 0)
	var err error
	err = b.Engine().Find(&dataAll)
	if orderby == "-id" {
		for i, j := 0, len(dataAll)-1; i < j; i, j = i+1, j-1 {
			dataAll[i], dataAll[j] = dataAll[j], dataAll[i]
		}
	}

	for _, d := range dataAll {
		if d.UserId == userId {
			data = append(data, d)
		}
	}

	if limit > 0 {
		res := make([]model.Charge, 0)
		for i, charge := range data {
			if i >= offset && i < offset+page*limit {
				res = append(res, charge)
			}
		}

		status = err == nil
		if err != nil {
			errinfo = err.Error()
		}
		return status, errinfo, res
	}

	status = err == nil
	if err != nil {
		errinfo = err.Error()
	}
	return status, errinfo, data
}

func (b *basicChargeService) GetSpec(ctx context.Context, chargeId string) (status bool, errinfo string, data interface{}) {
	var err error
	charge := model.Charge{
		Id: chargeId,
	}
	status, err = b.Engine().Get(&charge)
	if status == false {
		return false, "The query charge is not existed.", nil
	}
	if err != nil {
		return false, err.Error(), nil
	}
	if charge.UserId != ctx.Value("id").(string) || ctx.Value("role") != 0 {
		return false, "You can not query charges info of others.", nil
	}
	return true, "", charge
}

func (b *basicChargeService) Post(ctx context.Context, charge interface{}) (status bool, errinfo string, data interface{}) {
	var err error
	c, ok := charge.(model.Charge)
	if !ok {
		return false, "The post data is not matching to the model.Charge.", nil
	}
	sess := b.Engine().NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return false, err.Error(), nil
	}
	c.Timestamp = time.Now()
	c.UserId = ctx.Value("id").(string)
	if _, err = sess.Insert(c); err != nil {
		sess.Rollback()
		return false, err.Error(), nil
	}
	lastCharge := &model.Charge{}
	if _, err = sess.Limit(1, 0).Desc("Id").Get(lastCharge); err != nil {
		sess.Rollback()
		return false, err.Error(), nil
	}
	user := &model.User{
		Id: lastCharge.UserId,
	}
	if _, err = sess.Get(user); err != nil {
		sess.Rollback()
		return false, err.Error(), nil
	}
	user.Balance = user.Balance + lastCharge.Amount
	if _, err = sess.Where("id = ?", user.Id).Update(user); err != nil {
		sess.Rollback()
		return false, err.Error(), nil
	}
	err = sess.Commit()
	if err != nil {
		return false, err.Error(), nil
	}
	return true, "", lastCharge
}

func (b *basicChargeService) Delete(ctx context.Context, chargeId string) (status bool, errinfo string, data *model.Charge) {
	var err error
	sess := b.Engine().NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return false, err.Error(), nil
	}

	charge := &model.Charge{
		Id: chargeId,
	}

	if _, err = sess.Delete(charge); err != nil {
		sess.Rollback()
		return false, err.Error(), nil
	}
	err = sess.Commit()
	if err != nil {
		return false, err.Error(), nil
	}
	return true, "", nil
}

// NewBasicChargeService returns a naive, stateless implementation of ChargeService.
func NewBasicChargeService(conf string) ChargeService {
	basicChargeSvc := &basicChargeService{
		&db.DBService{},
	}
	err := basicChargeSvc.Bind(conf)
	if err != nil {
		log.Printf("The ChargeService failed to bind with mysql")
	}
	return basicChargeSvc
}
