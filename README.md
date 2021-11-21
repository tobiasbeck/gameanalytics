# gameanalytics-go

gameanalytics-go is an inofficial sdk for gameanalytics which is written in go.

It currently can send the following events:
- Business
- Design
- Error
- Progression
- Ressource
- Session Events (user, session_end)

[gameanalytics Docs](https://gameanalytics.com/docs/s/article/Collection-API)

## Installation

Install via go get

```bash
go get tobiasbeck/gameanalytics-go
```

## Usage

```go
import gameanalytics tobiasbeck/gameanalytics-go

func main() {
  client :=  gameanalytics.NewClient(key, secret)
  ctx := context.TODO()
  client.StartSession(ctx, "123456")

}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)