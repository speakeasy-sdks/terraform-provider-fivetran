// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"context"
	"fmt"
	"net/http"
	"terraform/internal/sdk/pkg/models/operations"
	"terraform/internal/sdk/pkg/utils"
)

type connectorSchemaManagement struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newConnectorSchemaManagement(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *connectorSchemaManagement {
	return &connectorSchemaManagement{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// ConnectorColumnConfig - Retrieve Source Table Columns Config
// Returns the source table columns config for an existing connector within your Fivetran account
func (s *connectorSchemaManagement) ConnectorColumnConfig(ctx context.Context, request operations.ConnectorColumnConfigRequest) (*operations.ConnectorColumnConfigResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/connectors/{connectorId}/schemas/{schema}/tables/{table}/columns", request, nil)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	utils.PopulateHeaders(ctx, req, request)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ConnectorColumnConfigResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ConnectorColumnConfig200ApplicationJSONAny = out
		}
	}

	return res, nil
}

// ConnectorSchemaConfig - Retrieve a Connector Schema Config
// Returns the connector schema config for an existing connector within your Fivetran account
func (s *connectorSchemaManagement) ConnectorSchemaConfig(ctx context.Context, request operations.ConnectorSchemaConfigRequest) (*operations.ConnectorSchemaConfigResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/connectors/{connectorId}/schemas", request, nil)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	utils.PopulateHeaders(ctx, req, request)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ConnectorSchemaConfigResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ConnectorSchemaConfig200ApplicationJSONAny = out
		}
	}

	return res, nil
}

// ModifyConnectorColumnConfig - Modify a Connector Column Config
// Updates the column config within your table for an existing connector within your Fivetran account
func (s *connectorSchemaManagement) ModifyConnectorColumnConfig(ctx context.Context, request operations.ModifyConnectorColumnConfigRequest) (*operations.ModifyConnectorColumnConfigResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/connectors/{connectorId}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}", request, nil)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "ColumnUpdateRequest", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PATCH", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	utils.PopulateHeaders(ctx, req, request)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ModifyConnectorColumnConfigResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ModifyConnectorColumnConfig200ApplicationJSONAny = out
		}
	}

	return res, nil
}

// ModifyConnectorDatabaseSchemaConfig - Modify a Connector Database Schema Config
// Updates the database schema config for an existing connector within your Fivetran account (for a single schema within a connector with multiple schemas)
func (s *connectorSchemaManagement) ModifyConnectorDatabaseSchemaConfig(ctx context.Context, request operations.ModifyConnectorDatabaseSchemaConfigRequest) (*operations.ModifyConnectorDatabaseSchemaConfigResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/connectors/{connectorId}/schemas/{schemaName}", request, nil)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "SchemaUpdateRequest", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PATCH", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	utils.PopulateHeaders(ctx, req, request)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ModifyConnectorDatabaseSchemaConfigResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ModifyConnectorDatabaseSchemaConfig200ApplicationJSONAny = out
		}
	}

	return res, nil
}

// ModifyConnectorSchemaConfig - Modify a Connector Schema Config
// Updates the schema config for an existing connector within your Fivetran account (for a single schema for a connector with multiple schemas)
func (s *connectorSchemaManagement) ModifyConnectorSchemaConfig(ctx context.Context, request operations.ModifyConnectorSchemaConfigRequest) (*operations.ModifyConnectorSchemaConfigResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/connectors/{connectorId}/schemas", request, nil)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "StandardConfigUpdateRequest", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PATCH", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	utils.PopulateHeaders(ctx, req, request)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ModifyConnectorSchemaConfigResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ModifyConnectorSchemaConfig200ApplicationJSONAny = out
		}
	}

	return res, nil
}

// ModifyConnectorTableConfig - Modify a Connector Table Config
// Updates the table config within your database schema for an existing connector within your Fivetran account
func (s *connectorSchemaManagement) ModifyConnectorTableConfig(ctx context.Context, request operations.ModifyConnectorTableConfigRequest) (*operations.ModifyConnectorTableConfigResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/connectors/{connectorId}/schemas/{schemaName}/tables/{tableName}", request, nil)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "TableUpdateRequest", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PATCH", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	utils.PopulateHeaders(ctx, req, request)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ModifyConnectorTableConfigResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ModifyConnectorTableConfig200ApplicationJSONAny = out
		}
	}

	return res, nil
}

// ReloadConnectorSchemaConfig - Reload a Connector Schema Config
// Reloads the connector schema config for an existing connector within your Fivetran account
func (s *connectorSchemaManagement) ReloadConnectorSchemaConfig(ctx context.Context, request operations.ReloadConnectorSchemaConfigRequest) (*operations.ReloadConnectorSchemaConfigResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/connectors/{connectorId}/schemas/reload", request, nil)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "ReloadStandardConfigRequest", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	utils.PopulateHeaders(ctx, req, request)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ReloadConnectorSchemaConfigResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ReloadConnectorSchemaConfig200ApplicationJSONAny = out
		}
	}

	return res, nil
}

// ResyncTables - Re-sync Connector Table Data
// Triggers a historical sync of all data for multiple schema tables within a connector. This action does not override the standard sync frequency you defined in the Fivetran dashboard.
func (s *connectorSchemaManagement) ResyncTables(ctx context.Context, request operations.ResyncTablesRequest) (*operations.ResyncTablesResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/connectors/{connectorId}/schemas/tables/resync", request, nil)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "RequestBody", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	utils.PopulateHeaders(ctx, req, request)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ResyncTablesResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ResyncTables200ApplicationJSONAny = out
		}
	}

	return res, nil
}
