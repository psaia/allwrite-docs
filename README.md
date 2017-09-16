# eggplant

Write and organize your documentation using Google Drive.

### Why?

1. Write and organize the hierarchical structure of your documentation using Google Drive. This gives you the benefit of using its wysiwyg and user permissions.
2. Get a nice API out of the box based on that Drive structure. Use a pre-made theme or make your own.

### Steps

1. Authenticate with Google Drive.
2. Build out your documentation.
3. Install or create a theme using the API this module creates.

### API

* `GET /menu` Returns a full menu based on the Drive directory.
* `GET /page(/:slug)` Pull a page based on its slug. If not provided, `|0|` page will be used.
