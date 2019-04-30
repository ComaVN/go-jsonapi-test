# go-jsonapi-test
Example code for go jsonapi library issue

```
$ go run main.go | jq .
{
  "data": {
    "type": "foo",
    "id": "1",
    "attributes": {
      "someattr": "a"
    },
    "relationships": {
      "rel": {
        "data": {
          "type": "bar",
          "id": "2"
        }
      },
      "relrio": {
        "data": {
          "type": "bar",
          "id": "3"
        }
      }
    }
  }
}
{
  "data": {
    "type": "foo",
    "id": "1",
    "attributes": {
      "someattr": "a"
    },
    "relationships": {
      "rel": {
        "data": {
          "type": "bar",
          "id": "2"
        }
      },
      "relrio": {
        "data": {
          "type": "bar",
          "id": "3"
        }
      }
    }
  },
  "included": [
    {
      "type": "bar",
      "id": "3"
    },
    {
      "type": "bar",
      "id": "2",
      "attributes": {
        "someattr": "b"
      }
    }
  ]
}
```
