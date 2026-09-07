package api

import (
	"strings"

	"github.com/NETWAYS/alertmanager-icinga-bridge/internal/icinga2"
)

type AlertFieldGetter interface {
	GetField(a *Alert, es icinga2.ExitStatus) string
}

type AnnotationsGetter []string

func (ag AnnotationsGetter) GetField(a *Alert, _ icinga2.ExitStatus) string {
	for _, key := range ag {
		if annotation, ok := a.Annotations[key]; ok {
			return annotation
		}
	}

	return ""
}

type LabelsGetter []string

func (lg LabelsGetter) GetField(a *Alert, _ icinga2.ExitStatus) string {
	for _, key := range lg {
		if label, ok := a.Labels[key]; ok {
			return label
		}
	}

	return ""
}

type AnnotationsPrefixGetter []string

func (apg AnnotationsPrefixGetter) GetField(a *Alert, es icinga2.ExitStatus) string {
	suffix := "_" + strings.ToLower(es.String())

	for _, key := range apg {
		if annotation, ok := a.Annotations[key+suffix]; ok {
			return annotation
		} else if annotation, ok := a.Annotations[key]; ok {
			return annotation
		}
	}

	return ""
}
