package main

import (
	"fmt"
	"time"
)

func main() {
	lot := &ParkingLot{
		HatchbackCars: make(map[string]Car),
		SUVCars:       make(map[string]Car),
	}
	lot.ParkTheCar("H0001", "hatchback")
	lot.ParkTheCar("S0001", "suv")
	fmt.Printf("%+v\n", lot)

	time.Sleep(1 * time.Minute)
	pay, _ := lot.removeCar("H0001")
	fmt.Printf("payment for H0001: %d rupees\n", pay)

}

type ParkingLot struct {
	HatchbackCount int
	SUVcount       int
	HatchbackCars  map[string]Car
	SUVCars        map[string]Car
}

type Car struct {
	Id       string
	Type     string
	ParkedAt time.Time
	Duration time.Duration
}

const (
	maxHatchbacks  = 5
	hatchbackPrize = 10
	suvPrize       = 20
)

func (p *ParkingLot) ParkTheCar(carId string, carType string) bool {
	if carType == "hatchback" {
		if p.HatchbackCount < maxHatchbacks {
			p.HatchbackCount++
			p.HatchbackCars[carId] = Car{Id: carId, Type: carType, ParkedAt: time.Now()}
			return true
		} else if len(p.SUVCars) < maxHatchbacks {
			p.SUVcount++
			p.SUVCars[carId] = Car{Id: carId, Type: carType, ParkedAt: time.Now()}
			return true

		}
	} else if carType == "suv" {
		p.SUVcount++
		p.SUVCars[carId] = Car{Id: carId, Type: carType, ParkedAt: time.Now()}
		return true
	}
	return false
}

func (p *ParkingLot) Calculatepayment(carId string) (int, bool) {
	if car, exists := p.HatchbackCars[carId]; exists {
		dur := time.Since(car.ParkedAt)
		hours := int(dur.Minutes())
		return hours * hatchbackPrize, true

	} else if car, exists := p.SUVCars[carId]; exists {
		dur := time.Since(car.ParkedAt)
		hours := int(dur.Minutes())
		return hours * suvPrize, true
	}
	return 0, false
}

func (p *ParkingLot) removeCar(carId string) (int, bool) {
	if _, exists := p.HatchbackCars[carId]; exists {
		payment, _ := p.Calculatepayment(carId)
		delete(p.HatchbackCars, carId)
		p.HatchbackCount--
		return payment, true
	} else if _, exists := p.SUVCars[carId]; exists {
		payment, _ := p.Calculatepayment(carId)
		delete(p.SUVCars, carId)
		p.SUVcount--
		return payment, true
	}
	return 0, false
}
