package helpers

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"loyalty_system/logger"
)

func getIDFromRequestString(url string) (int, error) {
	vars := strings.Split(url, "/")

	return strconv.Atoi(vars[len(vars)-1])
}

func formatResponse(w http.ResponseWriter, httpStatus int, category string) {
	w.WriteHeader(httpStatus)

	data := ResponseData{
		"error": "Unsuccessfull request",
	}
	SendResponse(w, data, category)
}

func checkError(w http.ResponseWriter, err error, category string) bool {
	if err != nil {
		logger.Errorf("Internal error %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		data := ResponseData{
			"error": "Internal error",
		}
		SendResponse(w, data, category)

		return true
	}

	return false
}

func isExists(query string, id int) bool {
	rows, err := DB.Prepare(query)
	if err != nil {
		logger.Error(err.Error())

		return false
	}

	defer rows.Close()
	err = rows.QueryRow(id).Scan()
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return false
	}

	return true
}
