package setting

import "time"

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type APPSettingS struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}

type DatabaseSettingS struct {
	DBType               string
	UserName             string
	Password             string
	Host                 string
	DBName               string
	TablePrefix          string
	Charset              string
	ParseTime            bool
	MaxIdleconns         int
	MaxOpenConns         int
	ConnMaxLifetimeHours int
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
