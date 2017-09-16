# eggplant

Writing in Google Drive is enjoyable. When something is more enjoyable, you tend to do it more often and with better quality.

This API connects with your Google Drive and provides endpoints which returns the folders and pages within Drive in a organized and usable format. With these endpoints, beautiful (or not) user interfaces can be created.

This is not a SaaS product and 100% open source. You just need to host it yourself!

# Table of Contents

* [Steps](#steps)
* [API](#api)

## Steps

1. Authenticate with Google Drive.
2. Build out your documentation.
3. Install or create a theme using the API this module creates.

## API

`GET /menu` Returns an array of page fragments.

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
        "children": false
      }
    ]
  }
]
```

`GET /page(/:slug)` Pull a page based on its slug. If not provided, `|0|` page will be used.

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
      "children": false
     }
  ]
}
```
