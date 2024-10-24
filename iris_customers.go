package iris

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CustomersResponse represents the response of the customers list endpoint
type CustomersResponse struct {
	Customers []Customer `json:"data"`
	ApiMeta
}

// CustomerResponse represents the response of a single customer from the /manage/customers/<id> endpoint
type CustomerResponse struct {
	Customer
	ApiMeta
}

// CustomerAddResponse represents the response of a single customer from the /add/customer endpoint
type CustomerAddResponse struct {
	Customer CustomerAddResponseObject
	ApiMeta
}

// CustomerContactAddResponse represents the response of a contact creation event
type CustomerContactAddResponse struct {
	Contact Contact
	ApiMeta
}

// Customer represents a single Customer object
type Customer struct {
	Contacts []Contact `json:"contacts"`
	CustomerDescription string `json:"customer_description"`
	CustomerID          int    `json:"customer_id"`
	CustomerName        string `json:"customer_name"`
	CustomerSLA         string `json:"customer_sla"`
	CustomerUUID        string `json:"customer_uuid"`
	CustomAttributes    map[string]interface{} `json:"custom_attributes"`
}

// UpdateCustomerRequest represents a struct for updating an existing customer
type UpdateCustomerRequest struct {
	CustomerName        string            `json:"customer_name"`
	CustomerDescription string            `json:"customer_description"`
	CustomerSLA         string            `json:"customer_sla"`
	CustomAttributes    map[string]interface{} `json:"custom_attributes"`
}

// AddCustomerRequest represents a struct for adding a new customer
type AddCustomerRequest struct {
	CustomerName        string      `json:"customer_name"`
	CustomerDescription string      `json:"customer_description"`
	CustomerSLA         string      `json:"customer_sla"`
	CustomAttributes    map[string]interface{} `json:"custom_attributes"`
}

// CustomerAddResponse represents a single Customer object after creation
type CustomerAddResponseObject struct {
	ClientUUID          string `json:"client_uuid"`
	CreationDate        string `json:"creation_date"`
	CustomAttributes    map[string]interface{}    `json:"custom_attributes"`
	CustomerDescription string `json:"customer_description"`
	CustomerID          int    `json:"customer_id"`
	CustomerName        string `json:"customer_name"`
	CustomerSLA         string `json:"customer_sla"`
	LastUpdateDate      string `json:"last_update_date"`
}

// Contact represents a contact object
type Contact struct {
	ClientID           int    `json:"client_id"`
	ContactEmail       string `json:"contact_email"`
	ContactMobilePhone string `json:"contact_mobile_phone"`
	ContactName        string `json:"contact_name"`
	ContactNote        string `json:"contact_note"`
	ContactRole        string `json:"contact_role"`
	ContactUUID        string `json:"contact_uuid"`
	ContactWorkPhone   string `json:"contact_work_phone"`
	CustomAttributes   map[string]interface{} `json:"custom_attributes"`
	ID                 int    `json:"id"`
}

// AddCustomerContactRequest represents a struct for adding a new contact
type AddCustomerContactRequest struct {
	ContactName        string `json:"contact_name"`
	ContactRole        string `json:"contact_role"`
	ContactEmail       string `json:"contact_email"`
	ContactMobilePhone string `json:"contact_mobile_phone"`
	ContactPhone			 string `json:"contact_work_phone"`
	ContactNote        string `json:"contact_note"`
	CustomAttributes   map[string]interface{} `json:"custom_attributes"`
}

// UpdateContactRequest represents a struct for updating an existing contact
type UpdateContactRequest struct {
	ContactName        string `json:"contact_name"`
	ContactRole        string `json:"contact_role"`
	ContactEmail       string `json:"contact_email"`
	ContactMobilePhone string `json:"contact_mobile_phone"`
	ContactWorkPhone   string `json:"contact_work_phone"`
	ContactNote        string `json:"contact_note"`
	CustomAttributes   map[string]interface{} `json:"custom_attributes"`
}


