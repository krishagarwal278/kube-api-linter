package b

type CustomMessage struct {
	// +kubebuilder:validation:Optional
	Field string // want "field CustomMessage.Field uses marker \"kubebuilder:validation:Optional\", should use preferred marker \"k8s:optional\" instead: custom message for k8s:optional"
}
