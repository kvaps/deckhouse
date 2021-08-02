---
# [<en>] section containing the parameters of instance class for worker nodes
# [<en>] version of the Deckhouse API
# [<ru>] секция, описывающая параметры инстанс-класса для узлов c компонентами, обеспечивающими рабочую нагрузку
# [<ru>] используемая версия API Deckhouse
apiVersion: deckhouse.io/v1
kind: GCPInstanceClass
metadata:
  # [<en>] name of instance class
  # [<ru>] имя инстанс класса
  name: worker
spec:
  diskSizeGb: 40
  # [<en>] Machine type in use for this instance class
  # [<ru>] используемый Machine type для данного инстанс-класса
  # [<en>] you might consider changing this
  # [<ru>] возможно, захотите изменить
  machineType: n2-standard-4
---
apiVersion: deckhouse.io/v1
kind: NodeGroup
metadata:
  name: worker
spec:
  cloudInstances:
    classReference:
      kind: GCPInstanceClass
      name: worker
    # [<en>] the minimum number of instances for the group in each zone
    # [<ru>] минимальное количество инстансов в зоне
    minPerZone: 1
    # [<en>] the maximum number of instances for the group in each zone
    # [<ru>] максимальное количество инстансов в зоне
    maxPerZone: 1
    # [<en>] you might consider changing this
    # [<ru>] возможно, захотите изменить
    zones:
    - europe-west3-a
  disruptions:
    approvalMode: Automatic
  nodeTemplate:
    # [<en>] similar to the standard metadata.labels field
    # [<ru>] аналогично стандартному полю metadata.labels
    labels:
      node-role.deckhouse.io/worker: ""
  nodeType: CloudEphemeral
---
# [<en>] section containing the parameters of nginx ingress controller
# [<en>] version of the Deckhouse API
# [<ru>] секция, описывающая параметры nginx ingress controller
# [<ru>] используемая версия API Deckhouse
apiVersion: deckhouse.io/v1
kind: IngressNginxController
metadata:
  name: nginx
spec:
  # [<en>] the name of the Ingress class to use with the Ingress nginx controller
  # [<ru>] имя Ingress-класса для обслуживания Ingress NGINX controller
  ingressClass: nginx
  # [<en>] the way traffic goes to cluster from the outer network
  # [<ru>] способ поступления трафика из внешнего мира
  inlet: LoadBalancer
  # [<en>] describes on which nodes the component will be located
  # [<ru>] описывает, на каких узлах будет находиться компонент
  nodeSelector:
    node-role.deckhouse.io/worker: ""
---
apiVersion: deckhouse.io/v1
kind: ClusterAuthorizationRule
metadata:
  name: admin
spec:
  # [<ru>] список учётных записей Kubernetes RBAC
  # [<en>] Kubernetes RBAC accounts list
  subjects:
  - kind: User
    name: admin@example.com
  # [<ru>] предустановленный шаблон уровня доступа
  # [<en>] pre-defined access template
  accessLevel: SuperAdmin
  # [<ru>] разрешить пользователю делать kubectl port-forward
  # [<en>] allow user to do kubectl port-forward
  portForwarding: true
---
# [<en>] section containing the parameters of the static user
# [<ru>] секция, описывающая параметры статического пользователя
# [<en>] version of the Deckhouse API
# [<ru>] используемая версия API Deckhouse
apiVersion: deckhouse.io/v1
kind: User
metadata:
  name: admin
spec:
  # [<en>] user e-mail
  # [<ru>] e-mail пользователя
  email: admin@example.com
  # [<en>] this is a hash for generated password: <GENERATED_PASSWORD>
  # [<ru>] это хэш сгенерированного пароля: <GENERATED_PASSWORD>
  # [<en>] you might consider changing this
  # [<ru>] возможно, захотите изменить
  password: <GENERATED_PASSWORD_HASH>