package jumpserver

import (
	"fmt"
)

// UsersService handles communication with the user related
// methods of the GitHub API.
//
// GitHub API docs: https://developer.github.com/v3/users/
type AssetsService struct {
	client *Client

	// Authentication type
	authType int

	// Basic auth username
	username string

	// Basic auth password
	password string
}

// User represents a GitHub user.
type Asset struct {
	Id           string   `json:"id"`
	OrgId        string   `json:"org_id"`
	Ip           string   `json:"ip"`
	Hostname     string   `json:"hostname"`
	Protocol     string   `json:"protocol"`
	Port         int      `json:"port"`
	Platform     string   `json:"platform"`
	IsActive     bool     `json:"is_active"`
	PublicIp     string   `json:"public_ip"`
	Number       string   `json:"number"`
	Vendor       string   `json:"vendor"`
	Model        string   `json:"model"`
	Sn           string   `json:"sn"`
	CpuModel     string   `json:"cpu_model"`
	CpuCount     string   `json:"cpu_count"`
	CpuCores     string   `json:"cpu_cores"`
	CpuVcpus     string   `json:"cpu_vcpus"`
	Memory       string   `json:"memory"`
	DiskTotal    string   `json:"disk_total"`
	DiskInfo     string   `json:"disk_info"`
	Os           string   `json:"os"`
	OsVersion    string   `json:"os_version"`
	OsArch       string   `json:"os_arch"`
	HostnameRaw  string   `json:"hostname_raw"`
	CreatedBy    string   `json:"created_by"`
	DateCreated  string   `json:"date_created"`
	Comment      string   `json:"comment"`
	Domain       string   `json:"domain"`
	AdminUser    string   `json:"admin_user"`
	Nodes        []string `json:"nodes"`
	Labels       []string `json:"labels"`
	HardwareInfo string   `json:"hardware_info"`
	Connectivity int      `json:"connectivity"`
	OrgName      string   `json:"org_name"`
}

func (s *AssetsService) GetList() ([]*Asset, *Response, error) {
	apiEndpoint := "/api/assets/v1/assets/"
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	var assets []*Asset
	resp, err := s.client.Do(req, &assets)
	if err != nil {
		return nil, resp, err
	}
	return assets, resp, nil
}

func (s *AssetsService) Search(assetName string) ([]*Asset, *Response, error) {
	apiEndpoint := fmt.Sprintf("/api/assets/v1/assets/?search=%s", assetName)
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	var assets []*Asset
	resp, err := s.client.Do(req, &assets)
	if err != nil {
		return nil, resp, err
	}
	return assets, resp, nil
}

func (s *AssetsService) Get(AssetId string) (*Asset, *Response, error) {
	apiEndpoint := fmt.Sprintf("/api/assets/v1/assets/%s/", AssetId)
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	var asset *Asset
	resp, err := s.client.Do(req, &asset)
	if err != nil {
		return nil, resp, err
	}
	return asset, resp, nil
}
