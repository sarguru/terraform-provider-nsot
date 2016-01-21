This is an experimental terraform provider for nsot. Currently it just supports site.

## Setup

Export the following variables

	export NSOT_EMAIL="<nsot_email>"
	export NSOT_SECRET="<nsot_secret>"
	export NSOT_URL="<nsot_url>"

example of nsot_url is "http://<NSOT_HOST>:<NSOT_PORT>/api"


## Example
```
resource "nsot_site" "test" {
        name = "test.me2"
        description = "bar"
}
```


