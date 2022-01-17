package main

import (
	"github.com/sirupsen/logrus"
	serverVersion "k8s.io/apimachinery/pkg/util/version"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

}

// KubeClientFromDefault
func KubeClientFromDefault() (kubernetes.Interface, error) {
	// incluster
	restConf, err := rest.InClusterConfig()
	if err != nil {
		// from default kubeconfig
		restConf, err = clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
		if err != nil {
			return nil, err
		}
	}

	client, err := kubernetes.NewForConfig(restConf)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// KubeClientFromDefaultOrDie
func KubeClientFromDefaultOrDie() kubernetes.Interface {
	// incluster
	restConf, err := rest.InClusterConfig()
	if err != nil {
		// from default kubeconfig
		restConf, err = clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
		if err != nil {
			panic(err)
		}
	}

	client, err := kubernetes.NewForConfig(restConf)
	if err != nil {
		panic(err)
	}
	return client
}

// KubeClientFromConfig return kubernetes client from kubeconfig
func KubeClientFromConfig(kubeconfigPath string) (kubernetes.Interface, error) {
	conf, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		logrus.Errorf("read kubeconfig file err: %s", err.Error())
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(conf)
	if err != nil {
		logrus.Errorf("new kubernetes clientset err: %s", err.Error())
		return nil, err
	}
	return clientset, nil
}

// KubeClientFromConfigOrDie return kubernetes client from kubeconfig and panic if there is an error occurred
func KubeClientFromConfigOrDie(kubeconfigPath string) kubernetes.Interface {
	conf, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		panic(err)
	}
	return kubernetes.NewForConfigOrDie(conf)
}

func KubeClientFromParser() (*kubernetes.Clientset, error) {
	return nil, nil
}

func kubeServerVersion() (*version.Info, error) {
	clientset, err := KubeClientFromDefault()
	if err != nil {
		return nil, err
	}
	return clientset.Discovery().ServerVersion()
}

// KubeVersion
func KubeVersion() string {
	serverVersion, err := kubeServerVersion()
	if err != nil {
		return ""
	}
	return serverVersion.GitVersion
}

// IsNetWorkingV1IngressSupported check if networking.k8s.io/v1 Ingress is supported
// k8s version >= v1.19.0
func IsNetWorkingV1IngressSupported() bool {
	v, err := kubeServerVersion()
	if err != nil {
		return false
	}

	sv, err := serverVersion.ParseGeneric(v.String())
	if err != nil {
		return false
	}

	return sv.AtLeast(serverVersion.MustParseGeneric("v1.19.0"))
}

// IsNetWorkingV1beta1IngressSupported check if networking.k8s.io/v1beta1 Ingress is supported
// k8s version >= v1.14.0
func IsNetWorkingV1beta1IngressSupported() bool {
	v, err := kubeServerVersion()
	if err != nil {
		return false
	}

	sv, err := serverVersion.ParseGeneric(v.String())
	if err != nil {
		return false
	}

	return sv.AtLeast(serverVersion.MustParseGeneric("v1.14.0"))
}

var MinVersionSupported = serverVersion.MustParseGeneric("v1.18.0")

// IsKubeVersionSupported
func IsKubeVersionSupported() bool {
	v, err := kubeServerVersion()
	if err != nil {
		return false
	}
	sv, err := serverVersion.ParseGeneric(v.String())
	if err != nil {
		return false
	}

	return sv.AtLeast(MinVersionSupported)
}
