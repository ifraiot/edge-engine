## Edge-Engine


### Example Commnad

Create Service
```json
[
  {
    "command": "create_service",
    "sender": "a92dd5f7-4cef-490c-8256-59d8390feb18",
    "parameters": {
      "from": "contrainer",
      "env": [
        {
          "name": "LOG_LEVEL",
          "value": "info"
        }
      ]
    }
  }
]
```

List Services
```json
[
  {
    "command": "list_service",
    "sender": "a92dd5f7-4cef-490c-8256-59d8390feb18"
  }
]
```