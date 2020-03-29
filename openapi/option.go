package openapi

type Option func(Options) Options

func OptionFilename(filename string) Option {
	return func(options Options) Options {
		newOptions := options
		newOptions.Filename = Filename(filename)
		return newOptions
	}
}

func OptionDir(path string) Option {
	return func(options Options) Options {
		newOptions := options
		newOptions.FileDir = path
		return newOptions
	}
}

func OptionVersion(version Version) Option {
	return func(options Options) Options {
		newOptions := options
		newOptions.Version = version
		return newOptions
	}
}

func OptionV2_0(options Options) Options {
	return OptionVersion(V2_0)(options)
}

func OptionV3_0_0(options Options) Options {
	return OptionVersion(V3_0_0)(options)
}

func OptionFormat(format Format) Option {
	return func(options Options) Options {
		newOptions := options
		newOptions.Format = format
		return newOptions
	}
}

func OptionJSON(options Options) Options {
	return OptionFormat(FormatJSON)(options)
}

func OptionYAML(options Options) Options {
	return OptionFormat(FormatYAML)(options)
}

func OptionMergeOnly(options Options) Options {
	newOptions := options
	newOptions.MergeOnly = true
	return options
}