// GetCustomers gets a list of all customers from the /manage/customers/list endpoint.
// It returns a Customer slice containing all customers registered on Iris.
// If the request fails or the response cannot be decoded,
// an error is returned.
//
// Example usage:
//
//	client := api.NewAPIClient(httpClient, "https://api.example.com")
//	customers, err := client.GetCustomers()
//	if err != nil {
//	    log.Fatalf("Failed to get API version: %v", err)
//	}
//	for customer := range customers {
//		fmt.Printf(customer.CustomerName)
//	}
//
// Returns:
// - *CustomersResponse*: The response from the API containing customer informations.
// - error: An error if the request fails or the response cannot be decoded.
func (client *APIClient) GetCustomers() (*CustomersResponse, error) {
	builder := NewRequestBuilder().
		SetURL("/manage/customers/list").
		SetMethod(http.MethodGet).
		Build()

	req, err := client.DoRequest(*builder)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()


	if req.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", req.StatusCode)
	}

	var customerResponse CustomersResponse
	if err := json.NewDecoder(req.Body).Decode(&customerResponse); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &customerResponse, nil
}

// GetCustomer returns a single Customer object from the /manage/customers/<customer-id> endpoint.
// It returns a Customer slice containing all customers registered on Iris.
// If the request fails or the response cannot be decoded,
// an error is returned.
//
// Example usage:
//
//		client := api.NewAPIClient(httpClient, "https://api.example.com")
//		customer, err := client.GetCustomer(2)
//		if err != nil {
//		    log.Fatalf("Failed to get API version: %v", err)
//		}
//
//	 fmt.Printf("%s", customer.CustomerName)
//
// Returns:
// - *CustomerResponse*: The response from the API containing customer informations.
// - error: An error if the request fails or the response cannot be decoded.
func (client *APIClient) GetCustomer(id int) (*CustomerResponse, error) {
	builder := NewRequestBuilder().
		SetURL(fmt.Sprintf("/manage/customers/%d", id)).
		SetMethod(http.MethodGet).
		Build()

	req, err := client.DoRequest(*builder)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", req.StatusCode)
	}

	type CustomerResponseWrapper struct {
		CustomerResponse `json:"data"`
		ApiMeta
	}

	var customerResponseWrapper CustomerResponseWrapper
	if err := json.NewDecoder(req.Body).Decode(&customerResponseWrapper); err != nil {
		return nil, err
	}

	customerResponse := CustomerResponse{
		Customer: customerResponseWrapper.Customer,
		ApiMeta:  customerResponseWrapper.ApiMeta,
	}

	return &customerResponse, nil
}

func (client *APIClient) DeleteCustomer(id int) error {
	builder := NewRequestBuilder().
		SetURL(fmt.Sprintf("/manage/customers/delete/%d", id)).
		SetMethod(http.MethodPost).
		AddHeader("Content-Type", "application/json").
		Build()

	req, err := client.DoRequest(*builder)
	if err != nil {
		return err
	}
	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", req.StatusCode)
	}

	return nil
}

// AddCustomer adds a customer to iris throught the /manage/customers/add endpoint.
// It returns a the customer object that was just created.
// If the request fails or the response cannot be decoded,
// an error is returned.
//
// Example usage:
//
//	client := api.NewAPIClient(httpClient, "https://api.example.com")
//	customers, err := client.GetCustomers()
//	if err != nil {
//	    log.Fatalf("Failed to get API version: %v", err)
//	}
//	for customer := range customers {
//		fmt.Printf(customer.CustomerName)
//	}
//
// Returns:
// - *CustomerResponse*: The response from the API containing customer informations.
// - error: An error if the request fails or the response cannot be decoded.
func (client *APIClient) AddCustomer(customer AddCustomerRequest) (*CustomerAddResponse, error) {
	jsondata, err := json.Marshal(customer)
	if err != nil {
		return nil, err
	}

	builder := NewRequestBuilder().
		SetURL("/manage/customers/add").
		SetMethod(http.MethodPost).
		AddHeader("Content-Type", "application/json").
		SetBody(jsondata).
		Build()

	req, err := client.DoRequest(*builder)
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", req.StatusCode)
	}

	type CustomerResponseWrapper struct {
		Customer CustomerAddResponseObject `json:"data"`
		ApiMeta
	}

	var customerResponseWrapper CustomerResponseWrapper
	if err := json.NewDecoder(req.Body).Decode(&customerResponseWrapper); err != nil {
		return nil, err
	}

	customerAddResponse := CustomerAddResponse{
		ApiMeta:  customerResponseWrapper.ApiMeta,
		Customer: customerResponseWrapper.Customer,
	}

	return &customerAddResponse, nil
}

