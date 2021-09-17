package consts

const (
	DefaultETCDTimeoutSeconds              = 5
	DefaultETCDDialKeepaliveTimeSeconds    = 30
	DefaultETCDDialKeepaliveTimeoutSeconds = 10

	AppNameMaxLength     = 20
	AppCompNameMaxLength = 20
	AppCompMaxReplicas   = 10

	YataiDebugImg             = "yatai.io/yatai-infras/debug"
	YataiKubectlNamespace     = "default"
	YataiKubectlContainerName = "main"
	YataiKubectlImage         = "yatai.io/yatai-infras/k8s"

	TracingContextKey = "tracing-context"
	// nolint: gosec
	YataiApiTokenHeaderName = "X-YATAI-API-TOKEN"

	BentoServicePort = 5000

	NoneStr = "None"
)