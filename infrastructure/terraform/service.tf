resource "kubernetes_service" "hf-client-service" { 
    metadata { 
        name = "hf-client-service"
        namespace = "dev"
    }

    spec {
    type = "NodePort"
    selector = {
      app = "hf-client-go-app"
    }

    port {
      name        = "hf-client-http-port"
      port        = 8080
      target_port = 8080
    }

    port {
      name        = "hf-client-rpc-port"
      port        = 8070
      target_port = 8070
      node_port   = 30070
    }
  }  
}