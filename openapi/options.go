package openapi

import (
	"fmt"
	"path"
)

type Options struct {
	Filename  Filename
	FileDir   string
	Version   Version
	Format    Format
	MergeOnly bool
}

func (o Options) fileAndPath() string {
	return path.Join(o.fileDir(), o.filename().String())
}

func (o Options) filename() Filename {
	if o.Filename != "" {
		return o.Filename
	}

	return Filename(fmt.Sprintf("%s%s", o.version().FileBasename(), o.format().Ext()))
}

func (o Options) fileDir() string {
	if o.FileDir != "" {
		return o.FileDir
	}

	return "."
}

func (o Options) version() Version {
	if o.Version != "" {
		return o.Version
	}

	if version := o.Filename.Version(); version != VersionInvalid {
		return version
	}

	return V3_0_0
}

func (o Options) format() Format {
	if o.Format != "" {
		return o.Format
	}

	if format := o.Filename.Format(); format != FormatInvalid {
		return format
	}

	return FormatYAML
}

func (o Options) filenameNotProvided() bool {
	return o.Filename == ""
}

func (o Options) versionNotProvided() bool {
	return o.Version == ""
}

func (o Options) formatNotProvided() bool {
	return o.Format == ""
}
