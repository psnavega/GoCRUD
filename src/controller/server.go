package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"my-app/src/db"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func CreateAlbum(w http.ResponseWriter, r *http.Request) {
	bodyReq, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Failed read body"))
		return
	}

	var album album
	if err = json.Unmarshal(bodyReq, &album); err != nil {
		w.Write([]byte("Failed parse album to struct"))
		return
	}

	conn, err := db.Connect()

	statement, err := conn.Prepare("INSERT INTO album (title, artist, price) VALUES (?, ?, ?);")
	if err != nil {
		w.Write([]byte("Failed prepare query"))
		return
	}

	defer statement.Close()

	result, err := statement.Exec(album.Title, album.Artist, album.Price)
	if err != nil {
		w.Write([]byte("Failed save at database"))
		return
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		w.Write([]byte("Failed return id saved at database"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Album inserted successfully ID: %d", lastId)))
}

func GetAllAlbuns(w http.ResponseWriter, r *http.Request) {
	conn, err := db.Connect()
	if err != nil {
		w.Write([]byte("Failed to connect at database"))
		return
	}

	rows, err := conn.Query("SELECT * FROM album;")
	if err != nil {
		w.Write([]byte("Failed to prepare query"))
		return
	}

	defer rows.Close()

	var albuns []album
	for rows.Next() {
		var album album

		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			w.Write([]byte("Failed to scan user"))
			return
		}

		albuns = append(albuns, album)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(albuns); err != nil {
		w.Write([]byte("Failed to convert datas to json"))
	}
}

func GetAlbum(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		w.Write([]byte("Error to convert to string"))
	}

	conn, err := db.Connect()
	if err != nil {
		w.Write([]byte("Connect with database was error"))
		return
	}

	rows, err := conn.Query("SELECT * FROM album WHERE id = ?;", id)
	if err != nil {
		w.Write([]byte("Failed to prepare query"))
		return
	}

	defer rows.Close()

	var album album
	for rows.Next() {
		if err := rows.Scan(&album.ID, &album.Artist, &album.Title, &album.Price); err != nil {
			w.Write([]byte("Failed inside next"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(album); err != nil {
		w.Write([]byte("Failed return encoded "))
		return
	}
}

func UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		w.Write([]byte("Error parsing id from body"))
		return
	}

	var album album

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Error parsing data from body"))
		return
	}

	err = json.Unmarshal(body, &album)
	if err != nil {
		w.Write([]byte("Error at cast param"))
		return
	}

	conn, err := db.Connect()
	if err != nil {
		w.Write([]byte("Error at connect with database"))
		return
	}

	stm, err := conn.Prepare("UPDATE album SET title=?, artist=?, price=? WHERE id=?;")
	if err != nil {
		w.Write([]byte("Error at update data"))
		return
	}

	defer stm.Close()

	_, err = stm.Exec(album.Title, album.Artist, album.Price, id)
	if err != nil {
		w.Write([]byte("Error at execute data"))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id, err := strconv.ParseUint(param["id"], 10, 64)
	if err != nil {
		w.Write([]byte("Error parsing id"))
		return
	}

	conn, err := db.Connect()
	if err != nil {
		w.Write([]byte("Error connecting at database"))
		return
	}

	_, err = conn.Query("DELETE FROM album WHERE id=?", id)
	if err != nil {
		w.Write([]byte("Error executing query"))
		return
	}

	w.WriteHeader(http.StatusOK)
}
