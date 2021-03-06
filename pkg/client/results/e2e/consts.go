/*
Copyright 2018 Heptio Inc.

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

// package e2e defines files and directories found in the e2e plugin results.
package e2e

const (
	// ResultsDirectory is the directory where the results will be found
	ResultsSubdirectory = "e2e/results/global/"

	// LegacyResultsSubdirectory was the directory where e2e results were found before
	// the global folder was added.
	LegacyResultsSubdirectory = "e2e/results/"

	JUnitResultsFile = "junit_01.xml"
)
