package repository

import (
	"context"
	"fmt"
	"strconv"

	"github.com/iLaurian/my-portfolio-vs-spx/entity"
	"github.com/redis/go-redis/v9"
)

type HoldingRepository interface {
	GetAll() ([]entity.Holding, error)
	Add(entity.Holding) error
	DeleteAll() error
}

type rdsHoldingRepository struct {
	Redis *redis.Client
}

func NewHoldingRepository(redisClient *redis.Client) HoldingRepository {
	return &rdsHoldingRepository{
		Redis: redisClient,
	}
}

func (r *rdsHoldingRepository) Add(holding entity.Holding) error {
	key := holding.Symbol

	fields := map[string]interface{}{
		"shares":      holding.Shares,
		"marketvalue": holding.MarketValue,
		"openprice":   holding.OpenPrice,
		"marketprice": holding.MarketPrice,
		"grossprofit": holding.GrossProfit,
		"winorloss":   holding.WinOrLoss,
	}

	err := r.Redis.HSet(context.Background(), key, fields).Err()
	if err != nil {
		return fmt.Errorf("failed to add holding to redis: %w", err)
	}

	return nil
}

func (r *rdsHoldingRepository) GetAll() ([]entity.Holding, error) {
	keys, err := r.Redis.Keys(context.Background(), "*").Result()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve keys from redis: %w", err)
	}

	var holdings []entity.Holding

	for _, key := range keys {
		fields, err := r.Redis.HGetAll(context.Background(), key).Result()
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve holding for key %s: %w", key, err)
		}

		holding, err := mapToHolding(key, fields)
		if err != nil {
			return nil, fmt.Errorf("failed to map fields to holding for key %s: %w", key, err)
		}

		holdings = append(holdings, holding)
	}

	return holdings, nil
}

func mapToHolding(symbol string, fields map[string]string) (entity.Holding, error) {
	holding := entity.Holding{
		Symbol: symbol,
	}

	if shares, err := strconv.ParseFloat(fields["shares"], 32); err == nil {
		holding.Shares = float32(shares)
	} else {
		return holding, fmt.Errorf("failed to parse shares: %w", err)
	}

	if marketValue, err := strconv.ParseFloat(fields["marketvalue"], 32); err == nil {
		holding.MarketValue = float32(marketValue)
	} else {
		return holding, fmt.Errorf("failed to parse marketvalue: %w", err)
	}

	if openPrice, err := strconv.ParseFloat(fields["openprice"], 32); err == nil {
		holding.OpenPrice = float32(openPrice)
	} else {
		return holding, fmt.Errorf("failed to parse openprice: %w", err)
	}

	if marketPrice, err := strconv.ParseFloat(fields["marketprice"], 32); err == nil {
		holding.MarketPrice = float32(marketPrice)
	} else {
		return holding, fmt.Errorf("failed to parse marketprice: %w", err)
	}

	if grossProfit, err := strconv.ParseFloat(fields["grossprofit"], 32); err == nil {
		holding.GrossProfit = float32(grossProfit)
	} else {
		return holding, fmt.Errorf("failed to parse grossprofit: %w", err)
	}

	if winOrLoss, err := strconv.ParseFloat(fields["winorloss"], 32); err == nil {
		holding.WinOrLoss = float32(winOrLoss)
	} else {
		return holding, fmt.Errorf("failed to parse winorloss: %w", err)
	}

	return holding, nil
}

func (r *rdsHoldingRepository) DeleteAll() error {
	err := r.Redis.FlushDB(context.Background()).Err()
	if err != nil {
		return fmt.Errorf("failed to delete all records from redis: %w", err)
	}

	return nil
}
