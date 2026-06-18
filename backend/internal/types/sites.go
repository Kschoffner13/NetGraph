package types

type Site struct {
	siteID   int
	siteCode string
	province string
	lat      float64
	long     float64
	network  string
}

type Circuit struct {
	site           Site
	capcity_Gbps   float64
	bandwidth_Gbps float64
}