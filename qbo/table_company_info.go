package qbo

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableQBOCompanyInfo(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "qbo_company_info",
		Description: "Company Information from QuickBooks Online",
		List: &plugin.ListConfig{
			Hydrate: listQBOCompanyInfo,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:       "id",
					Require:    plugin.Optional,
					CacheMatch: "exact",
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "Unique identifier for this object. Sort order is ASC by default.",
				Transform:   transform.FromField("ID").NullIfZero(),
			},
			{
				Name:        "sync_token",
				Type:        proto.ColumnType_STRING,
				Description: `Version number of the object. It is used to lock an object for use by one app at a time. As soon as an application modifies an object, its SyncToken is incremented. Attempts to modify an object specifying an older SyncToken fails. Only the latest version of the object is maintained by QuickBooks Online.`,
				Transform:   transform.FromField("SyncToken").NullIfZero(),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: `The name of the company.`,
				Transform:   transform.FromField("CompanyName").NullIfZero(),
			},
			{
				Name: "address_line1",
				Type: proto.ColumnType_STRING,
				Description: `Company Address as described in preference.
				If a physical address is updated from within the transaction object, the QuickBooks Online API flows individual address components differently into the Line elements of the transaction response then when the transaction was first created:
				Line1 and Line2 elements are populated with the customer name and company name.
				Original Line1 through Line 5 contents, City, SubDivisionCode, and PostalCode flow into Line3 through Line5as a free format strings.`,
				Transform: transform.FromField("CompanyAddress.Line1").NullIfZero(),
			},
			{
				Name: "address_line2",
				Type: proto.ColumnType_STRING,
				Description: `Company Address as described in preference.
				If a physical address is updated from within the transaction object, the QuickBooks Online API flows individual address components differently into the Line elements of the transaction response then when the transaction was first created:
				Line1 and Line2 elements are populated with the customer name and company name.
				Original Line1 through Line 5 contents, City, SubDivisionCode, and PostalCode flow into Line3 through Line5as a free format strings.`,
				Transform: transform.FromField("CompanyAddress.Line2").NullIfZero(),
			},
			{
				Name: "address_line3",
				Type: proto.ColumnType_STRING,
				Description: `Company Address as described in preference.
				If a physical address is updated from within the transaction object, the QuickBooks Online API flows individual address components differently into the Line elements of the transaction response then when the transaction was first created:
				Line1 and Line2 elements are populated with the customer name and company name.
				Original Line1 through Line 5 contents, City, SubDivisionCode, and PostalCode flow into Line3 through Line5as a free format strings.`,
				Transform: transform.FromField("CompanyAddress.Line3").NullIfZero(),
			},
			{
				Name: "address_line4",
				Type: proto.ColumnType_STRING,
				Description: `Company Address as described in preference.
				If a physical address is updated from within the transaction object, the QuickBooks Online API flows individual address components differently into the Line elements of the transaction response then when the transaction was first created:
				Line1 and Line2 elements are populated with the customer name and company name.
				Original Line1 through Line 5 contents, City, SubDivisionCode, and PostalCode flow into Line3 through Line5as a free format strings.`,
				Transform: transform.FromField("CompanyAddress.Line4").NullIfZero(),
			},
			{
				Name: "address_line5",
				Type: proto.ColumnType_STRING,
				Description: `Company Address as described in preference.
				If a physical address is updated from within the transaction object, the QuickBooks Online API flows individual address components differently into the Line elements of the transaction response then when the transaction was first created:
				Line1 and Line2 elements are populated with the customer name and company name.
				Original Line1 through Line 5 contents, City, SubDivisionCode, and PostalCode flow into Line3 through Line5as a free format strings.`,
				Transform: transform.FromField("CompanyAddress.Line5").NullIfZero(),
			},
			{
				Name:        "address_city",
				Type:        proto.ColumnType_STRING,
				Description: `Company Address as described in preference. City name.`,
				Transform:   transform.FromField("CompanyAddress.City").NullIfZero(),
			},
			{
				Name:        "address_region",
				Type:        proto.ColumnType_STRING,
				Description: `Company Address as described in preference. Region within a country. For example, state name for USA, province name for Canada.`,
				Transform:   transform.FromField("CompanyAddress.Region").NullIfZero(),
			},
			{
				Name:        "address_country",
				Type:        proto.ColumnType_STRING,
				Description: `Company Address as described in preference. Country name. For international addresses - countries should be passed as 3 ISO alpha-3 characters or the full name of the country.`,
				Transform:   transform.FromField("CompanyAddress.Country").NullIfZero(),
			},
			{
				Name:        "address_latitude",
				Type:        proto.ColumnType_STRING,
				Description: `Company Address as described in preference. Latitude coordinate of Geocode (Geospacial Entity Object Code). INVALIDis returned for invalid addresses.`,
				Transform:   transform.FromField("CompanyAddress.Latitude").NullIfZero(),
			},
			{
				Name:        "address_longitude",
				Type:        proto.ColumnType_STRING,
				Description: `Company Address as described in preference.Longitude coordinate of Geocode (Geospacial Entity Object Code). INVALIDis returned for invalid addresses.`,
				Transform:   transform.FromField("CompanyAddress.Longitude").NullIfZero(),
			},
			{
				Name: "legal_address_line1",
				Type: proto.ColumnType_STRING,
				Description: `Legal Address given to the government for any government communication.
				If a physical address is updated from within the transaction object, the QuickBooks Online API flows individual address components differently into the Line elements of the transaction response then when the transaction was first created:
				Line1 and Line2 elements are populated with the customer name and company name.
				Original Line1 through Line 5 contents, City, SubDivisionCode, and PostalCode flow into Line3 through Line5as a free format strings.`,
				Transform: transform.FromField("LegalAddress.Line1").NullIfZero(),
			},
			{
				Name: "legal_address_line2",
				Type: proto.ColumnType_STRING,
				Description: `Legal Address given to the government for any government communication.
				If a physical address is updated from within the transaction object, the QuickBooks Online API flows individual address components differently into the Line elements of the transaction response then when the transaction was first created:
				Line1 and Line2 elements are populated with the customer name and company name.
				Original Line1 through Line 5 contents, City, SubDivisionCode, and PostalCode flow into Line3 through Line5as a free format strings.`,
				Transform: transform.FromField("LegalAddress.Line2").NullIfZero(),
			},
			{
				Name: "legal_address_line3",
				Type: proto.ColumnType_STRING,
				Description: `Legal Address given to the government for any government communication.
				If a physical address is updated from within the transaction object, the QuickBooks Online API flows individual address components differently into the Line elements of the transaction response then when the transaction was first created:
				Line1 and Line2 elements are populated with the customer name and company name.
				Original Line1 through Line 5 contents, City, SubDivisionCode, and PostalCode flow into Line3 through Line5as a free format strings.`,
				Transform: transform.FromField("LegalAddress.Line3").NullIfZero(),
			},
			{
				Name: "legal_address_line4",
				Type: proto.ColumnType_STRING,
				Description: `Legal Address given to the government for any government communication.
				If a physical address is updated from within the transaction object, the QuickBooks Online API flows individual address components differently into the Line elements of the transaction response then when the transaction was first created:
				Line1 and Line2 elements are populated with the customer name and company name.
				Original Line1 through Line 5 contents, City, SubDivisionCode, and PostalCode flow into Line3 through Line5as a free format strings.`,
				Transform: transform.FromField("LegalAddress.Line4").NullIfZero(),
			},
			{
				Name: "legal_address_line5",
				Type: proto.ColumnType_STRING,
				Description: `Legal Address given to the government for any government communication.
				If a physical address is updated from within the transaction object, the QuickBooks Online API flows individual address components differently into the Line elements of the transaction response then when the transaction was first created:
				Line1 and Line2 elements are populated with the customer name and company name.
				Original Line1 through Line 5 contents, City, SubDivisionCode, and PostalCode flow into Line3 through Line5as a free format strings.`,
				Transform: transform.FromField("LegalAddress.Line5").NullIfZero(),
			},
			{
				Name:        "legal_address_city",
				Type:        proto.ColumnType_STRING,
				Description: `Legal Address given to the government for any government communication. City name.`,
				Transform:   transform.FromField("LegalAddress.City").NullIfZero(),
			},
			{
				Name:        "legal_address_region",
				Type:        proto.ColumnType_STRING,
				Description: `Legal Address given to the government for any government communication. Region within a country. For example, state name for USA, province name for Canada.`,
				Transform:   transform.FromField("LegalAddress.Region").NullIfZero(),
			},
			{
				Name:        "legal_address_country",
				Type:        proto.ColumnType_STRING,
				Description: `Legal Address given to the government for any government communication. Country name. For international addresses - countries should be passed as 3 ISO alpha-3 characters or the full name of the country.`,
				Transform:   transform.FromField("LegalAddress.Country").NullIfZero(),
			},
			{
				Name:        "legal_address_latitude",
				Type:        proto.ColumnType_STRING,
				Description: `Legal Address given to the government for any government communication. Latitude coordinate of Geocode (Geospacial Entity Object Code). INVALIDis returned for invalid addresses.`,
				Transform:   transform.FromField("LegalAddress.Latitude").NullIfZero(),
			},
			{
				Name:        "legal_address_longitude",
				Type:        proto.ColumnType_STRING,
				Description: `Legal Address given to the government for any government communication. Longitude coordinate of Geocode (Geospacial Entity Object Code). INVALIDis returned for invalid addresses.`,
				Transform:   transform.FromField("LegalAddress.Longitude").NullIfZero(),
			},
			{
				Name: "communication_address_line1",
				Type: proto.ColumnType_STRING,
				Description: `Address of the company as given to their customer, sometimes the address given to the customer mail address is different from Company address. 
				If a physical address is updated from within the transaction object, the QuickBooks Online API flows individual address components differently into the Line elements of the transaction response then when the transaction was first created:
				Line1 and Line2 elements are populated with the customer name and company name.
				Original Line1 through Line 5 contents, City, SubDivisionCode, and PostalCode flow into Line3 through Line5as a free format strings.`,
				Transform: transform.FromField("CustomerCommunicationAddress.Line1").NullIfZero(),
			},
			{
				Name: "communication_address_line2",
				Type: proto.ColumnType_STRING,
				Description: `Address of the company as given to their customer, sometimes the address given to the customer mail address is different from Company address. 				
				If a physical address is updated from within the transaction object, the QuickBooks Online API flows individual address components differently into the Line elements of the transaction response then when the transaction was first created:
				Line1 and Line2 elements are populated with the customer name and company name.
				Original Line1 through Line 5 contents, City, SubDivisionCode, and PostalCode flow into Line3 through Line5as a free format strings.`,
				Transform: transform.FromField("CustomerCommunicationAddress.Line2").NullIfZero(),
			},
			{
				Name: "communication_address_line3",
				Type: proto.ColumnType_STRING,
				Description: `Address of the company as given to their customer, sometimes the address given to the customer mail address is different from Company address. 
				If a physical address is updated from within the transaction object, the QuickBooks Online API flows individual address components differently into the Line elements of the transaction response then when the transaction was first created:
				Line1 and Line2 elements are populated with the customer name and company name.
				Original Line1 through Line 5 contents, City, SubDivisionCode, and PostalCode flow into Line3 through Line5as a free format strings.`,
				Transform: transform.FromField("CustomerCommunicationAddress.Line3").NullIfZero(),
			},
			{
				Name: "communication_address_line4",
				Type: proto.ColumnType_STRING,
				Description: `Address of the company as given to their customer, sometimes the address given to the customer mail address is different from Company address. 
				If a physical address is updated from within the transaction object, the QuickBooks Online API flows individual address components differently into the Line elements of the transaction response then when the transaction was first created:
				Line1 and Line2 elements are populated with the customer name and company name.
				Original Line1 through Line 5 contents, City, SubDivisionCode, and PostalCode flow into Line3 through Line5as a free format strings.`,
				Transform: transform.FromField("CustomerCommunicationAddress.Line4").NullIfZero(),
			},
			{
				Name: "communication_address_line5",
				Type: proto.ColumnType_STRING,
				Description: `Address of the company as given to their customer, sometimes the address given to the customer mail address is different from Company address. 
				If a physical address is updated from within the transaction object, the QuickBooks Online API flows individual address components differently into the Line elements of the transaction response then when the transaction was first created:
				Line1 and Line2 elements are populated with the customer name and company name.
				Original Line1 through Line 5 contents, City, SubDivisionCode, and PostalCode flow into Line3 through Line5as a free format strings.`,
				Transform: transform.FromField("CustomerCommunicationAddress.Line5").NullIfZero(),
			},
			{
				Name:        "communication_address_city",
				Type:        proto.ColumnType_STRING,
				Description: `Address of the company as given to their customer, sometimes the address given to the customer mail address is different from Company address. City name.`,
				Transform:   transform.FromField("CustomerCommunicationAddress.City").NullIfZero(),
			},
			{
				Name:        "communication_address_region",
				Type:        proto.ColumnType_STRING,
				Description: `Address of the company as given to their customer, sometimes the address given to the customer mail address is different from Company address. Region within a country. For example, state name for USA, province name for Canada.`,
				Transform:   transform.FromField("CustomerCommunicationAddress.Region").NullIfZero(),
			},
			{
				Name:        "communication_address_country",
				Type:        proto.ColumnType_STRING,
				Description: `Address of the company as given to their customer, sometimes the address given to the customer mail address is different from Company address. Country name. For international addresses - countries should be passed as 3 ISO alpha-3 characters or the full name of the country.`,
				Transform:   transform.FromField("CustomerCommunicationAddress.Country").NullIfZero(),
			},
			{
				Name:        "communication_address_latitude",
				Type:        proto.ColumnType_STRING,
				Description: `Address of the company as given to their customer, sometimes the address given to the customer mail address is different from Company address. Latitude coordinate of Geocode (Geospacial Entity Object Code). INVALIDis returned for invalid addresses.`,
				Transform:   transform.FromField("CustomerCommunicationAddress.Latitude").NullIfZero(),
			},
			{
				Name:        "communication_address_longitude",
				Type:        proto.ColumnType_STRING,
				Description: `Address of the company as given to their customer, sometimes the address given to the customer mail address is different from Company address. Longitude coordinate of Geocode (Geospacial Entity Object Code). INVALIDis returned for invalid addresses.`,
				Transform:   transform.FromField("CustomerCommunicationAddress.Longitude").NullIfZero(),
			},
			{
				Name:        "supported_languages",
				Type:        proto.ColumnType_STRING,
				Description: `Comma separated list of languages.`,
				Transform:   transform.FromField("SupportedLanguages").NullIfZero(),
			},
			{
				Name:        "country",
				Type:        proto.ColumnType_STRING,
				Description: `Country name to which the company belongs for financial calculations.`,
				Transform:   transform.FromField("Country").NullIfZero(),
			},
			{
				Name:        "email",
				Type:        proto.ColumnType_STRING,
				Description: `Default email address.`,
				Transform:   transform.FromField("Email.Address").NullIfZero(),
			},
			{
				Name:        "web",
				Type:        proto.ColumnType_STRING,
				Description: `Web site address.`,
				Transform:   transform.FromField("Web.URI").NullIfZero(),
			},
			// TODO: Maybe flatten this out?
			{
				Name: "attributes",
				Type: proto.ColumnType_JSON,
				Description: `Any other preference not covered with the standard set of attributes. See Data Services Extensions, below, for special reserved name/value pairs. NameValue.Name--Name of the element. NameValue.Value--Value of the element.
				NeoEnabled
				The type of company, classic or Harmony.
				NameValue.Name="NeoEnabled"
				NameValue.Value="neoFlag"
				where neoFlag is defined as:
				true for Harmony company
				false for Classic company

				firstTxnDate
				The date of the first transaction for the company.
				NameValue.Name="firsttxndate"
				NameValue.Value="date"
				where date is of the format yyyy-mm-dd
				This extension is avaliable when the include=firsttxndatequery parameter is include in the endpoint URI:
				A GET request looks like the following
				baseURL/company
				/213316401/companyinfo/213316401?
				include=firsttxndate
				A Query Request looks like the following
				baseURL/company/213316401/query?query=select * from CompanyInfo include=firsttxndate
				
				IndustryType
				The industry type for the company. This is defined when the company is first created.
				
				IndustryCode
				The NAICS/SIC industry code for the company. This is defined when the company is first created.
				
				CompanyType
				The company type as defined when the company is first created. Possible values include:
				Sole Proprietor
				Partnership
				Limited Liability
				Corporation
				Organization
				
				OfferingSKU
				The specific QuickBooks Online product. Possible values include:
				QuickBooks Online Plus
				QuickBooks Online Simple Start
				QuickBooks Online Essentials
				When CompanyInfo endpoint is invoked with minorversion=29, possible values include:
				QuickBooks Online Advanced - Advanced companies will return QuickBooks Online Plus as OfferingSKU if minor version lower than 29 is used.
				SubscriptionStatus
				minorVersion: 3
				The QuickBooks subscription status.
				Possible values, prior to minor version 3:
				TRIAL-Company is in trial
				PAID-For any other state
				When CompanyInfo endpoint is invoked with minorversion=3, possible values include:
				TRIAL-Company is in trial.
				SUBSCRIBED-Company is subscribed.
				TRIALOPTIN-Company is in trial and user has provided credit card info.
				RESTRICTED-The customer's subscription payment failed and QuickBooks services is waiting for the customer to update their payment information. During this state, customers have read and write access to their company file. If the customer does not update the payment information within a week, the state moves to suspended and write access is revoked.
				SUSPENDED-Company in a lock-out mode, for instance due to payment failure.
				EXPIRED-Company in a lock-out mode due to missing payment information.
				CANCELLED-Company is cancelled by the user or support agent.
				UNKNOWN-Context of the company is not available.

				PayrollFeature
				minorVersion: 3
				Whether subscription is enabled for the payroll feature.
				true is Enabled.
				false is Disabled.

				AccountantFeature
				Whether subscription is enabled for the accountant feature.
				true is Enabled.
				false is Disabled.

				ItemCategoriesFeature
				Whether a company is category enabled. Currently available for sandbox companies, only. This functionality will be rolled out to all companies in the coming months.
				true is Enabled.
				false is Disabled.

				NonTracking
				minorVersion: 28
				Property to determine whether the company is 'NonTracking' enabled. Based on this flag, the appropriate fields should be used while querying General Ledger or Profilt and Loss Detail report.
				true is Enabled.
				false is Disabled.`,
				Transform: transform.FromField("Attributes").NullIfZero(),
			},
			{
				Name:        "primary_phone",
				Type:        proto.ColumnType_STRING,
				Description: `Primary phone number.`,
				Transform:   transform.FromField("PrimaryPhone").NullIfZero(),
			},
			{
				Name:        "legal_name",
				Type:        proto.ColumnType_STRING,
				Description: `The legal name of the company`,
				Transform:   transform.FromField("LegalName").NullIfZero(),
			},
			{
				Name:        "create_time",
				Type:        proto.ColumnType_STRING,
				Description: `Time the entity was created in the source domain.`,
				Transform:   transform.FromField("MetaData.CreateTime").NullIfZero(),
			},
			{
				Name:        "last_updated_time",
				Type:        proto.ColumnType_STRING,
				Description: `Time the entity was last updated in the source domain.`,
				Transform:   transform.FromField("MetaData.LastUpdatedTime").NullIfZero(),
			},
			{
				Name:        "company_start_date",
				Type:        proto.ColumnType_STRING,
				Description: `DateTime when company file was created. This field and Metadata.CreateTimecontain the same value.`,
				Transform:   transform.FromField("CompanyStartDate").NullIfZero(),
			},
		},
	}
}

func listQBOCompanyInfo(
	ctx context.Context,
	d *plugin.QueryData,
	_ *plugin.HydrateData,
) (interface{}, error) {

	companyInfoApi := new(ApiCompanyInfo)
	_, err := qboApiCall(&companyInfoApi,
		"{{.baseURL}}/v3/company/{{.realmId}}/companyinfo/{{.realmId}}",
		ctx,
		d,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get company info: %v", err)
	}

	plugin.Logger(ctx).Info("Company Info: ", companyInfoApi.GetResponse())
	d.StreamListItem(ctx, *companyInfoApi.GetResponse())
	return nil, nil
}
