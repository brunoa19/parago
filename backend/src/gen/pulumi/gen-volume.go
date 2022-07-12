package pulumi

import (
	"fmt"
	"strings"

	"shipa-gen/src/shipa"
)

func GenerateVolume(cfg shipa.VolumeConfig) *shipa.Result {
	header := genMain()

	var content string
	if hasVolume(cfg) {
		content = genVolume(cfg)
	}

	if len(content) == 0 {
		return nil
	}

	return &shipa.Result{
		Filename: "index.ts",
		Header:   header,
		Content:  content,
	}
}

func hasVolume(cfg shipa.VolumeConfig) bool {
	return cfg.Name != "" && cfg.Capacity != "" && cfg.Plan != ""
}

func genVolume(cfg shipa.VolumeConfig) string {
	const indent = "   "
	fields := []string{
		fmt.Sprintf(`%s name: "%s"`, indent, cfg.Name),
		fmt.Sprintf(`%s capacity: "%s"`, indent, cfg.Capacity),
		fmt.Sprintf(`%s plan: "%s"`, indent, cfg.Plan),
	}

	if cfg.AccessModes != "" {
		fields = append(fields, fmt.Sprintf(`%s accessModes: "%s"`, indent, cfg.AccessModes))
	}

	if opts := genOpts(cfg.Opts); opts != "" {
		fields = append(fields, opts)
	}

	return fmt.Sprintf(`
const %s = new shipa.Volume("%s", {
%s
}%s);
`, genVolumeVarName(cfg.Name), genVolumeResourceName(cfg.Name), strings.Join(fields, ",\n"), genDepends(cfg.DependsOn, genVolumeVarName))
}

func genOpts(opts *shipa.VolumeOptions) string {
	if opts == nil {
		return ""
	}

	const indent = "       "
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

	return fmt.Sprintf(`    opts: {
%s
    }`, strings.Join(fields, ",\n"))
}

func genVolumeVarName(name string) string {
	return genVarName("volume", name)
}

func genVolumeResourceName(name string) string {
	return genResourceName("volume-", name)
}
