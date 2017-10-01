# Allwrite Docs

An incredibly fast API server powered by Google Drive written purely in Go.

This API connects with your Google Drive and provides RESTful endpoints which return the pages within Drive in a organized and usable format. With this API, beautiful (or ugly) user interfaces can be created and reused anywhere you need to display documentation online.

This is not a SaaS product and 100% open source.

# Table of Contents

* [Workflow](#workflow)
* [Formatting](#formatting)
* [Examples](#examples)
* [Themes](#themes)
* [Installation](#installation)
* [API](#api)
* [CLI](#cli)

## Workflow

Give authors access to the activated folder using the typical means of doing so within Drive. Then authors can create pages and folders within the activated folder. Pages can have children infinitely deep by using directories. Pages and folders are both made public and ordered by using a special format in their page name.

```
|n| Page Title
```

The number between the two pipes (`n`) is what the menu is sorted by. It can be any number. If a page or directory does not have this format, it will not be public. This is useful when writing pages that should not yet be public (aka drafts).

Lastly, if zero is used (`|0|`), this is considered to be the landing page of the sub directory. If there is no directory, it is used as the default response, or "homepage". If a `|0|` is not provided for the root or sub directory, the `slug` property will be false. This is common when there isn't a landing page, only sub pages.

```
|0| Getting Started ◀────────── Accessible via "/" or "/getting-started".
|1| Who we are ◀─────────────── Accessible via "getting-started".
|2| Configuration/◀──────────── Accessible via "/configuration".
  |0| Why? ◀─────────────────── Accessible via "/configuration" or "/configuration/why".
  |1| Using the CLI ◀────────── Accessible via "/configuration/using-the-cli".
|3| The Dashboard/ ◀─────────── NOT accessible! Header only.
  |1| Mathematics ◀──────────── Accessible via "/the-dashboard/mathematics".
  |2| Machine Learning ◀─────── Accessible via "/the-dashboard/machine-learning".
  |3| A.I. ◀─────────────────── Accessible via "/the-dashboard/ai".
  About Us ◀─────────────────── Not public.
```

## Formatting

Formatting within the document is done by using Google's wysiwyg, as usual. Allwrite then translates the content to well formatted html and markdown. Both the html and markdown formats are returned from the API so you may use which ever works best for you.

Formatting guide:

* Images just work. Plus they're hosted by Google for free!
* Format code (`<code>`) by using the "Courier" font.
* Indent a block of text for a `<blockquote>`.
* Unordered and ordered lists are treated as such.
* Colors, text alignment, and other frills will have no effect. Create an issue if you have a suggestion.
* Headers will be treated as such (`<h1>`, `<h2>`, and `<h3>`).

## Examples

Coming soon.

## Themes

Coming soon.

## Installation

```bash
# Install allwrite-docs executable (/usr/local/bin/allwrite-docs).
curl -L https://github.com/LevInteractive/allwrite-docs/blob/master/install.sh?raw=true | sh

# Import database schema. Note, you'll need to have postgres setup and know what
# your username and password is.
curl -O https://raw.githubusercontent.com/LevInteractive/allwrite-docs/master/store/postgres/sql/pages.sql
psql < pages.sql

# Download the environmental variables. These need to available to the
# user/shell so allwrite can connect.
curl https://raw.githubusercontent.com/LevInteractive/allwrite-docs/master/creds.example.sh > creds

# Configure.
vim creds

# Put the vars in your shell.
source creds

# Authenticate with google Drive. You'll be prompted to click a link and copy
# and paste a code back in the terminal.
allwrite setup

# Run the server. You'll eventually want to run this in the background and use
# something like nginx to create a reverse proxy.
allwrite start
```


## API

There are three endpoints.

* /:slug
* /menu
* /?s=any+escaped+string

#### GET /menu

Returns a collection of page fragments.

```json
{
  "code":200,
  "result":[
    {
      "name":"Homepage",
      "type":"file",
      "slug":"",
      "order":0,
      "updated":"2017-09-30T23:35:31.663365Z",
      "created":"2017-09-30T23:35:31.663365Z"
    },
    {
      "name":"Only one deep",
      "type":"file",
      "slug":"another-sub-directory",
      "order":0,
      "updated":"2017-09-30T23:35:31.663365Z",
      "created":"2017-09-30T23:35:31.663365Z",
      "children":[
        {
          "name":"This is a deep file",
          "type":"file",
          "slug":"another-sub-directory/a-deeper-directory",
          "order":0,
          "updated":"2017-09-30T23:35:31.663365Z",
          "created":"2017-09-30T23:35:31.663365Z"
        }
      ]
    },
    {
      "name":"A Sub Directory",
      "type":"dir",
      "slug":"a-sub-directory",
      "order":1,
      "updated":"2017-09-30T23:35:31.663365Z",
      "created":"2017-09-30T23:35:31.663365Z",
      "children":[
        {
          "name":"How to be a friend",
          "type":"file",
          "slug":"a-sub-directory/how-to-be-a-friend",
          "order":1,
          "updated":"2017-09-30T23:35:31.663365Z",
          "created":"2017-09-30T23:35:31.663365Z"
        }
      ]
    },
    {
      "name":"Images!",
      "type":"file",
      "slug":"images",
      "order":2,
      "updated":"2017-09-30T23:35:31.663365Z",
      "created":"2017-09-30T23:35:31.663365Z"
    }
  ]
}
```

#### GET /page(/:slug)

Pull a page based on its slug. If not provided, `|0|` page will be used.

```json
{
  "code":200,
  "result":{
    "name":"Homepage",
    "type":"file",
    "slug":"",
    "order":0,
    "updated":"2017-09-30T23:21:35.044194Z",
    "created":"2017-09-30T23:21:35.044194Z",
    "doc_id":"1V5G8XmX6ggLVu09QJXqONQkLKfIix-2bMuefFYbmTmE",
    "html":"[full clean html]",
    "md":"[full clean markdown]"
  }
}
```

#### GET /?q=my+search

This will search for results based on your q parameter. The string needs to be URL encoded.

```json
{
    "code": 200,
    "result": [
        {
            "name": "This is a deep file",
            "type": "file",
            "slug": "another-sub-directory/a-deeper-directory",
            "order": 0,
            "updated": "2017-09-30T23:35:31.663365Z",
            "created": "2017-09-30T23:35:31.663365Z"
        }
    ]
}
```

#### Page not found

If a page is not found, an error will be returned with error code `404`.

```json
{
  "code":400,
  "result":null,
  "error":"not found"
}
```

## CLI

After installing, you'll have access to the API.

$ allwrite-docs

```
NAME:
   Allwrite Docs | Publish your documentation with Drive. - A new cli application

USAGE:
   allwrite-docs [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     start, s  Start the server in the foreground. This will authenticate with Google if it's the first time you're running.
     setup     Only authenticate with Google and do not run the allwrite server.
     pull, p   Pull the latest content from Google Drive.
     reset, r  Reset any saved authentication credentials for Google. You will need to re-authenticate after doing this.
     info, i   Display environmental variables. Useful for making sure everything is setup correctly.
     help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

