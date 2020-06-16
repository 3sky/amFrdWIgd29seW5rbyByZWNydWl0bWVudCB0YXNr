# Developer recruitment task - Golang

## Description

Let’s say GogoApps is building a small application enabling users to retrieve information
about the weather in the places of their choosing. Your task is creating a microservice
responsible for fetching current weather conditions in cities specified in the requests.
Specification:

- As a source of the weather information you should use a free API described [here][1].

- Service should expose one HTTP endpoint that takes a list of city names as a
  query parameter and returns information about current weather in each city.

- Since free tier account of the OpenWeather API has limited number of API
  calls, the service has to have some kind of caching layer that would
  prevent subsequent calls for the same city in short time interval.

- The application has to expose some mechanism of configuration.
  An option to specify the HTTP port of the server and an API
  key is a minimum.

Nice to have:

- Provide a dockerfile that can be used to build and run the
  application without the need of having the Go toolchain installed.

Code must be deployed onto some remote Git repository. Preferably Github,
Bitbucket, Gitlab.
Name of the repository must be as follows:
base64 of (name + last name + “recruitment task”)
"jakub wolynko recruitment task”

[1]: https://openweathermap.org/current

## How to

- Get name

    ```bash
    echo -n 'jakub wolynko recruitment task' | base64
    > amFrdWIgd29seW5rbyByZWNydWl0bWVudCB0YXNr
    ```

- Go mod

    ```bash
    go mod init github.com/3sky/amFrdWIgd29seW5rbyByZWNydWl0bWVudCB0YXNr
    ```

- Push initial image

    ```bash
    gcloud auth configure-docker
    docker images
    docker tag gogoapp:latest gcr.io/tokyo-baton-256120/weather_app:0.0.1
    docker push gcr.io/tokyo-baton-256120/weather_app:0.0.1
    ```

- Build Cloud Run Infra

    ```bash
    cd infra
    terraform init
    terraform apply -var 'authfile=../../../../auth.json'
    ```

- Commit new tag to run CI/CD pipeline
- APPID is GitHub Secret
    
