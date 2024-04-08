package postgres

import (
	"Eff_Mob/models"
	"Eff_Mob/tools"
	"fmt"
	"log/slog"
)

func (s *Storage) GetCarInfo(log *slog.Logger, request string) ([]models.CarInfo, error) {

	carexec, err := s.Db.Query(request)
	if err != nil {
		log.Warn("get exec mistake: \n", err)
		return nil, err
	}

	var cars []models.Car

	for carexec.Next() {
		car := models.Car{}
		err := carexec.Scan(&car.RegNum, &car.Mark, &car.Model, &car.Year, &car.OwnerId)
		if err != nil {
			fmt.Println(err)
			continue
		}
		cars = append(cars, car)

	}

	query := tools.QueryOwnerSelect(cars)

	ownerexec, err := s.Db.Query(query)
	if err != nil {
		log.Warn("get owner prepare mistake\n", err)
		return nil, err
	}
	var ownerStr []models.Owner

	for ownerexec.Next() {
		o := models.Owner{}
		err := ownerexec.Scan(&o.Id, &o.Name, &o.Surname, &o.Patronymic)
		if err != nil {
			fmt.Println(err)
			continue
		}
		ownerStr = append(ownerStr, o)
	}

	var arrCarInfo []models.CarInfo
	for i := range cars {
		carInfo := new(models.CarInfo)
		carInfo.RegNum = cars[i].RegNum
		carInfo.Mark = cars[i].Mark
		carInfo.Model = cars[i].Model
		carInfo.Year = cars[i].Year
		for j := range ownerStr {
			if ownerStr[j].Id == cars[i].OwnerId {
				carInfo.CurrentOwner = ownerStr[j]
				break
			}
		}
		arrCarInfo = append(arrCarInfo, *carInfo)
	}

	return arrCarInfo, nil
}
