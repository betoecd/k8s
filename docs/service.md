# Service

Realize o trabalho de service discover e load balancer entre seus pods.
O `service` para a idenficação dos pods que compõe um serviço pelo nome do app
descrito na inicialização de cada pod.

## Tipos de service :construction:

### ClusterIP

Gera um endereço de IP interna que possibilita o acesso aos recursos criados
internamente.

### NodePort

Tipo de serviço mais arcaico quando queremos acessor um cluster k8s de fora.
Na prática será muito raro de utilizar.

### LoadBalancer

Gera um IP para acessar o seu serviço. Normalmente utilizado com provedores
em nuvem.

## Port e TargetPort

* port: define qual será a porta que estabelecerá conexão com os pods.
* targetPort: quando omitido assume o mesmo valor de `port`, e indica dentro do
container em execução qual é a porta do serviço ou funcionalidade que deve ser
buscada.
