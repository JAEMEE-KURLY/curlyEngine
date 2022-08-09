package appmmain

import (
	. "almcm.poscoict.com/scm/pme/curly-engine/configure"
	. "almcm.poscoict.com/scm/pme/curly-engine/log"
	"sync"
	"almcm.poscoict.com/scm/pme/curly-engine/database/gorm/userdb"
	"almcm.poscoict.com/scm/pme/curly-engine/database/gorm/logdb"
	utility "almcm.poscoict.com/scm/pme/curly-engine/library"
)

func Run(conf *Values) {
	if conf.CurlyEngine.Enabled == false {
		Logi("CurlyEngine Application Disabled")
		return
	}
	Logi("CurlyEngine Application Started")

	wg := new(sync.WaitGroup)

	// TEST DATABASE
	userdb.InitUserTable()
	logdb.InitLoggingTable()

	// HTTP
	if conf.Net.EnableHttp && conf.CurlyEngine.EnableHttpServer {
		// go httpserver.HttpServer(conf.CurlyEngine.HttpServerPort)
	}

	// Websocket
	if conf.Net.EnableWebsocket && conf.CurlyEngine.EnableWebsocketServer {

	}

	// gRPC
	if conf.Net.EnableGrpc && conf.CurlyEngine.EnableGrpcServer {
		// go grpcserver.GrpcServer(conf.CurlyEngine.GrpcServerPort)
	}

	// TCP Socket
	if conf.Net.EnableTcp && conf.CurlyEngine.EnableTcpServer {
		// socket.TcpServerRun("CurlyEngine", conf.CurlyEngine.TcpServerPortCurlyEngine)
	}

	// Serial
	if conf.Net.EnableSerial && conf.CurlyEngine.EnableSerial {

	}

	// Message Queue
	if conf.Net.EnableMqueue && conf.CurlyEngine.EnableMqueue {

	}

	// Start Scheduler
	if conf.Time.SchedulerMinute {
		go utility.SchedulerMin(SchedulerMin)
	}
	if conf.Time.SchedulerHour {
		go utility.SchedulerHour(SchedulerHour)
	}
	if conf.Time.SchedulerDay {
		go utility.SchedulerDay(SchedulerDay)
	}

	wg.Wait()
}
