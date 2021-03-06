// Code generated by protoc-gen-goext. DO NOT EDIT.

package containerregistry

func (m *ScanRequest) SetImageId(v string) {
	m.ImageId = v
}

func (m *ScanMetadata) SetScanResultId(v string) {
	m.ScanResultId = v
}

func (m *GetScanResultRequest) SetScanResultId(v string) {
	m.ScanResultId = v
}

func (m *GetLastScanResultRequest) SetImageId(v string) {
	m.ImageId = v
}

type ListScanResultsRequest_Id = isListScanResultsRequest_Id

func (m *ListScanResultsRequest) SetId(v ListScanResultsRequest_Id) {
	m.Id = v
}

func (m *ListScanResultsRequest) SetImageId(v string) {
	m.Id = &ListScanResultsRequest_ImageId{
		ImageId: v,
	}
}

func (m *ListScanResultsRequest) SetRepositoryId(v string) {
	m.Id = &ListScanResultsRequest_RepositoryId{
		RepositoryId: v,
	}
}

func (m *ListScanResultsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListScanResultsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListScanResultsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListScanResultsRequest) SetOrderBy(v string) {
	m.OrderBy = v
}

func (m *ListScanResultsResponse) SetScanResults(v []*ScanResult) {
	m.ScanResults = v
}

func (m *ListScanResultsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *ListVulnerabilitiesRequest) SetScanResultId(v string) {
	m.ScanResultId = v
}

func (m *ListVulnerabilitiesRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListVulnerabilitiesRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListVulnerabilitiesRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListVulnerabilitiesRequest) SetOrderBy(v string) {
	m.OrderBy = v
}

func (m *ListVulnerabilitiesResponse) SetVulnerabilities(v []*Vulnerability) {
	m.Vulnerabilities = v
}

func (m *ListVulnerabilitiesResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
