package pulumi

import (
	"fmt"
	"strings"

	"shipa-gen/src/shipa"

	"github.com/iancoleman/strcase"
)

func genVarName(prefix, name string) string {
	return prefix + strcase.ToCamel(name)
}

func genResourceName(prefix, name string) string {
	return prefix + strcase.ToKebab(name)
}

func genAppDeployVarName(name string) string {
	return genVarName("appDeploy", name)
}

func genAppDeployResourceName(name string) string {
	return genResourceName("app-deploy-", name)
}

func hasAppDeploy(cfg shipa.Config) bool {
	return cfg.AppName != "" && cfg.Image != ""
}

func genAppDeploy(cfg shipa.Config) string {
	return fmt.Sprintf(`
const %s = new shipa.AppDeploy("%s", {
    app: "%s",
%s
}%s);
`, genAppDeployVarName(cfg.AppName), genAppDeployResourceName(cfg.AppName), cfg.AppName, genAppDeployParams(cfg), genDepends(cfg.DependsOn, genAppDeployVarName))
}

func genDepends(dependsOn []string, genVar func(string) string) string {
	if dependsOn == nil {
		return ""
	}

	var deps []string
	for _, app := range dependsOn {
		deps = append(deps, genVar(app))
	}

	return fmt.Sprintf(", { dependsOn: [%s] }", strings.Join(deps, ", "))
}

func genAppDeployParams(cfg shipa.Config) string {
	const indent = "   "
	out := []string{
		fmt.Sprintf(`%s image: "%s"`, indent, cfg.Image),
		genAppDeployConfig(cfg),
	}

	if cfg.Port != 0 {
		out = append(out, genAppDeployPort(cfg))
	}

	if cfg.RegistryUser != "" && cfg.RegistrySecret != "" {
		out = append(out, genAppDeployRegistry(cfg))
	}

	if len(cfg.Volumes) > 0 {
		out = append(out, genAppDeployVolumes(cfg))
	}

	if cfg.CanarySettings != nil {
		out = append(out, genAppDeployCanarySettings(cfg))
	}

	if cfg.PodAutoScaler != nil {
		out = append(out, genAppDeployPodAutoScaler(cfg))
	}

	return strings.Join(out, ",\n")
}

func genAppDeployPodAutoScaler(cfg shipa.Config) string {
	const indent = "       "
	var out []string

	if cfg.PodAutoScaler.MaxReplicas > 0 {
		out = append(out, fmt.Sprintf(`%s max_replicas: %d`, indent, cfg.PodAutoScaler.MaxReplicas))
	}

	if cfg.PodAutoScaler.MinReplicas > 0 {
		out = append(out, fmt.Sprintf(`%s min_replicas: %d`, indent, cfg.PodAutoScaler.MinReplicas))
	}

	if cfg.PodAutoScaler.TargetCPUUtilizationPercentage > 0 {
		out = append(out, fmt.Sprintf(`%s target_cpu_utilization_percentage: %d`, indent, cfg.PodAutoScaler.TargetCPUUtilizationPercentage))
	}

	params := strings.Join(out, ",\n")
	return fmt.Sprintf(`    pod_auto_scaler: {
%s
    }`, params)
}

func genAppDeployCanarySettings(cfg shipa.Config) string {
	const indent = "       "
	var out []string

	if cfg.CanarySettings.Steps > 0 {
		out = append(out, fmt.Sprintf(`%s steps: %d`, indent, cfg.CanarySettings.Steps))
	}

	if cfg.CanarySettings.StepWeight > 0 {
		out = append(out, fmt.Sprintf(`%s step_weight: %d`, indent, cfg.CanarySettings.StepWeight))
	}

	if cfg.CanarySettings.StepInterval > 0 {
		out = append(out, fmt.Sprintf(`%s step_interval: %d`, indent, cfg.CanarySettings.StepInterval))
	}

	params := strings.Join(out, ",\n")
	return fmt.Sprintf(`    canary_settings: {
%s
    }`, params)
}

func genAppDeployConfig(cfg shipa.Config) string {
	const indent = "       "
	out := []string{
		fmt.Sprintf(`%s team: "%s"`, indent, cfg.Team),
		fmt.Sprintf(`%s framework: "%s"`, indent, cfg.Framework),
	}

	if cfg.Plan != "" {
		out = append(out, fmt.Sprintf(`%s plan: "%s"`, indent, cfg.Plan))
	}

	tags := genTags(cfg)
	if tags != "" {
		out = append(out, fmt.Sprintf(`%s %s`, indent, tags))
	}

	envs := genEnvs(cfg)
	if envs != "" {
		out = append(out, fmt.Sprintf(`%s %s`, indent, envs))
	}

	params := strings.Join(out, ",\n")
	return fmt.Sprintf(`    appConfig: {
%s
    }`, params)
}

func genAppDeployPort(cfg shipa.Config) string {
	if cfg.Port == 0 {
		return ""
	}

	return fmt.Sprintf(`    port: {
        number: %d,
        protocol: "TCP"
    }`, cfg.Port)
}

func genAppDeployRegistry(cfg shipa.Config) string {
	if cfg.RegistryUser == "" || cfg.RegistrySecret == "" {
		return ""
	}

	return fmt.Sprintf(`    registry: {
        user: "%s",
        secret: "%s"
    }`, cfg.RegistryUser, cfg.RegistrySecret)
}

func genAppDeployVolumes(cfg shipa.Config) string {
	if len(cfg.Volumes) == 0 {
		return ""
	}

	return fmt.Sprintf(`    volumes: [
%s
    ]`, genVolumes(cfg.Volumes))
}

func genVolumes(volumes []*shipa.Volume) string {
	var items []string
	for _, vol := range volumes {
		items = append(items, genAppDeployVolume(vol))
	}

	return strings.Join(items, ",\n")
}

func genAppDeployVolume(vol *shipa.Volume) string {
	return fmt.Sprintf(`        {
%s
        }`, genVolumeFields(vol))
}

func genVolumeFields(vol *shipa.Volume) string {
	const indent = "           "
	fields := []string{
		fmt.Sprintf(`%s name: "%s"`, indent, vol.Name),
		fmt.Sprintf(`%s mountPath: "%s"`, indent, vol.Path),
	}

	if opts := genVolumeOpts(vol.Opts); opts != "" {
		fields = append(fields, opts)
	}

	return strings.Join(fields, ",\n")
}

func genVolumeOpts(opts *shipa.VolumeOptions) string {
	if opts == nil {
		return ""
	}

	const indent = "               "
	var fields []string
	if opts.Prop1 != "" {
		fields = append(fields, fmt.Sprintf(`%s additionalProp1: "%s"`, indent, opts.Prop1))
	}
	if opts.Prop2 != "" {
		fields = append(fields, fmt.Sprintf(`%s additionalProp2: "%s"`, indent, opts.Prop2))
	}
	if opts.Prop3 != "" {
		fields = append(fields, fmt.Sprintf(`%s additionalProp3: "%s"`, indent, opts.Prop3))
	}

	if len(fields) == 0 {
		return ""
	}

	return fmt.Sprintf(`            mountOptions: {
%s
            }`, strings.Join(fields, ",\n"))
}

func genEnvs(cfg shipa.Config) string {
	if len(cfg.Envs) == 0 {
		return ""
	}

	var envs []string
	for _, env := range cfg.Envs {
		envs = append(envs, fmt.Sprintf("%s=%s", env.Name, env.Value))
	}

	return fmt.Sprintf(`envs: ["%s"]`, strings.Join(envs, `", "`))
}
