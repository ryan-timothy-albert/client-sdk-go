// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CodatFile struct {
	Content  []byte `multipartForm:"content"`
	FileName string `multipartForm:"name=fileName"`
}

func (o *CodatFile) GetContent() []byte {
	if o == nil {
		return []byte{}
	}
	return o.Content
}

func (o *CodatFile) GetFileName() string {
	if o == nil {
		return ""
	}
	return o.FileName
}
