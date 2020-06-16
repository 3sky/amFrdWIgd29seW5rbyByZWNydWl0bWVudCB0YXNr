locals {
  auth_file = file(var.path)
}

provider "google" {
 credentials = local.auth_file
 project     = var.project
}

resource "google_cloud_run_service" "default" {
  name     = "weather-app"
  location = "europe-west1"

  template {
    spec {
      containers {
        image = "gcr.io/tokyo-baton-256120/weather_app:0.0.1"
        env {
          name = "APP_PORT"
          value = "8080"
        }
        env {
          name = "APPID"
          value = "439d4b804bc8187953eb36d2a8c26a02"
        }
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}
data "google_iam_policy" "noauth" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers",
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "noauth" {
  location    = google_cloud_run_service.default.location
  project     = google_cloud_run_service.default.project
  service     = google_cloud_run_service.default.name

  policy_data = data.google_iam_policy.noauth.policy_data
}
