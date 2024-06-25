package model

type Vulnerability struct {
	VulnerabilityID  string   `json:"vulnerability_id"`
	PkgName          string   `json:"pkg_name"`
	InstalledVersion string   `json:"installed_version"`
	FixedVersion     string   `json:"fixed_version"`
	Severity         string   `json:"severity"`
	Description      string   `json:"description"`
	References       []string `json:"references"`
}

type ScanResult struct {
	Target        string          `json:"target"`
	Vulnerability []Vulnerability `json:"vulnerabilities"`
}
