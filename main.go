package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/engine"
)

const (
	konveyorDirectoryName = "konveyor"
)

var (
	chartFullPath       = filepath.Join("chart", "foo")
	konveyorPartialPath = filepath.Join("files", konveyorDirectoryName)
)

func main() {

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	e := engine.Engine{}
	l, err := loader.Loader(filepath.Join(pwd, chartFullPath))
	if err != nil {
		log.Fatal(err)
	}
	c, err := l.Load()
	if err != nil {
		log.Fatal(err)
	}
	options := chartutil.ReleaseOptions{
		Name:      c.Name(),
		Namespace: "default",
		Revision:  1,
		IsInstall: false,
		IsUpgrade: false,
	}
	valuesToRender, err := chartutil.ToRenderValues(c, c.Values, options, chartutil.DefaultCapabilities.Copy())
	if err != nil {
		log.Fatal(err)
	}

	rendered, err := e.Render(c, valuesToRender)
	if err != nil {
		log.Fatalf("Failed to render templates for chart %s:%s", c.Name(), err)
	}
	fmt.Printf("templates rendered in templates directory: %d\n", len(rendered))
	c2 := *c
	c2.Templates = filterTemplatesByPath(konveyorPartialPath, c2.Files)
	rendered, err = e.Render(&c2, valuesToRender)
	if err != nil {
		log.Fatalf("Failed to render templates for chart %s:%s", c.Name(), err)
	}
	fmt.Printf("templates rendered in files/konveyor directory: %d\n", len(rendered))
	for k, v := range rendered {
		fmt.Printf("%s\n%v\n", k, v)
	}

}

func filterTemplatesByPath(pathPrefix string, files []*chart.File) []*chart.File {
	ret := []*chart.File{}
	for _, f := range files {
		if strings.HasPrefix(f.Name, pathPrefix) {
			ret = append(ret, f)
		}
	}
	return ret
}
