package main

import (
	"flag"
	"fmt"
	conf "almcm.poscoict.com/scm/pme/curly-engine/configure"
	. "almcm.poscoict.com/scm/pme/curly-engine/log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	app "almcm.poscoict.com/scm/pme/curly-engine/application" // insert your application name
)

const CurlyEngineVersion = "0.9.0"

// global variables
var curlyEngine *CurlyEngine

// CurlyEngine application main structure
type CurlyEngine struct {
	args    *Arguments
	config 	*conf.Values
	chans 	*Channels
}

// Arguments process arguments
type Arguments struct {
	printConfig *string
	singleShot  *bool
	iniFile     *string
	test 	    *string

	/* insert your application arguments */
}

// Channels go channels
type Channels struct {
	doneChan chan bool			// true: exit application
	sigChan  chan os.Signal		// signal channel (SIGINT...)
}

func NewChannels() *Channels {
	return &Channels{}
}

func NewCurlyEngine() *CurlyEngine {
	if curlyEngine == nil {
		curlyEngine = &CurlyEngine{}
		curlyEngine.config = conf.NewValues()
		curlyEngine.chans = NewChannels()
	}
	return curlyEngine
}

func NewArguments() *Arguments {
	return &Arguments{}
}

func parseArguments(args *Arguments) bool {
	/* common arguments */
	args.printConfig = flag.String("pc", "", "Print config values [all|core|timer|log|net|...]")
	args.singleShot  = flag.Bool("ss", false, "Run once and exit the application")
	args.iniFile     = flag.String("ini", "curlyEngine.ini", "Set configuration file")
	args.test        = flag.String("test", "", "Input a string argument for the test function")

	/* insert your application arguments */

	flag.Parse()

	return true
}

func runSingleShot() {
	Logi("SingleShot Function Started")

	if len(*curlyEngine.args.test) > 0 {
		runTestCode(*curlyEngine.args.test)
	}

	// insert your singleshot application logic

	Logi("SingleShot Function Finished")
}

func runInfinite() {

	if curlyEngine.config.CurlyEngine.Enabled {
		go app.Run(curlyEngine.config)
	}

	// start your application

	ticker := time.NewTicker(time.Second * time.Duration(curlyEngine.config.Time.IdleTimeout))
	for {
		select {
		case sig := <-curlyEngine.chans.sigChan:
			Logd("\nSignal received: %s", sig)
			curlyEngine.chans.doneChan <- true
		case <-ticker.C:
			Logd("IDLE Timeout %d sec", curlyEngine.config.Time.IdleTimeout)
		default:
			//Logd("Log rotate test...")
			time.Sleep(20 * time.Millisecond)
		}
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // max number of cpu on this PC

	curlyEngine = NewCurlyEngine()

	// parse application arguments
	curlyEngine.args = NewArguments()
	if parseArguments(curlyEngine.args) == false {
		return
	}

	// Exit the application after executing this function
	defer func() {
		if *curlyEngine.args.singleShot == false {
			fmt.Println("### CurlyEngine Application Finished ###")
		}
		Logi("### CurlyEngine Application Finished ###")
		Close()
	}()

	// load configuration
	_, err := curlyEngine.config.GetValueALL(*curlyEngine.args.iniFile)
	if err != nil {
		fmt.Printf("Error, open ini file, [%v]\n", err)
		return
	}

	// init log system
	InitLogger(curlyEngine.config.Log)

	// print application start message
	Logi("### CurlyEngine Application Started, Version[%s], CPU Core Num: %d ###", CurlyEngineVersion, runtime.GOMAXPROCS(0))

	// print config values
	if len(*curlyEngine.args.printConfig) > 0 {
		values := curlyEngine.config.PrintValues(*curlyEngine.args.printConfig)
		fmt.Printf(values)
		return
	}

	// single-shot logic
	if *curlyEngine.args.singleShot == true {
		runSingleShot()
		return
	}

	// make default channels
	curlyEngine.chans.sigChan = make(chan os.Signal, 1)
	curlyEngine.chans.doneChan = make(chan bool, 1)
	signal.Notify(curlyEngine.chans.sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Main Infinite Loop
	go runInfinite()
	<-curlyEngine.chans.doneChan
}

func runTestCode(str string) {
	Logd("Start TEST Logic, String: [%s]", str)
}