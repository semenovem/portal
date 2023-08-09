package logger

type Setter struct {
	logger *Pen
}

func (s *Setter) SetLevel(level int8) {
	s.logger.setLevel(level)
}

func (p *Pen) setLevel(lev int8) {
	p.level = lev
}

func (s *Setter) SetCli(on bool) {
	s.logger.cli = on

	prefixBytesErr = []byte("\033[0;31m" + prefixErr + "\033[0m")
	prefixBytesInfo = []byte("\033[0;32m" + prefixInfo + "\033[0m")
	prefixBytesDebug = []byte("\033[1;34m" + prefixDebug + "\033[0m")

	prefixLen = len(prefixBytesErr)
}

func (s *Setter) SetShowTime(on bool) {
	s.logger.hideTime = on
}