
# SNAPFD App

## API Reference

#### You can this apis to work with app

```bash
  POST /api/v1/order - create order
```

## Deployment

To deploy this project clone it and run

```bash
  git clone https://github.com/farhadiis/snapfd-go.git
  cd snapfd-go
  sudo ./build.sh
  sudo ./run.sh
```

## Running Tests

To run tests, run the following command

```bash
  go test ./order/tests -v
  go test ./order-process/tests -v
```