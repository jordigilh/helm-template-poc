# helm-template-poc

This repository contains a Proof of Concept that uses the Helm template engine to render files outside the _templates_ subdirectory inside the chart. Helm expects files contained in the _templates_ subdirectory to be in yaml, tpl or json format, therefore restricting the use of the engine to these formats. This PoC validates that it's possible to reuse the helm code to enable processing of files in other formats and the chart still maintain compatibility with helm by storing these files in the _files_ subdirectory.

The process involves overwriting the values in the `chart.Template` field to contain files located in the _files/konveyor_ subdirectory before calling the `helm.Render` function, so that the engine processes these files despite the fact that they are located outside _templates_ and are not in the json, yaml or tpl formats.
