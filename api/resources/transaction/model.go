package transaction

import (
	"github.com/gofiber/fiber/v2"
)

func (h Handler) CreateTransaction(ctx *fiber.Ctx, trx interface{}) error {
	err := h.Db.WithContext(ctx.UserContext()).Table(h.TableName).Create(trx).Error
	if err != nil {
		return err
	}

	return nil
}
func (h Handler) FindTransactionById(ctx *fiber.Ctx, trxId string) (interface{}, error) {
	trx := map[string]interface{}{}
	err := h.Db.WithContext(ctx.UserContext()).
		Table(h.TableName).Find(&trx, "\"id\" = ?", trxId).Error
	if err != nil {
		return nil, err
	}

	return trx, nil
}
func (h Handler) UpdateTransactionById(ctx *fiber.Ctx, trxId string, update interface{}) error {
	err := h.Db.WithContext(ctx.UserContext()).Table(h.TableName).
		Where("\"id\" = ?", trxId).Updates(update).Error
	if err != nil {
		return err
	}

	return nil
}
func (h Handler) FindTransactionWithLimit(ctx *fiber.Ctx, limit int) (interface{}, error) {
	trx := map[string]interface{}{}
	err := h.Db.WithContext(ctx.UserContext()).
		Table(h.TableName).Limit(limit).Find(&trx).Error
	if err != nil {
		return nil, err
	}

	return trx, nil
}
