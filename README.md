# DFIR-IRIS go library

Library to support DFIR-IRIS tasks using go.

**NOTE**: This library is pretty much work in progress. 
A lot of functions are missing and will potentially added over time.


## Functionality

- [x] Customer management
  - Get Customer
  - Add Customer
  - Update Customer
  - Delete Customer
  - Get Contact
  - Add Contact
  - Update Contact
  - Delete Contact
- [ ] Template management
- [ ] User management
- [ ] Module management
- [ ] Case management

## Basic setup

```
conf := iris.GetInstance()
conf.BaseUrl = "https://iris.lab"
conf.AuthToken = "{ReplaceMe}"

authStrategy := &iris.ApiKeyAuth{ApiKey: conf.AuthToken}

irisClient := &iris.APIClient{
        AuthStrategy: authStrategy,
        BaseURL:      conf.BaseUrl,
        Client:       *iris.NewConfiguredHttpClient(iris.ClientConfig{IgnoreTLS: true}),
    }
```
