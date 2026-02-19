package services
type Pipeline struct{S *Summarizer}
func NewPipeline(s *Summarizer)*Pipeline{return &Pipeline{S:s}}