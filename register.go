package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
)

func register(c *gin.Context) {
	var newuser user
	newuser.FName = c.PostForm("register_first_name")
	newuser.LName = c.PostForm("register_last_name")
	newuser.Email = c.PostForm("register_email")
	newuser.Pass = c.PostForm("register_password")
	newuser.SecretKey = ksuid.New()
	newuser.Verified = false
	newuser.Role = "user"
	newuser.CreationDate = time.Now()
	fmt.Println(newuser)
	err := insertToDB(newuser)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/register")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/auth/userpage")

	}

}

func insertToDB(p user) error {
	query := "INSERT INTO DropItUsersDB(Email, First_Name,Last_Name,PassW, Secret_Key,URole, Creation_Date, Is_Verified ) VALUES (?, ?, ?, ?, ?, ?, ?, ? )"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := usersDb.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, p.Email, p.FName, p.LName, p.Pass, p.SecretKey, p.Role, p.CreationDate, p.Verified)
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d products created ", rows)
	return nil
}
