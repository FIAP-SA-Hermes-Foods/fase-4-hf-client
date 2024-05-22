resource "kubernetes_deployment" "hf_client_deployment" {
  metadata {
    name      = "hf-client-deployment"
    namespace = "dev"
  }

  spec {
    selector {
      match_labels = {
        app = "hf-client-go-app"
      }
    }
    template {
      metadata {
        labels = {
          app = "hf-client-go-app"
        }
      }
      spec {
        container {
          name            = "hf-client-go-app-http"
          image           = "localhost:5000/app-go-http:latest"
          image_pull_policy = "IfNotPresent"
          resources {
            limits = {
              cpu    = "500m"
              memory = "1Gi"
            }
            requests = {
              cpu    = "250m"
              memory = "512Mi"
            }
          }
          port {
            container_port = 8080
          }
        }

        container {
          name            = "hf-client-go-app-rpc"
          image           = "localhost:5000/app-go-rpc:latest"
          image_pull_policy = "IfNotPresent"
          resources {
            limits = {
              cpu    = "500m"
              memory = "1Gi"
            }
            requests = {
              cpu    = "250m"
              memory = "512Mi"
            }
          }
        }
      } 
    }
  }
}
