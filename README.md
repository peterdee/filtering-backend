## filtering-backend

This backend application is a [Brille module](https://github.com/julyskies/brille) demo

Frontend application code is available at https://github.com/peterdee/filtering-frontend

DEV: http://localhost:8910

PRODUCTION: https://brille-demo.onrender.com

### Deploy

Minimal required Golang version: **v1.18**

```shell script
git clone https://github.com/peterdee/filtering-backend
cd ./filtering-backend
gvm use 1.18
go mod download
```

### Environment variables

The `.env` file is required unless this application is launched on Render

See [.env.example](.env.example) for details

### Launch

```shell script
go run ./
```

Required to be used with [AIR](https://github.com/cosmtrek/air)

### Render deployment

This application (`release` branch) is automatically deployed to [Render](https://render.com)

Available at https://brille-demo.onrender.com

### License

[MIT](./LICENSE.md)
