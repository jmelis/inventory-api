Possible example to define relations in a resource-type/repoter-type yaml:

```yaml
relation_properties:
- relation: is_fedramp
  expr: {{ .acm.fedramp }}
  subject: user:*
- relation: has_addon_strimzi
  expr: {{ "strimzi" in .acm.addons }}
  subject: plan:{{.acm.subscription.plan}}
```

