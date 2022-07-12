package shipa_client

type ShipaTokenRequest struct {
	Password string `json:"password"`
}

type ShipaTokenResponse struct {
	Token string `json:"token"`
}

//type ShipaUserResponse struct {
//	Data  []ShipaUserInfo `json:"data"`
//}

type ShipaUserInfo struct {
	Id          string `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Deactivated bool   `json:"deactivated"`
	OrgId       string `json:"orgId"`
	OrgName     string `json:"orgName"`
}

/*
Shipa server response example

{
  "token": "18d2a552acd9f3d71721e385d551d24b2fa1917b",
}
{
  "data": [
    {
      "id": "626b9c07cba30d0001bd8125",
      "email": "haftomt71@gmail.com",
      "name": "bees wax",
      "deactivated": false,
      "phone": "",
      "Roles": [
        {
          "Name": "AllowAllOrg",
          "ContextType": "organization",
          "ContextValue": "b6667a8cc74c4d77ba1a6cb82af7a3c5",
          "Org": ""
        },
        {
          "Name": "TeamAdmin",
          "ContextType": "team",
          "ContextValue": "shipa-team",
          "Org": ""
        }
      ],
      "Permissions": [
        {
          "Name": "",
          "ContextType": "organization",
          "ContextValue": "b6667a8cc74c4d77ba1a6cb82af7a3c5",
          "Org": ""
        },
        {
          "Name": "app",
          "ContextType": "team",
          "ContextValue": "shipa-team",
          "Org": ""
        },
        {
          "Name": "cluster",
          "ContextType": "team",
          "ContextValue": "shipa-team",
          "Org": ""
        },
        {
          "Name": "team",
          "ContextType": "team",
          "ContextValue": "shipa-team",
          "Org": ""
        }
      ],
      "quota": {
        "limit": 8,
        "inuse": 0
      },
      "orgId": "b6667a8cc74c4d77ba1a6cb82af7a3c5",
      "orgName": "bees wax incorporated",
      "orgActivationDate": "2022-04-29T08:04:41.655Z",
      "orgAppLimit": 5
    }
  ]
}*/
