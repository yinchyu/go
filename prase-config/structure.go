package main

import "encoding/xml"

type AutoGenerated struct {
	ScrapeConfigs []ScrapeConfigs `yaml:"scrape_configs" json:"scrape_configs"`
}

type StaticConfigs struct {
	Targets []string `yaml:"targets" json:"targets"`
}

type Params struct {
	Format []string `yaml:"format" json:"format"`
}

type ScrapeConfigs struct {
	JobName       string          `yaml:"job_name" json:"job_name"`
	StaticConfigs []StaticConfigs `yaml:"static_configs" json:"static_configs"`
	Params        Params          `yaml:"params,omitempty" json:"params,omitempty"`
	HonorLabels   bool            `yaml:"honor_labels,omitempty" json:"honor_labels,omitempty"`
	MetricsPath   string          `yaml:"metrics_path,omitempty" json:"metrics_path,omitempty"`
}

type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Server  []struct {
		Text       string `xml:",chardata"`
		ServerName struct {
			Text string `xml:",chardata"`
		} `xml:"serverName"`
		ServerIP struct {
			Text string `xml:",chardata"`
		} `xml:"serverIP"`
	} `xml:"server"`
}