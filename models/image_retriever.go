package models

type ImageStoreService struct {
	ImageNounPathMap map[string][]string
}

func NewImageStoreService() *ImageStoreService {
	return &ImageStoreService{
		ImageNounPathMap: make(map[string][]string),
	}
}

func (s *ImageStoreService) GetImageNounPathMap() map[string][]string {
	return s.ImageNounPathMap
}

func (s *ImageStoreService) AddPathToNoun(noun, path string) {
	s.ImageNounPathMap[noun] = append(s.ImageNounPathMap[noun], path)
}

func (s *ImageStoreService) GetPathsForNoun(noun string) []string {
	return s.ImageNounPathMap[noun]
}
