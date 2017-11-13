# API Responses

### GET /menu

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

### GET /page(/:slug)

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

### GET /?q=my+search

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

### Page not found

If a page is not found, an error will be returned with error code `404`.

```json
{
  "code":400,
  "result":null,
  "error":"not found"
}
```
