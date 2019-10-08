package ssl

type Sslprofile struct {
	Builtin                           interface{} `json:"builtin,omitempty"`
	Ciphername                        string      `json:"ciphername,omitempty"`
	Cipherpriority                    int         `json:"cipherpriority,omitempty"`
	Cipherredirect                    string      `json:"cipherredirect,omitempty"`
	Cipherurl                         string      `json:"cipherurl,omitempty"`
	Cleartextport                     int         `json:"cleartextport,omitempty"`
	Clientauth                        string      `json:"clientauth,omitempty"`
	Clientauthuseboundcachain         string      `json:"clientauthuseboundcachain,omitempty"`
	Clientcert                        string      `json:"clientcert,omitempty"`
	Commonname                        string      `json:"commonname,omitempty"`
	Crlcheck                          string      `json:"crlcheck,omitempty"`
	Denysslreneg                      string      `json:"denysslreneg,omitempty"`
	Dh                                string      `json:"dh,omitempty"`
	Dhcount                           int         `json:"dhcount,omitempty"`
	Dhekeyexchangewithpsk             string      `json:"dhekeyexchangewithpsk,omitempty"`
	Dhfile                            string      `json:"dhfile,omitempty"`
	Dhkeyexpsizelimit                 string      `json:"dhkeyexpsizelimit,omitempty"`
	Dropreqwithnohostheader           string      `json:"dropreqwithnohostheader,omitempty"`
	Encrypttriggerpktcount            int         `json:"encrypttriggerpktcount,omitempty"`
	Ersa                              string      `json:"ersa,omitempty"`
	Ersacount                         int         `json:"ersacount,omitempty"`
	Hsts                              string      `json:"hsts,omitempty"`
	Includesubdomains                 string      `json:"includesubdomains,omitempty"`
	Insertionencoding                 string      `json:"insertionencoding,omitempty"`
	Invoke                            bool        `json:"invoke,omitempty"`
	Labeltype                         string      `json:"labeltype,omitempty"`
	Maxage                            int         `json:"maxage,omitempty"`
	Name                              string      `json:"name,omitempty"`
	Nonfipsciphers                    string      `json:"nonfipsciphers,omitempty"`
	Ocspcheck                         string      `json:"ocspcheck,omitempty"`
	Ocspstapling                      string      `json:"ocspstapling,omitempty"`
	Preload                           string      `json:"preload,omitempty"`
	Prevsessionkeylifetime            int         `json:"prevsessionkeylifetime,omitempty"`
	Pushenctrigger                    string      `json:"pushenctrigger,omitempty"`
	Pushenctriggertimeout             int         `json:"pushenctriggertimeout,omitempty"`
	Pushflag                          int         `json:"pushflag,omitempty"`
	Quantumsize                       string      `json:"quantumsize,omitempty"`
	Redirectportrewrite               string      `json:"redirectportrewrite,omitempty"`
	Sendclosenotify                   string      `json:"sendclosenotify,omitempty"`
	Serverauth                        string      `json:"serverauth,omitempty"`
	Service                           int         `json:"service,omitempty"`
	Sessionkeylifetime                int         `json:"sessionkeylifetime,omitempty"`
	Sessionticket                     string      `json:"sessionticket,omitempty"`
	Sessionticketkeydata              string      `json:"sessionticketkeydata,omitempty"`
	Sessionticketkeyrefresh           string      `json:"sessionticketkeyrefresh,omitempty"`
	Sessionticketlifetime             int         `json:"sessionticketlifetime,omitempty"`
	Sessreuse                         string      `json:"sessreuse,omitempty"`
	Sesstimeout                       int         `json:"sesstimeout,omitempty"`
	Skipcaname                        bool        `json:"skipcaname,omitempty"`
	Skipclientcertpolicycheck         string      `json:"skipclientcertpolicycheck,omitempty"`
	Snicert                           bool        `json:"snicert,omitempty"`
	Snienable                         string      `json:"snienable,omitempty"`
	Ssl3                              string      `json:"ssl3,omitempty"`
	Sslimaxsessperserver              int         `json:"sslimaxsessperserver,omitempty"`
	Sslinterception                   string      `json:"sslinterception,omitempty"`
	Ssliocspcheck                     string      `json:"ssliocspcheck,omitempty"`
	Sslireneg                         string      `json:"sslireneg,omitempty"`
	Ssliverifyservercertforreuse      string      `json:"ssliverifyservercertforreuse,omitempty"`
	Ssllogprofile                     string      `json:"ssllogprofile,omitempty"`
	Sslpfobjecttype                   int         `json:"sslpfobjecttype,omitempty"`
	Sslprofiletype                    string      `json:"sslprofiletype,omitempty"`
	Sslredirect                       string      `json:"sslredirect,omitempty"`
	Ssltriggertimeout                 int         `json:"ssltriggertimeout,omitempty"`
	Strictcachecks                    string      `json:"strictcachecks,omitempty"`
	Strictsigdigestcheck              string      `json:"strictsigdigestcheck,omitempty"`
	Tls1                              string      `json:"tls1,omitempty"`
	Tls11                             string      `json:"tls11,omitempty"`
	Tls12                             string      `json:"tls12,omitempty"`
	Tls13                             string      `json:"tls13,omitempty"`
	Tls13sessionticketsperauthcontext int         `json:"tls13sessionticketsperauthcontext,omitempty"`
	Zerorttearlydata                  string      `json:"zerorttearlydata,omitempty"`
}
