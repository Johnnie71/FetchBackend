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

func CalculatePoints(reciept models.Reciept) int {
	points := 0

	v := reflect.ValueOf(reciept)

	for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fieldName := v.Type().Field(i).Name
	
			switch fieldName {
					case "Retailer":
							for _, c := range field.String() {
									if unicode.IsLetter(c) || unicode.IsDigit(c) {
											points += 1
									}
									
							}
					case "PurchaseDate":
							day := strings.Split(field.String(), "-")[2]
							dayInt , err := strconv.Atoi(day)

							if err != nil {
									fmt.Println("Error converting day to integer", err)
							} else {

									if dayInt % 2 != 0 {
											points += 6
									}
							}
					
					case "PurchaseTime":
							purchaseTime, err := time.Parse("15:04", field.String())
			
							if err != nil {
									fmt.Println("Error parsing time:", err)
							} else {
									startTime, _ := time.Parse("15:04", "14:00") // 2:00pm
									end_time, _ := time.Parse("15:04", "16:00") // 4:00pm
									
									if purchaseTime.After(startTime) && purchaseTime.Before(end_time) {
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
													fmt.Println("Error converting string to float:", err)
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
									fmt.Print("Error converting string total to float", err)
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

	return points
}