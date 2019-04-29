package types


type UToC struct {
	DN       string         `json:"dn"        gencodec:"required"`
	ET       string         `json:"et"        gencodec:"required"`
	CAS      string	        `json:"cas"        gencodec:"required"`
}

func (u *UToC) Dn() string            { return u.DN }
func (u *UToC) Et() string            { return u.ET }
func (u *UToC) Cas() string            { return u.CAS }


