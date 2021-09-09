package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var discord_up bool

var known_servers map[string]server //server name -> server struct

type server struct {
	name        string //simply identifier
	addr        string //for byond queries
	comm_key    string //password
	webhook_key string //password for bot
	admins_page string //where to get admins from
	color       int    //for embeds
	mode        int    //encoding
}

func add_server(server server) {
	known_servers[server.name] = server
}

func check_server(server string) bool {
	_, ok := known_servers[server]
	return ok
}

func init() {
	known_servers = make(map[string]server)
}

func populate_servers() {
	for k := range known_servers {
		delete(known_servers, k)
	}
	defer logging_recover("DB PS ERR:")
	rows, err := Database.Query("select SRVNAME, SRVADDR, COMMKEY, WEBKEY, ADMINS_PAGE, COLOR, MODE from STATION_SERVERS ;")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var srvname, srvaddr, commkey, webkey, admp string
		var clr, encmode int
		if terr := rows.Scan(&srvname, &srvaddr, &commkey, &webkey, &admp, &clr, &encmode); terr != nil {
			panic(terr)
		}
		srvname = trim(srvname)
		srvaddr = trim(srvaddr)
		commkey = trim(commkey)
		webkey = trim(webkey)
		admp = trim(admp)
		add_server(server{
			name:        srvname,
			addr:        srvaddr,
			comm_key:    commkey,
			webhook_key: webkey,
			admins_page: admp,
			color:       clr,
			mode:        encmode,
		})
	}
}

var get_time func() string

func init_time() {
	defer logging_recover("init_time")
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		panic(err)
	}
	get_time = func() string {
		return time.Now().In(loc).Format("15:04:05")
	}
}

func start_ticker(tick_seconds int, callback func()) chan int {
	quit := make(chan int)
	go func() {
		tick := time.Tick(time.Duration(tick_seconds) * time.Second)
		for {
			select {
			case <-quit:
				return
			case <-tick:
				callback()
			}
		}
	}()
	return quit
}

func stop_ticker(ch chan int) {
	ch <- 0
}

func main() {
	db_init()
	log.Println("DB inited")
	init_time()
	log.Println("time inited")
	discord_init()
	log.Println("discord inited")
	populate_servers()
	log.Println("servers populated")
	populate_server_embeds()
	log.Println("server embeds populated")
	launch_ss_tickers()
	log.Println("ss tickers started")
	Dopen() //start discord
	log.Println("discord up")
	discord_up = true
	shell_repo_init()
	srv := Http_server() //start web server
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc //wait for SIGINT or kinda it
	discord_up = false
	Dclose()              //stop discord
	db_deinit()           //clean db templates
	http_server_stop <- 1 //stop server ticker
	//graceful shutdown for web server
	if err := srv.Shutdown(nil); err != nil {
		log.Fatal("Failed to shutdown webserver: ", err)
	}
	log.Println("Stoped correctly")
}
