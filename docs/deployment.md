## Deployment

Estrutura responsável por atualizar a imagem base dos pods em execução.

Hierarquia de estruturas:
`Deployment -> ReplicaSet -> Pod`

Quando aplicamos um arquivo de configuração do tipo `deployment` ele será
responsável por criar um novo `replicaset` e atualizar todos os `pods` existentes.
Esse processo acontece ser gerar downtime de serviço, pois o próprio kubernets
gerencia esse procedimento atualizando pequenos grupos de `pods` até aplicar
a nova configuração para todos.

### Voltando versão da imagem

Para listar o histórico de versões de `deployment` utilize o seguinte comando:
`kubectl rollout history deployment [DEPLOYMENT_NAME]`

Para voltar uma versão
`kubectl rollout undo deployment [DEPLOYMENT_NAME]`

Para voltar para uma versão específica
`kubectl rollout undo deployment [DEPLOYMENT_NAME] --to-revision=[REVISION_NUMBER]`

