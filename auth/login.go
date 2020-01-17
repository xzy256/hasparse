package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type HasConfig struct {
	HttpsHost      string
	HttpsPort      string
	UserName	   string
	Password	   string
	Type		   string
}

func NewHasConfig() *HasConfig {
	return &HasConfig{
	}
}

func getUserAndPassword(clientUgi string) (string, string) {
	ugiPath := ""
	if strings.Contains(clientUgi, ",") {
		nameAndPass := strings.Split(clientUgi, ",")
		return nameAndPass[0], nameAndPass[1]
	} else {
		if strings.HasPrefix(clientUgi, "/") {
			ugiPath = clientUgi
		} else {
			ugiPath = os.Getenv("HOME") + "/" + clientUgi
		}
	}
	return readUgiConfig(ugiPath)
}

// read first line of ugi_config
func readUgiConfig(ugiFile string) (string, string) {
	fi, err := os.Open(ugiFile)
	if err != nil {
		fmt.Println("Open ugi_config:", ugiFile, err)
		return "", ""
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	line, _, err := br.ReadLine()
	if err != nil {
		fmt.Println("Read ugi_config:", ugiFile, err)
		return "", ""
	}
	userAndPasswd := strings.Split(string(line), ",")

	return userAndPasswd[0], userAndPasswd[1]
}

// load has-client.conf
func loadHasClientConf(krb5conf []byte) (*HasConfig, error) {
	var hscfg *HasConfig
	var err error
	if krb5conf == nil || len(krb5conf) == 0 {
		configPath := os.Getenv("KRB5_CONFIG")
		if configPath == "" {
			configPath = "/etc/has/has-client.conf"
		}
		cfg, err := NewHasConfigFromString(configPath)
		if err != nil {
			return nil, err
		}
		return cfg, err
	}
	reader := bytes.NewReader(krb5conf)
	hscfg, err = NewHasConfigFromReader(reader)
	if err != nil {
		return nil, err
	}

	return hscfg, nil
}

func NewHasConfigFromReader(r io.Reader) (*HasConfig, error) {
	scanner := bufio.NewScanner(r)
	return NewHasConfigFromScanner(scanner)
}

func NewHasConfigFromString(confPath string) (*HasConfig, error) {
	fh, err := os.Open(confPath)
	if err != nil {
		return nil, errors.New("configuration file could not be opened: " + confPath + " " + err.Error())
	}
	defer fh.Close()
	scanner := bufio.NewScanner(fh)
	return NewHasConfigFromScanner(scanner)
}

func NewHasConfigFromScanner(scanner *bufio.Scanner) (*HasConfig, error) {
	c := NewHasConfig()
	var e error
	sections := make(map[int]string)
	var sectionLineNum []int
	var lines []string
	for scanner.Scan() {
		// Skip comments and blank lines
		if matched, _ := regexp.MatchString(`^\s*(#|;|\n)`, scanner.Text()); matched {
			continue
		}
		if matched, _ := regexp.MatchString(`^\s*\[HAS\]\s*`, scanner.Text()); matched {
			sections[len(lines)] = "HAS"
			sectionLineNum = append(sectionLineNum, len(lines))
			continue
		}
		if matched, _ := regexp.MatchString(`^\s*\[PLUGIN\]\s*`, scanner.Text()); matched {
			sections[len(lines)] = "PLUGIN"
			sectionLineNum = append(sectionLineNum, len(lines))
			continue
		}
		lines = append(lines, scanner.Text())
	}
	for i, start := range sectionLineNum {
		var end int
		if i+1 >= len(sectionLineNum) {
			end = len(lines)
		} else {
			end = sectionLineNum[i+1]
		}
		switch section := sections[start]; section {
		case "HAS":
			err := c.parseHasSection(lines[start:end])
			if err != nil {
				return nil, fmt.Errorf("error processing libdefaults section: %v", err)
			}
		case "PLUGIN":
			// do nothing
		default:
			continue
		}
	}
	return c, e
}

// Parse the lines of the [HAS] section of the configuration into the HasConfig struct.
func (this *HasConfig) parseHasSection(lines []string) error {
	for _, line := range lines {
		//Remove comments after the values
		if idx := strings.IndexAny(line, "#;"); idx != -1 {
			line = line[:idx]
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if !strings.Contains(line, "=") {
			return errors.New("libdefaults section line:" + line)
		}

		p := strings.Split(line, "=")
		key := strings.TrimSpace(strings.ToLower(p[0]))
		switch key {
		case "https_host":
			this.HttpsHost = strings.TrimSpace(p[1])
		case "https_port":
			this.HttpsPort = strings.TrimSpace(p[1])
		default:
			//Ignore the line
			continue
		}
	}

	return nil
}

