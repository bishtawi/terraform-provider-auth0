## v0.2.0-8 (September 25, 2019)

BUG FIXES:

* resource/auth0_tenant: Fix `flags`

## v0.2.0-7 (September 24, 2019)

BUG FIXES:

* resource/auth0_email_template: Fix optionality of fields

## v0.2.0-6 (September 24, 2019)

BUG FIXES:

* resource/auth0_tenant: Fix boolean parsing in `flags`

## v0.2.0-5 (September 24, 2019)

BUG FIXES:

* resource/auth0_tenant: Remove read-only attributes from `flags`

## v0.2.0-4 (September 24, 2019)

BUG FIXES:

* resource/auth0_connection: Fix `password_dictionary`
* resource/auth0_tenant: Fix `flags`

## v0.2.0-3 (September 23, 2019)

ENHANCEMENTS:

* resource/auth0_tenant: Add `flags`, `universal_login` and `idle_session_lifetime`

## v0.2.0-2 (September 18, 2019)

BUG FIXES:

* resource/auth0_user: Fix `user_metadata` and `app_metadata`
* resource/auth0_user: Make `password` sensitive

NOTES:

* Add example of auth0_tenant resource

## v0.2.0-1 (September 17, 2019)

ENHANCEMENTS:

* resource/auth0_connection: Add `password_complexity_options`
* resource/auth0_resource_server: Add `enforce_policies`
* resource/auth0_resource_server: Add `token_dialect`

BUG FIXES:

* resource/auth0_connection: Fix `password_dictionary`

NOTES:

* Update `go` to v1.13
* Update `go-auth0` to v1.2.5-1

## v0.2.0 (June 27, 2019)

ENHANCEMENTS:

* resource/auth0_user: Add support for user attribute `nickname`

BUG FIXES:

* resource/auth0_connection: Fix icorrect schema of `password_no_personal_info`

NOTES:

* Provider is compatible with Terraform v0.12.3

## v0.1.20 (May 17, 2019)

FEATURES:

* resource/auth0_connection: Add twillio for guardian mfa

## v0.1.19 (May 14, 2019)

ENHANCEMENTS

* resource/auth0_connection: Add `adfs_server` field.

## v0.1.18 (March 26, 2019)

NOTES:

* Extract version from changelog for release notes.

## v0.1.17 (March 26, 2019)

NOTES:

* Update Travis to build on tag push.

## v0.1.16 (March 25, 2019)

NOTES:

* Added contributing, code of conduct, issue templates to the project.

## v0.1.15 (March 25, 2019)

FEATURES:

* **New Resource:** auth0_tenant ([#79](https://github.com/yieldr/terraform-provider-auth0/pull/79))

ENHANCEMENTS:

* resource/auth0_connection: `enabled_clients` will suppress diff if the list order is different.
* resource/auth0_connection: set `client_secret` to sensitive ([#91](https://github.com/yieldr/terraform-provider-auth0/pull/91)).
* resource/auth0_resource_server: introduce `token_lifetime_for_web` ([#84](https://github.com/yieldr/terraform-provider-auth0/pull/84)).

NOTES:

* Re-imported Auth0 SDK from `gopkg.in/auth0.v1`.
