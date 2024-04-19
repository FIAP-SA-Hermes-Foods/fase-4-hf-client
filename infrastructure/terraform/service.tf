resource "kubernetes_service" "hf-client-service" { 
    metadata { 
        name = "hf-client-service"
        namespace = "dev"
    }

    spec { 
        type = "LoadBalancer"
        selector = { 
            app = "hf-client-go-app"
        }
        port { 
            protocol = "TCP"
            name = "hf-client-http-port"
            port = 8080
            target_port = 8080
        }
        port { 
            protocol = "TCP"
            name = "hf-client-rpc-port"
            port = 8070
            target_port = 8070
        }
    }
}
