package terraform

import (
	"fmt"
	"strings"

	"shipa-gen/src/shipa"

	"github.com/iancoleman/strcase"
)

func GenerateVolume(cfg shipa.VolumeConfig) *shipa.Result {
	header := genMain()

	content := genVolume(cfg)

	if len(content) == 0 {
		return nil
	}

	return &shipa.Result{
		Filename: "main.tf",
		Header:   header,
		Content:  content,
	}
}

func genVolume(cfg shipa.VolumeConfig) string {
	return fmt.Sprintf(`
# Set volume
resource "shipa_volume" "%s" {
%s
%s
}
`, genVolumeName(cfg.Name), genVolumeParams(cfg), genDepends(cfg.DependsOn, genVolumeName))
}

func genVolumeName(name string) string {
	return "shipa_volume." + strcase.ToSnake(name)
}

func genVolumeParams(cfg shipa.VolumeConfig) string {
	const indent = " "
	out := []string{
		fmt.Sprintf(`%s name = "%s"`, indent, cfg.Name),
		fmt.Sprintf(`%s capacity = "%s"`, indent, cfg.Capacity),
		fmt.Sprintf(`%s plan = "%s"`, indent, cfg.Plan),
	}

	if cfg.AccessModes != "" {
		out = append(out, fmt.Sprintf(`%s access_modes = "%s"`, indent, cfg.AccessModes))
	}

	opts := genVolumeOpts(cfg)
	if opts != "" {
		out = append(out, opts)
	}

	return strings.Join(out, "\n")
}

func genVolumeOpts(cfg shipa.VolumeConfig) string {
	if cfg.Opts == nil {
		return ""
	}

	const indent = "   "
	var out []string
	if cfg.Opts.Prop1 != "" {
		out = append(out, fmt.Sprintf(`%s additional_prop_1 = "%s"`, indent, cfg.Opts.Prop1))
	}
	if cfg.Opts.Prop2 != "" {
		out = append(out, fmt.Sprintf(`%s additional_prop_2 = "%s"`, indent, cfg.Opts.Prop2))
	}
	if cfg.Opts.Prop3 != "" {
		out = append(out, fmt.Sprintf(`%s additional_prop_3 = "%s"`, indent, cfg.Opts.Prop3))
	}

	if len(out) == 0 {
		return ""
	}

	return fmt.Sprintf(`  opts {
%s
  }`, strings.Join(out, "\n"))
}
