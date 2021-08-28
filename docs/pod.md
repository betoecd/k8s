## Pods

Pods são as menores unidades dentro do k8s, onde os containers serão executados.
O recomendável, é utilizar arquivos `.yaml` por são declarativos e podem ser
executados em qualquer ambiente e o kubernetes conseguirá aplicar todas as 
configuraçẽos.

As comunicações dentro do kubernetes são realizadas através de uma api, assim
sendo, cada recurso pode possuir uma versão diferente, nesses casos a documentação
sempre ajudará.

Para aplicar um arquivo `.yaml`
`kubeclt aplly -f config/pod.yaml`

Vericando o novo pod criado
`kubectl get pods`

É possível acessar um pod sem definir outras estruturas utilizando
`kubectl port-forward pod/goserver 8000:8000`

Para deletar um pod
`kubectl delete pod [NAME]'`