package schema

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
	"github.com/zarf-dev/zarf/src/pkg/message"
)

//go:embed schemas/*
var schemasFS embed.FS

func GenerateSchemas(baseDir string) error {
	valuesFiles, err := findFiles(baseDir, "values.yaml")
	if err != nil {
		return err
	}

	for _, valuesFile := range valuesFiles {
		outputFile := filepath.Join(filepath.Dir(valuesFile), "values.schema.json")
		message.Infof("Generating schema for %s...\n", valuesFile)

		// Convert YAML to JSON
		jsonData, err := yamlToJSON(valuesFile)
		if err != nil {
			return fmt.Errorf("failed to convert YAML to JSON for %s: %w", valuesFile, err)
		}

		// Generate schema
		baseSchema, err := generateBaseSchema(jsonData)
		if err != nil {
			return fmt.Errorf("failed to generate schema for %s: %w", valuesFile, err)
		}

		// Add custom schema
		customSchema, err := loadCustomSchema()
		if err != nil {
			return err
		}
		baseSchema["properties"].(map[string]interface{})["additionalNetworkAllow"] = customSchema

		// Write schema to file
		outputData, err := json.MarshalIndent(baseSchema, "", "  ")
		if err != nil {
			return err
		}
		if err := os.WriteFile(outputFile, outputData, 0644); err != nil {
			return err
		}
		message.Successf("Schema generated at %s\n", outputFile)
	}

	return nil
}

func CheckSchemas(baseDir string) error {
	valuesFiles, err := findFiles(baseDir, "values.yaml")
	if err != nil {
		return err
	}

	differencesFound := false

	for _, valuesFile := range valuesFiles {
		outputFile := filepath.Join(filepath.Dir(valuesFile), "values.schema.json")
		message.Infof("Checking schema for %s...\n", valuesFile)

		// Convert YAML to JSON
		jsonData, err := yamlToJSON(valuesFile)
		if err != nil {
			return fmt.Errorf("failed to convert YAML to JSON for %s: %w", valuesFile, err)
		}

		// Generate schema
		baseSchema, err := generateBaseSchema(jsonData)
		if err != nil {
			return fmt.Errorf("failed to generate schema for %s: %w", valuesFile, err)
		}

		// Add custom schema
		customSchema, err := loadCustomSchema()
		if err != nil {
			return err
		}
		baseSchema["properties"].(map[string]interface{})["additionalNetworkAllow"] = customSchema

		// Read existing schema
		if _, err := os.Stat(outputFile); errors.Is(err, os.ErrNotExist) {
			message.Warnf("Existing schema not found at %s\n", outputFile)
			differencesFound = true
			continue
		}
		existingSchema, err := os.ReadFile(outputFile)
		if err != nil {
			return err
		}

		// Compare schemas
		generatedSchema, err := json.MarshalIndent(baseSchema, "", "  ")
		if err != nil {
			return err
		}
		if !strings.EqualFold(string(existingSchema), string(generatedSchema)) {
			message.Warnf("Schemas do not match for %s.\nDifferences:\n", valuesFile)
			fmt.Println(diffSchemas(string(existingSchema), string(generatedSchema)))
			differencesFound = true
		}
	}

	if differencesFound {
		message.Warn("Schema differences found.")
		return nil
	}

	message.Success("All schemas match.")
	return nil
}

func diffSchemas(existing, generated string) string {
	// Parse existing schema into a map
	var existingMap map[string]interface{}
	if err := json.Unmarshal([]byte(existing), &existingMap); err != nil {
		return fmt.Sprintf("Error parsing existing schema: %v", err)
	}

	// Parse generated schema into a map
	var generatedMap map[string]interface{}
	if err := json.Unmarshal([]byte(generated), &generatedMap); err != nil {
		return fmt.Sprintf("Error parsing generated schema: %v", err)
	}

	// Create a new JSON differ
	differ := gojsondiff.New()
	diff := differ.CompareObjects(existingMap, generatedMap)

	// Check if differences exist
	if !diff.Modified() {
		return "No differences found."
	}

	// Format differences as ASCII
	formatter := formatter.NewAsciiFormatter(existingMap, formatter.AsciiFormatterConfig{})
	difference, err := formatter.Format(diff)
	if err != nil {
		return fmt.Sprintf("Error formatting differences: %v", err)
	}

	return difference
}

func findFiles(baseDir, pattern string) ([]string, error) {
	var files []string
	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Base(path) == pattern {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func yamlToJSON(filePath string) (map[string]interface{}, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var yamlData map[string]interface{}
	if err := yaml.Unmarshal(data, &yamlData); err != nil {
		return nil, err
	}

	return yamlData, nil
}

func generateBaseSchema(data map[string]interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{
		"type":                 "object",
		"properties":           mapToSchema(data),
		"additionalProperties": false,
	}, nil
}

func mapToSchema(data map[string]interface{}) map[string]interface{} {
	properties := make(map[string]interface{})
	for key, value := range data {
		switch v := value.(type) {
		case map[string]interface{}:
			properties[key] = map[string]interface{}{
				"type":                 "object",
				"properties":           mapToSchema(v),
				"additionalProperties": false,
			}
		case []interface{}:
			// Handle arrays
			items := map[string]interface{}{}
			if len(v) > 0 {
				// Infer schema from the first item
				items = mapToSchema(map[string]interface{}{"item": v[0]})["item"].(map[string]interface{})
			}
			properties[key] = map[string]interface{}{
				"type":  "array",
				"items": items,
			}
		case string:
			properties[key] = map[string]interface{}{
				"type": "string",
			}
		case float64, int, uint64:
			// All numeric values are mapped to "type": "number"
			properties[key] = map[string]interface{}{
				"type": "number",
			}
		case bool:
			properties[key] = map[string]interface{}{
				"type": "boolean",
			}
		case nil:
			properties[key] = map[string]interface{}{
				"type": "null",
			}
		default:
			// Ensure that raw values are not left as-is
			properties[key] = map[string]interface{}{
				"type": fmt.Sprintf("%T", value), // Fallback to type detection
			}
		}
	}
	return properties
}

// LoadCustomSchema loads the `custom_schema.json` file from the embedded filesystem.
func loadCustomSchema() (map[string]interface{}, error) {
	// The path within the embedded filesystem
	customSchemaPath := "schemas/custom_schema.json"

	// Read the file using the embedded filesystem
	data, err := schemasFS.ReadFile(customSchemaPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load custom schema: %w", err)
	}

	// Parse the JSON data
	var schema map[string]interface{}
	if err := json.Unmarshal(data, &schema); err != nil {
		return nil, fmt.Errorf("failed to parse custom schema: %w", err)
	}

	return schema, nil
}
