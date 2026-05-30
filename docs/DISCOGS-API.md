# Discogs API Documentation

## Home

Here’s your place to code all things Discogs! The Discogs API lets developers build their own Discogs-powered applications for the web, desktop, and mobile devices. We hope the API will connect and empower a community of music lovers around the world!

The Discogs API v2.0 is a RESTful interface to Discogs data. You can access JSON-formatted information about Database objects such as Artists, Releases, and Labels. Your application can also manage User Collections and Wantlists, create Marketplace Listings, and more.

Some Discogs data is made available under the CC0 No Rights Reserved license, and some is restricted data, as defined in our API Terms of Use.

Our monthly data dumps are available under the the CC0 No Rights Reserved license.

If you utilize the Discogs API, you are subject to the API Terms of Use. Please also ensure that any application you develop follows the Discogs Application Name and Description Policy.

### Quickstart

If you just want to see some results right now, issue this curl command:

```bash
curl https://api.discogs.com/releases/249504 --user-agent "FooBarApp/3.0"
```

For community-maintained client libraries and example code, see the links below:

| Language | Type | Maintainer | URL |
| ---------- | ------ | ------------ | ----- |
| Node.js | Client | bartve | [https://github.com/bartve/disconnect](https://github.com/bartve/disconnect) |
| PHP | Client | ricbra | [https://github.com/ricbra/php-discogs-api](https://github.com/ricbra/php-discogs-api) |
| Python | Client | joalla | [https://github.com/joalla/discogs_client](https://github.com/joalla/discogs_client) |
| Python | Example | jesseward | [https://github.com/jesseward/discogs-oauth-example](https://github.com/jesseward/discogs-oauth-example) |
| Ruby | Client | buntine | [https://github.com/buntine/discogs](https://github.com/buntine/discogs) |

### General Information

Your application must provide a User-Agent string that identifies itself – preferably something that follows RFC 1945. Some good examples include:

```bash
AwesomeDiscogsBrowser/0.1 +http://adb.example.com
LibraryMetadataEnhancer/0.3 +http://example.com/lime
MyDiscogsClient/1.0 +http://mydiscogsclient.org
```

Please don’t just copy one of those! Make it unique so we can let you know if your application starts to misbehave – the alternative is that we just silently block it, which will confuse and infuriate your users.

Here are some bad examples that are unclear or obscure the nature of the application:

```bash
curl/7.9.8 (i686-pc-linux-gnu) libcurl 7.9.8 (OpenSSL 0.9.6b)
Mozilla/5.0 (X11; Linux i686; rv:6.0.2) Gecko/20100101 Firefox/6.0.2
my app
```

### JSONP

When a callback query string parameter is supplied, the API can return responses in JSONP format.

JSONP, by its nature, cannot access HTTP information like headers or the status code, so the API supplies that information in the response, like so:

```bash
GET https://api.discogs.com/artists/1?callback=callbackname
```

```bash
200 OK
Content-Type: text/javascript
```

```json
callbackname({
    "meta": {
        "status": 200,
    },
    "data": {
        "id": 1,
        "name": "Persuader, The"
        // ... and so on
    }
})
```

### Rate Limiting

**Requests are throttled by the server by source IP to 60 per minute for authenticated requests, and 25 per minute for unauthenticated requests, with some exceptions.**

Your application should identify itself to our servers via a unique user agent string in order to achieve the maximum number of requests per minute.

Our rate limiting tracks your requests using a moving average over a 60 second window. If no requests are made in 60 seconds, your window will reset.

We attach the following headers to responses to help you track your rate limit use:

`X-Discogs-Ratelimit`: The total number of requests you can make in a one minute window.

`X-Discogs-Ratelimit-Used` : The number of requests you’ve made in your existing rate limit window.

`X-Discogs-Ratelimit-Remaining`: The number of remaining requests you are able to make in the existing rate limit window.

Your application should take our global limit into account and throttle its requests locally.

In the future, we may update these rate limits at any time in order to provide service for all users.

### Pagination

Some resources represent collections of objects and may be paginated. By default, 50 items per page are shown.

To browse different pages, or change the number of items per page (up to 100), use the page and per_page query string parameters:

```bash
GET https://api.discogs.com/artists/1/releases?page=2&per_page=75
```

Responses include a Link header:

```bash
Link: <https://api.discogs.com/artists/1/releases?page=3&per_page=75>; rel=next,
<https://api.discogs.com/artists/1/releases?page=1&per_page=75>; rel=first,
<https://api.discogs.com/artists/1/releases?page=30&per_page=75>; rel=last,
<https://api.discogs.com/artists/1/releases?page=1&per_page=75>; rel=prev
```

And a pagination object in the response body:

```json
{
    "pagination": {
        "page": 2,
        "pages": 30,
        "items": 2255,
        "per_page": 75,
        "urls":
            {
                "first": "https://api.discogs.com/artists/1/releases?page=1&per_page=75",
                "prev": "https://api.discogs.com/artists/1/releases?page=1&per_page=75",
                "next": "https://api.discogs.com/artists/1/releases?page=3&per_page=75",
                "last": "https://api.discogs.com/artists/1/releases?page=30&per_page=75"
            }
    },
    "releases":
        [ ... ]
}
```

### Versioning and Media Types

Currently, our API only supports one version: v2. However, you can specify a version in your requests to future-proof your application. By adding an `Accept` header with the version and media type, you can guarantee your requests will receive data from the correct version you develop your app on.

A standard `Accept` header may look like this:
`application/vnd.discogs.v2.html+json`

If you are requesting information from an endpoint that may have text formatting in it, you can choose which kind of formatting you want to be returned by changing that section of the `Accept` header. We currently support 3 types:
`application/vnd.discogs.v2.html+json`
`application/vnd.discogs.v2.plaintext+json`
`application/vnd.discogs.v2.discogs+json`

If no Accept header is supplied, or if the Accept header differs from one of the three previous options, we default to `application/vnd.discogs.v2.discogs+json`.

### FAQ

#### 1. Why am I getting an empty response from the server?

This generally happens when you forget to add a User-Agent header to your requests.

#### 2. How do I get updates about the API?

Subscribe to our API Announcements forum thread. For larger, breaking changes, we will send out an email notice to all developers with a registered Discogs application.

#### 3. Where can I register a Discogs application?

You can register a Discogs application on the Developer Settings.

#### 4. If I have a question/issue with the API, should I file a Support Request?

It's generally best to start out with a forum post on the API topic since other developers may have had similar issues and can point you in the right direction. If the issue requires privacy, then a support request is the best way to go.

#### 5. I'm getting a 404 response when trying to fetch images; what gives?

This may seem obvious, but make sure you aren't doing anything to change the URL. The URLs returned are signed URLs, so trying to change one part of the URL (e.g., Release ID number) will generally not work.

#### 6. What are the authentication requirements for requesting images?

Please see the Images documentation page.

#### 7. Why am I getting a particular HTTP response?

##### 200 OK

The request was successful, and the requested data is provided in the response body.

##### 201 Continue

You’ve sent a POST request to a list of resources to create a new one. The ID of the newly-created resource will be provided in the body of the response.

##### 204 No Content

The request was successful, and the server has no additional information to convey, so the response body is empty.

##### 401 Unauthorized

You’re attempting to access a resource that first requires authentication. See Authenticating with OAuth.

##### 403 Forbidden

You’re not allowed to access this resource. Even if you authenticated, or already have, you simply don’t have permission. Trying to modify another user’s profile, for example, will produce this error.

##### 404 Not Found

The resource you requested doesn’t exist.

##### 405 Method Not Allowed

You’re trying to use an HTTP verb that isn’t supported by the resource. Trying to PUT to /artists/1, for example, will fail because Artists are read-only.

##### 422 Unprocessable Entity

Your request was well-formed, but there’s something semantically wrong with the body of the request. This can be due to malformed JSON, a parameter that’s missing or the wrong type, or trying to perform an action that doesn’t make any sense. Check the response body for specific information about what went wrong.

##### 500 Internal Server Error

Something went wrong on our end while attempting to process your request. The response body’s message field will contain an error code that you can send to Discogs Support (which will help us track down your specific issue).

## Authentication

This section describes the various methods of authenticating with the Discogs API.

### Registration

In order to access protected endpoints, you’ll need to register for either a consumer key and secret or user token, depending on your situation:

- To easily access your own user account information, use a **User token**.
- To get access to an endpoint that requires authentication and build 3rd party apps, use a **Consumer Key and Secret**.

To get one of these, sign up for a Discogs account if you don’t have one already, and go to your Developer Settings. From there, you can create a new application/token or edit the metadata for an existing app.

When you create a new application, you’ll be granted a **Consumer Key** and **Consumer Secret**, which you can plug into your application and start making authenticated requests. It’s important that you don’t disclose the Consumer Secret to anyone.

Generating a user token is as easy as clicking the **Generate token** button on the developer settings. Keep this token private, as it allows users to access your information.

### OAuth Endpoints

The OAuth 1.0a flow involves three server-side endpoints:

Request token URL: `https://api.discogs.com/oauth/request_token`
Authorize URL: `https://www.discogs.com/oauth/authorize`
Access token URL: `https://api.discogs.com/oauth/access_token`

For convenience, these are also listed on the Edit Application page.

Once authenticated, you can test that everything’s working correctly by requesting the Identity resource.

### Discogs Auth Flow

If you do not plan on building an app which others can log into on their behalf, you should use this authentication method as it is much simpler and is still secure. Discogs Auth requires that requests be made over HTTPS, as you are sending your app’s key/secret or token in the request.
To send requests with Discogs Auth, you have two options: sending your credentials in the query string with `key` and `secret` parameters or a `token` parameter, for example:

```shell
curl "https://api.discogs.com/database/search?q=Nirvana&key=foo123&secret=bar456"
```

or

```shell
curl "https://api.discogs.com/database/search?q=Nirvana&token=abcxyz123456"
```

Your other option is to send your credentials in the request as an Authorization header, as follows:

```shell
curl "https://api.discogs.com/database/search?q=Nirvana" -H "Authorization: Discogs key=foo123, secret=bar456"
```

or

```shell
curl "https://api.discogs.com/database/search?q=Nirvana" -H "Authorization: Discogs token=abcxyz123456"
```

What differentiates the `key`/`secret` pair option from the `token` option? Let’s look at this table:

|Credentials in request|Rate limiting?|Image URLs?|Authenticated as user?|
|---|---|---|---|
|None|🐢 Low tier|❌ No|❌ No|
|Only Consumer key/secret|🐰 High tier|✔️ Yes|❌ No|
|Full OAuth 1.0a with access token/secret (see below)|🐰 High tier|✔️ Yes|✔️ Yes, on behalf of any user 🌍|
|Personal access token|🐰 High tier|✔️ Yes|✔️ Yes, for token holder only 👩|

In other words:

- Using the `key` and `secret` will get you image URLs. These are unavailable to unauthenticated requests.
- Using the `key` and `secret` will up your rate limit, unlike unauthenticated requests.
- **BUT** using the `key` and `secret` does **not** identify the requester as any particular user, and as such will not grant access to any resources users should be able to see on their own account (e.g. marketplace orders, private inventory fields, private collections). You will need to use either of the `token` options for these resources.

That’s it! Continue sending the key/secret pair or user token with the rest of your requests.

### OAuth Flow

OAuth is a protocol that allows users to grant access to their data without having to share their password.

This is an explanation of how a web application may work with Discogs using OAuth 1.0a. We highly suggest you use an OAuth library/wrapper to simplify the process of authenticating.

#### 1. Obtain consumer key and consumer secret from Developer Settings

Application registration can be found here: [https://www.discogs.com/settings/developers](https://www.discogs.com/settings/developers)

You only need to register once per application you make. You should not share your consumer secret, as it acts as a sort of password for your requests.

#### 2. Send a GET request to the Discogs request token URL

```bash
GET https://api.discogs.com/oauth/request_token
```

Include the following headers with your request:

```bash
Content-Type: application/x-www-form-urlencoded
Authorization:
        OAuth oauth_consumer_key="your_consumer_key",
        oauth_nonce="random_string_or_timestamp",
        oauth_signature="your_consumer_secret&",
        oauth_signature_method="PLAINTEXT",
        oauth_timestamp="current_timestamp",
        oauth_callback="your_callback"
User-Agent: some_user_agent
```

**Note**: using an OAuth library instead of generating these requests manually will likely save you a headache if you are new to OAuth. Please refer to your OAuth library’s documentation if you choose to do so.

As the example shows, we suggest sending requests with HTTPS and the PLAINTEXT signature method over HMAC-SHA1 due to its simple yet secure nature. This involves setting your oauth_signature_method to “PLAINTEXT” and your oauth_signature to be your consumer secret followed by an ampersand (&).

If all goes well with the request, you should get an `HTTP 200 OK` response. If an invalid OAuth request is sent, you will receive an `HTTP 400 Bad Request` response.

Be sure to include a unique User-Agent in the request headers.

A successful request will return a response that contains the following content: OAuth request token (`oauth_token`), an OAuth request token secret (`oauth_token_secret`), and a callback confirmation (`oauth_callback_confirmed`). These will be used in the following steps.

#### 3. Redirect your user to the Discogs Authorize page

Authorization page:

```bash
https://discogs.com/oauth/authorize?oauth_token=<your_oauth_request_token>
```

This page will ask the user to authorize your app on behalf of their Discogs account. If they accept and authorize, they will receive a verifier key to use as verification. This key is used in the next phase of OAuth authentication.

If you added a callback URL to your Discogs application registration, the user will be redirected to that URL, and you can capture their verifier from the response. The verifier will be used to generate the access token in the next step. You can always edit your application settings to include a callback URL (i.e., you don’t need to re-create a new application).

#### 4. Send a POST request to the Discogs access token URL

```curl
POST https://api.discogs.com/oauth/access_token
```

Include the following headers in your request:

```curl
Content-Type: application/x-www-form-urlencoded
Authorization:
        OAuth oauth_consumer_key="your_consumer_key",
        oauth_nonce="random_string_or_timestamp",
        oauth_token="oauth_token_received_from_step_2"
        oauth_signature="your_consumer_secret&",
        oauth_signature_method="PLAINTEXT",
        oauth_timestamp="current_timestamp",
        oauth_verifier="users_verifier"
User-Agent: some_user_agent
```

If the OAuth access token is not created within 15 minutes of when you receive the OAuth request token, your OAuth request token and verifier will expire, and you will need to re-create them. If you try to POST to the access token URL with an expired verifier or your request is malformed, you will receive an HTTP 400 Bad Request response.

As the example shows, we suggest sending requests with HTTPS and the PLAINTEXT signature method over HMAC-SHA1 due to its simple yet secure nature. This involves setting your oauth_signature_method to “PLAINTEXT” and your oauth_signature to be your consumer secret followed by an ampersand (&).

Be sure to include a unique User-Agent in the header.

A successful request will return a response that contains an OAuth access token (`oauth_token`) and an OAuth access token secret (`oauth_token_secret`). These tokens do not expire (unless the user revokes access from your app), so you should store these tokens in a database or persistent storage to make future requests signed with OAuth. All requests that require OAuth will require these two tokens to be sent in the request.

#### 5. Send authenticated requests to Discogs endpoints

You are now ready to send authenticated requests with Discogs through OAuth. Be sure to attach the user’s OAuth access token and OAuth access token secret to each request.

To test that you are ready to send authenticated requests, send a GET request to the identity URL. A successful request will yield a response that contains information about the authenticated user.

```curl
GET https://api.discogs.com/oauth/identity
```

### OAuth Endpoints

#### Request Token URL

Generate the request token

```url
GET  /oauth/request_token
```

**Request**

Headers

```curl
Content-Type: application/x-www-form-urlencoded
Authorization: OAuth oauth_consumer_key="your_consumer_key", oauth_nonce="random_string_or_timestamp", oauth_signature="your_consumer_secret&", oauth_signature_method="PLAINTEXT", oauth_timestamp="current_timestamp", oauth_callback="your_callback"
User-Agent: some_user_agent
```

**Response  `200`**

Headers

```url
oauth_token: abc123
oauth_token_secret: xyz789
oauth_callback_confirmed: 'true'
```

#### Access Token URL

Generate the access token

```url
POST /oauth/access_token
```

**Request**

Headers

```curl
Content-Type: application/x-www-form-urlencoded
Authorization: OAuth oauth_consumer_key="your_consumer_key", oauth_nonce="random_string_or_timestamp", oauth_token="oauth_token_received_from_step_2" oauth_signature="your_consumer_secret&", oauth_signature_method="PLAINTEXT", oauth_timestamp="current_timestamp", oauth_verifier="users_verifier"
User-Agent: some_user_agent
```

**Response  `200`**

Headers

```curl
oauth_token: abc123
oauth_token_secret: xyz789
```

## Database

### Release

The Release resource represents a particular physical or digital object released by one or more Artists.

```url
GET /releases/{release_id}{?curr_abbr}
```

Get a release

**Parameters**

- **release_id**: `number` (required) Example: 249504. The Release ID
- **curr_abbr**: `string` (optional) Example: USD. Currency for marketplace data. Defaults to the authenticated users currency. Must be one of the following: `USD` `GBP` `EUR` `CAD` `AUD` `JPY` `CHF` `MXN` `BRL` `NZD` `SEK` `ZAR`

**Response  `200`**

Headers

```
Status: HTTP/1.1 200 OK
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Content-Type: application/json
Server: lighttpd
Content-Length: 6223
Date: Tue, 01 Jul 2014 00:59:34 GMT
X-Varnish: 1465474310
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```

Body

```json
{
    "title": "Never Gonna Give You Up",
    "id": 249504,
    "artists": [
        {
            "anv": "",
            "id": 72872,
            "join": "",
            "name": "Rick Astley",
            "resource_url": "https://api.discogs.com/artists/72872",
            "role": "",
            "tracks": ""
        }
    ],
    "data_quality": "Correct",
    "thumb": "https://api-img.discogs.com/kAXVhuZuh_uat5NNr50zMjN7lho=/fit-in/300x300/filters:strip_icc():format(jpeg):mode_rgb()/discogs-images/R-249504-1334592212.jpeg.jpg",
    "community": {
        "contributors": [
            {
                "resource_url": "https://api.discogs.com/users/memory",
                "username": "memory"
            },
            {
                "resource_url": "https://api.discogs.com/users/_80_",
                "username": "_80_"
            }
        ],
        "data_quality": "Correct",
        "have": 252,
        "rating": {
            "average": 3.42,
            "count": 45
        },
        "status": "Accepted",
        "submitter": {
            "resource_url": "https://api.discogs.com/users/memory",
            "username": "memory"
        },
        "want": 42
    },
    "companies": [
        {
            "catno": "",
            "entity_type": "13",
            "entity_type_name": "Phonographic Copyright (p)",
            "id": 82835,
            "name": "BMG Records (UK) Ltd.",
            "resource_url": "https://api.discogs.com/labels/82835"
        },
        {
            "catno": "",
            "entity_type": "29",
            "entity_type_name": "Mastered At",
            "id": 266218,
            "name": "Utopia Studios",
            "resource_url": "https://api.discogs.com/labels/266218"
        }
    ],
    "country": "UK",
    "date_added": "2004-04-30T08:10:05-07:00",
    "date_changed": "2012-12-03T02:50:12-07:00",
    "estimated_weight": 60,
    "extraartists": [
        {
            "anv": "Me Co",
            "id": 547352,
            "join": "",
            "name": "Me Company",
            "resource_url": "https://api.discogs.com/artists/547352",
            "role": "Design",
            "tracks": ""
        },
        {
            "anv": "Stock / Aitken / Waterman",
            "id": 20942,
            "join": "",
            "name": "Stock, Aitken & Waterman",
            "resource_url": "https://api.discogs.com/artists/20942",
            "role": "Producer, Written-By",
            "tracks": ""
        }
    ],
    "format_quantity": 1,
    "formats": [
        {
            "descriptions": [
                "7\"",
                "Single",
                "45 RPM"
            ],
            "name": "Vinyl",
            "qty": "1"
        }
    ],
    "genres": [
        "Electronic",
        "Pop"
    ],
    "identifiers": [
        {
            "type": "Barcode",
            "value": "5012394144777"
        },
    ],
    "images": [
        {
            "height": 600,
            "resource_url": "https://api-img.discogs.com/z_u8yqxvDcwVnR4tX2HLNLaQO2Y=/fit-in/600x600/filters:strip_icc():format(jpeg):mode_rgb():quality(96)/discogs-images/R-249504-1334592212.jpeg.jpg",
            "type": "primary",
            "uri": "https://api-img.discogs.com/z_u8yqxvDcwVnR4tX2HLNLaQO2Y=/fit-in/600x600/filters:strip_icc():format(jpeg):mode_rgb():quality(96)/discogs-images/R-249504-1334592212.jpeg.jpg",
            "uri150": "https://api-img.discogs.com/0ZYgPR4X2HdUKA_jkhPJF4SN5mM=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb()/discogs-images/R-249504-1334592212.jpeg.jpg",
            "width": 600
        },
        {
            "height": 600,
            "resource_url": "https://api-img.discogs.com/EnQXaDOs5T6YI9zq-R5I_mT7hSk=/fit-in/600x600/filters:strip_icc():format(jpeg):mode_rgb():quality(96)/discogs-images/R-249504-1334592228.jpeg.jpg",
            "type": "secondary",
            "uri": "https://api-img.discogs.com/EnQXaDOs5T6YI9zq-R5I_mT7hSk=/fit-in/600x600/filters:strip_icc():format(jpeg):mode_rgb():quality(96)/discogs-images/R-249504-1334592228.jpeg.jpg",
            "uri150": "https://api-img.discogs.com/abk0FWgWsRDjU4bkCDwk0gyMKBo=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb()/discogs-images/R-249504-1334592228.jpeg.jpg",
            "width": 600
        }
    ],
    "labels": [
        {
            "catno": "PB 41447",
            "entity_type": "1",
            "id": 895,
            "name": "RCA",
            "resource_url": "https://api.discogs.com/labels/895"
        }
    ],
    "lowest_price": 0.63,
    "master_id": 96559,
    "master_url": "https://api.discogs.com/masters/96559",
    "notes": "UK Release has a black label with the text \"Manufactured In England\" printed on it.\r\n\r\nSleeve:\r\n\u2117 1987 \u2022 BMG Records (UK) Ltd. \u00a9 1987 \u2022 BMG Records (UK) Ltd.\r\nDistributed in the UK by BMG Records \u2022  Distribu\u00e9 en Europe par BMG/Ariola \u2022 Vertrieb en Europa d\u00fcrch BMG/Ariola.\r\n\r\nCenter labels:\r\n\u2117 1987 Pete Waterman Ltd.\r\nOriginal Sound Recording made by PWL.\r\nBMG Records (UK) Ltd. are the exclusive licensees for the world.\r\n\r\nDurations do not appear on the release.\r\n",
    "num_for_sale": 58,
    "released": "1987",
    "released_formatted": "1987",
    "resource_url": "https://api.discogs.com/releases/249504",
    "series": [],
    "status": "Accepted",
    "styles": [
        "Synth-pop"
    ],
    "tracklist": [
        {
            "duration": "3:32",
            "position": "A",
            "title": "Never Gonna Give You Up",
            "type_": "track"
        },
        {
            "duration": "3:30",
            "position": "B",
            "title": "Never Gonna Give You Up (Instrumental)",
            "type_": "track"
        }
    ],
    "uri": "https://www.discogs.com/Rick-Astley-Never-Gonna-Give-You-Up/release/249504",
    "videos": [
        {
            "description": "Rick Astley - Never Gonna Give You Up (Extended Version)",
            "duration": 330,
            "embed": true,
            "title": "Rick Astley - Never Gonna Give You Up (Extended Version)",
            "uri": "https://www.youtube.com/watch?v=te2jJncBVG4"
        },
    ],
    "year": 1987
}
```

**Response  `404`**

Headers

```
Status: HTTP/1.1 404 Not Found
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Content-Type: application/json
Server: lighttpd
Content-Length: 33
Date: Tue, 01 Jul 2014 01:03:20 GMT
X-Varnish: 1465521729
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```

Body

```json
{
  "message": "Release not found."
}
```

### Release Rating By User

The Release Rating endpoint retrieves, updates, or deletes the rating of a release for a given user.

```url
GET /releases/{release_id}/rating/{username}
```

Retrieves the release’s rating for a given user.

**Parameters**

- **release_id**: number (required) Example: 249504. The Release ID
- **username**: string (required) Example: memory. The username of the rating you are trying to request.

**Response  `200`**

Headers

```
Status: HTTP/1.1 200 OK
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Content-Type: application/json
Server: lighttpd
Content-Length: 6223
Date: Tue, 01 Jul 2014 00:59:34 GMT
X-Varnish: 1465474310
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```

Body

```json
{
  "username": "memory",
  "release_id": 249504,
  "rating": 5
}
```

```
PUT /releases/{release_id}/rating/{username}
```

Updates the release’s rating for a given user. Authentication as the user is required.

**Parameters**

- **release_id**: number (required) Example: 249504.  The Release ID
- **username**: string (required) Example: memory. The username of the rating you are trying to request.
- **rating**: int (required) Example: 5. The new rating for a release between 1 and 5.

**Response  `200`**

Headers

```
Status: HTTP/1.1 200 OK
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Content-Type: application/json
Server: lighttpd
Content-Length: 6223
Date: Tue, 01 Jul 2014 00:59:34 GMT
X-Varnish: 1465474310
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```

Body

```json
{
  "username": "memory",
  "release_id": 249504,
  "rating": 5
}
```

```
DELETE /releases/{release_id}/rating/{username}
```

Deletes the release’s rating for a given user. Authentication as the user is required.

**Parameters**

- **release_id**: number (required) Example: 249504. The Release ID
- **username**: string (required) Example: memory. The username of the rating you are trying to request.

### Community Release Rating

The Community Release Rating endpoint retrieves the average rating and the total number of user ratings for a given release.

```
GET /releases/{release_id}/rating
```

Retrieves the community release rating average and count.

**Parameters**

- **release_id**: number (required) Example: 249504
The Release ID

****Response  200****

Headers

```
Status: HTTP/1.1 200 OK
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Content-Type: application/json
Server: lighttpd
Content-Length: 6223
Date: Tue, 01 Jul 2014 00:59:34 GMT
X-Varnish: 1465474310
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```

Body

```
{
    "rating": {
        "count": 47
        "average": 4.19
    },
    "release_id": 249504
}
```

### Release Stats

The Release Stats endpoint retrieves the total number of “haves” (in the community’s collections) and “wants” (in the community’s wantlists) for a given release.

```
GET /releases/{release_id}/stats
```

Retrieves the release’s “have” and “want” counts.

**Parameters**

- **release_id**: number (required) Example: 249504. The Release ID

****Response  200****

Headers

```
Status: HTTP/1.1 200 OK
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Content-Type: application/json
Server: lighttpd
Content-Length: 35
Date: Mon, 31 Aug 2020 23:22:16 GMT
X-Varnish: 1465474310
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```

Body

```
{
  "num_have": 2315,
  "num_want": 467
}
```

### Master Release

The Master resource represents a set of similar Releases. Masters (also known as “master releases”) have a “main release” which is often the chronologically earliest.

```
GET /masters/{master_id}
```

Get a master release

**Parameters**

- **master_id**: number (required) Example: 1000: The Master ID

****Response  200****

Headers

```
Status: HTTP/1.1 200 OK
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Content-Type: application/json
Server: lighttpd
Content-Length: 7083
Date: Tue, 01 Jul 2014 01:11:23 GMT
X-Varnish: 1465622695
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```

Body

```
{
  "styles": [
    "Goa Trance"
  ],
  "genres": [
    "Electronic"
  ],
  "videos": [
    {
      "duration": 421,
      "description": "Electric Universe - Alien Encounter Part 2 (Spirit Zone 97)",
      "embed": true,
      "uri": "https://www.youtube.com/watch?v=n1LGinzMDi8",
      "title": "Electric Universe - Alien Encounter Part 2 (Spirit Zone 97)"
    }
  ],
  "title": "Stardiver",
  "main_release": 66785,
  "main_release_url": "https://api.discogs.com/releases/66785",
  "uri": "https://www.discogs.com/Electric-Universe-Stardiver/master/1000",
  "artists": [
    {
      "join": "",
      "name": "Electric Universe",
      "anv": "",
      "tracks": "",
      "role": "",
      "resource_url": "https://api.discogs.com/artists/21849",
      "id": 21849
    }
  ],
  "versions_url": "https://api.discogs.com/masters/1000/versions",
  "year": 1997,
  "images": [
    {
      "height": 569,
      "resource_url": "https://api-img.discogs.com/_0K5t_iLs6CzLPKTB4mwHVI3Vy0=/fit-in/600x569/filters:strip_icc():format(jpeg):mode_rgb():quality(96)/discogs-images/R-66785-1213949871.jpeg.jpg",
      "type": "primary",
      "uri": "https://api-img.discogs.com/_0K5t_iLs6CzLPKTB4mwHVI3Vy0=/fit-in/600x569/filters:strip_icc():format(jpeg):mode_rgb():quality(96)/discogs-images/R-66785-1213949871.jpeg.jpg",
      "uri150": "https://api-img.discogs.com/sSWjXKczZseDjX2QohG1Lc77F-w=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb()/discogs-images/R-66785-1213949871.jpeg.jpg",
      "width": 600
    },
    {
      "height": 296,
      "resource_url": "https://api-img.discogs.com/1iD31iOWgfgb2DpROI4_MvmceFw=/fit-in/600x296/filters:strip_icc():format(jpeg):mode_rgb():quality(96)/discogs-images/R-66785-1213950065.jpeg.jpg",
      "type": "secondary",
      "uri": "https://api-img.discogs.com/1iD31iOWgfgb2DpROI4_MvmceFw=/fit-in/600x296/filters:strip_icc():format(jpeg):mode_rgb():quality(96)/discogs-images/R-66785-1213950065.jpeg.jpg",
      "uri150": "https://api-img.discogs.com/Cm4Q_1S784pQeRfwa0lN2jsj47Y=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb()/discogs-images/R-66785-1213950065.jpeg.jpg",
      "width": 600
    }
  ],
  "resource_url": "https://api.discogs.com/masters/1000",
  "tracklist": [
    {
      "duration": "7:00",
      "position": "1",
      "type_": "track",
      "title": "Alien Encounter (Part 2)"
    },
    {
      "duration": "7:13",
      "position": "2",
      "type_": "track",
      "extraartists": [
        {
          "join": "",
          "name": "DJ Sangeet",
          "anv": "",
          "tracks": "",
          "role": "Written-By, Producer",
          "resource_url": "https://api.discogs.com/artists/25460",
          "id": 25460
        }
      ],
      "title": "From The Heart"
    },
    {
      "duration": "6:45",
      "position": "3",
      "type_": "track",
      "title": "Radio S.P.A.C.E."
    }
  ],
  "id": 1000,
  "num_for_sale": 9,
  "lowest_price": 9.36,
  "data_quality": "Correct"
}
```

******Response  404******

Headers

```
Status: HTTP/1.1 404 Not Found
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Content-Type: application/json
Server: lighttpd
Content-Length: 40
Date: Tue, 01 Jul 2014 01:11:21 GMT
X-Varnish: 1465622316
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```

Body

```
{
  "message": "Master Release not found."
}
```

### Master Release Versions

```
GET /masters/{master_id}/versions{?page,per_page}
```

Retrieves a list of all Releases that are versions of this master. Accepts Pagination parameters.

**Parameters**

- **master_id**: number (required) Example: 1000. The Master ID

- **page**
number (optional) Example: 3
The page you want to request

- **per_page**
number (optional) Example: 25
The number of items per page

- **format**
string (optional) Example: Vinyl
The format to filter

- **label**
string (optional) Example: Scorpio Music
The label to filter

- **released**
string (optional) Example: 1992
The release year to filter

- **country**
string (optional) Example: Belgium
The country to filter

- **sort**
  - string (optional) Example: released
  - Sort items by this field:
  - released (i.e. year of the release)
  - title (i.e. title of the release)
  - format
  - label
  - catno (i.e. catalog number of the release)
  - country

**sort_order**
string (optional) Example: asc
Sort items in a particular order (one of asc, desc)

****Response  200****

Headers

```
Status: HTTP/1.1 200 OK
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Content-Type: application/json
Server: lighttpd
Content-Length: 2834
Date: Tue, 01 Jul 2014 01:16:01 GMT
X-Varnish: 1465678820
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```

Body

```
{
  "pagination": {
    "per_page": 50,
    "items": 4,
    "page": 1,
    "urls": {},
    "pages": 1
  },
  "versions": [
    {
      "status": "Accepted",
      "stats": {
        "user": {
          "in_collection": 0,
          "in_wantlist": 0
        },
        "community": {
          "in_collection": 1067,
          "in_wantlist": 765
        }
      },
      "thumb": "https://img.discogs.com/wV56xo0Ak0M2bTCC6B_heD7Dx_o=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/R-18926-1381225536-8716.jpeg.jpg",
      "format": "12\", 33 ⅓ RPM",
      "country": "US",
      "title": "Plastic Dreams",
      "label": "Epic",
      "released": "1993",
      "major_formats": [
        "Vinyl"
      ],
      "catno": "49 74992",
      "resource_url": "https://api.discogs.com/releases/18926",
      "id": 18926
    },
    {
      "status": "Accepted",
      "stats": {
        "user": {
          "in_collection": 0,
          "in_wantlist": 0
        },
        "community": {
          "in_collection": 842,
          "in_wantlist": 762
        }
      },
      "thumb": "https://img.discogs.com/KAm38-Op5VlkvuJOaTGZieKwRVg=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/R-34228-1215760312.jpeg.jpg",
      "format": "12\", 33 ⅓ RPM",
      "country": "UK",
      "title": "Plastic Dreams (Mixes)",
      "label": "R & S Records",
      "released": "1993",
      "major_formats": [
        "Vinyl"
      ],
      "catno": "RSGB 101T",
      "resource_url": "https://api.discogs.com/releases/34228",
      "id": 34228
    },
    {
      "status": "Accepted",
      "stats": {
        "user": {
          "in_collection": 0,
          "in_wantlist": 0
        },
        "community": {
          "in_collection": 20,
          "in_wantlist": 285
        }
      },
      "thumb": "https://img.discogs.com/CnovcvhAYIle0tYTTZHo42wPz88=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/R-1873038-1249267784.jpeg.jpg",
      "format": "12\", White Label, 33 ⅓ RPM, Promo",
      "country": "UK",
      "title": "Plastic Dreams Mixes",
      "label": "R & S Records",
      "released": "1993",
      "major_formats": [
        "Vinyl"
      ],
      "catno": "RSGB 101 T",
      "resource_url": "https://api.discogs.com/releases/1873038",
      "id": 1873038
    }
  ]
}
```

******Response  404******

Headers

```
Status: HTTP/1.1 404 Not Found
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Content-Type: application/json
Server: lighttpd
Content-Length: 40
Date: Tue, 01 Jul 2014 01:20:23 GMT
X-Varnish: 1465732620
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```

Body

```
{
  "message": "Master Release not found."
}
```

### Artist

The Artist resource represents a person in the Discogs database who contributed to a Release in some capacity.

```
GET /artists/{artist_id}
```

Get an artist

**Parameters**

- artist_id
number (required) Example: 108713
The Artist ID

****Response  200****

Headers

```
Status: HTTP/1.1 200 OK
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Content-Type: application/json
Server: lighttpd
Content-Length: 1258
Date: Tue, 01 Jul 2014 01:06:53 GMT
X-Varnish: 1465566651
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```

Body

```
{
  "namevariations": [
    "Nickleback"
  ],
  "profile": "Nickelback is a Canadian rock band from Hanna, Alberta formed in 1995. Nickelback's music is classed as hard rock and alternative metal. Nickelback is one of the most commercially successful Canadian groups, having sold almost 50 million albums worldwide, ranking as the 11th best selling music act of the 2000s, and is the 2nd best selling foreign act in the U.S. behind The Beatles for the 2000's.",
  "releases_url": "https://api.discogs.com/artists/108713/releases",
  "resource_url": "https://api.discogs.com/artists/108713",
  "uri": "https://www.discogs.com/artist/108713-Nickelback",
  "urls": [
    "http://www.nickelback.com/",
    "http://en.wikipedia.org/wiki/Nickelback"
  ],
  "data_quality": "Needs Vote",
  "id": 108713,
  "images": [
    {
      "height": 260,
      "resource_url": "https://api-img.discogs.com/9xJ5T7IBn23DDMpg1USsDJ7IGm4=/330x260/smart/filters:strip_icc():format(jpeg):mode_rgb():quality(96)/discogs-images/A-108713-1110576087.jpg.jpg",
      "type": "primary",
      "uri": "https://api-img.discogs.com/9xJ5T7IBn23DDMpg1USsDJ7IGm4=/330x260/smart/filters:strip_icc():format(jpeg):mode_rgb():quality(96)/discogs-images/A-108713-1110576087.jpg.jpg",
      "uri150": "https://api-img.discogs.com/--xqi8cBtaBZz3qOjVcvzGvNRmU=/150x150/smart/filters:strip_icc():format(jpeg):mode_rgb()/discogs-images/A-108713-1110576087.jpg.jpg",
      "width": 330
    },
    {
      "height": 500,
      "resource_url": "https://api-img.discogs.com/r1jRG8b9-nlqTHPlJ-t8JR5ugoA=/493x500/smart/filters:strip_icc():format(jpeg):mode_rgb():quality(96)/discogs-images/A-108713-1264273865.jpeg.jpg",
      "type": "secondary",
      "uri": "https://api-img.discogs.com/r1jRG8b9-nlqTHPlJ-t8JR5ugoA=/493x500/smart/filters:strip_icc():format(jpeg):mode_rgb():quality(96)/discogs-images/A-108713-1264273865.jpeg.jpg",
      "uri150": "https://api-img.discogs.com/6K-cI5xDgsurmc-2OX6uCygzDgw=/150x150/smart/filters:strip_icc():format(jpeg):mode_rgb()/discogs-images/A-108713-1264273865.jpeg.jpg",
      "width": 493
    }
  ],
  "members": [
    {
      "active": true,
      "id": 270222,
      "name": "Chad Kroeger",
      "resource_url": "https://api.discogs.com/artists/270222"
    },
    {
      "active": true,
      "id": 685755,
      "name": "Daniel Adair",
      "resource_url": "https://api.discogs.com/artists/685755"
    },
    {
      "active": true,
      "id": 685754,
      "name": "Mike Kroeger",
      "resource_url": "https://api.discogs.com/artists/685754"
    },
    {
      "active": true,
      "id": 685756,
      "name": "Ryan \"Vik\" Vikedal",
      "resource_url": "https://api.discogs.com/artists/685756"
    },
    {
      "active": true,
      "id": 685757,
      "name": "Ryan Peake",
      "resource_url": "https://api.discogs.com/artists/685757"
    }
  ],
}
```

******Response  404******

Headers

```
Status: HTTP/1.1 404 Not Found
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Content-Type: application/json
Server: lighttpd
Content-Length: 32
Date: Tue, 01 Jul 2014 01:08:31 GMT
X-Varnish: 1465587583
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```

Body

```
{
  "message": "Artist not found."
}
```

### Artist Releases

Returns a list of Releases and Masters associated with the Artist.
Accepts Pagination.

```
GET /artists/{artist_id}/releases{?sort,sort_order}
```

Get an artist’s releases

**Parameters**

- **artist_id**
number (required) Example: 108713
The Artist ID

- **sort**
  - string (optional) Example: year
  Sort items by this field:
  - year (i.e. year of the release)
  - title (i.e. title of the release)
  - format

- **sort_order**
string (optional) Example: asc
Sort items in a particular order (one of asc, desc)

****Response  200****

Headers

```
Status: HTTP/1.1 200 OK
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Content-Type: application/json
Server: lighttpd
Content-Length: 1258
Date: Tue, 01 Jul 2014 01:06:53 GMT
X-Varnish: 1465566651
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```

Body

```
{
  "pagination": {
    "per_page": 50,
    "items": 9,
    "page": 1,
    "urls": {},
    "pages": 1
  },
  "releases": [
    {
      "artist": "Nickelback",
      "id": 173765,
      "main_release": 3128432,
      "resource_url": "http://api.discogs.com/masters/173765",
      "role": "Main",
      "thumb": "https://api-img.discogs.com/lb0zp7--FLaRP0LmJ4W6DhfweNc=/fit-in/90x90/filters:strip_icc():format(jpeg):mode_rgb()/discogs-images/R-5557864-1396493975-7618.jpeg.jpg",
      "title": "Curb",
      "type": "master",
      "year": 1996
    },
    {
      "artist": "Nickelback",
      "format": "CD, EP",
      "id": 4299404,
      "label": "Not On Label (Nickelback Self-released)",
      "resource_url": "http://api.discogs.com/releases/4299404",
      "role": "Main",
      "status": "Accepted",
      "thumb": "https://api-img.discogs.com/eFRGD78N7UhtvRjhdLZYXs2QJXk=/fit-in/90x90/filters:strip_icc():format(jpeg):mode_rgb()/discogs-images/R-4299404-1361106117-3632.jpeg.jpg",
      "title": "Hesher",
      "type": "release",
      "year": 1996
    },
    {
      "artist": "Nickelback",
      "id": 173767,
      "main_release": 1905922,
      "resource_url": "http://api.discogs.com/masters/173767",
      "role": "Main",
      "thumb": "https://api-img.discogs.com/12LXbUV44IHjyb6drFZOTQxgCqs=/fit-in/90x90/filters:strip_icc():format(jpeg):mode_rgb()/discogs-images/R-1905922-1251540516.jpeg.jpg",
      "title": "Leader Of Men",
      "type": "master",
      "year": 1999
    }
  ]
}
```

****Response  404****

Headers

```
Status: HTTP/1.1 404 Not Found
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Content-Type: application/json
Server: lighttpd
Content-Length: 32
Date: Tue, 01 Jul 2014 01:08:31 GMT
X-Varnish: 1465587583
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```

Body

```
{
  "message": "Artist not found."
}
```

### Label

The Label resource represents a label, company, recording studio, location, or other entity involved with Artists and Releases. Labels were recently expanded in scope to include things that aren’t labels – the name is an artifact of this history.

```
GET /labels/{label_id}
```

Get a label

**Parameters**

- label_id
number (required) Example: 1
The Label ID

****Response  200****

Headers

```
Status: HTTP/1.1 200 OK
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Content-Type: application/json
Server: lighttpd
Content-Length: 2834
Date: Tue, 01 Jul 2014 01:16:01 GMT
X-Varnish: 1465678820
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```

Body

```
{
  "profile": "Classic Techno label from Detroit, USA.\r\n[b]Label owner:[/b] [a=Carl Craig].\r\n",
  "releases_url": "https://api.discogs.com/labels/1/releases",
  "name": "Planet E",
  "contact_info": "Planet E Communications\r\nP.O. Box 27218\r\nDetroit, 48227, USA\r\n\r\np: 313.874.8729 \r\nf: 313.874.8732\r\n\r\nemail: info AT Planet-e DOT net\r\n",
  "uri": "https://www.discogs.com/label/1-Planet-E",
  "sublabels": [
    {
      "resource_url": "https://api.discogs.com/labels/86537",
      "id": 86537,
      "name": "Antidote (4)"
    },
    {
      "resource_url": "https://api.discogs.com/labels/41841",
      "id": 41841,
      "name": "Community Projects"
    }
  ],
  "urls": [
    "http://www.planet-e.net",
    "http://planetecommunications.bandcamp.com",
    "http://twitter.com/planetedetroit"
  ],
  "images": [
    {
      "height": 24,
      "resource_url": "https://api-img.discogs.com/85-gKw4oEXfDp9iHtqtCF5Y_ZgI=/fit-in/132x24/filters:strip_icc():format(jpeg):mode_rgb():quality(96)/discogs-images/L-1-1111053865.png.jpg",
      "type": "primary",
      "uri": "https://api-img.discogs.com/85-gKw4oEXfDp9iHtqtCF5Y_ZgI=/fit-in/132x24/filters:strip_icc():format(jpeg):mode_rgb():quality(96)/discogs-images/L-1-1111053865.png.jpg",
      "uri150": "https://api-img.discogs.com/cYmCut4Yh99FaLFHyoqkFo-Md1E=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb()/discogs-images/L-1-1111053865.png.jpg",
      "width": 132
    }
  ],
  "resource_url": "https://api.discogs.com/labels/1",
  "id": 1,
  "data_quality": "Needs Vote"
}
```

****Response  404****

Headers

```
Status: HTTP/1.1 404 Not Found
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Content-Type: application/json
Server: lighttpd
Content-Length: 30
Date: Tue, 01 Jul 2014 01:17:27 GMT
X-Varnish: 1465696276
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```

Body

```
{
  "message": "Label not found."
}
```

### All Label Releases

```
GET /labels/{label_id}/releases{?page,per_page}
```

Returns a list of Releases associated with the label. Accepts Pagination parameters.

**Parameters**

- label_id
number (required) Example: 1
The Label ID

- page
number (optional) Example: 3
The page you want to request

- per_page
number (optional) Example: 25
The number of items per page

****Response  200****

Headers

```
Status: HTTP/1.1 200 OK
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Content-Type: application/json
Server: lighttpd
Content-Length: 2834
Date: Tue, 01 Jul 2014 01:16:01 GMT
X-Varnish: 1465678820
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```

Body

```
{
  "pagination": {
    "per_page": 5,
    "pages": 68,
    "page": 1,
    "urls": {
      "last": "https://api.discogs.com/labels/1/releases?per_page=5&page=68",
      "next": "https://api.discogs.com/labels/1/releases?per_page=5&page=2"
    },
    "items": 338
  },
  "releases": [
    {
      "artist": "Andrea Parker",
      "catno": "!K7071CD",
      "format": "CD, Mixed",
      "id": 2801,
      "resource_url": "http://api.discogs.com/releases/2801",
      "status": "Accepted",
      "thumb": "https://api-img.discogs.com/cYmCut4Yh99FaLFHyoqkFo-Md1E=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb()/discogs-images/L-1-1111053865.png.jpg",
      "title": "DJ-Kicks",
      "year": 1998
    },
    {
      "artist": "Naomi Daniel",
      "catno": "2INZ 00140",
      "format": "12\"",
      "id": 65040,
      "resource_url": "http://api.discogs.com/releases/65040",
      "status": "Accepted",
      "thumb": "https://api-img.discogs.com/cYmCut4Yh99FaLFHyoqkFo-Md1E=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb()/discogs-images/L-1-1111053865.png.jpg",
      "title": "Stars",
      "year": 1993
    },
    {
      "artist": "Innerzone Orchestra",
      "catno": "546 137-2",
      "format": "CD, Album, P/Mixed",
      "id": 9922,
      "resource_url": "http://api.discogs.com/releases/9922",
      "status": "Accepted",
      "thumb": "https://api-img.discogs.com/cYmCut4Yh99FaLFHyoqkFo-Md1E=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb()/discogs-images/L-1-1111053865.png.jpg",
      "title": "Programmed",
      "year": 1999
    }
  ]
}
```

****Response  404****

Headers

```
Status: HTTP/1.1 404 Not Found
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Content-Type: application/json
Server: lighttpd
Content-Length: 30
Date: Tue, 01 Jul 2014 01:17:27 GMT
X-Varnish: 1465696276
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```

Body

```
{
  "message": "Label not found."
}
```

### Search

Issue a search query to our database. This endpoint accepts pagination parameters

Authentication (as any user) is required.

```
GET /database/search?q={query}&{?type,title,release_title,credit,artist,anv,label,genre,style,country,year,format,catno,barcode,track,submitter,contributor}
```

Issue a search query

**Parameters**

- query
string (optional) Example: nirvana
Your search query

- type
string (optional) Example: release
String. One of release, master, artist, label

- title
string (optional) Example: nirvana - nevermind
Search by combined “Artist Name - Release Title” title field.

- release_title
string (optional) Example: nevermind
Search release titles.

- credit
string (optional) Example: kurt
Search release credits.

- artist
string (optional) Example: nirvana
Search artist names.

- anv
string (optional) Example: nirvana
Search artist ANV.

- label
string (optional) Example: dgc
Search label names.

- genre
string (optional) Example: rock
Search genres.

- style
string (optional) Example: grunge
Search styles.

- country
string (optional) Example: canada
Search release country.

- year
string (optional) Example: 1991
Search release year.

- format
string (optional) Example: album
Search formats.

- catno
string (optional) Example: DGCD-24425
Search catalog number.

- barcode
string (optional) Example: 7 2064-24425-2 4
Search barcodes.

- track
string (optional) Example: smells like teen spirit
Search track titles.

- submitter
string (optional) Example: milKt
Search submitter username.

- contributor
string (optional) Example: jerome99
Search contributor usernames.

**Request**

Body

```
GET https://api.discogs.com/database/search?release_title=nevermind&artist=nirvana&per_page=3&page=1
```

****Response  200****

Headers

```
Status: HTTP/1.1 200 OK
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Cache-Control: public, must-revalidate
Content-Type: application/json
Content-Encoding: gzip
Server: lighttpd
Content-Length: 623
Date: Tue, 15 Jul 2014 18:44:17 GMT
X-Varnish: 1701844380 1701819611
Age: 101
Via: 1.1 varnish
Connection: keep-alive
```

Body

```
{
  "pagination": {
    "per_page": 3,
    "pages": 66,
    "page": 1,
    "urls": {
      "last": "http://api.discogs.com/database/search?per_page=3&artist=nirvana&release_title=nevermind&page=66",
      "next": "http://api.discogs.com/database/search?per_page=3&artist=nirvana&release_title=nevermind&page=2"
    },
    "items": 198
  },
  "results": [
    {
      "style": [
        "Interview",
        "Grunge"
      ],
      "thumb": "",
      "title": "Nirvana - Nevermind",
      "country": "Australia",
      "format": [
        "DVD",
        "PAL"
      ],
      "uri": "/Nirvana-Nevermind-Classic-Albums/release/2028757",
      "community": {
        "want": 1,
        "have": 5
      },
      "label": [
        "Eagle Vision",
        "Rajon Vision",
        "Classic Albums"
      ],
      "catno": "RV0296",
      "year": "2005",
      "genre": [
        "Non-Music",
        "Rock"
      ],
      "resource_url": "http://api.discogs.com/releases/2028757",
      "type": "release",
      "id": 2028757
    },
    {
      "style": [
        "Interview",
        "Grunge"
      ],
      "thumb": "",
      "title": "Nirvana - Nevermind",
      "country": "France",
      "format": [
        "DVD",
        "PAL"
      ],
      "uri": "/Nirvana-Nevermind-Classic-Albums/release/1852962",
      "community": {
        "want": 4,
        "have": 21
      },
      "label": [
        "Eagle Vision",
        "Classic Albums"
      ],
      "catno": "EV 426200",
      "year": "2005",
      "genre": [
        "Non-Music",
        "Rock"
      ],
      "resource_url": "http://api.discogs.com/releases/1852962",
      "type": "release",
      "id": 1852962
    },
    {
      "style": [
        "Hard Rock",
        "Classic Rock"
      ],
      "thumb": "",
      "format": [
        "UMD"
      ],
      "country": "Europe",
      "barcode": [
        "5 034504 843646"
      ],
      "uri": "/Nirvana-Nevermind/release/3058947",
      "community": {
        "want": 10,
        "have": 3
      },
      "label": [
        "Eagle Vision"
      ],
      "catno": "ERUMD436",
      "genre": [
        "Rock"
      ],
      "title": "Nirvana - Nevermind",
      "resource_url": "http://api.discogs.com/releases/3058947",
      "type": "release",
      "id": 3058947
    }
  ]
}
```

**Response  500**

Headers

```
Status: HTTP/1.1 500 Server Error
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Content-Type: application/json
Server: lighttpd
Content-Length: 32
Date: Tue, 01 Jul 2014 01:08:31 GMT
X-Varnish: 1465587583
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```

Body

```
{
  "message": "Query time exceeded. Please try a simpler query."
}
```

**Response  500**

Headers

```
Status: HTTP/1.1 500 Server Error
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Content-Type: application/json
Server: lighttpd
Content-Length: 32
Date: Tue, 01 Jul 2014 01:08:31 GMT
X-Varnish: 1465587583
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```

Body

```
{
  "message": "An internal server error occurred. (Malformed query?)"
}
```

### Videos
If your application integrates YouTube videos, then third party cookies may be used. You can view YouTube and Google’s cookie policy here.

## Images

The Image resource represents a user-contributed image of a database object, such as Artists or Releases. Image requests require authentication and are subject to rate limiting.

It’s unlikely that you’ll ever have to construct an image URL; images keys on other resources use fully-qualified URLs, including hostname and protocol. To retrieve images, authenticate via OAuth or Discogs Auth and fetch the object that contains the image of interest (e.g., the release, user profile, etc.). The image URL will be in the response using the HTTPS protocol, and requesting that URL should succeed.

## Marketplace

### Inventory 

Returns the list of listings in a user’s inventory. Accepts Pagination parameters.
Basic information about each listing and the corresponding release is provided, suitable for display in a list. For detailed information about the release, make another API call to fetch the corresponding Release.

If you are not authenticated as the inventory owner, only items that have a status of For Sale will be visible.
If you are authenticated as the inventory owner you will get additional weight, format_quantity, external_id, location, and quantity keys. Note that quantity is a read-only field for NearMint users, who will see 1 for all quantity values, regardless of their actual count. If the user is authorized, the listing will contain a in_cart boolean field indicating whether or not this listing is in their cart.

```
GET /users/{username}/inventory{?status,sort,sort_order}
```

Get a seller’s inventory

**Parameters**

- username
string (required) Example: 360vinyl
The username for whose inventory you are fetching

- status
string (optional) Example: for sale
Only show items with this status.

- sort
  - string (optional) Example: price
  Sort items by this field:
  - listed
  - price
  - item (i.e. the title of the release)
  - artist
  - label
  - catno
  - audio
  - status (when authenticated as the inventory owner)
  - location (when authenticated as the inventory owner)

- sort_order
string (optional) Example: asc
Sort items in a particular order (one of asc, desc)

****Response  200****

Headers

```
Status: HTTP/1.1 200 OK
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Cache-Control: public, must-revalidate
Content-Type: application/json
Link: <https://api.discogs.com/users/360vinyl/inventory?per_page=50&page=18>; rel="last", <https://api.discogs.com/users/360vinyl/inventory?per_page=50&page=2>; rel="next"
Server: lighttpd
Content-Length: 36813
Date: Tue, 15 Jul 2014 18:53:23 GMT
X-Varnish: 1701983958
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```

Body

```
{
  "pagination": {
    "per_page": 50,
    "items": 4,
    "page": 1,
    "urls": {},
    "pages": 1
  },
  "listings": [
    {
      "status": "For Sale",
      "price": {
        "currency": "USD",
        "value": 149.99
      },
      "allow_offers": true,
      "sleeve_condition": "Near Mint (NM or M-)",
      "id": 150899904,
      "condition": "Near Mint (NM or M-)",
      "posted": "2014-07-01T10:20:17-07:00",
      "ships_from": "United States",
      "uri": "https://www.discogs.com/sell/item/150899904",
      "comments": "Includes promotional booklet from original purchase!",
      "seller": {
        "username": "rappcats",
        "resource_url": "https://api.discogs.com/users/rappcats",
        "id": 2098225
      },
      "release": {
        "catalog_number": "509990292346, TMR092",
        "resource_url": "https://api.discogs.com/releases/2992668",
        "year": 2011,
        "id": 2992668,
        "description": "Danger Mouse & Daniele Luppi - Rome (LP, Ora + LP, Whi + Album, Ltd, Tip)",
        "artist": "Danger Mouse & Daniele Luppi",
        "title": "Rome",
        "format": "(LP, Ora + LP, Whi + Album, Ltd, Tip)",
        "thumbnail": "https://api-img.discogs.com/CFEw018vfc3LvUQDFtsvkh9FTyA=/fit-in/322x320/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/R-2992668-1310811542.jpeg.jpg"
      },
      "resource_url": "https://api.discogs.com/marketplace/listings/150899904",
      "audio": false
    },
    {
      "status": "For Sale",
      "price": {
        "currency": "USD",
        "value": 49.99
      },
      "allow_offers": false,
      "sleeve_condition": "Very Good Plus (VG+)",
      "id": 155473349,
      "condition": "Very Good Plus (VG+)",
      "posted": "2014-07-01T10:20:17-07:00",
      "ships_from": "United States",
      "uri": "https://www.discogs.com/sell/item/155473349",
      "comments": "Includes slipmats",
      "seller": {
        "username": "rappcats",
        "resource_url": "https://api.discogs.com/users/rappcats",
        "id": 2098225
      },
      "release": {
        "catalog_number": "STH 2222",
        "resource_url": "https://api.discogs.com/releases/1900152",
        "year": 2009,
        "id": 1900152,
        "description": "Various - Stones Throw X Serato (2x12\", Ltd, Cle)",
        "thumbnail": "https://api-img.discogs.com/BfviIBw5nZOA2BHd0xn8Vfu1X_g=/fit-in/600x600/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/R-1900152-1315429257.jpeg.jpg"
      },
      "resource_url": "https://api.discogs.com/marketplace/listings/155473349",
      "audio": false
    },
    {
      "status": "For Sale",
      "price": {
        "currency": "USD",
        "value": 39.99
      },
      "allow_offers": true,
      "sleeve_condition": "Near Mint (NM or M-)",
      "id": 150899171,
      "condition": "Very Good Plus (VG+)",
      "posted": "2014-07-07T11:40:08-07:00",
      "ships_from": "United States",
      "uri": "https://www.discogs.com/sell/item/150899171",
      "comments": "",
      "seller": {
        "username": "rappcats",
        "resource_url": "https://api.discogs.com/users/rappcats",
        "id": 2098225
      },
      "release": {
        "catalog_number": "STH 2172",
        "resource_url": "https://api.discogs.com/releases/1842118",
        "year": 2009,
        "id": 1842118,
        "description": "Last Electro-Acoustic Space Jazz & Percussion Ensemble, The - Summer Suite (CD, MiniAlbum, Ltd)",
        "thumbnail": "https://api-img.discogs.com/pm6PIqf4vEK8S8rCkySA9eKNFgk=/fit-in/455x455/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/R-1842118-1247162514.jpeg.jpg"
      },
      "resource_url": "https://api.discogs.com/marketplace/listings/150899171",
      "audio": false
    },
    {
      "status": "For Sale",
      "price": {
        "currency": "USD",
        "value": 229.99
      },
      "allow_offers": false,
      "sleeve_condition": "Near Mint (NM or M-)",
      "id": 171931719,
      "condition": "Near Mint (NM or M-)",
      "posted": "2014-07-12T16:23:14-07:00",
      "ships_from": "United States",
      "uri": "https://www.discogs.com/sell/item/171931719",
      "comments": "Includes poster, includes download card, includes 7\", in original bag w/ hype sticker. Complete set!",
      "seller": {
        "username": "rappcats",
        "resource_url": "https://api.discogs.com/users/rappcats",
        "id": 2098225
      },
      "release": {
        "catalog_number": "NSD-120",
        "resource_url": "https://api.discogs.com/releases/2791275",
        "year": 2011,
        "id": 2791275,
        "description": "Metal Fingers - Presents Special Herbs The Box Set Vol. 0-9 (Box, Comp, Ltd + 10xLP + 7\")",
        "thumbnail": "https://api-img.discogs.com/wyy8_nChnz_ergzK9gd4wxqr-K0=/fit-in/600x600/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/R-2791275-1301188174.jpeg.jpg"
      },
      "resource_url": "https://api.discogs.com/marketplace/listings/171931719",
      "audio": false
    }
  ]
}
```

**Response  404**

Headers

```
Status: HTTP/1.1 404 Not Found
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Content-Type: application/json
Server: lighttpd
Content-Length: 33
Date: Tue, 01 Jul 2014 01:03:20 GMT
X-Varnish: 1465521729
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```

Body

```
{
  "message": "The request resource was not found."
}
```

**Response  422**

Headers

```
Status: HTTP/1.1 422 Unprocessable Entity
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Cache-Control: public, must-revalidate
Content-Type: application/json
Server: lighttpd
Content-Length: 114
Date: Tue, 15 Jul 2014 19:16:59 GMT
X-Varnish: 1702310957
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```
Body

```
{
  "message": "Invalid status: expected one of All, Deleted, Draft, Expired, For Sale, Sold, Suspended, Violation."
}
```

### Listing

View the data associated with a listing.
If the authorized user is the listing owner the listing will include the weight, format_quantity, external_id, location, and quantity keys. Note that quantity is a read-only field for NearMint users, who will see 1 for all quantity values, regardless of their actual count. If the user is authorized, the listing will contain a in_cart boolean field indicating whether or not this listing is in their cart.

```
GET /marketplace/listings/{listing_id}{?curr_abbr}
```

The Listing resource allows you to view Marketplace listings.

**Parameters**

- listing_id
number (required) Example: 172723812
The ID of the listing you are fetching

- curr_abbr
string (optional) Example: USD
Currency for marketplace data. Defaults to the authenticated users currency. Must be one of the following:
USD GBP EUR CAD AUD JPY CHF MXN BRL NZD SEK ZAR

**Response  200**

Headers

```
Status: HTTP/1.1 200 OK
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Cache-Control: public, must-revalidate
Content-Type: application/json
Server: lighttpd
Content-Length: 780
Date: Tue, 15 Jul 2014 19:59:59 GMT
X-Varnish: 1702965334
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```

Body

```
{
  "status": "For Sale",
  "price": {
    "currency": "USD",
    "value": 120
  },
  "original_price": {
    "curr_abbr": "USD",
    "curr_id": 1,
    "formatted": "$120.00",
    "value": 120.0
  },
  "allow_offers": false,
  "sleeve_condition": "Mint (M)",
  "id": 172723812,
  "condition": "Mint (M)",
  "posted": "2014-07-15T12:55:01-07:00",
  "ships_from": "United States",
  "uri": "https://www.discogs.com/sell/item/172723812",
  "comments": "Brand new... Still sealed!",
  "seller": {
    "username": "Booms528",
    "avatar_url": "https://secure.gravatar.com/avatar/8aa676fcfa2be14266d0ccad88da2cc4?s=500&r=pg&d=mm",
    "resource_url": "https://api.discogs.com/users/Booms528",
    "url": "https://api.discogs.com/users/Booms528",
    "id": 1369620
    "shipping": "Buyer responsible for shipping. Price depends on distance but is usually $5-10.",
    "payment": "PayPal",
    "stats": {
      "rating": "100",
      "stars": 5.0,
      "total": 15
    }
  },
  "shipping_price": {
    "currency": "USD",
    "value": 2.50
  },
  "original_shipping_price": {
    "curr_abbr": "USD",
    "curr_id": 1,
    "formatted": "$2.50",
    "value": 2.5
  },
  "release": {
    "catalog_number": "541125-1, 1-541125 (K1)",
    "resource_url": "https://api.discogs.com/releases/5610049",
    "year": 2014,
    "id": 5610049,
    "description": "LCD Soundsystem - The Long Goodbye: LCD Soundsystem Live At Madison Square Garden (5xLP + Box)",
    "thumbnail": "https://api-img.discogs.com/UsvcarhmrXb0km4QH_dRP8gEf3E=/fit-in/600x600/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/R-5610049-1399500556-9283.jpeg.jpg"
  },
  "resource_url": "https://api.discogs.com/marketplace/listings/172723812",
  "audio": false
}
```

**Response  404**

Headers

```
Status: HTTP/1.1 404 Not Found
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Content-Type: application/json
Server: lighttpd
Content-Length: 33
Date: Tue, 01 Jul 2014 01:03:20 GMT
X-Varnish: 1465521729
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```

Body

```
{
  "message": "The request resource was not found."
}
```

```
POST /marketplace/listings/{listing_id}{?curr_abbr}
```

Edit the data associated with a listing.

If the listing’s status is not For Sale, Draft, or Expired, it cannot be modified – only deleted. To re-list a Sold listing, a new listing must be created.

Authentication as the listing owner is required.

**Parameters**

- listing_id
number (required) Example: 172723812
The ID of the listing you are fetching

- release_id
number (required) Example: 1
The ID of the release you are posting

- condition
  - string (required) Example: Mint
    The condition of the release you are posting. Must be one of the following:
  - Mint (M)
  - Near Mint (NM or M-)
  - Very Good Plus (VG+)
  - Very Good (VG)
  - Good Plus (G+)
  - Good (G)
  - Fair (F)
  - Poor (P)

- sleeve_condition
  - string (optional) Example: Fair
    The condition of the sleeve of the item you are posting. Must be one of the following:
  - Mint (M)
  - Near Mint (NM or M-)
  - Very Good Plus (VG+)
  - Very Good (VG)
  - Good Plus (G+)
  - Good (G)
  - Fair (F)
  - Poor (P)
  - Generic
  - Not Graded
  - No Cover

- price
number (required) Example: 10.00
The price of the item (in the seller’s currency).

- comments
string (optional) Example: This item is wonderful
Any remarks about the item that will be displayed to buyers.

- allow_offers
boolean (optional) Example: true
Whether or not to allow buyers to make offers on the item. Defaults to false.

- status
string (required) Example: Draft
The status of the listing. Defaults to For Sale. Options are For Sale (the listing is ready to be shown on the Marketplace) and Draft (the listing is not ready for public display).

- external_id
string (optional) Example: 10.00
A freeform field that can be used for the seller’s own reference. Information stored here will not be displayed to anyone other than the seller. This field is called “Private Comments” on the Discogs website.

- location
string (optional) Example: 10.00
A freeform field that is intended to help identify an item’s physical storage location. Information stored here will not be displayed to anyone other than the seller. This field will be visible on the inventory management page and will be available in inventory exports via the website.

- weight
number (optional) Example: 10.00
The weight, in grams, of this listing, for the purpose of calculating shipping. Set this field to auto to have the weight automatically estimated for you.

- format_quantity
number (optional) Example: 10.00
The number of items this listing counts as, for the purpose of calculating shipping. This field is called “Counts As” on the Discogs website. Set this field to auto to have the quantity automatically estimated for you.

**Request**

Headers

```
Content-Type: application/json
```

**Response  204**

```
DELETE /marketplace/listings/{listing_id}{?curr_abbr}
```

Permanently remove a listing from the Marketplace.
Authentication as the listing owner is required.

**Parameters**

- listing_id
number (required) Example: 172723812
The ID of the listing you are fetching

**Request**

Headers

```
Content-Type: application/json
```

**Response  204**

### New Listing

```
POST /marketplace/listings{?release_id,condition,sleeve_condition,price,comments,allow_offers,status,external_id,location,weight,format_quantity}
```

Create a Marketplace listing.
Authentication is required; the listing will be added to the authenticated user’s Inventory.

**Parameters**

- release_id
number (required) Example: 1
The ID of the release you are posting

- ondition
  - string (required) Example: Mint
    The condition of the release you are posting. Must be one of the following:
  - Mint (M)
  - Near Mint (NM or M-)
  - Very Good Plus (VG+)
  - Very Good (VG)
  - Good Plus (G+)
  - Good (G)
  - Fair (F)
  - Poor (P)

- sleeve_condition
  - string (optional) Example: Fair
  The condition of the sleeve of the item you are posting. Must be one of the following:
  - Mint (M)
  - Near Mint (NM or M-)
  - Very Good Plus (VG+)
  - Very Good (VG)
  - Good Plus (G+)
  - Good (G)
  - Fair (F)
  - Poor (P)
  - Generic
  - Not Graded
  - No Cover

- price
number (required) Example: 10.00
The price of the item (in the seller’s currency).

- comments
string (optional) Example: This item is wonderful
Any remarks about the item that will be displayed to buyers.

- allow_offers
boolean (optional) Example: true
Whether or not to allow buyers to make offers on the item. Defaults to false.

- status
string (required) Example: Draft
The status of the listing. Defaults to For Sale. Options are For Sale (the listing is ready to be shown on the Marketplace) and Draft (the listing is not ready for public display).

- external_id
string (optional) Example: 10.00
A freeform field that can be used for the seller’s own reference. Information stored here will not be displayed to anyone other than the seller. This field is called “Private Comments” on the Discogs website.

- location
string (optional) Example: 10.00
A freeform field that is intended to help identify an item’s physical storage location. Information stored here will not be displayed to anyone other than the seller. This field will be visible on the inventory management page and will be available in inventory exports via the website.

- weight
number (optional) Example: 10.00
The weight, in grams, of this listing, for the purpose of calculating shipping. Set this field to auto to have the weight automatically estimated for you.

- format_quantity
number (optional) Example: 10.00
The number of items this listing counts as, for the purpose of calculating shipping. This field is called “Counts As” on the Discogs website. Set this field to auto to have the quantity automatically estimated for you.

**Request**

Headers

```
Content-Type: application/json
```

**Response  201**

Body

```
{
  "listing_id": 41578241,
  "resource_url": "https://api.discogs.com/marketplace/listings/41578241"
}
```

**Response  403**

Body

```
{
  "message": "You don't have permission to access this resource."
}
```

### Order 

The Order resource allows you to manage a seller’s Marketplace orders.

```
GET /marketplace/orders/{order_id}
```

View the data associated with an order.
Authentication as the seller is required.

**Parameters**

- order_id
number (required) Example: 1-1
The ID of the order you are fetching

**Response  200**

Headers

```
Status: HTTP/1.1 200 OK
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Cache-Control: public, must-revalidate
Content-Type: application/json
Server: lighttpd
Content-Length: 780
Date: Tue, 15 Jul 2014 19:59:59 GMT
X-Varnish: 1702965334
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```

Body

```
{
  "id": "1-1",
  "resource_url": "https://api.discogs.com/marketplace/orders/1-1",
  "messages_url": "https://api.discogs.com/marketplace/orders/1-1/messages",
  "uri": "https://www.discogs.com/sell/order/1-1",
  "status": "New Order",
  "next_status": [
    "New Order",
    "Buyer Contacted",
    "Invoice Sent",
    "Payment Pending",
    "Payment Received",
    "In Progress",
    "Shipped",
    "Refund Sent",
    "Cancelled (Non-Paying Buyer)",
    "Cancelled (Item Unavailable)",
    "Cancelled (Per Buyer's Request)"
  ],
  "fee": {
    "currency": "USD",
    "value": 2.52
  },
  "created": "2011-10-21T09:25:17-07:00",
  "items": [
    {
      "release": {
        "id": 1,
        "description": "Persuader, The - Stockholm (2x12\")"
      },
      "price": {
        "currency": "USD",
        "value": 42
      },
      "media_condition": "Mint (M)",
      "sleeve_condition": "Mint (M)",
      "id": 41578242
    }
  ],
  "shipping": {
    "currency": "USD",
    "method": "Standard",
    "value": 0
  },
  "shipping_address": "Asdf Exampleton\n234 NE Asdf St.\nAsdf Town, Oregon, 14423\nUnited States\n\nPhone: 555-555-2733\nPaypal address: asdf@example.com",
  "additional_instructions": "please use sturdy packaging.",
  "archived": false,
  "seller": {
    "resource_url": "https://api.discogs.com/users/example_seller",
    "username": "example_seller",
    "id": 1
  },
  "last_activity": "2011-10-21T09:25:17-07:00",
  "buyer": {
    "resource_url": "https://api.discogs.com/users/example_buyer",
    "username": "example_buyer",
    "id": 2
  },
  "total": {
    "currency": "USD",
    "value": 42
  },
  "tracking": {
    "number": "1Z999999999999999",
    "carrier": "UPS",
    "url": "https://www.ups.com/track?tracknum=1Z999999999999999"
  }
}
```

**Response  401**

Headers

```
Status: HTTP/1.1 401 Unauthorized
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Cache-Control: public, must-revalidate
Content-Type: application/json
WWW-Authenticate: OAuth realm="https://api.discogs.com"
Server: lighttpd
Content-Length: 61
Date: Tue, 15 Jul 2014 20:37:49 GMT
X-Varnish: 1703540564
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```

Body

```
{
  "message": "You must authenticate to access this resource."
}
```

```
POST /marketplace/orders/{order_id}
```

Edit the data associated with an order.
Authentication as the seller is required.
The response contains a next_status key – an array of valid next statuses for this order, which you can display to the user in (for example) a dropdown control. This also renders your application more resilient to any future changes in the order status logic.
Changing the order status using this resource will always message the buyer with:

Seller changed status from Old Status to New Status

and does not provide a facility for including a custom message along with the change. For more fine-grained control, use the Add a new message resource, which allows you to simultaneously add a message and change the order status.
If the order status is neither cancelled, Payment Received, nor Shipped, you can change the shipping. Doing so will send an invoice to the buyer and set the order status to Invoice Sent. (For that reason, you cannot set the shipping and the order status in the same request.) Sellers can also set tracking information for an order by providing a tracking object with a tracking number and optionally a carrier. The tracking URL will be automatically generated based on the carrier.

**Parameters**

- order_id
number (required) Example: 1-1
The ID of the order you are fetching

- status
  - string (optional) Example: New Order
  The status of the Order you are updating. Must be one of the following:
  - New Order
  - Buyer Contacted
  - Invoice Sent
  - Payment Pending
  - Payment Received In Progress Shipped
  - Refund Sent
  - Cancelled (Non-Paying Buyer)
  - Cancelled (Item Unavailable)
  - Cancelled (Per Buyer's Request) the order’s current status

Furthermore, the new status must be present in the order’s next_status list. For more information about order statuses, see Edit an order.

- shipping
number (optional) Example: 5.00
The order shipping amount. As a side-effect of setting this value, the buyer is invoiced and the order status is set to Invoice Sent.

- tracking
object (optional)
Set tracking information for the order. Only the seller can set tracking information; buyers will receive a 403 Forbidden response. The tracking object has the following properties:

  - number (required, string): The tracking number provided by the carrier.
  - carrier (optional, string): The shipping carrier. Must be one of: UPS, USPS, DHL, Deutsche Post, La Poste, Royal Mail, PostNL, DHL Germany, Other. If not provided, defaults to no carrier specified.


**Request**

Headers

```
Content-Type: application/json
```

**Response  200**

Body

```
{
  "id": "1-1",
  "resource_url": "https://api.discogs.com/marketplace/orders/1-1",
  "messages_url": "https://api.discogs.com/marketplace/orders/1-1/messages",
  "uri": "https://www.discogs.com/sell/order/1-1",
  "status": "Invoice Sent",
  "next_status": [
    "New Order",
    "Buyer Contacted",
    "Invoice Sent",
    "Payment Pending",
    "Payment Received",
    "In Progress",
    "Shipped",
    "Refund Sent",
    "Cancelled (Non-Paying Buyer)",
    "Cancelled (Item Unavailable)",
    "Cancelled (Per Buyer's Request)"
  ],
  "fee": {
    "currency": "USD",
    "value": 2.52
  },
  "created": "2011-10-21T09:25:17-07:00",
  "items": [
    {
      "release": {
        "id": 1,
        "description": "Persuader, The - Stockholm (2x12\")"
      },
      "price": {
        "currency": "USD",
        "value": 42
      },
      "id": 41578242
    }
  ],
  "shipping": {
    "currency": "USD",
    "method": "Standard",
    "value": 5
  },
  "shipping_address": "Asdf Exampleton\n234 NE Asdf St.\nAsdf Town, Oregon, 14423\nUnited States\n\nPhone: 555-555-2733\nPaypal address: asdf@example.com",
  "additional_instructions": "please use sturdy packaging.",
  "archived": false,
  "seller": {
    "resource_url": "https://api.discogs.com/users/example_seller",
    "username": "example_seller",
    "id": 1
  },
  "last_activity": "2011-10-22T19:18:53-07:00",
  "buyer": {
    "resource_url": "https://api.discogs.com/users/example_buyer",
    "username": "example_buyer",
    "id": 2
  },
  "total": {
    "currency": "USD",
    "value": 47
  },
  "tracking": {
    "number": "1Z999999999999999",
    "carrier": "UPS",
    "url": "https://www.ups.com/track?tracknum=1Z999999999999999"
  }
}
```

**Response  400**

Body

```
{
  "message": "Tracking number is required."
}
```

**Response  403**

Body

```
{
  "message": "Only the seller can set tracking information."
}
```

### List Orders

Returns a list of the authenticated user’s orders. Accepts Pagination parameters.

```
GET /marketplace/orders{?status,created_after,created_before,sort,sort_order}
```

Returns a list of the authenticated user’s orders. Accepts Pagination parameters.

**Parameters**

- status
  - string (optional) Example: 1-1
  Only show orders with this status. Valid status keys are:
  - All
  - New Order
  - Buyer Contacted
  - Invoice Sent
  - Payment Pending
  - Payment Received In Progress Shipped
  - Merged
  - Order Changed
  - Refund Sent
  - Cancelled
  - Cancelled (Non-Paying Buyer)
  - Cancelled (Item Unavailable)
  - Cancelled (Per Buyer's Request) Cancelled (Refund Received)

- created_after
string (optional) Example: 2019-06-24T20:58:58Z
Only show orders created after this ISO 8601 timestamp.

- created_before
string (optional) Example: 2019-06-24T20:58:58Z
Only show orders created before this ISO 8601 timestamp.

- archived
boolean (optional) Example: true
Only show orders with a specific archived status. If no key is provided, both statuses are returned. Valid archived keys are:
  - true
  - false

- sort
string (optional) Example: 1-1
Sort items by this field (see below). Valid sort keys are:
  - id
  - buyer
  - created
  - status
  - last_activity

- sort_order
string (optional) Example: 1-1
Sort items in a particular order (one of asc, desc)

**Response  200**

Headers

```
Status: HTTP/1.1 200 OK
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Cache-Control: public, must-revalidate
Content-Type: application/json
Server: lighttpd
Content-Length: 780
Date: Tue, 15 Jul 2014 19:59:59 GMT
X-Varnish: 1702965334
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```

Body

```
{
  "pagination": {
    "per_page": 50,
    "pages": 1,
    "page": 1,
    "items": 1,
    "urls": {}
  },
  "orders": [
    {
      "status": "New Order",
      "fee": {
        "currency": "USD",
        "value": 2.52
      },
      "created": "2011-10-21T09:25:17-07:00",
      "items": [
        {
          "release": {
            "id": 1,
            "description": "Persuader, The - Stockholm (2x12\")",
            "resource_url": "https://api.discogs.com/releases/1",
            "thumbnail": "http://api-img.discogs.com/souG2t4I8ZFVK3kHVtD3zjGvd_Y=/fit-in/300x300/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/R-1-1193812031.jpeg.jpg"
          },
          "price": {
            "currency": "USD",
            "value": 42.0
          },
          "id": 41578242
        }
      ],
      "shipping": {
        "currency": "USD",
        "method": "Standard",
        "value": 0.0
      },
      "shipping_address": "Asdf Exampleton\n234 NE Asdf St.\nAsdf Town, Oregon, 14423\nUnited States\n\nPhone: 555-555-2733\nPaypal address: asdf@example.com",
      "additional_instructions": "please use sturdy packaging.",
      "archived": false,
      "seller": {
        "resource_url": "https://api.discogs.com/users/example_seller",
        "username": "example_seller",
        "id": 1
      },
      "last_activity": "2011-10-21T09:25:17-07:00",
      "buyer": {
        "resource_url": "https://api.discogs.com/users/example_buyer",
        "username": "example_buyer",
        "id": 2
      },
      "total": {
        "currency": "USD",
        "value": 42.0
      },
      "id": "1-1"
      "resource_url": "https://api.discogs.com/marketplace/orders/1-1",
      "messages_url": "https://api.discogs.com/marketplace/orders/1-1/messages",
      "uri": "https://www.discogs.com/sell/order/1-1",
      "next_status": [
        "New Order",
        "Buyer Contacted",
        "Invoice Sent",
        "Payment Pending",
        "Payment Received",
        "In Progress",
        "Shipped",
        "Refund Sent",
        "Cancelled (Non-Paying Buyer)",
        "Cancelled (Item Unavailable)",
        "Cancelled (Per Buyer's Request)"
      ]
    }
  ]
}
```

### List Order Messages 

```
GET /marketplace/orders/{order_id}/messages
```

Returns a list of the order’s messages with the most recent first. Accepts Pagination parameters.
Authentication as the seller is required.

**Parameters**

- order_id
string (required) Example: 1-1
The ID of the order you are fetching

**Response  200**

Headers

```
Status: HTTP/1.1 200 OK
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Cache-Control: public, must-revalidate
Content-Type: application/json
Server: lighttpd
Content-Length: 780
Date: Tue, 15 Jul 2014 19:59:59 GMT
X-Varnish: 1702965334
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```

Body

```
{
  "pagination": {
    "per_page": 50,
    "items": 8,
    "page": 1,
    "urls": {},
    "pages": 1
  },
  "messages": [
    {
      "refund": {
        "amount": 5,
        "order": {
          "resource_url": "https://api.discogs.com/marketplace/orders/845236-9",
          "id": "845236-9"
        }
      },
      "timestamp": "2015-06-02T13:17:54-07:00",
      "message": "example_buyer received refund of $5.00.",
      "type": "refund_received",
      "order": {
        "resource_url": "https://api.discogs.com/marketplace/orders/845236-9",
        "id": "845236-9"
      },
      "subject": ""
    },
    {
      "refund": {
        "amount": 5,
        "order": {
          "resource_url": "https://api.discogs.com/marketplace/orders/845236-9",
          "id": "845236-9"
        }
      },
      "timestamp": "2015-06-02T13:17:44-07:00",
      "message": "example_seller sent refund of $5.00.",
      "type": "refund_sent",
      "order": {
        "resource_url": "https://api.discogs.com/marketplace/orders/845236-9",
        "id": "845236-9"
      },
      "subject": ""
    },
    {
      "from": {
        "id": 1001,
        "username": "example_seller",
        "avatar_url": "https://secure.gravatar.com/avatar/1ddcc19fb43551fb86c143465f773282?s=300&r=pg&d=mm",
        "resource_url": "https://api.discogs.com/users/example_seller"
      },
      "timestamp": "2015-06-02T13:17:07-07:00",
      "message": "Thank you for your order!",
      "type": "message",
      "order": {
        "resource_url": "https://api.discogs.com/marketplace/orders/845236-9",
        "id": "845236-9"
      },
      "subject": "New Message - Order #845236-9 - TZ Goes Beyond 10! + 1 more item"
    },
    {
      "status_id": 6,
      "timestamp": "2015-06-02T13:16:57-07:00",
      "actor": {
        "username": "example_seller",
        "resource_url": "https://api.discogs.com/users/example_seller"
      },
      "message": "example_buyer changed the order status to Shipped.",
      "type": "status",
      "order": {
        "resource_url": "https://api.discogs.com/marketplace/orders/845236-9",
        "id": "845236-9"
      },
      "subject": ""
    },
    {
      "status_id": 5,
      "timestamp": "2015-06-02T13:16:51-07:00",
      "actor": {
        "username": "example_seller",
        "resource_url": "https://api.discogs.com/users/example_seller"
      },
      "message": "example_buyer changed the order status to Payment Received.",
      "type": "status",
      "order": {
        "resource_url": "https://api.discogs.com/marketplace/orders/845236-9",
        "id": "845236-9"
      },
      "subject": ""
    },
    {
      "status_id": 3,
      "timestamp": "2015-06-02T13:16:27-07:00",
      "actor": {
        "username": "example_seller",
        "resource_url": "https://api.discogs.com/users/example_seller"
      },
      "message": "example_buyer changed the order status to Invoice Sent.",
      "type": "status",
      "order": {
        "resource_url": "https://api.discogs.com/marketplace/orders/845236-9",
        "id": "845236-9"
      },
      "subject": ""
    },
    {
      "timestamp": "2015-06-02T13:16:27-07:00",
      "original": 0,
      "new": 5,
      "message": "example_seller set the shipping price to $5.00.",
      "type": "shipping",
      "order": {
        "resource_url": "https://api.discogs.com/marketplace/orders/845236-9",
        "id": "845236-9"
      },
      "subject": ""
    },
    {
      "status_id": 1,
      "timestamp": "2015-06-02T13:16:12-07:00",
      "actor": {
        "username": "example_seller",
        "resource_url": "https://api.discogs.com/users/example_seller"
      },
      "message": "example_seller created the order by merging orders #845236-7, #845236-8.",
      "type": "status",
      "order": {
        "resource_url": "https://api.discogs.com/marketplace/orders/845236-9",
        "id": "845236-9"
      },
      "subject": ""
    }
  ]
}
```

```
POST /marketplace/orders/{order_id}/messages
```

Adds a new message to the order’s message log.
When posting a new message, you can simultaneously change the order status. If you do, the message will automatically be prepended with:
Seller changed status from Old Status to New Status
While message and status are each optional, one or both must be present.

**Parameters**

- order_id
string (required) Example: 1-1
The ID of the order you are fetching

- message
string (optional) Example: hello world

- status
string (optional) Example: New Order

**Request**

Headers

```
Content-Type: application/json
```

**Response  201**

Body

```
{
  "from": {
    "username": "example_seller",
    "resource_url": "https://api.discogs.com/users/example_seller"
  },
  "message": "Seller changed status from Payment Received to Shipped\n\nYour order is on its way, tracking number #foobarbaz!",
  "order": {
    "resource_url": "https://api.discogs.com/marketplace/orders/1-1",
    "id": "1-1"
  },
  "timestamp": "2011-11-18T15:32:42-07:00",
  "subject": "Discogs Order #1-1, Stockholm"
}
```

**Response  403**

Body

```
{
  "message": "You don't have permission to access this resource."
}
```

### Fee

```
GET /marketplace/fee/{price}
```

The Fee resource allows you to quickly calculate the fee for selling an item on the Marketplace.

**Parameters**

- price
number (optional) Example: 10.00
The price to calculate a fee from

**Response  200**

Headers

```
Status: HTTP/1.1 200 OK
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Cache-Control: public, must-revalidate
Content-Type: application/json
Server: lighttpd
Content-Length: 780
Date: Tue, 15 Jul 2014 19:59:59 GMT
X-Varnish: 1702965334
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```

Body

```
{
    "value": 0.42,
    "currency": "USD",
}
```

### Fee with currency

```
GET /marketplace/fee/{price}/{currency}
```

The Fee resource allows you to quickly calculate the fee for selling an item on the Marketplace given a particular currency.

**Parameters**

- price
number (optional) Example: 10.00
The price to calculate a fee from

- currency
string (optional) Example: USD
Defaults to USD. Must be one of the following:
USD GBP EUR CAD AUD JPY CHF MXN BRL NZD SEK ZAR

**Response  200**

Headers

```
Status: HTTP/1.1 200 OK
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Cache-Control: public, must-revalidate
Content-Type: application/json
Server: lighttpd
Content-Length: 780
Date: Tue, 15 Jul 2014 19:59:59 GMT
X-Varnish: 1702965334
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```

Body

```
{
    "value": 0.42,
    "currency": "USD",
}
```

### Price Suggestions

```
GET /marketplace/price_suggestions/{release_id}
```

Retrieve price suggestions for the provided Release ID. If no suggestions are available, an empty object will be returned.
Authentication is required, and the user needs to have filled out their seller settings. Suggested prices will be denominated in the user’s selling currency.

**Parameters**
- release_id
number (required) Example: 1
The release ID to calculate a price from.

**Response  200**

Headers

```
Status: HTTP/1.1 200 OK
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Cache-Control: public, must-revalidate
Content-Type: application/json
Server: lighttpd
Content-Length: 780
Date: Tue, 15 Jul 2014 19:59:59 GMT
X-Varnish: 1702965334
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```

Body

```
{
  "Very Good (VG)": {
    "currency": "USD",
    "value": 6.7827501
  },
  "Good Plus (G+)": {
    "currency": "USD",
    "value": 3.7681945000000003
  },
  "Near Mint (NM or M-)": {
    "currency": "USD",
    "value": 12.8118613
  },
  "Good (G)": {
    "currency": "USD",
    "value": 2.2609167
  },
  "Very Good Plus (VG+)": {
    "currency": "USD",
    "value": 9.7973057
  },
  "Mint (M)": {
    "currency": "USD",
    "value": 14.319139100000001
  },
  "Fair (F)": {
    "currency": "USD",
    "value": 1.5072778000000002
  },
  "Poor (P)": {
    "currency": "USD",
    "value": 0.7536389000000001
  }
}
```

### Release Statistics 

```
GET /marketplace/stats/{release_id}{?curr_abbr}
```

Retrieve marketplace statistics for the provided Release ID. These statistics reflect the state of the release in the marketplace currently, and include the number of items currently for sale, lowest listed price of any item for sale, and whether the item is blocked for sale in the marketplace.

Authentication is optional. Authenticated users will by default have the lowest currency expressed in their own buyer currency, configurable in buyer settings, in the absence of the curr_abbr query parameter to specify the currency. Unauthenticated users will have the price expressed in US Dollars, if no curr_abbr is provided.

Releases that have no items for sale in the marketplace will return a body with null data in the lowest_price and num_for_sale keys. Releases that are blocked for sale will also have null data for these keys.

**Parameters**

- release_id
number (required) Example: 1
The release ID whose stats are desired

- curr_abbr
string (optional) Example: USD
Currency for marketplace data. Defaults to the authenticated users currency. Must be one of the following:
USD GBP EUR CAD AUD JPY CHF MXN BRL NZD SEK ZAR

**Response  200**

Headers

```
Status: HTTP/1.1 200 OK
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Cache-Control: public, must-revalidate
Content-Type: application/json
Server: lighttpd
Content-Length: 780
Date: Tue, 15 Jul 2014 19:59:59 GMT
X-Varnish: 1702965334
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```

Body

```
{
  "lowest_price": {
    "currency": "USD",
    "value": 2.09
  },
  "num_for_sale": 26,
  "blocked_from_sale": false
}
```

## Inventory Export 

### Export your inventory 

```
POST /inventory/export
```

Request an export of your inventory as a CSV.

**Response  200**

Headers

```
Content-Type: application/json
Location: https://api.discogs.com/inventory/export/599632
```

**Response  401**

```
Headers
Content-Type: application/json
```

**Response  409**

Headers

```
Content-Type: application/json
```

### Get recent exports

```
GET /inventory/export
```

Get a list of all recent exports of your inventory. Accepts Pagination parameters.

**Response  200**

Headers

```
Content-Type: application/json
```

Body

```
{
  "items": [
    {
      "status": "success",
      "created_ts": "2018-09-27T12:59:02",
      "url": "https://api.discogs.com/inventory/export/599632",
      "finished_ts": "2018-09-27T12:59:02",
      "download_url": "https://api.discogs.com/inventory/export/599632/download",
      "filename": "cburmeister-inventory-20180927-1259.csv",
      "id": 599632
    }
  ],
  "pagination": {
    "per_page": 50,
    "items": 15,
    "page": 1,
    "urls": {},
    "pages": 1
  }
}
```

**Response  401**

Headers

```
Content-Type: application/json
```

### Get an export

```
GET /inventory/export/{id}
```

Get details about the status of an inventory export.

**Parameters**

- id
number (required) 
Id of the export.

**Request**

Headers

```
Content-Type: multipart/form-data
If-Modified-Since: Thu, 27 Sep 2018 12:50:39 GMT
```

**Response  200**

Headers

```
Content-Type: application/json
Last-Modified: Thu, 27 Sep 2018 12:59:02 GMT
```

Body

```
{
  "status": "success",
  "created_ts": "2018-09-27T12:50:39",
  "url": "https://api.discogs.com/inventory/export/599632",
  "finished_ts": "2018-09-27T12:59:02",
  "download_url": "https://api.discogs.com/inventory/export/599632/download",
  "filename": "cburmeister-inventory-20180927-1259.csv",
  "id": 599632
}
```

**Response  304**

**Response  401**

Headers

```
Content-Type: application/json
```

**Response  404**

Headers

```
Content-Type: application/json
```

### Download an export

```
GET /inventory/export/{id}/download
```

Download the results of an inventory export.

**Parameters**

- id
number (required) 
Id of the export.

**Response  200**

Headers

```
Content-Type: text/csv; charset=utf-8
Content-Disposition: attachment; filename=cburmeister-inventory-20180927-1259.csv
```

**Response  401**

Headers

```
Content-Type: application/json
```

**Response  404**

Headers

```
Content-Type: application/json
```

## Inventory Upload

### Add inventory

```
POST /inventory/upload/add
```

Upload a CSV of listings to add to your inventory.

#### File structure

The file you upload must be a comma-separated CSV. The first row must be a header with lower case field names.

Here’s an example:

```
release_id,price,media_condition,comments
123,1.50,Mint (M),Comments about this release for sale
456,2.50,Near Mint (NM or M-),More comments
7890,3.50,Good (G),Signed vinyl copy
```

#### Things to note:
These listings will be marked “For Sale” immediately. Currency information will be pulled from your marketplace settings. Any fields that aren’t optional or required will be ignored.

#### Required fields
- release_id - Must be a number. This value corresponds with the Discogs database Release ID.

- price - Must be a number.

- media_condition - Must be a valid condition (see below).

#### Optional fields
- sleeve_condition - Must be a valid condition (see below).

- comments accept_offer - Must be Y or N.

- location - Private free-text field to help identify an item’s physical location.

- external_id - Private notes or IDs for your own reference.

- weight - In grams. Must be a non-negative integer.

- format_quantity - Number of items that this item counts as (for shipping).

#### Conditions
When you specify a media condition, it must exactly match one of these:

- Mint (M)

- Near Mint (NM or M-)

- Very Good Plus (VG+)

- Very Good (VG)

- Good Plus (G+)

- Good (G)

- Fair (F)

- Poor (P)

Sleeve condition may be any of the above, or:

- Not Graded

- Generic

- No Cover

**Parameters**

- upload
file (required) 
The CSV file of items to add to your inventory.

**Response  200**

Headers
```
Content-Type: application/json
Location: https://api.discogs.com/inventory/upload/599632
```
**Response  401**
Headers
```
Content-Type: application/json
```
**Response  409**
Headers
```
Content-Type: application/json
```
**Response  415**
Headers
```
Content-Type: application/json
```
**Response  422**
Headers
```
Content-Type: application/json
```

### Change inventory 

```
POST /inventory/upload/change
```

Upload a CSV of listings to change in your inventory.

#### File structure
The file you upload must be a comma-separated CSV. The first row must be a header with lower case field names.

Here’s an example:

```
release_id,price,media_condition,comments
123,1.50,Mint (M),Comments about this release for sale
456,2.50,Near Mint (NM or M-),More comments
7890,3.50,Good (G),Signed vinyl copy
```

#### Things to note:
These listings will be marked “For Sale” immediately. Currency information will be pulled from your marketplace settings. Any fields that aren’t optional or required will be ignored.

#### Required fields
- release_id - Must be a number. This value corresponds with the Discogs database Release ID.
#### Optional fields (at least one required)
- price -

- media_condition - Must be a valid condition (see below).

- sleeve_condition - Must be a valid condition (see below).

- comments

- accept_offer - Must be Y or N.

- external_id - Private notes or IDs for your own reference.

- location - Private free-text field to help identify an item’s physical location.

- weight - In grams. Must be a non-negative integer.

- format_quantity - Number of items that this item counts as (for shipping).

#### Conditions
When you specify a media condition, it must exactly match one of these:

- Mint (M)

- Near Mint (NM or M-)

- Very Good Plus (VG+)

- Very Good (VG)

- Good Plus (G+)

- Good (G)

 -Fair (F)

 -Poor (P)

Sleeve condition may be any of the above, or:

- Not Graded

- Generic

- No Cover

**Parameters**
- upload file (required). The CSV file of items to alter in your inventory.

**Response  200**
Headers
```
Content-Type: application/json
Location: https://api.discogs.com/inventory/upload/599632
```
**Response  401**
Headers
```
Content-Type: application/json
```
**Response  409**
Headers
```
Content-Type: application/json
```
**Response  415**
Headers
```
Content-Type: application/json
```
**Response  422**
Headers
```
Content-Type: application/json
```

### Delete inventory 

```
POST /inventory/upload/delete
```

Upload a CSV of listings to delete from your inventory.

#### File structure
The file you upload must be a comma-separated CSV. The first row must be a header with lower case field names.

Here’s an example:

```
listing_id
12345678
98765432
31415926
```

#### Required fields
- listing_id - Must be a number. This is the ID of the listing you wish to delete.
**Parameters**
- upload `file` (required). The CSV file listing items to remove from your inventory.

**Response  200**
Headers
```
Content-Type: application/json
Location: https://api.discogs.com/inventory/upload/599632
```
**Response  401**
Headers
```
Content-Type: application/json
```
**Response  409**
Headers
```
Content-Type: application/json
```
**Response  415**
Headers
```
Content-Type: application/json
```
**Response  422**
Headers
```
Content-Type: application/json
```

### Get recent uploads 
```
GET /inventory/upload
```
Get a list of all recent inventory uploads. Accepts Pagination parameters.

**Response  200**
Headers
```
Content-Type: application/json
```
Body
```
{
  "items": [
    {
      "status": "success",
      "results": "CSV file contains 1 records.<p>Processed 1 records.",
      "created_ts": "2017-12-18T09:20:28",
      "finished_ts": "2017-12-18T09:20:29",
      "filename": "add.csv",
      "type": "change",
      "id": 119615
    }
  ],
  "pagination": {
    "per_page": 50,
    "items": 1,
    "page": 1,
    "urls": {},
    "pages": 1
  }
}
```
**Response  401**
Headers
```
Content-Type: application/json
```

### Get an upload 
```
GET /inventory/upload/{id}
```
Get details about the status of an inventory upload.

**Parameters**
- id `number` (required). Id of the export.

**Request**
Headers
```
Content-Type: multipart/form-data
If-Modified-Since: Thu, 27 Sep 2018 09:20:28 GMT
```
**Response  200**
Headers
```
Content-Type: application/json
```
Body
```
{
  "status": "success",
  "results": "CSV file contains 1 records.<p>Processed 1 records.",
  "created_ts": "2017-12-18T09:20:28",
  "finished_ts": "2017-12-18T09:20:29",
  "filename": "add.csv",
  "type": "change",
  "id": 119615
}
```
**Response  304**
**Response  401**
Headers
```
Content-Type: application/json
```
**Response  404**
Headers
```
Content-Type: application/json
```

## User Identity 
### Identity 
```
GET /oauth/identity
```
Retrieve basic information about the authenticated user.
You can use this resource to find out who you’re authenticated as, and it also doubles as a good sanity check to ensure that you’re using OAuth correctly.
For more detailed information, make another request for the user’s Profile.

**Request**
Body
```
GET https://api.discogs.com/oauth/identity
```
**Response  200**
Headers
```
Status: HTTP/1.1 200 OK
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Cache-Control: public, must-revalidate
Content-Type: application/json
Content-Encoding: gzip
Content-Location: https://api.discogs.com/oauth/identity?oauth_body_hash=some_hash&oauth_nonce=...
Server: lighttpd
Content-Length: 127
Date: Tue, 15 Jul 2014 18:44:17 GMT
X-Varnish: 1718276369
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```
Body
```
{
  "id": 1,
  "username": "example",
  "resource_url": "https://api.discogs.com/users/example",
  "consumer_name": "Your Application Name"
}
```
**Response  401**
Headers
```
Status: HTTP/1.1 401 Unauthorized
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Cache-Control: public, must-revalidate
Content-Type: application/json
WWW-Authenticate: OAuth realm="https://api.discogs.com"
Server: lighttpd
Content-Length: 61
Date: Wed, 16 Jul 2014 18:42:38 GMT
X-Varnish: 1718433436
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```
Body
```
{
  "message": "You must authenticate to access this resource."
}
```
### Profile 
```
GET /users/{username}
```
Retrieve a user by username.
If authenticated as the requested user, the email key will be visible, and the num_list count will include the user’s private lists.
If authenticated as the requested user or the user’s collection/wantlist is public, the num_collection / num_wantlist keys will be visible.

**Parameters**
- username `string` (required) Example: rodneyfool. The username of whose profile you are requesting.

**Response  200**
Headers
```
Status: HTTP/1.1 200 OK
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Cache-Control: public, must-revalidate
Content-Type: application/json
Server: lighttpd
Content-Length: 858
Date: Wed, 16 Jul 2014 18:46:21 GMT
X-Varnish: 1718492795
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```
Body
```
{
  "profile": "I am a software developer for Discogs.\r\n\r\n[img=http://i.imgur.com/IAk3Ukk.gif]",
  "wantlist_url": "https://api.discogs.com/users/rodneyfool/wants",
  "rank": 149,
  "num_pending": 61,
  "id": 1578108,
  "num_for_sale": 0,
  "home_page": "",
  "location": "I live in the good ol' Pacific NW",
  "collection_folders_url": "https://api.discogs.com/users/rodneyfool/collection/folders",
  "username": "rodneyfool",
  "collection_fields_url": "https://api.discogs.com/users/rodneyfool/collection/fields",
  "releases_contributed": 5,
  "registered": "2012-08-15T21:13:36-07:00",
  "rating_avg": 3.47,
  "num_collection": 78,
  "releases_rated": 116,
  "num_lists": 0,
  "name": "Rodney",
  "num_wantlist": 160,
  "inventory_url": "https://api.discogs.com/users/rodneyfool/inventory",
  "avatar_url": "http://www.gravatar.com/avatar/55502f40dc8b7c769880b10874abc9d0?s=52&r=pg&d=mm",
  "banner_url": "https://img.discogs.com/dhuJe-pRJmod7hN3cdVi2PugEh4=/1600x400/filters:strip_icc():format(jpeg)/discogs-banners/B-1578108-user-1436314164-9231.jpg.jpg",
  "uri": "https://www.discogs.com/user/rodneyfool",
  "resource_url": "https://api.discogs.com/users/rodneyfool",
  "buyer_rating": 100.00,
  "buyer_rating_stars": 5,
  "buyer_num_ratings": 144,
  "seller_rating": 100.00,
  "seller_rating_stars": 5,
  "seller_num_ratings": 21,
  "curr_abbr": "USD",
}
```

```
POST /users/{username}
```
Edit a user’s profile data.

Authentication as the user is required.

**Parameters**
- username `string` (required) Example: vreon
The username of the user.

- name `string` (optional) Example: Nicolas Cage
The real name of the user.

- home_page `string` (optional) Example: www.discogs.com
The user’s website.

- location `string` (optional) Example: Portland
The geographical location of the user.

- profile `string` (optional) Example: I am a Discogs user!
Biographical information about the user.

- curr_abbr `string` (optional) Example: USD
Currency for marketplace data. Must be one of the following:
USD GBP EUR CAD AUD JPY CHF MXN BRL NZD SEK ZAR

**Response  200**
Body
```
{
  "id": 1,
  "username": "example",
  "name": "Example Sampleman",
  "email": "sampleman@example.com",
  "resource_url": "https://api.discogs.com/users/example",
  "inventory_url": "https://api.discogs.com/users/example/inventory",
  "collection_folders_url": "https://api.discogs.com/users/example/collection/folders",
  "collection_fields_url": "https://api.discogs.com/users/example/collection/fields",
  "wantlist_url": "https://api.discogs.com/users/example/wants",
  "uri": "https://www.discogs.com/user/example",
  "profile": "This is some [b]different sample[/b] data!",
  "home_page": "http://www.example.com",
  "location": "Australia",
  "registered": "2011-08-30 14:21:45-07:00",
  "num_lists": 0,
  "num_for_sale": 6,
  "num_collection": 4,
  "num_wantlist": 5,
  "num_pending": 10,
  "releases_contributed": 15,
  "rank": 30,
  "releases_rated": 4,
  "rating_avg": 2.5
}
```
**Response  403**
Body
```
{
  "message": "You don't have permission to access this resource."
}
```

### User Submissions 
The Submissions resource represents all edits that a user makes to releases, labels, and artist.

```
GET /users/{username}/submissions
```

Retrieve a user’s submissions by username. Accepts Pagination parameters.

**Parameters**
- username `string` (required) Example: shooezgirl
The username of the submissions you are trying to fetch.

**Response  200**
Body
```
{
  "pagination": {
    "items": 3,
    "page": 1,
    "pages": 1,
    "per_page": 50,
    "urls": {}
  },
  "submissions": {
    "artists": [
      {
        "data_quality": "Needs Vote",
        "id": 240177,
        "name": "Grimy",
        "namevariations": [
          "Grimmy"
        ],
        "releases_url": "https://api.discogs.com/artists/240177/releases",
        "resource_url": "https://api.discogs.com/artists/240177",
        "uri": "https://www.discogs.com/artist/240177-Grimy"
      }
    ],
    "labels": [],
    "releases": [
      {
        "artists": [
          {
            "anv": "",
            "id": 17035,
            "join": "",
            "name": "Chaka Khan",
            "resource_url": "https://api.discogs.com/artists/17035",
            "role": "",
            "tracks": ""
          }
        ],
        "community": {
          "contributors": [
            {
              "resource_url": "https://api.discogs.com/users/shooezgirl",
              "username": "shooezgirl"
            },
            {
              "resource_url": "https://api.discogs.com/users/Diognes_The_Fox",
              "username": "Diognes_The_Fox"
            }
          ],
          "data_quality": "Needs Vote",
          "have": 0,
          "rating": {
            "average": 0,
            "count": 0
          },
          "status": "Accepted",
          "submitter": {
            "resource_url": "https://api.discogs.com/users/shooezgirl",
            "username": "shooezgirl"
          },
          "want": 0
        },
        "companies": [],
        "country": "US",
        "data_quality": "Needs Vote",
        "date_added": "2014-03-25T14:52:18-07:00",
        "date_changed": "2014-05-14T13:28:21-07:00",
        "estimated_weight": 60,
        "format_quantity": 1,
        "formats": [
          {
            "descriptions": [
              "7\"",
              "45 RPM",
              "Promo"
            ],
            "name": "Vinyl",
            "qty": "1"
          }
        ],
        "genres": [
          "Funk / Soul"
        ],
        "id": 5531861,
        "images": [
          {
            "height": 594,
            "resource_url": "https://api-img.discogs.com/i3jf6S4S7LMNBuWxstCxAQs2Rw0=/fit-in/600x594/filters:strip_icc():format(jpeg):mode_rgb():quality(96)/discogs-images/R-5531861-1400099290-9223.jpeg.jpg",
            "type": "primary",
            "uri": "https://api-img.discogs.com/i3jf6S4S7LMNBuWxstCxAQs2Rw0=/fit-in/600x594/filters:strip_icc():format(jpeg):mode_rgb():quality(96)/discogs-images/R-5531861-1400099290-9223.jpeg.jpg",
            "uri150": "https://api-img.discogs.com/xMSwaqP2T8SNwDUTO-gXmHXWt6s=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb()/discogs-images/R-5531861-1400099290-9223.jpeg.jpg",
            "width": 600
          },
          {
            "height": 600,
            "resource_url": "https://api-img.discogs.com/sT0plDMoOcfJz-1JEui8XQr69kw=/fit-in/600x600/filters:strip_icc():format(jpeg):mode_rgb():quality(96)/discogs-images/R-5531861-1400099290-9749.jpeg.jpg",
            "type": "secondary",
            "uri": "https://api-img.discogs.com/sT0plDMoOcfJz-1JEui8XQr69kw=/fit-in/600x600/filters:strip_icc():format(jpeg):mode_rgb():quality(96)/discogs-images/R-5531861-1400099290-9749.jpeg.jpg",
            "uri150": "https://api-img.discogs.com/f2QfdjBjpuP-Eht4DVSlCfPtTe8=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb()/discogs-images/R-5531861-1400099290-9749.jpeg.jpg",
            "width": 600
          }
        ],
        "labels": [
          {
            "catno": "7-28671",
            "entity_type": "1",
            "id": 1000,
            "name": "Warner Bros. Records",
            "resource_url": "https://api.discogs.com/labels/1000"
          }
        ],
        "master_id": 174698,
        "master_url": "https://api.discogs.com/masters/174698",
        "notes": "Promotion Not for Sale",
        "released": "1986",
        "released_formatted": "1986",
        "resource_url": "https://api.discogs.com/releases/5531861",
        "series": [],
        "status": "Accepted",
        "styles": [
          "Rhythm & Blues"
        ],
        "thumb": "https://api-img.discogs.com/xMSwaqP2T8SNwDUTO-gXmHXWt6s=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb()/discogs-images/R-5531861-1400099290-9223.jpeg.jpg",
        "title": "Love Of A Lifetime",
        "uri": "https://www.discogs.com/Chaka-Khan-Love-Of-A-Lifetime/release/5531861",
        "videos": [
          {
            "description": "Chaka Khan - Love you all my lifetime (live)",
            "duration": 285,
            "embed": true,
            "title": "Chaka Khan - Love you all my lifetime (live)",
            "uri": "https://www.youtube.com/watch?v=WOrCYzfchTI"
          },
          {
            "description": "Chaka Khan - Love Of A Lifetime [Official Video]",
            "duration": 257,
            "embed": true,
            "title": "Chaka Khan - Love Of A Lifetime [Official Video]",
            "uri": "https://www.youtube.com/watch?v=5K6-q2VqJdE"
          },
          {
            "description": "CHAKA KHAN - COLTRANE DREAMS  ~{The 45 VERSION}~",
            "duration": 151,
            "embed": true,
            "title": "CHAKA KHAN - COLTRANE DREAMS  ~{The 45 VERSION}~",
            "uri": "https://www.youtube.com/watch?v=11eaG7KdS9g"
          }
        ],
        "year": 1986
      }
    ]
  }
}
```

### User Contributions 
The Contributions resource represents releases, labels, and artists submitted by a user.

```
GET /users/{username}/contributions{?sort,sort_order}
```

Retrieve a user’s contributions by username. Accepts Pagination parameters.

Valid `sort` keys are:
`label`
`artist`
`title`
`catno`
`format`
`rating`
`year`
`added`

Valid `sort_order` keys are:
`asc`
`desc`

**Parameters**
- username `string` (required) Example: shooezgirl
The username of the submissions you are trying to fetch.

- sort `string` (optional) Example: artist
Sort items by this field (see below for all valid sort keys.

- sort_order `string` (optional) Example: desc
Sort items in a particular order (asc or desc)

**Response  200**
Body

```
{
  "contributions": [
    {
      "artists": [
        {
          "anv": "",
          "id": 17961,
          "join": "And",
          "name": "Cher",
          "resource_url": "https://api.discogs.com/artists/17961",
          "role": "",
          "tracks": ""
        }
      ],
      "community": {
        "contributors": [
          {
            "resource_url": "https://api.discogs.com/users/shooezgirl",
            "username": "shooezgirl"
          },
          {
            "resource_url": "https://api.discogs.com/users/Aquacrash_Dj",
            "username": "Aquacrash_Dj"
          }
        ],
        "data_quality": "Needs Vote",
        "have": 0,
        "rating": {
          "average": 0,
          "count": 0
        },
        "status": "Accepted",
        "submitter": {
          "resource_url": "https://api.discogs.com/users/shooezgirl",
          "username": "shooezgirl"
        },
        "want": 0
      },
      "companies": [],
      "country": "US",
      "data_quality": "Needs Vote",
      "date_added": "2014-03-25T15:16:13-07:00",
      "date_changed": "2014-05-14T13:36:00-07:00",
      "estimated_weight": 60,
      "format_quantity": 1,
      "formats": [
        {
          "descriptions": [
            "7\"",
            "45 RPM",
            "Single",
            "Promo"
          ],
          "name": "Vinyl",
          "qty": "1"
        }
      ],
      "genres": [
        "Rock",
        "Pop"
      ],
      "id": 5531933,
      "images": [
        {
          "height": 605,
          "resource_url": "https://api-img.discogs.com/e0X1tNZv6nkdOOTPJAn-dtCbFa0=/fit-in/600x605/filters:strip_icc():format(jpeg):mode_rgb():quality(96)/discogs-images/R-5531933-1400099758-6444.jpeg.jpg",
          "type": "primary",
          "uri": "https://api-img.discogs.com/e0X1tNZv6nkdOOTPJAn-dtCbFa0=/fit-in/600x605/filters:strip_icc():format(jpeg):mode_rgb():quality(96)/discogs-images/R-5531933-1400099758-6444.jpeg.jpg",
          "uri150": "https://api-img.discogs.com/8et0xtf9REFloKoqi6NSJ6AJvFI=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb()/discogs-images/R-5531933-1400099758-6444.jpeg.jpg",
          "width": 600
        }
      ],
      "labels": [
        {
          "catno": "7-27529-DJ",
          "entity_type": "1",
          "id": 821,
          "name": "Geffen Records",
          "resource_url": "https://api.discogs.com/labels/821"
        }
      ],
      "master_id": 223379,
      "master_url": "https://api.discogs.com/masters/223379",
      "notes": "Promotion Not For Sale",
      "released": "1989",
      "released_formatted": "1989",
      "resource_url": "https://api.discogs.com/releases/5531933",
      "series": [],
      "status": "Accepted",
      "styles": [
        "Ballad"
      ],
      "thumb": "https://api-img.discogs.com/8et0xtf9REFloKoqi6NSJ6AJvFI=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb()/discogs-images/R-5531933-1400099758-6444.jpeg.jpg",
      "title": "After All",
      "uri": "https://www.discogs.com/Cher-And-Peter-Cetera-After-All/release/5531933",
      "videos": [
        {
          "description": "Cher & Peter Cetera - After All [On-Screen Lyrics]",
          "duration": 247,
          "embed": true,
          "title": "Cher & Peter Cetera - After All [On-Screen Lyrics]",
          "uri": "https://www.youtube.com/watch?v=qy717J3Iscw"
        }
      ],
      "year": 1989
    }
  ],
  "pagination": {
    "items": 2,
    "page": 1,
    "pages": 1,
    "per_page": 50,
    "urls": {}
  }
}
```

## User Collection 
### Collection 
The Collection resource allows you to view and manage a user’s collection.
A collection is arranged into folders. Every user has two permanent folders already:

ID `0`, the “All” folder, which cannot have releases added to it, and
ID `1`, the “Uncategorized” folder.

Because it’s possible to own more than one copy of a release, each with its own notes, grading, and so on, each instance of a release in a folder has an instance ID.

Through the Discogs website, users can create custom notes fields. There is not yet an API method for creating and deleting these fields, but they can be listed, and the values of the fields on any instance can be modified.

```
GET /users/{username}/collection/folders
```

Retrieve a list of folders in a user’s collection. If the collection has been made private by its owner, authentication as the collection owner is required. If you are not authenticated as the collection owner, only folder ID 0 (the “All” folder) will be visible (if the requested user’s collection is public).

**Parameters**
- username `string` (required) Example: rodneyfool
The username of the collection you are trying to retrieve.

**Response  200**
Headers
```
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Cache-Control: public, must-revalidate
Content-Type: application/json
Server: lighttpd
Content-Length: 132
Date: Wed, 16 Jul 2014 23:20:21 GMT
X-Varnish: 1722533701
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```
Body
```
{
  "folders": [
    {
      "id": 0,
      "count": 23,
      "name": "All",
      "resource_url": "https://api.discogs.com/users/example/collection/folders/0"
    },
    {
      "id": 1,
      "count": 20,
      "name": "Uncategorized",
      "resource_url": "https://api.discogs.com/users/example/collection/folders/1"
    }
  ]
}
```

```
POST /users/{username}/collection/folders
```

Create a new folder in a user’s collection.

Authentication as the collection owner is required.

**Parameters**
- username `string` (required) Example: rodneyfool
The username of the collection you are trying to retrieve.

- name `string` (optional) Example: My favorites
The name of the newly-created folder.

**Response  201**
Body
```
{
  "id": 232842,
  "name": "My Music",
  "count": 0,
  "resource_url": "https://api.discogs.com/users/example/collection/folders/232842"
}
```

### Collection Folder 

```
GET /users/{username}/collection/folders/{folder_id}
```

Retrieve metadata about a folder in a user’s collection.

If `folder_id` is not `0`, authentication as the collection owner is required.

**Parameters**
- username `string` (required) Example: rodneyfool
The username of the collection you are trying to request.

- folder_id `number` (required) Example: 3
The ID of the folder to request.

**Response  200**
Headers
```
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Cache-Control: public, must-revalidate
Content-Type: application/json
Server: lighttpd
Content-Length: 132
Date: Wed, 16 Jul 2014 23:20:21 GMT
X-Varnish: 1722533701
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```
Body
```
{
  "id": 1,
  "count": 20,
  "name": "Uncategorized",
  "resource_url": "https://api.discogs.com/users/example/collection/folders/1"
}
```

```
POST /users/{username}/collection/folders/{folder_id}
```

Edit a folder’s metadata. Folders 0 and 1 cannot be renamed.

Authentication as the collection owner is required.

**Parameters**
- username `string` (required) Example: rodneyfool
The username of the collection you are trying to modify.

- folder_id `number` (required) Example: 3
The ID of the folder to modify.

**Response  200**
Body
```
{
  "id": 392,
  "count": 3,
  "name": "An Example Folder",
  "resource_url": "https://api.discogs.com/users/example/collection/folders/392"
}
```

```
DELETE /users/{username}/collection/folders/{folder_id}
```

Delete a folder from a user’s collection. A folder must be empty before it can be deleted.

Authentication as the collection owner is required.

**Parameters**
- username `string` (required) Example: rodneyfool
The username of the collection you are trying to modify.

- folder_id `number` (required) Example: 3
The ID of the folder to delete.

**Response  204**

### Collection Items By Release 

```
GET /users/{username}/collection/releases/{release_id}
```

View the user’s collection folders which contain a specified release. This will also show information about each release instance.

The `release_id` must be non-zero.

Authentication as the collection owner is required if the owner’s collection is private.

**Parameters**
- username `string` (required) Example: susan.salkeld
The username of the collection you are trying to view.

- release_id `number` (required) Example: 7781525
The ID of the release to request.

**Response  200**
Body
```
{
  "pagination": {
    "per_page": 50,
    "items": 28,
    "page": 1,
    "urls": {},
    "pages": 1
  },
  "releases": [
    {
      "instance_id": 148842233,
      "rating": 4,
      "basic_information": {
        "labels": [
          {
            "name": "Varèse Vintage",
            "entity_type": "1",
            "catno": "302 067 361 1",
            "resource_url": "http://api.discogs.com/labels/151979",
            "id": 151979,
            "entity_type_name": "Label"
          }
        ],
        "formats": [
          {
            "descriptions": [
              "LP",
              "Album",
              "Compilation",
              "Limited Edition",
              "Mono"
            ],
            "name": "Vinyl",
            "qty": "2"
          }
        ],
        "thumb": "http://api-img.discogs.com/FGmDbZ6M9wNPwEAsn0yWz1jzQuI=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb():quality(90)/discogs-images/R-7781525-1449187274-5587.jpeg.jpg",
        "title": "The BBC Radio Sessions",
        "artists": [
          {
            "join": ",",
            "name": "The Zombies",
            "anv": "",
            "tracks": "",
            "role": "",
            "resource_url": "http://api.discogs.com/artists/262221",
            "id": 262221
          }
        ],
        "resource_url": "http://api.discogs.com/releases/7781525",
        "year": 2015,
        "id": 7781525
      },
      "folder_id": 688739,
      "date_added": "2015-11-30T10:54:13-08:00",
      "id": 7781525
    },
    {
      "instance_id": 181301430,
      "rating": 4,
      "basic_information": {
        "labels": [
          {
            "name": "Varèse Vintage",
            "entity_type": "1",
            "catno": "302 067 361 1",
            "resource_url": "http://api.discogs.com/labels/151979",
            "id": 151979,
            "entity_type_name": "Label"
          }
        ],
        "formats": [
          {
            "descriptions": [
              "LP",
              "Album",
              "Compilation",
              "Limited Edition",
              "Mono"
            ],
            "name": "Vinyl",
            "qty": "2"
          }
        ],
        "thumb": "http://api-img.discogs.com/FGmDbZ6M9wNPwEAsn0yWz1jzQuI=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb():quality(90)/discogs-images/R-7781525-1449187274-5587.jpeg.jpg",
        "title": "The BBC Radio Sessions",
        "artists": [
          {
            "join": ",",
            "name": "The Zombies",
            "anv": "",
            "tracks": "",
            "role": "",
            "resource_url": "http://api.discogs.com/artists/262221",
            "id": 262221
          }
        ],
        "resource_url": "http://api.discogs.com/releases/7781525",
        "year": 2015,
        "id": 7781525
      },
      "folder_id": 1,
      "date_added": "2016-07-27T08:11:29-07:00",
      "id": 7781525
    },
    {
      "instance_id": 181301442,
      "rating": 4,
      "basic_information": {
        "labels": [
          {
            "name": "Varèse Vintage",
            "entity_type": "1",
            "catno": "302 067 361 1",
            "resource_url": "http://api.discogs.com/labels/151979",
            "id": 151979,
            "entity_type_name": "Label"
          }
        ],
        "formats": [
          {
            "descriptions": [
              "LP",
              "Album",
              "Compilation",
              "Limited Edition",
              "Mono"
            ],
            "name": "Vinyl",
            "qty": "2"
          }
        ],
        "thumb": "http://api-img.discogs.com/FGmDbZ6M9wNPwEAsn0yWz1jzQuI=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb():quality(90)/discogs-images/R-7781525-1449187274-5587.jpeg.jpg",
        "title": "The BBC Radio Sessions",
        "artists": [
          {
            "join": ",",
            "name": "The Zombies",
            "anv": "",
            "tracks": "",
            "role": "",
            "resource_url": "http://api.discogs.com/artists/262221",
            "id": 262221
          }
        ],
        "resource_url": "http://api.discogs.com/releases/7781525",
        "year": 2015,
        "id": 7781525
      },
      "folder_id": 688739,
      "date_added": "2016-07-27T08:11:35-07:00",
      "id": 7781525
    }
  ]
}
```

### Collection Items By Folder 

```
GET /users/{username}/collection/folders/{folder_id}/releases
```

Returns the list of item in a folder in a user’s collection. Accepts Pagination parameters.

Basic information about each release is provided, suitable for display in a list. For detailed information, make another API call to fetch the corresponding release.

If `folder_id` is not `0`, or the collection has been made private by its owner, authentication as the collection owner is required.

If you are not authenticated as the collection owner, only public notes fields will be visible.

Valid `sort` keys are:
`label`
`artist`
`title`
`catno`
`format`
`rating`
`added`
`year`

**Parameters**
- username `string` (required) Example: rodneyfool
The username of the collection you are trying to request.

- folder_id `number` (required) Example: 3
The ID of the folder to request.

- sort `string` (optional) Example: artist
Sort items by this field (see below for all valid sort keys.

- sort_order `string` (optional) Example: desc
Sort items in a particular order (asc or desc)

**Response  200**
Headers
```
Reproxy-Status: yes
Access-Control-Allow-Origin: *
Cache-Control: public, must-revalidate
Content-Type: application/json
Server: lighttpd
Content-Length: 132
Date: Wed, 16 Jul 2014 23:20:21 GMT
X-Varnish: 1722533701
Age: 0
Via: 1.1 varnish
Connection: keep-alive
```
Body
```
{
  "pagination": {
    "per_page": 1,
    "pages": 14,
    "page": 1,
    "items": 14,
    "urls": {
      "next": "https://api.discogs.com/users/example/collection/folders/1/releases?page=2&per_page=1",
      "last": "https://api.discogs.com/users/example/collection/folders/1/releases?page=2&per_page=14",
    }
  },
  "releases": [
    {
      "id": 2464521,
      "instance_id": 1,
      "folder_id": 1,
      "rating": 0,
      "basic_information": {
        "id": 2464521,
        "title": "Information Chase",
        "year": 2006,
        "resource_url": "https://api.discogs.com/releases/2464521",
        "thumb": "https://api-img.discogs.com/vzpYq4_kc52GZFs14c0SCJ0ZE84=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb()/discogs-images/R-2464521-1285519861.jpeg.jpg",
        "cover_image": "https://api-img.discogs.com/vzpYq4_kc52GZFs14c0SCJ0ZE84/fit-in/500x500/filters:strip_icc():format(jpeg):mode_rgb():quality(90)/discogs-images/R-2464521-1285519861.jpeg.jpg",
        "formats": [
          {
            "qty": "1",
            "descriptions": [ "Mini", "EP" ],
            "name": "CDr"
          }
        ],
        "labels": [
          {
            "resource_url": "https://api.discogs.com/labels/11647",
            "entity_type": "",
            "catno": "8BP059",
            "id": 11647,
            "name": "8bitpeoples"
          }
        ],
        "artists": [
          {
            "id": 103906,
            "name": "Bit Shifter",
            "join": "",
            "resource_url": "https://api.discogs.com/artists/103906",
            "anv": "",
            "tracks": "",
            "role": ""
          }
        ],
        "genres": [
            "Electronic", 
            "Pop"
        ],
        "styles": [
            "Chiptune"
        ]
      },
      "notes": [
        {
          "field_id": 3,
          "value": "bleep bloop blorp."
        }
      ]
    }
  ]
}
```

### Add To Collection Folder 

```
POST /users/{username}/collection/folders/{folder_id}/releases/{release_id}
```

Add a release to a folder in a user’s collection.

The `folder_id` must be non-zero – you can use 1 for “Uncategorized”.

Authentication as the collection owner is required.

**Parameters**
- username `string` (required) Example: rodneyfool
The username of the collection you are trying to modify.

- folder_id `number` (required) Example: 3
The ID of the folder to modify.

- release_id `number` (required) Example: 130076
The ID of the release you are adding.

**Response  201**
Body
```
{
  "instance_id": 3,
  "resource_url": "https://api.discogs.com/users/example/collection/folders/1/release/1/instance/3"
}
```

### Change Rating Of Release 

```
POST /users/{username}/collection/folders/{folder_id}/releases/{release_id}/instances/{instance_id}
```

Change the rating on a release and/or move the instance to another folder.

This endpoint potentially takes 2 folder_id parameters: one in the URL (which is the folder you are requesting, and is required), and one in the request body (representing the folder you want to move the instance to, which is optional)

Authentication as the collection owner is required.

**Parameters**
- username `string` (required) Example: rodneyfool
The username of the collection you are trying to modify.

- folder_id `number` (optional) Example: 4
The ID of the folder to modify (this parameter is set in the request body and is set if you want to move the instance to this folder).

- release_id `number` (required) Example: 130076
The ID of the release you are modifying.

- instance_id `number` (required) Example: 1
The ID of the instance.

- rating `number` (optional) Example: 5
The rating of the instance you are supplying.

**Response  204**

### Delete Instance From Folder 

```
DELETE /users/{username}/collection/folders/{folder_id}/releases/{release_id}/instances/{instance_id}
```

Remove an instance of a release from a user’s collection folder.

To move the release to the “Uncategorized” folder instead, use the POST method.

Authentication as the collection owner is required.

**Parameters**
- username `string` (required) Example: rodneyfool
The username of the collection you are trying to modify.

- folder_id `number` (required) Example: 3
The ID of the folder to modify.

- release_id `number` (required) Example: 130076
The ID of the release you are modifying.

- instance_id `number` (required) Example: 1
The ID of the instance.

**Response  204**
**Response  403**
Body
```
{
  "message": "You don't have permission to access this resource."
}
```

### List Custom Fields 

```
GET /users/{username}/collection/fields
```

Retrieve a list of user-defined collection notes fields. These fields are available on every release in the collection.

If the collection has been made private by its owner, authentication as the collection owner is required.

If you are not authenticated as the collection owner, only fields with public set to true will be visible.

**Parameters**
- username `string` (required) Example: rodneyfool
The username of the collection you are trying to modify.

**Response  200**
Body
```
{
  "fields": [
    {
      "name": "Media",
      "options": [
        "Mint (M)",
        "Near Mint (NM or M-)",
        "Very Good Plus (VG+)",
        "Very Good (VG)",
        "Good Plus (G+)",
        "Good (G)",
        "Fair (F)",
        "Poor (P)"
      ],
      "id": 1,
      "position": 1,
      "type": "dropdown",
      "public": true
    },
    {
      "name": "Sleeve",
      "options": [
        "Generic",
        "No Cover",
        "Mint (M)",
        "Near Mint (NM or M-)",
        "Very Good Plus (VG+)",
        "Very Good (VG)",
        "Good Plus (G+)",
        "Good (G)",
        "Fair (F)",
        "Poor (P)"
      ],
      "id": 2,
      "position": 2,
      "type": "dropdown",
      "public": true
    },
    {
      "name": "Notes",
      "lines": 3,
      "id": 3,
      "position": 3,
      "type": "textarea",
      "public": true
    }
  ]
}
```

### Edit Fields Instance 

```
POST /users/{username}/collection/folders/{folder_id}/releases/{release_id}/instances/{instance_id}/fields/{field_id}{?value}
```

Change the value of a notes field on a particular instance.

**Parameters**
- username `string` (required) Example: rodneyfool
The username of the collection you are trying to modify.

- value `string` (required) Example: foo
The new value of the field. If the field’s type is dropdown, the value must match one of the values in the field’s list of options.

- folder_id `number` (required) Example: 3
The ID of the folder to modify.

- release_id `number` (required) Example: 130076
The ID of the release you are modifying.

- instance_id `number` (required) Example: 1
The ID of the instance.

- field_id `number` (required) Example: 8
The ID of the field.

**Response  204**

### Collection Value 

```
GET /users/{username}/collection/value
```

Returns the minimum, median, and maximum value of a user’s collection.

Authentication as the collection owner is required.

**Parameters**
- username `string` (required) Example: rodneyfool
The username of the collection you are trying to modify.

**Response  200**
Body
```
{
  "maximum": "$250.00",
  "median": "$100.25",
  "minimum": "$75.50"
}
```

## User Wantlist 
### Wantlist 
The Wantlist resource allows you to view and manage a user’s wantlist.

```
GET /users/{username}/wants
```

Returns the list of releases in a user’s wantlist. Accepts Pagination parameters.

Basic information about each release is provided, suitable for display in a list. For detailed information, make another API call to fetch the corresponding release.

If the wantlist has been made private by its owner, you must be authenticated as the owner to view it.

The `notes` field will be visible if you are authenticated as the wantlist owner.

**Parameters**
- username `string` (required) Example: rodneyfool
The username of the wantlist you are trying to fetch.

**Response  200**
Body
```
{
  "pagination": {
    "per_page": 50,
    "pages": 1,
    "page": 1,
    "items": 2,
    "urls": {}
  },
  "wants": [
    {
      "rating": 4,
      "basic_information": {
        "formats": [
          {
            "text": "Digipak",
            "qty": "1",
            "descriptions": [
              "Album"
            ],
            "name": "CD"
          }
        ],
        "thumb": "https://api-img.discogs.com/PsLAcp_I0-EPPkSBaHx2t7dmXTg=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb()/discogs-images/R-1867708-1248886216.jpeg.jpg",
        "cover_image": "https://api-img.discogs.com/PsLAcp_I0-EPPkSBaHx2t7dmXTg=/fit-in/500x500/filters:strip_icc():format(jpeg):mode_rgb()/discogs-images/R-1867708-1248886216.jpeg.jpg",
        "title": "Year Zero",
        "labels": [
          {
            "resource_url": "https://api.discogs.com/labels/2311",
            "entity_type": "",
            "catno": "B0008764-02",
            "id": 2311,
            "name": "Interscope Records"
          }
        ],
        "year": 2007,
        "artists": [
          {
            "join": "",
            "name": "Nine Inch Nails",
            "anv": "",
            "tracks": "",
            "role": "",
            "resource_url": "https://api.discogs.com/artists/3857",
            "id": 3857
          }
        ],
        "resource_url": "https://api.discogs.com/releases/1867708",
        "id": 1867708
      },
      "resource_url": "https://api.discogs.com/users/example/wants/1867708",
      "id": 1867708
    },
    {
      "rating": 0,
      "basic_information": {
        "formats": [
          {
            "qty": "1",
            "descriptions": [
              "Album"
            ],
            "name": "CDr"
          }
        ],
        "thumb": "https://api-img.discogs.com/w1cVy7ppMYEDlqY9sjoAojC3MhQ=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb()/discogs-images/R-1675174-1236118359.jpeg.jpg",
        "cover_image": "https://api-img.discogs.com/w1cVy7ppMYEDlqY9sjoAojC3MhQ=/fit-in/347x352/filters:strip_icc():format(jpeg):mode_rgb()/discogs-images/R-1675174-1236118359.jpeg.jpg",
        "title": "Dawn Metropolis",
        "labels": [
          {
            "resource_url": "https://api.discogs.com/labels/141550",
            "entity_type": "",
            "catno": "NORM007",
            "id": 141550,
            "name": "Normative"
          }
        ],
        "year": 2009,
        "artists": [
          {
            "join": "",
            "name": "Anamanaguchi",
            "anv": "",
            "tracks": "",
            "role": "",
            "resource_url": "https://api.discogs.com/artists/667233",
            "id": 667233
          }
        ],
        "resource_url": "https://api.discogs.com/releases/1675174",
        "id": 1675174
      },
      "notes": "Sample notes.",
      "resource_url": "https://api.discogs.com/users/example/wants/1675174",
      "id": 1675174
    }
  ]
}
```

### Add To Wantlist 
Add a release to a user’s wantlist.

Authentication as the wantlist owner is required.

```
PUT /users/{username}/wants/{release_id}{?notes,rating}
```

**Parameters**
- username `string` (required) Example: rodneyfool
The username of the wantlist you are trying to fetch.

- release_id `number` (required) Example: 130076
The ID of the release you are adding.

- notes `string` (optional) Example: My favorite release
User notes to associate with this release.

- rating `number` (optional) Example: 5
User’s rating of this release, from 0 (unrated) to 5 (best). Defaults to 0.

**Response  201**
Body
```
{
  "id": 1,
  "rating": 0,
  "notes": "",
  "resource_url": "https://api.discogs.com/users/example/wants/1",
  "basic_information": {
    "id": 1,
    "resource_url": "https://api.discogs.com/releases/1",
    "thumb": "https://api-img.discogs.com/7HGTQzTb7os1duruukQElELEapk=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb()/discogs-images/R-1-1193812031.jpeg.jpg",
    "cover_image": "https://api-img.discogs.com/7HGTQzTb7os1duruukQElELEapk=/fit-in/347x352/filters:strip_icc():format(jpeg):mode_rgb()/discogs-images/R-1-1193812031.jpeg.jpg",
    "title": "Stockholm",
    "year": 1999,
    "formats": [
      {
        "qty": "2",
        "descriptions": [
          "12\""
        ],
        "name": "Vinyl"
      }
    ],
    "labels": [
      {
        "name": "Svek",
        "entity_type": "1",
        "catno": "SK032",
        "resource_url": "https://api.discogs.com/labels/5",
        "id": 5,
        "entity_type_name": "Label"
      }
    ],
    "artists": [
      {
        "join": "",
        "name": "Persuader, The",
        "anv": "",
        "tracks": "",
        "role": "",
        "resource_url": "https://api.discogs.com/artists/1",
        "id": 1
      }
    ]
  }
}
```

```
POST /users/{username}/wants/{release_id}{?notes,rating}ç
```

**Parameters**
- username `string` (required) Example: rodneyfool
The username of the wantlist you are trying to fetch.

- notes `string` (optional) Example: My favorite release
User notes to associate with this release.

- rating `number` (optional) Example: 5
User’s rating of this release, from 0 (unrated) to 5 (best). Defaults to 0.

**Response  200**
Body
```
{
  "id": 1,
  "rating": 0,
  "notes": "I've added some notes!",
  "resource_url": "https://api.discogs.com/users/example/wants/1",
  "basic_information": {
    "id": 1,
    "resource_url": "https://api.discogs.com/releases/1",
    "thumb": "https://api-img.discogs.com/7HGTQzTb7os1duruukQElELEapk=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb()/discogs-images/R-1-1193812031.jpeg.jpg",
    "cover_image": "https://api-img.discogs.com/7HGTQzTb7os1duruukQElELEapk=/fit-in/500x500/filters:strip_icc():format(jpeg):mode_rgb()/discogs-images/R-1-1193812031.jpeg.jpg",
    "title": "Stockholm",
    "year": 1999,
    "formats": [
      {
        "qty": "2",
        "descriptions": [
          "12\""
        ],
        "name": "Vinyl"
      }
    ],
    "labels": [
      {
        "name": "Svek",
        "entity_type": "1",
        "catno": "SK032",
        "resource_url": "https://api.discogs.com/labels/5",
        "id": 5,
        "entity_type_name": "Label"
      }
    ],
    "artists": [
      {
        "join": "",
        "name": "Persuader, The",
        "anv": "",
        "tracks": "",
        "role": "",
        "resource_url": "https://api.discogs.com/artists/1",
        "id": 1
      }
    ]
  }
}
```

```
DELETE /users/{username}/wants/{release_id}{?notes,rating}
```

**Parameters**
- username `string` (required) Example: rodneyfool
The username of the wantlist you are trying to fetch.

**Response  204**

## User Lists 
### User Lists 
The List resource allows you to view a User’s Lists.

```
GET /users/{username}/lists
```

Returns a User’s Lists. Private Lists will only display when authenticated as the owner. Accepts Pagination parameters.

**Parameters**
- username `string` (required) Example: rodneyfool
The username of the Lists you are trying to fetch.

**Response  200**
Body
```
{
  "pagination": {
    "per_page": 1,
    "items": 2,
    "page": 2,
    "urls": {
      "prev": "https://api.discogs.com/users/rodneyfool/lists?per_page=1&page=1",
      "first": "https://api.discogs.com/users/rodneyfool/lists?per_page=1&page=1"
    },
    "pages": 2
  },
  "lists": [
    {
      "date_added": "2016-05-31T10:30:51-07:00",
      "date_changed": "2016-05-31T10:30:51-07:00",
      "name": "rodneyfool",
      "id": 1,
      "uri": "https://www.discogs.com/lists/rodneyfool/1",
      "resource_url": "https://api.discogs.com/lists/1",
      "public": false,
      "description": "foo test description"
    }
  ]
}
```

### List 

```
GET /lists/{list_id}
```

Returns items from a specified List. Private Lists will only display when authenticated as the owner.

**Parameters**
- list_id `string` (required) Example: 123
The ID of the List you are trying to fetch.

**Response  200**
Body
```
{
  "created_ts": "2016-05-31T10:36:30-07:00",
  "modified_ts": "2016-05-31T13:46:12-07:00",
  "name": "new list",
  "list_id": 2,
  "url": "https://www.discogs.com/lists/new-list/2",
  "items": [
    {
      "comment": "My list comment",
      "display_title": "Silent Phase - The Rewired Mixes",
      "uri": "https://www.discogs.com/Silent-Phase-The-Rewired-Mixes/release/4674",
      "image_url": "https://api-img.discogs.com/-06gF81ykx-Ok1PCpNR7B7Rt_Dc=/300x300/smart/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/A-3227-1132807172.jpeg.jpg",
      "resource_url": "https://api.discogs.com/releases/4674",
      "type": "release",
      "id": 4674
    },
    {
      "comment": "item comment",
      "display_title": "Various - Artificial Intelligence II",
      "uri": "https://www.discogs.com/Various-Artificial-Intelligence-II/release/2964",
      "image_url": "http://api-img.discogs.com/euixsynJwQxJelre_kQNV-ZtX0Y=/fit-in/300x300/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/R-2964-1215444984.jpeg.jpg",
      "resource_url": "https://api.discogs.com/releases/2964",
      "type": "release",
      "id": 2964
    },
    {
      "comment": "This is an artist",
      "display_title": "Silent Phase",
      "uri": "https://www.discogs.com/artist/3227-Silent-Phase",
      "image_url": "http://api-img.discogs.com/-06gF81ykx-Ok1PCpNR7B7Rt_Dc=/300x300/smart/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/A-3227-1132807172.jpeg.jpg",
      "resource_url": "https://api.discogs.com/artists/3227",
      "type": "artist",
      "id": 3227
    }
  ],
  "resource_url": "https://api.discogs.com/lists/2",
  "public": false,
  "description": "What a cool list!"
}
```
