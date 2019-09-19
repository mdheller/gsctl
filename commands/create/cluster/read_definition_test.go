package cluster

import (
	"strconv"
	"testing"

	"github.com/giantswarm/gsctl/commands/types"
	"github.com/google/go-cmp/cmp"
	"github.com/spf13/afero"
	yaml "gopkg.in/yaml.v2"
)

// Test_ReadDefinitionFiles tests the readDefinitionFromFile with all
// YAML files in the testdata directory.
func Test_ReadDefinitionFiles(t *testing.T) {
	basePath := "testdata"
	fs := afero.NewOsFs()
	files, _ := afero.ReadDir(fs, basePath)

	for i, f := range files {
		t.Logf("Case %d, file %s", i, f.Name())
		path := basePath + "/" + f.Name()
		_, err := readDefinitionFromFile(fs, path)
		if err != nil {
			t.Errorf("Unexpected error in case %d, file %s: %s", i, f.Name(), err)
		}
	}
}

// Test_ParseYAMLDefinitionV4 tests parsing v4 YAML definition files.
func Test_ParseYAMLDefinitionV4(t *testing.T) {
	var testCases = []struct {
		inputYAML      []byte
		expectedOutput *types.ClusterDefinitionV4
	}{
		// Minimal YAML.
		{
			[]byte(`owner: myorg`),
			&types.ClusterDefinitionV4{
				Owner: "myorg",
			},
		},
		// More details.
		{
			[]byte(`owner: myorg
name: My cluster
release_version: 1.2.3
availability_zones: 3
scaling:
  min: 3
  max: 5`),
			&types.ClusterDefinitionV4{
				Owner:             "myorg",
				Name:              "My cluster",
				ReleaseVersion:    "1.2.3",
				AvailabilityZones: 3,
				Scaling: types.ScalingDefinition{
					Min: 3,
					Max: 5,
				},
			},
		},
		// KVM worker details.
		{
			[]byte(`owner: myorg
workers:
- memory:
    size_gb: 16.5
  cpu:
    cores: 4
  storage:
    size_gb: 100
- memory:
    size_gb: 32
  cpu:
    cores: 8
  storage:
    size_gb: 50
`),
			&types.ClusterDefinitionV4{
				Owner: "myorg",
				Workers: []types.NodeDefinition{
					types.NodeDefinition{
						Memory:  types.MemoryDefinition{SizeGB: 16.5},
						CPU:     types.CPUDefinition{Cores: 4},
						Storage: types.StorageDefinition{SizeGB: 100},
					},
					types.NodeDefinition{
						Memory:  types.MemoryDefinition{SizeGB: 32},
						CPU:     types.CPUDefinition{Cores: 8},
						Storage: types.StorageDefinition{SizeGB: 50},
					},
				},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			def, err := readDefinitionFromYAML(tc.inputYAML)
			if err != nil {
				t.Errorf("Case %d - Unexpected error %v", i, err)
			}

			if diff := cmp.Diff(tc.expectedOutput, def); diff != "" {
				t.Errorf("Case %d - Resulting definition unequal. (-expected +got):\n%s", i, diff)
			}
		})
	}
}

// Test_ParseYAMLDefinitionV5 tests parsing v5 YAML definition files.
func Test_ParseYAMLDefinitionV5(t *testing.T) {
	var testCases = []struct {
		inputYAML      []byte
		expectedOutput *types.ClusterDefinitionV5
	}{
		// Minimal YAML.
		{
			[]byte(`api_version: v5
owner: myorg`),
			&types.ClusterDefinitionV5{
				Owner: "myorg",
			},
		},
		// More details.
		{
			[]byte(`api_version: v5
owner: myorg
name: My cluster
release_version: 1.2.3
`),
			&types.ClusterDefinitionV5{
				Owner:          "myorg",
				Name:           "My cluster",
				ReleaseVersion: "1.2.3",
			},
		},
		// Node pools.
		{
			[]byte(`api_version: v5
owner: myorg
master:
  availability_zone: my-zone-1a
nodepools:
- name: General purpose
  availability_zones:
    number: 2
- name: Database
  availability_zones:
    zones:
    - my-zone-1a
    - my-zone-1b
    - my-zone-1c
  scaling:
    min: 3
    max: 10
  node_spec:
    aws:
      instance_type: "m5.superlarge"
- name: Batch
`),
			&types.ClusterDefinitionV5{
				Owner:  "myorg",
				Master: &types.MasterDefinition{AvailabilityZone: "my-zone-1a"},
				NodePools: []*types.NodePoolDefinition{
					&types.NodePoolDefinition{
						Name:              "General purpose",
						AvailabilityZones: &types.AvailabilityZonesDefinition{Number: 2},
					},
					&types.NodePoolDefinition{
						Name:              "Database",
						AvailabilityZones: &types.AvailabilityZonesDefinition{Zones: []string{"my-zone-1a", "my-zone-1b", "my-zone-1c"}},
						Scaling:           &types.ScalingDefinition{Min: 3, Max: 10},
						NodeSpec:          &types.NodeSpec{AWS: &types.AWSSpecificDefinition{InstanceType: "m5.superlarge"}},
					},
					&types.NodePoolDefinition{
						Name: "Batch",
					},
				},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			def, err := readDefinitionFromYAML(tc.inputYAML)
			if err != nil {
				t.Errorf("Case %d - Unexpected error %v", i, err)
			}

			if diff := cmp.Diff(tc.expectedOutput, def); diff != "" {
				t.Errorf("Case %d - Resulting definition unequal. (-expected +got):\n%s", i, diff)
			}
		})
	}
}

// Test_CreateFromBadYAML01 tests how non-conforming YAML is treated.
func Test_CreateFromBadYAML01(t *testing.T) {
	data := []byte(`o: myorg`)
	def := types.ClusterDefinitionV4{}

	err := yaml.Unmarshal(data, &def)
	if err != nil {
		t.Fatalf("expected error to be empty, got %#v", err)
	}

	if def.Owner != "" {
		t.Fatalf("expected owner to be empty, got %q", def.Owner)
	}
}