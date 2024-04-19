resource "kubernetes_deployment" "hf-client-deployment" {

    metadata {  
        name = "hf-client-deployment"
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
                    name = "hf-client-go-app-http"
                    image = "${data.external.env_vars.result.image_api_http_url}:${data.external.env_vars.result.image_tag}"
                    image_pull_policy = "Always"
                    env_from { 
                        secret_ref { 
                            name = "hf-deploy-secret"
                        }
                    }
                    port { 
                        container_port = 8080
                    }
                }
                container { 
                    name = "hf-client-go-app-rpc"
                    image = "${data.external.env_vars.result.image_api_rpc_url}:${data.external.env_vars.result.image_tag}"
                    image_pull_policy = "Always"
                    env_from { 
                        secret_ref { 
                            name = "hf-deploy-secret"
                        }
                    }
                    port { 
                        container_port = 8080
                    }
                }
            image_pull_secrets { 
                    name = "hfregcred"
                }
            }
        }
    }
}

