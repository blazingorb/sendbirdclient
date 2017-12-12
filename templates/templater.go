package templates

import (
	"bytes"
	"html/template"
)

var globalEngine = createEngine()

func GetUsersTemplate(templateData interface{}, key templateKeysUsers) (string, error) {
	return globalEngine.execute(globalEngine.templateUsers, string(key), templateData)
}

func GetOpenChannelTemplate(templateData interface{}, key templateKeysOpenChannels) (string, error) {
	return globalEngine.execute(globalEngine.templateOpenChannels, string(key), templateData)
}

func GetGroupChannelTemplate(templateData interface{}, key templateKeysGroupChannels) (string, error) {
	return globalEngine.execute(globalEngine.templateGroupChannels, string(key), templateData)
}

func GetChannelMetadataTemplate(templateData interface{}, key templateKeysChannelMetadata) (string, error) {
	return globalEngine.execute(globalEngine.templateChannelMetadata, string(key), templateData)
}

func GetMessagesTemplate(templateData interface{}, key templateKeysMessages) (string, error) {
	return globalEngine.execute(globalEngine.templateMessages, string(key), templateData)
}

type engine struct {
	templateUsers           *template.Template
	templateOpenChannels    *template.Template
	templateGroupChannels   *template.Template
	templateChannelMetadata *template.Template
	templateMessages        *template.Template
}

func createEngine() *engine {
	e := &engine{}
	e.templateUsers = template.Must(template.New("Users").Parse(SendbirdURLUsersTemplate))
	e.templateOpenChannels = template.Must(template.New("OpenChannels").Parse(SendbirdURLOpenChannelsTemplate))
	e.templateGroupChannels = template.Must(template.New("GroupChannels").Parse(SendbirdURLGroupChannelsTemplate))
	e.templateChannelMetadata = template.Must(template.New("ChannelMetadata").Parse(SendbirdURLChannelMetadataTemplate))
	e.templateMessages = template.Must(template.New("Messages").Parse(SendbirdURLMessagesTemplate))

	return e
}

func (e *engine) execute(t *template.Template, key string, templateData interface{}) (string, error) {
	var buffer bytes.Buffer
	err := t.ExecuteTemplate(&buffer, key, templateData)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
