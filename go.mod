module kubeform.dev/provider-rediscloud-gen

go 1.16

require (
	github.com/RedisLabs/terraform-provider-rediscloud v0.0.0-00010101000000-000000000000
	k8s.io/apimachinery v0.18.3
	kubeform.dev/generator-v2 v0.0.18
)

replace github.com/RedisLabs/terraform-provider-rediscloud => github.com/kubeform/terraform-provider-rediscloud v0.2.5-0.20210730144149-20f14a5c8cdc
