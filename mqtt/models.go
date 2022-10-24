package mqtt

type ACL struct {
	ACLType  string `json:"acltype"`
	Topic    string `json:"topic"`
	Priority int    `json:"priority"`
	Allow    bool   `json:"allow"`
}

type CreateRoleReq struct {
	Command  string  `json:"command"`
	Rolename *string `json:"rolename"`
	ACLs     []*ACL  `json:"acls"`
}

type Role struct {
	Rolename *string `json:"rolename"`
	Priority int     `json:"priority"`
}

type CreateClientReq struct {
	Command  string  `json:"command"`
	Username *string `json:"username"`
	Password *string `json:"password"`
	Roles    []*Role `json:"roles"`
}

type ChangeACLReq struct {
	Command  string  `json:"command"`
	Rolename *string `json:"rolename"`
	ACLtype  string  `json:"acltype"`
	Topic    string  `json:"topic"`
	Priority int     `json:"priority,omitempty"`
	Allow    bool    `json:"allow,omitempty"`
}

type EnvelopeRole struct {
	Commands []*CreateRoleReq `json:"commands"`
}

type EnvelopeClient struct {
	Commands []*CreateClientReq `json:"commands"`
}

type EnvelopeACL struct {
	Commands []*ChangeACLReq `json:"commands"`
}
