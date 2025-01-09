# helm-template-poc

This repository contains a Proof of Concept that uses the Helm template engine to render files outside the _templates_ subdirectory inside the chart. Helm expects files contained in the _templates_ subdirectory to be in yaml, tpl or json format, therefore restricting the use of the engine to a specific format. This PoC validates that it's possible to reuse the helm code to enable processing of files located in the _files_ subdirectory by overwriting the values in the `chart.Template` field to contain files located in the _files/konveyor_ subdirectory.
