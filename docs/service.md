## Service

Realize o trabalho de service discover e load balancer entre seus pods.
O `service` para a idenficação dos pods que compõe um serviço pelo nome do app
descrito na inicialização de cada pod.

### Tipos de service

:construction:

### Port e TargetPort

* port: define qual será a porta que estabelecerá conexão com os pods.
* targetPort: quando omitido assume o mesmo valor de `port`, e indica dentro do
container em execução qual é a porta do serviço ou funcionalidade que deve ser
buscada.
