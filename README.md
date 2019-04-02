# golang_keycloak_grpc_service
Keycloak integration for IAM (Identity and ccess Management) gRPC Service to Secure Web Services/gRPC Microservices.

# Deploy Keycloak in kubernetes using Helm chart
Keycloak is an open source identity and access management for modern applications and services.

```
$ helm install stable/keycloak
```

This chart bootstraps a Keycloak StatefulSet on a Kubernetes cluster using the Helm package manager. It provisions a fully featured Keycloak installation. For more information on Keycloak and its capabilities.

The chart has an optional dependency on the PostgreSQL chart. By default, the PostgreSQL chart requires PV support on underlying infrastructure (may be disabled).

To install the chart with the release name keycloak:
```
$ helm install --name keycloak stable/keycloak
```
