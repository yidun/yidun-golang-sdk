package envidence

type MediaAntispamEvidence struct {
	// 融媒体机审证据信息
	Texts       []*MediaTextEvidence       `json:"texts,omitempty"`
	Images      []*MediaImageEvidence      `json:"images,omitempty"`
	Audios      []*MediaAudioEvidence      `json:"audios,omitempty"`
	Audiovideos []*MediaAudioVideoEvidence `json:"audiovideos,omitempty"`
	Files       []*MediaFileEvidence       `json:"files,omitempty"`
}

func (m *MediaAntispamEvidence) GetTexts() []*MediaTextEvidence {
	return m.Texts
}

func (m *MediaAntispamEvidence) SetTexts(texts []*MediaTextEvidence) {
	m.Texts = texts
}

func (m *MediaAntispamEvidence) GetImages() []*MediaImageEvidence {
	return m.Images
}

func (m *MediaAntispamEvidence) SetImages(images []*MediaImageEvidence) {
	m.Images = images
}

func (m *MediaAntispamEvidence) GetAudios() []*MediaAudioEvidence {
	return m.Audios
}

func (m *MediaAntispamEvidence) SetAudios(audios []*MediaAudioEvidence) {
	m.Audios = audios
}

func (m *MediaAntispamEvidence) GetAudiovideos() []*MediaAudioVideoEvidence {
	return m.Audiovideos
}

func (m *MediaAntispamEvidence) SetAudiovideos(audiovideos []*MediaAudioVideoEvidence) {
	m.Audiovideos = audiovideos
}

func (m *MediaAntispamEvidence) GetFiles() []*MediaFileEvidence {
	return m.Files
}

func (m *MediaAntispamEvidence) SetFiles(files []*MediaFileEvidence) {
	m.Files = files
}