func (client *APIClient) UpdateCustomer(id int, customer UpdateCustomerRequest) (*CustomerResponse, error) {
	jsondata, err := json.Marshal(customer)
	if err != nil {
		return nil, err
	}


	builder := NewRequestBuilder().
		SetURL(fmt.Sprintf("/manage/customers/update/%d", id)).
		SetMethod(http.MethodPost).
		AddHeader("Content-Type", "application/json").
		SetBody(jsondata).
		Build()

	req, err := client.DoRequest(*builder)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", req.StatusCode)
	}

	type CustomerResponseWrapper struct {
		Customer Customer `json:"data"`
		ApiMeta
	}

	var customerResponseWrapper CustomerResponseWrapper
	if err := json.NewDecoder(req.Body).Decode(&customerResponseWrapper); err != nil {
		return nil, err
	}


	customerResponse := CustomerResponse{
		ApiMeta:  customerResponseWrapper.ApiMeta,
		Customer: customerResponseWrapper.Customer,
	}


	return &customerResponse, nil
}



// {
//
//     "customer_name": "New customer",
//     "customer_description": "New customer description",
//     "customer_sla": "New customer SLA",
//     "custom_attributes": { }
//
// }

// AddCustomerContact adds a customer contact to iris throught the /manage/customers/{customer_id}/contacts/add endpoint.
// It returns a the customer object that was just created.
// If the request fails or the response cannot be decoded,
// an error is returned.
//
// Example usage:
//
//	client := api.NewAPIClient(httpClient, "https://api.example.com")
//	customers, err := client.GetCustomers()
//	if err != nil {
//	    log.Fatalf("Failed to get API version: %v", err)
//	}
//	for customer := range customers {
//		fmt.Printf(customer.CustomerName)
//	}
//
// Returns:
// - *CustomerResponse*: The response from the API containing customer informations.
// - error: An error if the request fails or the response cannot be decoded.
func (client *APIClient) AddCustomerContact(customerId int, contact AddCustomerContactRequest) (*CustomerContactAddResponse, error) {
	jsondata, err := json.Marshal(contact)
	if err != nil {
		return nil, err
	}

	builder := NewRequestBuilder().
		SetURL(fmt.Sprint("/manage/customers/", customerId,"/contacts/add")).
		SetMethod(http.MethodPost).
		AddHeader("Content-Type", "application/json").
		SetBody(jsondata).
		Build()

	req, err := client.DoRequest(*builder)
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", req.StatusCode)
	}

	type CustomerContactResponseWrapper struct {
		Contact Contact `json:"data"`
		ApiMeta
	}

	var customerContactResponseWrapper CustomerContactResponseWrapper
	if err := json.NewDecoder(req.Body).Decode(&customerContactResponseWrapper); err != nil {
		return nil, err
	}

	customerContactAddResponse := CustomerContactAddResponse {
		ApiMeta:  customerContactResponseWrapper.ApiMeta,
		Contact: customerContactResponseWrapper.Contact,
	}
	return &customerContactAddResponse, nil
}

func (client *APIClient) UpdateCustomerContact(customerId int, contactId int, contact UpdateContactRequest) (*CustomerContactAddResponse, error) {
	jsondata, err := json.Marshal(contact)
	if err != nil {
		return nil, err
	}


	builder := NewRequestBuilder().
		SetURL(fmt.Sprintf("/manage/customers/%d/contacts/%d/update", customerId, contactId)).
		SetMethod(http.MethodPost).
		AddHeader("Content-Type", "application/json").
		SetBody(jsondata).
		Build()

	req, err := client.DoRequest(*builder)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", req.StatusCode)
	}

	var customerContactResponse CustomerContactAddResponse
	if err := json.NewDecoder(req.Body).Decode(&customerContactResponse); err != nil {
		return nil, err
	}

	return &customerContactResponse, nil
}

func (client *APIClient) DeleteCustomerContact(customerId int, contactId int) error {
	builder := NewRequestBuilder().
		SetURL(fmt.Sprintf("/manage/customers/%d/contacts/%d/delete", customerId, contactId)).
		SetMethod(http.MethodPost).
		AddHeader("Content-Type", "application/json").
		Build()

	req, err := client.DoRequest(*builder)
	if err != nil {
		return err
	}
	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		return err
	}

	return nil
}

