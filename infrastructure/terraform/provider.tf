terraform { 
    required_providers { 
        kubernetes = { 
            source = "hashicorp/kubernetes"
        } 
    }
}

provider "kubernetes" { 
    config_path = "~/.kube/config"
}

/*
o que estava causando erro era que estavamos usando um "mycloud" ao inves de "kubernetes" na linha 3
*/