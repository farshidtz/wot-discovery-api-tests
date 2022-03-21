package directory

const (
	MediaTypeJSON             = "application/json"
	MediaTypeJSONLD           = "application/ld+json"
	MediaTypeThingDescription = "application/td+json"
	MediaTypeMergePatch       = "application/merge-patch+json"
)

// All TDD assertions
// Updated on 2020.05.21
var tddAssertions = []string{
	"tdd-anonymous-td-identifier",
	"tdd-anonymous-td-local-uuid",
	"tdd-http-error-response",
	"tdd-http-head",
	"tdd-http-missing-api-endpoint",
	"tdd-https",
	"tdd-notification",
	"tdd-notification-data",
	"tdd-notification-data-create-full",
	"tdd-notification-data-delete-diff",
	"tdd-notification-data-diff-unsupported",
	"tdd-notification-data-td-id",
	"tdd-notification-data-update-diff",
	"tdd-notification-data-update-id",
	"tdd-notification-event-id",
	"tdd-notification-event-types",
	"tdd-notification-filter-type",
	"tdd-notification-sse",
	"tdd-registrationinfo-expiry-config",
	"tdd-registrationinfo-expiry-purge",
	"tdd-registrationinfo-vocab-created",
	"tdd-registrationinfo-vocab-expires",
	"tdd-registrationinfo-vocab-modified",
	"tdd-registrationinfo-vocab-retrieved",
	"tdd-registrationinfo-vocab-ttl",
	"tdd-search-large-tdds",
	"tdd-search-sparql",
	"tdd-search-sparql-error",
	"tdd-search-sparql-federation",
	"tdd-search-sparql-federation-version",
	"tdd-search-sparql-method-get",
	"tdd-search-sparql-method-post",
	"tdd-search-sparql-resp-describe-construct",
	"tdd-search-sparql-resp-select-ask",
	"tdd-search-sparql-version",
	"tdd-things-additional-representation",
	"tdd-things-create-anonymous-contenttype",
	"tdd-things-create-anonymous-td",
	"tdd-things-create-anonymous-td-resp",
	"tdd-things-create-known-contenttype",
	"tdd-things-create-known-td",
	"tdd-things-create-known-td-resp",
	"tdd-things-create-known-vs-anonymous",
	"tdd-things-crud",
	"tdd-things-crudl",
	"tdd-things-default-representation",
	"tdd-things-delete",
	"tdd-things-delete-resp",
	"tdd-things-list-method",
	"tdd-things-list-only",
	"tdd-things-list-pagination",
	"tdd-things-list-pagination-collection",
	"tdd-things-list-pagination-header-canonicallink",
	"tdd-things-list-pagination-header-nextlink",
	"tdd-things-list-pagination-header-nextlink-attr",
	"tdd-things-list-pagination-header-nextlink-base",
	"tdd-things-list-pagination-limit",
	"tdd-things-list-pagination-order",
	"tdd-things-list-pagination-order-default",
	"tdd-things-list-pagination-order-nextlink",
	"tdd-things-list-pagination-order-unsupported",
	"tdd-things-list-resp",
	"tdd-things-read-only-auth",
	"tdd-things-retrieve",
	"tdd-things-retrieve-resp",
	"tdd-things-update",
	"tdd-things-update-contenttype",
	"tdd-things-update-partial",
	"tdd-things-update-partial-contenttype",
	"tdd-things-update-partial-mergepatch",
	"tdd-things-update-partial-partialtd",
	"tdd-things-update-partial-resp",
	"tdd-things-update-resp",
	"tdd-validation-jsonschema",
	"tdd-validation-response",
	"tdd-validation-result",
	"tdd-validation-syntactic",
}