# Golang Weather App

## How implementation proces looks

- Go mod

    ```bash
    go mod init github.com/3sky/amFrdWIgd29seW5rbyByZWNydWl0bWVudCB0YXNr
    ```

- Push initial image

    ```bash
    docker build -t weather_app .
    gcloud auth configure-docker
    docker images
    docker tag weather_app:latest gcr.io/tokyo-baton-256120/weather_app:0.0.1
    docker push gcr.io/tokyo-baton-256120/weather_app:0.0.1
    ```

- Build Cloud Run Infra

    ```bash
    cd infra
    terraform init
    terraform apply -var 'authfile=../../../../auth.json'
    ```

- Commit new tag to run CI/CD pipeline

    ```bash
    # for docker build
    # everything without tag started by `v`
    git push
    # for deploy to Cloud Run
    # tag started by `v`
    git push origin v0.0.1
    ```

- APPID is set as  environment variable in Cloud Run manifest
