Random notes for my short memory.

# Install cert-manager with kustomize

- Github: https://github.com/jetstack/cert-manager
- Release note: https://github.com/jetstack/cert-manager/releases
- getting started: https://docs.cert-manager.io/en/latest/getting-started/install/kubernetes.html

The following procedure is mostly based on getting started document.

## download the deployment manifests and generate Kustomization.yaml
```shell
$ mkdir -p deployments/cert-manager
$ wget https://github.com/jetstack/cert-manager/releases/download/v0.8.0/cert-manager.yaml -O ./deployments/cert-manager/cert-manager.yaml
$ cat <<-EOF > ./deployments/cert-manager/cert-manager.yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
resources:
- cert-manager.yaml
EOF
```

Do not add 'namespace: cert-manager' to `Kustomization.yaml`, otherwise you accidentally overwrite `namespace: kube-system` in `cert-manager.yaml` and experience bellow error in `pod/cert-manager-webhook-xxxxx`.
```shell
Error: configmaps "extension-apiserver-authentication" is forbidden: User "system:serviceaccount:cert-manager:cert-manager-webhook" cannot get resource "configmaps" in API group "" in the namespace "kube-system"
```

## Create a clusterrolebinding
This is a required specifically for GKE. See https://docs.cert-manager.io/en/latest/getting-started/install/kubernetes.html
TODO: make it as a descriptive deployment code instead of a command
```shell
$ kubectl create clusterrolebinding cluster-admin-binding \
  --clusterrole=cluster-admin \
  --user=$(gcloud config get-value core/account)
Your active configuration is: [<your active configuration name>]
clusterrolebinding.rbac.authorization.k8s.io/cluster-admin-binding created
```

```shell
$ kustomize build deployments/cert-manager \
    | kubectl apply -f -
namespace/cert-manager created
customresourcedefinition.apiextensions.k8s.io/certificates.certmanager.k8s.io created
customresourcedefinition.apiextensions.k8s.io/challenges.certmanager.k8s.io created
customresourcedefinition.apiextensions.k8s.io/clusterissuers.certmanager.k8s.io created
customresourcedefinition.apiextensions.k8s.io/issuers.certmanager.k8s.io created
customresourcedefinition.apiextensions.k8s.io/orders.certmanager.k8s.io created
validatingwebhookconfiguration.admissionregistration.k8s.io/cert-manager-webhook created
serviceaccount/cert-manager-cainjector created
serviceaccount/cert-manager-webhook created
serviceaccount/cert-manager created
clusterrole.rbac.authorization.k8s.io/cert-manager-edit created
clusterrole.rbac.authorization.k8s.io/cert-manager-view created
clusterrole.rbac.authorization.k8s.io/cert-manager-webhook:webhook-requester created
clusterrole.rbac.authorization.k8s.io/cert-manager-cainjector created
clusterrole.rbac.authorization.k8s.io/cert-manager created
rolebinding.rbac.authorization.k8s.io/cert-manager-webhook:webhook-authentication-reader created
clusterrolebinding.rbac.authorization.k8s.io/cert-manager-cainjector created
clusterrolebinding.rbac.authorization.k8s.io/cert-manager-webhook:auth-delegator created
clusterrolebinding.rbac.authorization.k8s.io/cert-manager created
service/cert-manager-webhook created
deployment.apps/cert-manager-cainjector created
deployment.apps/cert-manager-webhook created
deployment.apps/cert-manager created
apiservice.apiregistration.k8s.io/v1beta1.admission.certmanager.k8s.io created
certificate.certmanager.k8s.io/cert-manager-webhook-ca created
certificate.certmanager.k8s.io/cert-manager-webhook-webhook-tls created
issuer.certmanager.k8s.io/cert-manager-webhook-ca created
issuer.certmanager.k8s.io/cert-manager-webhook-selfsign created
```

## Verifying the installation

Create a ClusterIssuer to test the webhook works okay.
```shell
$ cat <<EOF > test-resources.yaml
apiVersion: v1
kind: Namespace
metadata:
  name: cert-manager-test
---
apiVersion: certmanager.k8s.io/v1alpha1
kind: Issuer
metadata:
  name: test-selfsigned
  namespace: cert-manager-test
spec:
  selfSigned: {}
---
apiVersion: certmanager.k8s.io/v1alpha1
kind: Certificate
metadata:
  name: selfsigned-cert
  namespace: cert-manager-test
spec:
  commonName: example.com
  secretName: selfsigned-cert-tls
  issuerRef:
    name: test-selfsigned
EOF
```

Create the test resources.
```
$ kubectl apply -f test-resources.yaml
namespace/cert-manager-test unchanged
issuer.certmanager.k8s.io/test-selfsigned created
certificate.certmanager.k8s.io/selfsigned-cert created
```

