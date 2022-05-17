/*
 * Copyright (c) 2022, Alibaba Group;
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package inject

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"io"
	"strings"

	"sigs.k8s.io/controller-tools/pkg/genall"
	"sigs.k8s.io/controller-tools/pkg/loader"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

type GeneratorInterface interface {
	RegisterMarkers(into *markers.Registry) error
	Generate(*genall.GenerationContext) error
}

var (
	enableIOCGoangAutowireMarker         = markers.Must(markers.MakeDefinition("ioc:autowire", markers.DescribesType, false))
	iocGolangAutowireTypeMarker          = markers.Must(markers.MakeDefinition("ioc:autowire:type", markers.DescribesType, ""))
	iocGolangAutowireBaseTypeMarker      = markers.Must(markers.MakeDefinition("ioc:autowire:baseType", markers.DescribesType, false))
	iocGolangAutowireInterfaceMarker     = markers.Must(markers.MakeDefinition("ioc:autowire:interface", markers.DescribesType, ""))
	iocGolangAutowireParamMarker         = markers.Must(markers.MakeDefinition("ioc:autowire:paramType", markers.DescribesType, ""))
	iocGolangAutowireParamLoaderMarker   = markers.Must(markers.MakeDefinition("ioc:autowire:paramLoader", markers.DescribesType, ""))
	iocGolangAutowireConstructFuncMarker = markers.Must(markers.MakeDefinition("ioc:autowire:constructFunc", markers.DescribesType, ""))
)

type Generator struct {
	HeaderFile string `marker:",optional"`
	Year       string `marker:",optional"`
}

func (Generator) CheckFilter() loader.NodeFilter {
	return func(node ast.Node) bool {
		// ignore interfaces
		_, isIface := node.(*ast.InterfaceType)
		return !isIface
	}
}

func (Generator) RegisterMarkers(into *markers.Registry) error {
	if err := markers.RegisterAll(into,
		enableIOCGoangAutowireMarker, iocGolangAutowireTypeMarker, iocGolangAutowireInterfaceMarker,
		iocGolangAutowireConstructFuncMarker, iocGolangAutowireParamLoaderMarker, iocGolangAutowireParamMarker, iocGolangAutowireBaseTypeMarker); err != nil {
		return err
	}
	return nil
}

func (d Generator) Generate(ctx *genall.GenerationContext) error {
	var headerText string

	if d.HeaderFile != "" {
		headerBytes, err := ctx.ReadFile(d.HeaderFile)
		if err != nil {
			return err
		}
		headerText = string(headerBytes)
	}

	objGenCtx := ObjectGenCtx{
		Collector:  ctx.Collector,
		Checker:    ctx.Checker,
		HeaderText: headerText,
	}

	for _, root := range ctx.Roots {
		outContents := objGenCtx.generateForPackage(root)
		if outContents == nil {
			continue
		}

		writeOut(ctx, root, outContents)
	}

	return nil
}

// ObjectGenCtx contains the common info for generating deepcopy implementations.
// It mostly exists so that generating for a package can be easily tested without
// requiring a full set of output rules, etc.
type ObjectGenCtx struct {
	Collector  *markers.Collector
	Checker    *loader.TypeChecker
	HeaderText string
}

// writeHeader writes out the build tag, package declaration, and imports
func writeHeader(pkg *loader.Package, out io.Writer, packageName string, imports *importsList, headerText string) {
	// NB(directxman12): blank line after build tags to distinguish them from comments
	_, err := fmt.Fprintf(out, `//go:build !ignore_autogenerated
// +build !ignore_autogenerated

%[3]s

// Code generated by ioc-go-cli

package %[1]s

import (
%[2]s
)

`, packageName, strings.Join(imports.ImportSpecs(), "\n"), headerText)
	if err != nil {
		pkg.AddError(err)
	}

}

// generateForPackage generates IOCGloang init and runtime.Object implementations for
// types in the given package, writing the formatted result to given writer.
// May return nil if source could not be generated.
func (ctx *ObjectGenCtx) generateForPackage(root *loader.Package) []byte {
	ctx.Checker.Check(root)

	root.NeedTypesInfo()

	imports := &importsList{
		byPath:  make(map[string]string),
		byAlias: make(map[string]string),
		pkg:     root,
	}
	// avoid confusing aliases by "reserving" the root package's name as an alias
	imports.byAlias[root.Name] = ""

	infos := make([]*markers.TypeInfo, 0)
	if err := markers.EachType(ctx.Collector, root, func(info *markers.TypeInfo) {
		infos = append(infos, info)
	}); err != nil {
		root.AddError(err)
		return nil
	}
	outContent := new(bytes.Buffer)
	copyCtx := &copyMethodMaker{
		pkg:         root,
		importsList: imports,
		codeWriter:  &codeWriter{out: outContent},
	}

	needGen := false
	for _, info := range infos {
		if len(info.Markers["ioc:autowire"]) != 0 {
			needGen = true
			break
		}
	}
	if !needGen {
		return nil
	}

	copyCtx.GenerateMethodsFor(root, imports, infos)

	outBytes := outContent.Bytes()

	outContent = new(bytes.Buffer)
	writeHeader(root, outContent, root.Name, imports, ctx.HeaderText)
	writeMethods(root, outContent, outBytes)

	outBytes = outContent.Bytes()
	formattedBytes, err := format.Source(outBytes)
	if err != nil {
		root.AddError(err)
		// we still write the invalid source to disk to figure out what went wrong
	} else {
		outBytes = formattedBytes
	}

	return outBytes
}

// writeMethods writes each method to the file, sorted by type name.
func writeMethods(pkg *loader.Package, out io.Writer, outBuffer []byte) {
	_, err := out.Write(outBuffer)
	if err != nil {
		pkg.AddError(err)
	}
}

// writeFormatted outputs the given code, after gofmt-ing it.  If we couldn't gofmt,
// we write the unformatted code for debugging purposes.
func writeOut(ctx *genall.GenerationContext, root *loader.Package, outBytes []byte) {
	outputFile, err := ctx.Open(root, "zz_generated.ioc.go")
	if err != nil {
		root.AddError(err)
		return
	}
	defer outputFile.Close()
	n, err := outputFile.Write(outBytes)
	if err != nil {
		root.AddError(err)
		return
	}
	if n < len(outBytes) {
		root.AddError(io.ErrShortWrite)
	}
}
