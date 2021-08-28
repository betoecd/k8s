## Notes

Init a cluster
`kind create cluster` (1)

Para acessar o cluster criado, é necessário informar o contexto
`kubectl cluster-info --context kind-kind` (2)

Obs: As credenciais e certificas ficam armazenas dentro do arquivo
`~/.kube/config`
Esse arquivo pode armezenar informações para diferentes clusters, e para
manipular cada cluster será necessário informar o context.

Ao executar o comando (2) já podemos manipular o cluster k8s.
Utilizando o comando `docker ps` é possível identificar um container em execução,
justamente o controle do cluster.
Outra forma, mais eficaz, de obter informações do cluster é através do
commando `kubectl get nodes`, ele retorna informações de todos os nodes em execução.

Para recuperar apenas clusters:
`kind get clusters`

Para deletar um cluster:
`kind delete clusters [NAME]`

Parâmetros extras para criar um cluster:
 - --config=[FILE]
 - --name=[CLUSTER-NAME]
`kind create cluster --config=config/kind.yaml --name=bits`

Para recuperar informações de todos os clusters configurados no `~/.kube/config`
`kubectl config get-clusters`

Para alterar de cluster
`kubectl config use-context [CLUSTER-NAME]`
