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
package toscalib

// TopologyTemplateType as described in appendix A 8
// This section defines the topology template of a cloud application. The main ingredients of the topology template are node templates representing components of the application and relationship templates representing links between the components. These elements are defined in the nested node_templates section and the nested relationship_templates sections, respectively.  Furthermore, a topology template allows for defining input parameters, output parameters as well as grouping of node templates.
type TopologyTemplateType struct {
	Inputs        map[string]PropertyDefinition `yaml:"inputs,omitempty" json:"inputs,omitempty"`
	NodeTemplates map[string]NodeTemplate       `yaml:"node_templates" json:"node_templates"`
	Outputs       map[string]Output             `yaml:"outputs,omitempty" json:"outputs,omitempty"`
}
