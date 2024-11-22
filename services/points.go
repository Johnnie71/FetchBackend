package services

import (
	"backend-service/models"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func CalculatePoints(reciept models.Reciept) (int, error) {
	points := 0

	v := reflect.ValueOf(reciept)

	for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fieldName := v.Type().Field(i).Name
	
			switch fieldName {
					case "Retailer":
							retailerName := field.String()

							if retailerName == "" {
								return 0, fmt.Errorf("retailer name is empty")
							}

							for _, c := range retailerName {
									if unicode.IsLetter(c) || unicode.IsDigit(c) {
											points += 1
									}
									
							}
					case "PurchaseDate":
							purchaseDate := field.String()

							if purchaseDate == "" {
								return 0, fmt.Errorf("purchase date is empty")
							}

							parts := strings.Split(purchaseDate, "-")

							if len(parts) < 3 {
								return 0, fmt.Errorf("invalid purchase date format, expected format: YYYY-MM-DD, got: %s", purchaseDate)
							}

							day := parts[2]
							dayInt , err := strconv.Atoi(day)

							if err != nil {
									return 0, fmt.Errorf("error converting day to integer: %v", err)
							} else {

									if dayInt % 2 != 0 {
											points += 6
									}
							}
					
					case "PurchaseTime":
							purchaseTime := field.String()

							if purchaseTime == "" {
								return 0, fmt.Errorf("purchase time is empty")
							}
							
							timeOfPurchase, err := time.Parse("15:04", purchaseTime)
			
							if err != nil {
								return 0, fmt.Errorf("Error parsing time: %v", err)
							} else {
									startTime, _ := time.Parse("15:04", "14:00") // 2:00pm
									end_time, _ := time.Parse("15:04", "16:00") // 4:00pm
									
									if timeOfPurchase.After(startTime) && timeOfPurchase.Before(end_time) {
											points += 10
									}
							}
					case "Items":
							// Finding amount of pairs of items
							numOfItems := field.Len()
							pairs := numOfItems / 2
							points += pairs * 5
							
							// Checking descriptions
							for i := 0; i < numOfItems; i ++ {
									item := field.Index(i)

									itemDescription := item.FieldByName("ShortDescription").String()
									trimmed := strings.TrimSpace(itemDescription)
									
									if len(trimmed) % 3 == 0 {
											itemPrice := item.FieldByName("Price").String()
											price, err := strconv.ParseFloat(itemPrice, 64)
											if err != nil {
												return 0, fmt.Errorf("error converting string to float for item price: %v", err)
											} else {
													addedPoints := math.Ceil(price * 0.2)
													points += int(addedPoints)
											}

									}
							}
					case "Total":
							t := field.String()
							total, err := strconv.ParseFloat(t, 64)
							if err != nil {
								return 0, fmt.Errorf("error converting string total to float: %v", err)
							} else {
									// Checking if there are any cents
									if total == math.Floor(total) {
											points += 50
									}

									// Checking if a multiple of .25
									if math.Mod(total * 4, 1) == 0 {
											points += 25
									}
							}
							
			}
	}

	return points, nil
}