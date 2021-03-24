This builds because

(a) The module imports a different version of every transitive dependency, e.g., `controller-runtime`:

```
sigs.k8s.io/controller-runtime v0.5.11/go.mod h1:OTqxLuz7gVcrq+BHGUgedRu6b2VIKCEc7Pu4Jbwui0A=
sigs.k8s.io/controller-runtime v0.5.14/go.mod h1:OTqxLuz7gVcrq+BHGUgedRu6b2VIKCEc7Pu4Jbwui0A=
sigs.k8s.io/controller-runtime v0.8.3 h1:GMHvzjTmaWHQB8HadW+dIvBoJuLvZObYJ5YoZruPRao=
sigs.k8s.io/controller-runtime v0.8.3/go.mod h1:U/l+DUopBc1ecfRZ5aviA9JDmGFQKvLf5YkZNx2e0sU=
```

(b) The module, so far, does not use call any mutually incompatible function. But if it does, the module will not build.
