package builds

import (
	"context"
	"fmt"
	"strings"

	g "github.com/onsi/ginkgo"
	o "github.com/onsi/gomega"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	e2e "k8s.io/kubernetes/test/e2e/framework"

	buildv1 "github.com/openshift/api/build/v1"
	exutil "github.com/openshift/origin/test/extended/util"
	"github.com/openshift/origin/test/extended/util/image"
)

var _ = g.Describe("[sig-builds][Feature:Builds] Optimized image builds", func() {
	defer g.GinkgoRecover()
	var (
		oc             = exutil.NewCLI("build-dockerfile-env")
		skipLayers     = buildv1.ImageOptimizationSkipLayers
		testDockerfile = fmt.Sprintf(`
FROM %s
RUN yum list installed
USER 1001
`, image.ShellImage())
	)

	g.Context("", func() {

		g.BeforeEach(func() {
			exutil.PreTestDump()
		})

		g.AfterEach(func() {
			if g.CurrentGinkgoTestDescription().Failed {
				exutil.DumpPodStates(oc)
				exutil.DumpConfigMapStates(oc)
				exutil.DumpPodLogsStartingWith("", oc)
			}
		})

		g.It("should succeed", func() {
			g.By("creating a build directly")
			build, err := oc.AdminBuildClient().BuildV1().Builds(oc.Namespace()).Create(context.Background(), &buildv1.Build{
				ObjectMeta: metav1.ObjectMeta{
					Name: "optimized",
				},
				Spec: buildv1.BuildSpec{
					CommonSpec: buildv1.CommonSpec{
						Source: buildv1.BuildSource{
							Dockerfile: &testDockerfile,
						},
						Strategy: buildv1.BuildStrategy{
							DockerStrategy: &buildv1.DockerBuildStrategy{
								ImageOptimizationPolicy: &skipLayers,
								Env: []corev1.EnvVar{
									{
										Name:  "BUILD_LOGLEVEL",
										Value: "2",
									},
								},
							},
						},
					},
				},
			}, metav1.CreateOptions{})
			o.Expect(err).NotTo(o.HaveOccurred())
			o.Expect(build.Spec.Strategy.DockerStrategy.ImageOptimizationPolicy).ToNot(o.BeNil())
			result := exutil.NewBuildResult(oc, build)
			err = exutil.WaitForBuildResult(oc.AdminBuildClient().BuildV1().Builds(oc.Namespace()), result)
			o.Expect(err).NotTo(o.HaveOccurred())
			o.Expect(result.BuildSuccess).To(o.BeTrue(), "Build did not succeed: %v", result)

			pod, err := oc.KubeClient().CoreV1().Pods(oc.Namespace()).Get(context.Background(), build.Name+"-build", metav1.GetOptions{})
			o.Expect(err).NotTo(o.HaveOccurred())
			if strings.HasSuffix(pod.Spec.Containers[0].Image, ":v3.6.0-alpha.0") {
				g.Skip(fmt.Sprintf("The currently selected builder image does not yet support optimized image builds: %s", pod.Spec.Containers[0].Image))
			}

			s, err := result.Logs()
			o.Expect(err).NotTo(o.HaveOccurred())
			o.Expect(s).To(o.ContainSubstring("Installed Packages"))
			o.Expect(s).To(o.ContainSubstring(fmt.Sprintf("\"OPENSHIFT_BUILD_NAMESPACE\"=\"%s\"", oc.Namespace())))
			o.Expect(s).To(o.ContainSubstring("Build complete, no image push requested"))
			e2e.Logf("Build logs:\n%s", result)
		})
	})
})
