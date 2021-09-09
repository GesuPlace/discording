package main

import (
	"fmt"
	"log"
	"strings"
)

func log_line_runtime(in string) {
	log.Println(in) // usual logging
	log_line(in, "runtimes")
}

func log_line(in, ch string) {
	if strings.HasSuffix(in, "connect: connection refused") || strings.HasSuffix(in, "i/o timeout") || strings.HasSuffix(in, "timed out") {
		return //don't want to log it
	}
	if discord_up {
		chans, ok := known_channels_s_t_id_m["generic"]
		if !ok {
			return
		}
		channels, ok := chans[ch]
		if !ok || len(channels) < 1 {
			return //no bound channels
		}
		for _, id := range channels {
			send_message(id, in)
		}
	}
}

func logging_crash(ctx string) {
	if r := recover(); r != nil {
		log.Fatalln("ERRF: ["+ctx+"]:", r)
	}
}

func logging_recover(ctx string) {
	if r := recover(); r != nil {
		log_line_runtime("ERR: [" + ctx + "]: " + fmt.Sprint(r))
	}
}

func recovering_callback(callback func()) {
	if r := recover(); r != nil {
		callback()
	}
}
func onerror(callback func()) {
	if r := recover(); r != nil {
		callback()
		panic(r)
	}
}

func logging_pass(ctx string) {
	if r := recover(); r != nil {
		log_line_runtime("ERR: [" + ctx + "]: " + fmt.Sprint(r))
		panic(r)
	}
}

func rise_error(app string) {
	if r := recover(); r != nil {
		panic(fmt.Sprintf("[%v]:%v", app, r))
	}
}

func noerror(err error) {
	if err != nil {
		panic(err)
	}
}

func assert(st bool, message string) {
	if !st {
		panic(message)
	}
}

func maybeerror(err error) {
	if err != nil {
		log_line_runtime(fmt.Sprintf("MERR: %v", err))
	}
}
