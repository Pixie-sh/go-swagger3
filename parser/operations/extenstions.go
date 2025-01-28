package operations

import (
	"encoding/json"
	"fmt"
	oas "github.com/pixie-sh/go-swagger3/openApi3Schema"
	"regexp"
	"strings"
)

func (p *parser) parseExtensionComment(operation *oas.OperationObject, comment string) error {
	sourceString := strings.TrimSpace(comment[len("@Extension"):])

	// /path extension-name {json_object}
	// /path x-lambda {"function":"my-lambda"}
	re := regexp.MustCompile(`([\w\.\/\-{}]+)[^\[]([-.\w]+)[\s]+(\{.*?\})*`)
	matches := re.FindStringSubmatch(sourceString)
	if len(matches) != 4 {
		return fmt.Errorf("Can not parse extention comment \"%s\", skipped", comment)
	}

	//for _, pathItemObject := range p.OpenAPI.Paths {
	//	if pathItemObject.Ref === operation.
	//}

	_, ok := p.OpenAPI.Paths[matches[1]]
	if !ok {
		p.OpenAPI.Paths[matches[1]] = &oas.PathItemObject{UnderlyingExtensions: make(map[string]interface{})}
	}

	// json -> matches[3]
	var raw map[string]interface{}
	err := json.Unmarshal([]byte(matches[3]), &raw)
	if err != nil {
		return err
	}

	// where to store -> operation[matches[1]]
	if p.OpenAPI.Paths[matches[1]].UnderlyingExtensions == nil {
		p.OpenAPI.Paths[matches[1]].UnderlyingExtensions = make(map[string]interface{})
	}

	p.OpenAPI.Paths[matches[1]].UnderlyingExtensions[matches[2]] = raw

	return nil
}
