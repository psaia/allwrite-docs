# Allwrite Docs

Writing in Google Drive is enjoyable. When something is enjoyable and accessible, we tend to do it more often and with better quality.

This API connects with your Google Drive and provides RESTful endpoints which return the pages within Drive in a organized and usable format. With this API, beautiful (or ugly) user interfaces can be created and reused anywhere you need to display documentation online.

This is not a SaaS product and 100% open source. You just need to host it yourself.

# Table of Contents

* [How it works](#how-it-works)
* [Installation](#installation)
* [Formatting](#formatting)
* [API](#api)

## Workflow

Give authors access to the activated folder using the typical means of doing so within Drive. Then authors can create pages and folders within the activated folder. Pages can have children infinitely deep by using directories. Pages and folders are both made public and ordered by using a special format in their page name.

```
|n| Page Title
```

The number between the two pipes (`n`) is what the menu is sorted by. It can be any number. If a page or directory does not have this format, it will not be public. This is useful when writing pages that should not yet be public (aka drafts).

Lastly, if zero is used (`|0|`), this is considered to be the landing page of the sub directory. If there is no directory, it is used as the default response, or "homepage". If a `|0|` is not provided for the root or sub directory, the `slug` property will be false. This is common when there isn't a landing page, only sub pages.

## Installation

TODO:

* The Go App
* Docker image
* Other quick and easy installation routes.

## Formatting

Formatting within the document is done by using Google's wysiwyg, as usual. Allwrite then translates the content to well formatted html and markdown. You may use which ever format you like.

Other tips:

* Format code (`<code>`) by using the "Courier" font.
* Indent a block of text for a `<blockquote>`.
* Unordered and ordered lists are treated as such.
* Colors will have no effect.
* Headers will be treated as such (`<h1>`, `<h2>`, and `<h3>`).

## API

#### GET /menu 200

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

#### GET /page(/:slug) 200

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

If a page is not found, an error will be returned with error code `404`.

```json
{
  "status": 404,
  "message": "We're sorry, but the page you were looking for could not be found."
}
```
