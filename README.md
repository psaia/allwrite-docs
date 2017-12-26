# Allwrite Docs

An incredibly fast documentation API powered by Google Drive written purely in Go.

This API connects with your Google Drive and provides RESTful endpoints which return the pages within Drive in a organized and usable format. With this API, beautiful (or ugly) user interfaces can be created and reused anywhere you need to display documentation online.

**Features:**

* Let anyone, technical or not, contribute to documentation.
* Full-text search.
* Auto-generating SSL via [certbot](https://certbot.eff.org/).
* Transforms Google docs to clean markdown and html.
* Images are directly referenced from Google so you don't need to worry about image storage.
* No dependencies other than Postgres.
* URL structure builds itself based on the directory structure.
* Did I mention that it's crazy fast since pages are pre-cached?

![screenshot](/docs/screenshot.png)

# Table of Contents

* [Workflow](#workflow)
* [Formatting](#formatting)
* [Examples](#examples)
* [Themes](#themes)
* [Installation](#installation)
* [API](#api)
* [Contributing](#contributing)
* [CLI](#cli)

## Workflow

Give authors access to the activated folder using the typical means of doing so
within Drive. Then authors can create pages and folders within the activated
folder. Pages can have children infinitely deep by using directories. Pages and
folders are both made public and ordered by using a special format in their page
name.

```
|n| Page Title
```

The number between the two pipes (`n`) is what the menu is sorted by. It can be
any number. If a page or directory does not have this format, it will not be
public. This is useful when writing pages that should not yet be public (aka
drafts).

Lastly, if zero is used (`|0|`), this is considered to be the landing page of
the sub directory. If there is no directory, it is used as the default response,
or "homepage". If a `|0|` is not provided for the root or sub directory, the
`slug` property will be false. This is common when there isn't a landing page,
only sub pages.

## Formatting

Formatting within the document is done by using Google's wysiwyg, as usual. Allwrite then translates the content to well formatted html and markdown. Both the html and markdown formats are returned from the API so you may use which ever works best for you.

Formatting guide:

* Images just work. Plus they're hosted by Google for free!
* Unordered and ordered lists are treated as such.
* Colors, text alignment, and other frills will have no effect. Create an issue if you have a suggestion.
* Headers will be treated as such (`<h1>`, `<h2>`, and `<h3>`).
* Format code as you normally would (3 backticks followed by the language).

## Examples

* [stackahoy.io](https://stackahoy.io) is going to use it.

## Themes

* [Spartan](https://github.com/LevInteractive/spartan-allwrite/)
* Make your own!

## Installation

First, you should generate a OAuth 2.0 json file [here](https://console.developers.google.com/projectselector/apis/credentials). Select
"other" for Application Type then place the client_secret.json file on the
server you'll be running the API.

Head to your server and run the following to **install or update** Allwrite:

```bash
curl -L https://github.com/LevInteractive/allwrite-docs/blob/master/install.sh?raw=true | sh
```

If it's you're installing for the first time import the postgres schema:

```bash
curl -O https://raw.githubusercontent.com/LevInteractive/allwrite-docs/master/store/postgres/sql/pages.sql
psql < pages.sql
```

Finally, try to run the server in the foreground:

```bash
# Download the environmental variables. These need to available to the
# user/shell so allwrite can connect.
curl https://raw.githubusercontent.com/LevInteractive/allwrite-docs/master/creds.example.sh > creds

# Configure. Make sure these variables are correct.
vim creds

# Load the variables.
source creds
```

Start the server.

```bash
# Run the server. You'll eventually want to run this in the background and use
# something like nginx to create a reverse proxy.
allwrite s
```

Once you confirmed that it works, setup something like supervisord to run it for
you. You can see an example configuration file for supervisord [here](/docs/supervisord.md)

## API

See response examples [here](/docs/api.md).

## Contributing

See docs for development [here](/docs/development.md).

## CLI

After installing, you'll have access to the CLI

```bash
$ allwrite-docs
```
