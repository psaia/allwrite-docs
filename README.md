# eggplant

Write and organize your documentation using Google Drive.

### Why?

1. Write and organize the hierarchical structure of your documentation using Google Drive. This gives you the benefit of using its wysiwyg and user permissions. Not everyone in an organization wants to write markdown in GitHub.
2. Get a nice API out of the box based on that Drive structure. Use a pre-made theme or make your own.

### Steps

1. Authenticate with Google Drive.
2. Build out your documentation.
3. Install or create a theme using the API this module creates.

### API

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