Check the status of the newly created certificate.
You may need to wait a few seconds before cert-manager processes the certificate request.
```
$ kubectl describe certificate -n cert-manager-test
Name:         selfsigned-cert
Namespace:    cert-manager-test
...
Spec:
  Common Name:  example.com
  Issuer Ref:
    Name:       test-selfsigned
  Secret Name:  selfsigned-cert-tls
Status:
  Conditions:
    Last Transition Time:  2019-06-08T10:30:35Z
    Message:               Certificate is up to date and has not expired
    Reason:                Ready
    Status:                True
    Type:                  Ready
  Not After:               2019-09-06T10:30:35Z
Events:
  Type    Reason      Age    From          Message
  ----    ------      ----   ----          -------
  Normal  CertIssued  3m18s  cert-manager  Certificate issued successfully
```

Clean up the test resources.
```
kubectl delete -f test-resources.yaml
namespace "cert-manager-test" deleted
issuer.certmanager.k8s.io "test-selfsigned" deleted
certificate.certmanager.k8s.io "selfsigned-cert" deleted
```

# Issue wildcard certificates using Let's Encrypt

## Understand ACME Issuers
Read:
- https://docs.cert-manager.io/en/latest/tasks/issuers/index.html
- https://docs.cert-manager.io/en/latest/tasks/issuers/setup-acme/index.html
- https://docs.cert-manager.io/en/latest/reference/issuers.html

> Let’s Encrypt does not support issuing wildcard certificates with HTTP-01 challenges. To issue wildcard certificates, you must use the DNS-01 challenge.

Also read https://letsencrypt.org/how-it-works/ and https://letsencrypt.org/docs/challenge-types/ to understand how Let's Encrypt works and what is "challenge".

You need to set up DNS providers accordingly depending on where you manage a domain.
https://docs.cert-manager.io/en/latest/tasks/issuers/setup-acme/dns01/index.html#supported-dns01-providers

Make sure clusterissuer.
```
$ kubectl get secrets letsencrypt-staging --namespace cert-manager -o json | jq .
$ kubectl describe clusterissuer letsencrypt # `Message` should be `The ACME account was registered with the ACME server`
```

# Issuing certificates
https://docs.cert-manager.io/en/latest/tasks/issuing-certificates/index.html
https://docs.cert-manager.io/en/release-0.4/tutorials/acme/dns-validation.html

Confirm issued certification.

```shell
$ kubectl describe certificate prd-gbookshelf-certificate --namespace cert-manager
...
Events:
  Type    Reason         Age    From          Message
  ----    ------         ----   ----          -------
  Normal  OrderCreated   8m13s  cert-manager  Created Order resource "prd-gbookshelf-certificate-1808351630"
  Normal  OrderComplete  4m24s  cert-manager  Order "prd-gbookshelf-certificate-1808351630" completed successfully
  Normal  CertIssued     4m24s  cert-manager  Certificate issued successfully
$ kubectl get secrets prd-gbookshelf-tls --namespace cert-manager -o json | jq -r '.data."tls.crt"' | base64 -D | openssl x509 -text
$ cfssl certinfo -domain example.com # https://github.com/cloudflare/cfssl
```

If you can not confirm `CertIssued` event, cert-manager pod is having a problem. Its logs should explain the reason why it is failing.
```shell
$ kubectl logs pod/cert-manager-xxxxxxxxxx-xxxx  --namespace cert-manager
```

# Set up an issued certificate for cloud load balancer

> An Ingress does not expose arbitrary ports or protocols. Exposing services other than HTTP and HTTPS to the internet typically uses a service of type Service.Type=NodePort or Service.Type=LoadBalancer.
> https://kubernetes.io/docs/concepts/services-networking/ingress/

Confirm cert propagation in Ingress. If Ingress still keeps temporal self signed certificate, curl returns an error.
```
$ curl -X GET https://example.com
curl: (60) SSL certificate problem: unable to get local issuer certificate
More details here: https://curl.haxx.se/docs/sslcerts.html
```

It took 30 minutes when I confirmed it. I'm not sure how long it takes time to propagate actual let's encrypt certificate to Ingress.

If you don't want to update domain upon Ingress creation every time, you can reserve a static external IP.
https://cloud.google.com/kubernetes-engine/docs/tutorials/configuring-domain-name-static-ip

# By the way, it doesn't work for gRPC

Need to take futher action to take an advantage of gRPC.
- https://cloud.google.com/load-balancing/
- https://cloud.google.com/load-balancing/docs/https/
- https://github.com/kubernetes/ingress-gce/issues/18#issuecomment-498669572

# Refs
- https://github.com/GoogleCloudPlatform/gke-managed-certs
