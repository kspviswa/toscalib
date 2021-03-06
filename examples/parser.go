/*
Copyright 2015 - Olivier Wulveryck

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"bytes"
	"fmt"
	"github.com/awalterschulze/gographviz"
	"github.com/owulveryck/toscalib"
	"github.com/owulveryck/toscalib/toscaexec"
	//"gopkg.in/yaml.v2"
	"log"
	"os"
	"text/template"
)

func main() {
	template, err := template.New("node").Parse(`{{define "NODE"}}<<table border="0" cellspacing="0">
		<tr ><td colspan="2" port="port1" border="1" bgcolor="lightblue">{{.Name}}</td></tr>
		<tr ><td colspan="2" port="port2" border="1">{{.Target}}</td></tr>
		<tr>
			<td port="port2" border="1">{{.Engine}}</td>
			<td port="port8" border="1">{{.Artifact}}</td>
		</tr>
		<tr ><td colspan="2" port="port2" border="1">{{.Args}}</td></tr>
		</table>>{{end}}`)
	var t toscalib.ServiceTemplateDefinition
	err = t.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	//out, err := yaml.Marshal(t)
	//fmt.Println(string(out))
	// Creates a new graph
	g := gographviz.NewGraph()
	g.AddAttr("", "rankdir", "LR")
	g.SetName("G")
	g.SetDir(true)
	e := toscaexec.GeneratePlaybook(t)
	for i, p := range e.Index {
		var out bytes.Buffer
		err = template.ExecuteTemplate(&out, "NODE", p.NodeTemplate)
		g.AddNode("G", fmt.Sprintf("%v", i),
			map[string]string{
				"id":    fmt.Sprintf("\"%v\"", i),
				"label": out.String(),
				//"label": fmt.Sprintf("\"%v|%v\"", p.NodeTemplate.Name, p.OperationName),
				"shape": "\"record\"",
			})
	}
	l := e.AdjacencyMatrix.Dim()
	for r := 0; r < l; r++ {
		for c := 0; c < l; c++ {
			if e.AdjacencyMatrix.At(r, c) == 1 {
				g.AddEdge(fmt.Sprintf("%v", r), fmt.Sprintf("%v", c), true, nil)
			}
		}

	}
	s := g.String()
	fmt.Println(s)
}
