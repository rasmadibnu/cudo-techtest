package service

import (
	"cudo-techtest/entity"
	"cudo-techtest/repository"
	"fmt"
	"math"
	"strconv"
	"time"
)

type TransactionService struct {
	repository repository.TransactionRepository
}

func NewTransactionService(r repository.TransactionRepository) TransactionService {
	return TransactionService{
		repository: r,
	}
}

// @Summary : Insert
// @Description : Insert data to repository
// @Author : rasmadibnu
func (s *TransactionService) GetDataTransaction() (interface{}, error) {

	transactions, err := s.repository.Find()

	if err != nil {
		return transactions, err
	}

	tx := frequencyCheck(transactions)

	// tx := DetectOutliers(transactions)

	return tx, nil
}

func frequencyCheck(transactions []entity.Transaction) []entity.Transaction {
	type key struct {
		UserID     int
		HourBucket string
	}

	// Map to count number of transactions per (user_id, hour)
	countMap := make(map[key]int)

	// Map to track the index of the *latest* transaction per (user_id, hour)
	latestTxIndex := make(map[key]int)

	for i, tx := range transactions {
		hourBucket := tx.TransactionDate.Truncate(time.Hour).Format("2006-01-02 15:00")
		k := key{UserID: tx.UserID, HourBucket: hourBucket}

		countMap[k]++

		// Check if this tx is the latest for the key
		if existingIdx, ok := latestTxIndex[k]; !ok || tx.TransactionDate.After(transactions[existingIdx].TransactionDate) {
			latestTxIndex[k] = i
		}
	}

	var freq []entity.Transaction

	// Attach hourly_order_count only to the latest transaction if count > 5
	for k, count := range countMap {
		if count > 5 {
			latestIdx := latestTxIndex[k]
			transactions[latestIdx].DetectionResults = append(transactions[latestIdx].DetectionResults, entity.DetectionResult{
				IsSupicious:     true,
				ConfidanceScore: float64(GetConfidenceScore(float64(count))),
				Triggers: []string{
					fmt.Sprintf("High order frequency: %d in 1 hour", count),
				},
			})

			freq = append(freq, transactions[latestIdx])
		}
	}

	return freq
}

func DetectOutliers(transactions []entity.Transaction) []entity.Transaction {
	var sum, sumSq float64
	var parsed []entity.Transaction

	// Step 1: Parse amounts and calculate total & total squares
	for _, tx := range transactions {
		amt, err := strconv.ParseFloat(tx.Amount, 64)
		if err != nil {
			continue
		}
		tx.ParsedAmount = amt
		sum += amt
		sumSq += amt * amt
		parsed = append(parsed, tx)
	}

	n := float64(len(parsed))
	if n == 0 {
		return parsed
	}

	// Step 2: Compute mean and stddev
	mean := sum / n
	variance := (sumSq / n) - (mean * mean)
	stddev := math.Sqrt(variance)

	// Step 3: Check Z-score and apply confidence score
	for i, tx := range parsed {
		z := (tx.ParsedAmount - mean) / stddev
		parsed[i].ZScore = z

		if z > 2 {
			// score := int(math.Min(100, math.Max(0, (z-2)*20)))
			// parsed[i].ConfidenceScore = score
		}
	}

	return parsed
}

func GetConfidenceScore(transactionsPerHour float64) int {
	if transactionsPerHour > 8 {
		return 90 // 90-100
	} else if transactionsPerHour >= 7 {
		return 80 // 80-89
	} else if transactionsPerHour >= 6 {
		return 70 // 70-79
	} else if transactionsPerHour >= 5 {
		return 50 // 50-69
	} else {
		return 40 // <50
	}
}

// @Summary : Find
// @Description : Find data by id to repository
// @Author : rasmadibnu
func (s *TransactionService) FindById(ID int) (entity.Transaction, error) {
	data, err := s.repository.FindById(ID)

	if err != nil {
		return data, err
	}

	return data, nil
}

// @Summary : Udate
// @Description : Update data by id to repository
// @Author : rasmadibnu
func (s *TransactionService) Update(req entity.Transaction, ID int) (entity.Transaction, error) {

	updateData, err := s.repository.Update(req, ID)
	if err != nil {
		return updateData, err
	}

	return updateData, nil
}

// @Summary : Delete
// @Description : Delete data by id to repository
// @Author : rasmadibnu
func (s *TransactionService) Delete(ID int) (bool, error) {
	data, err := s.repository.Delete(ID)

	if err != nil {
		return data, err
	}

	return data, nil
}
