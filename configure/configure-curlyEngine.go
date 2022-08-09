package configure

// exmaple structure
type ValueCurlyEngine struct {
	sectionName string
	// Insert configure values here
	Enabled 				bool

	EnableHttpServer 		bool
	EnableWebsocketServer   bool
	EnableGrpcServer		bool
	EnableTcpServer	 		bool
	EnableSerial	 		bool
	EnableMqueue			bool

	HttpServerPort 	 		string
	GrpcServerPort			string

	TcpServerPortCurlyEngine 	string
	TcpClientPortCurlyEngine 	string

	SerialPortCurlyEngine		string

	SwaggerAddress          string
}

func (c *Values) GetValueCurlyEngine(filePath string) (*Values, error) {
	_, err := c.chkCfgFile(filePath)
	if err != nil {
		return nil, err
	}

	c.CurlyEngine.sectionName = SECT_CurlyEngine
	c.CurlyEngine.Enabled = Config.File.Section(SECT_CurlyEngine).Key("Enabled").MustBool(false)

	// Insert configure values here
	c.CurlyEngine.EnableHttpServer = Config.File.Section(SECT_CurlyEngine).Key("EnableHttpServer").MustBool(false)
	c.CurlyEngine.EnableWebsocketServer = Config.File.Section(SECT_CurlyEngine).Key("EnableWebsocketServer").MustBool(false)
	c.CurlyEngine.EnableGrpcServer = Config.File.Section(SECT_CurlyEngine).Key("EnableGrpcServer").MustBool(false)
	c.CurlyEngine.EnableTcpServer = Config.File.Section(SECT_CurlyEngine).Key("EnableTcpServer").MustBool(false)
	c.CurlyEngine.EnableSerial = Config.File.Section(SECT_CurlyEngine).Key("EnableSerial").MustBool(false)
	c.CurlyEngine.EnableMqueue = Config.File.Section(SECT_CurlyEngine).Key("EnableMqueue").MustBool(false)

	c.CurlyEngine.HttpServerPort = Config.File.Section(SECT_CurlyEngine).Key("HttpServerPort").MustString("1323")
	c.CurlyEngine.GrpcServerPort = Config.File.Section(SECT_CurlyEngine).Key("GrpcServerPort").MustString("1324")

	c.CurlyEngine.TcpServerPortCurlyEngine = Config.File.Section(SECT_CurlyEngine).Key("TcpServerPort_CurlyEngine").MustString("9900")
	c.CurlyEngine.TcpClientPortCurlyEngine = Config.File.Section(SECT_CurlyEngine).Key("TcpClientPort_CurlyEngine").MustString("9901")

	c.CurlyEngine.SerialPortCurlyEngine = Config.File.Section(SECT_CurlyEngine).Key("SerialPort_CurlyEngine").MustString("/dev/ttyUSB0")

	c.CurlyEngine.SwaggerAddress = Config.File.Section(SECT_CurlyEngine).Key("SwaggerAddress").MustString("localhost")

	return c, nil
}
