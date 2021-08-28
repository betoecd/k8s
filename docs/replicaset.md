## ReplicaSet

Responsável por ditar as regras de um ou mais pods que irão compartilhar das 
mesmas diretrizes de execução. 
E quando ocorrer algum problema com algum dos pods, o replicaset atua para
disponibilizar aquele recurso novamente.

Para aplicar uma replicaset
`kubectl apply -f config/replicaset.yaml`

Para obter mais informações
`kubectl get replicasets`

Note que se deletarmos, deliberadamente um de nossos pods, a replicaset iniciará
um novo pod para ocupar o seu lugar.