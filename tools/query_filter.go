package tools

import (
	"Eff_Mob/models"
	"net/http"
	"strconv"
)

//Functions for making queries:

func QueryFilterSelect(req models.Car, page string) string {
	query := "select * from car "
	flag := false
	if req.RegNum != "" {
		query += "where " + "regnum = " + "'" + req.RegNum + "'"
		flag = true
	}

	if req.Mark != "" {
		if flag {
			query += " and mark = " + "'" + req.Mark + "'"
		} else {
			query += " where mark = " + "'" + req.Mark + "'"
		}

		flag = true
	}

	if req.Model != "" {
		if flag {
			query += " and model = " + "'" + req.Mark + "'"
		}
		query += " where model = " + "'" + req.Model + "'"
		flag = true
	}

	if req.Year != 0 {
		if flag {
			query += " and year = " + "'" + req.Mark + "'"
		}
		query += " where year = " + strconv.Itoa(req.Year)
		flag = true
	}

	if req.OwnerId != 0 {
		if flag {
			query += " and owner = " + "'" + req.Mark + "'"
		}
		query += " where owner = " + strconv.Itoa(req.OwnerId)
		flag = true
	}
	perPage := 2
	p, _ := strconv.Atoi(page)
	offset := (p - 1) * perPage
	query += " limit " + strconv.Itoa(perPage) + " offset " + strconv.Itoa(offset)
	return query

}

func QueryFilterUpdateCar(w http.ResponseWriter, req models.CarInfo) string {

	if (req.Mark == "") || (req.Model == "") || (req.Year == 0) {
		w.WriteHeader(http.StatusBadRequest)
		return ""
	} else {
		query := "update car set" + " mark = " + "'" + req.Mark + "', " + "model = " + "'" + req.Model + "', " +
			"year = " + strconv.Itoa(req.Year) + " where regnum = " + "'" + req.RegNum + "'"
		w.WriteHeader(http.StatusOK)
		return query
	}

}

func QueryOwnerSelect(cars []models.Car) string {
	query := "select * from people " + "where id in ("
	for i := range cars {
		if i == len(cars)-1 {
			query += strconv.Itoa(cars[i].OwnerId)
		} else {
			query += strconv.Itoa(cars[i].OwnerId) + ", "
		}

	}
	query = query + ")"
	return query
}
