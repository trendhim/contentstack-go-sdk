// Package delivery provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package delivery

// AccessToken defines model for access_token.
type AccessToken string

// ApiKey defines model for api_key.
type ApiKey string

// AssetUid defines model for asset_uid.
type AssetUid string

// ContentTypeUid defines model for content_type_uid.
type ContentTypeUid string

// ContentTypeUidQuery defines model for content_type_uid_query.
type ContentTypeUidQuery string

// EntryUid defines model for entry_uid.
type EntryUid string

// Environment defines model for environment.
type Environment string

// GlobalFieldUid defines model for global_field_uid.
type GlobalFieldUid string

// IncludeCount defines model for include_count.
type IncludeCount string

// IncludeDimension defines model for include_dimension.
type IncludeDimension string

// IncludeFallback defines model for include_fallback.
type IncludeFallback string

// Init defines model for init.
type Init string

// Locale defines model for locale.
type Locale string

// PaginationToken defines model for pagination_token.
type PaginationToken string

// Query defines model for query.
type Query string

// RelativeUrls defines model for relative_urls.
type RelativeUrls string

// StartFrom defines model for start_from.
type StartFrom string

// SyncToken defines model for sync_token.
type SyncToken string

// Type defines model for type.
type Type string

// GetallassetsParams defines parameters for Getallassets.
type GetallassetsParams struct {
	// Enter the name of the environment if you want to retrieve the assets published in a particular environment.
	Environment string `json:"environment"`

	// Enter 'true' to include the published asset details from the fallback locale when the specified locale does not contain published content.
	IncludeFallback string `json:"include_fallback"`

	// Enter 'true' to include the relative URLs of the assets in the response.
	RelativeUrls RelativeUrls `json:"relative_urls"`

	// Enter 'true' to include the dimensions (height and width) of the image in the response. Supported image types: JPG, GIF, PNG, WebP, BMP, TIFF, SVG, and PSD.
	IncludeDimension IncludeDimension `json:"include_dimension"`
	ApiKey           string           `json:"api_key"`
	AccessToken      string           `json:"access_token"`
}

// GetasingleassetParams defines parameters for Getasingleasset.
type GetasingleassetParams struct {
	// Enter the name of the environment if you want to retrieve an asset published in a particular environment.
	Environment string `json:"environment"`

	// Specify the version number of the asset that you want to retrieve. To retrieve a specific version, keep the environment parameter blank.&nbsp;If the version is not specified, the details of the latest version will be retrieved.
	Version *float32 `json:"version,omitempty"`

	// Enter 'true' to include published asset details from the fallback locale when the specified locale does not contain published information.
	IncludeFallback string `json:"include_fallback"`

	// Enter 'true' to include the dimensions (height and width)&nbsp;of the image&nbsp;in the response.&nbsp;Supported image&nbsp;types: JPG, GIF, PNG, WebP, BMP, TIFF, SVG, and PSD.
	IncludeDimension string `json:"include_dimension"`

	// Enter 'true' to include the relative URL of the asset in the response.
	RelativeUrls string `json:"relative_urls"`
	ApiKey       string `json:"api_key"`
	AccessToken  string `json:"access_token"`
}

// GetallcontenttypesParams defines parameters for Getallcontenttypes.
type GetallcontenttypesParams struct {
	// Set this to 'true' to include in response the total count of content types available in your stack.
	IncludeCount IncludeCount `json:"include_count"`
	ApiKey       ApiKey       `json:"api_key"`
	AccessToken  AccessToken  `json:"access_token"`
}

// GetasinglecontenttypeParams defines parameters for Getasinglecontenttype.
type GetasinglecontenttypeParams struct {
	ApiKey      string `json:"api_key"`
	AccessToken string `json:"access_token"`
}

// GetallentriesParams defines parameters for Getallentries.
type GetallentriesParams struct {
	// Enter the name of the environment of which the entries needs to be included.
	Environment Environment `json:"environment"`

	// Enter the code of the language of which the entries needs to be included. Only the entries published in this locale will be displayed.
	Locale Locale `json:"locale"`

	// Enter 'true' to include the published localized content from the fallback locale when the specified locale does not contain published content.
	IncludeFallback IncludeFallback `json:"include_fallback"`
	ApiKey          string          `json:"api_key"`
	AccessToken     string          `json:"access_token"`
}

// QueriesParams defines parameters for Queries.
type QueriesParams struct {
	// Enter the name of the environment of which the entries needs to be included.
	Environment Environment `json:"environment"`

	// Enter the code of the language of which the entries needs to be included. Only the entries published in this locale will be displayed.
	Locale Locale `json:"locale"`

	// Enter the actual query that will be executed to retrieve entries. This query should be in JSON format.
	Query       *Query  `json:"query,omitempty"`
	ApiKey      *string `json:"api_key,omitempty"`
	AccessToken *string `json:"access_token,omitempty"`
}

// GetasingleentryParams defines parameters for Getasingleentry.
type GetasingleentryParams struct {
	// Enter the name of the environment of which the entries needs to be included.
	Environment string `json:"environment"`

	// Enter the code of the language of which you want to include&nbsp;the entries. Only the published localized entries will be displayed.
	Locale string `json:"locale"`

	// Enter 'true' to include the published localized content from the fallback locale when the specified locale does not contain published content.
	IncludeFallback string `json:"include_fallback"`
	ApiKey          string `json:"api_key"`
	AccessToken     string `json:"access_token"`
}

// GetallglobalfieldsParams defines parameters for Getallglobalfields.
type GetallglobalfieldsParams struct {
	ApiKey      string `json:"api_key"`
	AccessToken string `json:"access_token"`
}

// GetasingleglobalfieldParams defines parameters for Getasingleglobalfield.
type GetasingleglobalfieldParams struct {
	ApiKey      string `json:"api_key"`
	AccessToken string `json:"access_token"`
}

// SynchronizationParams defines parameters for Synchronization.
type SynchronizationParams struct {
	// Enter ‘true’ to perform a complete sync of all your app data.
	Init Init `json:"init"`

	// Enter the environment to retrieve and sync the content published on a specific environment.
	Environment    string               `json:"environment"`
	ContentTypeUid *ContentTypeUidQuery `json:"content_type_uid,omitempty"`

	// Enter the locale to retrieve and sync the content published on a specific locale.
	Locale string `json:"locale"`

	// Enter the type(s) of content you want to retrieve and sync. You can pass multiple types as comma-separated values.
	Type Type `json:"type"`

	// Specify the start date, if you want to retrieve and sync data starting from a specific date.
	StartFrom StartFrom `json:"start_from"`

	// Enter the pagination token that you received in the response body of the previous sync process.
	PaginationToken *PaginationToken `json:"pagination_token,omitempty"`

	// Enter the sync token that you received in the response body of the previous completed Synchronization process to get the delta updates
	SyncToken   *SyncToken `json:"sync_token,omitempty"`
	ApiKey      string     `json:"api_key"`
	AccessToken string     `json:"access_token"`
}

