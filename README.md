# Allwrite Docs

Writing in Google Drive is enjoyable. When something is enjoyable and accessible, we tend to do it more often and with better quality.

This API connects with your Google Drive and provides RESTful endpoints which return the pages within Drive in a organized and usable format. With this API, beautiful (or ugly) user interfaces can be created and reused anywhere you need to display documentation online.

This is not a SaaS product and 100% open source. You just need to host it yourself.

# Table of Contents

* [How it works](#how-it-works)
* [Installation](#installation)
* [Formatting](#formatting)
* [API](#api)
* [Examples](#examples)

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

## Installation

TODO:

* The Go App
* Docker image
* Other quick and easy installation methods.

## Formatting

Formatting within the document is done by using Google's wysiwyg, as usual. Allwrite then translates the content to well formatted html and markdown. Both the html and markdown formats are returned from the API so you may use which ever works best for you.

Formatting guide:

* Format code (`<code>`) by using the "Courier" font.
* Indent a block of text for a `<blockquote>`.
* Unordered and ordered lists are treated as such.
* Colors, text alignment, and other frills will have no effect. Create an issue if you have a suggestion.
* Headers will be treated as such (`<h1>`, `<h2>`, and `<h3>`).

## API

There are only two endpoints!

#### GET /menu 

Returns a collection of page fragments.

```json
[
  {
    "name": "Getting Start",
    "slug": "getting-started",
    "updated": 1500057521,
    "children": []
  },
  {
    "name": "Configure",
    "slug": false,
    "updated": 1500057521,
    "children": [
      {
        "name": "Hello World",
        "slug": "configure/hello-world",
        "updated": 1500057521,
        "children": []
      }
    ]
  }
]
```

#### GET /page(/:slug) 

Pull a page based on its slug. If not provided, `|0|` page will be used. Page fragments will be included as children if they exist.

```json
{
  "name": "Configure",
  "slug": "configure",
  "updated": 1500057521,
  "html": "<html here>",
  "md": "<markdown here>",
  "children": [
    {
      "name": "Hello World",
      "slug": "configure/hello-world",
      "updated": 1500057521,
      "children": []
     }
  ]
}
```

#### Page not found

If a page is not found, an error will be returned with error code `404`.

```json
{
  "status": 404,
  "message": "We're sorry, but the page you were looking for could not be found."
}
```

## Examples

Coming soon.
