resource "kubernetes_horizontal_pod_autoscaler_v2" "hf_client_hpa" {
  metadata {
    name      = "hf-client-hpa"
    namespace = "dev"
  }

  spec {
    scale_target_ref {
      api_version = "apps/v1"
      kind        = "Deployment"
      name        = "hf-client-deployment"
    }

    min_replicas = 1
    max_replicas = 2

  behavior {
      scale_up {
        stabilization_window_seconds = 60
        select_policy = "Max"  # Add this line
        policy {
          type          = "Percent"
          value         = 10
          period_seconds = 60
        }
      }
      scale_down {
        stabilization_window_seconds = 120
        select_policy = "Max"  # Add this line
        policy {
          type          = "Percent"
          value         = 10
          period_seconds = 120
        }
      }
    }

    metric {
      type = "Resource"
      resource {
        name = "cpu"
        target {
          type               = "Utilization"
          average_utilization = 70
        }
      }
    }
  }
}
