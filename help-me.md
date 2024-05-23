# Setup

**IMPORTANTE**
Caso tenha já algum cluster com as mesmas funções que serão criados certifique-se de limpar seu ambiente: `kind delete cluster`

Também vale lembrar que se você já setou em algum momento uma docker network para o seu kind é melhor excluir para evitar erros. Sugerimos apagar a imagem do registry e as imagens que já *buildou*.  

Caso queira um fluxo pra se orientar:
`Pare todos os clusters -> Pare todos os containers -> apague a network kind -> apague as imagens`

### 1.1 Build do cluster com .yaml

Precisamos criar um arquivo **YAML** com o seguinte conteúdo:
```yaml
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
name: hf-cluster
containerdConfigPatches:
  - |-
    [plugins."io.containerd.grpc.v1.cri".registry.mirrors."localhost:5000"]
      endpoint = ["http://registry:5000"]
nodes:
  - role: control-plane
    kubeadmConfigPatches:
      - |-
        kind: ClusterConfiguration
        apiVersion: kubeadm.k8s.io/v1beta2
        controllerManager:
          extraArgs:
            cloud-provider: external
        featureGates:
          CloudControllerManager: true
    extraPortMappings:
      - containerPort: 30080
        hostPort: 30080
        protocol: TCP
      - containerPort: 30070
        hostPort: 30070
        protocol: TCP
  - role: worker


```

**VERIFIQUE O PASSO 1.2 ANTES DE PROSSEGUIR**

Ao fechar o seu editor de texto vá ate o terminal e digite `kind create cluster --config kind-config.yaml` 

**Obs: Certifique-se de que o seu arquivo de configuração chama-se `kind-config.yaml`, senão substitua pelo nome que deu.**

**Obs2: Não precisa passar a flag --name pois esta referenciada no YAML**

Verifique seu cluster com o comando:
```sh
kubectl cluster-info --context kind-hf-cluster
kubectl get nodes
```

### 1.2 Instale o Cloud Provider KIND

Siga essas instruções para instalar: 
1. Acesse a página de lançamentos do projeto: [baixe aqui](https://github.com/kubernetes-sigs/cloud-provider-kind/releases)
2. Baixe o binário correspondente ao seu sistema operacional
3. Descompacte o arquivo e mova o binário `cloud-provider-kind` para um diretório no seu PATH (por exemplo, `/usr/local/bin`).
4. Volte ao 1.1 e configure seu cluster, ou se quiser so *rebuildar* o cluster também que da na mesma :D

*Help*: [Doc oficial do Cloud Provider](https://github.com/kubernetes-sigs/cloud-provider-kind)
